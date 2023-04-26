package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

// RegisterCodec registers the necessary x/incentives interfaces and concrete types on the provided
// LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGauge{}, "duality/incentives/create-gauge", nil)
	cdc.RegisterConcrete(&MsgAddToGauge{}, "duality/incentives/add-to-gauge", nil)
	cdc.RegisterConcrete(&MsgStake{}, "duality/stakeup/stake-tokens", nil)
	cdc.RegisterConcrete(&MsgUnstake{}, "duality/stakeup/begin-unstake-period-stake", nil)
}

// RegisterInterfaces registers interfaces and implementations of the incentives module.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateGauge{},
		&MsgAddToGauge{},
		&MsgStake{},
		&MsgUnstake{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
