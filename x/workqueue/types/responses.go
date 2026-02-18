package types

// MsgSubmitWorkResponse is the response for MsgSubmitWork
type MsgSubmitWorkResponse struct {
	// WorkID is the ID of the submitted work
	WorkID string `protobuf:"bytes,1,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
}

// MsgValidateWorkResponse is the response for MsgValidateWork
type MsgValidateWorkResponse struct {
	// Success indicates if the validation was successful
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

// MsgRejectWorkResponse is the response for MsgRejectWork
type MsgRejectWorkResponse struct {
	// Success indicates if the rejection was successful
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}
