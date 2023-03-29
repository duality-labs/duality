package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawFilledLimitOrder = "withdrawl_withdrawn_limit_order"

var _ sdk.Msg = &MsgWithdrawFilledLimitOrder{}

func NewMsgWithdrawFilledLimitOrder(creator string, trancheKey string) *MsgWithdrawFilledLimitOrder {
	return &MsgWithdrawFilledLimitOrder{
		Creator: creator,
		TrancheKey: trancheKey,
	}
}

func (msg *MsgWithdrawFilledLimitOrder) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawFilledLimitOrder) Type() string {
	return TypeMsgWithdrawFilledLimitOrder
}

func (msg *MsgWithdrawFilledLimitOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawFilledLimitOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawFilledLimitOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
