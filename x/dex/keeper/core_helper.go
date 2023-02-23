package keeper

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/duality-labs/duality/utils"
	"github.com/duality-labs/duality/x/dex/types"
)

///////////////////////////////////////////////////////////////////////////////
//                           GETTERS & INITIALIZERS                          //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) GetOrInitPoolReserves(ctx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64, fee uint64) (*types.PoolReserves, error) {
	tickLiq, tickFound := k.GetPoolReserves(
		ctx,
		pairId,
		tokenIn,
		tickIndex,
		fee,
	)
	if tickFound {
		return tickLiq, nil
	} else if IsTickOutOfRange(tickIndex) {
		return nil, types.ErrTickOutsideRange
	} else {
		return &types.PoolReserves{
			PairId:    pairId,
			TokenIn:   tokenIn,
			TickIndex: tickIndex,
			Fee:       fee,
			Reserves:  sdk.ZeroInt(),
		}, nil
	}

}

func NewLimitOrderTranche(pairId *types.PairId, tokenIn string, tickIndex int64, trancheKey string) (types.LimitOrderTranche, error) {
	if IsTickOutOfRange(tickIndex) {
		return types.LimitOrderTranche{}, types.ErrTickOutsideRange
	}
	return types.LimitOrderTranche{
		PairId:           pairId,
		TokenIn:          tokenIn,
		TickIndex:        tickIndex,
		TrancheKey:       trancheKey,
		ReservesTokenIn:  sdk.ZeroInt(),
		ReservesTokenOut: sdk.ZeroInt(),
		TotalTokenIn:     sdk.ZeroInt(),
		TotalTokenOut:    sdk.ZeroInt(),
	}, nil

}

func (k Keeper) GetOrInitLimitOrderTrancheUser(
	ctx sdk.Context,
	pairId *types.PairId,
	tickIndex int64,
	tokenIn string,
	currentLimitOrderKey string,
	receiver string,
) types.LimitOrderTrancheUser {
	UserShareData, UserShareDataFound := k.GetLimitOrderTrancheUser(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey, receiver)

	if !UserShareDataFound {
		return types.LimitOrderTrancheUser{
			TrancheKey:      currentLimitOrderKey,
			Address:         receiver,
			SharesOwned:     sdk.ZeroInt(),
			SharesWithdrawn: sdk.ZeroInt(),
			SharesCancelled: sdk.ZeroInt(),
			TickIndex:       tickIndex,
			Token:           tokenIn,
			PairId:          pairId,
		}
	}

	return UserShareData
}

///////////////////////////////////////////////////////////////////////////////
//                          STATE CALCULATIONS                               //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) GetCurrTick1To0(ctx sdk.Context, pairId *types.PairId) (tickIdx int64, found bool) {

	ti := k.NewTickIterator(ctx, pairId, pairId.Token0)

	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if tick.HasToken() {
			return tick.TickIndex(), true
		}
	}
	return math.MinInt64, false

}

func (k Keeper) GetCurrTick0To1(ctx sdk.Context, pairId *types.PairId) (tickIdx int64, found bool) {
	ti := k.NewTickIterator(ctx, pairId, pairId.Token1)
	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if tick.HasToken() {
			return tick.TickIndex(), true
		}
	}

	return math.MaxInt64, false
}

func (k Keeper) IsBehindEnemyLines(ctx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) bool {
	if tokenIn == pairId.Token0 {
		curr0To1, _ := k.GetCurrTick0To1(ctx, pairId)
		if tickIndex >= curr0To1 {
			return true
		}
	} else {

		curr1To0, _ := k.GetCurrTick1To0(ctx, pairId)
		if tickIndex <= curr1To0 {
			return true
		}
	}
	return false
}

///////////////////////////////////////////////////////////////////////////////
//                            TOKENIZER UTILS                                //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) MintShares(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int, sharesId string) error {
	// mint share tokens
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesId, amount)}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// transfer them to addr
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sharesCoins); err != nil {
		return err
	}
	return nil
}

func (k Keeper) BurnShares(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int, sharesId string) error {
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesId, amount)}
	// transfer tokens to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// burn tokens
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	return nil
}
