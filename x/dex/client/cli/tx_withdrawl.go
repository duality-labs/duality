package cli

import (
	"strconv"
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWithdrawl() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "withdrawl [receiver] [token-a] [token-b] [list of shares-to-remove] [list of tick-index] [list of fee indexes] ",
		Short: "Broadcast message withdrawl",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argTokenA := args[1]
			argTokenB := args[2]

			argSharesToRemove := strings.Split(args[3], ",")
			argTickIndexes := strings.Split(args[4], ",")
			argFeeIndexes := strings.Split(args[5], ",")

			var SharesToRemoveInt []sdk.Int
			var TicksIndexesInt []int64
			var FeeIndexesUint []uint64
			for _, s := range argSharesToRemove {
				sharesToRemoveInt, ok := sdk.NewIntFromString(s)

				if ok != true {
					return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Integer Overflow for shares-to-remove")
				}

				SharesToRemoveInt = append(SharesToRemoveInt, sharesToRemoveInt)
			}

			for _, s := range argTickIndexes {
				TickIndexInt, err := strconv.Atoi(s)

				if err != nil {
					return err
				}

				TicksIndexesInt = append(TicksIndexesInt, int64(TickIndexInt))

			}

			for _, s := range argFeeIndexes {
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
				SharesToRemoveInt,
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

	return cmd
}
