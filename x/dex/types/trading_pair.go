package types

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper interface {
	TickHasToken0(sdk.Context, *Tick) bool
	TickHasToken1(sdk.Context, *Tick) bool
	NewTickIterator(context.Context, int64, int64, *PairId, bool, codec.BinaryCodec) TickIteratorI
	GetCdc() codec.BinaryCodec
}

func PairIdToTokens(pairId *PairId) (token0 string, token1 string) {

	return pairId.Token0, pairId.Token1
}

func (p TradingPair) ToTokens() (token0 string, token1 string) {
	return PairIdToTokens(p.PairId)
}
