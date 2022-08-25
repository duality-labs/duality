package keeper

import (
	"context"
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SingleDeposit(goCtx context.Context, token0 string, token1 string, amount sdk.Dec, price sdk.Dec, msg *types.MsgAddLiquidity, callerAddr sdk.AccAddress, receiver sdk.AccAddress) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	PairOld, PairFound := k.GetPairs(ctx, token0, token1)

	fmt.Println(PairOld.Tickmap)
	if !PairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Valid pair not found")
	}

	fee, err := sdk.NewDecFromStr(msg.Fee)
	fee = fee.Quo(sdk.NewDec(10000))
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

	IndexQueue, IndexQueueFound := k.GetIndexQueue(ctx, token0, token1, msg.Index)

	// Tick from the tick store
	Tick, TickFound := k.GetTicks(ctx, token0, token1, msg.Price, msg.Fee, msg.OrderType)

	var NewTick types.Ticks
	var oldAmount sdk.Dec //Event variable
	var shares sdk.Dec

	// TODO: Confirm this is the correct way to calculate price
	if msg.TokenDirection == token0 {
		shares = amount.Mul(price.Mul(fee))
	} else {
		shares = amount.Mul(sdk.OneDec().Quo(fee))
	}

	// Index Queue Logic

	if !IndexQueueFound {

		NewQueue := []*types.IndexQueueType{
			&types.IndexQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: shares,
				},
			},
		}
		IndexQueue = types.IndexQueue{
			Index: msg.Index,
			Queue: NewQueue,
		}

	} else {

		if !TickFound {

			// Add tick to the IndexQueue
			IndexQueue.Queue = k.enqueue(ctx, IndexQueue.Queue, types.IndexQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: shares,
				},
			})

		} else {
			tickIndex := -1
			// Do a linear search over the queue to find the tick with the matching price + fee
			for i, tick := range IndexQueue.Queue {
				if tick.Price.Equal(price) && tick.Fee.Equal(fee) {
					tickIndex = i
					break
				}
			}
			if tickIndex == -1 {
				return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Tick not found in queue")
			}

			// Update the existing tick with the new amount
			// Multiple deposits can go to the same tick
			// Need to do this as tick mapping is not tied to an address/unique to a deposit
			IndexQueue.Queue[tickIndex] = &types.IndexQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: Tick.TotalShares.Add(shares),
				},
			}
		}
	}
	//// Tick Logic
	if !TickFound {

		if msg.TokenDirection == token0 {
			NewTick = types.Ticks{
				Price:       msg.Price,
				Fee:         msg.Fee,
				OrderType:   msg.OrderType,
				Reserve0:    amount,
				Reserve1:    sdk.ZeroDec(),
				PairPrice:   price,
				PairFee:     fee,
				TotalShares: shares,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: shares,
				},
			}

			oldAmount = sdk.ZeroDec()
		} else {
			NewTick = types.Ticks{
				Price:       msg.Price,
				Fee:         msg.Fee,
				OrderType:   msg.OrderType,
				Reserve0:    sdk.ZeroDec(),
				Reserve1:    amount,
				PairPrice:   price,
				PairFee:     fee,
				TotalShares: shares,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: shares,
				},
			}
			oldAmount = sdk.ZeroDec()
		}

	} else {
		// If the tick is found, add it to the existing reserve for the tick storage

		if msg.TokenDirection == token0 {
			oldAmount = Tick.Reserve0
			NewTick = types.Ticks{
				Price:       msg.Price,
				Fee:         msg.Fee,
				OrderType:   msg.OrderType,
				Reserve0:    Tick.Reserve0.Add(amount),
				Reserve1:    Tick.Reserve1,
				PairPrice:   price,
				PairFee:     fee,
				TotalShares: Tick.TotalShares.Add(shares),
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: Tick.TotalShares.Add(shares),
				},
			}

		} else {
			oldAmount = Tick.Reserve1
			NewTick = types.Ticks{
				Price:       msg.Price,
				Fee:         msg.Fee,
				OrderType:   msg.OrderType,
				Reserve0:    Tick.Reserve0,
				Reserve1:    Tick.Reserve1.Add(amount),
				PairPrice:   price,
				PairFee:     fee,
				TotalShares: Tick.TotalShares.Add(shares),
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: Tick.TotalShares.Add(shares),
				},
			}
		}

	}

	// Update the storage
	k.SetTicks(ctx, token0, token1, NewTick)
	k.SetIndexQueue(ctx, token0, token1, IndexQueue)

	PairNew, PairFound := k.GetPairs(ctx, token0, token1)

	fmt.Println("Pair New", PairNew, PairFound)
	NewPairs := types.Pairs{
		Token0:       token0,
		Token1:       token1,
		CurrentIndex: PairOld.CurrentIndex,
		TickSpacing:  PairOld.TickSpacing,
		Tickmap:      PairNew.Tickmap,
		IndexMap:     PairNew.IndexMap,
	}

	k.SetPairs(ctx, NewPairs)

	// Sending tokens from the user to the module, might be necessary to do this before the rest of logic to avoid reentrancy/failure attacks
	if msg.TokenDirection == token0 {
		if amount.GT(sdk.ZeroDec()) {
			coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amount.BigInt()))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
				return err
			}
		} else {
			return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Cannnot send zero amount")
		}

	} else {
		if amount.GT(sdk.ZeroDec()) {
			coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount.BigInt()))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
				return err
			}
		} else {
			return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Cannnot send zero amount")
		}
	}

	ctx.EventManager().EmitEvent(types.CreateDepositEvent(msg.Creator,
		token0, token1, price.String(), fee.String(), msg.TokenDirection,
		oldAmount.String(), oldAmount.Add(amount).String(),
		sdk.NewAttribute(types.DepositEventSharesMinted, shares.String()),
	))

	return nil

}

// Can take amount or shares here, depends on what we want to calculate

// Withdraws shares from given price, fee
// Makes more sense, as calculating price & fee can be difficult

// TODO: If withdrawing from one tick with two tokens (i.e. currentTick), will require two withdraw operations

// TODO: Confirm price is always token1/token0, otherwise oldAmount calculation will not work
// TODO: Remove tokenDirection from msg, as it is redundant

/*
Remove Liquidity needs to have verification that the user has enough shares to withdraw & must check re-entrancy attacks
*/
func (k Keeper) SingleWithdraw(goCtx context.Context, token0 string, token1 string, shares sdk.Dec, price sdk.Dec, msg *types.MsgRemoveLiquidity, callerAddr sdk.AccAddress, receiver sdk.AccAddress) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	PairOld, PairFound := k.GetPairs(ctx, token0, token1)

	if !PairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Valid pair not found")
	}

	fee, err := sdk.NewDecFromStr(msg.Fee)
	// Error checking for valid sdk.Dec
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	IndexQueue, IndexQueueFound := k.GetIndexQueue(ctx, token0, token1, msg.Index)

	// Tick from the tick store
	Tick, TickFound := k.GetTicks(ctx, token0, token1, msg.Price, msg.Fee, msg.OrderType)

	var NewTick types.Ticks
	// Index Queue Logic
	removeTick := false
	// Check if tick exists
	if !IndexQueueFound || !TickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Can't withdraw liquidity from a tick that does not exist!, %s", err)

	} else {

		tickIndex := -1
		// Do a linear search over the queue to find the tick with the matching price + fee
		for i, tick := range IndexQueue.Queue {
			if tick.Price.Equal(price) && tick.Fee.Equal(fee) {
				tickIndex = i
				break
			}
		}
		if tickIndex == -1 {
			return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Tick not found in queue")
		}

		// Update the existing tick with the new amount
		// Multiple deposits can go to the same tick
		// Need to do this as tick mapping is not tied to an address/unique to a deposit

		if Tick.TotalShares.GT(shares) {
			IndexQueue.Queue[tickIndex] = &types.IndexQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: Tick.TotalShares.Sub(shares),
				},
			}
		} else {
			// TODO: We should confirm that shares matches the tick amount (to ensure we're not withdrawing more than we have)

			if !Tick.TotalShares.Equal(shares) {
				return sdkerrors.Wrapf(types.ErrNotEnoughShares, "Trying to withdraw more shares than available")
			}
			removeTick = true

			// Remove tick from queue
			IndexQueue.Queue = append(IndexQueue.Queue[:tickIndex], IndexQueue.Queue[tickIndex+1:]...)
		}
	}
	//// Updating Tick Logic
	oldReserve0 := Tick.Reserve0
	oldReserve1 := Tick.Reserve1
	amount0toRemove := Tick.Reserve0
	amount1toRemove := Tick.Reserve1
	if !removeTick {
		// TODO: Decimal precision checks on quotient
		ratio := Tick.Reserve1.Quo(Tick.Reserve0.Add(Tick.Reserve1))
		// r0 * price * 1/(r1/r0+r1)
		amount0toRemove := Tick.Reserve0.Mul(price).Mul(sdk.NewDec(1).Sub(ratio))
		amount1toRemove := Tick.Reserve1.Mul(ratio)

		NewTick = types.Ticks{
			Price:       msg.Price,
			Fee:         msg.Fee,
			OrderType:   msg.OrderType,
			Reserve0:    Tick.Reserve0.Sub(amount0toRemove),
			Reserve1:    Tick.Reserve1.Sub(amount1toRemove),
			PairPrice:   price,
			PairFee:     fee,
			TotalShares: Tick.TotalShares.Sub(shares),
			Orderparams: &types.OrderParams{
				OrderRule:   "",
				OrderType:   msg.OrderType,
				OrderShares: Tick.TotalShares.Sub(shares),
			},
		}

	}

	k.SetIndexQueue(ctx, token0, token1, IndexQueue)
	if removeTick {
		k.RemoveTicks(ctx, token0, token1, msg.Price, msg.Fee, msg.OrderType)
	} else {
		k.SetTicks(ctx, token0, token1, NewTick)
	}

	PairNew, _ := k.GetPairs(ctx, token0, token1)

	NewPairs := types.Pairs{
		Token0:       token0,
		Token1:       token1,
		CurrentIndex: PairOld.CurrentIndex,
		TickSpacing:  PairOld.TickSpacing,
		Tickmap:      PairNew.Tickmap,
		IndexMap:     PairNew.IndexMap,
	}

	k.SetPairs(ctx, NewPairs)

	if !amount0toRemove.GT(sdk.ZeroDec()) && !amount1toRemove.GT(sdk.ZeroDec()) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Cannnot send zero amount")
	}

	// TODO: Sending tokens from the user to the module, will be necessary to do this before the rest of logic to avoid reentrancy/failure attacks
	if amount0toRemove.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amount0toRemove.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	if amount1toRemove.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount1toRemove.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coin1}); err != nil {
			return err
		}
	}

	// TODO: Is this the best format for events with liquidity?
	ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(msg.Creator,
		token0, token1, price.String(), fee.String(), oldReserve0.String(), oldReserve1.String(),
		NewTick.Reserve0.String(), NewTick.Reserve1.String(),
		sdk.NewAttribute(types.WithdrawEventSharesRemoved, shares.String()),
	))

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
