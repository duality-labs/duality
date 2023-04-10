package types

import (
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateGauge       = "create_gauge"
	TypeMsgAddToGauge        = "add_to_gauge"
	TypeMsgLockTokens        = "lock_tokens"
	TypeMsgBeginUnlockingAll = "begin_unlocking_all"
	TypeMsgBeginUnlocking    = "begin_unlocking"
	TypeMsgExtendLockup      = "edit_lockup"
)

var _ sdk.Msg = &MsgCreateGauge{}

// NewMsgCreateGauge creates a message to create a gauge with the provided parameters.
func NewMsgCreateGauge(isPerpetual bool, owner sdk.AccAddress, distributeTo QueryCondition, coins sdk.Coins, startTime time.Time, numEpochsPaidOver uint64) *MsgCreateGauge {
	return &MsgCreateGauge{
		IsPerpetual:       isPerpetual,
		Owner:             owner.String(),
		DistributeTo:      distributeTo,
		Coins:             coins,
		StartTime:         startTime,
		NumEpochsPaidOver: numEpochsPaidOver,
	}
}

// Route takes a create gauge message, then returns the RouterKey used for slashing.
func (m MsgCreateGauge) Route() string { return RouterKey }

// Type takes a create gauge message, then returns a create gauge message type.
func (m MsgCreateGauge) Type() string { return TypeMsgCreateGauge }

// ValidateBasic checks that the create gauge message is valid.
func (m MsgCreateGauge) ValidateBasic() error {
	if m.Owner == "" {
		return errors.New("owner should be set")
	}
	// TODO
	if m.StartTime.Equal(time.Time{}) {
		return errors.New("distribution start time should be set")
	}
	if m.NumEpochsPaidOver == 0 {
		return errors.New("distribution period should be at least 1 epoch")
	}
	if m.IsPerpetual && m.NumEpochsPaidOver != 1 {
		return errors.New("distribution period should be 1 epoch for perpetual gauge")
	}

	return nil
}

// GetSignBytes takes a create gauge message and turns it into a byte array.
func (m MsgCreateGauge) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners takes a create gauge message and returns the owner in a byte array.
func (m MsgCreateGauge) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgAddToGauge{}

// NewMsgAddToGauge creates a message to add rewards to a specific gauge.
func NewMsgAddToGauge(owner sdk.AccAddress, gaugeId uint64, rewards sdk.Coins) *MsgAddToGauge {
	return &MsgAddToGauge{
		Owner:   owner.String(),
		GaugeId: gaugeId,
		Rewards: rewards,
	}
}

// Route takes an add to gauge message, then returns the RouterKey used for slashing.
func (m MsgAddToGauge) Route() string { return RouterKey }

// Type takes an add to gauge message, then returns an add to gauge message type.
func (m MsgAddToGauge) Type() string { return TypeMsgAddToGauge }

// ValidateBasic checks that the add to gauge message is valid.
func (m MsgAddToGauge) ValidateBasic() error {
	if m.Owner == "" {
		return errors.New("owner should be set")
	}
	if m.Rewards.Empty() {
		return errors.New("additional rewards should not be empty")
	}

	return nil
}

// GetSignBytes takes an add to gauge message and turns it into a byte array.
func (m MsgAddToGauge) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners takes an add to gauge message and returns the owner in a byte array.
func (m MsgAddToGauge) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgLockTokens{}

// NewMsgLockTokens creates a message to lock tokens.
func NewMsgSetupLock(owner sdk.AccAddress, duration time.Duration, coins sdk.Coins) *MsgLockTokens {
	return &MsgLockTokens{
		Owner: owner.String(),
		Coins: coins,
	}
}

func (m MsgLockTokens) Route() string { return RouterKey }
func (m MsgLockTokens) Type() string  { return TypeMsgLockTokens }
func (m MsgLockTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid owner address (%s)", err)
	}

	// we only allow locks with one denom for now
	if m.Coins.Len() != 1 {
		return fmt.Errorf("lockups can only have one denom per lock ID, got %v", m.Coins)
	}

	if !m.Coins.IsAllPositive() {
		return fmt.Errorf("cannot lock up a zero or negative amount")
	}

	return nil
}

func (m MsgLockTokens) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgLockTokens) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgBeginUnlockingAll{}

// NewMsgBeginUnlockingAll creates a message to begin unlocking tokens.
func NewMsgBeginUnlockingAll(owner sdk.AccAddress) *MsgBeginUnlockingAll {
	return &MsgBeginUnlockingAll{
		Owner: owner.String(),
	}
}

func (m MsgBeginUnlockingAll) Route() string { return RouterKey }
func (m MsgBeginUnlockingAll) Type() string  { return TypeMsgBeginUnlockingAll }
func (m MsgBeginUnlockingAll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid owner address (%s)", err)
	}
	return nil
}

func (m MsgBeginUnlockingAll) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgBeginUnlockingAll) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgBeginUnlocking{}

// NewMsgBeginUnlocking creates a message to begin unlocking the tokens of a specific lock.
func NewMsgBeginUnlocking(owner sdk.AccAddress, id uint64, coins sdk.Coins) *MsgBeginUnlocking {
	return &MsgBeginUnlocking{
		Owner: owner.String(),
		ID:    id,
		Coins: coins,
	}
}

func (m MsgBeginUnlocking) Route() string { return RouterKey }
func (m MsgBeginUnlocking) Type() string  { return TypeMsgBeginUnlocking }
func (m MsgBeginUnlocking) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid owner address (%s)", err)
	}

	if m.ID == 0 {
		return fmt.Errorf("invalid lockup ID, got %v", m.ID)
	}

	// only allow unlocks with a single denom or empty
	if m.Coins.Len() > 1 {
		return fmt.Errorf("can only unlock one denom per lock ID, got %v", m.Coins)
	}

	if !m.Coins.Empty() && !m.Coins.IsAllPositive() {
		return fmt.Errorf("cannot unlock a zero or negative amount")
	}

	return nil
}

func (m MsgBeginUnlocking) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgBeginUnlocking) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{owner}
}

// // NewMsgExtendLockup creates a message to edit the properties of existing locks
// func NewMsgExtendLockup(owner sdk.AccAddress, id uint64, duration time.Duration) *MsgExtendLockup {
// 	return &MsgExtendLockup{
// 		Owner:    owner.String(),
// 		ID:       id,
// 		Duration: duration,
// 	}
// }

// func (m MsgExtendLockup) Route() string { return RouterKey }
// func (m MsgExtendLockup) Type() string  { return TypeMsgExtendLockup }
// func (m MsgExtendLockup) ValidateBasic() error {
// 	_, err := sdk.AccAddressFromBech32(m.Owner)
// 	if err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid owner address (%s)", err)
// 	}
// 	if m.ID == 0 {
// 		return fmt.Errorf("id is empty")
// 	}
// 	if m.Duration <= 0 {
// 		return fmt.Errorf("duration should be positive: %d < 0", m.Duration)
// 	}
// 	return nil
// }

// func (m MsgExtendLockup) GetSignBytes() []byte {
// 	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON((&m)))
// }

// func (m MsgExtendLockup) GetSigners() []sdk.AccAddress {
// 	owner, _ := sdk.AccAddressFromBech32(m.Owner)
// 	return []sdk.AccAddress{owner}
// }
