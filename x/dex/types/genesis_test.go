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

				NodesList: []types.Nodes{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				NodesCount: 2,
				VirtualPriceTickQueueList: []types.VirtualPriceTickQueue{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				VirtualPriceTickQueueCount: 2,
				TicksList: []types.Ticks{
					{
						Price:     "0",
						Fee:       "0",
						Direction: "0",
						OrderType: "0",
					},
					{
						Price:     "1",
						Fee:       "1",
						Direction: "1",
						OrderType: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated nodes",
			genState: &types.GenesisState{
				NodesList: []types.Nodes{
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
			desc: "invalid nodes count",
			genState: &types.GenesisState{
				NodesList: []types.Nodes{
					{
						Id: 1,
					},
				},
				NodesCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated virtualPriceTickQueue",
			genState: &types.GenesisState{
				VirtualPriceTickQueueList: []types.VirtualPriceTickQueue{
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
			desc: "invalid virtualPriceTickQueue count",
			genState: &types.GenesisState{
				VirtualPriceTickQueueList: []types.VirtualPriceTickQueue{
					{
						Id: 1,
					},
				},
				VirtualPriceTickQueueCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated ticks",
			genState: &types.GenesisState{
				TicksList: []types.Ticks{
					{
						Price:     "0",
						Fee:       "0",
						Direction: "0",
						OrderType: "0",
					},
					{
						Price:     "0",
						Fee:       "0",
						Direction: "0",
						OrderType: "0",
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
