package types_test

import (
	"testing"

	"github.com/NicholasDotSol/duality/x/dex/types"
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

				TickList: []types.Tick{
					{
						TickIndex: 0,
					},
					{
						TickIndex: 1,
					},
				},
				TradingPairList: []types.TradingPair{
					{
						PairId: "0",
					},
					{
						PairId: "1",
					},
				},
				TokensList: []types.Tokens{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				TokensCount: 2,
				TokenMapList: []types.TokenMap{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
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
						Count:   0,
						Address: "0",
					},
					{
						Count:   1,
						Address: "1",
					},
				},
				LimitOrderTrancheList: []types.LimitOrderTranche{
					{
						TickIndex: 0,
					},
					{
						TickIndex: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated Tick",
			genState: &types.GenesisState{
				TickList: []types.Tick{
					{
						TickIndex: 0,
					},
					{
						TickIndex: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated TradingPair",
			genState: &types.GenesisState{
				TradingPairList: []types.TradingPair{
					{
						PairId: "0",
					},
					{
						PairId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated tokens",
			genState: &types.GenesisState{
				TokensList: []types.Tokens{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid tokens count",
			genState: &types.GenesisState{
				TokensList: []types.Tokens{
					{
						Id: 1,
					},
				},
				TokensCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated tokenMap",
			genState: &types.GenesisState{
				TokenMapList: []types.TokenMap{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated FeeTier",
			genState: &types.GenesisState{
				FeeTierList: []types.FeeTier{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid FeeTier count",
			genState: &types.GenesisState{
				FeeTierList: []types.FeeTier{
					{
						Id: 1,
					},
				},
				FeeTierCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated LimitOrderTrancheUser",
			genState: &types.GenesisState{
				LimitOrderTrancheUserList: []types.LimitOrderTrancheUser{
					{
						Count:   0,
						Address: "0",
					},
					{
						Count:   0,
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated LimitOrderTranche",
			genState: &types.GenesisState{
				LimitOrderTrancheList: []types.LimitOrderTranche{
					{
						TickIndex: 0,
					},
					{
						TickIndex: 0,
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
