package types

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewPairID(token0, token1 string) (*PairID, error) {
	if token0 == token1 {
		return nil, sdkerrors.Wrapf(ErrInvalidTradingPair, "%s<>%s", token0, token1)
	}
	return &PairID{
		Token0: token0,
		Token1: token1,
	}, nil
}

func MustNewPairID(token0, token1 string) *PairID {
	pairID, err := NewPairID(token0, token1)
	if err != nil {
		panic(err)
	}
	return pairID
}

func NewPairIDFromUnsorted(tokenA, tokenB string) (*PairID, error) {
	token0, token1 := SortTokens(tokenA, tokenB)
	return NewPairID(token0, token1)
}

func (p *PairID) CanonicalString() string {
	return fmt.Sprintf("%s<>%s", p.Token0, p.Token1)
}

func (p *PairID) OppositeToken(token string) (oppToken string, ok bool) {
	switch token {
	case p.Token0:
		return p.Token1, true
	case p.Token1:
		return p.Token0, true
	default:
		return "", false
	}
}

func (p *PairID) MustOppositeToken(token string) string {
	if oppToken, ok := p.OppositeToken(token); ok {
		return oppToken
	}
	panic("Supplied token matches neither side of pair")
}

func NewPairIDFromCanonicalString(pairIDStr string) (*PairID, error) {
	tokens := strings.Split(pairIDStr, "<>")

	if len(tokens) == 2 {
		return NewPairID(tokens[0], tokens[1])
	}

	return &PairID{}, sdkerrors.Wrapf(ErrInvalidPairIDStr, pairIDStr)
}

func SortTokens(tokenA, tokenB string) (string, string) {
	if tokenA < tokenB {
		return tokenA, tokenB
	} else {
		return tokenB, tokenA
	}
}

func (p *PairID) MustTradePairIDFromMaker(maker string) *TradePairID {
	if p.Token0 == maker {
		return MustNewTradePairID(p.Token1, p.Token0)
	} else if p.Token1 == maker {
		return MustNewTradePairID(p.Token0, p.Token1)
	} else {
		panic(fmt.Errorf("pair.TradePairIDFromMaker(maker string) called where maker does not equal either pair.Token0 or pair.Token1"))
	}
}

func (p *PairID) MustTradePairIDFromTaker(taker string) *TradePairID {
	if p.Token0 == taker {
		return MustNewTradePairID(p.Token0, p.Token1)
	} else if p.Token1 == taker {
		return MustNewTradePairID(p.Token1, p.Token0)
	} else {
		panic(fmt.Errorf("pair.TradePairIDFromMaker(maker string) called where maker does not equal either pair.Token0 or pair.Token1"))
	}
}
