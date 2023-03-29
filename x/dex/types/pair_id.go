package types

import (
	"fmt"
)

func (p *PairID) Stringify() string {
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
	} else {
		panic("Supplied token matches neither side of pair")
	}
}
