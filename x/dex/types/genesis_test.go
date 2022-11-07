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

				TickObjectList: []types.TickObject{
					{
						TickIndex: 0,
					},
					{
						TickIndex: 1,
					},
				},
				PairObjectList: []types.PairObject{
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
				TokenObjectList: []types.TokenObject{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				SharesList: []types.Shares{
					{
						Address:   "0",
						PairId:    "0",
						TickIndex: 0,
						FeeIndex:  0,
					},
					{
						Address:   "1",
						PairId:    "1",
						TickIndex: 1,
						FeeIndex:  1,
					},
				},
				FeeListList: []types.FeeList{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				FeeListCount: 2,

				LimitOrderPoolUserShareObjectList: []types.LimitOrderPoolUserShareObject{
					{
						Count:   0,
						Address: "0",
					},
					{
						Count:   1,
						Address: "1",
					},
				},
				LimitOrderPoolUserSharesWithdrawnList: []types.LimitOrderPoolUserSharesWithdrawn{
					{
						Count:   0,
						Address: "0",
					},
					{
						Count:   1,
						Address: "1",
					},
				},
				LimitOrderPoolTotalSharesObjectList: []types.LimitOrderPoolTotalSharesObject{
					{
						Count: 0,
					},
					{
						Count: 1,
					},
				},
				LimitOrderPoolReserveObjectList: []types.LimitOrderPoolReserveObject{
					{
						Count: 0,
					},
					{
						Count: 1,
					},
				},
				LimitOrderPoolFillObjectList: []types.LimitOrderPoolFillObject{
					{
						Count: 0,
					},
					{
						Count: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated tickObject",
			genState: &types.GenesisState{
				TickObjectList: []types.TickObject{
					{
						TickIndex: 0,
					},
					{
						TickIndex: 1,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pairObject",
			genState: &types.GenesisState{
				PairObjectList: []types.PairObject{
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
			desc: "duplicated tokenObject",
			genState: &types.GenesisState{
				TokenObjectList: []types.TokenObject{
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
			desc: "duplicated shares",
			genState: &types.GenesisState{
				SharesList: []types.Shares{
					{
						Address:   "0",
						PairId:    "0",
						TickIndex: 0,
						FeeIndex:  0,
					},
					{
						Address:   "0",
						PairId:    "0",
						TickIndex: 0,
						FeeIndex:  0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated feeList",
			genState: &types.GenesisState{
				FeeListList: []types.FeeList{
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
			desc: "invalid feeList count",
			genState: &types.GenesisState{
				FeeListList: []types.FeeList{
					{
						Id: 1,
					},
				},
				FeeListCount: 0,
			},
			valid: false,
		},

		{
			desc: "duplicated limitOrderPoolUserShareObject",
			genState: &types.GenesisState{
				LimitOrderPoolUserShareObjectList: []types.LimitOrderPoolUserShareObject{
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
			desc: "duplicated limitOrderPoolUserSharesWithdrawn",
			genState: &types.GenesisState{
				LimitOrderPoolUserSharesWithdrawnList: []types.LimitOrderPoolUserSharesWithdrawn{
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
			desc: "duplicated limitOrderPoolTotalSharesObject",
			genState: &types.GenesisState{
				LimitOrderPoolTotalSharesObjectList: []types.LimitOrderPoolTotalSharesObject{
					{
						Count: 0,
					},
					{
						Count: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated limitOrderPoolReserveObject",
			genState: &types.GenesisState{
				LimitOrderPoolReserveObjectList: []types.LimitOrderPoolReserveObject{
					{
						Count: 0,
					},
					{
						Count: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated limitOrderPoolFillObject",
			genState: &types.GenesisState{
				LimitOrderPoolFillObjectList: []types.LimitOrderPoolFillObject{
					{
						Count: 0,
					},
					{
						Count: 0,
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
