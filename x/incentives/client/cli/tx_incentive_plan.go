package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/incentives/types"
	"github.com/spf13/cobra"
)

func CmdCreateIncentivePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-incentive-plan [index] [start-date] [end-date] [trading-pair] [total-amount] [start-tick] [end-tick]",
		Short: "Create a new IncentivePlan",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argStartDate, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			argEndDate, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			argTradingPair := args[3]
			argTotalAmount, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}
			argStartTick, err := strconv.ParseInt(args[5], 10, 64)
			if err != nil {
				return err
			}
			argEndTick, err := strconv.ParseInt(args[6], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateIncentivePlan(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argStartDate,
				argEndDate,
				argTradingPair,
				argTotalAmount,
				argStartTick,
				argEndTick,
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

func CmdUpdateIncentivePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-incentive-plan [index] [start-date] [end-date] [trading-pair] [total-amount] [start-tick] [end-tick]",
		Short: "Update a IncentivePlan",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argStartDate, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			argEndDate, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			argTradingPair := args[3]
			argTotalAmount, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}
			argStartTick, err := strconv.ParseInt(args[5], 10, 64)
			if err != nil {
				return err
			}
			argEndTick, err := strconv.ParseInt(args[6], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateIncentivePlan(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argStartDate,
				argEndDate,
				argTradingPair,
				argTotalAmount,
				argStartTick,
				argEndTick,
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

func CmdDeleteIncentivePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-incentive-plan [index]",
		Short: "Delete a IncentivePlan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteIncentivePlan(
				clientCtx.GetFromAddress().String(),
				indexIndex,
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
