package types

import (
	fmt "fmt"
	"regexp"
	"strconv"
)

const LPsharesRegexpStr = "^" + DepositSharesPrefix + "([0-9]{2,127})$"

var LPSharesRegexp = regexp.MustCompile(LPsharesRegexpStr)

func ParsePoolIDFromDepositDenom(denom string) (poolID uint64, err error) {
	matchArr := LPSharesRegexp.FindAllStringSubmatch(denom, -1)
	if matchArr == nil {
		return 0, ErrInvalidDepositDenom
	}

	poolID, err = strconv.ParseUint(matchArr[0][1], 10, 0)
	if err != nil {
		return 0, err
	}

	return poolID, nil
}

func NewDepositDenom(poolID uint64) string {
	return fmt.Sprintf("%s%d", DepositSharesPrefix, poolID)
}
