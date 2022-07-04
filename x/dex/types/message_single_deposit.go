package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSingleDeposit = "single_deposit"

var _ sdk.Msg = &MsgSingleDeposit{}

func NewMsgSingleDeposit(creator string, token0 string, token1 string, price string, fee uint64, amounts0 uint64, amounts1 uint64, receiver string) *MsgSingleDeposit {
	return &MsgSingleDeposit{
		Creator:  creator,
		Token0:   token0,
		Token1:   token1,
		Price:    price,
		Fee:      fee,
		Amounts0: amounts0,
		Amounts1: amounts1,
		Receiver: receiver,
	}
}

func (msg *MsgSingleDeposit) Route() string {
	return RouterKey
}

func (msg *MsgSingleDeposit) Type() string {
	return TypeMsgSingleDeposit
}

func (msg *MsgSingleDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSingleDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSingleDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
