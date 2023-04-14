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
	return cmd
}

// GetCmdGetModuleStatus returns status of incentive module.
func GetCmdGetModuleStatus() (*osmocli.QueryDescriptor, *types.GetModuleStatusRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "module-status",
		Short: "Query module status..",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} gauge-by-id 1
`,
	}, &types.GetModuleStatusRequest{}
}

// GetCmdGetGaugeByID returns a gauge by ID.
func GetCmdGetGaugeByID() (*osmocli.QueryDescriptor, *types.GetGaugeByIDRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "gauge-by-id [id]",
		Short: "Query gauge by id.",
		Long:  `{{.Short}}{{.ExampleHeader}}`,
	}, &types.GetGaugeByIDRequest{}
}

func parseFilterStatus(arg string, _ *pflag.FlagSet) (any, osmocli.FieldReadLocation, error) {
	filterStatusInt, ok := types.StatusFilter_value[arg]
	if !ok {
		return 0, osmocli.UsedArg, types.ErrInvalidGaugeStatus
	}
	filterStatus := types.StatusFilter(filterStatusInt)

	return filterStatus, osmocli.UsedArg, nil
}

// GetCmdGauges returns all gauges for a given filter.
func GetCmdGauges() (*osmocli.QueryDescriptor, *types.GetGaugesRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "list-gauges [StatusFilter] [denom]",
		Short: "Query all gauges",
		Long:  `{{.Short}}{{.ExampleHeader}}`,
		CustomFieldParsers: map[string]osmocli.CustomFieldParserFn{
			"StatusFilter": parseFilterStatus,
		},
	}, &types.GetGaugesRequest{}
}

// GetCmdGetLockByID returns a lock by ID.
func GetCmdGetLockByID() (*osmocli.QueryDescriptor, *types.GetLockByIDRequest) {
	return &osmocli.QueryDescriptor{
		Use:   "lock-by-id [LockId]",
		Short: "Query lock by id.",
		Long:  `{{.Short}}{{.ExampleHeader}}`,
	}, &types.GetLockByIDRequest{}
}
