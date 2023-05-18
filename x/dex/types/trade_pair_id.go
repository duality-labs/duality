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

func (p TradePairID) PriceMakerToTaker(normalizedTickIndex int64) (priceMakerToTaker sdk.Dec, err error) {
	pairID, err := p.PairID()
	if err != nil {
		return sdk.ZeroDec(), err
	}

	if pairID.Token0 == p.MakerDenom {
		return CalcPrice0To1(normalizedTickIndex)
	} else {
		return CalcPrice1To0(normalizedTickIndex)
	}
}

func (p TradePairID) PriceTakerToMaker(normalizedTickIndex int64) (priceTakerToMaker sdk.Dec, err error) {
	pairID, err := p.PairID()
	if err != nil {
		return sdk.ZeroDec(), err
	}

	if pairID.Token0 == p.MakerDenom {
		return CalcPrice1To0(normalizedTickIndex)
	} else {
		return CalcPrice0To1(normalizedTickIndex)
	}
}

func (p TradePairID) MustPriceTakerToMaker(normalizedTickIndex int64) (priceTakerToMaker sdk.Dec) {
	price, err := p.PriceTakerToMaker(normalizedTickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

func (p TradePairID) MustPriceMakerToTaker(normalizedTickIndex int64) (priceMakerToTaker sdk.Dec) {
	price, err := p.PriceMakerToTaker(normalizedTickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

func (p TradePairID) Reversed() *TradePairID {
	return &TradePairID{
		MakerDenom: p.TakerDenom,
		TakerDenom: p.MakerDenom,
	}
}
