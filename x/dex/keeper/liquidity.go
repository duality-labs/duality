package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type Liquidity interface {
	Swap(maxAmountIn sdk.Int, maxAmountOut sdk.Int) (inAmount, outAmount sdk.Int)
	Price() sdk.Dec
}

func (k Keeper) Swap(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	maxAmountTakerDenom sdk.Int,
	maxAmountMakerDenom sdk.Int,
	limitPrice *sdk.Dec,
) (totalTakerCoin, totalMakerCoin sdk.Coin, err error) {
	useMaxOut := !maxAmountMakerDenom.IsZero()

	remainingTakerDenom := maxAmountTakerDenom
	remainingMakerDenom := maxAmountMakerDenom
	totalMakerDenom := sdk.ZeroInt()

	// verify that amount left is not zero and that there are additional valid ticks to check
	liqIter := k.NewLiquidityIterator(ctx, tradePairID)
	defer liqIter.Close()
	for remainingTakerDenom.GT(sdk.ZeroInt()) {
		liq := liqIter.Next()
		if liq == nil {
			break
		}

		// break as soon as we iterated past limitPrice
		if limitPrice != nil && liq.Price().LT(*limitPrice) {
			break
		}

		amountTakerDenom, amountMakerDenom := liq.Swap(remainingTakerDenom, remainingMakerDenom)

		remainingTakerDenom = remainingTakerDenom.Sub(amountTakerDenom)
		totalMakerDenom = totalMakerDenom.Add(amountMakerDenom)
		if useMaxOut {
			remainingMakerDenom = remainingMakerDenom.Sub(amountMakerDenom)
		}

		k.SaveLiquidity(ctx, liq)

		// if remaining out has been used up then exit
		if useMaxOut && remainingMakerDenom.LTE(sdk.ZeroInt()) {
			break
		}
	}
	totalTakerDenom := maxAmountTakerDenom.Sub(remainingTakerDenom)

	return sdk.NewCoin(tradePairID.TakerDenom, totalTakerDenom), sdk.NewCoin(tradePairID.MakerDenom, totalMakerDenom), nil
}

func (k Keeper) SwapExactAmountIn(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	amountTakerDenom sdk.Int,
	maxAmountMakerDenom sdk.Int,
	limitPrice *sdk.Dec,
) (totalTaker, totalMaker sdk.Coin, err error) {
	swapAmountTakerDenom, swapAmountMakerDenom, err := k.Swap(ctx, tradePairID, amountTakerDenom, maxAmountMakerDenom, limitPrice)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}
	if !swapAmountTakerDenom.Amount.Equal(amountTakerDenom) {
		return sdk.Coin{}, sdk.Coin{}, types.ErrInsufficientLiquidity
	}

	return swapAmountTakerDenom, swapAmountMakerDenom, err
}

func (k Keeper) SwapWithCache(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	maxAmountTakerDenom sdk.Int,
	maxAmountMakerDenom sdk.Int,
	limitPrice sdk.Dec,
) (totalTakerDenom, totalMakerDenom sdk.Coin, err error) {
	cacheCtx, writeCache := ctx.CacheContext()
	totalTakerDenom, totalMakerDenom, err = k.Swap(cacheCtx, tradePairID, maxAmountTakerDenom, maxAmountMakerDenom, &limitPrice)

	writeCache()

	// NOTE: in current version events from the cache are never passed to the
	// parent context. This is fixed in cosmos v0.46.4
	// Once we update, the below code can be removed
	ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())

	return totalTakerDenom, totalMakerDenom, err
}

func (k Keeper) SaveLiquidity(sdkCtx sdk.Context, liquidityI Liquidity) {
	switch liquidity := liquidityI.(type) {
	case *types.LimitOrderTranche:
		k.SaveTranche(sdkCtx, liquidity)

	case *PoolLiquidity:
		k.SavePool(sdkCtx, *liquidity.pool)
	default:
		panic("Invalid liquidity type")
	}
}
