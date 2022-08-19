package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Decide whether to addLiquidity if pair exists
// TODO: Add current tick specification for pair multi-deposit
// TODO: Determine how we plan to set tick spacing for pair
func (k Keeper) CreateNewPair(goCtx context.Context, token0 string, token1 string, amount sdk.Dec, msg *types.MsgCreatePair, callerAdr sdk.AccAddress, receiver sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	/*
		1) Check if pair exists
		   a) If so, output pair
		   b) Else, init pair
		       i) If nodes do not exist, init nodes
			   ii) Add tokenA, tokenB to eachother's outgoingEdges
		2) Call SingleDeposit on pool & set currTick equivalent to corresponding virtualTick (for price, fee)
	*/

	_ = ctx
	return nil
}

// price is in terms of token1/token0, whereas for msg.Price we have no guarantees on whether tokenA == token0, so we need to use price
func (k Keeper) SingleWithdraw(goCtx context.Context, token0 string, token1 string, shares sdk.Dec, price sdk.Dec, msg *types.MsgRemoveLiquidity, callerAdr sdk.AccAddress, receiver sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	/*
			1) Find Pair
			   a) If pair exists, get the pair
			   b) If pair does not exist, error

			2) Find Tick
		       a) If virtual price tick/index does not exist, exit
			   b) If exists
					i) Update tick (-= amount) in virtual_price_tick_list
					ii) Subtract amount from existing tick in corresponding queue
					iii) If tick cleared, uninitialize index in bitmap, remove from virtual_price_tick_list & remove from queue
			3) Update Shares
				i) TBD
	*/
	_ = ctx
	return nil
}

// Need to figure out logic for route vs. swap
func (k Keeper) SingleSwapIn(goCtx context.Context, token0 string, token1 string, amountIn sdk.Dec, msg *types.MsgSwap, callerAdr sdk.AccAddress, receiver sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	/*
		1) Find Pair
		   a) If pair exists, get the pair
		   b) If pair does not exist, error
		2) Get CurrTick & corresponding list for direction
		3) Attempt to swap amount through the ticks in pair
			i) Loop through queue for virtual tick & empty ticks
			ii) If queue empty, query next virtualTick from bitmap
			iii) Continue looping until amount == 0
			iv) Store last tick, will be new currTick
		4) Perform swap
		5) Update CurrTick
		6) Update Shares
			i) TBD
	*/

	_ = ctx
	return nil
}
