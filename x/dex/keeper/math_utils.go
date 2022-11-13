package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// Return the base value for price, 1.0001
func BasePrice() sdk.Dec {
	prec := sdk.NewDecFromIntWithPrec(sdk.NewIntFromUint64(1), 14)
	return sdk.OneDec().Add(prec)
}

// Iterative implementation of exponentiation by squaring algorithm, taken from Wikipedia (https://en.wikipedia.org/wiki/Exponentiation_by_squaring#With_constant_auxiliary_memory)
// Note: sdk.Dec will panic on overflow
func Pow(x sdk.Dec, n int64) sdk.Dec {
	if n == 0 {
		return sdk.OneDec()
	}

	// convert n to nonnegative exponent
	var exp uint64
	if n < 0 {
		exp = uint64(-1 * n)
		x = sdk.OneDec().QuoInt(sdk.NewIntFromUint64(exp))
	} else {
		exp = uint64(n)
	}
	y := sdk.OneDec()
	for exp > 1 {
		if exp%2 == 0 {
			x = x.Mul(x)
			exp /= 2
		} else {
			y = x.Mul(y)
			x = x.Mul(x)
			exp = (exp - 1) / 2
		}
	}

	return x.Mul(y)
}
