package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLimitOrderPoolUserSharesWithdrawn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-limit-order-pool-user-shares-withdrawn",
		Short: "list all LimitOrderPoolUserSharesWithdrawn",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLimitOrderPoolUserSharesWithdrawnRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LimitOrderPoolUserSharesWithdrawnAll(context.Background(), params)
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

func CmdShowLimitOrderPoolUserSharesWithdrawn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-limit-order-pool-user-shares-withdrawn [count] [address]",
		Short: "shows a LimitOrderPoolUserSharesWithdrawn",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			tmpargCount := args[0]

			argCount, err := strconv.Atoi(tmpargCount)

			if err != nil {
				return err
			}
			argAddress := args[1]

			params := &types.QueryGetLimitOrderPoolUserSharesWithdrawnRequest{
				Count:   uint64(argCount),
				Address: argAddress,
			}

			res, err := queryClient.LimitOrderPoolUserSharesWithdrawn(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
