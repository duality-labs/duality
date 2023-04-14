package cli_test

import (
	"github.com/duality-labs/duality/osmoutils"
)

var testAddresses = osmoutils.CreateRandomAccounts(3)

// func TestGetCmdGauges(t *testing.T) {
// 	desc, _ := cli.GetCmdGauges()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.GetGaugesActiveUpcomingRequest]{
// 		"basic test": {
// 			Cmd: "--offset=2",
// 			ExpectedQuery: &types.GetGaugesActiveUpcomingRequest{
// 				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
// 			},
// 		},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestGetCmdToDistributeCoins(t *testing.T) {
// 	desc, _ := cli.GetCmdToDistributeCoins()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.GetModuleCoinsToBeDistributedRequest]{
// 		"basic test": {
// 			Cmd: "", ExpectedQuery: &types.GetModuleCoinsToBeDistributedRequest{},
// 		},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestGetCmdGetGaugeByID(t *testing.T) {
// 	desc, _ := cli.GetCmdGetGaugeByID()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.GetGaugeByIDRequest]{
// 		"basic test": {
// 			Cmd: "1", ExpectedQuery: &types.GetGaugeByIDRequest{Id: 1},
// 		},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestGetCmdActiveGauges(t *testing.T) {
// 	desc, _ := cli.GetCmdActiveGauges()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.ActiveGetGaugesActiveUpcomingRequest]{
// 		"basic test": {
// 			Cmd: "--offset=2",
// 			ExpectedQuery: &types.ActiveGetGaugesActiveUpcomingRequest{
// 				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
// 			}},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestGetCmdActiveGaugesPerDenom(t *testing.T) {
// 	desc, _ := cli.GetCmdActiveGaugesPerDenom()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.ActiveGaugesPerDenomRequest]{
// 		"basic test": {
// 			Cmd: "uosmo --offset=2",
// 			ExpectedQuery: &types.ActiveGaugesPerDenomRequest{
// 				Denom:      "uosmo",
// 				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
// 			}},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestGetCmdUpcomingGauges(t *testing.T) {
// 	desc, _ := cli.GetCmdUpcomingGauges()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.UpcomingGetGaugesActiveUpcomingRequest]{
// 		"basic test": {
// 			Cmd: "--offset=2",
// 			ExpectedQuery: &types.UpcomingGetGaugesActiveUpcomingRequest{
// 				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
// 			}},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestGetCmdUpcomingGaugesPerDenom(t *testing.T) {
// 	desc, _ := cli.GetCmdUpcomingGaugesPerDenom()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.UpcomingGaugesPerDenomRequest]{
// 		"basic test": {
// 			Cmd: "uosmo --offset=2",
// 			ExpectedQuery: &types.UpcomingGaugesPerDenomRequest{
// 				Denom:      "uosmo",
// 				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
// 			}},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestStakeTokensCmd(t *testing.T) {
// 	desc, _ := cli.NewStakeTokensCmd()
// 	tcs := map[string]osmocli.TxCliTestCase[*types.MsgStake]{
// 		"stake 201stake tokens for 1 day": {
// 			Cmd: "201uosmo --from=" + testAddresses[0].String(),
// 			ExpectedMsg: &types.MsgStake{
// 				Owner: testAddresses[0].String(),
// 				Coins: sdk.NewCoins(sdk.NewInt64Coin("uosmo", 201)),
// 			},
// 		},
// 	}
// 	osmocli.RunTxTestCases(t, desc, tcs)
// }

// func TestBeginUnstakingAllCmd(t *testing.T) {
// 	desc, _ := cli.NewBeginUnstakingAllCmd()
// 	tcs := map[string]osmocli.TxCliTestCase[*types.MsgBeginUnstakingAll]{
// 		"basic test": {
// 			Cmd: "--from=" + testAddresses[0].String(),
// 			ExpectedMsg: &types.MsgBeginUnstakingAll{
// 				Owner: testAddresses[0].String(),
// 			},
// 		},
// 	}
// 	osmocli.RunTxTestCases(t, desc, tcs)
// }

// func TestBeginUnstakingByIDCmd(t *testing.T) {
// 	desc, _ := cli.NewUnstakeByIDCmd()
// 	tcs := map[string]osmocli.TxCliTestCase[*types.MsgUnstake]{
// 		"basic test no coins": {
// 			Cmd: "10 --from=" + testAddresses[0].String(),
// 			ExpectedMsg: &types.MsgUnstake{
// 				Owner: testAddresses[0].String(),
// 				ID:    10,
// 				Coins: sdk.Coins(nil),
// 			},
// 		},
// 		"basic test w/ coins": {
// 			Cmd: "10 --amount=5uosmo --from=" + testAddresses[0].String(),
// 			ExpectedMsg: &types.MsgUnstake{
// 				Owner: testAddresses[0].String(),
// 				ID:    10,
// 				Coins: sdk.NewCoins(sdk.NewInt64Coin("uosmo", 5)),
// 			},
// 		},
// 	}
// 	osmocli.RunTxTestCases(t, desc, tcs)
// }

// func TestModuleBalanceCmd(t *testing.T) {
// 	desc, _ := cli.GetCmdModuleBalance()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.ModuleBalanceRequest]{
// 		"basic test": {
// 			Cmd:           "",
// 			ExpectedQuery: &types.ModuleBalanceRequest{},
// 		},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestAccountUnstakingCoinsCmd(t *testing.T) {
// 	desc, _ := cli.GetCmdAccountUnstakingCoins()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.AccountUnstakingCoinsRequest]{
// 		"basic test": {
// 			Cmd: testAddresses[0].String(),
// 			ExpectedQuery: &types.AccountUnstakingCoinsRequest{
// 				Owner: testAddresses[0].String(),
// 			},
// 		},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestCmdAccountStakedPastTime(t *testing.T) {
// 	desc, _ := cli.GetCmdAccountStakedPastTime()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.AccountStakedPastTimeRequest]{
// 		"basic test": {
// 			Cmd: testAddresses[0].String() + " 1670431012",
// 			ExpectedQuery: &types.AccountStakedPastTimeRequest{
// 				Owner:     testAddresses[0].String(),
// 				Timestamp: time.Unix(1670431012, 0),
// 			},
// 		},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestCmdAccountStakedPastTimeNotUnstakingOnly(t *testing.T) {
// 	desc, _ := cli.GetCmdAccountStakedPastTimeNotUnstakingOnly()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.AccountStakedPastTimeNotUnstakingOnlyRequest]{
// 		"basic test": {
// 			Cmd: testAddresses[0].String() + " 1670431012",
// 			ExpectedQuery: &types.AccountStakedPastTimeNotUnstakingOnlyRequest{
// 				Owner:     testAddresses[0].String(),
// 				Timestamp: time.Unix(1670431012, 0),
// 			},
// 		},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }

// func TestCmdTotalStakedByDenom(t *testing.T) {
// 	desc, _ := cli.GetCmdTotalStakedByDenom()
// 	tcs := map[string]osmocli.QueryCliTestCase[*types.StakedDenomRequest]{
// 		"basic test": {
// 			Cmd: "uosmo --min-duration=1s",
// 			ExpectedQuery: &types.StakedDenomRequest{
// 				Denom:    "uosmo",
// 				Duration: time.Second,
// 			},
// 		},
// 	}
// 	osmocli.RunQueryTestCases(t, desc, tcs)
// }
