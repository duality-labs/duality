package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/duality-labs/duality/osmoutils/osmocli"
	"github.com/duality-labs/duality/x/lockup/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := osmocli.TxIndexCmd(types.ModuleName)
	osmocli.AddTxCmd(cmd, NewLockTokensCmd)
	osmocli.AddTxCmd(cmd, NewBeginUnlockingAllCmd)
	osmocli.AddTxCmd(cmd, NewBeginUnlockByIDCmd)

	return cmd
}

func NewLockTokensCmd() (*osmocli.TxCliDesc, *types.MsgLockTokens) {
	return &osmocli.TxCliDesc{
		Use:   "lock-tokens [tokens]",
		Short: "lock tokens into lockup pool from user account",
		CustomFlagOverrides: map[string]string{
			"duration": FlagDuration,
		},
		Flags: osmocli.FlagDesc{RequiredFlags: []*pflag.FlagSet{FlagSetLockTokens()}},
	}, &types.MsgLockTokens{}
}

// TODO: We should change the Use string to be unlock-all
func NewBeginUnlockingAllCmd() (*osmocli.TxCliDesc, *types.MsgBeginUnlockingAll) {
	return &osmocli.TxCliDesc{
		Use:   "begin-unlock-tokens",
		Short: "begin unlock not unlocking tokens from lockup pool for sender",
	}, &types.MsgBeginUnlockingAll{}
}

// NewBeginUnlockByIDCmd unlocks individual period lock by ID.
func NewBeginUnlockByIDCmd() (*osmocli.TxCliDesc, *types.MsgBeginUnlocking) {
	return &osmocli.TxCliDesc{
		Use:   "begin-unlock-by-id [id]",
		Short: "begin unlock individual period lock by ID",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
		Flags: osmocli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnlockTokens()}},
	}, &types.MsgBeginUnlocking{}
}
