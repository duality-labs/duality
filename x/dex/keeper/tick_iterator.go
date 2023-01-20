package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type TickIterator struct {
	iter sdk.Iterator
	cdc  codec.BinaryCodec
}

func (k Keeper) NewTickIterator(
	ctx sdk.Context,
	pairId *types.PairId,
	tokenIn string,
) TickIterator {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickLiquidityPrefix(pairId, tokenIn))

	return TickIterator{
		iter: prefixStore.Iterator(nil, nil),
		cdc:  k.cdc,
	}
}

func (ti TickIterator) Valid() bool {
	return ti.iter.Valid()
}

func (ti TickIterator) Close() error {
	return ti.iter.Close()
}

func (ti TickIterator) Value() types.TickLiquidity {
	var tick types.TickLiquidity
	ti.cdc.MustUnmarshal(ti.iter.Value(), &tick)
	return tick
}

func (ti TickIterator) Next() {
	ti.iter.Next()
}
