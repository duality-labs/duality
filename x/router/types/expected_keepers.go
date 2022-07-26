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
	SortTokens(ctx sdk.Context, token0 string, token1 string) (string, string, error)
	Pop0to1(pools *([]*dextypes.Pool)) dextypes.Pool
	Update0to1(pools *([]*dextypes.Pool), pool *dextypes.Pool, reserve0, reserve1, fee, totalShares, price sdk.Dec)
	Pop1to0(pools *([]*dextypes.Pool)) dextypes.Pool
	Update1to0(pools *([]*dextypes.Pool), pool *dextypes.Pool, reserve0, reserve1, fee, totalShares, price sdk.Dec)
	Push0to1(pools *([]*dextypes.Pool), newPool *dextypes.Pool)
	Push1to0(pools *([]*dextypes.Pool), newPool *dextypes.Pool)
	GetPool(pools *([]*dextypes.Pool), Fee, Price sdk.Dec) (dextypes.Pool, bool)

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
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}
