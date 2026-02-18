package types

const (
	EventTypeWorkSubmitted = "work_submitted"
	EventTypeWorkValidated = "work_validated"
	EventTypeWorkRejected  = "work_rejected"

	AttributeKeyWorkID     = "work_id"
	AttributeKeyWorkType   = "work_type"
	AttributeKeyValidator  = "validator"
	AttributeKeyStatus     = "status"
	AttributeKeyConfidence = "confidence"
	AttributeKeySubmittedAt = "submitted_at"
	AttributeKeyReason     = "reason"
)
