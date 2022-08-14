package cli

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListTicks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-ticks",
		Short: "list all ticks",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTicksRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TicksAll(context.Background(), params)
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

func CmdShowTicks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-ticks [price] [fee] [direction] [order-type]",
		Short: "shows a ticks",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPrice := args[0]
			argFee := args[1]
			argDirection := args[2]
			argOrderType := args[3]

			params := &types.QueryGetTicksRequest{
				Price:     argPrice,
				Fee:       argFee,
				Direction: argDirection,
				OrderType: argOrderType,
			}

			res, err := queryClient.Ticks(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
