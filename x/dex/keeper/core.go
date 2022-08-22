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

	ctx := sdk.UnwrapSDKContext(goCtx)

	PairOld, PairFound := k.GetPairs(ctx, token0, token1)

	if !PairFound {
		sdkerrors.Wrapf(types.ErrValidPairNotFound, "Valid pair not found")
	}

	fee, err := sdk.NewDecFromStr(msg.Fee)
	// Error checking for valid sdk.Dec
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	// Can only deposit amount0 where vPrice >= CurrentPrice
	if msg.Index < (PairOld.CurrentIndex) && msg.TokenDirection == token0 {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit token0 at a price/fee pair less than the current price")
		// Can only deposit amount1 where CurrentPrice >= vPrice
	} else if PairOld.CurrentIndex < msg.Index && msg.TokenDirection == token1 {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit token1 at a price/fee pair greater than the current price")
	}

	IndexQueueOld, IndexQueueFound := k.GetIndexQueue(ctx, token0, token1, msg.Index)

	//FIX ME
	shares := amount.Mul(price)

	if !IndexQueueFound {

		NewQueue := []*types.IndexQueueType{
			{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: shares,
				},
			},
		}
		IndexQueueOld = types.IndexQueue{
			Index: msg.Index,
			Queue: NewQueue,
		}

	} else {
		TickIndexFound := -1
		for i := 0; i < len(IndexQueueOld.Queue); i++ {
			if IndexQueueOld.Queue[i].Price.Equal(price) && IndexQueueOld.Queue[i].Fee.Equal(fee) && IndexQueueOld.Queue[i].Orderparams.OrderType == msg.OrderType {
				TickIndexFound = i
				break
			}
		}

		if TickIndexFound != -1 {

			IndexQueueOld.Queue[TickIndexFound] = &types.IndexQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: IndexQueueOld.Queue[TickIndexFound].Orderparams.OrderShares.Add(shares),
				},
			}
		} else {

			IndexQueueOld.Queue = k.enqueue(ctx, IndexQueueOld.Queue, types.IndexQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: shares,
				},
			})
		}
	}

	k.SetPairs(ctx, types.Pairs{
		Token0:       token0,
		Token1:       token1,
		TickSpacing:  PairOld.TickSpacing,
		CurrentIndex: PairOld.CurrentIndex,
		Tickmap:      PairOld.Tickmap,
		IndexMap:     &IndexQueueOld,
	},
	)

	return nil
}

// price is in terms of token1/token0, whereas for msg.Price we have no guarantees on whether tokenA == token0, so we need to use price

// Updates the internal data structs for withdraw operations
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
	PairOld, PairFound := k.GetPairs(ctx, token0, token1)

	if !PairFound {
		sdkerrors.Wrapf(types.ErrValidPairNotFound, "Valid pair not found")
	}

	fee, err := sdk.NewDecFromStr(msg.Fee)
	// Error checking for valid sdk.Dec
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	IndexQueueOld, IndexQueueFound := k.GetIndexQueue(ctx, token0, token1, msg.Index)

	//FIX ME
	if !IndexQueueFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Cannot withdraw shares from a price/fee pair that does not exist")

	} else {
		TickIndexFound := -1
		for i := 0; i < len(IndexQueueOld.Queue); i++ {
			if IndexQueueOld.Queue[i].Price.Equal(price) && IndexQueueOld.Queue[i].Fee.Equal(fee) && IndexQueueOld.Queue[i].Orderparams.OrderType == msg.OrderType {
				TickIndexFound = i
				break
			}
		}

		if TickIndexFound != -1 {
			// Assumes we've already done pre-verification on the number of shares to withdraw & that the user
			// If withdrawing more shares than available, this should be caught before this function is called
			if IndexQueueOld.Queue[TickIndexFound].Orderparams.OrderShares.Sub(shares).Equal(sdk.ZeroDec()) {
				// No simple way to remove element from queue, so we'll just set the tick to 0 & remove from list
				tickToRemove := IndexQueueOld.Queue[TickIndexFound]
				// Set this to empty for now
				_ = tickToRemove

				// Remove from queue
				IndexQueueOld.Queue = append(IndexQueueOld.Queue[:TickIndexFound], IndexQueueOld.Queue[TickIndexFound+1:]...)
			} else {
				// Subtract respective shares from this tick
				IndexQueueOld.Queue[TickIndexFound] = &types.IndexQueueType{
					Price: price,
					Fee:   fee,
					Orderparams: &types.OrderParams{
						OrderRule:   "",
						OrderType:   msg.OrderType,
						OrderShares: IndexQueueOld.Queue[TickIndexFound].Orderparams.OrderShares.Sub(shares),
					},
				}
			}
		} else {
			sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found in corresponding queue (index is incorrectly called)")
		}
	}

	k.SetPairs(ctx, types.Pairs{
		Token0:       token0,
		Token1:       token1,
		TickSpacing:  PairOld.TickSpacing,
		CurrentIndex: PairOld.CurrentIndex,
		Tickmap:      PairOld.Tickmap,
		IndexMap:     &IndexQueueOld,
	},
	)

	return nil
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
