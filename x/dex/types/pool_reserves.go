package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p PoolReserves) HasToken() bool {
	return p.ReservesMakerDenom.GT(sdk.ZeroInt())
}

func NewPoolReservesFromCounterpart(
	counterpart *PoolReserves,
) *PoolReserves {
	thisID := counterpart.Key.Counterpart()
	return &PoolReserves{
		PoolId:                    counterpart.PoolId,
		Key:                       thisID,
		ReservesMakerDenom:        sdk.ZeroInt(),
		PriceTakerToMaker:         counterpart.PriceOppositeTakerToMaker,
		PriceOppositeTakerToMaker: counterpart.PriceTakerToMaker,
	}
}

func NewPoolReserves(
	poolReservesID *PoolReservesKey,
) (*PoolReserves, error) {
	priceTakerToMaker, err := poolReservesID.PriceTakerToMaker()
	if err != nil {
		return nil, err
	}
	counterpartID := poolReservesID.Counterpart()
	priceOppositeTakerToMaker, err := counterpartID.PriceTakerToMaker()
	if err != nil {
		return nil, err
	}

	return &PoolReserves{
		// PoolId is set on pool save
		Key:                       poolReservesID,
		ReservesMakerDenom:        sdk.ZeroInt(),
		PriceTakerToMaker:         priceTakerToMaker,
		PriceOppositeTakerToMaker: priceOppositeTakerToMaker,
	}, nil
}

func MustNewPoolReserves(
	poolID uint64,
	poolReservesID *PoolReservesKey,
) *PoolReserves {
	poolReserves, err := NewPoolReserves(poolReservesID)
	poolReserves.PoolId = poolID
	if err != nil {
		panic(err)
	}
	return poolReserves
}
