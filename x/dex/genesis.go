package dex

import (
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the nodes
	for _, elem := range genState.NodesList {
		k.SetNodes(ctx, elem)
	}

	// Set nodes count
	k.SetNodesCount(ctx, genState.NodesCount)
	// Set all the virtualPriceTickQueue

	// Set virtualPriceTickQueue count

	// Set all the ticks
	for _, elem := range genState.TicksList {
		k.SetTicks(ctx, elem)
	}
	// Set all the virtualPriceTickList

	// Set all the bitArr
	for _, elem := range genState.BitArrList {
		k.SetBitArr(ctx, elem)
	}

	// Set bitArr count
	k.SetBitArrCount(ctx, genState.BitArrCount)
	// Set all the pairs
	for _, elem := range genState.PairsList {
		k.SetPairs(ctx, elem)
	}
	// Set all the virtualPriceQueue
	for _, elem := range genState.VirtualPriceQueueList {
		k.SetVirtualPriceQueue(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.NodesList = k.GetAllNodes(ctx)
	genesis.NodesCount = k.GetNodesCount(ctx)

	genesis.TicksList = k.GetAllTicks(ctx)

	genesis.BitArrList = k.GetAllBitArr(ctx)
	genesis.BitArrCount = k.GetBitArrCount(ctx)
	genesis.PairsList = k.GetAllPairs(ctx)
	genesis.VirtualPriceQueueList = k.GetAllVirtualPriceQueue(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
