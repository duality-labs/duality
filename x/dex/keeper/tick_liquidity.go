package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// SetTickLiquidity set a specific tickLiquidity in the store from its index
func (k Keeper) SetTickLiquidityPoolReserves(ctx sdk.Context, pool types.PoolReserves) {
	//Wrap pool back into TickLiquidity
	tick := types.TickLiquidity{
		Liquidity: &types.TickLiquidity_PoolReserves{
			PoolReserves: &pool,
		},
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := k.cdc.MustMarshal(&tick)
	store.Set(types.TickLiquidityKey(
		pool.PairId,
		pool.TokenIn,
		pool.TickIndex,
		types.LiquidityTypeLP,
		pool.Fee,
	), b)
}

func (k Keeper) SetTickLiquidityLO(ctx sdk.Context, tranche types.LimitOrderTranche) {
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
		types.LiquidityTypeLO,
		tranche.TrancheIndex,
	), b)
}

func (k Keeper) SetTickLiquidity(ctx sdk.Context, tick types.TickLiquidity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := k.cdc.MustMarshal(&tick)
	store.Set(types.TickLiquidityKey(
		tick.PairId(),
		tick.TokenIn(),
		tick.TickIndex(),
		types.LiquidityTypeLO,
		tick.LiquidityIndex(),
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

func (k Keeper) GetTickLiquidityInterface(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	liquidityType string,
	liquidityIndex uint64,

) (val types.TickLiquidityI, found bool) {
	tick, found := k.GetTickLiquidity(ctx, pairId, tokenIn, tickIndex, liquidityType, liquidityIndex)
	if !found {
		return nil, false
	}
	return MustTickLiquidityToInterface(tick), true
}

func MustTickLiquidityToInterface(tick types.TickLiquidity) types.TickLiquidityI {
	switch liquidity := tick.Liquidity.(type) {

	case *types.TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche

	case *types.TickLiquidity_PoolReserves:
		return liquidity.PoolReserves
	default:
		panic("Tick does not contain valid liqudityType")
	}
}

func (k Keeper) GetTickLiquidityLP(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	fee uint64,

) (val types.TickLiquidityI, found bool) {
	return k.GetTickLiquidityInterface(ctx, pairId, tokenIn, tickIndex, types.LiquidityTypeLP, fee)
}

func (k Keeper) GetTickLiquidityLO(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
	tickIndex int64,
	trancheIndex uint64,

) (val types.TickLiquidityI, found bool) {
	return k.GetTickLiquidityInterface(ctx, pairId, tokenIn, tickIndex, types.LiquidityTypeLO, trancheIndex)
}

// RemoveTickLiquidity removes a tickLiquidity from the store
func (k Keeper) RemoveTickLiquidity(ctx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64, liqudityType string, liquidityIndex uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	store.Delete(types.TickLiquidityKey(
		pairId,
		tokenIn,
		tickIndex,
		liqudityType,
		liquidityIndex,
	))
}

func (k Keeper) RemovePoolReserves(ctx sdk.Context, pool types.PoolReserves) {
	k.RemoveTickLiquidity(ctx, pool.PairId, pool.TokenIn, pool.TickIndex, types.LiquidityTypeLP, pool.Fee)
}

func (k Keeper) RemoveLimitOrder(ctx sdk.Context, tranche types.LimitOrderTranche) {
	k.RemoveTickLiquidity(ctx, tranche.PairId, tranche.TokenIn, tranche.TickIndex, types.LiquidityTypeLO, tranche.TrancheIndex)
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

func (k Keeper) GetPlaceTranche(sdkCtx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) (types.LimitOrderTranche, bool) {
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLOPrefix(pairId, tokenIn, tickIndex))
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
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLOPrefix(pairId, tokenIn, tickIndex))
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
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickLiquidityLOPrefix(pairId, tokenIn, tickIndex))
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
		return NewTickLO(pairId, tokenIn, tickIndex, newTrancheIndex)
	}
	newestFilledTranche, found := k.GetNewestFilledLimitOrderTranche(sdkCtx, pairId, tokenIn, tickIndex)

	if found {
		newTrancheIndex := newestFilledTranche.TrancheIndex + 1
		return NewTickLO(pairId, tokenIn, tickIndex, newTrancheIndex)
	}

	return NewTickLO(pairId, tokenIn, tickIndex, 0)
}
