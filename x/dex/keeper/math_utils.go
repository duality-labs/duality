package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Return the base value for price, 1.0001
func BasePrice() sdk.Dec {
	return sdk.MustNewDecFromStr("1.0001")
}

func Pow(a sdk.Dec, n uint64) sdk.Dec {
	if n == 0 {
		return sdk.OneDec()
	}
	if n&1 == 0 {
		return Pow(a.Mul(a), n>>1)
	} else {
		return a.Mul(Pow(a.Mul(a), n>>1))
	}
}

func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	} else {
		return a
	}
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MinDec(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}

func MaxDec(a, b sdk.Dec) sdk.Dec {
	if a.GT(b) {
		return a
	}
	return b
}
