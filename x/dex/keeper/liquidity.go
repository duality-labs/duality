package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

func (k Keeper) Swap(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	maxAmountTakerDenom sdk.Int,
	maxAmountMakerDenom *sdk.Int,
	limitPrice *sdk.Dec,
) (totalTakerCoin, totalMakerCoin sdk.Coin, orderFilled bool, err error) {

	useMaxOut := maxAmountMakerDenom != nil
	var remainingMakerDenom *sdk.Int
	if useMaxOut {
		copy := *maxAmountMakerDenom
		remainingMakerDenom = &copy
	}

	remainingTakerDenom := maxAmountTakerDenom
	totalMakerDenom := sdk.ZeroInt()
	orderFilled = false

	// verify that amount left is not zero and that there are additional valid ticks to check
	liqIter := k.NewLiquidityIterator(ctx, tradePairID)
	defer liqIter.Close()
	for {
		liq := liqIter.Next()
		if liq == nil {
			break
		}

		// break as soon as we iterated past limitPrice
		if limitPrice != nil && liq.Price().LT(*limitPrice) {
			break
		}

		inAmount, outAmount := liq.Swap(remainingTakerDenom, remainingMakerDenom)
		k.SaveLiquidity(ctx, liq)

		remainingTakerDenom = remainingTakerDenom.Sub(inAmount)
		totalMakerDenom = totalMakerDenom.Add(outAmount)

		// break if remainingTakerDenom will yield less than 1 tokenOut at current price
		// this avoids unnecessary iteration since outAmount will always be 0 going forward
		// this also catches the normal exit case where remainingTakerDenom == 0
		if liq.Price().MulInt(remainingTakerDenom).LT(sdk.OneDec()) {
			orderFilled = true
			break
		}

		if useMaxOut {
			copy := remainingMakerDenom.Sub(outAmount)
			remainingMakerDenom = &copy

			// if maxAmountOut has been used up then exit
			if remainingMakerDenom.LTE(sdk.ZeroInt()) {
				orderFilled = true
				break
			}
		}
	}
	totalTakerDenom := maxAmountTakerDenom.Sub(remainingTakerDenom)

	return sdk.NewCoin(
			tradePairID.TakerDenom,
			totalTakerDenom,
		), sdk.NewCoin(
			tradePairID.MakerDenom,
			totalMakerDenom,
		), orderFilled, nil
}

func (k Keeper) SwapWithCache(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	maxAmountIn sdk.Int,
	maxAmountOut *sdk.Int,
	limitPrice *sdk.Dec,
) (totalIn, totalOut sdk.Coin, orderFilled bool, err error) {
	cacheCtx, writeCache := ctx.CacheContext()
	totalIn, totalOut, orderFilled, err = k.Swap(
		cacheCtx,
		tradePairID,
		maxAmountIn,
		maxAmountOut,
		limitPrice,
	)

	writeCache()

	// NOTE: in current version events from the cache are never passed to the
	// parent context. This is fixed in cosmos v0.46.4
	// Once we update, the below code can be removed
	ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())

	return totalIn, totalOut, orderFilled, err
}

func (k Keeper) SaveLiquidity(sdkCtx sdk.Context, liquidityI types.Liquidity) {
	switch liquidity := liquidityI.(type) {
	case *types.LimitOrderTranche:
		k.SaveTranche(sdkCtx, liquidity)

	case *types.PoolLiquidity:
		k.SetPool(sdkCtx, liquidity.Pool)
	default:
		panic("Invalid liquidity type")
	}
}
