package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

func createNPoolReserves(k *keeper.Keeper, ctx sdk.Context, n int) []*types.PoolReserves {
	items := make([]*types.PoolReserves, n)
	for i := range items {
		pool := keeper.MustNewPool(types.MustNewPairID("TokenA", "TokenB"), int64(i), uint64(i))
		pool.Deposit(sdk.NewInt(10), sdk.NewInt(0), sdk.ZeroInt(), true)
		k.SetPool(ctx, pool)
		items[i] = pool.LowerTick0
	}

	return items
}

func createNPools(k *keeper.Keeper, ctx sdk.Context, n int) []struct {
	Pool      *types.Pool
	TickIndex int64
	Fee       uint64
} {
	items := make([]struct {
		Pool      *types.Pool
		TickIndex int64
		Fee       uint64
	}, n)
	for i := range items {
		pool := keeper.MustNewPool(types.MustNewPairID("TokenA", "TokenB"), int64(i), uint64(i))
		pool.Deposit(sdk.NewInt(10), sdk.NewInt(0), sdk.ZeroInt(), true)
		k.SetPool(ctx, pool)
		items[i] = struct {
			Pool      *types.Pool
			TickIndex int64
			Fee       uint64
		}{
			Pool:      pool,
			TickIndex: int64(i),
			Fee:       uint64(i),
		}
	}

	return items
}

func TestGetPoolReserves(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPoolReserves(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPoolReserves(ctx, item.Key)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestRemovePoolReserves(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPoolReserves(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePoolReserves(ctx, item.Key)
		_, found := keeper.GetPoolReserves(ctx, item.Key)
		require.False(t, found)
	}
}
