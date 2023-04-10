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
	FlagLockIds   = "lock-ids"
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

// FlagSetLockTokens returns flags for LockTokens msg builder.
func FlagSetSetupLock() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	return fs
}

func FlagSetUnSetupLock() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagAmount, "", "The amount to be unlocked. e.g. 1osmo")
	return fs
}
