package types

const (
	// WorkStatusPending is the status for pending work
	WorkStatusPending = "pending"
	// WorkStatusValidating is the status for validating work
	WorkStatusValidating = "validating"
	// WorkStatusValidated is the status for validated work
	WorkStatusValidated = "validated"
	// WorkStatusRejected is the status for rejected work
	WorkStatusRejected = "rejected"
)

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
