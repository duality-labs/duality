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

func CmdDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [token-a] [token-b] [amount-0] [amount-1] [tick-index] [fee] [receiver]",
		Short: "Broadcast message deposit",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenA := args[0]
			argTokenB := args[1]
			argAmount0 := args[2]
			argAmount1 := args[3]
			argTickIndex := args[4]
			argFee := args[5]

			tmpArgTickIndex, err := strconv.Atoi(argTickIndex)

			if err != nil {
				return err
			}

			tmpArgFeeIndex, err := strconv.Atoi(argFee)

			if err != nil {
				return err
			}

			argReceiver := args[7]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				argAmount0,
				argAmount1,
				int64(tmpArgTickIndex),
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
