package types

type DirectionalTradingPair struct {
	*PairId
	TokenIn  string
	TokenOut string
}

func NewDirectionalTradingPair(pairId *PairId, tokenIn, tokenOut string) DirectionalTradingPair {
	return DirectionalTradingPair{
		PairId:   pairId,
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
	}
}

func (dp DirectionalTradingPair) IsTokenInToken0() bool {
	return dp.TokenIn == dp.PairId.Token0
}

func (dp DirectionalTradingPair) IsTokenOutToken0() bool {
	return !dp.IsTokenInToken0()
}
