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

func createNLimitOrderPoolUser(keeper *keeper.Keeper, ctx sdk.Context, pairId string, tickIndex int64, token string, n int) []types.LimitOrderPoolUser {
	items := make([]types.LimitOrderPoolUser, n)
	for i := range items {
		items[i].Count = uint64(i)
		items[i].Address = strconv.Itoa(i)
		items[i].PairId = pairId
		items[i].Token = token
		items[i].TickIndex = tickIndex

		keeper.SetLimitOrderPoolUser(ctx, items[i])
		items[i].SharesOwned = sdk.ZeroDec()
		items[i].SharesWithdrawn = sdk.ZeroDec()
		items[i].SharesCancelled = sdk.ZeroDec()

	}
	return items
}

func TestLimitOrderPoolUserGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolUser(keeper, ctx, "TokenA<>TokenB", 0, "TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderPoolUser(ctx,
			"TokenA<>TokenB",
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
func TestLimitOrderPoolUserRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolUser(keeper, ctx, "TokenA<>TokenB", 0, "TokenA", 10)
	for _, item := range items {
		keeper.RemoveLimitOrderPoolUser(ctx,
			"TokenA<>TokenB",
			0,
			"TokenA",
			item.Count,
			item.Address,
		)
		_, found := keeper.GetLimitOrderPoolUser(ctx,
			"TokenA<>TokenB",
			0,
			"TokenA",
			item.Count,
			item.Address,
		)
		require.False(t, found)
	}
}
