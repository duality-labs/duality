package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/duality-labs/duality/osmoutils/osmocli"
	"github.com/duality-labs/duality/x/incentives/types"
)

// GetQueryCmd returns the query commands for this module.
func GetQueryCmd() *cobra.Command {
	// group incentives queries under a subcommand
	cmd := osmocli.QueryIndexCmd(types.ModuleName)
	qcGetter := types.NewQueryClient
	osmocli.AddQueryCmd(cmd, qcGetter, GetCmdGetModuleStatus)
	osmocli.AddQueryCmd(cmd, qcGetter, GetCmdGetGaugeByID)
	osmocli.AddQueryCmd(cmd, qcGetter, GetCmdGauges)
	osmocli.AddQueryCmd(cmd, qcGetter, GetCmdGetLockByID)
	osmocli.AddQueryCmd(cmd, qcGetter, GetCmdLocks)
	osmocli.AddQueryCmd(cmd, qcGetter, GetCmdGetFutureRewardEstimate)

	return cmd
}

// GetCmdGetModuleStatus returns status of incentive module.
func GetCmdGetModuleStatus() (*osmocli.QueryDescriptor, *types.GetModuleStatusRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "module-status",
		Short: "Query module status.",
		Long:  `{{.Short}}`,
	}, &types.GetModuleStatusRequest{}
}

// GetCmdGetGaugeByID returns a gauge by ID.
func GetCmdGetGaugeByID() (*osmocli.QueryDescriptor, *types.GetGaugeByIDRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "gauge-by-id [id]",
		Short: "Query gauge by id.",
		Long:  `{{.Short}}{{.ExampleHeader}} gauge-by-id 1`,
	}, &types.GetGaugeByIDRequest{}
}

func parseGaugeStatus(arg string, _ *pflag.FlagSet) (any, osmocli.FieldReadLocation, error) {
	gaugeStatusInt, ok := types.GaugeStatus_value[arg]
	if !ok {
		return 0, osmocli.UsedArg, types.ErrInvalidGaugeStatus
	}
	gaugeStatus := types.GaugeStatus(gaugeStatusInt)

	return gaugeStatus, osmocli.UsedArg, nil
}

// GetCmdGauges returns all gauges for a given status and denom.
func GetCmdGauges() (*osmocli.QueryDescriptor, *types.GetGaugesRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "list-gauges [status] [denom]",
		Short: "Query gauges",
		Long:  `{{.Short}}{{.ExampleHeader}} list-gauges UPCOMING TokenA`,
		CustomFieldParsers: map[string]osmocli.CustomFieldParserFn{
			"Status": parseGaugeStatus,
		},
	}, &types.GetGaugesRequest{}
}

// GetCmdGetLockByID returns a lock by ID.
func GetCmdGetLockByID() (*osmocli.QueryDescriptor, *types.GetLockByIDRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "lock-by-id [LockId]",
		Short: "Query lock by id.",
		Long:  `{{.Short}}{{.ExampleHeader}} lock-by-id 1`,
	}, &types.GetLockByIDRequest{}
}

func parseLockStatus(arg string, _ *pflag.FlagSet) (any, osmocli.FieldReadLocation, error) {
	lockStatusInt, ok := types.LockStatus_value[arg]
	if !ok {
		return 0, osmocli.UsedArg, types.ErrInvalidLockStatus
	}
	lockStatus := types.LockStatus(lockStatusInt)

	return lockStatus, osmocli.UsedArg, nil
}

// GetCmdLocks returns all gauges for a given status and owner.
func GetCmdLocks() (*osmocli.QueryDescriptor, *types.GetLocksRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "list-locks [status] [owner]",
		Short: "Query locks",
		Long:  `{{.Short}}{{.ExampleHeader}} list-locks ALL cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx`,
		CustomFieldParsers: map[string]osmocli.CustomFieldParserFn{
			"Status": parseLockStatus,
		},
	}, &types.GetLocksRequest{}
}

// GetCmdGetFutureRewardsEstimate returns a rewards estimate for a given set of locks.
func GetCmdGetFutureRewardEstimate() (*osmocli.QueryDescriptor, *types.GetFutureRewardEstimateRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "reward-estimate [owner] [lockIDs] [endEpoch]",
		Short: "Get rewards estimate for set of locks",
		Long:  `{{.Short}}{{.ExampleHeader}} reward-estimate cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx [1,2,3] 1681450672`,
		CustomFieldParsers: map[string]osmocli.CustomFieldParserFn{
			"LockIDs": osmocli.ParseUintArray,
		},
	}, &types.GetFutureRewardEstimateRequest{}
}
