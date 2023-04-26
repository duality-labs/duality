package types_test

import (
	"testing"

	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				LimitOrderTrancheUserList: []types.LimitOrderTrancheUser{
					{
						TrancheKey: "0",
						Address:    "0",
						PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
					},
					{
						TrancheKey: "1",
						Address:    "1",
						PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
					},
				},
				TickLiquidityList: []types.TickLiquidity{
					{
						Liquidity: &types.TickLiquidity_LimitOrderTranche{
							LimitOrderTranche: &types.LimitOrderTranche{
								PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
								TokenIn:    "0",
								TickIndex:  0,
								TrancheKey: "0",
							},
						},
					},
					{
						Liquidity: &types.TickLiquidity_PoolReserves{
							PoolReserves: &types.PoolReserves{
								PairID:    &types.PairID{Token0: "TokenA", Token1: "TokenB"},
								TokenIn:   "0",
								TickIndex: 0,
								Fee:       0,
							},
						},
					},
				},
				InactiveLimitOrderTrancheList: []types.LimitOrderTranche{
					{
						PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
						TokenIn:    "0",
						TickIndex:  0,
						TrancheKey: "0",
					},
					{
						PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
						TokenIn:    "1",
						TickIndex:  1,
						TrancheKey: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated LimitOrderTrancheUser",
			genState: &types.GenesisState{
				LimitOrderTrancheUserList: []types.LimitOrderTrancheUser{
					{
						TrancheKey: "0",
						Address:    "0",
						PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
					},
					{
						TrancheKey: "0",
						Address:    "0",
						PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated tickLiquidity",
			genState: &types.GenesisState{
				TickLiquidityList: []types.TickLiquidity{
					{
						Liquidity: &types.TickLiquidity_LimitOrderTranche{
							LimitOrderTranche: &types.LimitOrderTranche{
								PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
								TokenIn:    "0",
								TickIndex:  0,
								TrancheKey: "0",
							},
						},
					},
					{
						Liquidity: &types.TickLiquidity_LimitOrderTranche{
							LimitOrderTranche: &types.LimitOrderTranche{
								PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
								TokenIn:    "0",
								TickIndex:  0,
								TrancheKey: "0",
							},
						},
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated inactiveLimitOrderTranche",
			genState: &types.GenesisState{
				InactiveLimitOrderTrancheList: []types.LimitOrderTranche{
					{
						PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
						TokenIn:    "0",
						TickIndex:  0,
						TrancheKey: "0",
					},
					{
						PairID:     &types.PairID{Token0: "TokenA", Token1: "TokenB"},
						TokenIn:    "0",
						TickIndex:  0,
						TrancheKey: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
