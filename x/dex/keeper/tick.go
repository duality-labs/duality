package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetOrInitTick(goCtx context.Context, pairId string, tickIndex int64) types.Tick {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tick, tickFound := k.GetTick(ctx, pairId, tickIndex)
	if !tickFound {
		numFees := k.GetFeeTierCount(ctx)
		tick := NewTick(pairId, tickIndex, numFees)
		k.SetTick(ctx, pairId, tick)
	}
	return tick
}

func NewLimitOrderTranche() *types.LimitOrderTranche {
	return &types.LimitOrderTranche{
		ReservesTokenIn:  sdk.ZeroDec(),
		ReservesTokenOut: sdk.ZeroDec(),
		TotalTokenIn:     sdk.ZeroDec(),
		TotalTokenOut:    sdk.ZeroDec(),
	}
}

func NewTick(pairId string, tickIndex int64, numFees uint64) types.Tick {
	tick := types.Tick{
		PairId:       pairId,
		TickIndex:    tickIndex,
		TickFeeTiers: make([]*types.TickFeeTier, numFees),
		LimitOrderBook0To1: &types.LimitOrderBook{
			FillTrancheIndex:  0,
			PlaceTrancheIndex: 0,
			Tranches:          []*types.LimitOrderTranche{},
		},
		LimitOrderBook1To0: &types.LimitOrderBook{
			FillTrancheIndex:  0,
			PlaceTrancheIndex: 0,
			Tranches: []*types.LimitOrderTranche{
				{
					ReservesTokenIn:  sdk.ZeroDec(),
					ReservesTokenOut: sdk.ZeroDec(),
					TotalTokenIn:     sdk.ZeroDec(),
					TotalTokenOut:    sdk.ZeroDec(),
				},
			},
		},
	}
	for i := 0; i < int(numFees); i++ {
		tick.TickFeeTiers[i] = &types.TickFeeTier{
			Reserve0:    sdk.ZeroDec(),
			TotalShares: sdk.ZeroDec(),
			Reserve1:    sdk.ZeroDec(),
		}
	}
	return tick
}

// SetTick set a specific Tick in the store from its index
func (k Keeper) SetTick(ctx sdk.Context, pairId string, Tick types.Tick) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickKeyPrefix))
	b := k.cdc.MustMarshal(&Tick)
	store.Set(types.TickKey(
		pairId,
		Tick.TickIndex,
	), b)
}

// GetTick returns a Tick from its index
func (k Keeper) GetTick(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
) (val types.Tick, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickKeyPrefix))

	b := store.Get(types.TickKey(
		pairId,
		tickIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTick removes a Tick from the store
func (k Keeper) RemoveTick(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickKeyPrefix))
	store.Delete(types.TickKey(
		pairId,
		tickIndex,
	))
}

// GetAllTick returns all Tick
func (k Keeper) GetAllTick(ctx sdk.Context) (list []types.Tick) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tick
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllTick returns all Tick
func (k Keeper) GetAllTickByPair(ctx sdk.Context, pairId string) (list []types.Tick) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickPrefix(pairId))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tick
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
