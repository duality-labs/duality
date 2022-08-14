package cli

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListVirtualPriceTickList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-virtual-price-tick-list",
		Short: "list all virtualPriceTickList",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVirtualPriceTickListRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VirtualPriceTickListAll(context.Background(), params)
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

func CmdShowVirtualPriceTickList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-virtual-price-tick-list [v-price] [direction] [order-type]",
		Short: "shows a virtualPriceTickList",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argVPrice := args[0]
			argDirection := args[1]
			argOrderType := args[2]

			params := &types.QueryGetVirtualPriceTickListRequest{
				VPrice:    argVPrice,
				Direction: argDirection,
				OrderType: argOrderType,
			}

			res, err := queryClient.VirtualPriceTickList(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
