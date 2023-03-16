package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

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
	tranche, found := k.GetInactiveLimitOrderTranche(ctx, pairId, token, tickIndex, trancheKey)
	if found {
		return tranche, true, true
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
	if tranche.HasTokenIn() {
		k.SetLimitOrderTranche(sdkCtx, tranche)
	} else {
		k.SetInactiveLimitOrderTranche(sdkCtx, tranche)
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

func (k Keeper) GetLimitOrderTrancheByKey(
	ctx sdk.Context,
	key []byte,

) (tranche *types.LimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := store.Get(key)

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

func (k Keeper) GetOrInitPlaceTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (placeTranche types.LimitOrderTranche, err error) {

	placeTranche, found := k.GetPlaceTranche(sdkCtx, pairId, tokenIn, tickIndex)
	if !found {
		placeTranche, err = NewLimitOrderTranche(sdkCtx, pairId, tokenIn, tickIndex, nil)
		if err != nil {
			return types.LimitOrderTranche{}, err
		}
	}
	return placeTranche, nil
}
