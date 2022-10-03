package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawlFilledLimitOrder = "withdrawl_filled_limit_order"

var _ sdk.Msg = &MsgWithdrawlFilledLimitOrder{}

func NewMsgWithdrawlFilledLimitOrder(creator string, tokenA string, tokenB string, tickIndex int64, keyToken string, key uint64) *MsgWithdrawlFilledLimitOrder {
	return &MsgWithdrawlFilledLimitOrder{
		Creator:   creator,
		TokenA:    tokenA,
		TokenB:    tokenB,
		TickIndex: tickIndex,
		KeyToken:  keyToken,
		Key:       key,
	}
}

func (msg *MsgWithdrawlFilledLimitOrder) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawlFilledLimitOrder) Type() string {
	return TypeMsgWithdrawlFilledLimitOrder
}

func (msg *MsgWithdrawlFilledLimitOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawlFilledLimitOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawlFilledLimitOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
