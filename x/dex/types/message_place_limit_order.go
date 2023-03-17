package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPlaceLimitOrder = "place_limit_order"

var _ sdk.Msg = &MsgPlaceLimitOrder{}

func NewMsgPlaceLimitOrder(creator string, receiver string, tokenA string, tokenB string, tickIndex int64, tokenIn string, amountIn sdk.Int, goodTil *time.Time) *MsgPlaceLimitOrder {
	return &MsgPlaceLimitOrder{
		Creator:        creator,
		Receiver:       receiver,
		TokenA:         tokenA,
		TokenB:         tokenB,
		TickIndex:      tickIndex,
		TokenIn:        tokenIn,
		AmountIn:       amountIn,
		ExpirationTime: goodTil,
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

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	if msg.TokenIn != msg.TokenA && msg.TokenIn != msg.TokenB {
		return ErrInvalidTokenIn
	}

	if msg.AmountIn.LTE(sdk.ZeroInt()) {
		return ErrZeroLimitOrder
	}

	if msg.OrderType.IsGoodTil() && msg.ExpirationTime == nil {
		return ErrGoodTilOrderWithoutExpiration
	}

	if !msg.OrderType.IsGoodTil() && msg.ExpirationTime != nil {
		return ErrExpirationOnWrongOrderType
	}
	return nil
}
