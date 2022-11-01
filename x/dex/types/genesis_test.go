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

				TickMapList: []types.TickMap{
					{
						TickIndex: 0,
					},
					{
						TickIndex: 1,
					},
				},
				PairMapList: []types.PairMap{
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
				EdgeRowList: []types.EdgeRow{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				EdgeRowCount: 2,
				AdjanceyMatrixList: []types.AdjanceyMatrix{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				AdjanceyMatrixCount: 2,
				LimitOrderPoolUserShareMapList: []types.LimitOrderPoolUserShareMap{
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
				LimitOrderPoolTotalSharesMapList: []types.LimitOrderPoolTotalSharesMap{
					{
						Count: 0,
					},
					{
						Count: 1,
					},
				},
				LimitOrderPoolReserveMapList: []types.LimitOrderPoolReserveMap{
					{
						Count: 0,
					},
					{
						Count: 1,
					},
				},
				LimitOrderPoolFillMapList: []types.LimitOrderPoolFillMap{
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
			desc: "duplicated tickMap",
			genState: &types.GenesisState{
				TickMapList: []types.TickMap{
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
			desc: "duplicated pairMap",
			genState: &types.GenesisState{
				PairMapList: []types.PairMap{
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
			desc: "duplicated edgeRow",
			genState: &types.GenesisState{
				EdgeRowList: []types.EdgeRow{
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
			desc: "invalid edgeRow count",
			genState: &types.GenesisState{
				EdgeRowList: []types.EdgeRow{
					{
						Id: 1,
					},
				},
				EdgeRowCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated adjanceyMatrix",
			genState: &types.GenesisState{
				AdjanceyMatrixList: []types.AdjanceyMatrix{
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
			desc: "invalid adjanceyMatrix count",
			genState: &types.GenesisState{
				AdjanceyMatrixList: []types.AdjanceyMatrix{
					{
						Id: 1,
					},
				},
				AdjanceyMatrixCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated limitOrderPoolUserShareMap",
			genState: &types.GenesisState{
				LimitOrderPoolUserShareMapList: []types.LimitOrderPoolUserShareMap{
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
			desc: "duplicated limitOrderPoolTotalSharesMap",
			genState: &types.GenesisState{
				LimitOrderPoolTotalSharesMapList: []types.LimitOrderPoolTotalSharesMap{
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
			desc: "duplicated limitOrderPoolReserveMap",
			genState: &types.GenesisState{
				LimitOrderPoolReserveMapList: []types.LimitOrderPoolReserveMap{
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
			desc: "duplicated limitOrderPoolFillMap",
			genState: &types.GenesisState{
				LimitOrderPoolFillMapList: []types.LimitOrderPoolFillMap{
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
