package cli

import (
	"context"

	"github.com/duality-labs/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListFilledLimitOrderTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-filled-limit-order-tranche",
		Short: "list all FilledLimitOrderTranche",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFilledLimitOrderTrancheRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.FilledLimitOrderTrancheAll(context.Background(), params)
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

func CmdShowFilledLimitOrderTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-filled-limit-order-tranche [pair-id] [token-in] [tick-index] [tranche-index]",
		Short: "shows a FilledLimitOrderTranche",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPairId := args[0]
			argTokenIn := args[1]
			argTickIndex, err := cast.ToInt64E(args[2])
			if err != nil {
				return err
			}
			argTrancheIndex, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			params := &types.QueryGetFilledLimitOrderTrancheRequest{
				PairId:       argPairId,
				TokenIn:      argTokenIn,
				TickIndex:    argTickIndex,
				TrancheIndex: argTrancheIndex,
			}

			res, err := queryClient.FilledLimitOrderTranche(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
