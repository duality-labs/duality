package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMultiHopSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multi-hop-swap [receiver] [hops] [amount-in] [exit-limit-price]",
		Short: "Broadcast message multiHopSwap",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiever := args[0]
			argHops := strings.Split(args[1], ",")
			argAmountIn := args[2]
			argExitLimitPrice := args[3]

			amountInInt, ok := sdk.NewIntFromString(argAmountIn)
			if !ok {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Invalid value for amount-in")
			}

			exitLimitPriceDec, err := sdk.NewDecFromStr(argExitLimitPrice)
			if err != nil {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Invalid value for exit-limit-price")
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMultiHopSwap(
				clientCtx.GetFromAddress().String(),
				argReceiever,
				argHops,
				amountInInt,
				exitLimitPriceDec,
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
