package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTradingPair set a specific TradingPair in the store from its index
func (k Keeper) SetTradingPair(ctx sdk.Context, TradingPair types.TradingPair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingPairKeyPrefix))
	b := k.cdc.MustMarshal(&TradingPair)
	store.Set(types.TradingPairKey(
		TradingPair.PairId,
	), b)
}

// GetTradingPair returns a TradingPair from its index
func (k Keeper) GetTradingPair(
	ctx sdk.Context,
	pairId string,
) (val types.TradingPair, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingPairKeyPrefix))

	b := store.Get(types.TradingPairKey(
		pairId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetDirectionalTradingPair(
	ctx sdk.Context,
	tokenIn string,
	tokenOut string,
) (pair types.DirectionalTradingPair, err error) {
	token0, token1, err := SortTokens(ctx, tokenIn, tokenOut)
	if err != nil {
		return pair, err
	}
	pairId := CreatePairId(token0, token1)
	rawPair, found := k.GetTradingPair(ctx, pairId)
	if !found {
		return pair, types.ErrValidPairNotFound
	} else {
		dPair := types.DirectionalTradingPair{
			TradingPair: rawPair,
			TokenIn:     tokenIn,
			TokenOut:    tokenOut,
			Token0:      token0,
		}
		return dPair, nil
	}
}

// RemoveTradingPair removes a TradingPair from the store
func (k Keeper) RemoveTradingPair(
	ctx sdk.Context,
	pairId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingPairKeyPrefix))
	store.Delete(types.TradingPairKey(
		pairId,
	))
}

// GetAllTradingPair returns all TradingPair
func (k Keeper) GetAllTradingPair(ctx sdk.Context) (list []types.TradingPair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingPairKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradingPair
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
