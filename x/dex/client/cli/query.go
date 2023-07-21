package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/duality-labs/duality/x/dex/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group dex queries under a subcommand
	cmd := &cobra.Command{
		Use: types.ModuleName,
		Short: fmt.Sprintf(
			"Querying commands for the %s module",
			types.ModuleName,
		),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListLimitOrderTrancheUser())
	cmd.AddCommand(CmdShowLimitOrderTrancheUser())
	cmd.AddCommand(CmdListLimitOrderTranche())
	cmd.AddCommand(CmdShowLimitOrderTranche())
	cmd.AddCommand(CmdShowUserPositions())
	cmd.AddCommand(CmdListUserDeposits())
	cmd.AddCommand(CmdListUserLimitOrders())
	cmd.AddCommand(CmdListTickLiquidity())
	cmd.AddCommand(CmdListInactiveLimitOrderTranche())
	cmd.AddCommand(CmdShowInactiveLimitOrderTranche())
	cmd.AddCommand(CmdListPoolReserves())
	cmd.AddCommand(CmdShowPoolReserves())
	cmd.AddCommand(CmdShowPool())

	// this line is used by starport scaffolding # 1

	return cmd
}
