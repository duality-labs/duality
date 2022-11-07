package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListTickObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tick-map",
		Short: "list all TickObject",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTickObjectRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TickObjectAll(context.Background(), params)
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

func CmdShowTickObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tick-map [tick-index]",
		Short: "shows a TickObject",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argTickIndex := args[0]

			tmpTickIndex, _ := strconv.Atoi(argTickIndex)
			params := &types.QueryGetTickObjectRequest{
				TickIndex: int64(tmpTickIndex),
			}

			res, err := queryClient.TickObject(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
