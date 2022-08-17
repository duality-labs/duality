package cli

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListVirtualPriceQueue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-virtual-price-queue",
		Short: "list all virtualPriceQueue",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVirtualPriceQueueRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VirtualPriceQueueAll(context.Background(), params)
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

func CmdShowVirtualPriceQueue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-virtual-price-queue [v-price] [direction] [order-type]",
		Short: "shows a virtualPriceQueue",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argVPrice := args[0]
			argDirection := args[1]
			argOrderType := args[2]

			params := &types.QueryGetVirtualPriceQueueRequest{
				VPrice:    argVPrice,
				Direction: argDirection,
				OrderType: argOrderType,
			}

			res, err := queryClient.VirtualPriceQueue(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
