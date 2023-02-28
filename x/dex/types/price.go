package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

// NOTE: -352,437 is the lowest possible tick at which price can be calculated with a < 1% error
// when using 18 digit decimal precision (via sdk.Dec)
const MaxTickExp uint64 = 352437

type Price struct {
	// Relative in this context means that if `price` represents a conversion factor
	// such that `x * price = y`, then `x * 1.0001^(-1 * RelativeTickIndex) = y`.
	// So for a pair `"tokenA<>tokenB"`, if you wanted to trade tokenB for tokenA
	// then the conversion factor would be `x * 1.0001^(-1 * -1 * StandardizedTickIndex) = y`
	// (the extra -1 factor is because we're trading from the right-hand token to the left-hand token).
	RelativeTickIndex int64
}

func MustNewPrice(relativeTickIndex int64) *Price {
	price, err := NewPrice(relativeTickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

func NewPrice(relativeTickIndex int64) (*Price, error) {
	if IsTickOutOfRange(relativeTickIndex) {
		// TODO: This is a bit weird that we can't return a types.ErrTickOutsideRange because of cyclical dependencesi
		//Also maybe don't need this at all since we already validate that the tick is in range upstream
		return nil, errors.New("Supplying a tick outside the range of [-352437, 352437] is not allowed")
	}
	return &Price{
		RelativeTickIndex: relativeTickIndex,
	}, nil
}

func (p Price) MulInt(other sdk.Int) sdk.Dec {
	return p.Mul(other.ToDec())
}

// We are careful not to use negative-valued exponents anywhere
func (p Price) Mul(other sdk.Dec) sdk.Dec {
	if p.RelativeTickIndex >= 0 {
		return other.Quo(utils.BasePrice().Power(uint64(p.RelativeTickIndex)))
	} else {
		return other.Mul(utils.BasePrice().Power(uint64(-1 * p.RelativeTickIndex)))
	}
}

func (p Price) Inv() *Price {
	return &Price{
		RelativeTickIndex: p.RelativeTickIndex * -1,
	}
}

// We don't use this for calculations because when the tick is positive
// then we calculate via a manual inversion 1 / 1.0001^X and this is lossy
func (p Price) ToDec() sdk.Dec {
	if p.RelativeTickIndex >= 0 {
		return sdk.OneDec().Quo(utils.BasePrice().Power(uint64(p.RelativeTickIndex)))
	} else {
		return utils.BasePrice().Power(uint64(-1 * p.RelativeTickIndex))
	}
}

func MustCalcPrice0To1(tickIndex int64) *Price {
	price, err := CalcPrice0To1(tickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

// Calculates the price for a swap from token 0 to token 1 given a tick
// tickIndex refers to the index of a specified tick
func CalcPrice0To1(tickIndex int64) (*Price, error) {
	return NewPrice(tickIndex)
}

func MustCalcPrice1To0(tickIndex int64) *Price {
	return MustNewPrice(-1 * tickIndex)
}

// Calculates the price for a swap from token 1 to token 0 given a tick
// tickIndex refers to the index of a specified tick
func CalcPrice1To0(tickIndex int64) (*Price, error) {
	return NewPrice(-1 * tickIndex)
}

func IsTickOutOfRange(tickIndex int64) bool {
	absTickIndex := utils.Abs(tickIndex)
	return absTickIndex > MaxTickExp
}
