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

func MinIntArr(vals []sdk.Int) sdk.Int {
	min := vals[0]
	for _, val := range vals {
		if val.LT(min) {
			min = val
		}
	}

	return min
}

func MaxIntArr(vals []sdk.Int) sdk.Int {
	max := vals[0]
	for _, val := range vals {
		if val.GT(max) {
			max = val
		}
	}

	return max
}

func Uint64ToSortableString(i uint64) string {
	// Converts a Uint to a string that sorts lexogrpahically in integer order
	intStr := strconv.FormatUint(i, 36)
	lenStr := len(intStr)
	lenChar := strconv.FormatUint(uint64(lenStr), 36)

	return fmt.Sprintf("%s%s", lenChar, intStr)
}

func SafeUint64ToInt64(in uint64) (out int64, overflow bool) {
	return int64(in), in > math.MaxInt64
}

func MustSafeUint64ToInt64(in uint64) (out int64) {
	safeInt64, overflow := SafeUint64ToInt64(in)
	if overflow {
		panic("Overflow while casting uint64 to int64")
	}

	return safeInt64
}
