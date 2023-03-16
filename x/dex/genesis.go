package dex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the FeeTier
	for _, elem := range genState.FeeTierList {
		k.SetFeeTier(ctx, elem)
	}

	// Set FeeTier count
	k.SetFeeTierCount(ctx, genState.FeeTierCount)

	// Set all the tickLiquidity
	for _, elem := range genState.TickLiquidityList {
		switch elem.Liquidity.(type) {
		case *types.TickLiquidity_PoolReserves:
			k.SetPoolReserves(ctx, *elem.GetPoolReserves())
		case *types.TickLiquidity_LimitOrderTranche:
			k.SetLimitOrderTranche(ctx, *elem.GetLimitOrderTranche())
		}
	}
	// Set all the inactiveLimitOrderTranche
	for _, elem := range genState.InactiveLimitOrderTrancheList {
		k.SetInactiveLimitOrderTranche(ctx, elem)
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
	genesis.FeeTierList = k.GetAllFeeTier(ctx)
	genesis.FeeTierCount = k.GetFeeTierCount(ctx)
	genesis.LimitOrderTrancheUserList = k.GetAllLimitOrderTrancheUser(ctx)
	genesis.TickLiquidityList = k.GetAllTickLiquidity(ctx)
	genesis.InactiveLimitOrderTrancheList = k.GetAllInactiveLimitOrderTranche(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
