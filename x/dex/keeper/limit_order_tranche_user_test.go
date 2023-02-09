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

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLimitOrderTrancheUser(keeper *keeper.Keeper, ctx sdk.Context, tickIndex int64, token string, n int) []types.LimitOrderTrancheUser {
	items := make([]types.LimitOrderTrancheUser, n)
	for i := range items {
		items[i].TrancheKey = uint64(i)
		items[i].Address = strconv.Itoa(i)
		items[i].PairId = &types.PairId{Token0: "TokenA", Token1: "TokenB"}
		items[i].Token = token
		items[i].TickIndex = tickIndex

		keeper.SetLimitOrderTrancheUser(ctx, items[i])
		items[i].SharesOwned = sdk.ZeroInt()
		items[i].SharesWithdrawn = sdk.ZeroInt()
		items[i].SharesCancelled = sdk.ZeroInt()

	}
	return items
}

func TestLimitOrderTrancheUserGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTrancheUser(keeper, ctx, 0, "TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderTrancheUser(ctx,
			defaultPairId,
			0,
			"TokenA",
			item.TrancheKey,
			item.Address,
		)
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
		keeper.RemoveLimitOrderTrancheUser(ctx,
			&types.PairId{Token0: "TokenA", Token1: "TokenB"},
			0,
			"TokenA",
			item.TrancheKey,
			item.Address,
		)
		_, found := keeper.GetLimitOrderTrancheUser(ctx,
			&types.PairId{Token0: "TokenA", Token1: "TokenB"},
			0,
			"TokenA",
			item.TrancheKey,
			item.Address,
		)
		require.False(t, found)
	}
}
