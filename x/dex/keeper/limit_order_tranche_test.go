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

func createNLimitOrderTranches(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LimitOrderTranche {
	items := make([]types.LimitOrderTranche, n)
	for i := range items {
		tranche := &types.LimitOrderTranche{
			PairID:           &types.PairID{Token0: "TokenA", Token1: "TokenB"},
			TokenIn:          "TokenA",
			TickIndex:        int64(i),
			TrancheKey:       strconv.Itoa(i),
			ReservesTokenIn:  sdk.NewInt(10),
			ReservesTokenOut: sdk.NewInt(10),
			TotalTokenIn:     sdk.NewInt(10),
			TotalTokenOut:    sdk.NewInt(10),
		}
		keeper.SetLimitOrderTranche(ctx, *tranche)
		items[i] = *tranche
	}

	return items
}

func TestGetLimitOrderTranche(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTranches(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLimitOrderTranche(ctx,
			item.PairID,
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

func TestRemoveLimitOrderTranche(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTranches(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLimitOrderTranche(ctx, item)
		_, found := keeper.GetLimitOrderTranche(ctx,
			item.PairID,
			item.TokenIn,
			item.TickIndex,
			item.TrancheKey,
		)
		require.False(t, found)
	}
}
