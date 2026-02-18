package types

const (
	// ModuleName defines the module name
	ModuleName = "workqueue"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_workqueue"
)

var (
	// KeyPrefixWorkUnit is the prefix for work units
	KeyPrefixWorkUnit = []byte{0x01}

	// KeyPrefixValidatorStats is the prefix for validator statistics
	KeyPrefixValidatorStats = []byte{0x02}

	// KeyWorkQueue stores the work queue state
	KeyWorkQueue = []byte{0x03}
)

// WorkUnitKey returns the key for a work unit
func WorkUnitKey(workID string) []byte {
	return append(KeyPrefixWorkUnit, []byte(workID)...)
}

// ValidatorStatsKey returns the key for validator statistics
func ValidatorStatsKey(validatorAddr string) []byte {
	return append(KeyPrefixValidatorStats, []byte(validatorAddr)...)
}
