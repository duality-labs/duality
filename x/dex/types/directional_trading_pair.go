package types

import (
	context "context"
)

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
