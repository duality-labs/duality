package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTickLiquidity set a specific tickLiquidity in the store from its index
func (k Keeper) SetTickLiquidity(ctx sdk.Context, tickLiquidity types.TickLiquidity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := k.cdc.MustMarshal(&tickLiquidity)
	store.Set(types.TickLiquidityKey(
		tickLiquidity.PairId,
		tickLiquidity.TokenIn,
		tickLiquidity.TickIndex,
		tickLiquidity.LiquidityType,
		tickLiquidity.LiquidityIndex,
	), b)
}

// GetTickLiquidity returns a tickLiquidity from its index
func (k Keeper) GetTickLiquidity(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	liquidityType string,
	liquidityIndex uint64,

) (val types.TickLiquidity, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))

	b := store.Get(types.TickLiquidityKey(
		pairId,
		tokenIn,
		tickIndex,
		liquidityType,
		liquidityIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetTickLiquidityLP(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	fee uint64,

) (val types.TickLiquidity, found bool) {
	return k.GetTickLiquidity(ctx, pairId, tokenIn, tickIndex, types.LiquidityTypeLP, fee)
}

func (k Keeper) GetTickLiquidityLO(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	trancheIndex uint64,

) (val types.TickLiquidity, found bool) {
	return k.GetTickLiquidity(ctx, pairId, tokenIn, tickIndex, types.LiquidityTypeLO, trancheIndex)
}

// RemoveTickLiquidity removes a tickLiquidity from the store
func (k Keeper) RemoveTickLiquidity(ctx sdk.Context, tick types.TickLiquidity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	store.Delete(types.TickLiquidityKey(
		tick.PairId,
		tick.TokenIn,
		tick.TickIndex,
		tick.LiquidityType,
		tick.LiquidityIndex,
	))
}

// GetAllTickLiquidity returns all tickLiquidity
func (k Keeper) GetAllTickLiquidity(ctx sdk.Context) (list []types.TickLiquidity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TickLiquidity
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetPlaceTrancheTick(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (types.TickLiquidity, bool) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLOPrefix(pairId, tokenIn, tickIndex))
	iter := prefixStore.Iterator(nil, nil)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)
		tranche := tick.LimitOrderTranche
		if tranche.IsPlaceTranche() {
			return tick, true

		}
	}
	return types.TickLiquidity{}, false
}

func (k Keeper) GetNewestLimitOrderTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (*types.LimitOrderTranche, bool) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLOPrefix(pairId, tokenIn, tickIndex))
	iter := sdk.KVStoreReversePrefixIterator(prefixStore, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)
		return tick.LimitOrderTranche, true
	}
	return &types.LimitOrderTranche{}, false
}

func (k Keeper) GetAllLimitOrderTrancheAtIndex(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (tickList []types.TickLiquidity) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLOPrefix(pairId, tokenIn, tickIndex))
	iter := sdk.KVStoreReversePrefixIterator(prefixStore, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var tick types.TickLiquidity
		k.cdc.MustUnmarshal(iter.Value(), &tick)
		tickList = append(tickList, tick)
	}
	return tickList
}
func (k Keeper) InitPlaceTrancheTick(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (types.TickLiquidity, error) {
	// NOTE: CONTRACT: There is no active place tranche (ie. GetPlaceTrancheTick has returned false)

	//TODO: This could probably be made more efficient since at this point it requires 3 lookups in the worst case
	// ideally we can find a way to generate trancheIds that are lexographically increasing witout any lookups
	// we can get close to this with sdkCtx.BlockTime(), but would have to track number of tranches created in a given block
	// to handle cases where multiple placeTranches are created in a single block

	newestActiveTranche, found := k.GetNewestLimitOrderTranche(sdkCtx, pairId, tokenIn, tickIndex)
	if found {
		newTrancheIndex := newestActiveTranche.TrancheIndex + 1
		return NewTickLO(pairId, tokenIn, tickIndex, newTrancheIndex)
	}
	newestFilledTranche, found := k.GetNewestFilledLimitOrderTranche(sdkCtx, pairId, tokenIn, tickIndex)

	if found {
		newTrancheIndex := newestFilledTranche.TrancheIndex + 1
		return NewTickLO(pairId, tokenIn, tickIndex, newTrancheIndex)
	}

	return NewTickLO(pairId, tokenIn, tickIndex, 0)
}
