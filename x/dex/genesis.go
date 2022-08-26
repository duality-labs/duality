package dex

import (
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	// Set all the pairs
	for _, elem := range genState.PairsList {
		k.SetPairs(ctx, elem)

		// // Set all the ticks
		// for _, elem2 := range genState.TicksList {
		// 	k.SetTicks(ctx, elem.Token0, elem.Token1, elem2)
		// }

		// // Set all the virtualPriceQueue
		// for _, elem2 := range genState.IndexQueueList {
		// 	k.SetIndexQueue(ctx, elem.Token0, elem.Token1, elem2)
		// }

	}

	// Set all the nodes
	for _, elem := range genState.NodesList {
		k.SetNodes(ctx, elem)
	}
	// Set all the shares
	for _, elem := range genState.SharesList {
		k.SetShares(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.TicksList = k.GetAllTicks(ctx)
	genesis.PairsList = k.GetAllPairs(ctx)
	genesis.IndexQueueList = k.GetAllIndexQueue(ctx)
	genesis.NodesList = k.GetAllNodes(ctx)
	genesis.SharesList = k.GetAllShares(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
