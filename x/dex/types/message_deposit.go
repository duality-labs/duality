package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeposit = "deposit"

var _ sdk.Msg = &MsgDeposit{}

func NewMsgDeposit(creator string, receiver string, tokenA string, tokenB string, amountsA []sdk.Int, amountsB []sdk.Int, tickIndexes []int64, fees []uint64, depositOptions []*DepositOptions) *MsgDeposit {
	return &MsgDeposit{
		Creator:     creator,
		Receiver:    receiver,
		TokenA:      tokenA,
		TokenB:      tokenB,
		AmountsA:    amountsA,
		AmountsB:    amountsB,
		TickIndexes: tickIndexes,
		Fees:        fees,
		Options:     depositOptions,
	}
}

func (msg *MsgDeposit) Route() string {
	return RouterKey
}

func (msg *MsgDeposit) Type() string {
	return TypeMsgDeposit
}

func (msg *MsgDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	// Verify that the lengths of TickIndexes, Fees, AmountsA, AmountsB are all equal
	if len(msg.Fees) != len(msg.TickIndexes) ||
		len(msg.AmountsA) != len(msg.AmountsB) ||
		len(msg.AmountsA) != len(msg.TickIndexes) {
		return ErrUnbalancedTxArray
	}

	if len(msg.AmountsA) == 0 {
		return ErrZeroDeposit
	}

	for i := 0; i < len(msg.AmountsA); i++ {
		if msg.AmountsA[i].LTE(sdk.ZeroInt()) && msg.AmountsB[i].LTE(sdk.ZeroInt()) {
			return ErrZeroDeposit
		}
	}
	return nil
}
