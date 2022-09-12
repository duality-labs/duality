package cli

import (
	"context"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListShares() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-shares",
		Short: "list all Shares",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSharesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SharesAll(context.Background(), params)
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

func CmdShowShares() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-shares [address] [pair-id] [price-index] [fee]",
		Short: "shows a Shares",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argPairId := args[1]
			argPriceIndex := args[2]
			argFeeIndex := args[3]

			tmpArgPriceIndex, err := strconv.Atoi(argPriceIndex)

			if err != nil {
				return err
			}

			tmpArgFeeIndex, err := strconv.Atoi(argFeeIndex)

			if err != nil {
				return err
			}

			params := &types.QueryGetSharesRequest{
				Address:    argAddress,
				PairId:     argPairId,
				PriceIndex: int64(tmpArgPriceIndex),
				Fee:        uint64(tmpArgFeeIndex),
			}

			res, err := queryClient.Shares(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
