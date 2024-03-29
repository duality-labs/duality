package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPlaceLimitOrder = "place_limit_order"

var _ sdk.Msg = &MsgPlaceLimitOrder{}

func NewMsgPlaceLimitOrder(
	creator,
	receiver,
	tokenIn,
	tokenOut string,
	tickIndex int64,
	amountIn sdk.Int,
	orderType LimitOrderType,
	goodTil *time.Time,
	maxAmountOut *sdk.Int,
) *MsgPlaceLimitOrder {
	return &MsgPlaceLimitOrder{
		Creator:          creator,
		Receiver:         receiver,
		TokenIn:          tokenIn,
		TokenOut:         tokenOut,
		TickIndexInToOut: tickIndex,
		AmountIn:         amountIn,
		OrderType:        orderType,
		ExpirationTime:   goodTil,
		MaxAmountOut:     maxAmountOut,
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

	if msg.AmountIn.LTE(sdk.ZeroInt()) {
		return ErrZeroLimitOrder
	}

	if msg.OrderType.IsGoodTil() && msg.ExpirationTime == nil {
		return ErrGoodTilOrderWithoutExpiration
	}

	if !msg.OrderType.IsGoodTil() && msg.ExpirationTime != nil {
		return ErrExpirationOnWrongOrderType
	}

	if msg.MaxAmountOut != nil {
		if !msg.MaxAmountOut.IsPositive() {
			return ErrZeroMaxAmountOut
		}
		if !msg.OrderType.IsTakerOnly() {
			return ErrInvalidMaxAmountOutForMaker
		}
	}

	return nil
}

func (msg *MsgPlaceLimitOrder) ValidateGoodTilExpiration(blockTime time.Time) error {
	if msg.OrderType.IsGoodTil() && !msg.ExpirationTime.After(blockTime) {
		return sdkerrors.Wrapf(ErrExpirationTimeInPast,
			"Current BlockTime: %s; Provided ExpirationTime: %s",
			blockTime.String(),
			msg.ExpirationTime.String(),
		)
	}

	return nil
}
