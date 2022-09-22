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

func CmdWithdrawal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Withdrawal [token-a] [token-b] [shares-to-remove] [price-index] [fee] [receiver]",
		Short: "Broadcast message Withdrawal",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenA := args[0]
			argTokenB := args[1]
			argSharesToRemove := args[2]
			argPriceIndex := args[3]
			argFee := args[4]
			argReceiver := args[5]

			tmpArgPriceIndex, err := strconv.Atoi(argPriceIndex)

			if err != nil {
				return err
			}

			tmpArgFeeIndex, err := strconv.Atoi(argFee)

			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawal(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				argSharesToRemove,
				int64(tmpArgPriceIndex),
				uint64(tmpArgFeeIndex),
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
