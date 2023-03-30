package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawal = "withdrawal"

var _ sdk.Msg = &MsgWithdrawal{}

func NewMsgWithdrawal(creator,
	receiver,
	tokenA,
	tokenB string,
	sharesToRemove []sdk.Int,
	tickIndexes []int64,
	fees []uint64,
) *MsgWithdrawal {
	return &MsgWithdrawal{
		Creator:         creator,
		Receiver:        receiver,
		TokenA:          tokenA,
		TokenB:          tokenB,
		SharesToRemove:  sharesToRemove,
		TickIndexesAToB: tickIndexes,
		Fees:            fees,
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

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	// Verify that the lengths of TickIndexes, Fees, SharesToRemove are all equal
	if len(msg.Fees) != len(msg.TickIndexesAToB) ||
		len(msg.SharesToRemove) != len(msg.TickIndexesAToB) {
		return ErrUnbalancedTxArray
	}

	if len(msg.Fees) == 0 {
		return ErrZeroWithdraw
	}

	for i := 0; i < len(msg.Fees); i++ {
		if msg.SharesToRemove[i].LTE(sdk.ZeroInt()) {
			return ErrZeroWithdraw
		}
	}

	return nil
}
