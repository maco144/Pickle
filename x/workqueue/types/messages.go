package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitWork = "submit_work"
const TypeMsgValidateWork = "validate_work"
const TypeMsgRejectWork = "reject_work"

var (
	_ sdk.Msg = &MsgSubmitWork{}
	_ sdk.Msg = &MsgValidateWork{}
	_ sdk.Msg = &MsgRejectWork{}
)

// MsgSubmitWork submits a new work unit for validation
type MsgSubmitWork struct {
	// Submitter is the address submitting the work
	Submitter string `protobuf:"bytes,1,opt,name=submitter,proto3" json:"submitter,omitempty"`

	// WorkType is the type of work (crypto, supply_chain, ml_data)
	WorkType string `protobuf:"bytes,2,opt,name=work_type,json=workType,proto3" json:"work_type,omitempty"`

	// WorkData is the raw data to validate
	WorkData []byte `protobuf:"bytes,3,opt,name=work_data,json=workData,proto3" json:"work_data,omitempty"`

	// WorkID is a unique identifier (optional, generated if not provided)
	WorkID string `protobuf:"bytes,4,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
}

func (msg *MsgSubmitWork) Route() string {
	return RouterKey
}

func (msg *MsgSubmitWork) Type() string {
	return TypeMsgSubmitWork
}

func (msg *MsgSubmitWork) GetSigners() []sdk.AccAddress {
	submitter, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{submitter}
}

func (msg *MsgSubmitWork) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submitter address (%s)", err)
	}

	if msg.WorkType == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "work type cannot be empty")
	}

	if len(msg.WorkData) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "work data cannot be empty")
	}

	return nil
}

// MsgValidateWork submits a validation result for a work unit
type MsgValidateWork struct {
	// Validator is the address of the validator
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`

	// WorkID is the ID of the work unit being validated
	WorkID string `protobuf:"bytes,2,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`

	// Valid indicates if the work unit passed validation
	Valid bool `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`

	// Confidence is the validator's confidence in this result (0-100)
	Confidence uint32 `protobuf:"varint,4,opt,name=confidence,proto3" json:"confidence,omitempty"`

	// Proof is optional proof of validation (e.g., hash of verification)
	Proof string `protobuf:"bytes,5,opt,name=proof,proto3" json:"proof,omitempty"`

	// Reason is an optional explanation if invalid
	Reason string `protobuf:"bytes,6,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (msg *MsgValidateWork) Route() string {
	return RouterKey
}

func (msg *MsgValidateWork) Type() string {
	return TypeMsgValidateWork
}

func (msg *MsgValidateWork) GetSigners() []sdk.AccAddress {
	validator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{validator}
}

func (msg *MsgValidateWork) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	if msg.WorkID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "work ID cannot be empty")
	}

	if msg.Confidence > 100 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "confidence cannot exceed 100")
	}

	return nil
}

// MsgRejectWork explicitly rejects a work unit
type MsgRejectWork struct {
	// Validator is the address of the validator
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`

	// WorkID is the ID of the work unit being rejected
	WorkID string `protobuf:"bytes,2,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`

	// Reason is the reason for rejection
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (msg *MsgRejectWork) Route() string {
	return RouterKey
}

func (msg *MsgRejectWork) Type() string {
	return TypeMsgRejectWork
}

func (msg *MsgRejectWork) GetSigners() []sdk.AccAddress {
	validator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{validator}
}

func (msg *MsgRejectWork) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	if msg.WorkID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "work ID cannot be empty")
	}

	return nil
}
