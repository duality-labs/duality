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
		Use:   "deposit [token-a] [token-b] [amount-0] [amount-1] [price-index] [fee]",
		Short: "Broadcast message deposit",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenA := args[0]
			argTokenB := args[1]
			argAmount0 := args[2]
			argAmount1 := args[3]
			argPriceIndex := args[4]
			argFee := args[5]

			tmpArgPriceIndex, err := strconv.Atoi(argPriceIndex)

			if err != nil {
				return err
			}

			tmpArgFee, err := strconv.Atoi(argFee)

			if err != nil {
				return err
			}

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
				int64(tmpArgPriceIndex),
				int64(tmpArgFee),
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
