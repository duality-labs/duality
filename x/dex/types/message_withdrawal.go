package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawal = "Withdrawal"

var _ sdk.Msg = &MsgWithdrawal{}

func NewMsgWithdrawal(creator string, tokenA string, tokenB string, sharesToRemove string, priceIndex int64, feeIndex uint64, receiver string) *MsgWithdrawal {
	return &MsgWithdrawal{
		Creator:        creator,
		TokenA:         tokenA,
		TokenB:         tokenB,
		SharesToRemove: sharesToRemove,
		TickIndex:      priceIndex,
		FeeIndex:       feeIndex,
		Receiver:       receiver,
	}
}

func (msg *MsgWithdrawal) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawal) Type() string {
	return TypeMsgWithdrawal
}

func (msg *MsgWithdrawal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
