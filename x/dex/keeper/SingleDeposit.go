package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CalculateVirtualPrice(token0 string, token1 string, tokenDirection string, amout sdk.Dec, fee sdk.Dec, price sdk.Dec) (sdk.Dec, error) {

	if token0 == tokenDirection {

		return fee.Quo(price.Mul(sdk.NewDec(10000))), nil
	} else if token1 == tokenDirection {
		// pools[j].Price.Mul(pools[j].Fee)).Quo(sdk.NewDec(10000)))
		return price.Mul(fee).Quo(sdk.NewDec(10000)), nil
	}
	return sdk.ZeroDec(), nil

}

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

	vprice, err := k.CalculateVirtualPrice(token0, token1, msg.TokenDirection, amount, fee, price)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Virtual Price Calculations resulted in a non-valid type: %s", err)
	}

	// Can only deposit amount0 where vPrice >= CurrentPrice
	if vprice.LT(PairOld.CurrentPrice) && msg.TokenDirection == token0 {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit token0 at a price/fee pair less than the current price")
		// Can only deposit amount1 where CurrentPrice >= vPrice
	} else if PairOld.CurrentPrice.LT(vprice) && msg.TokenDirection == token1 {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot deposit token1 at a price/fee pair greater than the current price")
	}

	VirtualPriceQueueOld, VirtualPriceQueueFound := k.GetVirtualPriceQueue(ctx, vprice.String(), msg.TokenDirection, msg.OrderType)

	shares := amount.Mul(vprice)

	if !VirtualPriceQueueFound {

		NewQueue := []*types.VirtualPriceQueueType{
			&types.VirtualPriceQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: shares,
				},
			},
		}
		VirtualPriceQueueOld = types.VirtualPriceQueue{
			VPrice:    vprice.String(),
			Direction: msg.TokenDirection,
			OrderType: msg.OrderType,
			Queue:     NewQueue,
		}

	} else {
		TickIndexFound := -1
		for i := 0; i < len(VirtualPriceQueueOld.Queue); i++ {
			if VirtualPriceQueueOld.Queue[i].Price.Equal(price) && VirtualPriceQueueOld.Queue[i].Fee.Equal(fee) && VirtualPriceQueueOld.Queue[i].Orderparams.OrderType == msg.OrderType {
				TickIndexFound = i
				break
			}
		}

		if TickIndexFound != -1 {

			VirtualPriceQueueOld.Queue[TickIndexFound] = &types.VirtualPriceQueueType{
				Price: price,
				Fee:   fee,
				Orderparams: &types.OrderParams{
					OrderRule:   "",
					OrderType:   msg.OrderType,
					OrderShares: VirtualPriceQueueOld.Queue[TickIndexFound].Orderparams.OrderShares.Add(shares),
				},
			}
		} else {

			VirtualPriceQueueOld.Queue = k.enqueue(ctx, VirtualPriceQueueOld.Queue, types.VirtualPriceQueueType{
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
		Token0:          token0,
		Token1:          token1,
		TickSpacing:     PairOld.TickSpacing,
		CurrentPrice:    PairOld.CurrentPrice,
		BitArray:        PairOld.BitArray,
		Tickmap:         PairOld.Tickmap,
		VirtualPriceMap: &VirtualPriceQueueOld,
	},
	)

	return nil
}
