package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLimitOrderTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-limit-order-pool-total-shares-map",
		Short: "list all LimitOrderTranche",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLimitOrderTrancheRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LimitOrderTrancheAll(context.Background(), params)
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

func CmdShowLimitOrderTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-limit-order-pool-total-shares-map [pairId] [tickIndex] [tokenIn] [TrancheIndex]",
		Short: "shows a LimitOrderTranche",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPairId := args[0]
			argTickIndex := args[1]
			argTokenIn := args[2]
			argTrancheIndex := args[3]

			argTrancheIndexInt, err := strconv.Atoi(argTrancheIndex)

			if err != nil {
				return err
			}

			argTickIndexInt, err := strconv.Atoi(argTickIndex)

			if err != nil {
				return err
			}

			params := &types.QueryGetLimitOrderTrancheRequest{
				PairId:       argPairId,
				TickIndex:    int64(argTickIndexInt),
				Token:        argTokenIn,
				TrancheIndex: uint64(argTrancheIndexInt),
			}

			res, err := queryClient.LimitOrderTranche(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
