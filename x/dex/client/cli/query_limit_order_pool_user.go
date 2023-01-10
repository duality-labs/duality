package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLimitOrderTrancheUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-limit-order-pool-user-share-map",
		Short: "list all LimitOrderTrancheUser",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLimitOrderTrancheUserRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LimitOrderTrancheUserAll(context.Background(), params)
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

func CmdShowLimitOrderTrancheUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-limit-order-pool-user-share-map [count] [address]",
		Short: "shows a LimitOrderTrancheUser",
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

			params := &types.QueryGetLimitOrderTrancheUserRequest{
				Count:   uint64(argCount),
				Address: argAddress,
			}

			res, err := queryClient.LimitOrderTrancheUser(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
