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

func createNTickObject(keeper *keeper.Keeper, ctx sdk.Context, pairId string, n int) []types.TickObject {
	items := make([]types.TickObject, n)
	for i := range items {

		items[i].TickData = &types.TickDataType{
			Reserve0AndShares: []*types.Reserve0AndSharesType{
				{Reserve0: sdk.OneDec(),
					TotalShares: sdk.ZeroDec(),
				}},
			Reserve1: []sdk.Dec{sdk.ZeroDec()},
		}

		items[i].TickIndex = int64(i)

		// testTickObject :=  &types.TickObject{0, &types.TickDataType{Reserve0AndShares: []*types.Reserve0AndSharesType{
		// 	{Reserve0: sdk.OneDec(),
		// 	TotalShares: sdk.ZeroDec(),
		// }},
		// Reserve1: []sdk.Dec{sdk.ZeroDec()},
		// }}

		keeper.SetTickObject(ctx, pairId, items[i])
	}

	return items
}

func TestTickObjectGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTickObject(keeper, ctx, "TokenB/TokenA", 10)
	for _, item := range items {
		rst, found := keeper.GetTickObject(ctx, "TokenB/TokenA",
			item.TickIndex,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTickObjectRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTickObject(keeper, ctx, "TokenB/TokenA", 10)
	for _, item := range items {
		keeper.RemoveTickObject(ctx,
			"TokenB/TokenA",
			item.TickIndex,
		)
		_, found := keeper.GetTickObject(ctx,
			"TokenB/TokenA",
			item.TickIndex,
		)
		require.False(t, found)
	}
}

func TestTickObjectGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTickObject(keeper, ctx, "TokenB/TokenA", 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTickObject(ctx)),
	)
}
