package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateIncentivePlan{}, "incentives/CreateIncentivePlan", nil)
	cdc.RegisterConcrete(&MsgUpdateIncentivePlan{}, "incentives/UpdateIncentivePlan", nil)
	cdc.RegisterConcrete(&MsgDeleteIncentivePlan{}, "incentives/DeleteIncentivePlan", nil)
	cdc.RegisterConcrete(&MsgCreateUserStake{}, "incentives/CreateUserStake", nil)
	cdc.RegisterConcrete(&MsgUpdateUserStake{}, "incentives/UpdateUserStake", nil)
	cdc.RegisterConcrete(&MsgDeleteUserStake{}, "incentives/DeleteUserStake", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateIncentivePlan{},
		&MsgUpdateIncentivePlan{},
		&MsgDeleteIncentivePlan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUserStake{},
		&MsgUpdateUserStake{},
		&MsgDeleteUserStake{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
