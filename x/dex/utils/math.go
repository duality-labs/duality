package utils

import (
	"fmt"
	"math"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Return the base value for price, 1.0001
func BasePrice() sdk.Dec {
	return sdk.MustNewDecFromStr("1.0001")
}

func Abs(x int64) uint64 {
	if x < 0 {
		return uint64(-x)
	}
	return uint64(x)
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

func Uint64ToSortableString(i uint64) string {
	// Converts a Uint to a string that sorts lexogrpahically in integer order
	intStr := strconv.FormatUint(i, 36)
	lenStr := len(intStr)
	lenChar := strconv.FormatUint(uint64(lenStr), 36)
	return fmt.Sprintf("%s%s", lenChar, intStr)
}

func SafeUint64(in uint64) (out int64, overflow bool) {
	return int64(in), in > math.MaxInt64
}

func MustSafeUint64(in uint64) (out int64) {
	int64, overflow := SafeUint64(in)
	if overflow {
		panic("Overflow while casting uint64 to int64")
	}
	return int64
}
