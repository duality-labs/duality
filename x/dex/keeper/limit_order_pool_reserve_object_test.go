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

func createNLimitOrderPoolReserveObject(keeper *keeper.Keeper, ctx sdk.Context, pairId string, tickIndex int64, token string, n int) []types.LimitOrderPoolReserveObject {
	items := make([]types.LimitOrderPoolReserveObject, n)
	for i := range items {
		items[i].Count = uint64(i)
		items[i].TickIndex = tickIndex
		items[i].Token = token
		items[i].PairId = pairId
		items[i].Reserves = sdk.ZeroDec()

		keeper.SetLimitOrderPoolReserveObject(ctx, items[i])
	}
	return items
}

func TestLimitOrderPoolReserveObjectGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolReserveObject(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderPoolReserveObject(ctx,
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
func TestLimitOrderPoolReserveObjectRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderPoolReserveObject(keeper, ctx, "TokenA/TokenB", 0, "TokenA", 10)
	for _, item := range items {
		keeper.RemoveLimitOrderPoolReserveObject(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
		)
		_, found := keeper.GetLimitOrderPoolReserveObject(ctx,
			"TokenA/TokenB",
			0,
			"TokenA",
			item.Count,
		)
		require.False(t, found)
	}
}
