package cli

import (
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCancelLimitOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-limit-order [receiver] [token-a] [token-b] [tick-index] [key-token] [key] [sharesOut]",
		Short: "Broadcast message CancelLimitOrder",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argTokenA := args[1]
			argTokenB := args[2]
			argTickIndex := args[3]

			argTickIndexInt, err := strconv.Atoi(argTickIndex)

			if err != nil {
				return err
			}

			argKeyToken := args[4]
			argKey := args[5]
			argSharesOut := args[6]

			argKeyInt, err := strconv.Atoi(argKey)

			if err != nil {
				return err
			}

			argSharesOutInt, ok := sdk.NewIntFromString(argSharesOut)
			if ok != true {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Integer overflow for sharesOut")
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
				argSharesOutInt,
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
