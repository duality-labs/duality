package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

func (k Keeper) SetTickLiquidityLimitOrder(ctx sdk.Context, tranche types.LimitOrderTranche) {
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

func (k Keeper) GetTickLiquidityLimitOrder(
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

func (k Keeper) RemoveLimitOrder(ctx sdk.Context, tranche types.LimitOrderTranche) {
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
