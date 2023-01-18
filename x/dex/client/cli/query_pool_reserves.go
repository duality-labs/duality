package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/spf13/cobra"
)

func CmdListPoolReserves() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-pool-reserves [pair-id] [token-in]",
		Short: "Query AllPoolReserves",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPairId := args[0]
			reqTokenIn := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPoolReservesRequest{

				PairId:  reqPairId,
				TokenIn: reqTokenIn,
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			res, err := queryClient.PoolReservesAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPoolReserves() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-pool-reserves [pair-id] [tick-index] [token-in] [fee]",
		Short: "shows a PoolReserves",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPairId := args[0]
			argTickIndex := args[1]
			argTokenIn := args[2]
			argFee := args[3]

			argTrancheIndexInt, err := strconv.Atoi(argFee)

			if err != nil {
				return err
			}

			argTickIndexInt, err := strconv.Atoi(argTickIndex)

			if err != nil {
				return err
			}

			params := &types.QueryGetPoolReservesRequest{
				PairId:    argPairId,
				TickIndex: int64(argTickIndexInt),
				TokenIn:   argTokenIn,
				Fee:       uint64(argTrancheIndexInt),
			}

			res, err := queryClient.PoolReserves(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
