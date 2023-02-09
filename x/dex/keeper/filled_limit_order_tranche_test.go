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

func createNFilledLimitOrderTranche(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FilledLimitOrderTranche {
	items := make([]types.FilledLimitOrderTranche, n)
	for i := range items {
		items[i] = types.FilledLimitOrderTranche{
			TrancheKey:       strconv.Itoa(i),
			PairId:           &types.PairId{Token0: "TokenA", Token1: "TokenB"},
			TickIndex:        int64(i),
			TokenIn:          "TokenA",
			TotalTokenIn:     sdk.ZeroInt(),
			TotalTokenOut:    sdk.ZeroInt(),
			ReservesTokenOut: sdk.ZeroInt(),
		}
		keeper.SetFilledLimitOrderTranche(ctx, items[i])
	}
	return items
}

func createNFilledLimitOrderTrancheSameTick(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FilledLimitOrderTranche {
	items := make([]types.FilledLimitOrderTranche, n)
	for i := range items {
		items[i] = types.FilledLimitOrderTranche{
			TrancheKey:       strconv.Itoa(i),
			PairId:           &types.PairId{Token0: "TokenA", Token1: "TokenB"},
			TickIndex:        0,
			TokenIn:          "TokenA",
			TotalTokenIn:     sdk.ZeroInt(),
			TotalTokenOut:    sdk.ZeroInt(),
			ReservesTokenOut: sdk.ZeroInt(),
		}
		keeper.SetFilledLimitOrderTranche(ctx, items[i])
	}
	return items
}

func TestFilledLimitOrderTrancheGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFilledLimitOrderTranche(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFilledLimitOrderTranche(ctx,
			item.PairId,
			item.TokenIn,
			item.TickIndex,
			item.TrancheKey,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFilledLimitOrderTrancheRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFilledLimitOrderTranche(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFilledLimitOrderTranche(ctx,
			item.PairId,
			item.TokenIn,
			item.TickIndex,
			item.TrancheKey,
		)
		_, found := keeper.GetFilledLimitOrderTranche(ctx,
			item.PairId,
			item.TokenIn,
			item.TickIndex,
			item.TrancheKey,
		)
		require.False(t, found)
	}
}

func TestFilledLimitOrderTrancheGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFilledLimitOrderTranche(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFilledLimitOrderTranche(ctx)),
	)
}

func TestGetNewestFilledLimitOrderTranche(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNFilledLimitOrderTrancheSameTick(keeper, ctx, 10)
	newestTranche, found := keeper.GetNewestFilledLimitOrderTranche(ctx, items[0].PairId, items[0].TokenIn, items[0].TickIndex)

	require.True(t, found)
	require.Equal(t, items[9], newestTranche)
}

func TestGetNewestFilledLimitOrderTrancheEmpty(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	_, found := keeper.GetNewestFilledLimitOrderTranche(ctx, defaultPairId, "TokenA", 0)

	require.False(t, found)
}
