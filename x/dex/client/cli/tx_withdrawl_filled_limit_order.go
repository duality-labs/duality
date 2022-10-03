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

func CmdWithdrawlFilledLimitOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdrawl-filled-limit-order [token-a] [token-b] [tick-index] [key-token] [key]",
		Short: "Broadcast message WithdrawlFilledLimitOrder",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenA := args[0]
			argTokenB := args[1]
			argTickIndex := args[2]
			argKeyToken := args[3]
			argKey := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawlFilledLimitOrder(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				argTickIndex,
				argKeyToken,
				argKey,
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
