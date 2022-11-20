package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxTickExp uint64 = 1048575

func BasePrice() sdk.Dec {
	return sdk.MustNewDecFromStr("1.0001")
}

func Pow(a sdk.Dec, n uint64) (sdk.Dec, error) {
	if n > MaxTickExp {
		return sdk.ZeroDec(), types.ErrTickAbsValTooHigh
	}
	sum := sdk.OneDec()
	for n > 0 {
		if n&1 == 1 {
			sum = sum.Mul(a)
		}
		a = a.Mul(a)
		n >>= 1
	}
	return sum, nil
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
