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

func createNInactiveLimitOrderTranche(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LimitOrderTranche {
	items := make([]types.LimitOrderTranche, n)
	for i := range items {
		items[i] = types.LimitOrderTranche{
			TrancheKey:         strconv.Itoa(i),
			TradePairID:        &types.TradePairID{MakerDenom: "TokenA", TakerDenom: "TokenB"},
			TickIndex:          int64(i),
			TotalMakerDenom:    sdk.ZeroInt(),
			TotalTakerDenom:    sdk.ZeroInt(),
			ReservesTakerDenom: sdk.ZeroInt(),
			ReservesMakerDenom: sdk.ZeroInt(),
		}
		keeper.SetInactiveLimitOrderTranche(ctx, &items[i])
	}

	return items
}

func TestInactiveLimitOrderTrancheGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNInactiveLimitOrderTranche(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetInactiveLimitOrderTranche(ctx,
			item.TradePairID,
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

func TestInactiveLimitOrderTrancheRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNInactiveLimitOrderTranche(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInactiveLimitOrderTranche(ctx,
			item.TradePairID,
			item.TickIndex,
			item.TrancheKey,
		)
		_, found := keeper.GetInactiveLimitOrderTranche(ctx,
			item.TradePairID,
			item.TickIndex,
			item.TrancheKey,
		)
		require.False(t, found)
	}
}

func TestInactiveLimitOrderTrancheGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNInactiveLimitOrderTranche(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllInactiveLimitOrderTranche(ctx)),
	)
}
