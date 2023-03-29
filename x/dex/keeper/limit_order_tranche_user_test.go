package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

func createNLimitOrderTrancheUser(keeper *keeper.Keeper, ctx sdk.Context, tickIndex int64, token string, n int) []types.LimitOrderTrancheUser {
	items := make([]types.LimitOrderTrancheUser, n)
	for i := range items {
		items[i].TrancheKey = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(i)
		items[i].PairID = &types.PairID{Token0: "TokenA", Token1: "TokenB"}
		items[i].Token = token
		items[i].TickIndex = tickIndex
		items[i].SharesOwned = sdk.ZeroInt()
		items[i].SharesWithdrawn = sdk.ZeroInt()
		items[i].SharesCancelled = sdk.ZeroInt()
		items[i].TakerReserves = sdk.ZeroInt()

		keeper.SetLimitOrderTrancheUser(ctx, items[i])

	}
	return items
}

func TestLimitOrderTrancheUserGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTrancheUser(keeper, ctx, 0, "TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderTrancheUser(ctx, item.Address, item.TrancheKey)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestLimitOrderTrancheUserRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTrancheUser(keeper, ctx, 0, "TokenA", 10)
	for _, item := range items {
		keeper.RemoveLimitOrderTrancheUserByKey(ctx,
			&types.PairID{Token0: "TokenA", Token1: "TokenB"},
			0,
			"TokenA",
			item.TrancheKey,
			item.Address,
		)
		_, found := keeper.GetLimitOrderTrancheUser(ctx, item.Address, item.TrancheKey)
		require.False(t, found)
	}
}
