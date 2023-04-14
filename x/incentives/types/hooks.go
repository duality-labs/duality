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
	AfterAddTokensToStake(ctx sdk.Context, address sdk.AccAddress, stakeID uint64, amount sdk.Coins)
	OnTokenStaked(ctx sdk.Context, address sdk.AccAddress, stakeID uint64, amount sdk.Coins, unstakeTime time.Time)
	OnTokenUnstaked(ctx sdk.Context, address sdk.AccAddress, stakeID uint64, amount sdk.Coins, unstakeTime time.Time)
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

func (h MultiIncentiveHooks) AfterAddTokensToStake(ctx sdk.Context, address sdk.AccAddress, stakeID uint64, amount sdk.Coins) {
	for i := range h {
		h[i].AfterAddTokensToStake(ctx, address, stakeID, amount)
	}
}

func (h MultiIncentiveHooks) OnTokenStaked(ctx sdk.Context, address sdk.AccAddress, stakeID uint64, amount sdk.Coins, unstakeTime time.Time) {
	for i := range h {
		h[i].OnTokenStaked(ctx, address, stakeID, amount, unstakeTime)
	}
}

func (h MultiIncentiveHooks) OnTokenUnstaked(ctx sdk.Context, address sdk.AccAddress, stakeID uint64, amount sdk.Coins, unstakeTime time.Time) {
	for i := range h {
		h[i].OnTokenUnstaked(ctx, address, stakeID, amount, unstakeTime)
	}
}

// func (h MultiIncentiveHooks) OnStakeupExtend(ctx sdk.Context, stakeID uint64, prevDuration, newDuration time.Duration) {
// 	for i := range h {
// 		h[i].OnStakeupExtend(ctx, stakeID, prevDuration, newDuration)
// 	}
// }
