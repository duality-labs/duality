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

func CmdWithdrawl() *cobra.Command {

	var sharesToRemove []string
	var TicksIndexes []string
	var FeesIndexes []string
	cmd := &cobra.Command{
		Use:   "withdrawl [receiver] [token-a] [token-b] [list of shares-to-remove] [list of tick-index] [list of fee indexes] ",
		Short: "Broadcast message withdrawl",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argTokenA := args[1]
			argTokenB := args[2]

			var sharesToRemoveDec []sdk.Dec
			var TicksIndexesInt []int64
			var FeeIndexesUint []uint64
			for _, s := range sharesToRemove {
				sharesDec, err := sdk.NewDecFromStr(s)

				if err != nil {
					return err
				}

				sharesToRemoveDec = append(sharesToRemoveDec, sharesDec)
			}

			for _, s := range TicksIndexes {
				TickIndexInt, err := strconv.Atoi(s)

				if err != nil {
					return err
				}

				TicksIndexesInt = append(TicksIndexesInt, int64(TickIndexInt))

			}

			for _, s := range FeesIndexes {
				FeeIndexInt, err := strconv.Atoi(s)

				if err != nil {
					return err
				}

				FeeIndexesUint = append(FeeIndexesUint, uint64(FeeIndexInt))
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawl(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argTokenA,
				argTokenB,
				sharesToRemoveDec,
				TicksIndexesInt,
				FeeIndexesUint,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().StringArrayVarP(&sharesToRemove, "sharesToRemove", "r", []string{}, "")
	cmd.Flags().StringArrayVarP(&TicksIndexes, "ticksIndexes", "t", []string{}, "")
	cmd.Flags().StringArrayVarP(&FeesIndexes, "feeIndexes", "f", []string{}, "")

	return cmd
}
