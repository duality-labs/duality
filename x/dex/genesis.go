package dex

import (
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	// Set all the pairMap
	for _, elem := range genState.PairMapList {
		k.SetPairMap(ctx, elem)
		// Set all the tickMap
		for _, elem2 := range genState.TickMapList {
			k.SetTickMap(ctx, elem.PairId, elem2)
		}
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
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TickMapList = k.GetAllTickMap(ctx)
	genesis.PairMapList = k.GetAllPairMap(ctx)
	genesis.TokensList = k.GetAllTokens(ctx)
	genesis.TokensCount = k.GetTokensCount(ctx)
	genesis.TokenMapList = k.GetAllTokenMap(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
