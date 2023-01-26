package dex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	// Set all the TradingPair
	for _, elem := range genState.TradingPairList {
		k.SetTradingPair(ctx, elem)
	}
	// Set all the tokens
	for _, elem := range genState.TokensList {
		k.SetTokens(ctx, elem)
	}

	// Set tokens count
	k.SetTokensCount(ctx, genState.TokensCount)
	// Set all the tokenMap
	for _, elem := range genState.TokenMapList {
		k.SetTokenMap(ctx, elem)
	}
	// Set all the FeeTier
	for _, elem := range genState.FeeTierList {
		k.SetFeeTier(ctx, elem)
	}

	// Set FeeTier count
	k.SetFeeTierCount(ctx, genState.FeeTierCount)

	// Set all the tickLiquidity
	for _, elem := range genState.TickLiquidityList {
		k.SetTickLiquidity(ctx, elem)
	}
	// Set all the filledLimitOrderTranche
	for _, elem := range genState.FilledLimitOrderTrancheList {
		k.SetFilledLimitOrderTranche(ctx, elem)
	}

	// Set all the LimitOrderTrancheUser
	for _, elem := range genState.LimitOrderTrancheUserList {
		k.SetLimitOrderTrancheUser(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TradingPairList = k.GetAllTradingPair(ctx)
	genesis.TokensList = k.GetAllTokens(ctx)
	genesis.TokensCount = k.GetTokensCount(ctx)
	genesis.TokenMapList = k.GetAllTokenMap(ctx)
	genesis.FeeTierList = k.GetAllFeeTier(ctx)
	genesis.FeeTierCount = k.GetFeeTierCount(ctx)
	genesis.LimitOrderTrancheUserList = k.GetAllLimitOrderTrancheUser(ctx)
	genesis.TickLiquidityList = k.GetAllTickLiquidity(ctx)
	genesis.FilledLimitOrderTrancheList = k.GetAllFilledLimitOrderTranche(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
