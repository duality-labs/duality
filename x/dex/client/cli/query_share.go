package cli

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-share",
		Short: "list all share",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllShareRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ShareAll(context.Background(), params)
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

func CmdShowShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-share [owner] [token-0] [token-1] [price] [fee]",
		Short: "shows a share",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argOwner := args[0]
			argToken0 := args[1]
			argToken1 := args[2]
			argPrice := args[3]
			argFee := args[4]

			params := &types.QueryGetShareRequest{
				Owner:  argOwner,
				Token0: argToken0,
				Token1: argToken1,
				Price:  argPrice,
				Fee:    argFee,
			}

			res, err := queryClient.Share(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
