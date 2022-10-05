package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawlWithdrawnLimitOrder = "withdrawl_withdrawn_limit_order"

var _ sdk.Msg = &MsgWithdrawlWithdrawnLimitOrder{}

func NewMsgWithdrawlWithdrawnLimitOrder(creator string, tokenA string, tokenB string, tickIndex int64, keyToken string, key uint64) *MsgWithdrawlWithdrawnLimitOrder {
	return &MsgWithdrawlWithdrawnLimitOrder{
		Creator:   creator,
		TokenA:    tokenA,
		TokenB:    tokenB,
		TickIndex: tickIndex,
		KeyToken:  keyToken,
		Key:       key,
	}
}

func (msg *MsgWithdrawlWithdrawnLimitOrder) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawlWithdrawnLimitOrder) Type() string {
	return TypeMsgWithdrawlWithdrawnLimitOrder
}

func (msg *MsgWithdrawlWithdrawnLimitOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawlWithdrawnLimitOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawlWithdrawnLimitOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
