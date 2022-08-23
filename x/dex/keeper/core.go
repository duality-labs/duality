package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddLiquidityVerification(goCtx context.Context, msg *types.MsgAddLiquidity) (string, string, sdk.AccAddress, sdk.AccAddress, sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	price, err := sdk.NewDecFromStr(msg.Price)
	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	token0, token1, priceDec, err := k.SortTokens(ctx, msg.TokenA, msg.TokenB, price)

	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Converts receiver address (string) to sdk.AccAddress
	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error checking for the valid receiver address
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	amount, err := sdk.NewDecFromStr(msg.Amount)
	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	if msg.TokenDirection != msg.TokenA && msg.TokenB != msg.TokenDirection {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Token Direction must be the same as either Token A or Token B")
	}
	//var decAmounts []sdk.Dec
	// for i := 0; i < len(msg.Amount); i++ {

	//amount, err := sdk.NewDecFromStr(msg.Amount)
	// // Error checking for valid sdk.Dec
	//if err != nil {
	// return "", "", nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	// }
	// decAmounts = append(decAmounts, amount)
	// }

	return token0, token1, callerAddr, receiverAddr, amount, priceDec, nil
}

func (k Keeper) SingleDeposit(goCtx context.Context, token0 string, token1 string, amount sdk.Dec, price sdk.Dec, msg *types.MsgAddLiquidity, callerAddr sdk.AccAddress, receiver sdk.AccAddress) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	PairOld, PairFound := k.GetPairs(ctx, token0, token1)

	if !PairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Valid pair not found")
	}

	// Fee is some value in {0, 10000} where fee/10000 is equal the exchange rate
	fee, err := sdk.NewDecFromStr(msg.Fee)
	// Error checking for valid sdk.Dec
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	fee = fee.Quo(sdk.NewDec(10000))

	// Can only deposit amount0 where vPrice >= CurrentPrice
	if msg.Index < (PairOld.CurrentIndex) && msg.TokenDirection == token0 {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit token0 at a price/fee pair less than the current price")
		// Can only deposit amount1 where CurrentPrice >= vPrice
	} else if PairOld.CurrentIndex < msg.Index && msg.TokenDirection == token1 {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit token1 at a price/fee pair greater than the current price")
	}

	IndexQueue, IndexQueueFound := k.GetIndexQueue(ctx, token0, token1, msg.Index)

	Tick, TickFound := k.GetTicks(ctx, token0, token1, msg.Price, msg.Fee, msg.OrderType)

	var NewTick types.Ticks
	var oldAmount sdk.Dec //Event variable
	var shares sdk.Dec

	if msg.TokenDirection == token0 {
		shares = amount.Mul(price.Mul(fee))
	} else {
		shares = amount.Mul(sdk.OneDec().Quo(fee))
	}
	// Index QUeue Logic

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

			IndexQueue.Queue = k.enqueue(ctx, IndexQueue.Queue, types.IndexQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: shares,
				},
			})

		} //else {

		// 	IndexQueue.Queue = k.enqueue(ctx, IndexQueue.Queue, types.IndexQueueType{
		// 		Price: price,
		// 		Fee:   fee,
		// 		Orderparams: &types.OrderParams{
		// 			OrderRule:   "",
		// 			OrderType:   msg.OrderType,
		// 			OrderShares: Tick.TotalShares.Add(shares),
		// 		},
		// 	})
		// }
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
		if msg.TokenDirection == token0 {
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

			oldAmount = Tick.Reserve0
		} else {
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
			oldAmount = Tick.Reserve1
		}

	}

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

	k.SetTicks(ctx, token0, token1, NewTick)
	k.SetIndexQueue(ctx, token0, token1, IndexQueue)

	PairNew, PairFound := k.GetPairs(ctx, token0, token1)

	NewPairs := types.Pairs{
		Token0:       token0,
		Token1:       token1,
		CurrentIndex: PairOld.CurrentIndex,
		TickSpacing:  PairOld.TickSpacing,
		Tickmap:      PairNew.Tickmap,
		IndexMap:     PairNew.IndexMap,
	}

	k.SetPairs(ctx, NewPairs)

	ctx.EventManager().EmitEvent(types.CreateDepositEvent(msg.Creator,
		token0, token1, price.String(), fee.String(), msg.TokenDirection,
		oldAmount.String(), oldAmount.Add(amount).String(),
		sdk.NewAttribute(types.DepositEventSharesMinted, shares.String()),
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
