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

func createNShares(keeper *keeper.Keeper, ctx sdk.Context, n int, token0 string, token1 string) []types.Shares {
	items := make([]types.Shares, n)
	for i := range items {
		items[i].SharesOwned = sdk.ZeroDec()
		items[i].Address = strconv.Itoa(i)
		items[i].Price = strconv.Itoa(i)
		items[i].Fee = strconv.Itoa(i)
		items[i].OrderType = strconv.Itoa(i)

		keeper.SetShares(ctx, token0, token1, items[i])
	}
	return items
}

func TestSharesGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShares(keeper, ctx, 10, "TokenB", "TokenA")
	for _, item := range items {
		rst, found := keeper.GetShares(ctx,
			"TokenB",
			"TokenA",
			item.Address,
			item.Price,
			item.Fee,
			item.OrderType,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSharesRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShares(keeper, ctx, 10, "TokenB", "TokenA")
	for _, item := range items {
		keeper.RemoveShares(ctx,
			"TokenB",
			"TokenA",
			item.Address,
			item.Price,
			item.Fee,
			item.OrderType,
		)
		_, found := keeper.GetShares(ctx,
			"TokenB",
			"TokenA",
			item.Address,
			item.Price,
			item.Fee,
			item.OrderType,
		)
		require.False(t, found)
	}
}

func TestSharesGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShares(keeper, ctx, 10, "Token0", "Token1")
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSharesByPair(ctx, "Token0", "Token1")),
	)
}
