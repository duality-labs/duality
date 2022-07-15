package keeper_test

import (
	"fmt"
	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTicks(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Ticks {
	items := make([]types.Ticks, n)
	for i := range items {
		items[i].Token0 = strconv.Itoa(i)
		items[i].Token1 = strconv.Itoa(i)
		items[i].PoolsOneToZero = []*types.Pool{
			&types.Pool{
				ReserveA:    "100",
				ReserveB:    "100",
				Fee:         "3",
				Price:       "2",
				TotalShares: "200",
			},
			&types.Pool{
				ReserveA:    "100",
				ReserveB:    "100",
				Fee:         "5",
				Price:       "2",
				TotalShares: "200",
			},
		}
		items[i].PoolsZeroToOne = []*types.Pool{
			&types.Pool{
				ReserveA:    "100",
				ReserveB:    "100",
				Fee:         "3",
				Price:       "2",
				TotalShares: "200",
			},
			&types.Pool{
				ReserveA:    "100",
				ReserveB:    "100",
				Fee:         "5",
				Price:       "2",
				TotalShares: "200",
			},
		}
		keeper.SetTicks(ctx, items[i])
	}
	return items
}

func TestTicksGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)

	for _, item := range items {
		rst, found := keeper.GetTicks(ctx,
			item.Token0,
			item.Token1,
		)
		//fmt.Println(item)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTicksRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTicks(ctx,
			item.Token0,
			item.Token1,
		)
		_, found := keeper.GetTicks(ctx,
			item.Token0,
			item.Token1,
		)
		require.False(t, found)
	}
}

func TestTicksGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTicks(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTicks(ctx)),
	)
}

func TestInit(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)

	pools := []*types.Pool{
		&types.Pool{
			ReserveA:    "100",
			ReserveB:    "100",
			Fee:         ".5",
			Price:       "2",
			TotalShares: "200",
			Index:       0,
		},
		&types.Pool{
			ReserveA:    "100",
			ReserveB:    "100",
			Fee:         ".2",
			Price:       "2",
			TotalShares: "200",
			Index:       0,
		},
	}
	fmt.Println(pools)
	keeper.Init(&pools)
	fmt.Println(pools)

	newPool := &types.Pool{
		ReserveA:    "100",
		ReserveB:    "100",
		Fee:         ".2",
		Price:       "3",
		TotalShares: "200",
		Index:       0,
	}

	keeper.Push(&pools, newPool)
	fmt.Println(pools)

	removedPool := keeper.Pop(&pools)
	fmt.Println(pools)
	fmt.Println("Popped Pool", removedPool)
	removedPool = keeper.Remove(&pools, 0)
	fmt.Println("Post Remove: ", pools)
	fmt.Println(removedPool)
	keeper.Push(&pools, newPool)
	keeper.Push(&pools, newPool)
	fmt.Println("HEAD", pools[1].Index)
	fmt.Println("Post Pushes:", pools)
	keeper.Update(&pools, pools[1], "100", "100", ".3", "200", "4")
	fmt.Println("Post Update: ", pools)
	removedPool = keeper.Pop(&pools)
	fmt.Println(pools)
	fmt.Println("Popped Pool", removedPool)
	removedPool = keeper.Pop(&pools)
	fmt.Println(pools)
	fmt.Println("Popped Pool", removedPool)

	newPool2 := &types.Pool{
		ReserveA:    "100",
		ReserveB:    "100",
		Fee:         ".2",
		Price:       "4",
		TotalShares: "200",
		Index:       0,
	}
	newPool3 := &types.Pool{
		ReserveA:    "100",
		ReserveB:    "100",
		Fee:         ".5",
		Price:       "3",
		TotalShares: "200",
		Index:       0,
	}
	fmt.Println("New Pushes")
	fmt.Println("")
	keeper.Push(&pools, newPool)
	keeper.Push(&pools, newPool2)
	keeper.Push(&pools, newPool3)
	keeper.Push(&pools, newPool2)
	fmt.Println(pools)
	removedPool = keeper.Pop(&pools)
	fmt.Println(pools)
	fmt.Println("Popped Pool", removedPool)
	removedPool = keeper.Pop(&pools)
	fmt.Println(pools)
	fmt.Println("Popped Pool", removedPool)
	removedPool = keeper.Pop(&pools)
	fmt.Println(pools)
	fmt.Println("Popped Pool", removedPool)
	removedPool = keeper.Pop(&pools)
	fmt.Println(pools)
	fmt.Println("Popped Pool", removedPool)

	if _, found := keeper.GetTicks(ctx, "tokenA", "tokenB"); found == false {
		keeper.SetTicks(ctx, types.Ticks{
			Token0:         "tokenA",
			Token1:         "tokenB",
			PoolsZeroToOne: pools,
			PoolsOneToZero: pools,
		})
	}

	tick, found := keeper.GetTicks(ctx, "tokenA", "tokenB")

	require.True(t, found)

	fmt.Println(tick)
	fmt.Println(tick.PoolsZeroToOne)
	removedPool = keeper.Pop(&tick.PoolsZeroToOne)
	fmt.Println(&tick.PoolsZeroToOne)
	fmt.Println("Popped Pool", removedPool)

	keeper.SetTicks(ctx, types.Ticks{
		Token0:         "tokenA",
		Token1:         "tokenB",
		PoolsZeroToOne: tick.PoolsZeroToOne,
		PoolsOneToZero: tick.PoolsOneToZero,
	})

	tick1, found := keeper.GetTicks(ctx, "tokenA", "tokenB")
	require.True(t, found)
	fmt.Println(tick1.PoolsZeroToOne)

	_ = ctx
}
