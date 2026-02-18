package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// QueryWorkRequest is the request for querying a specific work unit
type QueryWorkRequest struct {
	WorkID string `protobuf:"bytes,1,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
}

// QueryWorkResponse is the response for querying a specific work unit
type QueryWorkResponse struct {
	Work *WorkUnit `protobuf:"bytes,1,opt,name=work,proto3" json:"work,omitempty"`
}

// QueryPendingWorkRequest is the request for querying pending work
type QueryPendingWorkRequest struct{}

// QueryPendingWorkResponse is the response for querying pending work
type QueryPendingWorkResponse struct {
	PendingWork []*WorkUnit `protobuf:"bytes,1,rep,name=pending_work,json=pendingWork,proto3" json:"pending_work"`
}

// QueryValidatorStatsRequest is the request for querying validator stats
type QueryValidatorStatsRequest struct {
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

// QueryValidatorStatsResponse is the response for querying validator stats
type QueryValidatorStatsResponse struct {
	Stats *ValidatorStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
}

// QueryTotalStatsRequest is the request for querying total statistics
type QueryTotalStatsRequest struct{}

// QueryTotalStatsResponse is the response for querying total statistics
type QueryTotalStatsResponse struct {
	TotalSubmitted uint64 `protobuf:"varint,1,opt,name=total_submitted,json=totalSubmitted,proto3" json:"total_submitted,omitempty"`
	TotalValidated uint64 `protobuf:"varint,2,opt,name=total_validated,json=totalValidated,proto3" json:"total_validated,omitempty"`
	TotalRejected  uint64 `protobuf:"varint,3,opt,name=total_rejected,json=totalRejected,proto3" json:"total_rejected,omitempty"`
}
