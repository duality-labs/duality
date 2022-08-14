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
	for _, elem := range genState.VirtualPriceTickQueueList {
		k.SetVirtualPriceTickQueue(ctx, elem)
	}

	// Set virtualPriceTickQueue count
	k.SetVirtualPriceTickQueueCount(ctx, genState.VirtualPriceTickQueueCount)
	// Set all the ticks
	for _, elem := range genState.TicksList {
		k.SetTicks(ctx, elem)
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
	genesis.VirtualPriceTickQueueList = k.GetAllVirtualPriceTickQueue(ctx)
	genesis.VirtualPriceTickQueueCount = k.GetVirtualPriceTickQueueCount(ctx)
	genesis.TicksList = k.GetAllTicks(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
