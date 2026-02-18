package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/maco144/pickle/x/workqueue/types"
)

type queryServer struct {
	Keeper
	types.UnimplementedQueryServer
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

// Work implements the Query.Work method
func (qs queryServer) Work(goCtx context.Context, req *types.QueryWorkRequest) (*types.QueryWorkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	work, found := qs.Keeper.GetWork(ctx, req.WorkId)
	if !found {
		return nil, status.Error(codes.NotFound, "work not found")
	}

	return &types.QueryWorkResponse{
		Work: work,
	}, nil
}

// PendingWork implements the Query.PendingWork method
func (qs queryServer) PendingWork(goCtx context.Context, req *types.QueryPendingWorkRequest) (*types.QueryPendingWorkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pending := qs.Keeper.GetPendingWork(ctx)

	return &types.QueryPendingWorkResponse{
		PendingWork: pending,
	}, nil
}

// ValidatorStats implements the Query.ValidatorStats method
func (qs queryServer) ValidatorStats(goCtx context.Context, req *types.QueryValidatorStatsRequest) (*types.QueryValidatorStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	stats, found := qs.Keeper.GetValidatorStats(ctx, req.ValidatorAddress)
	if !found {
		return nil, status.Error(codes.NotFound, "validator stats not found")
	}

	return &types.QueryValidatorStatsResponse{
		Stats: stats,
	}, nil
}

// TotalStats implements the Query.TotalStats method
func (qs queryServer) TotalStats(goCtx context.Context, req *types.QueryTotalStatsRequest) (*types.QueryTotalStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryTotalStatsResponse{
		TotalSubmitted: qs.Keeper.GetTotalWorkSubmitted(ctx),
		TotalValidated: qs.Keeper.GetTotalWorkValidated(ctx),
		TotalRejected:  qs.Keeper.GetTotalWorkRejected(ctx),
	}, nil
}
