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

var _ = strconv.Itoa(0)

func CmdCancelLimitOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cancel-limit-order [receiver] [token-a] [token-b] [tick-index] [key-token] [key]",
		Short:   "Broadcast message CancelLimitOrder",
		Example: "cancel-limit-order alice tokenA tokenB [-10] tokenA 0 --from alice",
		Args:    cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argTokenA := args[1]
			argTokenB := args[2]

			if strings.HasPrefix(args[3], "[") && strings.HasSuffix(args[3], "]") {
				args[3] = strings.TrimPrefix(args[3], "[")
				args[3] = strings.TrimSuffix(args[3], "]")
			}
			argTickIndex := args[3]

			argTickIndexInt, err := strconv.Atoi(argTickIndex)
			if err != nil {
				return err
			}

			argKeyToken := args[4]
			argKey := args[5]

			argKeyInt, err := strconv.Atoi(argKey)

			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelLimitOrder(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argTokenA,
				argTokenB,
				int64(argTickIndexInt),
				argKeyToken,
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
