package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"runtime"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// TEMP: Temporary warning until CLI commands have been fully tested
	PrintCLIWarning()

	cmd.AddCommand(CmdDeposit())
	cmd.AddCommand(CmdWithdrawl())
	cmd.AddCommand(CmdSwap())
	cmd.AddCommand(CmdPlaceLimitOrder())
	cmd.AddCommand(CmdWithdrawFilledLimitOrder())
	cmd.AddCommand(CmdCancelLimitOrder())
	// this line is used by starport scaffolding # 1

	return cmd
}

func PrintCLIWarning(){
	colorChar := "\033[31m"
	colorEnd := "\033[0m"
	if runtime.GOOS == "windows" {
		colorChar, colorEnd = "", ""
	}
	fmt.Printf(colorChar + "\n!! WARNING CLI COMMANDS ARE STILL IN EARLY BETA !! \nNOT ALL FUNCTIONALITY HAS BEEN TESTED\n" + colorEnd)
}
