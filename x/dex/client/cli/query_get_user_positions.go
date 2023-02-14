package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdShowUserPositions() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "show-user-positions [address]",
		Short:   "shows a users current positions",
		Example: "show-user-positions alice",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetUserPositionsRequest{

				Address: reqAddress,
			}

			res, err := queryClient.GetUserPositions(cmd.Context(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
