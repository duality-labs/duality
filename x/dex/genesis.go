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
		// Set all the tickObject
		for _, elem2 := range genState.TickObjectList {
			k.SetTickObject(ctx, elem.PairId, elem2)
		}

	}
	// Set all the tokens
	for _, elem := range genState.TokensList {
		k.SetTokens(ctx, elem)
	}

	// Set tokens count
	k.SetTokensCount(ctx, genState.TokensCount)
	// Set all the tokenObject
	for _, elem := range genState.TokenObjectList {
		k.SetTokenObject(ctx, elem)
	}
	// Set all the shares
	for _, elem := range genState.SharesList {
		k.SetShares(ctx, elem)
	}
	// Set all the feeList
	for _, elem := range genState.FeeListList {
		k.SetFeeList(ctx, elem)
	}

	// Set feeList count
	k.SetFeeListCount(ctx, genState.FeeListCount)
	// Set all the edgeRow

	// Set all the limitOrderPoolUserShareMap
	for _, elem := range genState.LimitOrderPoolUserShareMapList {
		k.SetLimitOrderPoolUserShareMap(ctx, elem)
	}
	// Set all the limitOrderPoolUserSharesWithdrawn
	for _, elem := range genState.LimitOrderPoolUserSharesWithdrawnList {
		k.SetLimitOrderPoolUserSharesWithdrawn(ctx, elem)
	}
	// Set all the limitOrderPoolTotalSharesMap
	for _, elem := range genState.LimitOrderPoolTotalSharesMapList {
		k.SetLimitOrderPoolTotalSharesMap(ctx, elem)
	}
	// Set all the limitOrderPoolReserveMap
	for _, elem := range genState.LimitOrderPoolReserveMapList {
		k.SetLimitOrderPoolReserveMap(ctx, elem)
	}
	// Set all the limitOrderPoolFillMap
	for _, elem := range genState.LimitOrderPoolFillMapList {
		k.SetLimitOrderPoolFillMap(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TickObjectList = k.GetAllTickObject(ctx)
	genesis.PairMapList = k.GetAllPairMap(ctx)
	genesis.TokensList = k.GetAllTokens(ctx)
	genesis.TokensCount = k.GetTokensCount(ctx)
	genesis.TokenObjectList = k.GetAllTokenObject(ctx)
	genesis.SharesList = k.GetAllShares(ctx)
	genesis.FeeListList = k.GetAllFeeList(ctx)
	genesis.FeeListCount = k.GetFeeListCount(ctx)
	genesis.LimitOrderPoolUserShareMapList = k.GetAllLimitOrderPoolUserShareMap(ctx)
	genesis.LimitOrderPoolUserSharesWithdrawnList = k.GetAllLimitOrderPoolUserSharesWithdrawn(ctx)
	genesis.LimitOrderPoolTotalSharesMapList = k.GetAllLimitOrderPoolTotalSharesMap(ctx)
	genesis.LimitOrderPoolReserveMapList = k.GetAllLimitOrderPoolReserveMap(ctx)
	genesis.LimitOrderPoolFillMapList = k.GetAllLimitOrderPoolFillMap(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
