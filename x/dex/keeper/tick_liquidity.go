package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

func (k Keeper) SetTickLiquidity(ctx sdk.Context, tick types.TickLiquidity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := k.cdc.MustMarshal(&tick)
	store.Set(types.TickLiquidityKey(
		tick.PairId(),
		tick.TokenIn(),
		tick.TickIndex(),
		types.LiquidityTypeLimitOrder,
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
