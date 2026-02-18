package pickle

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	tmdb "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	tmtypes "github.com/cometbft/cometbft/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/maco144/pickle/x/workqueue"
	workqueuekeeper "github.com/maco144/pickle/x/workqueue/keeper"
	workqueuetypes "github.com/maco144/pickle/x/workqueue/types"
)

const (
	Name = "pickle"
)

var (
	// DefaultNodeHome default home directories for the application
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		workqueue.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
	}
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with some additional fields
// to track tendermint commits and module-specific keepers
type App struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyCodec
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	// keepers
	keys map[string]*storetypes.KVStoreKey
	tKeys map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	AccountKeeper      authkeeper.AccountKeeper
	BankKeeper         bankkeeper.Keeper
	ParamsKeeper       paramskeeper.Keeper
	WorkqueueKeeper    workqueuekeeper.Keeper

	// the module manager
	mm *module.Manager

	// the configurator
	configurator module.Configurator
}

// NewApp returns a reference to an initialized Pickle application.
func NewApp(
	logger log.Logger,
	db tmdb.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	// Setup codec
	var cdc codec.Codec
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc = codec.NewProtoCodec(interfaceRegistry)
	legacyAmino := codec.NewLegacyAmino()

	// Register module codecs
	ModuleBasics.RegisterInterfaces(interfaceRegistry)

	// Create base app
	bApp := baseapp.NewBaseApp(Name, logger, db, nil, append(
		[]func(*baseapp.BaseApp){
			baseapp.SetChainID("pickle-1"),
		},
		baseAppOptions...,
	)...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	// Declare store keys
	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		paramstypes.StoreKey,
		workqueuetypes.StoreKey,
	)

	tKeys := sdk.NewTransientStoreKeys(
		paramstypes.TStoreKey,
	)

	memKeys := sdk.NewMemoryStoreKeys(
		workqueuetypes.MemStoreKey,
	)

	// Create app
	app := &App{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          cdc,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tKeys:             tKeys,
		memKeys:           memKeys,
	}

	// Initialize keepers
	app.ParamsKeeper = paramskeeper.NewKeeper(
		cdc,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tKeys[paramstypes.TStoreKey],
	)

	app.AccountKeeper = authkeeper.NewAccountKeeper(
		cdc,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		maccPerms,
		sdk.Bech32MainPrefix,
		"pickle",
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		cdc,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		map[string]bool{},
		"pickle",
	)

	app.WorkqueueKeeper = workqueuekeeper.NewKeeper(
		cdc,
		keys[workqueuetypes.StoreKey],
		memKeys[workqueuetypes.MemStoreKey],
	)

	// Create module manager
	app.mm = module.NewManager(
		auth.NewAppModule(cdc, app.AccountKeeper, nil, nil),
		bank.NewAppModule(cdc, app.BankKeeper, app.AccountKeeper, nil),
		params.NewAppModule(app.ParamsKeeper),
		workqueue.NewAppModule(cdc, app.WorkqueueKeeper),
	)

	// Set module order
	app.mm.SetOrderBeginBlockers()

	app.mm.SetOrderEndBlockers()

	app.mm.SetOrderInitGenesis(
		authtypes.ModuleName,
		banktypes.ModuleName,
		paramstypes.ModuleName,
		workqueuetypes.ModuleName,
	)

	// Register upgrade handlers
	app.registerUpgradeHandlers()

	// Mount stores
	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)
	app.MountMemoryStores(memKeys)

	// Initialize the app
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// Set antehandler
	app.SetAnteHandler(nil)

	// Create the configurator
	app.configurator = module.NewConfigurator(cdc, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// Load latest version
	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			panic(err)
		}
	}

	// Initialize and seal capabilities
	app.Seal()

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return Name }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) error {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) ([]abci.ValidatorUpdate, error) {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState map[string]json.RawMessage

	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LegacyAmino returns SimApp's amino codec.
func (app *App) LegacyAmino() *codec.LegacyCodec {
	return app.legacyAmino
}

// AppCodec returns Pickle's app codec.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Pickle's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// Configurator returns the configurator
func (app *App) Configurator() module.Configurator {
	return app.configurator
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	// This is handled by the SDK's built-in routing
}

// GetKey returns the KVStoreKey for the provided store key.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tKeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided store key.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// registerUpgradeHandlers registers upgrade handlers for the app
func (app *App) registerUpgradeHandlers() {
	// Upgrade handlers would go here
}

// RegisterTxService registers the tx service for the app
func (app *App) RegisterTxService(clientCtx client.Context) {
	// Tx service registration
}

// RegisterTendermintService registers the tendermint service for the app
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	// Tendermint service registration
}

// ExportAppStateAndValidators exports the state of the application for a genesis file.
func (app *App) ExportAppStateAndValidators(
	forZeroHeight bool,
	jailAllowedAddrs []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	ctx := app.NewContext(true)

	// Export genesis state
	genesisState := app.mm.ExportGenesisForModules(
		ctx,
		app.appCodec,
		[]string{},
	)

	appState, err := json.MarshalIndent(genesisState, "", "  ")
	if err != nil {
		return nil, nil, err
	}

	return appState, []tmtypes.GenesisValidator{}, nil
}
