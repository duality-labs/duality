package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLimitOrderPoolFillMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-limit-order-pool-fill-map",
		Short: "list all LimitOrderPoolFillMap",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLimitOrderPoolFillMapRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LimitOrderPoolFillMapAll(context.Background(), params)
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

func CmdShowLimitOrderPoolFillMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-limit-order-pool-fill-map [count]",
		Short: "shows a LimitOrderPoolFillMap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			tmpargCount := args[0]

			argCount, err := strconv.Atoi(tmpargCount)

			if err != nil {
				return err
			}

			params := &types.QueryGetLimitOrderPoolFillMapRequest{
				Count: uint64(argCount),
			}

			res, err := queryClient.LimitOrderPoolFillMap(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}