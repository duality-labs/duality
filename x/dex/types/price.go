package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

// NOTE: -352,437 is the lowest possible tick at which price can be calculated with a < 1% error
// when using 18 digit decimal precision (via sdk.Dec)
const MaxTickExp uint64 = 352437

// Calculates the price for a swap from token 0 to token 1 given a relative tick
// tickIndex refers to the index of a specified tick such that x * 1.0001 ^(-1 * t) = y
// Lower ticks offer better prices.
func CalcPrice(relativeTickIndex int64) (sdk.Dec, error) {
	if IsTickOutOfRange(relativeTickIndex) {
		return sdk.ZeroDec(), ErrTickOutsideRange
	}
	if relativeTickIndex < 0 {
		return utils.BasePrice().Power(uint64(-1 * relativeTickIndex)), nil
	} else {
		return sdk.OneDec().Quo(utils.BasePrice().Power(uint64(relativeTickIndex))), nil
	}
}

func MustCalcPrice(relativeTickIndex int64) sdk.Dec {
	price, err := CalcPrice(relativeTickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

func IsTickOutOfRange(tickIndex int64) bool {
	absTickIndex := utils.Abs(tickIndex)
	return absTickIndex > MaxTickExp
}
