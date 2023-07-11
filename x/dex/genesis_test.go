package dex_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: *types.DefaultParams(),
		LimitOrderTrancheUserList: []*types.LimitOrderTrancheUser{
			{
				TradePairID:           &types.TradePairID{TakerDenom: "TokenA", MakerDenom: "TokenB"},
				TickIndexTakerToMaker: 1,
				TrancheKey:            "0",
				Address:               "fakeAddr",
				SharesOwned:           sdk.NewInt(10),
				SharesWithdrawn:       sdk.NewInt(0),
				SharesCancelled:       sdk.NewInt(0),
			},
			{
				TradePairID:           &types.TradePairID{TakerDenom: "TokenB", MakerDenom: "TokenA"},
				TickIndexTakerToMaker: 20,
				TrancheKey:            "0",
				Address:               "fakeAddr",
				SharesOwned:           sdk.NewInt(10),
				SharesWithdrawn:       sdk.NewInt(0),
				SharesCancelled:       sdk.NewInt(0),
			},
		},
		TickLiquidityList: []*types.TickLiquidity{
			{
				Liquidity: &types.TickLiquidity_LimitOrderTranche{
					LimitOrderTranche: types.MustNewLimitOrderTranche(
						"TokenB",
						"TokenA",
						"0",
						0,
						sdk.ZeroInt(),
						sdk.ZeroInt(),
						sdk.ZeroInt(),
						sdk.ZeroInt(),
					),
				},
			},
			{
				Liquidity: &types.TickLiquidity_LimitOrderTranche{
					LimitOrderTranche: types.MustNewLimitOrderTranche(
						"TokenB",
						"TokenA",
						"0",
						0,
						sdk.ZeroInt(),
						sdk.ZeroInt(),
						sdk.ZeroInt(),
						sdk.ZeroInt(),
					),
				},
			},
		},
		InactiveLimitOrderTrancheList: []*types.LimitOrderTranche{
			{
				Key: &types.LimitOrderTrancheKey{
					TradePairID:           &types.TradePairID{TakerDenom: "TokenA", MakerDenom: "TokenB"},
					TickIndexTakerToMaker: 0,
					TrancheKey:            "0",
				},
			},
			{
				Key: &types.LimitOrderTrancheKey{
					TradePairID:           &types.TradePairID{TakerDenom: "TokenA", MakerDenom: "TokenB"},
					TickIndexTakerToMaker: 1,
					TrancheKey:            "1",
				},
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LimitOrderTrancheUserList, got.LimitOrderTrancheUserList)
	require.ElementsMatch(t, genesisState.TickLiquidityList, got.TickLiquidityList)
	require.ElementsMatch(t, genesisState.InactiveLimitOrderTrancheList, got.InactiveLimitOrderTrancheList)
	// this line is used by starport scaffolding # genesis/test/assert
}
