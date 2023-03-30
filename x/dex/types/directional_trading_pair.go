package types

type DirectionalTradingPair struct {
	*PairID
	TokenIn  string
	TokenOut string
}

func NewDirectionalTradingPair(pairID *PairID, tokenIn, tokenOut string) DirectionalTradingPair {
	return DirectionalTradingPair{
		PairID:   pairID,
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
	}
}

func (dp DirectionalTradingPair) IsTokenInToken0() bool {
	return dp.TokenIn == dp.PairID.Token0
}

func (dp DirectionalTradingPair) IsTokenOutToken0() bool {
	return !dp.IsTokenInToken0()
}
