package keeper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

const LPsharesRegexpStr = "^" + types.DepositSharesPrefix + "-" +
	// Token0 (regexp from cosmos-sdk.types.coin.reDnmString)
	"([a-zA-Z][a-zA-Z0-9/-]{2,127})" + "-" +
	// Token1
	"([a-zA-Z][a-zA-Z0-9/-]{2,127})" + "-" +
	// Tickindex
	`t(\d+)` + "-" +
	// fee
	`f(\d+)`

var LPSharesRegexp = regexp.MustCompile(LPsharesRegexpStr)

func CreateSharesId(token0 string, token1 string, tickIndex int64, fee uint64) (denom string) {
	t0 := strings.ReplaceAll(token0, "-", "")
	t1 := strings.ReplaceAll(token1, "-", "")
	return fmt.Sprintf("%s-%s-%s-t%d-f%d", types.DepositSharesPrefix, t0, t1, tickIndex, fee)
}

func ParseDepositShares(shares sdk.Coin) (matches []string, valid bool) {
	// NOTE: Since dashes are removed as part of CreateSharesId, if either side of the LP position are denoms that contain dashes
	// they will not be parsed correctly and the correct dneom will not be returned
	matchArr := LPSharesRegexp.FindAllStringSubmatch(shares.Denom, -1)
	if matchArr == nil {
		return nil, false
	}
	return matchArr[0][1:5], true
}

func DepositSharesToData(shares sdk.Coin) (types.DepositRecord, error) {
	matches, valid := ParseDepositShares(shares)

	if !valid {
		return types.DepositRecord{}, types.ErrInvalidDepositShares
	}

	pairId := CreatePairId(matches[0], matches[1])
	tickIndex, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		return types.DepositRecord{}, types.ErrInvalidDepositShares
	}
	fee, err := strconv.ParseUint(matches[3], 10, 0)
	if err != nil {
		return types.DepositRecord{}, types.ErrInvalidDepositShares
	}
	feeUint := utils.MustSafeUint64(fee)
	return types.DepositRecord{
		PairId:          pairId,
		SharesOwned:     shares.Amount,
		CenterTickIndex: tickIndex,
		LowerTickIndex:  tickIndex - feeUint,
		UpperTickIndex:  tickIndex + feeUint,
		Fee:             fee,
	}, nil
}
