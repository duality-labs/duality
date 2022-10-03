package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelLimitOrder = "cancel_limit_order"

var _ sdk.Msg = &MsgCancelLimitOrder{}

func NewMsgCancelLimitOrder(creator string, tokenA string, tokenB string, tickIndex int64, keyToken string, key uint64) *MsgCancelLimitOrder {
	return &MsgCancelLimitOrder{
		Creator:   creator,
		TokenA:    tokenA,
		TokenB:    tokenB,
		TickIndex: tickIndex,
		KeyToken:  keyToken,
		Key:       key,
	}
}

func (msg *MsgCancelLimitOrder) Route() string {
	return RouterKey
}

func (msg *MsgCancelLimitOrder) Type() string {
	return TypeMsgCancelLimitOrder
}

func (msg *MsgCancelLimitOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelLimitOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelLimitOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
