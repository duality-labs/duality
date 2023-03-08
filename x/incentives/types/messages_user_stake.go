package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateUserStake = "create_user_stake"
	TypeMsgUpdateUserStake = "update_user_stake"
	TypeMsgDeleteUserStake = "delete_user_stake"
)

var _ sdk.Msg = &MsgCreateUserStake{}

func NewMsgCreateUserStake(
	creator string,
	index string,
	amount sdk.Coins,
	startDate uint64,
	endDate uint64,

) *MsgCreateUserStake {
	return &MsgCreateUserStake{
		Creator:   creator,
		Index:     index,
		Amount:    amount,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func (msg *MsgCreateUserStake) Route() string {
	return RouterKey
}

func (msg *MsgCreateUserStake) Type() string {
	return TypeMsgCreateUserStake
}

func (msg *MsgCreateUserStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateUserStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUserStake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUserStake{}

func NewMsgUpdateUserStake(
	creator string,
	index string,
	amount sdk.Coins,
	startDate uint64,
	endDate uint64,

) *MsgUpdateUserStake {
	return &MsgUpdateUserStake{
		Creator:   creator,
		Index:     index,
		Amount:    amount,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func (msg *MsgUpdateUserStake) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUserStake) Type() string {
	return TypeMsgUpdateUserStake
}

func (msg *MsgUpdateUserStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUserStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUserStake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUserStake{}

func NewMsgDeleteUserStake(
	creator string,
	index string,

) *MsgDeleteUserStake {
	return &MsgDeleteUserStake{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteUserStake) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUserStake) Type() string {
	return TypeMsgDeleteUserStake
}

func (msg *MsgDeleteUserStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUserStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUserStake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
