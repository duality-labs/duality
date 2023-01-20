package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type LimitOrderTranche struct {
	Tranche           *types.LimitOrderTranche
	Tick              *types.Tick
	PriceTakerToMaker sdk.Dec
	PriceMakerToTaker sdk.Dec
}

func NewLimitOrderTranche(
	tick *types.Tick,
	tranche *types.LimitOrderTranche,
	priceMakerToTaker sdk.Dec,
	priceTakerToMaker sdk.Dec,
) *LimitOrderTranche {
	return &LimitOrderTranche{
		Tranche:           tranche,
		Tick:              tick,
		PriceTakerToMaker: priceTakerToMaker,
		PriceMakerToTaker: priceMakerToTaker,
	}
}

func (t *LimitOrderTranche) Swap(maxAmountTaker sdk.Int) (
	inAmount sdk.Int,
	outAmount sdk.Int,
	initedTick *types.Tick,
	deinitedTick *types.Tick,
) {
	tranche := t.Tranche
	reservesTokenOut := &tranche.ReservesTokenIn
	fillTokenIn := &tranche.ReservesTokenOut
	totalTokenIn := &tranche.TotalTokenOut
	amountFilledTokenOut := maxAmountTaker.ToDec().Mul(t.PriceTakerToMaker).TruncateInt()
	if reservesTokenOut.LTE(amountFilledTokenOut) {
		inAmount = reservesTokenOut.ToDec().Mul(t.PriceMakerToTaker).TruncateInt()
		outAmount = *reservesTokenOut
		*reservesTokenOut = sdk.ZeroInt()
		*fillTokenIn = fillTokenIn.Add(inAmount)
		*totalTokenIn = totalTokenIn.Add(inAmount)
		deinitedTick = t.Tick
	} else {
		inAmount = maxAmountTaker
		outAmount = amountFilledTokenOut
		*fillTokenIn = fillTokenIn.Add(maxAmountTaker)
		*totalTokenIn = totalTokenIn.Add(maxAmountTaker)
		*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
	}
	return inAmount, outAmount, nil, deinitedTick
}

func (t *LimitOrderTranche) Save(ctx context.Context, keeper Keeper) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	keeper.SetTick(sdkCtx, t.Tick.PairId, *t.Tick)
	keeper.SetLimitOrderTranche(sdkCtx, *t.Tranche)
}

func (t *LimitOrderTranche) Price() sdk.Dec {
	return t.PriceTakerToMaker
}

func (t LimitOrderTranche) HasLiquidity() bool {
	return t.Tranche.ReservesTokenIn.GT(sdk.ZeroInt())
}

// SetLimitOrderTranche set a specific LimitOrderTranche in the store from its index
func (k Keeper) SetLimitOrderTranche(ctx sdk.Context, LimitOrderTranche types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))
	b := k.cdc.MustMarshal(&LimitOrderTranche)
	store.Set(types.LimitOrderTrancheKey(
		LimitOrderTranche.PairId,
		LimitOrderTranche.TickIndex,
		LimitOrderTranche.TokenIn,
		LimitOrderTranche.TrancheIndex,
	), b)
}

func (k Keeper) GetOrInitLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tickIndex int64,
	tokenIn string,
	trancheIndex uint64,
) types.LimitOrderTranche {
	tranche, found := k.GetLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, trancheIndex)
	if !found {
		tranche = types.LimitOrderTranche{
			TrancheIndex:     trancheIndex,
			TickIndex:        tickIndex,
			TokenIn:          tokenIn,
			PairId:           pairId,
			ReservesTokenIn:  sdk.ZeroInt(),
			ReservesTokenOut: sdk.ZeroInt(),
			TotalTokenIn:     sdk.ZeroInt(),
			TotalTokenOut:    sdk.ZeroInt(),
		}
		k.SetLimitOrderTranche(ctx, tranche)
	}

	return tranche
}

// GetLimitOrderTranche returns a LimitOrderTranche from its index
func (k Keeper) GetLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tickIndex int64,
	token string,
	tranchIndex uint64,

) (val types.LimitOrderTranche, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))

	b := store.Get(types.LimitOrderTrancheKey(
		pairId,
		tickIndex,
		token,
		tranchIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLimitOrderTranche removes a LimitOrderTranche from the store
func (k Keeper) RemoveLimitOrderTranche(
	ctx sdk.Context,
	pairId *types.PairId,
	tickIndex int64,
	token string,
	trancheIndex uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))
	store.Delete(types.LimitOrderTrancheKey(
		pairId,
		tickIndex,
		token,
		trancheIndex,
	))
}

// GetAllLimitOrderTranche returns all LimitOrderTrancheUser
func (k Keeper) GetAllLimitOrderTranche(ctx sdk.Context) (list []types.LimitOrderTranche) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderTranche
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
