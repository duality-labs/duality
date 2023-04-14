package cli

import (
	"errors"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/duality-labs/duality/osmoutils/osmocli"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := osmocli.TxIndexCmd(types.ModuleName)

	osmocli.AddTxCmd(cmd, NewCreateGaugeCmd)
	osmocli.AddTxCmd(cmd, NewAddToGaugeCmd)
	osmocli.AddTxCmd(cmd, NewLockTokensCmd)
	osmocli.AddTxCmd(cmd, NewBeginUnlockingAllCmd)
	osmocli.AddTxCmd(cmd, NewBeginUnlockByIDCmd)

	return cmd
}

func CreateGaugeCmdBuilder(clientCtx client.Context, args []string, flags *pflag.FlagSet) (sdk.Msg, error) {
	// "create-gauge [pairID] [startTick] [endTick] [coins] [numEpochs] [pricingTick]"
	pairID, err := dextypes.StringToPairID(args[0])
	if err != nil {
		return &types.MsgCreateGauge{}, err
	}

	startTick, err := osmocli.ParseIntMaybeNegative(args[1], "startTick")
	if err != nil {
		return &types.MsgCreateGauge{}, err
	}

	endTick, err := osmocli.ParseIntMaybeNegative(args[2], "endTick")
	if err != nil {
		return &types.MsgCreateGauge{}, err
	}

	coins, err := sdk.ParseCoinsNormalized(args[3])
	if err != nil {
		return &types.MsgCreateGauge{}, err
	}

	var startTime time.Time
	timeStr, err := flags.GetString(FlagStartTime)
	if err != nil {
		return &types.MsgCreateGauge{}, err
	}
	if timeStr == "" { // empty start time
		startTime = time.Unix(0, 0)
	} else if timeUnix, err := strconv.ParseInt(timeStr, 10, 64); err == nil { // unix time
		startTime = time.Unix(timeUnix, 0)
	} else if timeRFC, err := time.Parse(time.RFC3339, timeStr); err == nil { // RFC time
		startTime = timeRFC
	} else { // invalid input
		return &types.MsgCreateGauge{}, errors.New("invalid start time format")
	}

	epochs, err := osmocli.ParseUint(args[4], "numEpochs")
	if err != nil {
		return &types.MsgCreateGauge{}, err
	}

	perpetual, err := flags.GetBool(FlagPerpetual)
	if err != nil {
		return &types.MsgCreateGauge{}, err
	}

	if perpetual {
		epochs = 1
	}

	pricingTick, err := osmocli.ParseIntMaybeNegative(args[5], "pricingTick")
	if err != nil {
		return &types.MsgCreateGauge{}, err
	}

	distributeTo := types.QueryCondition{
		PairID:    pairID,
		StartTick: startTick,
		EndTick:   endTick,
	}

	msg := types.NewMsgCreateGauge(
		epochs == 1,
		clientCtx.GetFromAddress(),
		distributeTo,
		coins,
		startTime,
		epochs,
		pricingTick,
	)

	return msg, nil
}

func NewCreateGaugeCmd() (*osmocli.TxCliDesc, *types.MsgCreateGauge) {
	return &osmocli.TxCliDesc{
		ParseAndBuildMsg: CreateGaugeCmdBuilder,
		Use:              "create-gauge [pairID] [startTick] [endTick] [coins] [numEpochs] [pricingTick]",
		Short:            "create a gauge to distribute rewards to users",
		Long: `{{.Short}}{{.ExampleHeader}}
TokenA<>TokenB [-10] 200 100TokenA,200TokenB 6 0 --start-time 2006-01-02T15:04:05Z07:00 --perpetual true`,
		Flags: osmocli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetCreateGauge()}},
	}, &types.MsgCreateGauge{}
}

func NewAddToGaugeCmd() (*osmocli.TxCliDesc, *types.MsgAddToGauge) {
	return &osmocli.TxCliDesc{
		Use:   "add-to-gauge [gauge_id] [coins]",
		Short: "add coins to gauge to distribute more rewards to users",
		Long:  `{{.Short}}{{.ExampleHeader}} add-to-gauge 1 TokenA,TokenB`,
	}, &types.MsgAddToGauge{}
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
