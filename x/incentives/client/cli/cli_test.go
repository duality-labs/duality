package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/osmoutils"
	"github.com/duality-labs/duality/osmoutils/osmocli"
	"github.com/duality-labs/duality/x/incentives/client/cli"
	"github.com/duality-labs/duality/x/incentives/types"
)

var testAddresses = osmoutils.CreateRandomAccounts(3)

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

func TestGetCmdGetLockByID(t *testing.T) {
	desc, _ := cli.GetCmdGetLockByID()
	tcs := map[string]osmocli.QueryCliTestCase[*types.GetLockByIDRequest]{
		"basic test": {
			Cmd: "1", ExpectedQuery: &types.GetLockByIDRequest{LockId: 1},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdLocks(t *testing.T) {
	desc, _ := cli.GetCmdLocks()
	tcs := map[string]osmocli.QueryCliTestCase[*types.GetLocksRequest]{
		"test ALL with pagination": {
			Cmd: fmt.Sprintf("ALL %s --offset=2", testAddresses[0]),
			ExpectedQuery: &types.GetLocksRequest{
				Status:     types.LockStatus_ALL,
				Owner:      testAddresses[0].String(),
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
		"test UNLOCKING": {
			Cmd: fmt.Sprintf("UNLOCKING %s", testAddresses[0]),
			ExpectedQuery: &types.GetLocksRequest{
				Status:     types.LockStatus_UNLOCKING,
				Owner:      testAddresses[0].String(),
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 0, Limit: 100},
			},
		},
		"test NOT_UNLOCKING": {
			Cmd: fmt.Sprintf("NOT_UNLOCKING %s", testAddresses[0]),
			ExpectedQuery: &types.GetLocksRequest{
				Status:     types.LockStatus_NOT_UNLOCKING,
				Owner:      testAddresses[0].String(),
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 0, Limit: 100},
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}
