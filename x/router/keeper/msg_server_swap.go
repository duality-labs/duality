package keeper

import (
	"context"
	"fmt"
	//"sort"
	//"fmt"

	dextypes "github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/NicholasDotSol/duality/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//k.dexKeeper.GetAllTicks(ctx)

	
	token0, token1, err := k.dexKeeper.SortTokens(ctx, msg.TokenIn, msg.TokenOut)

	if err != nil {
		return nil, err
	}

	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	amountIn, err := sdk.NewDecFromStr(msg.AmountIn)
	if err != nil {
		return nil, err
	}


	AccountsTokenInBalance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.TokenIn).Amount)
	if AccountsTokenInBalance.LT(amountIn) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s does not have enough  of token In", callerAddr)
	}

	oldTick, tickFound := k.dexKeeper.GetTicks(ctx, token0, token1)

	if !tickFound {
		return nil, err
	}

	// Token A Token B
	// PoolZeroToOne [ {ReserveA : 100, ReserveB: 100, Fee = .2} ,  {ReserveA : 50, ReserveB: 0, Fee=.8}]
	//PoolOneToZero  [{ReserveA : 100, ReserveB: 100, Fee= .2} ]

	// Swap AtoB
	// TokenIN: A
	// TokenOut: B
	// AmountIN: 75
	// minAmount: 50
	// RemainingAmount = AmountIN
	// Head of PQ
	// while (remaingAmount != 0) 

	// RequiredToDeplete = ReserveB + ReserveB * (fee / (Pricec * 10000))

	//  if RemainingAmount < RequiredToDeplete
		
		// AmountOut =  RemainingAmount - RemainingAmount * (fee /  (Price * 10000))
		// NewReserveB = ReserveB - AmountOut
		
		// NewReserveA = ReserveA +  RemainingAmount
		// TotalAmountOut += AmountOut
		// Update0to1
		// RemainingAmount = 0 
		// Update1to0 

	// else RemainigAmount >= RequiredToDeplete
		// PQ.POP()

		// NewReserveB = 0
		// AmountOut =  ReserveB - ReserveB * Fee / (Price * 10000)
		// NewReserveA = ReserveA + AmountOut
		// TotalAmountOut += AmountOut
		// RemainingAmount = RemainingAmount -= AmountOut
		// Update1to0

	//if TotalAmountOut < minOut
		// return nil, error
	
	// Transfer totalAmountIN
	// transfer TotalAmountOUt
	

	remainingAmount, err := sdk.NewDecFromStr(msg.AmountIn)
	if err != nil {
		return nil, err
	}
	TotalAmountOut := sdk.ZeroDec()

	
	if token0 == msg.TokenIn {
		if len(oldTick.PoolsZeroToOne) != 0 {
			
			for ( !(remainingAmount.Equal(sdk.ZeroDec())) && len(oldTick.PoolsZeroToOne) !=0 ) {
				if (remainingAmount.LT(oldTick.PoolsZeroToOne[0].Reserve1)) {
					
					AmountOut := remainingAmount.Sub( remainingAmount.Mul(oldTick.PoolsZeroToOne[0].Fee.Quo(oldTick.PoolsZeroToOne[0].Price.Mul(sdk.NewDec(10000))))  )
					NewReserve1 := oldTick.PoolsZeroToOne[0].Reserve1.Sub(AmountOut)
					
					
					k.dexKeeper.Update0to1(&oldTick.PoolsZeroToOne, oldTick.PoolsZeroToOne[0],  
						oldTick.PoolsZeroToOne[0].Reserve0.Add(remainingAmount), NewReserve1, oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].TotalShares, oldTick.PoolsZeroToOne[0].Price)
					
					oldOneToZeroPool, OneToZeroPoolFound := k.dexKeeper.GetPool(&oldTick.PoolsOneToZero, oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].Price )

					if OneToZeroPoolFound {
						k.dexKeeper.Update1to0(&oldTick.PoolsOneToZero, &oldOneToZeroPool,
							oldTick.PoolsZeroToOne[0].Reserve0.Add(remainingAmount), NewReserve1, oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].TotalShares, oldTick.PoolsZeroToOne[0].Price)
					
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

					NewTick := dextypes.Ticks {
						token0,
						token1,
						oldTick.PoolsZeroToOne,
						oldTick.PoolsOneToZero,
					}

					k.dexKeeper.SetTicks(
						ctx,
						NewTick,
					)

					TotalAmountOut = TotalAmountOut.Add(AmountOut)
					remainingAmount = sdk.ZeroDec()
					
				} else {
					
					AmountOut := oldTick.PoolsZeroToOne[0].Reserve1.Sub( oldTick.PoolsZeroToOne[0].Reserve1.Mul(oldTick.PoolsZeroToOne[0].Fee.Quo(oldTick.PoolsZeroToOne[0].Price.Mul(sdk.NewDec(10000))))  )
					

					

					oldOneToZeroPool, OneToZeroPoolFound := k.dexKeeper.GetPool(&oldTick.PoolsOneToZero, oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].Price )

					if OneToZeroPoolFound {
						k.dexKeeper.Update1to0(&oldTick.PoolsOneToZero, &oldOneToZeroPool,
							oldTick.PoolsZeroToOne[0].Reserve0.Add(oldTick.PoolsZeroToOne[0].Reserve1), sdk.ZeroDec(), oldTick.PoolsZeroToOne[0].Fee, oldTick.PoolsZeroToOne[0].TotalShares, oldTick.PoolsZeroToOne[0].Price)
					
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

					NewTick := dextypes.Ticks {
						token0,
						token1,
						oldTick.PoolsZeroToOne,
						oldTick.PoolsOneToZero,
					}

					k.dexKeeper.Pop0to1(&oldTick.PoolsZeroToOne)

					k.dexKeeper.SetTicks(
						ctx,
						NewTick,
					)

					TotalAmountOut = TotalAmountOut.Add(AmountOut)
					remainingAmount = remainingAmount.Sub(AmountOut)				


				}
			}
			
		} else {
			return nil, sdkerrors.Wrapf(types.ErrValidPathNotFound, "Valid Path not found")
		}

	} else {
		fmt.Println("Do I go here?")
		if len(oldTick.PoolsOneToZero) != 0 {
			for (!(remainingAmount.Equal(sdk.ZeroDec())) || len(oldTick.PoolsOneToZero) ==0 ) {
				if (remainingAmount.LT(oldTick.PoolsZeroToOne[0].Reserve0)) {
					AmountOut := remainingAmount.Sub( remainingAmount.Mul(oldTick.PoolsOneToZero[0].Fee.Quo(oldTick.PoolsOneToZero[0].Price.Mul(sdk.NewDec(10000))))  )
					NewReserve0 := oldTick.PoolsOneToZero[0].Reserve0.Sub(AmountOut)
					
					
					k.dexKeeper.Update1to0(&oldTick.PoolsOneToZero, oldTick.PoolsOneToZero[0],  NewReserve0,
						oldTick.PoolsOneToZero[0].Reserve1.Add(remainingAmount), oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].TotalShares, oldTick.PoolsOneToZero[0].Price)
					
					oldZeroToOnePool, ZeroToOnePoolFound := k.dexKeeper.GetPool(&oldTick.PoolsZeroToOne, oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].Price )

					if ZeroToOnePoolFound {
						k.dexKeeper.Update0to1(&oldTick.PoolsOneToZero, &oldZeroToOnePool, NewReserve0,
							oldTick.PoolsOneToZero[0].Reserve1.Add(remainingAmount),  oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].TotalShares, oldTick.PoolsOneToZero[0].Price)
					
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

					NewTick := dextypes.Ticks {
						token0,
						token1,
						oldTick.PoolsZeroToOne,
						oldTick.PoolsOneToZero,
					}

					k.dexKeeper.SetTicks(
						ctx,
						NewTick,
					)

					TotalAmountOut = TotalAmountOut.Add(AmountOut)
					remainingAmount = sdk.ZeroDec()
					
				} else {
					
					AmountOut := oldTick.PoolsOneToZero[0].Reserve0.Sub( oldTick.PoolsOneToZero[0].Reserve0.Mul(oldTick.PoolsOneToZero[0].Fee.Quo(oldTick.PoolsOneToZero[0].Price.Mul(sdk.NewDec(10000))))  )
					

					

					oldZeroToOnePool, ZeroToOnePoolFound := k.dexKeeper.GetPool(&oldTick.PoolsZeroToOne, oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].Price )

					if ZeroToOnePoolFound {
						k.dexKeeper.Update1to0(&oldTick.PoolsOneToZero, &oldZeroToOnePool, sdk.ZeroDec(),
							oldTick.PoolsOneToZero[0].Reserve1.Add(oldTick.PoolsOneToZero[0].Reserve0),  oldTick.PoolsOneToZero[0].Fee, oldTick.PoolsOneToZero[0].TotalShares, oldTick.PoolsOneToZero[0].Price)
					
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

					NewTick := dextypes.Ticks {
						token0,
						token1,
						oldTick.PoolsZeroToOne,
						oldTick.PoolsOneToZero,
					}

					k.dexKeeper.Pop0to1(&oldTick.PoolsZeroToOne)

					k.dexKeeper.SetTicks(
						ctx,
						NewTick,
					)

					TotalAmountOut = TotalAmountOut.Add(AmountOut)
					remainingAmount = remainingAmount.Sub(AmountOut)				


				}
			}
		} else {
			return nil, sdkerrors.Wrapf(types.ErrValidPathNotFound, "Valid Path not found")
		}

	}

	fmt.Println(TotalAmountOut)
	minOut, err := sdk.NewDecFromStr(msg.MinOut)
	if err != nil {
		return nil, err
	}

	if TotalAmountOut.LT(minOut) {
		//TODO Custom ERROR Type here
		return nil, err
	}

	
	if amountIn.GT(sdk.ZeroDec()) {
		coinIn := sdk.NewCoin(msg.TokenIn, sdk.NewIntFromBigInt(amountIn.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, dextypes.ModuleName, sdk.Coins{coinIn}); err != nil {
			return nil, err
		}
	}

	if TotalAmountOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(msg.TokenOut, sdk.NewIntFromBigInt(TotalAmountOut.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, dextypes.ModuleName, callerAddr, sdk.Coins{coinOut}); err != nil {
			return nil, err
		}
	}
	// A B C D

	// A B
	// TokenIN: A
	// TokenOut: B
	//sort(tokenIN, tokenOut)
	//TokenIN = token0
	// TokenOut == token1

	// Mapping(token0, token1) -> PoolZeroToOne, PoolsOneToZero

	// TokenIN: B
	// TOkenOut: A
	//TokenOut == token0
	//TokenIN == token1
	// Mapping(token0, token1) -> PoolZeroToOne, PoolsOneToZero

	_ = ctx

	return &types.MsgSwapResponse{}, nil
}
