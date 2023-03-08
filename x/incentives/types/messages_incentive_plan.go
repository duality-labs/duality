package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateIncentivePlan = "create_incentive_plan"
	TypeMsgUpdateIncentivePlan = "update_incentive_plan"
	TypeMsgDeleteIncentivePlan = "delete_incentive_plan"
)

var _ sdk.Msg = &MsgCreateIncentivePlan{}

func NewMsgCreateIncentivePlan(
	creator string,
	index string,
	startDate uint64,
	endDate uint64,
	tradingPair string,
	totalAmount sdk.Coin,
	startTick int64,
	endTick int64,
) *MsgCreateIncentivePlan {
	return &MsgCreateIncentivePlan{
		Creator:     creator,
		Index:       index,
		StartDate:   startDate,
		EndDate:     endDate,
		TradingPair: tradingPair,
		TotalAmount: totalAmount,
		StartTick:   startTick,
		EndTick:     endTick,
	}
}

func (msg *MsgCreateIncentivePlan) Route() string {
	return RouterKey
}

func (msg *MsgCreateIncentivePlan) Type() string {
	return TypeMsgCreateIncentivePlan
}

func (msg *MsgCreateIncentivePlan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateIncentivePlan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIncentivePlan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateIncentivePlan{}

func NewMsgUpdateIncentivePlan(
	creator string,
	index string,
	startDate uint64,
	endDate uint64,
	tradingPair string,
	totalAmount sdk.Coin,
	startTick int64,
	endTick int64,

) *MsgUpdateIncentivePlan {
	return &MsgUpdateIncentivePlan{
		Creator:     creator,
		Index:       index,
		StartDate:   startDate,
		EndDate:     endDate,
		TradingPair: tradingPair,
		TotalAmount: totalAmount,
		StartTick:   startTick,
		EndTick:     endTick,
	}
}

func (msg *MsgUpdateIncentivePlan) Route() string {
	return RouterKey
}

func (msg *MsgUpdateIncentivePlan) Type() string {
	return TypeMsgUpdateIncentivePlan
}

func (msg *MsgUpdateIncentivePlan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateIncentivePlan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateIncentivePlan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteIncentivePlan{}

func NewMsgDeleteIncentivePlan(
	creator string,
	index string,

) *MsgDeleteIncentivePlan {
	return &MsgDeleteIncentivePlan{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteIncentivePlan) Route() string {
	return RouterKey
}

func (msg *MsgDeleteIncentivePlan) Type() string {
	return TypeMsgDeleteIncentivePlan
}

func (msg *MsgDeleteIncentivePlan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteIncentivePlan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteIncentivePlan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
