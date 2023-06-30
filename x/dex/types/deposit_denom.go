package types

import (
	fmt "fmt"
	"regexp"
	"strconv"
	"strings"
)

const LPsharesRegexpStr = "^" + DepositSharesPrefix + "-" +
	// Token0 (regexp from cosmos-sdk.types.coin.reDnmString)
	"([a-zA-Z][a-zA-Z0-9/-]{2,127})" + "-" +
	// Token1
	"([a-zA-Z][a-zA-Z0-9/-]{2,127})" + "-" +
	// Tickindex
	`t(-?\d+)` + "-" +
	// fee
	`f(\d+)`

var LPSharesRegexp = regexp.MustCompile(LPsharesRegexpStr)

type DepositDenom struct {
	PairID *PairID
	Tick   int64
	Fee    uint64
}

func NewDepositDenom(pairID *PairID, tick int64, fee uint64) *DepositDenom {
	return &DepositDenom{
		PairID: pairID,
		Tick:   tick,
		Fee:    fee,
	}
}

func NewDepositDenomFromString(denom string) (depositDenom *DepositDenom, err error) {
	// NOTE: Since dashes are removed as part of CreateSharesId, if either side of the LP position are denoms that contain dashes
	// they will not be parsed correctly and the correct dneom will not be returned
	matchArr := LPSharesRegexp.FindAllStringSubmatch(denom, -1)
	if matchArr == nil {
		return nil, ErrInvalidDepositDenom
	}

	matches := matchArr[0][1:5]
	tick, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		return nil, ErrInvalidDepositDenom
	}

	fee, err := strconv.ParseUint(matches[3], 10, 0)
	if err != nil {
		return nil, ErrInvalidDepositDenom
	}

	return &DepositDenom{
		PairID: &PairID{
			Token0: matches[0],
			Token1: matches[1],
		},
		Tick: tick,
		Fee:  fee,
	}, nil
}

func (d DepositDenom) String() string {
	// TODO: Revist security of this.
	prefix := DepositDenomPairIDPrefix(d.PairID.Token0, d.PairID.Token1)
	return fmt.Sprintf("%s-t%d-f%d", prefix, d.Tick, d.Fee)
}

func DepositDenomPairIDPrefix(token0, token1 string) string {
	t0 := strings.ReplaceAll(token0, "-", "")
	t1 := strings.ReplaceAll(token1, "-", "")
	return fmt.Sprintf("%s-%s-%s", DepositSharesPrefix, t0, t1)
}
