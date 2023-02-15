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

var defaultPairId *types.PairId = &types.PairId{Token0: "TokenA", Token1: "TokenB"}

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		FeeTierList: []types.FeeTier{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		FeeTierCount: 2,
		LimitOrderTrancheUserList: []types.LimitOrderTrancheUser{
			{
				PairId:          &types.PairId{Token0: "TokenA", Token1: "TokenB"},
				Token:           "TokenB",
				TickIndex:       1,
				TrancheKey:      "0",
				Address:         "fakeAddr",
				SharesOwned:     sdk.NewInt(10),
				SharesWithdrawn: sdk.NewInt(0),
				SharesCancelled: sdk.NewInt(0),
			},
			{
				PairId:          &types.PairId{Token0: "TokenA", Token1: "TokenB"},
				Token:           "TokenA",
				TickIndex:       20,
				TrancheKey:      "0",
				Address:         "fakeAddr",
				SharesOwned:     sdk.NewInt(10),
				SharesWithdrawn: sdk.NewInt(0),
				SharesCancelled: sdk.NewInt(0),
			},
		},
		TickLiquidityList: []types.TickLiquidity{
			{
				Liquidity: &types.TickLiquidity_LimitOrderTranche{
					LimitOrderTranche: &types.LimitOrderTranche{
						PairId:           &types.PairId{Token0: "TokenA", Token1: "TokenB"},
						TokenIn:          "0",
						TickIndex:        0,
						TrancheKey:       "0",
						ReservesTokenIn:  sdk.NewInt(0),
						ReservesTokenOut: sdk.NewInt(0),
						TotalTokenIn:     sdk.NewInt(0),
						TotalTokenOut:    sdk.NewInt(0),
					},
				},
			},
			{
				Liquidity: &types.TickLiquidity_LimitOrderTranche{
					LimitOrderTranche: &types.LimitOrderTranche{
						PairId:     &types.PairId{Token0: "TokenA", Token1: "TokenB"},
						TokenIn:    "0",
						TickIndex:  0,
						TrancheKey: "0",
					},
				},
			},
		},
		FilledLimitOrderTrancheList: []types.FilledLimitOrderTranche{
			{
				PairId:     &types.PairId{Token0: "TokenA", Token1: "TokenB"},
				TokenIn:    "0",
				TickIndex:  0,
				TrancheKey: "0",
			},
			{
				PairId:     &types.PairId{Token0: "TokenA", Token1: "TokenB"},
				TokenIn:    "1",
				TickIndex:  1,
				TrancheKey: "1",
			},
		},
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FeeTierList, got.FeeTierList)
	require.Equal(t, genesisState.FeeTierCount, got.FeeTierCount)
	require.ElementsMatch(t, genesisState.LimitOrderTrancheUserList, got.LimitOrderTrancheUserList)
	require.ElementsMatch(t, genesisState.TickLiquidityList, got.TickLiquidityList)
	require.ElementsMatch(t, genesisState.FilledLimitOrderTrancheList, got.FilledLimitOrderTrancheList)
	// this line is used by starport scaffolding # genesis/test/assert
}
