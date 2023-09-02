package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateHook = "create_hook"
	TypeMsgDeleteHook = "delete_hook"
)

var _ sdk.Msg = &MsgCreateHook{}

func NewMsgCreateHook(creator string, contractID string, args string, persistent bool, triggerKey string, triggerValue string) *MsgCreateHook {
	return &MsgCreateHook{
		Creator:      creator,
		ContractID:   contractID,
		Args:         args,
		Persistent:   persistent,
		TriggerKey:   triggerKey,
		TriggerValue: triggerValue,
	}
}

func (msg *MsgCreateHook) Route() string {
	return RouterKey
}

func (msg *MsgCreateHook) Type() string {
	return TypeMsgCreateHook
}

func (msg *MsgCreateHook) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHook) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHook{}

func NewMsgDeleteHook(creator string, id uint64) *MsgDeleteHook {
	return &MsgDeleteHook{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteHook) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHook) Type() string {
	return TypeMsgDeleteHook
}

func (msg *MsgDeleteHook) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHook) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
