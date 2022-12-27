package cli

import (
	"strconv"

	"github.com/NicholasDotSol/duality/x/mev/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWithdrawFunds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-funds [amount-out] [token-out]",
		Short: "Broadcast message withdrawFunds",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmountOut := args[0]
			argTokenOut := args[1]

			argAmountOutInt, ok := sdk.NewIntFromString(argAmountOut)

			if !ok {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "AmountIn Overflower error")
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawFunds(
				clientCtx.GetFromAddress().String(),
				argAmountOutInt,
				argTokenOut,
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
