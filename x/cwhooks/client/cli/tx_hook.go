package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/duality-labs/duality/x/cwhooks/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateHook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-hook [contract-id] [args] [persistent] [trigger-key] [trigger-value]",
		Short: "Create a new hook",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContractID := args[0]
			argArgs := args[1]
			argPersistent, err := cast.ToBoolE(args[2])
			if err != nil {
				return err
			}
			argTriggerKey := args[3]
			argTriggerValue := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateHook(clientCtx.GetFromAddress().String(), argContractID, argArgs, argPersistent, argTriggerKey, argTriggerValue)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteHook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-hook [id]",
		Short: "Delete a hook by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteHook(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
