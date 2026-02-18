package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maco144/pickle/x/workqueue/types"
)

// InitGenesis initializes the module's state from a genesis state
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// Import initial work units
	for _, work := range genState.InitialWorkUnits {
		workCopy := work
		if err := k.SubmitWork(ctx, &workCopy); err != nil {
			panic(err)
		}
	}

	// Import initial validator stats
	for _, stats := range genState.InitialValidatorStats {
		statsCopy := stats
		k.SetValidatorStats(ctx, &statsCopy)
	}

	// Set parameters (if params module integration is added later)
	// For now, parameters are stored in genesis state
}

// ExportGenesis exports the module's state to a genesis state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genState := &types.GenesisState{
		InitialWorkUnits:      []types.WorkUnit{},
		InitialValidatorStats: []types.ValidatorStats{},
		Params: types.Params{
			MaxWorkDataSize: 1024 * 1024,
			MinConfidence:   50,
		},
	}

	// Export all work units
	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(types.KeyPrefixWorkUnit, append(types.KeyPrefixWorkUnit, 0xff))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var work types.WorkUnit
		k.cdc.MustUnmarshal(iterator.Value(), &work)
		genState.InitialWorkUnits = append(genState.InitialWorkUnits, work)
	}

	// Export all validator stats
	iterator = store.Iterator(types.KeyPrefixValidatorStats, append(types.KeyPrefixValidatorStats, 0xff))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var stats types.ValidatorStats
		k.cdc.MustUnmarshal(iterator.Value(), &stats)
		genState.InitialValidatorStats = append(genState.InitialValidatorStats, stats)
	}

	return genState
}
