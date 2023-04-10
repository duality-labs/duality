package keeper

import (
	"math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

///////////////////////////////////////////////////////////////////////////////
//                           GETTERS & INITIALIZERS                          //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) GetOrInitPoolReserves(
	ctx sdk.Context,
	pairID *types.PairID,
	tokenIn string,
	tickIndex int64,
	fee uint64,
) (*types.PoolReserves, error) {
	tickLiq, tickFound := k.GetPoolReserves(
		ctx,
		pairID,
		tokenIn,
		tickIndex,
		fee,
	)
	switch {
	case tickFound:
		return tickLiq, nil
	case types.IsTickOutOfRange(tickIndex):
		return nil, types.ErrTickOutsideRange
	default:
		return &types.PoolReserves{
			PairID:    pairID,
			TokenIn:   tokenIn,
			TickIndex: tickIndex,
			Fee:       fee,
			Reserves:  sdk.ZeroInt(),
		}, nil
	}
}

func NewLimitOrderExpiration(tranche types.LimitOrderTranche) types.LimitOrderExpiration {
	trancheExpiry := tranche.ExpirationTime
	if trancheExpiry == nil {
		panic("Cannot create LimitOrderExpiration from tranche with nil ExpirationTime")
	}

	return types.LimitOrderExpiration{
		TrancheRef:     tranche.Ref(),
		ExpirationTime: *tranche.ExpirationTime,
	}
}

func NewLimitOrderTranche(
	sdkCtx sdk.Context,
	pairID *types.PairID,
	tokenIn string,
	tickIndex int64,
	goodTil *time.Time,
) (types.LimitOrderTranche, error) {
	// NOTE: CONTRACT: There is no active place tranche (ie. GetPlaceTrancheTick has returned false)
	if types.IsTickOutOfRange(tickIndex) {
		return types.LimitOrderTranche{}, types.ErrTickOutsideRange
	}
	trancheKey := NewTrancheKey(sdkCtx)

	return types.LimitOrderTranche{
		PairID:           pairID,
		TokenIn:          tokenIn,
		TickIndex:        tickIndex,
		TrancheKey:       trancheKey,
		ReservesTokenIn:  sdk.ZeroInt(),
		ReservesTokenOut: sdk.ZeroInt(),
		TotalTokenIn:     sdk.ZeroInt(),
		TotalTokenOut:    sdk.ZeroInt(),
		ExpirationTime:   goodTil,
	}, nil
}

func (k Keeper) GetOrInitLimitOrderTrancheUser(
	ctx sdk.Context,
	pairID *types.PairID,
	tickIndex int64,
	tokenIn string,
	trancheKey string,
	orderType types.LimitOrderType,
	receiver string,
) types.LimitOrderTrancheUser {
	UserShareData, UserShareDataFound := k.GetLimitOrderTrancheUser(ctx, receiver, trancheKey)

	if !UserShareDataFound {
		return types.LimitOrderTrancheUser{
			TrancheKey:      trancheKey,
			Address:         receiver,
			SharesOwned:     sdk.ZeroInt(),
			SharesWithdrawn: sdk.ZeroInt(),
			SharesCancelled: sdk.ZeroInt(),
			TickIndex:       tickIndex,
			Token:           tokenIn,
			PairID:          pairID,
			OrderType:       orderType,
		}
	}

	return UserShareData
}

///////////////////////////////////////////////////////////////////////////////
//                          STATE CALCULATIONS                               //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) GetCurrTick1To0(ctx sdk.Context, pairID *types.PairID) (tickIdx int64, found bool) {
	ti := k.NewTickIterator(ctx, pairID, pairID.Token0)

	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if tick.HasToken() {
			return tick.TickIndex(), true
		}
	}

	return math.MinInt64, false
}

func (k Keeper) GetCurrTick0To1(ctx sdk.Context, pairID *types.PairID) (tickIdx int64, found bool) {
	ti := k.NewTickIterator(ctx, pairID, pairID.Token1)
	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if tick.HasToken() {
			return tick.TickIndex(), true
		}
	}

	return math.MaxInt64, false
}

func (k Keeper) IsBehindEnemyLines(ctx sdk.Context, pairID *types.PairID, tokenIn string, tickIndex int64) bool {
	if tokenIn == pairID.Token0 {
		curr0To1, _ := k.GetCurrTick0To1(ctx, pairID)
		if tickIndex >= curr0To1 {
			return true
		}
	} else {
		curr1To0, _ := k.GetCurrTick1To0(ctx, pairID)
		if tickIndex <= curr1To0 {
			return true
		}
	}

	return false
}

func CalcAmountAsToken0(amount0, amount1 sdk.Int, price1To0 types.Price) sdk.Dec {
	amount0Dec := amount0.ToDec()

	return amount0Dec.Add(price1To0.MulInt(amount1))
}

///////////////////////////////////////////////////////////////////////////////
//                            TOKENIZER UTILS                                //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) MintShares(ctx sdk.Context, addr sdk.AccAddress, shareCoin sdk.Coin) error {
	// mint share tokens
	sharesCoins := sdk.Coins{shareCoin}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// transfer them to addr
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sharesCoins)

	return err
}

func (k Keeper) BurnShares(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int, sharesID string) error {
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesID, amount)}
	// transfer tokens to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// burn tokens
	err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sharesCoins)

	return err
}
