package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetShares set a specific shares in the store from its index
func (k Keeper) SetShares(ctx sdk.Context, shares types.Shares) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SharesKeyPrefix))
	b := k.cdc.MustMarshal(&shares)
	store.Set(types.SharesKey(
		shares.Address,
		shares.PairId,
		shares.TickIndex,
		shares.FeeIndex,
	), b)
}

// GetShares returns a shares from its index
func (k Keeper) GetShares(
	ctx sdk.Context,
	address string,
	pairId string,
	tickIndex int64,
	feeIndex uint64,
) (val types.Shares, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SharesKeyPrefix))

	b := store.Get(types.SharesKey(
		address,
		pairId,
		tickIndex,
		feeIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShares removes a shares from the store
func (k Keeper) RemoveShares(
	ctx sdk.Context,
	address string,
	pairId string,
	tickIndex int64,
	feeIndex uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SharesKeyPrefix))
	store.Delete(types.SharesKey(
		address,
		pairId,
		tickIndex,
		feeIndex,
	))
}

// GetAllShares returns all shares
func (k Keeper) GetAllShares(ctx sdk.Context) (list []types.Shares) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SharesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Shares
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
