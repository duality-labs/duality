package types

import (
	"fmt"
	"regexp"
)

const PoolDenomRegexpStr = PoolNamePrefix + `\d+`

var PoolDenomRegexp = regexp.MustCompile(PoolDenomRegexpStr)

func NewPoolDenom(poolIdx uint64) string {
	return fmt.Sprintf("%s%d", PoolNamePrefix, poolIdx)
}

func ValidatePoolDenom(denom string) error {
	if !PoolDenomRegexp.MatchString(denom) {
		return ErrInvalidDepositDenom
	}

	return nil
}
