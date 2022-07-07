package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSingleWithdraw = "single_withdraw"

var _ sdk.Msg = &MsgSingleWithdraw{}

func NewMsgSingleWithdraw(creator string, token0 string, token1 string, price string, fee uint64, sharesRemoving string, receiver string) *MsgSingleWithdraw {
	return &MsgSingleWithdraw{
		Creator:        creator,
		Token0:         token0,
		Token1:         token1,
		Price:          price,
		Fee:            fee,
		SharesRemoving: sharesRemoving,
		Receiver:       receiver,
	}
}

func (msg *MsgSingleWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgSingleWithdraw) Type() string {
	return TypeMsgSingleWithdraw
}

func (msg *MsgSingleWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSingleWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSingleWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
