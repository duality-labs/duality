package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

// NOTE: -352,437 is the lowest possible tick at which price can be calculated with a < 1% error
// when using 18 digit decimal precision (via sdk.Dec)
const MaxTickExp uint64 = 352437

func MustCalcPrice0To1(tickIndex int64) sdk.Dec {
	price, err := CalcPrice0To1(tickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

// Calculates the price for a swap from token 0 to token 1 given a tick
// tickIndex refers to the index of a specified tick
func CalcPrice0To1(tickIndex int64) (sdk.Dec, error) {
	if IsTickOutOfRange(tickIndex) {
		// TODO: This is a bit weird that we can't return a types.ErrTickOutsideRange because of cyclical dependencesi
		//Also maybe don't need this at all since we already validate that the tick is in range upstream
		return sdk.ZeroDec(), errors.New("Supplying a tick outside the range of [-352437, 352437] is not allowed")
	}

	if 0 <= tickIndex {
		return sdk.OneDec().Quo(BasePrice().Power(uint64(tickIndex))), nil
	} else {
		return BasePrice().Power(uint64(-1 * tickIndex)), nil
	}
}

func MustCalcPrice1To0(tickIndex int64) sdk.Dec {
	price, err := CalcPrice1To0(tickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

// Calculates the price for a swap from token 1 to token 0 given a tick
// tickIndex refers to the index of a specified tick
func CalcPrice1To0(tickIndex int64) (sdk.Dec, error) {

	// TODO: see above
	if IsTickOutOfRange(tickIndex) {
		return sdk.ZeroDec(), errors.New("Supplying a tick outside the range of [-352437, 352437] is not allowed")
	}
	if 0 <= tickIndex {
		return BasePrice().Power(uint64(tickIndex)), nil
	} else {
		return sdk.OneDec().Quo(BasePrice().Power(uint64(-1 * tickIndex))), nil
	}
}

func IsTickOutOfRange(tickIndex int64) bool {
	absTickIndex := Abs(tickIndex)
	return absTickIndex > MaxTickExp
}
