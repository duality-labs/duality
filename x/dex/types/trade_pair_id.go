package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewTradePairID(takerDenom, makerDenom string) (*TradePairID, error) {
	if takerDenom == makerDenom {
		return nil, sdkerrors.Wrapf(ErrInvalidTradingPair, "%s, %s", takerDenom, makerDenom)
	}
	return &TradePairID{
		TakerDenom: takerDenom,
		MakerDenom: makerDenom,
	}, nil
}

func MustNewTradePairID(takerDenom, makerDenom string) *TradePairID {
	tradePairID, err := NewTradePairID(takerDenom, makerDenom)
	if err != nil {
		panic(err)
	}
	return tradePairID
}

func NewTradePairIDFromMaker(pairID *PairID, makerDenom string) *TradePairID {
	var takerDenom string
	if pairID.Token0 == makerDenom {
		takerDenom = pairID.Token1
	} else {
		takerDenom = pairID.Token0
	}
	return &TradePairID{
		TakerDenom: takerDenom,
		MakerDenom: makerDenom,
	}
}

func NewTradePairIDFromTaker(pairID *PairID, takerDenom string) *TradePairID {
	var makerDenom string
	if pairID.Token0 == takerDenom {
		makerDenom = pairID.Token1
	} else {
		makerDenom = pairID.Token0
	}
	return &TradePairID{
		TakerDenom: takerDenom,
		MakerDenom: makerDenom,
	}
}

func (p TradePairID) IsTakerDenomToken0() bool {
	return p.TakerDenom == p.MustPairID().Token0
}

func (p TradePairID) IsMakerDenomToken0() bool {
	return p.MakerDenom == p.MustPairID().Token0
}

func (p TradePairID) MustPairID() *PairID {
	pairID, err := p.PairID()
	if err != nil {
		panic(err)
	}

	return pairID
}

func (p TradePairID) PairID() (*PairID, error) {
	return NewPairIDFromUnsorted(p.MakerDenom, p.TakerDenom)
}

func (p TradePairID) Reversed() *TradePairID {
	return &TradePairID{
		MakerDenom: p.TakerDenom,
		TakerDenom: p.MakerDenom,
	}
}

func (p TradePairID) TickIndexTakerToMaker(tickIndexNormalized int64) int64 {
	pairID := p.MustPairID()
	if pairID.Token1 == p.MakerDenom {
		return tickIndexNormalized
	} else {
		return -1 * tickIndexNormalized
	}
}

func (p TradePairID) TickIndexNormalized(tickIndexTakerToMaker int64) int64 {
	return p.TickIndexTakerToMaker(tickIndexTakerToMaker)
}

func (p TradePairID) PriceTakerToMaker(tickIndexNormalized int64) (priceTakerToMaker sdk.Dec, err error) {
	return CalcPrice(p.TickIndexTakerToMaker(tickIndexNormalized))
}

func (p TradePairID) MustPriceTakerToMaker(tickIndexNormalized int64) (priceTakerToMaker sdk.Dec) {
	price, err := p.PriceTakerToMaker(tickIndexNormalized)
	if err != nil {
		panic(err)
	}
	return price
}
