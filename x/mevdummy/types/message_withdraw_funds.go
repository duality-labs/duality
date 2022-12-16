package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawFunds = "withdraw_funds"

var _ sdk.Msg = &MsgWithdrawFunds{}

func NewMsgWithdrawFunds(creator string, amountOut sdk.Int, tokenOut string) *MsgWithdrawFunds {
	return &MsgWithdrawFunds{
		Creator:   creator,
		AmountOut: amountOut,
		TokenOut:  tokenOut,
	}
}

func (msg *MsgWithdrawFunds) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawFunds) Type() string {
	return TypeMsgWithdrawFunds
}

func (msg *MsgWithdrawFunds) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawFunds) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawFunds) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
