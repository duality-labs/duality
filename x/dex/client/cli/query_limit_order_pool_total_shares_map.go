package cli

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLimitOrderPoolTotalSharesMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-limit-order-pool-total-shares-map",
		Short: "list all LimitOrderPoolTotalSharesMap",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLimitOrderPoolTotalSharesMapRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LimitOrderPoolTotalSharesMapAll(context.Background(), params)
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

func CmdShowLimitOrderPoolTotalSharesMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-limit-order-pool-total-shares-map [count]",
		Short: "shows a LimitOrderPoolTotalSharesMap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCount := args[0]

			params := &types.QueryGetLimitOrderPoolTotalSharesMapRequest{
				Count: argCount,
			}

			res, err := queryClient.LimitOrderPoolTotalSharesMap(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
