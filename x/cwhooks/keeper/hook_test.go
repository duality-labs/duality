package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/cwhooks/keeper"
	"github.com/duality-labs/duality/x/cwhooks/types"
	"github.com/stretchr/testify/require"
)

func createNHook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Hook {
	items := make([]types.Hook, n)
	for i := range items {
		items[i].TriggerKey = "eventName"
		items[i].TriggerValue = fmt.Sprintf("event_value%d", i)
		items[i].Id = keeper.AppendHook(ctx, items[i])
	}

	return items
}

func TestHookGet(t *testing.T) {
	keeper, ctx := keepertest.CWHooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHook(ctx, item.TriggerKey, item.TriggerValue, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHookGetByID(t *testing.T) {
	keeper, ctx := keepertest.CWHooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHookByID(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHookRemove(t *testing.T) {
	keeper, ctx := keepertest.CWHooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHook(ctx, item)
		_, found := keeper.GetHookByID(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHookRemoveByID(t *testing.T) {
	keeper, ctx := keepertest.CWHooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHookByID(ctx, item.Id)
		_, found := keeper.GetHookByID(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHookGetAll(t *testing.T) {
	keeper, ctx := keepertest.CWHooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHook(ctx)),
	)
}

func TestHookCount(t *testing.T) {
	keeper, ctx := keepertest.CWHooksKeeper(t)
	items := createNHook(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHookCount(ctx))
}

// JCP TODO add tests for hook by creator

func TestGetAllHooksForKeyValue(t *testing.T) {
	keeper, ctx := keepertest.CWHooksKeeper(t)
	triggerKey := "k1"
	triggerValue := "v1"
	items := []types.Hook{
		// matching hooks
		{
			TriggerKey:      triggerKey,
			TriggerValue:    triggerValue,
			ContractAddress: "addr",
		},
		{
			TriggerKey:      triggerKey,
			TriggerValue:    triggerValue,
			ContractAddress: "addr2",
		},
		// No-matching hook
		{
			TriggerKey:      "k2",
			TriggerValue:    "v2",
			ContractAddress: "addr",
		},
	}
	for i, hook := range items {
		items[i].Id = keeper.AppendHook(ctx, hook)
	}

	foundHooks := keeper.GetAllHooksForKeyValue(ctx, triggerKey, triggerValue)
	require.ElementsMatch(t,
		nullify.Fill(items[0:2]),
		nullify.Fill(foundHooks),
	)
}
