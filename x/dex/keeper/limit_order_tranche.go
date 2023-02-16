package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/utils"
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
	return &LimitOrderTrancheWrapper{
		LimitOrderTranche: tranche,
		PriceTakerToMaker: tranche.PriceTakerToMaker(),
		PriceMakerToTaker: tranche.PriceMakerToTaker(),
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

func (t *LimitOrderTrancheWrapper) Withdraw(trancheUser types.LimitOrderTrancheUser) (sdk.Int, sdk.Dec) {
	reservesTokenOutDec := sdk.NewDecFromInt(t.ReservesTokenOut)

	amountFilled := t.PriceTakerToMaker.MulInt(t.TotalTokenOut)
	ratioFilled := amountFilled.QuoInt(t.TotalTokenIn)
	maxAllowedToWithdraw := ratioFilled.MulInt(trancheUser.SharesOwned).TruncateInt()

	amountOutTokenIn := maxAllowedToWithdraw.Sub(trancheUser.SharesWithdrawn)
	amountOutTokenOut := t.PriceMakerToTaker.MulInt(amountOutTokenIn)
	t.ReservesTokenOut = reservesTokenOutDec.Sub(amountOutTokenOut).TruncateInt()

	return amountOutTokenIn, amountOutTokenOut

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
	if !t.IsFilled() {
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
	trancheKey string,
) (val types.LimitOrderTranche, fromFilled bool, found bool) {

	// Try to find the tranche in the active liq index
	tick, found := k.GetLimitOrderTranche(ctx, pairId, token, tickIndex, trancheKey)
	if found {
		return *tick, false, true
	}
	// Look for filled limit orders
	tranche, found := k.GetFilledLimitOrderTranche(ctx, pairId, token, tickIndex, trancheKey)
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
		tranche.TrancheKey,
	), b)
}

func (k Keeper) GetLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	trancheKey string,

) (tranche *types.LimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := store.Get(types.TickLiquidityKey(
		pairId,
		tokenIn,
		tickIndex,
		types.LiquidityTypeLimitOrder,
		trancheKey,
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
		tranche.TrancheKey,
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

func (k Keeper) GetFillTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (*types.LimitOrderTranche, bool) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLimitOrderPrefix(pairId, tokenIn, tickIndex))
	iter := sdk.KVStorePrefixIterator(prefixStore, []byte{})

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

func NewTrancheKey(sdkCtx sdk.Context) string {

	blockHeight := sdkCtx.BlockHeight()
	txGas := sdkCtx.GasMeter().GasConsumed()
	blockGas := sdkCtx.BlockGasMeter().GasConsumed()
	totalGas := blockGas + txGas

	blockStr := utils.Uint64ToSortableString(uint64(blockHeight))
	gasStr := utils.Uint64ToSortableString(totalGas)

	return fmt.Sprintf("%s%s", blockStr, gasStr)

}

func (k Keeper) InitPlaceTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (types.LimitOrderTranche, error) {
	// NOTE: CONTRACT: There is no active place tranche (ie. GetPlaceTrancheTick has returned false)

	trancheKey := NewTrancheKey(sdkCtx)
	return NewLimitOrderTranche(pairId, tokenIn, tickIndex, trancheKey)
}
