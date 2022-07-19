package keeper

import (
	//"fmt"

	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k Keeper) depositHelperSub(pool *types.Pool, amount0, amount1 sdk.Int, Fee, Price string ) (sdk.Dec, sdk.Dec, sdk.Int, error){

	fee, err := sdk.NewDecFromStr(Fee)
	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdk.NewInt(0), err
	}

	price, err := sdk.NewDecFromStr(Price)
	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdk.NewInt(0), err
	}

	if pool.Reserves0 > 0 {
		trueAmounts0 = k.min(uint(msg.Amounts1), uint((uint(tickOld.Reserves1)*uint(msg.Amounts0))/uint(tickOld.Reserves0)))
	}

	if tickOld.Reserves1 > 0 {
		trueAmounts0 = k.min(uint(msg.Amounts0), uint((uint(tickOld.Reserves0)*uint(msg.Amounts1))/uint(tickOld.Reserves1)))
	}

	if trueAmounts0 == uint(msg.Amounts0) && trueAmounts1 != uint(msg.Amounts1) {
		trueAmounts1 = uint(msg.Amounts1) + (((uint(msg.Amounts1) - trueAmounts1) * uint(msg.Fee)) / uint(10000-msg.Fee))
	} else if trueAmounts1 == uint(msg.Amounts1) && trueAmounts0 != uint(msg.Amounts0) {
		trueAmounts0 = uint(msg.Amounts0) + (((uint(msg.Amounts0) - trueAmounts0) * uint(msg.Fee)) / uint(10000-msg.Fee))
	}

	if tickOld.TotalShares == 0 {
		SharesMinted = uint(float64(msg.Amounts0) + float64(msg.Amounts1)*price)
	} else {
		SharesMinted =
			uint(float64(tickOld.TotalShares) * ((float64(msg.Amounts0) + float64(msg.Amounts1)*price) / (float64(tickOld.Reserves0) + float64(tickOld.Reserves1)*price)))
	}
}