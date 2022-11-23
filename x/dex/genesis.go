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
	for _, elem := range genState.EdgeRowList {
		k.SetEdgeRow(ctx, elem)
	}

	// Set edgeRow count
	k.SetEdgeRowCount(ctx, genState.EdgeRowCount)
	// Set all the adjanceyMatrix
	for _, elem := range genState.AdjanceyMatrixList {
		k.SetAdjanceyMatrix(ctx, elem)
	}

	// Set adjanceyMatrix count
	k.SetAdjanceyMatrixCount(ctx, genState.AdjanceyMatrixCount)
	// Set all the LimitOrderTrancheUser
	for _, elem := range genState.LimitOrderTrancheUserList {
		k.SetLimitOrderTrancheUser(ctx, elem)
	}
	// Set all the LimitOrderTranche
	for _, elem := range genState.LimitOrderTrancheList {
		k.SetLimitOrderTranche(ctx, elem)
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
	genesis.SharesList = k.GetAllShares(ctx)
	genesis.FeeListList = k.GetAllFeeList(ctx)
	genesis.FeeListCount = k.GetFeeListCount(ctx)
	genesis.EdgeRowList = k.GetAllEdgeRow(ctx)
	genesis.EdgeRowCount = k.GetEdgeRowCount(ctx)
	genesis.AdjanceyMatrixList = k.GetAllAdjanceyMatrix(ctx)
	genesis.AdjanceyMatrixCount = k.GetAdjanceyMatrixCount(ctx)
	genesis.LimitOrderTrancheUserList = k.GetAllLimitOrderTrancheUser(ctx)
	genesis.LimitOrderTrancheList = k.GetAllLimitOrderTranche(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
