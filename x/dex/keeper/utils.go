package keeper

import (
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// tokenDirection
func (k Keeper) CalculateVirtualPrice(token0 string, token1 string, tokenOut string, fee sdk.Dec, price sdk.Dec) (sdk.Dec, error) {

	if token0 == tokenOut {

		return fee.Quo(price.Mul(sdk.NewDec(10000))), nil
	} else if token1 == tokenOut {
		// pools[j].Price.Mul(pools[j].Fee)).Quo(sdk.NewDec(10000)))
		return price.Mul(fee).Quo(sdk.NewDec(10000)), nil
	}
	return sdk.ZeroDec(), nil

}

// Returns corresponding tick index for a given virtualPrice, always rounds down
func (k Keeper) CalculateTick(virtualPrice sdk.Dec) (uint64, error) {

	// Ticks are 1bp apart
	fVirtualPrice, _ := virtualPrice.Float64()
	index := math.Log(fVirtualPrice) / math.Log(1.0001)
	// Always takes floor of index
	return uint64(index), nil

}

// Always gives price as y/x
// TODO: Confirm scaling is correct
func (k Keeper) GetVirtualPriceFromTick(tick int32) (sdk.Dec, error) {

	vPrice := math.Pow(1.0001, float64(tick))
	// Always takes floor of index
	// TODO: Need to do testing on conversions
	return sdk.NewDec(int64(vPrice)), nil

}

// TODO: Set up a tick-map or a quick way to query the next index?
// Might make sense to sort this, but we'll see especially depending on spacing near the center of the tick
func (k Keeper) GetNextIndex(ctx sdk.Context, Pair types.Pairs, swapToToken0 bool) (index int32, IndexQueueFound bool) {
	currIdx := Pair.CurrentIndex
	currIdx += 1
	for math.MinInt32 < currIdx && currIdx < math.MaxInt32 {
		_, TickFound := k.GetIndexQueue(ctx, Pair.Token0, Pair.Token1, currIdx)
		if TickFound {
			return currIdx, true
		} else {
			if swapToToken0 {
				currIdx += 1
			} else {
				currIdx -= 1
			}
		}
	}
	return 0, false
}
