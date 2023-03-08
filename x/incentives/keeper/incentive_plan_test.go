package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/incentives/keeper"
	"github.com/duality-labs/duality/x/incentives/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNIncentivePlan(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.IncentivePlan {
	items := make([]types.IncentivePlan, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetIncentivePlan(ctx, items[i])
	}
	return items
}

func TestIncentivePlanGet(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNIncentivePlan(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetIncentivePlan(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestIncentivePlanRemove(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNIncentivePlan(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIncentivePlan(ctx,
			item.Index,
		)
		_, found := keeper.GetIncentivePlan(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestIncentivePlanGetAll(t *testing.T) {
	keeper, ctx := keepertest.IncentivesKeeper(t)
	items := createNIncentivePlan(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllIncentivePlan(ctx)),
	)
}
