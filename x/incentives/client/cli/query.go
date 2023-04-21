package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/duality-labs/duality/utils/dcli"
	"github.com/duality-labs/duality/x/incentives/types"
)

// GetQueryCmd returns the query commands for this module.
func GetQueryCmd() *cobra.Command {
	// group incentives queries under a subcommand
	cmd := dcli.QueryIndexCmd(types.ModuleName)
	qcGetter := types.NewQueryClient
	dcli.AddQueryCmd(cmd, qcGetter, GetCmdGetModuleStatus)
	dcli.AddQueryCmd(cmd, qcGetter, GetCmdGetGaugeByID)
	dcli.AddQueryCmd(cmd, qcGetter, GetCmdGauges)
	dcli.AddQueryCmd(cmd, qcGetter, GetCmdGetStakeByID)
	dcli.AddQueryCmd(cmd, qcGetter, GetCmdStakes)
	dcli.AddQueryCmd(cmd, qcGetter, GetCmdGetFutureRewardEstimate)

	return cmd
}

// GetCmdGetModuleStatus returns status of incentive module.
func GetCmdGetModuleStatus() (*dcli.QueryDescriptor, *types.GetModuleStatusRequest) {
	return &dcli.QueryDescriptor{
		Use:   "module-status",
		Short: "Query module status.",
		Long:  `{{.Short}}`,
	}, &types.GetModuleStatusRequest{}
}

// GetCmdGetGaugeByID returns a gauge by ID.
func GetCmdGetGaugeByID() (*dcli.QueryDescriptor, *types.GetGaugeByIDRequest) {
	return &dcli.QueryDescriptor{
		Use:   "gauge-by-id [id]",
		Short: "Query gauge by id.",
		Long:  `{{.Short}}{{.ExampleHeader}} gauge-by-id 1`,
	}, &types.GetGaugeByIDRequest{}
}

func parseGaugeStatus(arg string, _ *pflag.FlagSet) (any, dcli.FieldReadLocation, error) {
	gaugeStatusInt, ok := types.GaugeStatus_value[arg]
	if !ok {
		return 0, dcli.UsedArg, types.ErrInvalidGaugeStatus
	}
	gaugeStatus := types.GaugeStatus(gaugeStatusInt)

	return gaugeStatus, dcli.UsedArg, nil
}

// GetCmdGauges returns all gauges for a given status and denom.
func GetCmdGauges() (*dcli.QueryDescriptor, *types.GetGaugesRequest) {
	return &dcli.QueryDescriptor{
		Use:   "list-gauges [status] [denom]",
		Short: "Query gauges",
		Long:  `{{.Short}}{{.ExampleHeader}} list-gauges UPCOMING DualityPoolShares-stake-token-t0-f1`,
		CustomFieldParsers: map[string]dcli.CustomFieldParserFn{
			"Status": parseGaugeStatus,
		},
	}, &types.GetGaugesRequest{}
}

// GetCmdGetStakeByID returns a lock by ID.
func GetCmdGetStakeByID() (*dcli.QueryDescriptor, *types.GetStakeByIDRequest) {
	return &dcli.QueryDescriptor{
		Use:   "stake-by-id [stakeID]",
		Short: "Query stake by id.",
		Long:  `{{.Short}}{{.ExampleHeader}} Stake-by-id 1`,
	}, &types.GetStakeByIDRequest{}
}

// GetCmdStakes returns all gauges for a given status and owner.
func GetCmdStakes() (*dcli.QueryDescriptor, *types.GetStakesRequest) {
	return &dcli.QueryDescriptor{
		Use:   "list-stakes [owner]",
		Short: "Query stakes",
		Long:  `{{.Short}}{{.ExampleHeader}} list-stakes cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx`,
	}, &types.GetStakesRequest{}
}

// GetCmdGetFutureRewardsEstimate returns a rewards estimate for a given set of stakes.
func GetCmdGetFutureRewardEstimate() (*dcli.QueryDescriptor, *types.GetFutureRewardEstimateRequest) {
	return &dcli.QueryDescriptor{
		Use:   "reward-estimate [owner] [stakeIDs] [endEpoch]",
		Short: "Get rewards estimate for set of stakes",
		Long:  `{{.Short}}{{.ExampleHeader}} reward-estimate cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx [1,2,3] 1681450672`,
		CustomFieldParsers: map[string]dcli.CustomFieldParserFn{
			"StakeIDs": dcli.ParseUintArray,
		},
	}, &types.GetFutureRewardEstimateRequest{}
}
