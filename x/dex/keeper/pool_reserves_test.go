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

func createNPoolReserves(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PoolReserves {
	items := make([]types.PoolReserves, n)
	for i := range items {
		poolReserves := types.MustNewPoolReserves(
			"TokenA", "TokenB", int64(i), uint64(i), sdk.NewInt(10),
		)
		keeper.SetPoolReserves(ctx, *poolReserves)
		items[i] = *poolReserves
	}

	return items
}

func TestGetPoolReserves(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPoolReserves(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPoolReserves(ctx,
			item.TradePairID,
			item.TickIndex,
			item.Fee,
		)
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
		keeper.RemovePoolReserves(ctx, item)
		_, found := keeper.GetPoolReserves(ctx,
			item.TradePairID,
			item.TickIndex,
			item.Fee,
		)
		require.False(t, found)
	}
}
