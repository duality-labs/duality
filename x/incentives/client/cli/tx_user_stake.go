package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/incentives/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateUserStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user-stake [index] [amount] [start-date] [end-date]",
		Short: "Create a new UserStake",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argAmount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}
			argStartDate, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argEndDate, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUserStake(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argAmount,
				argStartDate,
				argEndDate,
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

func CmdUpdateUserStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-user-stake [index] [amount] [start-date] [end-date]",
		Short: "Update a UserStake",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argAmount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}
			argStartDate, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argEndDate, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateUserStake(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argAmount,
				argStartDate,
				argEndDate,
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

func CmdDeleteUserStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-user-stake [index]",
		Short: "Delete a UserStake",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteUserStake(
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
