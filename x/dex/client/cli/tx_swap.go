package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/spf13/cobra"
)

func CmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "swap [receiver] [amount-in] [tokenIn] [tokenOut]",
		Short:   "Broadcast message swap",
		Example: "swap alice 50 tokenA tokenB --from alice",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argAmountIn := args[1]

			amountInInt, ok := sdk.NewIntFromString(argAmountIn)
			if ok != true {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Integer overflow for amount-in")
			}

			argTokenIn := args[2]
			argTokenOut := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				argTokenIn,
				argTokenOut,
				amountInInt,
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
