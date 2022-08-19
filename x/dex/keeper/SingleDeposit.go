package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SingleDeposit(goCtx context.Context, token0 string, token1 string, amount sdk.Dec, price sdk.Dec, msg *types.MsgAddLiquidity, callerAdr sdk.AccAddress, receiver sdk.AccAddress) error {

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
	shares := amount.Mul(price.Mul(fee))

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
