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

func CmdWithdrawFilledLimitOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdrawl-withdrawn-limit-order [receiver] [token-a] [token-b] [tick-index] [key-token] [key]",
		Short: "Broadcast message WithdrawFilledLimitOrder",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argKeyToken := args[1]
			argTokenOut := args[2]
			argTickIndex := args[3]
			argTickIndexInt, err := strconv.Atoi(argTickIndex)

			if err != nil {
				return err
			}

			argKey := args[4]

			argKeyInt, err := strconv.Atoi(argKey)

			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawFilledLimitOrder(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argKeyToken,
				argTokenOut,
				int64(argTickIndexInt),
				uint64(argKeyInt),
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
