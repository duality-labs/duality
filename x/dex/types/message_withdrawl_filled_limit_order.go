package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawFilledLimitOrder = "withdrawl_withdrawn_limit_order"

var _ sdk.Msg = &MsgWithdrawFilledLimitOrder{}

func NewMsgWithdrawFilledLimitOrder(creator string, receiver string, tokenA string, tokenB string, tickIndex int64, keyToken string, key uint64) *MsgWithdrawFilledLimitOrder {
	return &MsgWithdrawFilledLimitOrder{
		Creator:   creator,
		Receiver:  receiver,
		TokenA:    tokenA,
		TokenB:    tokenB,
		TickIndex: tickIndex,
		KeyToken:  keyToken,
		Key:       key,
	}
}

func (msg *MsgWithdrawFilledLimitOrder) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawFilledLimitOrder) Type() string {
	return TypeMsgWithdrawFilledLimitOrder
}

func (msg *MsgWithdrawFilledLimitOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawFilledLimitOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawFilledLimitOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	if msg.KeyToken != msg.TokenA && msg.KeyToken != msg.TokenB {
		return sdkerrors.Wrapf(ErrInvalidTradingPair, "KeyToken must be either TokenA or TokenB")
	}

	return nil
}
