package keeper

import (
	"fmt"
	"time"

	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ DistributorKeeper = Keeper{}

func (k Keeper) ValueForShares(ctx sdk.Context, coin sdk.Coin, tick int64) (sdk.Int, error) {
	totalShares := k.bk.GetSupply(ctx, coin.Denom).Amount
	depositDenom, err := dexkeeper.NewDepositDenomFromString(coin.Denom)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	pool, err := k.dk.GetOrInitPool(
		ctx,
		depositDenom.PairID,
		depositDenom.Tick,
		depositDenom.Fee,
	)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	amount0, amount1 := pool.RedeemValue(coin.Amount, totalShares)
	price1To0Center := dextypes.MustNewPrice(-1 * tick)
	return amount0.ToDec().Add(price1To0Center.MulInt(amount1)).TruncateInt(), nil
}

// Distribute distributes coins from an array of gauges to all eligible locks.
func (k Keeper) Distribute(ctx sdk.Context, gauges types.Gauges) (types.DistributionSpec, error) {
	distSpec := types.DistributionSpec{}
	for _, gauge := range gauges {
		filterLocks := k.GetLocksByQueryCondition(ctx, &gauge.DistributeTo)
		gaugeDistSpec, err := k.distributor.Distribute(ctx, gauge, filterLocks)

		if err != nil {
			return nil, err
		}
		distSpec = distSpec.Add(gaugeDistSpec)

		err = k.setGauge(ctx, gauge)
		if err != nil {
			return nil, err
		}
		if gauge.IsFinishedGauge(ctx.BlockTime()) {
			if err := k.moveActiveGaugeToFinishedGauge(ctx, gauge); err != nil {
				return nil, err
			}
		}
	}

	ctx.Logger().Debug(fmt.Sprintf("Beginning distribution to %d users", len(distSpec)))
	for addr, rewards := range distSpec {
		decodedAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		err = k.bk.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			decodedAddr,
			rewards)
		if err != nil {
			return nil, err
		}
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.TypeEvtDistribution,
				sdk.NewAttribute(types.AttributeReceiver, addr),
				sdk.NewAttribute(types.AttributeAmount, rewards.String()),
			),
		})
	}
	ctx.Logger().Debug(fmt.Sprintf("Finished Distributing to %d users", len(distSpec)))
	k.hooks.AfterEpochDistribution(ctx)
	return distSpec, nil
}

// GetModuleCoinsToBeDistributed returns sum of coins yet to be distributed for all of the module.
func (k Keeper) GetModuleCoinsToBeDistributed(ctx sdk.Context) sdk.Coins {
	activeGaugesDistr := k.GetActiveGauges(ctx).GetCoinsRemaining()
	upcomingGaugesDistr := k.GetUpcomingGauges(ctx).GetCoinsRemaining()
	return activeGaugesDistr.Add(upcomingGaugesDistr...)
}

// GetModuleDistributedCoins returns sum of coins that have been distributed so far for all of the module.
func (k Keeper) GetModuleDistributedCoins(ctx sdk.Context) sdk.Coins {
	activeGaugesDistr := k.GetActiveGauges(ctx).GetCoinsDistributed()
	finishedGaugesDistr := k.GetFinishedGauges(ctx).GetCoinsDistributed()
	return activeGaugesDistr.Add(finishedGaugesDistr...)
}

// GetRewardsEstimate returns rewards estimation at a future specific time (by epoch)
// If locks are nil, it returns the rewards between now and the end epoch associated with address.
// If locks are not nil, it returns all the rewards for the given locks between now and end epoch.
func (k Keeper) GetRewardsEstimate(ctx sdk.Context, addr sdk.AccAddress, filterLocks types.Locks, endEpoch int64) (sdk.Coins, error) {
	// if locks are nil, populate with all locks associated with the address
	if len(filterLocks) == 0 {
		filterLocks = k.GetLocksByAccount(ctx, addr)
	}

	// for each specified lock get associated pairs
	pairSet := map[dextypes.PairID]bool{}
	for _, l := range filterLocks {
		for _, c := range l.Coins {
			depositDenom, err := dexkeeper.NewDepositDenomFromString(c.Denom)
			if err != nil {
				panic("all locks should be valid deposit denoms")
			}
			pairSet[*depositDenom.PairID] = true
		}
	}

	// for each pair get associated gauges
	gauges := types.Gauges{}
	for s := range pairSet {
		gauges = append(gauges, k.GetGaugesByPair(ctx, &s)...)
	}

	// estimate rewards
	estimatedRewards := sdk.Coins{}
	epochInfo := k.GetEpochInfo(ctx)

	// ensure we don't change storage while doing estimation
	cacheCtx, _ := ctx.CacheContext()
	for _, gauge := range gauges {
		distrBeginEpoch := epochInfo.CurrentEpoch
		blockTime := ctx.BlockTime()
		if gauge.StartTime.After(blockTime) {
			distrBeginEpoch = epochInfo.CurrentEpoch + 1 + int64(gauge.StartTime.Sub(blockTime)/epochInfo.Duration)
		}

		// TODO: Make more efficient by making it possible to call distribute with this
		// gaugeLocks := k.GetLocksByQueryCondition(cacheCtx, &gauge.DistributeTo)
		for epoch := distrBeginEpoch; epoch <= endEpoch; epoch++ {
			epochTime := epochInfo.StartTime.Add(time.Duration(epoch-epochInfo.CurrentEpoch) * epochInfo.Duration)
			if !gauge.IsActiveGauge(epochTime) {
				break
			}

			futureCtx := cacheCtx.WithBlockTime(epochTime)
			distSpec, err := k.distributor.Distribute(futureCtx, gauge, filterLocks)
			if err != nil {
				return nil, err
			}

			estimatedRewards = estimatedRewards.Add(distSpec.GetTotal()...)
		}
	}

	return estimatedRewards, nil
}
