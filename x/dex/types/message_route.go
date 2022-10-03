package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRoute = "route"

var _ sdk.Msg = &MsgRoute{}

func NewMsgRoute(creator string, receiver string, tokenA string, tokenB string, amountIn string, tokenIn string, minOut string) *MsgRoute {
	return &MsgRoute{
		Creator:  creator,
		Receiver: receiver,
		TokenA:   tokenA,
		TokenB:   tokenB,
		AmountIn: amountIn,
		TokenIn:  tokenIn,
		MinOut:   minOut,
	}
}

func (msg *MsgRoute) Route() string {
	return RouterKey
}

func (msg *MsgRoute) Type() string {
	return TypeMsgRoute
}

func (msg *MsgRoute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRoute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRoute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
