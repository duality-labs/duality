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

func createNShares(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Shares {
	items := make([]types.Shares, n)
	for i := range items {
		items[i].SharesOwned = sdk.ZeroInt()
		items[i].Address = strconv.Itoa(i)
		items[i].PairId = strconv.Itoa(i)
		items[i].TickIndex = int64(i)
		items[i].FeeIndex = uint64(i)

		keeper.SetShares(ctx, items[i])
	}
	return items
}

func TestSharesGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShares(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetShares(ctx,
			item.Address,
			item.PairId,
			item.TickIndex,
			item.FeeIndex,
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
	items := createNShares(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShares(ctx,
			item.Address,
			item.PairId,
			item.TickIndex,
			item.FeeIndex,
		)
		_, found := keeper.GetShares(ctx,
			item.Address,
			item.PairId,
			item.TickIndex,
			item.FeeIndex,
		)
		require.False(t, found)
	}
}

func TestSharesGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShares(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllShares(ctx)),
	)
}
