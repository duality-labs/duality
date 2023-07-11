package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
)

type MultihopStep struct {
	BestPrice   sdk.Dec
	tradePairID *types.TradePairID
}

func (k Keeper) HopsToRouteData(
	ctx sdk.Context,
	hops []string,
	exitLimitPrice sdk.Dec,
) ([]MultihopStep, error) {
	nPairs := len(hops) - 1
	routeArr := make([]MultihopStep, nPairs)
	priceUpperbound := sdk.OneDec()
	for i := range routeArr {
		tokenIn := hops[i]
		tokenOut := hops[i+1]
		tradePairID, err := types.NewTradePairID(tokenIn, tokenOut)
		if err != nil {
			return routeArr, err
		}
		price, found := k.GetCurrPrice(ctx, tradePairID)
		if !found {
			return routeArr, types.ErrInsufficientLiquidity
		}
		priceUpperbound = price.Mul(priceUpperbound)
		routeArr[i] = MultihopStep{
			tradePairID: tradePairID,
			BestPrice:   price,
		}
	}

	if exitLimitPrice.GT(priceUpperbound) {
		return routeArr, types.ErrExitLimitPriceHit
	}

	return routeArr, nil
}

func CalcMultihopPriceUpperbound(currentPrice sdk.Dec, remainingSteps []MultihopStep) sdk.Dec {
	price := currentPrice
	for _, step := range remainingSteps {
		price = step.BestPrice.Mul(price)
	}

	return price
}

type StepResult struct {
	Ctx     types.BranchableCache
	CoinOut sdk.Coin
	Err     error
}

type multihopCacheKey struct {
	TokenIn  string
	TokenOut string
	InAmount sdk.Int
}

func newCacheKey(tokenIn, tokenOut string, inAmount sdk.Int) multihopCacheKey {
	return multihopCacheKey{
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
		InAmount: inAmount,
	}
}

func (k Keeper) MultihopStep(
	bctx types.BranchableCache,
	step MultihopStep,
	inCoin sdk.Coin,
	exitLimitPrice sdk.Dec,
	currentPrice sdk.Dec,
	remainingSteps []MultihopStep,
	stepCache map[multihopCacheKey]StepResult,
) (sdk.Coin, types.BranchableCache, error) {
	priceUpperbound := CalcMultihopPriceUpperbound(currentPrice, remainingSteps)
	if exitLimitPrice.GT(priceUpperbound) {
		// If we can't hit the best possible price we can greedily abort
		return sdk.Coin{}, bctx, types.ErrExitLimitPriceHit
	}
	cacheKey := newCacheKey(step.tradePairID.TakerDenom, step.tradePairID.MakerDenom, inCoin.Amount)
	val, ok := stepCache[cacheKey]
	if ok {
		ctxBranchCopy := val.Ctx.Branch()
		return val.CoinOut, ctxBranchCopy, val.Err
	}

	// TODO: Due to rounding on swap it is possible to leak tokens at each hop.
	// In these cases the user will end up with trace amounts of tokens from intermediary steps.
	// To fix this we would have to pre-calculate the route such that the amount
	// in will be used completely at each step
	_, coinOut, err := k.SwapExactAmountIn(
		bctx.Ctx,
		step.tradePairID,
		inCoin.Amount,
		nil,
		nil,
	)
	ctxBranch := bctx.Branch()
	stepCache[cacheKey] = StepResult{Ctx: bctx, CoinOut: coinOut, Err: err}
	if err != nil {
		return sdk.Coin{}, bctx, err
	}

	return coinOut, ctxBranch, nil
}

func (k Keeper) RunMultihopRoute(
	ctx sdk.Context,
	route types.MultiHopRoute,
	initialInCoin sdk.Coin,
	exitLimitPrice sdk.Dec,
	stepCache map[multihopCacheKey]StepResult,
) (sdk.Coin, func(), error) {
	routeData, err := k.HopsToRouteData(ctx, route.Hops, exitLimitPrice)
	if err != nil {
		return sdk.Coin{}, nil, err
	}
	currentPrice := sdk.OneDec()

	var currentOutCoin sdk.Coin
	inCoin := initialInCoin
	bCacheCtx := types.NewBranchableCache(ctx)

	for i, step := range routeData {
		currentOutCoin, bCacheCtx, err = k.MultihopStep(
			bCacheCtx,
			step,
			inCoin,
			exitLimitPrice,
			currentPrice,
			routeData[i:],
			stepCache,
		)
		if err != nil {
			return sdk.Coin{}, nil, sdkerrors.Wrapf(
				err,
				"Failed at pair: %s",
				step.tradePairID.MustPairID().CanonicalString(),
			)
		}
		currentPrice = sdk.NewDecFromInt(currentOutCoin.Amount).
			Quo(sdk.NewDecFromInt(initialInCoin.Amount))
	}

	if exitLimitPrice.GT(currentPrice) {
		return sdk.Coin{}, nil, types.ErrExitLimitPriceHit
	}

	return currentOutCoin, bCacheCtx.Write, nil
}

func (k Keeper) SwapExactAmountIn(ctx sdk.Context,
	tradePairID *types.TradePairID,
	amountIn sdk.Int,
	maxAmountOut *sdk.Int,
	limitPrice *sdk.Dec,
) (totalIn, totalOut sdk.Coin, err error) {
	swapAmountTakerDenom, swapAmountMakerDenom, orderFilled, err := k.Swap(
		ctx,
		tradePairID,
		amountIn,
		maxAmountOut,
		limitPrice,
	)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}
	if !orderFilled {
		return sdk.Coin{}, sdk.Coin{}, types.ErrInsufficientLiquidity
	}

	return swapAmountTakerDenom, swapAmountMakerDenom, err
}
