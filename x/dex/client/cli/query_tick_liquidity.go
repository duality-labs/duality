package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/spf13/cobra"
)

func CmdListTickLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tick-liquidity [pair-id] [token-in]",
		Short: "list all tickLiquidity",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			argPairId := args[0]
			argTokenIn := args[1]

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTickLiquidityRequest{
				PairId:     argPairId,
				TokenIn:    argTokenIn,
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
