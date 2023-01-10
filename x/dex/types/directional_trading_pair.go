package types

type DirectionalTradingPair struct {
	TradingPair
	TokenIn  string
	TokenOut string
}

func NewDirectionalTradingPair(pair TradingPair, tokenIn string, tokenOut string) DirectionalTradingPair {
	return DirectionalTradingPair{
		TradingPair: pair,
		TokenIn:     tokenIn,
		TokenOut:    tokenOut,
	}
}

func (dp DirectionalTradingPair) IsTokenInToken0() bool {
	return dp.TokenIn == dp.PairId.Token0
}

func (dp DirectionalTradingPair) IsTokenOutToken0() bool {
	return !dp.IsTokenInToken0()
}
