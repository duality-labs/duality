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
		Use:     "swap [receiver] [amount-in] [token-in] [token-out] ?[max-amount-aut]",
		Short:   "Broadcast message swap",
		Example: "swap alice 50 tokenA tokenB --from alice",
		Args:    cobra.RangeArgs(4, 5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argAmountIn := args[1]
			argTokenIn := args[2]
			argTokenOut := args[3]

			amountInInt, ok := sdk.NewIntFromString(argAmountIn)
			if !ok {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Integer overflow for amount-in")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			var maxAmountOutInt sdk.Int
			if len(args) == 5 {
				maxAmountOutInt, ok = sdk.NewIntFromString(args[4])
				if !ok {
					return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Integer overflow for max-amount-out")
				}
			}

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				argTokenIn,
				argTokenOut,
				amountInInt,
				&maxAmountOutInt,
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
