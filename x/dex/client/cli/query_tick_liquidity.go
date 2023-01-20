package cli

import (
	"context"
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListTickLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tick-liquidity",
		Short: "list all tickLiquidity",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTickLiquidityRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TickLiquidityAll(context.Background(), params)
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

func CmdShowTickLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tick-liquidity [pair-id] [token-in] [tick-index] [liquidity-type] [liquidity-index]",
		Short: "shows a tickLiquidity",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPairId := args[0]
			argTokenIn := args[1]

			tickIndexString := strings.Trim(args[2], "\"")
			argTickIndex, err := cast.ToInt64E(tickIndexString)
			if err != nil {
				return err
			}
			argLiquidityType := args[3]
			argLiquidityIndex, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			params := &types.QueryGetTickLiquidityRequest{
				PairId:         argPairId,
				TokenIn:        argTokenIn,
				TickIndex:      argTickIndex,
				LiquidityType:  argLiquidityType,
				LiquidityIndex: argLiquidityIndex,
			}

			res, err := queryClient.TickLiquidity(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
