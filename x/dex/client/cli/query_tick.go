package cli

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListTick() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tick",
		Short: "list all tick",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTickRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TickAll(context.Background(), params)
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

func CmdShowTick() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tick [token-0] [token-1] [price] [fee]",
		Short: "shows a tick",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argToken0 := args[0]
			argToken1 := args[1]
			argPrice := args[2]
			argFee, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			params := &types.QueryGetTickRequest{
				Token0: argToken0,
				Token1: argToken1,
				Price:  argPrice,
				Fee:    argFee,
			}

			res, err := queryClient.Tick(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
