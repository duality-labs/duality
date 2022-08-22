package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SingleDeposit(goCtx context.Context, token0 string, token1 string, amount sdk.Dec, price sdk.Dec, msg *types.MsgAddLiquidity, callerAddr sdk.AccAddress, receiver sdk.AccAddress) error {

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

		} else {

			IndexQueue.Queue = k.enqueue(ctx, IndexQueue.Queue, types.IndexQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: Tick.TotalShares.Add(shares),
				},
			})
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
