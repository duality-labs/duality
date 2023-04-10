package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

type DistributorKeeper interface {
	ValueForShares(ctx sdk.Context, coin sdk.Coin, tick int64) (sdk.Int, error)
	GetLocksByQueryCondition(ctx sdk.Context, distrTo *types.QueryCondition) types.Locks
}

type Distributor struct {
	keeper DistributorKeeper
}

func NewDistributor(keeper DistributorKeeper) Distributor {
	return Distributor{
		keeper: keeper,
	}
}

func (d Distributor) Distribute(ctx sdk.Context, gauge *types.Gauge, filterLocks types.Locks) (types.DistributionSpec, error) {
	if !gauge.IsActiveGauge(ctx.BlockTime()) {
		return nil, types.ErrGaugeNotActive
	}

	distSpec := types.DistributionSpec{}

	rewardsNextEpoch := gauge.RewardsNextEpoch()

	adjustedGaugeTotal := sdk.ZeroInt()
	lockSumCache := map[uint64]sdk.Int{}
	gaugeLocks := d.keeper.GetLocksByQueryCondition(ctx, &gauge.DistributeTo)
	for _, lock := range gaugeLocks {
		lockCoins := lock.CoinsPassingQueryCondition(gauge.DistributeTo)
		lockTotal := sdk.ZeroInt()
		for _, lockCoin := range lockCoins {
			adjustedPositionValue, err := d.keeper.ValueForShares(ctx, lockCoin, gauge.PricingTick)
			if err != nil {
				return nil, err
			}
			lockTotal = lockTotal.Add(adjustedPositionValue)
			adjustedGaugeTotal = adjustedGaugeTotal.Add(adjustedPositionValue)
		}
		lockSumCache[lock.ID] = lockTotal
	}
	if adjustedGaugeTotal.IsZero() {
		return distSpec, nil
	}

	for _, lock := range filterLocks {
		distCoins := sdk.Coins{}
		for _, epochRewards := range rewardsNextEpoch {
			// distribution amount = gauge_size * denom_lock_amount / (total_denom_lock_amount * remain_epochs)
			lockAmt := lockSumCache[lock.ID]
			amount := epochRewards.Amount.Mul(lockAmt).Quo(adjustedGaugeTotal)
			reward := sdk.Coin{Denom: epochRewards.Denom, Amount: amount}
			distCoins = distCoins.Add(reward)
		}

		// update the amount for that address
		if distCoins.Empty() {
			continue
		}

		gauge.DistributedCoins = gauge.DistributedCoins.Add(distCoins...)
		if spec, ok := distSpec[lock.Owner]; ok {
			distSpec[lock.Owner] = spec.Add(distCoins...)
		} else {
			distSpec[lock.Owner] = distCoins
		}
	}

	gauge.FilledEpochs += 1
	return distSpec, nil
}
