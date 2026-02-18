package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WorkStatus represents the status of a work unit
type WorkStatus string

const (
	WorkStatusPending    WorkStatus = "pending"
	WorkStatusValidating WorkStatus = "validating"
	WorkStatusValidated  WorkStatus = "validated"
	WorkStatusRejected   WorkStatus = "rejected"
)

// WorkType represents the type of work being validated
type WorkType string

const (
	WorkTypeCrypto      WorkType = "crypto"
	WorkTypeSupplyChain WorkType = "supply_chain"
	WorkTypeMLData      WorkType = "ml_data"
)

// WorkUnit represents a single unit of work to be validated
type WorkUnit struct {
	// ID is a unique identifier for this work unit
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`

	// Type is the type of work (crypto, supply_chain, ml_data)
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`

	// Data is the raw data that needs to be validated
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`

	// SubmittedAt is the block height when the work was submitted
	SubmittedAt int64 `protobuf:"varint,4,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`

	// ValidatedAt is the block height when the work was validated
	ValidatedAt int64 `protobuf:"varint,5,opt,name=validated_at,json=validatedAt,proto3" json:"validated_at,omitempty"`

	// Validator is the address of the validator that handled this work
	Validator string `protobuf:"bytes,6,opt,name=validator,proto3" json:"validator,omitempty"`

	// Status is the current status of the work
	Status string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`

	// Confidence is the validator's confidence in the result (0-100)
	Confidence uint32 `protobuf:"varint,8,opt,name=confidence,proto3" json:"confidence,omitempty"`

	// Proof is optional proof of validation
	Proof string `protobuf:"bytes,9,opt,name=proof,proto3" json:"proof,omitempty"`
}

// ValidateBasic performs basic validation of a WorkUnit
func (w WorkUnit) ValidateBasic() error {
	if w.ID == "" {
		return fmt.Errorf("work unit ID cannot be empty")
	}

	if w.Type != string(WorkTypeCrypto) && w.Type != string(WorkTypeSupplyChain) && w.Type != string(WorkTypeMLData) {
		return fmt.Errorf("invalid work type: %s", w.Type)
	}

	if len(w.Data) == 0 {
		return fmt.Errorf("work unit data cannot be empty")
	}

	if w.Confidence > 100 {
		return fmt.Errorf("confidence cannot exceed 100")
	}

	return nil
}

// WorkQueue stores the queue of work units
type WorkQueue struct {
	// PendingWork is a list of work units waiting to be validated
	PendingWork []WorkUnit `protobuf:"bytes,1,rep,name=pending_work,json=pendingWork,proto3" json:"pending_work"`

	// TotalSubmitted is the total number of work units ever submitted
	TotalSubmitted uint64 `protobuf:"varint,2,opt,name=total_submitted,json=totalSubmitted,proto3" json:"total_submitted,omitempty"`

	// TotalValidated is the total number of work units validated
	TotalValidated uint64 `protobuf:"varint,3,opt,name=total_validated,json=totalValidated,proto3" json:"total_validated,omitempty"`

	// TotalRejected is the total number of work units rejected
	TotalRejected uint64 `protobuf:"varint,4,opt,name=total_rejected,json=totalRejected,proto3" json:"total_rejected,omitempty"`
}

// ValidatorStats tracks performance metrics for a validator
type ValidatorStats struct {
	// Address is the validator's address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`

	// TotalWorkValidated is the number of work units validated
	TotalWorkValidated uint64 `protobuf:"varint,2,opt,name=total_work_validated,json=totalWorkValidated,proto3" json:"total_work_validated,omitempty"`

	// TotalWorkRejected is the number of work units rejected
	TotalWorkRejected uint64 `protobuf:"varint,3,opt,name=total_work_rejected,json=totalWorkRejected,proto3" json:"total_work_rejected,omitempty"`

	// Specializations tracks count of work per type
	Specializations map[string]uint64 `protobuf:"bytes,4,rep,name=specializations,proto3" json:"specializations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`

	// AverageConfidence is the average confidence of validations
	AverageConfidence uint32 `protobuf:"varint,5,opt,name=average_confidence,json=averageConfidence,proto3" json:"average_confidence,omitempty"`

	// LastActiveAt is the block height when last active
	LastActiveAt int64 `protobuf:"varint,6,opt,name=last_active_at,json=lastActiveAt,proto3" json:"last_active_at,omitempty"`
}

// IncrementWorkType increments the work count for a specific type
func (vs *ValidatorStats) IncrementWorkType(workType string) {
	if vs.Specializations == nil {
		vs.Specializations = make(map[string]uint64)
	}
	vs.Specializations[workType]++
}

// GetAccuracy returns the accuracy percentage (validated / total)
func (vs *ValidatorStats) GetAccuracy() uint64 {
	total := vs.TotalWorkValidated + vs.TotalWorkRejected
	if total == 0 {
		return 0
	}
	return (vs.TotalWorkValidated * 100) / total
}
