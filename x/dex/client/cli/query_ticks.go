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
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argToken0 := args[0]
			argToken1 := args[1]
			argPrice := args[2]
			argFee := args[3]
			argOrderType := args[4]

			params := &types.QueryGetTicksRequest{
				Token0:    argToken0,
				Token1:    argToken1,
				Price:     argPrice,
				Fee:       argFee,
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
