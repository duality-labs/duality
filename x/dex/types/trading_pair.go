package types

func PairIdToTokens(pairId *PairId) (token0 string, token1 string) {

	return pairId.Token0, pairId.Token1
}

func (p TradingPair) ToTokens() (token0 string, token1 string) {
	return PairIdToTokens(p.PairId)
}
