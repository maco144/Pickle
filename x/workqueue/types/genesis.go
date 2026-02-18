package types

import "fmt"

// GenesisState defines the workqueue module's genesis state.
type GenesisState struct {
	// InitialWorkUnits are work units to be imported at genesis
	InitialWorkUnits []WorkUnit `protobuf:"bytes,1,rep,name=initial_work_units,json=initialWorkUnits,proto3" json:"initial_work_units"`

	// InitialValidatorStats are validator stats to be imported at genesis
	InitialValidatorStats []ValidatorStats `protobuf:"bytes,2,rep,name=initial_validator_stats,json=initialValidatorStats,proto3" json:"initial_validator_stats"`

	// Params defines the parameters of the module
	Params Params `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

// Params defines the parameters for the workqueue module
type Params struct {
	// MaxWorkDataSize is the maximum size of work data in bytes
	MaxWorkDataSize uint64 `protobuf:"varint,1,opt,name=max_work_data_size,json=maxWorkDataSize,proto3" json:"max_work_data_size,omitempty"`

	// MinConfidence is the minimum confidence required for validation (0-100)
	MinConfidence uint32 `protobuf:"varint,2,opt,name=min_confidence,json=minConfidence,proto3" json:"min_confidence,omitempty"`
}

// DefaultGenesisState returns the default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		InitialWorkUnits:      []WorkUnit{},
		InitialValidatorStats: []ValidatorStats{},
		Params: Params{
			MaxWorkDataSize: 1024 * 1024,       // 1MB default
			MinConfidence:   50,                 // 50% minimum confidence
		},
	}
}

// Validate performs basic validation of genesis state
func (gs GenesisState) Validate() error {
	// Validate work units
	for _, work := range gs.InitialWorkUnits {
		if err := work.ValidateBasic(); err != nil {
			return fmt.Errorf("invalid work unit: %w", err)
		}
	}

	// Validate parameters
	if gs.Params.MaxWorkDataSize == 0 {
		return fmt.Errorf("max work data size must be greater than 0")
	}

	if gs.Params.MinConfidence > 100 {
		return fmt.Errorf("min confidence cannot exceed 100")
	}

	return nil
}
