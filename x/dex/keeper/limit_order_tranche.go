package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type LimitOrderTrancheWrapper struct {
	*types.LimitOrderTranche
	PriceTakerToMaker sdk.Dec
	PriceMakerToTaker sdk.Dec
}

func NewLimitOrderTrancheWrapper(
	tranche *types.LimitOrderTranche,
) *LimitOrderTrancheWrapper {
	var priceMakerToTaker, priceTakerToMaker sdk.Dec
	if tranche.TokenIn == tranche.PairId.Token0 {
		priceMakerToTaker = MustCalcPrice0To1(tranche.TickIndex)
		priceTakerToMaker = MustCalcPrice1To0(tranche.TickIndex)
	} else {
		priceMakerToTaker = MustCalcPrice1To0(tranche.TickIndex)
		priceTakerToMaker = MustCalcPrice0To1(tranche.TickIndex)
	}
	return &LimitOrderTrancheWrapper{
		LimitOrderTranche: tranche,
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
	if t.HasToken() {
		keeper.SetLimitOrderTranche(sdkCtx, *t.LimitOrderTranche)
	} else {
		filledTranche := t.LimitOrderTranche.CreateFilledTranche()
		keeper.SetFilledLimitOrderTranche(sdkCtx, filledTranche)
		keeper.RemoveLimitOrderTranche(sdkCtx, *t.LimitOrderTranche)
	}

}

func (t *LimitOrderTrancheWrapper) Price() sdk.Dec {
	return t.PriceTakerToMaker
}

func (t LimitOrderTrancheWrapper) HasLiquidity() bool {
	return t.ReservesTokenIn.GT(sdk.ZeroInt())
}

func (k Keeper) FindLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tickIndex int64,
	token string,
	trancheIndex uint64,
) (val types.LimitOrderTranche, fromFilled bool, found bool) {

	// Try to find the tranche in the active liq index
	tick, found := k.GetLimitOrderTranche(ctx, pairId, token, tickIndex, trancheIndex)
	if found {
		return *tick, false, true
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

func (k Keeper) SaveTranche(sdkCtx sdk.Context, tranche types.LimitOrderTranche) {
	if tranche.HasToken() {
		k.SetLimitOrderTranche(sdkCtx, tranche)
	} else {
		filledTranche := tranche.CreateFilledTranche()
		k.SetFilledLimitOrderTranche(sdkCtx, filledTranche)
		k.RemoveLimitOrderTranche(sdkCtx, tranche)
	}

}

func (k Keeper) SetLimitOrderTranche(ctx sdk.Context, tranche types.LimitOrderTranche) {
	//Wrap tranche back into TickLiquidity
	tick := types.TickLiquidity{
		Liquidity: &types.TickLiquidity_LimitOrderTranche{
			LimitOrderTranche: &tranche,
		},
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := k.cdc.MustMarshal(&tick)
	store.Set(types.TickLiquidityKey(
		tranche.PairId,
		tranche.TokenIn,
		tranche.TickIndex,
		types.LiquidityTypeLimitOrder,
		tranche.TrancheIndex,
	), b)
}

func (k Keeper) GetLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	trancheIndex uint64,

) (tranche *types.LimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := store.Get(types.TickLiquidityKey(
		pairId,
		tokenIn,
		tickIndex,
		types.LiquidityTypeLimitOrder,
		trancheIndex,
	))

	if b == nil {
		return nil, false
	}

	var tick types.TickLiquidity
	k.cdc.MustUnmarshal(b, &tick)
	return tick.GetLimitOrderTranche(), true
}

func (k Keeper) RemoveLimitOrderTranche(ctx sdk.Context, tranche types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	store.Delete(types.TickLiquidityKey(
		tranche.PairId,
		tranche.TokenIn,
		tranche.TickIndex,
		types.LiquidityTypeLimitOrder,
		tranche.TrancheIndex,
	))
}

func (k Keeper) GetPlaceTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (types.LimitOrderTranche, bool) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLimitOrderPrefix(pairId, tokenIn, tickIndex))
	iter := prefixStore.Iterator(nil, nil)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)
		tranche := tick.GetLimitOrderTranche()
		if tranche.IsPlaceTranche() {
			return *tranche, true

		}
	}
	return types.LimitOrderTranche{}, false
}

func (k Keeper) GetNewestLimitOrderTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (*types.LimitOrderTranche, bool) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLimitOrderPrefix(pairId, tokenIn, tickIndex))
	iter := sdk.KVStoreReversePrefixIterator(prefixStore, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)
		return tick.GetLimitOrderTranche(), true
	}
	return &types.LimitOrderTranche{}, false
}

func (k Keeper) GetAllLimitOrderTrancheAtIndex(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (trancheList []types.LimitOrderTranche) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLimitOrderPrefix(pairId, tokenIn, tickIndex))
	iter := sdk.KVStoreReversePrefixIterator(prefixStore, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)
		trancheList = append(trancheList, *tick.GetLimitOrderTranche())
	}
	return trancheList
}
func (k Keeper) InitPlaceTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (types.LimitOrderTranche, error) {
	// NOTE: CONTRACT: There is no active place tranche (ie. GetPlaceTrancheTick has returned false)

	//TODO: This could probably be made more efficient since at this point it requires 3 lookups in the worst case
	// ideally we can find a way to generate trancheIds that are lexographically increasing witout any lookups
	// we can get close to this with sdkCtx.BlockTime(), but would have to track number of tranches created in a given block
	// to handle cases where multiple placeTranches are created in a single block

	newestActiveTranche, found := k.GetNewestLimitOrderTranche(sdkCtx, pairId, tokenIn, tickIndex)
	if found {
		newTrancheIndex := newestActiveTranche.TrancheIndex + 1
		return NewLimitOrderTranche(pairId, tokenIn, tickIndex, newTrancheIndex)
	}
	newestFilledTranche, found := k.GetNewestFilledLimitOrderTranche(sdkCtx, pairId, tokenIn, tickIndex)

	if found {
		newTrancheIndex := newestFilledTranche.TrancheIndex + 1
		return NewLimitOrderTranche(pairId, tokenIn, tickIndex, newTrancheIndex)
	}

	return NewLimitOrderTranche(pairId, tokenIn, tickIndex, 0)
}
