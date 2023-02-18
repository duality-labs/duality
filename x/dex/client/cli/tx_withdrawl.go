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

func CmdWithdrawl() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "withdrawl [receiver] [token-a] [token-b] [list of shares-to-remove] [list of tick-index] [list of fee indexes] ",
		Short:   "Broadcast message withdrawl",
		Example: "withdrawl alice tokenA tokenB 100,50 [-10,5] 1,1 --from alice",
		Args:    cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argTokenA := args[1]
			argTokenB := args[2]
			argSharesToRemove := strings.Split(args[3], ",")

			if strings.HasPrefix(args[4], "[") && strings.HasSuffix(args[4], "]") {
				args[4] = strings.TrimPrefix(args[4], "[")
				args[4] = strings.TrimSuffix(args[4], "]")
			}
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
				TickIndexInt, err := strconv.ParseInt(s, 10, 0)
				if err != nil {
					return err
				}

				TicksIndexesInt = append(TicksIndexesInt, TickIndexInt)

			}

			for _, s := range argFeeIndexes {
				FeeIndexInt, err := strconv.ParseUint(s, 10, 0)
				if err != nil {
					return err
				}

				FeeIndexesUint = append(FeeIndexesUint, FeeIndexInt)
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
