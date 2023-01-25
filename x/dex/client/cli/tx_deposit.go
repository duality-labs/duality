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

func CmdDeposit() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "deposit [receiver] [token-a] [token-b] [list of amount-0] [list of amount-1] [list of tick-index] [list of fee] [deposit option parameters]",
		Short: "Broadcast message deposit",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argTokenA := args[1]
			argTokenB := args[2]
			argAmountsA := strings.Split(args[3], ",")
			argAmountsB := strings.Split(args[4], ",")
			argTicksIndexes := strings.Split(args[5], ",")
			argFeesIndexes := strings.Split(args[6], ",")
			argDepositOptions := strings.Split(args[7], ",")

			var AmountsA []sdk.Int
			var AmountsB []sdk.Int
			var TicksIndexesInt []int64
			var FeesIndexesUint []uint64
			var DepositOptions []*types.DepositOptions

			for _, s := range argAmountsA {
				amountA, ok := sdk.NewIntFromString(s)
				if ok != true {
					return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Integer overflow for amountsA")
				}

				AmountsA = append(AmountsA, amountA)
			}

			for _, s := range argAmountsB {
				amountB, ok := sdk.NewIntFromString(s)
				if ok != true {
					return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Integer overflow for amountsB")
				}

				AmountsB = append(AmountsB, amountB)
			}

			for _, s := range argTicksIndexes {
				str := strings.Trim(s, "\"")
				TickIndexInt, err := strconv.Atoi(str)
				if err != nil {
					return err
				}

				TicksIndexesInt = append(TicksIndexesInt, int64(TickIndexInt))

			}

			for _, s := range argFeesIndexes {
				FeeIndexInt, err := strconv.Atoi(s)
				if err != nil {
					return err
				}

				FeesIndexesUint = append(FeesIndexesUint, uint64(FeeIndexInt))
			}

			for _, s := range argDepositOptions {
				autoswap, err := strconv.ParseBool(s)
				if err != nil {
					return err
				}
				DepositOptions = append(DepositOptions, &types.DepositOptions{Autoswap: autoswap})
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argTokenA,
				argTokenB,
				AmountsA,
				AmountsB,
				TicksIndexesInt,
				FeesIndexesUint,
				DepositOptions,
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
