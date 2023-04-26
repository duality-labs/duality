package cli_test

import (
	"testing"

	"github.com/duality-labs/duality/utils/dcli"
	"github.com/duality-labs/duality/x/epochs/client/cli"
	"github.com/duality-labs/duality/x/epochs/types"
)

func TestGetCmdCurrentEpoch(t *testing.T) {
	desc, _ := cli.GetCmdCurrentEpoch()
	tcs := map[string]dcli.QueryCliTestCase[*types.QueryCurrentEpochRequest]{
		"basic test": {
			Cmd: "day",
			ExpectedQuery: &types.QueryCurrentEpochRequest{
				Identifier: "day",
			},
		},
	}
	dcli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdEpochsInfo(t *testing.T) {
	desc, _ := cli.GetCmdEpochInfos()
	tcs := map[string]dcli.QueryCliTestCase[*types.QueryEpochsInfoRequest]{
		"basic test": {
			Cmd:           "",
			ExpectedQuery: &types.QueryEpochsInfoRequest{},
		},
	}
	dcli.RunQueryTestCases(t, desc, tcs)
}
