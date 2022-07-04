package cli

import (
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSingleDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "single-deposit [token-0] [token-1] [price] [fee] [amounts-0] [amounts-1] [receiver]",
		Short: "Broadcast message single_deposit",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToken0 := args[0]
			argToken1 := args[1]
			argPrice := args[2]
			argFee, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argAmounts0, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argAmounts1, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argReceiver := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSingleDeposit(
				clientCtx.GetFromAddress().String(),
				argToken0,
				argToken1,
				argPrice,
				argFee,
				argAmounts0,
				argAmounts1,
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
