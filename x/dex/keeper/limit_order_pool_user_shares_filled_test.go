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

func createNLimitOrderPoolUserSharesWithdrawn(keeper *keeper.Keeper, ctx sdk.Context, pairId string, tickIndex int64, token string, n int) []types.LimitOrderPoolUserSharesWithdrawn {
	items := make([]types.LimitOrderPoolUserSharesWithdrawn, n)
	for i := range items {
		items[i].Count = uint64(i)
		items[i].Address = strconv.Itoa(i)
		items[i].PairId = pairId
		items[i].TickIndex = tickIndex
		items[i].Token = token

		keeper.SetLimitOrderPoolUserSharesWithdrawn(ctx, items[i])
	}
	return items
}

func TestLimitOrderPoolUserSharesWithdrawnGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolUserSharesWithdrawn(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderPoolUserSharesWithdrawn(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLimitOrderPoolUserSharesWithdrawnRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolUserSharesWithdrawn(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	for _, item := range items {
		keeper.RemoveLimitOrderPoolUserSharesWithdrawn(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
			item.Address,
		)
		_, found := keeper.GetLimitOrderPoolUserSharesWithdrawn(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestLimitOrderPoolUserSharesWithdrawnGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolUserSharesWithdrawn(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLimitOrderPoolUserSharesWithdrawn(ctx)),
	)
}
