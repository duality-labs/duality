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

				TicksList: []types.Ticks{
					{
						Price:     "0",
						Fee:       "0",
						OrderType: "0",
					},
					{
						Price:     "1",
						Fee:       "1",
						OrderType: "1",
					},
				},

				PairsList: []types.Pairs{
					{
						Token0: "0",
						Token1: "0",
					},
					{
						Token0: "1",
						Token1: "1",
					},
				},
				IndexQueueList: []types.IndexQueue{
					{
						Index: 0,
					},
					{
						Index: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},

		{
			desc: "duplicated ticks",
			genState: &types.GenesisState{
				TicksList: []types.Ticks{
					{
						Price:     "0",
						Fee:       "0",
						OrderType: "0",
					},
					{
						Price:     "0",
						Fee:       "0",
						OrderType: "0",
					},
				},
			},
			valid: false,
		},

		{
			desc: "duplicated pairs",
			genState: &types.GenesisState{
				PairsList: []types.Pairs{
					{
						Token0: "0",
						Token1: "0",
					},
					{
						Token0: "0",
						Token1: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated virtualPriceQueue",
			genState: &types.GenesisState{
				IndexQueueList: []types.IndexQueue{
					{
						Index: 0,
					},
					{
						Index: 0,
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
