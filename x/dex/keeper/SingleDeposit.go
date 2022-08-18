package keeper

import (
	"context"

	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CalculateVirtualPrice(token0 string, token1 string, tokenDirection string, amount sdk.Dec, fee sdk.Dec, price sdk.Dec) (sdk.Dec, error) {

	if token0 == tokenDirection {

		return fee.Quo(price.Mul(sdk.NewDec(10000))), nil
	} else if token1 == tokenDirection {
		// pools[j].Price.Mul(pools[j].Fee)).Quo(sdk.NewDec(10000)))
		return price.Mul(fee).Quo(sdk.NewDec(10000)), nil
	}
	return sdk.ZeroDec(), nil

}

// Returns corresponding tick index for a given virtualPrice
func (k Keeper) CalculateTick(virtualPrice sdk.Dec) (sdk.Dec, error) {

	// Ticks are 1bp apart
	fVirtualPrice, _ := virtualPrice.Float64()
	index := math.Log(fVirtualPrice) / math.Log(1.0001)
	// Always takes floor of index
	return sdk.NewDec(int64(index)), nil

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
	_ = vprice
	_ = PairOld

	return nil
}
