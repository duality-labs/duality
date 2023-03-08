package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/duality-labs/duality/x/incentives/types"
	"github.com/spf13/cobra"
)

func CmdListIncentivePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-incentive-plan",
		Short: "list all IncentivePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllIncentivePlanRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.IncentivePlanAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowIncentivePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-incentive-plan [index]",
		Short: "shows a IncentivePlan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetIncentivePlanRequest{
				Index: argIndex,
			}

			res, err := queryClient.IncentivePlan(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
