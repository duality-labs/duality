package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/duality-labs/duality/x/dex/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	epochstypes "github.com/duality-labs/duality/x/epochs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin

	HasSupply(ctx sdk.Context, denom string) bool
	GetSupply(ctx sdk.Context, denom string) sdk.Coin

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// EpochKeeper defines the expected interface needed to retrieve epoch info.
type EpochKeeper interface {
	GetEpochInfo(ctx sdk.Context, identifier string) epochstypes.EpochInfo
}

type AccountKeeper interface {
	GetAllAccounts(ctx sdk.Context) []authtypes.AccountI
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
}

type DexKeeper interface {
	GetOrInitPool(ctx sdk.Context, pairID *types.PairID, centerTickIndex int64, fee uint64) (*dextypes.Pool, error)
	GetPoolParamsByDenom(ctx sdk.Context, id string) (types.PoolParams, error)
}
