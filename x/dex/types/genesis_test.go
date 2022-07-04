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

				ShareList: []types.Share{
					{
						Owner:  "0",
						Token0: "0",
						Token1: "0",
						Price:  "0",
						Fee:    0,
					},
					{
						Owner:  "1",
						Token0: "1",
						Token1: "1",
						Price:  "1",
						Fee:    1,
					},
				},
				TickList: []types.Tick{
					{
						Token0: "0",
						Token1: "0",
						Price:  "0",
						Fee:    0,
					},
					{
						Token0: "1",
						Token1: "1",
						Price:  "1",
						Fee:    1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated share",
			genState: &types.GenesisState{
				ShareList: []types.Share{
					{
						Owner:  "0",
						Token0: "0",
						Token1: "0",
						Price:  "0",
						Fee:    0,
					},
					{
						Owner:  "0",
						Token0: "0",
						Token1: "0",
						Price:  "0",
						Fee:    0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated tick",
			genState: &types.GenesisState{
				TickList: []types.Tick{
					{
						Token0: "0",
						Token1: "0",
						Price:  "0",
						Fee:    0,
					},
					{
						Token0: "0",
						Token1: "0",
						Price:  "0",
						Fee:    0,
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
