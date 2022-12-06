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

func createNLimitOrderTranche(keeper *keeper.Keeper, ctx sdk.Context, pairId string, tickIndex int64, token string, n int) []types.LimitOrderTranche {
	items := make([]types.LimitOrderTranche, n)
	for i := range items {
		items[i] = types.LimitOrderTranche{
			TrancheIndex:     uint64(i),
			PairId:           pairId,
			TickIndex:        tickIndex,
			TokenIn:          token,
			ReservesTokenIn:  sdk.ZeroInt(),
			ReservesTokenOut: sdk.ZeroInt(),
			TotalTokenIn:     sdk.ZeroInt(),
			TotalTokenOut:    sdk.ZeroInt(),
		}
		keeper.SetLimitOrderTranche(ctx, items[i])
	}
	return items
}

func TestLimitOrderTrancheGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTranche(keeper, ctx, "TokenA<>TokenB", 0, "TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderTranche(ctx,
			"TokenA<>TokenB",
			0,
			"TokenA",
			item.TrancheIndex,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLimitOrderTrancheRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTranche(keeper, ctx, "TokenA<>TokenB", 0, "TokenA", 10)
	for _, item := range items {
		keeper.RemoveLimitOrderTranche(ctx,
			"TokenA<>TokenB",
			0,
			"TokenA",
			item.TrancheIndex,
		)
		_, found := keeper.GetLimitOrderTranche(ctx,
			"TokenA<>TokenB",
			0,
			"TokenA",
			item.TrancheIndex,
		)
		require.False(t, found)
	}
}

func TestLimitOrderTrancheGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTranche(keeper, ctx, "TokenA<>TokenB", 0, "TokenA", 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLimitOrderTranche(ctx)),
	)
}
