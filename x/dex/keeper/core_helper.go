package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	"golang.org/x/exp/slices"
)

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

// Returns a takerToMaker tick index
func (k Keeper) GetCurrTickIndexTakerToMaker(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
) (int64, bool) {
	liq := k.GetCurrLiq(ctx, tradePairID)
	if liq != nil {
		return liq.TickIndex(), true
	}
	return 0, false
}

// Returns a takerToMaker tick index
func (k Keeper) GetCurrTickIndexTakerToMakerNormalized(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
) (int64, bool) {
	tickIndexTakerToMaker, found := k.GetCurrTickIndexTakerToMaker(ctx, tradePairID)
	if found {
		tickIndexTakerToMakerNormalized := tradePairID.TickIndexNormalized(tickIndexTakerToMaker)
		return tickIndexTakerToMakerNormalized, true
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

func (k Keeper) GetValidFees(ctx sdk.Context) []uint64 {
	return k.GetParams(ctx).FeeTiers
}

func (k Keeper) ValidateFee(ctx sdk.Context, fee uint64) error {
	validFees := k.GetValidFees(ctx)
	if !slices.Contains(validFees, fee) {
		return sdkerrors.Wrapf(types.ErrInvalidFee, "%s", validFees)
	}

	return nil
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

func (k Keeper) BurnShares(
	ctx sdk.Context,
	addr sdk.AccAddress,
	amount sdk.Int,
	sharesID string,
) error {
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesID, amount)}
	// transfer tokens to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// burn tokens
	err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sharesCoins)

	return err
}
