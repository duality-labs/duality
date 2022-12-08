package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawl = "withdrawl"

var _ sdk.Msg = &MsgWithdrawl{}

func NewMsgWithdrawl(creator string, receiver string, tokenA string, tokenB string, sharesToRemove []sdk.Int, tickIndexes []int64, feeIndexes []uint64) *MsgWithdrawl {
	return &MsgWithdrawl{
		Creator:        creator,
		Receiver:       receiver,
		TokenA:         tokenA,
		TokenB:         tokenB,
		SharesToRemove: sharesToRemove,
		TickIndexes:    tickIndexes,
		FeeIndexes:     feeIndexes,
	}
}

func (msg *MsgWithdrawl) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawl) Type() string {
	return TypeMsgWithdrawl
}

func (msg *MsgWithdrawl) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawl) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawl) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
