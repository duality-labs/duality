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
	osmocli.AddTxCmd(cmd, NewStakeTokensCmd)
	osmocli.AddTxCmd(cmd, NewUnstakeByIDCmd)

	return cmd
}

// TODO
// NewCreateGaugeCmd broadcasts a CreateGauge message.
func NewCreateGaugeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-gauge [stakeup_denom] [reward] [flags]",
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

func NewStakeTokensCmd() (*osmocli.TxCliDesc, *types.MsgStake) {
	return &osmocli.TxCliDesc{
		Use:   "stake-tokens [tokens]",
		Short: "stake tokens into stakeup pool from user account",
	}, &types.MsgStake{}
}

// NewUnstakeByIDCmd unstakes individual period stake by ID.
func NewUnstakeByIDCmd() (*osmocli.TxCliDesc, *types.MsgUnstake) {
	return &osmocli.TxCliDesc{
		Use:   "begin-unstake-by-id [id]",
		Short: "begin unstake individual period stake by ID",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
		Flags: osmocli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnSetupStake()}},
	}, &types.MsgUnstake{}
}
