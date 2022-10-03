package cli

import (
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPlaceLimitOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "place-limit-order [token-a] [token-b] [tick-index] [token-in] [amount-in]",
		Short: "Broadcast message PlaceLimitOrder",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenA := args[0]
			argTokenB := args[1]
			argTickIndex := args[2]
			argTokenIn := args[3]
			argAmountIn := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPlaceLimitOrder(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				argTickIndex,
				argTokenIn,
				argAmountIn,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
