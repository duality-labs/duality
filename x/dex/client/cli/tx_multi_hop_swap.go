package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMultiHopSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multi-hop-swap [receiever] [hops] [amount-in] [exit-limit-price]",
		Short: "Broadcast message multiHopSwap",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiever := args[0]
			argHops := args[1]
			argAmountIn := args[2]
			argExitLimitPrice := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMultiHopSwap(
				clientCtx.GetFromAddress().String(),
				argReceiever,
				argHops,
				argAmountIn,
				argExitLimitPrice,
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
