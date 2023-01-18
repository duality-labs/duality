package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LimitOrderTrancheWrapper struct {
	*types.LimitOrderTranche
	Tick              *types.TickLiquidity
	PriceTakerToMaker sdk.Dec
	PriceMakerToTaker sdk.Dec
}

func NewLimitOrderTranche(
	tick *types.TickLiquidity,
) *LimitOrderTrancheWrapper {
	var priceMakerToTaker, priceTakerToMaker sdk.Dec
	tranche := tick.LimitOrderTranche
	if tranche.TokenIn == tick.PairId.Token0 {
		priceMakerToTaker = MustCalcPrice0To1(tick.TickIndex)
		priceTakerToMaker = MustCalcPrice1To0(tick.TickIndex)
	} else {
		priceMakerToTaker = MustCalcPrice1To0(tick.TickIndex)
		priceTakerToMaker = MustCalcPrice0To1(tick.TickIndex)
	}
	return &LimitOrderTrancheWrapper{
		LimitOrderTranche: tranche,
		Tick:              tick,
		PriceTakerToMaker: priceTakerToMaker,
		PriceMakerToTaker: priceMakerToTaker,
	}
}

func (t *LimitOrderTrancheWrapper) Cancel(trancheUser types.LimitOrderTrancheUser) (amountToCancel sdk.Int) {
	totalTokenInDec := sdk.NewDecFromInt(t.TotalTokenIn)
	totalTokenOutDec := sdk.NewDecFromInt(t.TotalTokenOut)

	filledAmount := t.PriceTakerToMaker.Mul(totalTokenOutDec)
	ratioNotFilled := totalTokenInDec.Sub(filledAmount).Quo(totalTokenInDec)

	amountToCancel = trancheUser.SharesOwned.ToDec().Mul(ratioNotFilled).TruncateInt()
	t.ReservesTokenIn = t.ReservesTokenIn.Sub(amountToCancel)

	return amountToCancel

}
func (t *LimitOrderTrancheWrapper) Swap(maxAmountTaker sdk.Int) (
	inAmount sdk.Int,
	outAmount sdk.Int,
) {
	reservesTokenOut := &t.ReservesTokenIn
	fillTokenIn := &t.ReservesTokenOut
	totalTokenIn := &t.TotalTokenOut
	amountFilledTokenOut := maxAmountTaker.ToDec().Mul(t.PriceTakerToMaker).TruncateInt()
	if reservesTokenOut.LTE(amountFilledTokenOut) {
		inAmount = reservesTokenOut.ToDec().Mul(t.PriceMakerToTaker).TruncateInt()
		outAmount = *reservesTokenOut
		*reservesTokenOut = sdk.ZeroInt()
		*fillTokenIn = fillTokenIn.Add(inAmount)
		*totalTokenIn = totalTokenIn.Add(inAmount)
	} else {
		inAmount = maxAmountTaker
		outAmount = amountFilledTokenOut
		*fillTokenIn = fillTokenIn.Add(maxAmountTaker)
		*totalTokenIn = totalTokenIn.Add(maxAmountTaker)
		*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
	}
	return inAmount, outAmount
}

func (t LimitOrderTrancheWrapper) IsFilled() bool {
	return t.ReservesTokenIn.IsZero()
}

func (t *LimitOrderTrancheWrapper) Save(sdkCtx sdk.Context, keeper Keeper) {
	if t.Tick.HasActiveLimitOrders() {
		keeper.SetTickLiquidity(sdkCtx, *t.Tick)
	} else {
		filledTranche := t.Tick.LimitOrderTranche.CreateFilledTranche()
		keeper.SetFilledLimitOrderTranche(sdkCtx, filledTranche)
		keeper.RemoveTickLiquidity(sdkCtx, *t.Tick)
	}

}

func (t *LimitOrderTrancheWrapper) Price() sdk.Dec {
	return t.PriceTakerToMaker
}

func (t LimitOrderTrancheWrapper) HasLiquidity() bool {
	return t.ReservesTokenIn.GT(sdk.ZeroInt())
}

func (k Keeper) GetLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tickIndex int64,
	token string,
	trancheIndex uint64,
) (val types.LimitOrderTranche, fromFilled bool, found bool) {

	// Try to find the tranche in the active liq index
	tick, found := k.GetTickLiquidityLO(ctx, pairId, token, tickIndex, trancheIndex)
	if found {
		return *tick.LimitOrderTranche, false, true
	}
	// Look for filled limit orders
	tranche, found := k.GetFilledLimitOrderTranche(ctx, pairId, token, tickIndex, trancheIndex)
	if found {
		return types.NewFromFilledTranche(tranche), true, true
	}
	return types.LimitOrderTranche{}, false, false
}

func (k Keeper) GetAllLimitOrderTranche(ctx sdk.Context) (list []types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderTranche
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
