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

func CmdRemoveLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-liquidity [token-a] [token-b] [shares] [price] [fee] [order-type]",
		Short: "Broadcast message RemoveLiquidity",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenA := args[0]
			argTokenB := args[1]
			argShares := args[2]
			argPrice := args[3]
			argFee := args[4]
			argOrderType := args[5]
			argReceiver := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveLiquidity(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				argShares,
				argPrice,
				argFee,
				argOrderType,
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
