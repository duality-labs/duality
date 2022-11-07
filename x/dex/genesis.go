package dex

import (
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	// Set all the pairObject
	for _, elem := range genState.PairObjectList {
		k.SetPairObject(ctx, elem)
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

	// Set all the limitOrderPoolUserShareObject
	for _, elem := range genState.LimitOrderPoolUserShareObjectList {
		k.SetLimitOrderPoolUserShareObject(ctx, elem)
	}
	// Set all the limitOrderPoolUserSharesWithdrawn
	for _, elem := range genState.LimitOrderPoolUserSharesWithdrawnList {
		k.SetLimitOrderPoolUserSharesWithdrawn(ctx, elem)
	}
	// Set all the limitOrderPoolTotalSharesObject
	for _, elem := range genState.LimitOrderPoolTotalSharesObjectList {
		k.SetLimitOrderPoolTotalSharesObject(ctx, elem)
	}
	// Set all the limitOrderPoolReserveObject
	for _, elem := range genState.LimitOrderPoolReserveObjectList {
		k.SetLimitOrderPoolReserveObject(ctx, elem)
	}
	// Set all the limitOrderPoolFillObject
	for _, elem := range genState.LimitOrderPoolFillObjectList {
		k.SetLimitOrderPoolFillObject(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TickObjectList = k.GetAllTickObject(ctx)
	genesis.PairObjectList = k.GetAllPairObject(ctx)
	genesis.TokensList = k.GetAllTokens(ctx)
	genesis.TokensCount = k.GetTokensCount(ctx)
	genesis.TokenObjectList = k.GetAllTokenObject(ctx)
	genesis.SharesList = k.GetAllShares(ctx)
	genesis.FeeListList = k.GetAllFeeList(ctx)
	genesis.FeeListCount = k.GetFeeListCount(ctx)
	genesis.LimitOrderPoolUserShareObjectList = k.GetAllLimitOrderPoolUserShareObject(ctx)
	genesis.LimitOrderPoolUserSharesWithdrawnList = k.GetAllLimitOrderPoolUserSharesWithdrawn(ctx)
	genesis.LimitOrderPoolTotalSharesObjectList = k.GetAllLimitOrderPoolTotalSharesObject(ctx)
	genesis.LimitOrderPoolReserveObjectList = k.GetAllLimitOrderPoolReserveObject(ctx)
	genesis.LimitOrderPoolFillObjectList = k.GetAllLimitOrderPoolFillObject(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
