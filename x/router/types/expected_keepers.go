package types

import (
	dextypes "github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type DexKeeper interface {
	GetAllTicks(ctx sdk.Context) (list []dextypes.Ticks)
	GetTicks(ctx sdk.Context, token0 string, token1 string) (val dextypes.Ticks, found bool)
	SetTicks(ctx sdk.Context, ticks dextypes.Ticks)
	SortTokens(ctx sdk.Context, tokens0 []string, tokens1 []string) ([]string, []string, error)

	// Methods imported from dex should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
