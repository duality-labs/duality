package dex

import (
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the ticks
	for _, elem := range genState.TicksList {
		k.SetTicks(ctx, elem)
	}
	// Set all the share
	for _, elem := range genState.ShareList {
		k.SetShare(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TicksList = k.GetAllTicks(ctx)
	genesis.ShareList = k.GetAllShare(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
