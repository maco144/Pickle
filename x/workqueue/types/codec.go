package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyCodec) {
	cdc.RegisterConcrete(&MsgSubmitWork{}, "workqueue/SubmitWork", nil)
	cdc.RegisterConcrete(&MsgValidateWork{}, "workqueue/ValidateWork", nil)
	cdc.RegisterConcrete(&MsgRejectWork{}, "workqueue/RejectWork", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitWork{},
		&MsgValidateWork{},
		&MsgRejectWork{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(Amino)
	Amino.Seal()
}
