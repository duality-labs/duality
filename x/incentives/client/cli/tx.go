package cli

import (
	"errors"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/duality-labs/duality/osmoutils/osmocli"
	"github.com/duality-labs/duality/x/incentives/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := osmocli.TxIndexCmd(types.ModuleName)
	cmd.AddCommand(
		NewCreateGaugeCmd(),
		NewAddToGaugeCmd(),
	)
	osmocli.AddTxCmd(cmd, NewLockTokensCmd)
	osmocli.AddTxCmd(cmd, NewBeginUnlockingAllCmd)
	osmocli.AddTxCmd(cmd, NewBeginUnlockByIDCmd)

	return cmd
}

// TODO
// NewCreateGaugeCmd broadcasts a CreateGauge message.
func NewCreateGaugeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-gauge [lockup_denom] [reward] [flags]",
		Short: "create a gauge to distribute rewards to users",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// denom := args[0]

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)
			coins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			var startTime time.Time
			timeStr, err := cmd.Flags().GetString(FlagStartTime)
			if err != nil {
				return err
			}
			if timeStr == "" { // empty start time
				startTime = time.Unix(0, 0)
			} else if timeUnix, err := strconv.ParseInt(timeStr, 10, 64); err == nil { // unix time
				startTime = time.Unix(timeUnix, 0)
			} else if timeRFC, err := time.Parse(time.RFC3339, timeStr); err == nil { // RFC time
				startTime = timeRFC
			} else { // invalid input
				return errors.New("invalid start time format")
			}

			epochs, err := cmd.Flags().GetUint64(FlagEpochs)
			if err != nil {
				return err
			}

			perpetual, err := cmd.Flags().GetBool(FlagPerpetual)
			if err != nil {
				return err
			}

			if perpetual {
				epochs = 1
			}

			distributeTo := types.QueryCondition{
				// Denom:     denom,
			}

			msg := types.NewMsgCreateGauge(
				epochs == 1,
				clientCtx.GetFromAddress(),
				distributeTo,
				coins,
				startTime,
				epochs,
			)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetCreateGauge())
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddToGaugeCmd() *cobra.Command {
	return osmocli.BuildTxCli[*types.MsgAddToGauge](&osmocli.TxCliDesc{
		Use:   "add-to-gauge [gauge_id] [rewards] [flags]",
		Short: "add coins to gauge to distribute more rewards to users",
	})
}

func NewLockTokensCmd() (*osmocli.TxCliDesc, *types.MsgLockTokens) {
	return &osmocli.TxCliDesc{
		Use:   "lock-tokens [tokens]",
		Short: "lock tokens into lockup pool from user account",
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
		Flags: osmocli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnSetupLock()}},
	}, &types.MsgBeginUnlocking{}
}
