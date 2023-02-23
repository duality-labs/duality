package cli

import (
	"context"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/duality-labs/duality/x/dex/types"
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
		Use:     "show-limit-order-pool-user-share-map [pairId] [tickIndex] [tokenIn] [tranche-key] [address]",
		Short:   "shows a LimitOrderTrancheUser",
		Example: "show-limit-order-pool-user-share-map tokenA<>tokenB [-5] tokenA 0 alice",
		Args:    cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPairId := args[0]
			argTickIndex := args[1]
			argTokenIn := args[2]
			if strings.HasPrefix(args[1], "[") && strings.HasSuffix(args[1], "]") {
				args[1] = strings.TrimPrefix(args[1], "[")
				args[1] = strings.TrimSuffix(args[1], "]")
			}
			argTrancheKey := args[3]

			argAddress := args[4]

			argTickIndexInt, err := strconv.ParseInt(argTickIndex, 10, 0)

			if err != nil {
				return err
			}

			params := &types.QueryGetLimitOrderTrancheUserRequest{
				PairId:     argPairId,
				TickIndex:  argTickIndexInt,
				Token:      argTokenIn,
				TrancheKey: argTrancheKey,
				Address:    argAddress,
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
