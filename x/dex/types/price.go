package types

import (
	"math/big"

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
		return utils.BasePriceDec.Power(uint64(-1 * relativeTickIndex)), nil
	} else {
		return sdk.OneDec().Quo(utils.BasePriceDec.Power(uint64(relativeTickIndex))), nil
	}
}

func MustCalcPrice(relativeTickIndex int64) sdk.Dec {
	price, err := CalcPrice(relativeTickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

// Power returns a the result of raising to a positive integer power
func RatPower(base *big.Rat, power uint64) *big.Rat {
	if power == 0 {
		return big.NewRat(1, 1)
	}
	tmp := big.NewRat(1, 1)

	for i := power; i > 1; {
		if i%2 != 0 {
			tmp = tmp.Mul(tmp, base)
		}
		i /= 2
		base = base.Mul(base, base)
	}

	return base.Mul(base, tmp)
}

func CalcPriceAsRat(relativeTickIndex int64) (*big.Rat, error) {
	if IsTickOutOfRange(relativeTickIndex) {
		return big.NewRat(0, 1), ErrTickOutsideRange
	}
	if relativeTickIndex < 0 {
		inverseResult := RatPower(utils.BasePriceRat, uint64(-1*relativeTickIndex))
		return inverseResult.Inv(inverseResult), nil
	} else {
		result := RatPower(utils.BasePriceRat, uint64(relativeTickIndex))
		return result, nil
	}
}

func MustCalcPriceAsRat(relativeTickIndex int64) *big.Rat {
	price, err := CalcPriceAsRat(relativeTickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

func CalcPrice0To1(tickIndex int64) (sdk.Dec, error) {
	return CalcPrice(tickIndex)
}

func MustCalcPrice0To1(tickIndex int64) sdk.Dec {
	return MustCalcPrice(tickIndex)
}

func CalcPrice1To0(tickIndex int64) (sdk.Dec, error) {
	return CalcPrice(-1 * tickIndex)
}

func MustCalcPrice1To0(tickIndex int64) sdk.Dec {
	return MustCalcPrice(-1 * tickIndex)
}

func IsTickOutOfRange(tickIndex int64) bool {
	absTickIndex := utils.Abs(tickIndex)
	return absTickIndex > MaxTickExp
}
