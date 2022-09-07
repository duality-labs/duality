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

func CmdWithdrawl() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdrawl [token-a] [token-b] [shares-to-remove] [price-index] [fee] [receiver]",
		Short: "Broadcast message withdrawl",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenA := args[0]
			argTokenB := args[1]
			argSharesToRemove := args[2]
			argPriceIndex := args[3]
			argFee := args[4]
			argReceiver := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawl(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				argSharesToRemove,
				argPriceIndex,
				argFee,
				argReceiver,
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
