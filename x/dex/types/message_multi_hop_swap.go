package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMultiHopSwap = "multi_hop_swap"

var _ sdk.Msg = &MsgMultiHopSwap{}

func NewMsgMultiHopSwap(creator string, receiever string, hops string, amountIn string, exitLimitPrice string) *MsgMultiHopSwap {
	return &MsgMultiHopSwap{
		Creator:        creator,
		Receiever:      receiever,
		Hops:           hops,
		AmountIn:       amountIn,
		ExitLimitPrice: exitLimitPrice,
	}
}

func (msg *MsgMultiHopSwap) Route() string {
	return RouterKey
}

func (msg *MsgMultiHopSwap) Type() string {
	return TypeMsgMultiHopSwap
}

func (msg *MsgMultiHopSwap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMultiHopSwap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMultiHopSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
