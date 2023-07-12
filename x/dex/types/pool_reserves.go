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
		Key:                       poolReservesID,
		ReservesMakerDenom:        sdk.ZeroInt(),
		PriceTakerToMaker:         priceTakerToMaker,
		PriceOppositeTakerToMaker: priceOppositeTakerToMaker,
	}, nil
}

func MustNewPoolReserves(
	poolReservesID *PoolReservesKey,
) *PoolReserves {
	poolReserves, err := NewPoolReserves(poolReservesID)
	if err != nil {
		panic(err)
	}
	return poolReserves
}
