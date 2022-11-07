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

func createNLimitOrderPoolFillObject(keeper *keeper.Keeper, ctx sdk.Context, pairId string, tickIndex int64, token string, n int) []types.LimitOrderPoolFillObject {
	items := make([]types.LimitOrderPoolFillObject, n)
	for i := range items {
		items[i].Count = uint64(i)
		items[i].PairId = pairId
		items[i].TickIndex = tickIndex
		items[i].Token = token
		items[i].FilledReserves = sdk.ZeroDec()
		keeper.SetLimitOrderPoolFillObject(ctx, items[i])
	}
	return items
}

func TestLimitOrderPoolFillObjectGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolFillObject(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderPoolFillObject(ctx,
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
func TestLimitOrderPoolFillObjectRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolFillObject(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	for _, item := range items {
		keeper.RemoveLimitOrderPoolFillObject(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
		)
		_, found := keeper.GetLimitOrderPoolFillObject(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
		)
		require.False(t, found)
	}
}
