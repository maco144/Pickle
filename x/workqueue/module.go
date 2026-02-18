package workqueue

import (
	"encoding/json"
	"fmt"

	"github.com/cometbft/cometbft/abci/types"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/maco144/pickle/x/workqueue/client/cli"
	"github.com/maco144/pickle/x/workqueue/keeper"
	workqueuetypes "github.com/maco144/pickle/x/workqueue/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the workqueue module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the workqueue module's name.
func (AppModuleBasic) Name() string {
	return workqueuetypes.ModuleName
}

// RegisterLegacyAminoCodec registers the module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// TODO: Register legacy amino codec if needed
}

// RegisterInterfaces registers the module's interface types
func (b AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	workqueuetypes.RegisterInterfaces(registry)
}

// DefaultGenesis returns default genesis state as raw bytes for the workqueue
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(&workqueuetypes.GenesisState{})
}

// ValidateGenesis performs genesis state validation for the workqueue module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState workqueuetypes.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", workqueuetypes.ModuleName, err)
	}
	return nil
}

// GetTxCmd returns the root tx command for the workqueue module.
func (b AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the root query command for the workqueue module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(workqueuetypes.StoreKey)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the workqueue module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	// TODO: Register gRPC Gateway routes if needed
}

// AppModule implements an application module for the workqueue module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates and returns a new workqueue module.
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// IsAppModule implements the appmodule.AppModule interface.
func (AppModule) IsAppModule() {}

// IsOnePerModuleType implements the appmodule.IsOnePerModuleType interface.
func (AppModule) IsOnePerModuleType() {}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	workqueuetypes.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	workqueuetypes.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServerImpl(am.keeper))
}

// InitGenesis performs genesis initialization for the workqueue module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []types.ValidatorUpdate {
	var genState workqueuetypes.GenesisState
	cdc.MustUnmarshalJSON(data, &genState)
	am.keeper.InitGenesis(ctx, &genState)
	return []types.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the workqueue
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(genState)
}

// RegisterInvariants registers the workqueue module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
	// TODO: Register invariants if needed
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }
