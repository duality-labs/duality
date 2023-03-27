package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/spf13/cobra"
)

func CmdCancelLimitOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cancel-limit-order [tokenIn] [tokenOut] [tick-index] [tranche-key]",
		Short:   "Broadcast message CancelLimitOrder",
		Example: "cancel-limit-order alice tokenA tokenB [-10] 0 --from alice",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenIn := args[0]
			argTokenOut := args[1]

			if strings.HasPrefix(args[2], "[") && strings.HasSuffix(args[2], "]") {
				args[2] = strings.TrimPrefix(args[2], "[")
				args[2] = strings.TrimSuffix(args[2], "]")
			}
			argTickIndex := args[2]
			argTrancheKey := args[3]

			argTickIndexInt, err := strconv.ParseInt(argTickIndex, 10, 0)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelLimitOrder(
				clientCtx.GetFromAddress().String(),
				argTokenIn,
				argTokenOut,
				argTickIndexInt,
				argTrancheKey,
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
