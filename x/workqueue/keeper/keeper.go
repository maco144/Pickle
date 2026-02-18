package keeper

import (
	"crypto/sha256"
	"fmt"

	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/maco144/pickle/x/workqueue/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey *storetypes.KVStoreKey
		memKey   *storetypes.MemoryStoreKey
	}
)

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, storeKey *storetypes.KVStoreKey, memKey *storetypes.MemoryStoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SubmitWork submits a new work unit for validation
func (k Keeper) SubmitWork(ctx sdk.Context, workUnit *types.WorkUnit) error {
	// Generate ID if not provided (use block height + hash)
	if workUnit.Id == "" {
		hash := sha256.Sum256(workUnit.Data)
		workUnit.Id = fmt.Sprintf("%d-%x", ctx.BlockHeight(), hash[:8])
	}

	// Set submission block height
	workUnit.SubmittedAt = ctx.BlockHeight()
	workUnit.Status = types.WorkStatusPending

	// Store the work unit
	k.SetWork(ctx, workUnit)

	// Update total submitted count
	k.IncrementTotalSubmitted(ctx)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWorkSubmitted,
			sdk.NewAttribute(types.AttributeKeyWorkID, workUnit.Id),
			sdk.NewAttribute(types.AttributeKeyWorkType, workUnit.Type),
			sdk.NewAttribute(types.AttributeKeySubmittedAt, fmt.Sprintf("%d", workUnit.SubmittedAt)),
		),
	)

	return nil
}

// GetWork retrieves a work unit by ID
func (k Keeper) GetWork(ctx sdk.Context, workID string) (*types.WorkUnit, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.WorkUnitKey(workID))
	if bz == nil {
		return nil, false
	}

	var work types.WorkUnit
	k.cdc.MustUnmarshal(bz, &work)
	return &work, true
}

// SetWork stores a work unit
func (k Keeper) SetWork(ctx sdk.Context, work *types.WorkUnit) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(work)
	store.Set(types.WorkUnitKey(work.Id), bz)
}

// GetValidatorStats retrieves statistics for a validator
func (k Keeper) GetValidatorStats(ctx sdk.Context, validatorAddr string) (*types.ValidatorStats, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ValidatorStatsKey(validatorAddr))
	if bz == nil {
		return nil, false
	}

	var stats types.ValidatorStats
	k.cdc.MustUnmarshal(bz, &stats)
	return &stats, true
}

// SetValidatorStats stores validator statistics
func (k Keeper) SetValidatorStats(ctx sdk.Context, stats *types.ValidatorStats) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(stats)
	store.Set(types.ValidatorStatsKey(stats.Address), bz)
}

// GetPendingWork returns a list of pending work units
func (k Keeper) GetPendingWork(ctx sdk.Context) []*types.WorkUnit {
	var pending []*types.WorkUnit
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixWorkUnit)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var work types.WorkUnit
		k.cdc.MustUnmarshal(iterator.Value(), &work)
		if work.Status == types.WorkStatusPending {
			pending = append(pending, &work)
		}
	}

	return pending
}

// ValidateWork marks a work unit as validated
func (k Keeper) ValidateWork(ctx sdk.Context, workID string, validatorAddr string, valid bool, confidence uint32, proof string) error {
	// Get the work unit
	work, found := k.GetWork(ctx, workID)
	if !found {
		return fmt.Errorf("work unit not found: %s", workID)
	}

	// Validate confidence
	if confidence > 100 {
		return fmt.Errorf("confidence cannot exceed 100")
	}

	// Update work unit
	work.Validator = validatorAddr
	work.ValidatedAt = ctx.BlockHeight()
	work.Confidence = confidence
	work.Proof = proof

	if valid {
		work.Status = types.WorkStatusValidated
		k.IncrementTotalValidated(ctx)
	} else {
		work.Status = types.WorkStatusRejected
		k.IncrementTotalRejected(ctx)
	}

	// Store updated work unit
	k.SetWork(ctx, work)

	// Update validator stats
	stats, _ := k.GetValidatorStats(ctx, validatorAddr)
	if stats == nil {
		stats = &types.ValidatorStats{
			Address:         validatorAddr,
			Specializations: make(map[string]uint64),
		}
	}

	if valid {
		stats.TotalWorkValidated++
	} else {
		stats.TotalWorkRejected++
	}

	// Update specialization tracking
	stats.IncrementWorkType(work.Type)

	// Update average confidence
	total := stats.TotalWorkValidated + stats.TotalWorkRejected
	if total > 0 {
		stats.AverageConfidence = uint32((uint64(stats.AverageConfidence)*(total-1) + uint64(confidence)) / total)
	} else {
		stats.AverageConfidence = confidence
	}

	stats.LastActiveAt = ctx.BlockHeight()

	// Store updated stats
	k.SetValidatorStats(ctx, stats)

	// Emit event
	statusStr := "validated"
	if !valid {
		statusStr = "rejected"
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWorkValidated,
			sdk.NewAttribute(types.AttributeKeyWorkID, workID),
			sdk.NewAttribute(types.AttributeKeyValidator, validatorAddr),
			sdk.NewAttribute(types.AttributeKeyStatus, statusStr),
			sdk.NewAttribute(types.AttributeKeyConfidence, fmt.Sprintf("%d", confidence)),
		),
	)

	return nil
}

// RejectWork marks a work unit as rejected
func (k Keeper) RejectWork(ctx sdk.Context, workID string, validatorAddr string, reason string) error {
	// Get the work unit
	work, found := k.GetWork(ctx, workID)
	if !found {
		return fmt.Errorf("work unit not found: %s", workID)
	}

	// Update work unit
	work.Validator = validatorAddr
	work.ValidatedAt = ctx.BlockHeight()
	work.Status = types.WorkStatusRejected
	work.Proof = reason

	// Store updated work unit
	k.SetWork(ctx, work)

	// Update validator stats
	stats, _ := k.GetValidatorStats(ctx, validatorAddr)
	if stats == nil {
		stats = &types.ValidatorStats{
			Address:         validatorAddr,
			Specializations: make(map[string]uint64),
		}
	}

	stats.TotalWorkRejected++
	stats.IncrementWorkType(work.Type)
	stats.LastActiveAt = ctx.BlockHeight()

	// Store updated stats
	k.SetValidatorStats(ctx, stats)

	// Increment total rejected
	k.IncrementTotalRejected(ctx)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWorkRejected,
			sdk.NewAttribute(types.AttributeKeyWorkID, workID),
			sdk.NewAttribute(types.AttributeKeyValidator, validatorAddr),
			sdk.NewAttribute(types.AttributeKeyReason, reason),
		),
	)

	return nil
}

// GetTotalWorkValidated returns the total number of validated work units
func (k Keeper) GetTotalWorkValidated(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("total_validated"))
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

// GetTotalWorkRejected returns the total number of rejected work units
func (k Keeper) GetTotalWorkRejected(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("total_rejected"))
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

// GetTotalWorkSubmitted returns the total number of submitted work units
func (k Keeper) GetTotalWorkSubmitted(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("total_submitted"))
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

// IncrementTotalValidated increments the total validated count
func (k Keeper) IncrementTotalValidated(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	total := k.GetTotalWorkValidated(ctx) + 1
	store.Set([]byte("total_validated"), sdk.Uint64ToBigEndian(total))
}

// IncrementTotalRejected increments the total rejected count
func (k Keeper) IncrementTotalRejected(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	total := k.GetTotalWorkRejected(ctx) + 1
	store.Set([]byte("total_rejected"), sdk.Uint64ToBigEndian(total))
}

// IncrementTotalSubmitted increments the total submitted count
func (k Keeper) IncrementTotalSubmitted(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	total := k.GetTotalWorkSubmitted(ctx) + 1
	store.Set([]byte("total_submitted"), sdk.Uint64ToBigEndian(total))
}

// IterateValidators iterates over all validators with stats
func (k Keeper) IterateValidators(ctx sdk.Context, cb func(validator string, stats *types.ValidatorStats) bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixValidatorStats)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var stats types.ValidatorStats
		k.cdc.MustUnmarshal(iterator.Value(), &stats)
		if !cb(stats.Address, &stats) {
			break
		}
	}
}
