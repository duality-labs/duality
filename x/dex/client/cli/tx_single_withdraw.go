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

func CmdSingleWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "single-withdraw [token-0] [token-1] [price] [fee] [shares-removing] [receiver]",
		Short: "Broadcast message single_withdraw",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToken0 := args[0]
			argToken1 := args[1]
			argPrice := args[2]
			argFee, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argSharesRemoving := args[4]
			argReceiver := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSingleWithdraw(
				clientCtx.GetFromAddress().String(),
				argToken0,
				argToken1,
				argPrice,
				argFee,
				argSharesRemoving,
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
