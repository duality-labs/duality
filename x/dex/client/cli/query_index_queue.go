package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListIndexQueue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-virtual-price-queue",
		Short: "list all IndexQueue",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllIndexQueueRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.IndexQueueAll(context.Background(), params)
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

func CmdShowIndexQueue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-virtual-price-queue  [token0] [token1] [index] ",
		Short: "shows a IndexQueue",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			token0 := args[0]
			token1 := args[1]
			i, err := strconv.ParseInt(args[0], 10, 32)

			if err != nil {
				return err
			}

			params := &types.QueryGetIndexQueueRequest{
				Token0: token0,
				Token1: token1,
				Index:  int32(i),
			}

			res, err := queryClient.IndexQueue(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
