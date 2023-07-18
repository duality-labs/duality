package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
)

type MultihopStep struct {
	RemainingBestPrice sdk.Dec
	tradePairID        *types.TradePairID
}

func (k Keeper) HopsToRouteData(
	ctx sdk.Context,
	hops []string,
	exitLimitPrice sdk.Dec,
) ([]MultihopStep, error) {
	nPairs := len(hops) - 1
	routeArr := make([]MultihopStep, nPairs)
	priceAcc := sdk.OneDec()
	for i := range routeArr {
		index := len(routeArr) - 1 - i
		tokenIn := hops[index]
		tokenOut := hops[index+1]
		tradePairID, err := types.NewTradePairID(tokenIn, tokenOut)
		if err != nil {
			return routeArr, err
		}
		price, found := k.GetCurrPrice(ctx, tradePairID)
		if !found {
			return routeArr, types.ErrInsufficientLiquidity
		}
		priceAcc = priceAcc.Mul(price)
		routeArr[index] = MultihopStep{
			tradePairID:        tradePairID,
			RemainingBestPrice: priceAcc,
		}
	}

	return routeArr, nil
}

type StepResult struct {
	Ctx     *types.BranchableCache
	CoinOut sdk.Coin
	CoinIn  sdk.Coin
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
	bctx *types.BranchableCache,
	step MultihopStep,
	inCoin sdk.Coin,
	exitLimitPrice sdk.Dec,
	currentPrice sdk.Dec,
	remainingSteps []MultihopStep,
	stepCache map[multihopCacheKey]StepResult,
) (sdk.Coin, *types.BranchableCache, error) {

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
	coinOut, err := k.SwapExactAmountIn(bctx.Ctx, step.tradePairID, inCoin.Amount)
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
		// If we can't hit the best possible price we can greedily abort
		priceUpperbound := currentPrice.Mul(step.RemainingBestPrice)
		if exitLimitPrice.GT(priceUpperbound) {
			return sdk.Coin{}, bCacheCtx.WriteCache, types.ErrExitLimitPriceHit
		}

		currentOutCoin, bCacheCtx, err = k.MultihopStep(
			bCacheCtx,
			step,
			inCoin,
			exitLimitPrice,
			currentPrice,
			routeData[i:],
			stepCache,
		)
		inCoin = currentOutCoin
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

	return currentOutCoin, bCacheCtx.WriteCache, nil
}

func (k Keeper) SwapExactAmountIn(ctx sdk.Context,
	tradePairID *types.TradePairID,
	amountIn sdk.Int,
) (totalOut sdk.Coin, err error) {
	_, swapAmountMakerDenom, orderFilled, err := k.Swap(
		ctx,
		tradePairID,
		amountIn,
		nil,
		nil,
	)
	if err != nil {
		return sdk.Coin{}, err
	}
	if !orderFilled {
		return sdk.Coin{}, types.ErrInsufficientLiquidity
	}

	return swapAmountMakerDenom, err
}
