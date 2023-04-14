package cli

import (
	flag "github.com/spf13/pflag"
)

// Flags for incentives module tx commands.
const (
	FlagStartTime = "start-time"
	FlagEpochs    = "epochs"
	FlagPerpetual = "perpetual"
	FlagTimestamp = "timestamp"
	FlagOwner     = "owner"
	FlagStakeIds  = "stake-ids"
	FlagEndEpoch  = "end-epoch"
	FlagAmount    = "amount"
)

// FlagSetCreateGauge returns flags for creating gauges.
func FlagSetCreateGauge() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagStartTime, "", "Timestamp to begin distribution")
	fs.Uint64(FlagEpochs, 0, "Total epochs to distribute tokens")
	fs.Bool(FlagPerpetual, false, "Perpetual distribution")
	return fs
}

// FlagSetStakeTokens returns flags for StakeTokens msg builder.
func FlagSetSetupStake() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	return fs
}

func FlagSetUnSetupStake() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagAmount, "", "The amount to be unstaked. e.g. 1osmo")
	return fs
}
