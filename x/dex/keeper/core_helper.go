package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

///////////////////////////////////////////////////////////////////////////////
//                           GETTERS & INITIALIZERS                          //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) GetOrInitPoolReserves(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	tickIndex int64,
	fee uint64,
) (*types.PoolReserves, error) {
	tickLiq, tickFound := k.GetPoolReserves(
		ctx,
		tradePairID,
		tickIndex,
		fee,
	)
	switch {
	case tickFound:
		return tickLiq, nil
	case types.IsTickOutOfRange(tickIndex):
		return nil, types.ErrTickOutsideRange
	default:
		var price0To1, price1To0, priceMakerToTaker, priceTakerToMaker sdk.Dec
		var err error
		price0To1, err = types.CalcPrice0To1(tickIndex)
		if err != nil {
			return nil, err
		}
		price1To0, err = types.CalcPrice1To0(tickIndex)
		if err != nil {
			return nil, err
		}
		if tradePairID.IsMakerDenomToken0() {
			priceMakerToTaker = price0To1
			priceTakerToMaker = price1To0
		} else {
			priceMakerToTaker = price1To0
			priceTakerToMaker = price0To1
		}
		return &types.PoolReserves{
			TradePairID:        tradePairID,
			TickIndex:          tickIndex,
			Fee:                fee,
			ReservesMakerDenom: sdk.ZeroInt(),
			PriceMakerToTaker:  priceMakerToTaker,
			PriceTakerToMaker:  priceTakerToMaker,
		}, nil
	}
}

func NewLimitOrderExpiration(tranche *types.LimitOrderTranche) types.LimitOrderExpiration {
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
	tradePairID *types.TradePairID,
	tickIndex int64,
	goodTil *time.Time,
) (*types.LimitOrderTranche, error) {
	priceMakerToTaker, err := tradePairID.PriceMakerToTaker(tickIndex)
	if err != nil {
		return nil, err
	}
	priceTakerToMaker, err := tradePairID.PriceTakerToMaker(tickIndex)
	if err != nil {
		return nil, err
	}
	return &types.LimitOrderTranche{
		TradePairID:        tradePairID,
		TickIndex:          tickIndex,
		TrancheKey:         NewTrancheKey(sdkCtx),
		ReservesMakerDenom: sdk.ZeroInt(),
		ReservesTakerDenom: sdk.ZeroInt(),
		TotalMakerDenom:    sdk.ZeroInt(),
		TotalTakerDenom:    sdk.ZeroInt(),
		ExpirationTime:     goodTil,
		PriceTakerToMaker:  priceTakerToMaker,
		PriceMakerToTaker:  priceMakerToTaker,
	}, nil
}

func (k Keeper) GetOrInitLimitOrderTrancheUser(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
	tickIndex int64,
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
			TradePairID:     tradePairID,
			OrderType:       orderType,
		}
	}

	return UserShareData
}

///////////////////////////////////////////////////////////////////////////////
//                          STATE CALCULATIONS                               //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) GetCurrPrice(ctx sdk.Context, tradePairID *types.TradePairID) (sdk.Dec, bool) {
	liq := k.GetCurrLiq(ctx, tradePairID)
	if liq != nil {
		return liq.Price(), true
	}
	return sdk.ZeroDec(), false
}

func (k Keeper) GetCurrTick(ctx sdk.Context, tradePairID *types.TradePairID) (int64, bool) {
	liq := k.GetCurrLiq(ctx, tradePairID)
	if liq != nil {
		return liq.TickIndex(), true
	}
	return 0, false
}

func (k Keeper) GetCurrLiq(ctx sdk.Context, tradePairID *types.TradePairID) *types.TickLiquidity {
	ti := k.NewTickIterator(ctx, tradePairID)
	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if tick.HasToken() {
			return &tick
		}
	}

	return nil
}

func CalcAmountAsToken0(amount0, amount1 sdk.Int, price1To0 sdk.Dec) sdk.Dec {
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
