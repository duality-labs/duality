package cli_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/osmoutils"
	"github.com/duality-labs/duality/osmoutils/osmocli"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/client/cli"
	"github.com/duality-labs/duality/x/incentives/types"
)

var testAddresses = osmoutils.CreateRandomAccounts(3)

// Queries ////////////////////////////////////////////////////////////////////

func TestGetCmdGetModuleStatus(t *testing.T) {
	desc, _ := cli.GetCmdGetModuleStatus()
	tcs := map[string]osmocli.QueryCliTestCase[*types.GetModuleStatusRequest]{
		"basic test": {
			ExpectedQuery: &types.GetModuleStatusRequest{},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdGetGaugeByID(t *testing.T) {
	desc, _ := cli.GetCmdGetGaugeByID()
	tcs := map[string]osmocli.QueryCliTestCase[*types.GetGaugeByIDRequest]{
		"basic test": {
			Cmd: "1", ExpectedQuery: &types.GetGaugeByIDRequest{Id: 1},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdGauges(t *testing.T) {
	desc, _ := cli.GetCmdGauges()
	tcs := map[string]osmocli.QueryCliTestCase[*types.GetGaugesRequest]{
		"test ACTIVE with pagination": {
			Cmd: "ACTIVE TokenA --offset=2",
			ExpectedQuery: &types.GetGaugesRequest{
				Status:     types.GaugeStatus_ACTIVE,
				Denom:      "TokenA",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
		"test ACTIVE_UPCOMING": {
			Cmd: "ACTIVE_UPCOMING TokenA",
			ExpectedQuery: &types.GetGaugesRequest{
				Status:     types.GaugeStatus_ACTIVE_UPCOMING,
				Denom:      "TokenA",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 0, Limit: 100},
			},
		},
		"test UPCOMING": {
			Cmd: "UPCOMING TokenA",
			ExpectedQuery: &types.GetGaugesRequest{
				Status:     types.GaugeStatus_UPCOMING,
				Denom:      "TokenA",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 0, Limit: 100},
			},
		},
		"test FINISHED": {
			Cmd: "FINISHED TokenA",
			ExpectedQuery: &types.GetGaugesRequest{
				Status:     types.GaugeStatus_FINISHED,
				Denom:      "TokenA",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 0, Limit: 100},
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdGetStakeByID(t *testing.T) {
	desc, _ := cli.GetCmdGetStakeByID()
	tcs := map[string]osmocli.QueryCliTestCase[*types.GetStakeByIDRequest]{
		"basic test": {
			Cmd: "1", ExpectedQuery: &types.GetStakeByIDRequest{StakeId: 1},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdStakes(t *testing.T) {
	desc, _ := cli.GetCmdStakes()
	tcs := map[string]osmocli.QueryCliTestCase[*types.GetStakesRequest]{
		"basic test": {
			Cmd: fmt.Sprintf("%s", testAddresses[0]),
			ExpectedQuery: &types.GetStakesRequest{
				Owner: testAddresses[0].String(),
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdFutureRewardEstimate(t *testing.T) {
	desc, _ := cli.GetCmdGetFutureRewardEstimate()
	tcs := map[string]osmocli.QueryCliTestCase[*types.GetFutureRewardEstimateRequest]{
		"basic test": {
			Cmd: fmt.Sprintf("%s [1,2,3] 1000", testAddresses[0]),
			ExpectedQuery: &types.GetFutureRewardEstimateRequest{
				Owner:    testAddresses[0].String(),
				StakeIds: []uint64{1, 2, 3},
				EndEpoch: 1000,
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

// TXS ////////////////////////////////////////////////////////////////////////

func TestNewCreateGaugeCmd(t *testing.T) {
	testTime := time.Unix(1681505514, 0).UTC()
	desc, _ := cli.NewCreateGaugeCmd()
	tcs := map[string]osmocli.TxCliTestCase[*types.MsgCreateGauge]{
		"basic test": {
			Cmd: fmt.Sprintf("TokenA<>TokenB 0 100 100TokenA,100TokenB 50 0 --from %s", testAddresses[0]),
			ExpectedMsg: &types.MsgCreateGauge{
				IsPerpetual: false,
				Owner:       testAddresses[0].String(),
				DistributeTo: types.QueryCondition{
					PairID:    &dextypes.PairID{Token0: "TokenA", Token1: "TokenB"},
					StartTick: 0,
					EndTick:   100,
				},
				Coins: sdk.NewCoins(
					sdk.NewCoin("TokenA", sdk.NewInt(100)),
					sdk.NewCoin("TokenB", sdk.NewInt(100)),
				),
				StartTime:         time.Unix(0, 0),
				NumEpochsPaidOver: 50,
				PricingTick:       0,
			},
		},
		"tests with time (RFC3339)": {
			Cmd: fmt.Sprintf("TokenA<>TokenB [-20] 20 100TokenA,100TokenB 50 0 --start-time %s --from %s", testTime.Format(time.RFC3339), testAddresses[0]),
			ExpectedMsg: &types.MsgCreateGauge{
				IsPerpetual: false,
				Owner:       testAddresses[0].String(),
				DistributeTo: types.QueryCondition{
					PairID:    &dextypes.PairID{Token0: "TokenA", Token1: "TokenB"},
					StartTick: -20,
					EndTick:   20,
				},
				Coins: sdk.NewCoins(
					sdk.NewCoin("TokenA", sdk.NewInt(100)),
					sdk.NewCoin("TokenB", sdk.NewInt(100)),
				),
				StartTime:         testTime,
				NumEpochsPaidOver: 50,
				PricingTick:       0,
			},
		},
		"tests with time (unix int)": {
			Cmd: fmt.Sprintf("TokenA<>TokenB [-20] 20 100TokenA,100TokenB 50 0 --start-time %d --from %s", testTime.Unix(), testAddresses[0]),
			ExpectedMsg: &types.MsgCreateGauge{
				IsPerpetual: false,
				Owner:       testAddresses[0].String(),
				DistributeTo: types.QueryCondition{
					PairID:    &dextypes.PairID{Token0: "TokenA", Token1: "TokenB"},
					StartTick: -20,
					EndTick:   20,
				},
				Coins: sdk.NewCoins(
					sdk.NewCoin("TokenA", sdk.NewInt(100)),
					sdk.NewCoin("TokenB", sdk.NewInt(100)),
				),
				StartTime:         testTime,
				NumEpochsPaidOver: 50,
				PricingTick:       0,
			},
		},
		"tests with perpetual": {
			Cmd: fmt.Sprintf("TokenA<>TokenB [-20] 20 100TokenA,100TokenB 50 0 --perpetual --from %s", testAddresses[0]),
			ExpectedMsg: &types.MsgCreateGauge{
				IsPerpetual: true,
				Owner:       testAddresses[0].String(),
				DistributeTo: types.QueryCondition{
					PairID:    &dextypes.PairID{Token0: "TokenA", Token1: "TokenB"},
					StartTick: -20,
					EndTick:   20,
				},
				Coins: sdk.NewCoins(
					sdk.NewCoin("TokenA", sdk.NewInt(100)),
					sdk.NewCoin("TokenB", sdk.NewInt(100)),
				),
				StartTime:         time.Unix(0, 0),
				NumEpochsPaidOver: 1,
				PricingTick:       0,
			},
		},
	}
	osmocli.RunTxTestCases(t, desc, tcs)
}

func TestNewAddToGaugeCmd(t *testing.T) {
	desc, _ := cli.NewAddToGaugeCmd()
	tcs := map[string]osmocli.TxCliTestCase[*types.MsgAddToGauge]{
		"basic test": {
			Cmd: fmt.Sprintf("1 1000TokenA --from %s", testAddresses[0]),
			ExpectedMsg: &types.MsgAddToGauge{
				Owner:   testAddresses[0].String(),
				GaugeId: 1,
				Rewards: sdk.NewCoins(sdk.NewCoin("TokenA", sdk.NewInt(1000))),
			},
		},
		"multiple tokens": {
			Cmd: fmt.Sprintf("1 1000TokenA,1TokenZ --from %s", testAddresses[0]),
			ExpectedMsg: &types.MsgAddToGauge{
				Owner:   testAddresses[0].String(),
				GaugeId: 1,
				Rewards: sdk.NewCoins(
					sdk.NewCoin("TokenA", sdk.NewInt(1000)),
					sdk.NewCoin("TokenZ", sdk.NewInt(1)),
				),
			},
		},
	}
	osmocli.RunTxTestCases(t, desc, tcs)
}

func TestNewStakeCmd(t *testing.T) {
	desc, _ := cli.NewStakeCmd()
	tcs := map[string]osmocli.TxCliTestCase[*types.MsgStake]{
		"basic test": {
			Cmd: fmt.Sprintf("1000TokenA --from %s", testAddresses[0]),
			ExpectedMsg: &types.MsgStake{
				Owner: testAddresses[0].String(),
				Coins: sdk.NewCoins(sdk.NewCoin("TokenA", sdk.NewInt(1000))),
			},
		},
		"multiple tokens": {
			Cmd: fmt.Sprintf("1000TokenA,1TokenZ --from %s", testAddresses[0]),
			ExpectedMsg: &types.MsgStake{
				Owner: testAddresses[0].String(),
				Coins: sdk.NewCoins(
					sdk.NewCoin("TokenA", sdk.NewInt(1000)),
					sdk.NewCoin("TokenZ", sdk.NewInt(1)),
				),
			},
		},
	}
	osmocli.RunTxTestCases(t, desc, tcs)
}

func TestNewUnstakeCmd(t *testing.T) {
	desc, _ := cli.NewUnstakeCmd()
	tcs := map[string]osmocli.TxCliTestCase[*types.MsgUnstake]{
		"basic test": {
			Cmd: fmt.Sprintf("--from %s", testAddresses[0]),
			ExpectedMsg: &types.MsgUnstake{
				Owner:    testAddresses[0].String(),
				Unstakes: []*types.MsgUnstake_UnstakeDescriptor{},
			},
		},
		"with coins": {
			Cmd: fmt.Sprintf("1:10TokenA 10:10TokenA,10TokenC --from %s", testAddresses[0]),
			ExpectedMsg: &types.MsgUnstake{
				Owner: testAddresses[0].String(),
				Unstakes: []*types.MsgUnstake_UnstakeDescriptor{
					{
						ID: 1,
						Coins: sdk.NewCoins(
							sdk.NewCoin("TokenA", sdk.NewInt(10)),
						),
					}, {
						ID: 10,
						Coins: sdk.NewCoins(
							sdk.NewCoin("TokenA", sdk.NewInt(10)),
							sdk.NewCoin("TokenC", sdk.NewInt(10)),
						),
					},
				},
			},
		},
	}
	osmocli.RunTxTestCases(t, desc, tcs)
}
