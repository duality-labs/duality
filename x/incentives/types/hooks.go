package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IncentiveHooks interface {
	AfterCreateGauge(ctx sdk.Context, gaugeId uint64)
	AfterAddToGauge(ctx sdk.Context, gaugeId uint64)
	AfterStartDistribution(ctx sdk.Context, gaugeId uint64)
	AfterFinishDistribution(ctx sdk.Context, gaugeId uint64)
	AfterEpochDistribution(ctx sdk.Context)
	AfterAddTokensToLock(ctx sdk.Context, address sdk.AccAddress, lockID uint64, amount sdk.Coins)
	OnTokenLocked(ctx sdk.Context, address sdk.AccAddress, lockID uint64, amount sdk.Coins, lockDuration time.Duration, unlockTime time.Time)
	OnStartUnlock(ctx sdk.Context, address sdk.AccAddress, lockID uint64, amount sdk.Coins, lockDuration time.Duration, unlockTime time.Time)
	OnTokenUnlocked(ctx sdk.Context, address sdk.AccAddress, lockID uint64, amount sdk.Coins, lockDuration time.Duration, unlockTime time.Time)
	// OnLockupExtend(ctx sdk.Context, lockID uint64, prevDuration time.Duration, newDuration time.Duration)
}

var _ IncentiveHooks = MultiIncentiveHooks{}

// MultiIncentiveHooks combines multiple incentive hooks. All hook functions are run in array sequence.
type MultiIncentiveHooks []IncentiveHooks

// NewMultiIncentiveHooks combines multiple incentive hooks into a single IncentiveHooks array.
func NewMultiIncentiveHooks(hooks ...IncentiveHooks) MultiIncentiveHooks {
	return hooks
}

func (h MultiIncentiveHooks) AfterCreateGauge(ctx sdk.Context, gaugeId uint64) {
	for i := range h {
		h[i].AfterCreateGauge(ctx, gaugeId)
	}
}

func (h MultiIncentiveHooks) AfterAddToGauge(ctx sdk.Context, gaugeId uint64) {
	for i := range h {
		h[i].AfterAddToGauge(ctx, gaugeId)
	}
}

func (h MultiIncentiveHooks) AfterStartDistribution(ctx sdk.Context, gaugeId uint64) {
	for i := range h {
		h[i].AfterStartDistribution(ctx, gaugeId)
	}
}

func (h MultiIncentiveHooks) AfterFinishDistribution(ctx sdk.Context, gaugeId uint64) {
	for i := range h {
		h[i].AfterFinishDistribution(ctx, gaugeId)
	}
}

func (h MultiIncentiveHooks) AfterEpochDistribution(ctx sdk.Context) {
	for i := range h {
		h[i].AfterEpochDistribution(ctx)
	}
}

func (h MultiIncentiveHooks) AfterAddTokensToLock(ctx sdk.Context, address sdk.AccAddress, lockID uint64, amount sdk.Coins) {
	for i := range h {
		h[i].AfterAddTokensToLock(ctx, address, lockID, amount)
	}
}

func (h MultiIncentiveHooks) OnTokenLocked(ctx sdk.Context, address sdk.AccAddress, lockID uint64, amount sdk.Coins, lockDuration time.Duration, unlockTime time.Time) {
	for i := range h {
		h[i].OnTokenLocked(ctx, address, lockID, amount, lockDuration, unlockTime)
	}
}

func (h MultiIncentiveHooks) OnStartUnlock(ctx sdk.Context, address sdk.AccAddress, lockID uint64, amount sdk.Coins, lockDuration time.Duration, unlockTime time.Time) {
	for i := range h {
		h[i].OnStartUnlock(ctx, address, lockID, amount, lockDuration, unlockTime)
	}
}

func (h MultiIncentiveHooks) OnTokenUnlocked(ctx sdk.Context, address sdk.AccAddress, lockID uint64, amount sdk.Coins, lockDuration time.Duration, unlockTime time.Time) {
	for i := range h {
		h[i].OnTokenUnlocked(ctx, address, lockID, amount, lockDuration, unlockTime)
	}
}

// func (h MultiIncentiveHooks) OnLockupExtend(ctx sdk.Context, lockID uint64, prevDuration, newDuration time.Duration) {
// 	for i := range h {
// 		h[i].OnLockupExtend(ctx, lockID, prevDuration, newDuration)
// 	}
// }
