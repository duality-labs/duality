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
	osmocli.AddQueryCmd(cmd, qcGetter, GetCmdGetStakeByID)
	osmocli.AddQueryCmd(cmd, qcGetter, GetCmdStakes)
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
		Long:  `{{.Short}}{{.ExampleHeader}} list-gauges UPCOMING DualityPoolShares-stake-token-t0-f1`,
		CustomFieldParsers: map[string]osmocli.CustomFieldParserFn{
			"Status": parseGaugeStatus,
		},
	}, &types.GetGaugesRequest{}
}

// GetCmdGetStakeByID returns a lock by ID.
func GetCmdGetStakeByID() (*osmocli.QueryDescriptor, *types.GetStakeByIDRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "stake-by-id [stakeID]",
		Short: "Query stake by id.",
		Long:  `{{.Short}}{{.ExampleHeader}} Stake-by-id 1`,
	}, &types.GetStakeByIDRequest{}
}

// GetCmdStakes returns all gauges for a given status and owner.
func GetCmdStakes() (*osmocli.QueryDescriptor, *types.GetStakesRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "list-stakes [owner]",
		Short: "Query stakes",
		Long:  `{{.Short}}{{.ExampleHeader}} list-stakes cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx`,
	}, &types.GetStakesRequest{}
}

// GetCmdGetFutureRewardsEstimate returns a rewards estimate for a given set of stakes.
func GetCmdGetFutureRewardEstimate() (*osmocli.QueryDescriptor, *types.GetFutureRewardEstimateRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "reward-estimate [owner] [stakeIDs] [endEpoch]",
		Short: "Get rewards estimate for set of stakes",
		Long:  `{{.Short}}{{.ExampleHeader}} reward-estimate cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx [1,2,3] 1681450672`,
		CustomFieldParsers: map[string]osmocli.CustomFieldParserFn{
			"StakeIDs": osmocli.ParseUintArray,
		},
	}, &types.GetFutureRewardEstimateRequest{}
}
