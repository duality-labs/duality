package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p PoolReserves) HasToken() bool {
	return p.ReservesMakerDenom.GT(sdk.ZeroInt())
}

func NewPoolReserves(makerDenom string, takerDenom string, tickIndex int64, fee uint64, reservesMakerDenom sdk.Int) (*PoolReserves, error) {
	tradePairID, err := NewTradePairID(takerDenom, makerDenom)
	if err != nil {
		return nil, err
	}
	priceMakerToTaker, err := tradePairID.PriceMakerToTaker(tickIndex)
	if err != nil {
		return nil, err
	}
	priceTakerToMaker, err := tradePairID.PriceTakerToMaker(tickIndex)
	if err != nil {
		return nil, err
	}
	return &PoolReserves{
		TradePairID:        tradePairID,
		TickIndex:          tickIndex,
		Fee:                fee,
		ReservesMakerDenom: reservesMakerDenom,
		PriceTakerToMaker:  priceTakerToMaker,
		PriceMakerToTaker:  priceMakerToTaker,
	}, nil
}

// Useful for testing
func MustNewPoolReserves(makerDenom string, takerDenom string, tickIndex int64, fee uint64, reservesMakerDenom sdk.Int) *PoolReserves {
	poolReserves, err := NewPoolReserves(makerDenom, takerDenom, tickIndex, fee, reservesMakerDenom)
	if err != nil {
		panic(err)
	}
	return poolReserves
}
