package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListTick() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tick-map",
		Short: "list all Tick",
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
		Use:   "show-tick-map [tick-index] [pairId]",
		Short: "shows a Tick",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argTickIndex := args[0]
			argPairId := args[1]

			tmpTickIndex, _ := strconv.Atoi(argTickIndex)
			params := &types.QueryGetTickRequest{
				TickIndex: int64(tmpTickIndex),
				PairId:    argPairId,
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
