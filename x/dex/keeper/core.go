package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) SingleDeposit(goCtx context.Context, token0 string, token1 string, amount sdk.Dec, price sdk.Dec, msg *types.MsgAddLiquidity, callerAdr sdk.AccAddress, receiver sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	/*
			// Check if the pair already exists
			1) Find Pair
			   a) If pair exists, get the pair
			   b) If pair does not exist, error
			2) Find Tick
		       a) If virtual price tick/index does not exist, add new index
				    i) Initialize index to 1 in bitmap
					ii) Create new tick w/ amount in virtual_price_tick_list
					iii) Add tick w/ amount to virtual_price_tick_queue
			   b) If exists
					i) Update tick (+= amount) in virtual_price_tick_list
					ii) Add amount to existing tick in corresponding queue
			3) Update Shares
				i) TBD
	*/

	PairOld, PairFound := k.GetPairs(ctx, token0, token1)

	if !PairFound {
		sdkerrors.Wrapf(types.ErrValidPairNotFound, "Valid pair not found")
	}

	fee, err := sdk.NewDecFromStr(msg.Fee)
	// Error checking for valid sdk.Dec
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	vprice, err := k.CalculateVirtualPrice(token0, token1, msg.TokenDirection, amount, fee, price)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Virtual Price Calculations resulted in a non-valid type: %s", err)
	}

	nearestTickIndex, _ := k.CalculateTick(vprice)
	// TODO:
	// This works if nearestTickIndex's corresponding queue is initialized
	// What is GetVirtualPriceTickQueue's behavior if queue is not initialized

	// TODO: We need to tie each tick queue to the pair, otherwise gets too complex
	// TODO: GetVirtualPriceQueue should take a uint, not a string
	q, _ := k.GetVirtualPriceQueue(ctx, string(nearestTickIndex), msg.TokenDirection, msg.OrderType)

	// TODO: What is ID here?
	// TODO: We need to make ID be tied to a pair, otherwise this recombination is too complex
	// t0 := []byte(token0)
	// t1 := []byte(token1)
	// t2 := []byte(nearestTickIndex)
	// byteArr := append(t0, t1, t2)

	// How do we ensure we're getting the nearest tick price?

	// Queue is uninitialized
	if q.Size() == 0 {
		// Initialize queue
		// newQ := types.VirtualPriceQueue{}
	}

	_ = vprice
	_ = PairOld

	return nil

	_ = ctx
	return nil
}

func (k Keeper) SingleWithdraw(goCtx context.Context, token0 string, token1 string, shares sdk.Dec, msg *types.MsgRemoveLiquidity, callerAdr sdk.AccAddress, receiver sdk.AccAddress) error {
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
