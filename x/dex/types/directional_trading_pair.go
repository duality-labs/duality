package types

import (
	context "context"
)

type DirectionalTradingPair struct {
	TradingPair
	TokenIn  string
	TokenOut string
	Token0   string
}

func NewDirectionalTradingPair(pair TradingPair, tokenIn string, tokenOut string) DirectionalTradingPair {
	token0, _ := pair.PairToTokens()
	return DirectionalTradingPair{
		TradingPair: pair,
		TokenIn:     tokenIn,
		TokenOut:    tokenOut,
		Token0:      token0,
	}
}

func (dp DirectionalTradingPair) IsTokenInToken0() bool {
	return dp.TokenIn == dp.Token0
}

func (dp DirectionalTradingPair) IsTokenOutToken0() bool {
	return !dp.IsTokenInToken0()
}

func (dp *DirectionalTradingPair) InitLiquidity(tickIndex int64) {
	if dp.IsTokenInToken0() {
		dp.InitLiquidityToken0(tickIndex)
	} else {
		dp.InitLiquidityToken1(tickIndex)
	}
}

func (dp *DirectionalTradingPair) DeinitLiquidity(ctx context.Context, k Keeper, tickIndex int64) {
	if dp.IsTokenOutToken0() {
		dp.DeinitLiquidityToken0(ctx, k, tickIndex)
	} else {
		dp.DeinitLiquidityToken1(ctx, k, tickIndex)
	}
}
