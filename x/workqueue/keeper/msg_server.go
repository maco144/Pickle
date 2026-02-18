package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maco144/pickle/x/workqueue/types"
)

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// SubmitWork implements the MsgServer.SubmitWork method
func (ms msgServer) SubmitWork(goCtx context.Context, msg *types.MsgSubmitWork) (*types.MsgSubmitWorkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create work unit from message
	work := &types.WorkUnit{
		Id:   msg.WorkId,
		Type: msg.WorkType,
		Data: msg.WorkData,
	}

	// Submit the work
	if err := ms.Keeper.SubmitWork(ctx, work); err != nil {
		return nil, err
	}

	return &types.MsgSubmitWorkResponse{
		WorkId: work.Id,
	}, nil
}

// ValidateWork implements the MsgServer.ValidateWork method
func (ms msgServer) ValidateWork(goCtx context.Context, msg *types.MsgValidateWork) (*types.MsgValidateWorkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the work
	if err := ms.Keeper.ValidateWork(ctx, msg.WorkId, msg.Validator, msg.Valid, msg.Confidence, msg.Proof); err != nil {
		return nil, err
	}

	return &types.MsgValidateWorkResponse{}, nil
}

// RejectWork implements the MsgServer.RejectWork method
func (ms msgServer) RejectWork(goCtx context.Context, msg *types.MsgRejectWork) (*types.MsgRejectWorkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Reject the work
	if err := ms.Keeper.RejectWork(ctx, msg.WorkId, msg.Validator, msg.Reason); err != nil {
		return nil, err
	}

	return &types.MsgRejectWorkResponse{}, nil
}
