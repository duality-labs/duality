package cli

import (
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [receiver] [amount-in] [tokenA] [tokenB] [token-in] [minOut] ",
		Short: "Broadcast message swap",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argAmountIn := args[1]

			amountInInt, err := strconv.Atoi(argAmountIn)
			if err != nil {
				return err
			}

			amountIn := sdk.NewInt(int64(amountInInt))

			argTokenA := args[2]
			argTokenB := args[3]
			argTokenIn := args[4]
			argMinOut := args[5]

			minOutInt, err := strconv.Atoi(argMinOut)
			if err != nil {
				return err
			}
			minOut := sdk.NewInt(int64(minOutInt))

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				amountIn,
				argTokenIn,
				minOut,
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
