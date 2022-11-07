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
	cmd.AddCommand(CmdListTickObject())
	cmd.AddCommand(CmdShowTickObject())
	cmd.AddCommand(CmdListPairObject())
	cmd.AddCommand(CmdShowPairObject())
	cmd.AddCommand(CmdListTokens())
	cmd.AddCommand(CmdShowTokens())
	cmd.AddCommand(CmdListTokenObject())
	cmd.AddCommand(CmdShowTokenObject())
	cmd.AddCommand(CmdListShares())
	cmd.AddCommand(CmdShowShares())
	cmd.AddCommand(CmdListFeeList())
	cmd.AddCommand(CmdShowFeeList())
	cmd.AddCommand(CmdListLimitOrderPoolUserShareObject())
	cmd.AddCommand(CmdShowLimitOrderPoolUserShareObject())
	cmd.AddCommand(CmdListLimitOrderPoolUserSharesWithdrawnObject())
	cmd.AddCommand(CmdShowLimitOrderPoolUserSharesWithdrawnObject())
	cmd.AddCommand(CmdListLimitOrderPoolTotalSharesObject())
	cmd.AddCommand(CmdShowLimitOrderPoolTotalSharesObject())
	cmd.AddCommand(CmdListLimitOrderPoolReserveObject())
	cmd.AddCommand(CmdShowLimitOrderPoolReserveObject())
	cmd.AddCommand(CmdListLimitOrderPoolFillObject())
	cmd.AddCommand(CmdShowLimitOrderPoolFillObject())
	// this line is used by starport scaffolding # 1

	return cmd
}
