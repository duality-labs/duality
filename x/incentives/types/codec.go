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
	cdc.RegisterConcrete(&MsgLockTokens{}, "duality/lockup/lock-tokens", nil)
	cdc.RegisterConcrete(&MsgBeginUnlockingAll{}, "duality/lockup/begin-unlock-tokens", nil)
	cdc.RegisterConcrete(&MsgBeginUnlocking{}, "duality/lockup/begin-unlock-period-lock", nil)
}

// RegisterInterfaces registers interfaces and implementations of the incentives module.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateGauge{},
		&MsgAddToGauge{},
		&MsgLockTokens{},
		&MsgBeginUnlockingAll{},
		&MsgBeginUnlocking{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
