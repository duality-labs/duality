package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

func (k Keeper) SetPoolReserves(ctx sdk.Context, pool types.PoolReserves) {
	// Wrap pool back into TickLiquidity
	tick := types.TickLiquidity{
		Liquidity: &types.TickLiquidity_PoolReserves{
			PoolReserves: &pool,
		},
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := k.cdc.MustMarshal(&tick)
	store.Set(types.TickLiquidityKey(
		pool.PairID,
		pool.TokenIn,
		pool.TickIndex,
		types.LiquidityTypePoolReserves,
		pool.Fee,
	), b)
}

func (k Keeper) GetPoolReserves(
	ctx sdk.Context,
	pairID *types.PairID,
	tokenIn string,
	tickIndex int64,
	fee uint64,
) (pool *types.PoolReserves, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	b := store.Get(types.TickLiquidityKey(
		pairID,
		tokenIn,
		tickIndex,
		types.LiquidityTypePoolReserves,
		fee,
	))
	if b == nil {
		return nil, false
	}

	var tick types.TickLiquidity
	k.cdc.MustUnmarshal(b, &tick)

	return tick.GetPoolReserves(), true
}

// RemoveTickLiquidity removes a tickLiquidity from the store
func (k Keeper) RemovePoolReserves(ctx sdk.Context, pool types.PoolReserves) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickLiquidityKeyPrefix))
	store.Delete(types.TickLiquidityKey(
		pool.PairID,
		pool.TokenIn,
		pool.TickIndex,
		types.LiquidityTypePoolReserves,
		pool.Fee,
	))
}
