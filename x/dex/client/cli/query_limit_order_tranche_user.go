package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLimitOrderTrancheUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-limit-order-pool-user-share-map",
		Short: "list all LimitOrderTrancheUser",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLimitOrderTrancheUserRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LimitOrderTrancheUserAll(context.Background(), params)
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

func CmdShowLimitOrderTrancheUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-limit-order-pool-user-share-map [pairId] [tickIndex] [tokenIn] [trancheIndex] [address]",
		Short: "shows a LimitOrderTrancheUser",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPairId := args[0]
			argTickIndex := args[1]
			argTokenIn := args[2]
			argTrancheIndex := args[3]
			argAddress := args[4]

			argTickIndexInt, err := strconv.Atoi(argTickIndex)

			if err != nil {
				return err
			}

			argTrancheIndexInt, err := strconv.Atoi(argTrancheIndex)

			if err != nil {
				return err
			}

			params := &types.QueryGetLimitOrderTrancheUserRequest{
				PairId:    argPairId,
				TickIndex: int64(argTickIndexInt),
				Token:     argTokenIn,
				Count:     uint64(argTrancheIndexInt),
				Address:   argAddress,
			}

			res, err := queryClient.LimitOrderTrancheUser(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
