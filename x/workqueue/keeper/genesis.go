package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maco144/pickle/x/workqueue/types"
)

// InitGenesis initializes the module's state from a genesis state
func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	if genState == nil {
		return
	}

	// Import initial work queue state
	if genState.WorkQueue != nil {
		// Initialize pending work
		for _, work := range genState.WorkQueue.PendingWork {
			workCopy := work
			if err := k.SubmitWork(ctx, workCopy); err != nil {
				panic(err)
			}
		}
	}

	// Import initial validator stats
	for _, stats := range genState.Validators {
		k.SetValidatorStats(ctx, stats)
	}
}

// ExportGenesis exports the module's state to a genesis state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genState := &types.GenesisState{
		WorkQueue:  &types.WorkQueue{},
		Validators: []*types.ValidatorStats{},
	}

	// Export all work units
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixWorkUnit)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var work types.WorkUnit
		k.cdc.MustUnmarshal(iterator.Value(), &work)
		if work.Status == types.WorkStatusPending {
			genState.WorkQueue.PendingWork = append(genState.WorkQueue.PendingWork, &work)
		}
	}

	// Export totals
	genState.WorkQueue.TotalSubmitted = k.GetTotalWorkSubmitted(ctx)
	genState.WorkQueue.TotalValidated = k.GetTotalWorkValidated(ctx)
	genState.WorkQueue.TotalRejected = k.GetTotalWorkRejected(ctx)

	// Export all validator stats
	k.IterateValidators(ctx, func(validator string, stats *types.ValidatorStats) bool {
		genState.Validators = append(genState.Validators, stats)
		return false
	})

	return genState
}
