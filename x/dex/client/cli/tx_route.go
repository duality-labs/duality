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

func CmdRoute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "route [receiver] [token-in] [token-out] [amount-in] [min-out]",
		Short: "Broadcast message route",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argTokenIn := args[1]
			argTokenOut := args[2]
			argAmountIn := args[3]
			argAmountInDec, err := sdk.NewDecFromStr(argAmountIn)

			if err != nil {
				return err
			}
			argMinOut := args[4]
			argMinOutDec, err := sdk.NewDecFromStr(argMinOut)

			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRoute(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argTokenIn,
				argTokenOut,
				argAmountInDec,
				argMinOutDec,
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
