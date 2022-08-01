package keeper

import (
	"context"

	dextypes "github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/NicholasDotSol/duality/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Specifications of types.Msg.Swap can be found in ../proto/router/tx.proto
func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	// Sorts token0, token1 for uniformity for use in internal mappings
	token0, token1, err := k.dexKeeper.SortTokens(ctx, msg.TokenIn, msg.TokenOut)

	// Error handling for sortTokens
	if err != nil {
		return nil, err
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)

	// Error checking for the calling address
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Converts msg.AmountIn (string) to sdk.Dec
	amountIn, err := sdk.NewDecFromStr(msg.AmountIn)
	// Error handling for sdk.Dec
	if err != nil {
		return nil, err
	}

	/* GetBalance is an external (x/bank) module function called to check the token balanced for a specified address
	  GetBalance returns a sdk.Coin: (string denom sdk.Int amount)

	  Note: sdk.Int is not the same as native Int but is a override of big/int

	  AccountsTokenInBalance is the sdk.Dec representation of the return sdk.Coin's amount of tokenIn.
	*/
	AccountsTokenInBalance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.TokenIn).Amount)

	// Error handling to verify the amount wished to swap is NOT more then the msg.creator holds in their accounts
	if AccountsTokenInBalance.LT(amountIn) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s does not have enough  of token In", callerAddr)
	}

	// Returns the token pair, and if a token pair was found
	oldTick, tickFound := k.dexKeeper.GetTicks(ctx, token0, token1)

	// For our naive (greedy) implementation if the token pair does not exist we return and error
	if !tickFound {
		return nil, sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found")
	}
	
	// Initializes remainingAmount to swap to amountIn
	remainingAmount := amountIn
	// Initializes TotalAmountOut to zero
	TotalAmountOut := sdk.ZeroDec()

	/*
	We first need to determine which priority queues are relevant: PoolsOneToZero, PoolsZeroToOne
	If token0 == TokenIn we look at the PoolsZeroToOne to determine a greedy swap
	If token1 == TokenIn we look at PoolsOneToZero to determine a greedy swap
	*/

	if token0 == msg.TokenIn {
		// Check to make sure the token pair has liquid pools
		if len(oldTick.PoolsZeroToOne) != 0 {

			// Calculate the amount of Reserve1 available given our virtual price
			RequiredToDeplete := oldTick.PoolsZeroToOne[0].Reserve1.Add(oldTick.PoolsZeroToOne[0].Reserve1.Mul(oldTick.PoolsZeroToOne[0].Fee.Quo(oldTick.PoolsZeroToOne[0].Price.Mul(sdk.NewDec(10000))))) // RequiredToDeplete = ReserveB + ReserveB * (fee / (Pricec * 10000))
			
			// While Loop, that loops until either:
			// 	1. remainingAmount is zero, the sender has provided all of tokenIn desired 
			//  2. All pools for that token pair have been emptied
			for ( !(remainingAmount.Equal(sdk.ZeroDec())) && len(oldTick.PoolsZeroToOne) !=0 ) {

				// All remaining amount of tokenIn can be depleted at this pool
				if (remainingAmount.LT(RequiredToDeplete)) {
					
					// Calculates the amount out at specified pool given, relevant pool price and fee
					AmountOut := remainingAmount.Sub( remainingAmount.Mul(oldTick.PoolsZeroToOne[0].Fee.Quo(oldTick.PoolsZeroToOne[0].Price.Mul(sdk.NewDec(10000))))  )
					// Calculates new Reserve1 to be Old Reserve1 - AmountOut
					NewReserve1 := oldTick.PoolsZeroToOne[0].Reserve1.Sub(AmountOut)
					
					// Update ZeroToOne Priority queue with changes to Reserve0, Reserve1,
					k.dexKeeper.Update0to1(&oldTick.PoolsZeroToOne, oldTick.PoolsZeroToOne[0],  
						oldTick.PoolsZeroToOne[0].Reserve0.Add(remainingAmount), NewReserve1, oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].TotalShares, oldTick.PoolsZeroToOne[0].Price)
					
					// Determine if the pool exists in alternate direction priority queue
					oldOneToZeroPool, OneToZeroPoolFound := k.dexKeeper.GetPool(&oldTick.PoolsOneToZero, oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].Price )

					// If the pool in alternate direction priority queue exists, we update changes to it as well
					// Note: As we have already added remainingAmount to reserve0 above, we only need to pass in reserve0 at this stage
					if OneToZeroPoolFound {
						k.dexKeeper.Update1to0(&oldTick.PoolsOneToZero, &oldOneToZeroPool,
							oldTick.PoolsZeroToOne[0].Reserve0, NewReserve1, oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].TotalShares, oldTick.PoolsZeroToOne[0].Price)
					
					// If pool does not exists, create a newPool and push it to the PoolsOneToZero priority queue
					} else {
						NewPool := dextypes.Pool{
							Reserve0: remainingAmount,
							Reserve1: NewReserve1,
							Price: oldTick.PoolsZeroToOne[0].Price,
							Fee: oldTick.PoolsZeroToOne[0].Fee,
							TotalShares:  oldTick.PoolsZeroToOne[0].TotalShares ,
							Index: 0,
						}

						k.dexKeeper.Push1to0(&oldTick.PoolsOneToZero, &NewPool)
					}

					// Initializes new tick object
					NewTick := dextypes.Ticks {
						token0,
						token1,
						oldTick.PoolsZeroToOne,
						oldTick.PoolsOneToZero,
					}

					// Sets changes to priority queue to memory
					k.dexKeeper.SetTicks(
						ctx,
						NewTick,
					)

					// Appends Amount out to the totalAmount of tokenOut 
					TotalAmountOut = TotalAmountOut.Add(AmountOut)
				
					remainingAmount = sdk.ZeroDec()
					
				 // If remainingAmount is not less than the amount required to deplete, we calculate the maximal amount we can take from this pool, and then pop the pool from our priority queue
				 // Note by conditionalizing based of the requiredToDeplete price we are able to remove exactly reserve1 from the pool (the maximal amount), allowing us to then pop it without any loss
				 } else {
					
					// Calculates the amount out at specified pool given, relevant pool price and fee
					AmountOut := oldTick.PoolsZeroToOne[0].Reserve1.Sub( oldTick.PoolsZeroToOne[0].Reserve1.Mul(oldTick.PoolsZeroToOne[0].Fee.Quo(oldTick.PoolsZeroToOne[0].Price.Mul(sdk.NewDec(10000))))  )

					// Returns whether the pool exists, and if so the pool in the alternate direction priority queue
					oldOneToZeroPool, OneToZeroPoolFound := k.dexKeeper.GetPool(&oldTick.PoolsOneToZero, oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].Price )

					// If the pool in the alternate direction exists we update it with our changes: adding reserves1 amount of token0 to reserves0 and setting reserves1 to zero
					// Note we must pop the pool after OneToZero to avoid null pointer errors
					if OneToZeroPoolFound {
						k.dexKeeper.Update1to0(&oldTick.PoolsOneToZero, &oldOneToZeroPool,
							oldTick.PoolsZeroToOne[0].Reserve0.Add(oldTick.PoolsZeroToOne[0].Reserve1), sdk.ZeroDec(), oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].TotalShares, oldTick.PoolsZeroToOne[0].Price)
					
					// If the pool in the alternate direction does not exists, we create a new pool object based off our changes and push it to the priority queue
					} else {
						NewPool := dextypes.Pool{
							Reserve0: oldTick.PoolsZeroToOne[0].Reserve1,
							Reserve1: sdk.ZeroDec(),
							Price: oldTick.PoolsZeroToOne[0].Price,
							Fee: oldTick.PoolsZeroToOne[0].Fee,
							TotalShares:  oldTick.PoolsZeroToOne[0].TotalShares ,
							Index: 0,
						}

						k.dexKeeper.Push1to0(&oldTick.PoolsOneToZero, &NewPool)
					}

					// Pop ZeroToOne pool
					k.dexKeeper.Pop0to1(&oldTick.PoolsZeroToOne)

					//Intializes a NewTick object to set in the memory
					NewTick := dextypes.Ticks {
						token0,
						token1,
						oldTick.PoolsZeroToOne,
						oldTick.PoolsOneToZero,
					}

					// Update the tick mapping in memory
					k.dexKeeper.SetTicks(
						ctx,
						NewTick,
					)

					// Append and Subtract AmountOut from TotalAmountOut and remainingAmount respectively 
					TotalAmountOut = TotalAmountOut.Add(AmountOut)
					remainingAmount = remainingAmount.Sub(AmountOut)				

				}
			}
		
		// If the TokenPair does not have any valid pools, return an error
		// Note this is purely the case for our naive greedy implementation.
		} else {
			return nil, sdkerrors.Wrapf(types.ErrValidPathNotFound, "Valid Path not found")
		}

	// If token1 == TokenIn we look at PoolsOneToZero to determine a greedy swap
	} else {

		// Check to make sure the token pair has liquid pools
		if len(oldTick.PoolsOneToZero) != 0 {
			// Calculate the amount of Reserve0 available given our virtual price
			RequiredToDeplete := oldTick.PoolsOneToZero[0].Reserve0.Add(oldTick.PoolsOneToZero[0].Reserve0.Mul(oldTick.PoolsOneToZero[0].Price.Mul(oldTick.PoolsOneToZero[0].Fee).Quo(sdk.NewDec(10000)))) 
			
			// While Loop, that loops until either:
			// 	1. remainingAmount is zero, the sender has provided all of tokenIn desired 
			//  2. All pools for that token pair have been emptied
			for (!(remainingAmount.Equal(sdk.ZeroDec())) || len(oldTick.PoolsOneToZero) ==0 ) {
				
				// All remaining amount of tokenIn can be depleted at this pool
				if (remainingAmount.LT(RequiredToDeplete)) {
					
					// Calculates the amount out at specified pool given relevant pool price and fee
					AmountOut := remainingAmount.Sub( remainingAmount.Mul(oldTick.PoolsOneToZero[0].Fee.Mul(oldTick.PoolsOneToZero[0].Price).Quo(sdk.NewDec(10000)))  )
					
					// Calculates new Reserve0 to be Old Reserve0 - AmountOut
					NewReserve0 := oldTick.PoolsOneToZero[0].Reserve0.Sub(AmountOut)
					
					// Update OneToZero Priority queue with changes to Reserve0, Reserve1,
					k.dexKeeper.Update1to0(&oldTick.PoolsOneToZero, oldTick.PoolsOneToZero[0],  NewReserve0,
						oldTick.PoolsOneToZero[0].Reserve1.Add(remainingAmount), oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].TotalShares, oldTick.PoolsOneToZero[0].Price)
					
					// Determine if the pool exists in alternate direction priority queue
					oldZeroToOnePool, ZeroToOnePoolFound := k.dexKeeper.GetPool(&oldTick.PoolsZeroToOne, oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].Price )

					// If the pool in alternate direction priority queue exists, we update changes to it as well
					// Note: As we have already added remainingAmount to reserve1 above, we only need to pass in reserve0 at this stage
					if ZeroToOnePoolFound {
						k.dexKeeper.Update0to1(&oldTick.PoolsOneToZero, &oldZeroToOnePool, NewReserve0,
							oldTick.PoolsOneToZero[0].Reserve1,  oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].TotalShares, oldTick.PoolsOneToZero[0].Price)
					
					// If pool does not exists, create a newPool and push it to the PoolsZeroToOne priority queue
					} else {
						NewPool := dextypes.Pool{
							Reserve0: NewReserve0,
							Reserve1: remainingAmount ,
							Price: oldTick.PoolsOneToZero[0].Price,
							Fee: oldTick.PoolsOneToZero[0].Fee,
							TotalShares:  oldTick.PoolsOneToZero[0].TotalShares ,
							Index: 0,
						}
						
						k.dexKeeper.Push1to0(&oldTick.PoolsZeroToOne, &NewPool)
					}

					// Initializes new tick object
					NewTick := dextypes.Ticks {
						token0,
						token1,
						oldTick.PoolsZeroToOne,
						oldTick.PoolsOneToZero,
					}

					// Sets changes to priority queue to memory
					k.dexKeeper.SetTicks(
						ctx,
						NewTick,
					)

					// Appends Amount out to the totalAmount of tokenOut 
					TotalAmountOut = TotalAmountOut.Add(AmountOut)

					remainingAmount = sdk.ZeroDec()
					
				 // If remainingAmount is not less than the amount required to deplete, we calculate the maximal amount we can take from this pool, and then pop the pool from our priority queue
				 // Note by conditionalizing based of the requiredToDeplete price we are able to remove exactly reserve0 from the pool (the maximal amount), allowing us to then pop it without any loss
				} else {
					
					// Calculates the amount out at specified pool for the relevant pool price and fee
					AmountOut := oldTick.PoolsOneToZero[0].Reserve0.Sub( oldTick.PoolsOneToZero[0].Reserve0.Mul(oldTick.PoolsOneToZero[0].Fee.Mul(oldTick.PoolsOneToZero[0].Price).Quo(sdk.NewDec(10000)))  )

					// Returns whether the pool exists, and if so the pool in the alternate direction priority queue
					oldZeroToOnePool, ZeroToOnePoolFound := k.dexKeeper.GetPool(&oldTick.PoolsZeroToOne, oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].Price )


					// If the pool in the alternate direction exists we update it with our changes: adding reserves1 amount of token0 to reserves0 and setting reserves1 to zero
					// Note we must pop the pool after OneToZero to avoid null pointer errors
					if ZeroToOnePoolFound {
						k.dexKeeper.Update1to0(&oldTick.PoolsOneToZero, &oldZeroToOnePool, sdk.ZeroDec(),
							oldTick.PoolsOneToZero[0].Reserve1.Add(oldTick.PoolsOneToZero[0].Reserve0),  oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].TotalShares, oldTick.PoolsOneToZero[0].Price)
					
							// If the pool in the alternate direction does not exists, we create a new pool object based off our changes and push it to the priority queue
					} else {
						NewPool := dextypes.Pool{
							Reserve0: sdk.ZeroDec(),
							Reserve1: oldTick.PoolsOneToZero[0].Reserve0,
							Price: oldTick.PoolsOneToZero[0].Price,
							Fee: oldTick.PoolsOneToZero[0].Fee,
							TotalShares:  oldTick.PoolsOneToZero[0].TotalShares ,
							Index: 0,
						}

						k.dexKeeper.Push1to0(&oldTick.PoolsZeroToOne, &NewPool)
					}

					// Pop ZeroToOne pool
					k.dexKeeper.Pop0to1(&oldTick.PoolsOneToZero)


					//Intializes a NewTick object to set in the memory
					NewTick := dextypes.Ticks {
						token0,
						token1,
						oldTick.PoolsZeroToOne,
						oldTick.PoolsOneToZero,
					}

					// Sets changes to priority queue to memory
					k.dexKeeper.SetTicks(
						ctx,
						NewTick,
					)

					// Appends Amount out to the totalAmount of tokenOut 
					TotalAmountOut = TotalAmountOut.Add(AmountOut)
					remainingAmount = remainingAmount.Sub(AmountOut)				


				}
			}
		} else {
			return nil, sdkerrors.Wrapf(types.ErrValidPathNotFound, "Valid Path not found")
		}

	}

	
	// Converts msg.MinOut (string) to sdk.Dec
	minOut, err := sdk.NewDecFromStr(msg.MinOut)
	// Error handling for sdk.Dec
	if err != nil {
		return nil, err
	}

	// Error handling if TotalAmount is less than minOut
	if TotalAmountOut.LT(minOut) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Total Amount is less than specified minimum amount out: %s", minOut.String())
	}

	
	// If AmountIn is greater than zero, send amountIn of tokenIn from caller Address to the dex module
	// Note this relies on a x/bank call SendCoinsFromAccountToModule
	if amountIn.GT(sdk.ZeroDec()) {
		coinIn := sdk.NewCoin(msg.TokenIn, sdk.NewIntFromBigInt(amountIn.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, dextypes.ModuleName, sdk.Coins{coinIn}); err != nil {
			return nil, err
		}
	}

	// Given that totalAmount is greater than zero, send totalAmountOut of tokenOut to the caller Address
	// Note this relies on a x/bank call SendCoinsFromModuleToAccount
	if TotalAmountOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(msg.TokenOut, sdk.NewIntFromBigInt(TotalAmountOut.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, dextypes.ModuleName, callerAddr, sdk.Coins{coinOut}); err != nil {
			return nil, err
		}
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, token0, token1, msg.AmountIn, TotalAmountOut.String()))

	
	_ = ctx

	return &types.MsgSwapResponse{}, nil
}
