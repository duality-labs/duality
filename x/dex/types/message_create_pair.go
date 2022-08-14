package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePair = "create_pair"

var _ sdk.Msg = &MsgCreatePair{}

func NewMsgCreatePair(creator string, tokenA string, tokenB string) *MsgCreatePair {
	return &MsgCreatePair{
		Creator: creator,
		TokenA:  tokenA,
		TokenB:  tokenB,
	}
}

func (msg *MsgCreatePair) Route() string {
	return RouterKey
}

func (msg *MsgCreatePair) Type() string {
	return TypeMsgCreatePair
}

func (msg *MsgCreatePair) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePair) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
