package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NicholasDotSol/duality/x/dex/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group dex queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListTick())
	cmd.AddCommand(CmdShowTick())
	cmd.AddCommand(CmdListTradingPair())
	cmd.AddCommand(CmdShowTradingPair())
	cmd.AddCommand(CmdListTokens())
	cmd.AddCommand(CmdShowTokens())
	cmd.AddCommand(CmdListTokenMap())
	cmd.AddCommand(CmdShowTokenMap())
	cmd.AddCommand(CmdListFeeTier())
	cmd.AddCommand(CmdShowFeeTier())
	cmd.AddCommand(CmdListLimitOrderTrancheUser())
	cmd.AddCommand(CmdShowLimitOrderTrancheUser())
	cmd.AddCommand(CmdListLimitOrderTranche())
	cmd.AddCommand(CmdShowLimitOrderTranche())
	// this line is used by starport scaffolding # 1

	return cmd
}
