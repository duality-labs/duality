package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
)

type MultihopStep struct {
	BestPrice   types.Price
	TradingPair types.DirectionalTradingPair
}

func (k Keeper) HopsToRouteData(ctx sdk.Context, hops []string, exitLimitPrice sdk.Dec) ([]MultihopStep, error) {
	nPairs := len(hops) - 1
	routeArr := make([]MultihopStep, nPairs)
	priceUpperbound := sdk.OneDec()
	for i := range routeArr {
		tokenIn := hops[i]
		tokenOut := hops[i+1]
		pairID, err := CreatePairIDFromUnsorted(tokenIn, tokenOut)
		dPair := types.NewDirectionalTradingPair(pairID, tokenIn, tokenOut)
		if err != nil {
			return routeArr, err
		}
		var price types.Price
		var found bool
		if pairID.Token0 == hops[i] {
			price, found = k.GetCurrPrice0To1(ctx, pairID)
		} else {
			price, found = k.GetCurrPrice1To0(ctx, pairID)
		}
		if !found {
			return routeArr, types.ErrInsufficientLiquidity
		}
		priceUpperbound = price.Mul(priceUpperbound)
		routeArr[i] = MultihopStep{
			TradingPair: dPair,
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
	cacheKey := newCacheKey(step.TradingPair.TokenIn, step.TradingPair.TokenOut, inCoin.Amount)
	val, ok := stepCache[cacheKey]
	if ok {
		ctxBranchCopy := val.Ctx.Branch()
		return val.CoinOut, ctxBranchCopy, val.Err
	}

	_, coinOut, err := k.SwapExactAmountIn(
		bctx.Ctx,
		step.TradingPair.PairID,
		step.TradingPair.TokenIn,
		step.TradingPair.TokenOut,
		inCoin.Amount,
		sdk.ZeroInt(),
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
			return sdk.Coin{}, nil, sdkerrors.Wrapf(err, "Failed at pair: %s", step.TradingPair.PairID.Stringify())
		}
		currentPrice = currentOutCoin.Amount.ToDec().Quo(initialInCoin.Amount.ToDec())
	}

	if exitLimitPrice.GT(currentPrice) {
		return sdk.Coin{}, nil, types.ErrExitLimitPriceHit
	}

	return currentOutCoin, bCacheCtx.Write, nil
}
