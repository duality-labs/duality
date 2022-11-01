package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLimitOrderPoolUserShareMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-limit-order-pool-user-share-map",
		Short: "list all LimitOrderPoolUserShareMap",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLimitOrderPoolUserShareMapRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LimitOrderPoolUserShareMapAll(context.Background(), params)
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

func CmdShowLimitOrderPoolUserShareMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-limit-order-pool-user-share-map [count] [address]",
		Short: "shows a LimitOrderPoolUserShareMap",
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

			params := &types.QueryGetLimitOrderPoolUserShareMapRequest{
				Count:   uint64(argCount),
				Address: argAddress,
			}

			res, err := queryClient.LimitOrderPoolUserShareMap(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
