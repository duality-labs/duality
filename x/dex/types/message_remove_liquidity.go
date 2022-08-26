package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveLiquidity = "remove_liquidity"

var _ sdk.Msg = &MsgRemoveLiquidity{}

func NewMsgRemoveLiquidity(creator string, tokenA string, tokenB string, shares string, price string, fee string, orderType string, receiver string) *MsgRemoveLiquidity {
	return &MsgRemoveLiquidity{
		Creator:   creator,
		TokenA:    tokenA,
		TokenB:    tokenB,
		Shares:    shares,
		Price:     price,
		Fee:       fee,
		OrderType: orderType,
		Receiver:  receiver,
	}
}

func (msg *MsgRemoveLiquidity) Route() string {
	return RouterKey
}

func (msg *MsgRemoveLiquidity) Type() string {
	return TypeMsgRemoveLiquidity
}

func (msg *MsgRemoveLiquidity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveLiquidity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
