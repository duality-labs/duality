package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPlaceLimitOrder = "place_limit_order"

var _ sdk.Msg = &MsgPlaceLimitOrder{}

func NewMsgPlaceLimitOrder(creator string, receiver string, tokenIn string, tokenOut string, tickIndex int64, amountIn sdk.Dec) *MsgPlaceLimitOrder {
	return &MsgPlaceLimitOrder{
		Creator:   creator,
		Receiver:  receiver,
		TokenIn:   tokenIn,
		TokenOut:  tokenOut,
		TickIndex: tickIndex,
		AmountIn:  amountIn,
	}
}

func (msg *MsgPlaceLimitOrder) Route() string {
	return RouterKey
}

func (msg *MsgPlaceLimitOrder) Type() string {
	return TypeMsgPlaceLimitOrder
}

func (msg *MsgPlaceLimitOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPlaceLimitOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPlaceLimitOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
