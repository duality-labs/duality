package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLimitOrderPoolTotalSharesMap(keeper *keeper.Keeper, ctx sdk.Context, pairId string, tickIndex int64, token string, n int) []types.LimitOrderPoolTotalSharesMap {
	items := make([]types.LimitOrderPoolTotalSharesMap, n)
	for i := range items {
		items[i].Count = uint64(i)

		keeper.SetLimitOrderPoolTotalSharesMap(ctx, pairId, tickIndex, token, items[i])
	}
	return items
}

func TestLimitOrderPoolTotalSharesMapGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolTotalSharesMap(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderPoolTotalSharesMap(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLimitOrderPoolTotalSharesMapRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolTotalSharesMap(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	for _, item := range items {
		keeper.RemoveLimitOrderPoolTotalSharesMap(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
		)
		_, found := keeper.GetLimitOrderPoolTotalSharesMap(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
		)
		require.False(t, found)
	}
}

func TestLimitOrderPoolTotalSharesMapGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolTotalSharesMap(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLimitOrderPoolTotalSharesMap(ctx)),
	)
}
