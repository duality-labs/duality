package keeper

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// func (k Keeper) CalculateVirtualPrice(token0 string, token1 string, tokenDirection string, amount sdk.Dec, fee sdk.Dec, price sdk.Dec) (sdk.Dec, error) {

// 	if token0 == tokenDirection {

// 		return fee.Quo(price.Mul(sdk.NewDec(10000))), nil
// 	} else if token1 == tokenDirection {
// 		// pools[j].Price.Mul(pools[j].Fee)).Quo(sdk.NewDec(10000)))
// 		return price.Mul(fee).Quo(sdk.NewDec(10000)), nil
// 	}
// 	return sdk.ZeroDec(), nil

// }

// Returns corresponding tick index for a given virtualPrice, always rounds down
func (k Keeper) CalculateTick(virtualPrice sdk.Dec) (uint64, error) {

	// Ticks are 1bp apart
	fVirtualPrice, _ := virtualPrice.Float64()
	index := math.Log(fVirtualPrice) / math.Log(1.0001)
	// Always takes floor of index
	return uint64(index), nil

}

func (k Keeper) GetVirtualPriceFromTick(tick uint64) (sdk.Dec, error) {

	vPrice := math.Pow(1.0001, float64(tick))
	// Always takes floor of index
	// TODO: Need to do testing on conversions
	return sdk.NewDec(int64(vPrice)), nil

}
