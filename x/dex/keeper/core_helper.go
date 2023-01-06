package keeper

import (
	"context"
	"math"

	. "github.com/NicholasDotSol/duality/utils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxTickExp uint64 = 1048575

///////////////////////////////////////////////////////////////////////////////
//                           GETTERS & INITIALIZERS                          //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) TokenInit(ctx sdk.Context, address string) {
	_, found := k.GetTokenMap(ctx, address)
	if !found {
		tokenIndex := k.GetTokensCount(ctx)
		newTokenCount := tokenIndex + 1
		// TODO: Consolidate TokenMap and Tokens into one type
		k.SetTokenMap(ctx, types.TokenMap{Address: address, Index: int64(tokenIndex)})
		k.AppendTokens(ctx, types.Tokens{Address: address, Id: tokenIndex})
		k.SetTokensCount(ctx, newTokenCount)
	}
}

// Handles initializing a new pair (token0/token1) if not found, adds token0, token1 to global list of tokens active on the dex
func (k Keeper) GetOrInitPair(goCtx context.Context, token0 string, token1 string) types.TradingPair {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.TokenInit(ctx, token0)
	k.TokenInit(ctx, token1)
	pairId := CreatePairId(token0, token1)
	pair, found := k.GetTradingPair(ctx, pairId)
	if !found {
		pair = types.TradingPair{
			PairId:          pairId,
			CurrentTick0To1: math.MaxInt64,
			CurrentTick1To0: math.MinInt64,
			MinTick:         math.MaxInt64,
			MaxTick:         math.MinInt64,
		}
		k.SetTradingPair(ctx, pair)
	}
	return pair
}

func (k Keeper) InitTick(ctx sdk.Context, pairId string, tickIndex int64) (types.Tick, error) {
	price0To1, err := CalcPrice0To1(tickIndex)
	if err != nil {
		return types.Tick{}, err
	}

	numFees := k.GetFeeTierCount(ctx)
	tick := types.Tick{
		PairId:    pairId,
		TickIndex: tickIndex,
		Price0To1: &price0To1,
		TickData: &types.TickDataType{
			// TODO: clean up tickdata proto
			Reserve0: make([]sdk.Int, numFees),
			Reserve1: make([]sdk.Int, numFees),
		},
		LimitOrderTranche0To1: &types.LimitTrancheIndexes{0, 0},
		LimitOrderTranche1To0: &types.LimitTrancheIndexes{0, 0},
	}
	for i := 0; i < int(numFees); i++ {
		// TODO: clean up tickdata proto
		tick.TickData.Reserve0[i] = sdk.ZeroInt()
		tick.TickData.Reserve1[i] = sdk.ZeroInt()
	}
	k.SetTick(ctx, pairId, tick)

	token0, token1 := PairToTokens(pairId)
	k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, token0, 0)
	k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, token1, 0)

	return tick, nil
}

func (k Keeper) GetOrInitTick(goCtx context.Context, pairId string, tickIndex int64) (types.Tick, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tick, tickFound := k.GetTick(ctx, pairId, tickIndex)
	if tickFound {
		return tick, nil
	} else {
		return k.InitTick(ctx, pairId, tickIndex)
	}
}

func (k Keeper) GetOrInitLimitOrderTrancheUser(
	goCtx context.Context,
	pairId string,
	tickIndex int64,
	tokenIn string,
	currentLimitOrderKey uint64,
	receiver string,
) types.LimitOrderTrancheUser {
	ctx := sdk.UnwrapSDKContext(goCtx)

	UserShareData, UserShareDataFound := k.GetLimitOrderTrancheUser(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey, receiver)

	if !UserShareDataFound {
		return types.LimitOrderTrancheUser{
			Count:           currentLimitOrderKey,
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

// Balance trueAmount1 to the pool ratio
func CalcTrueAmounts(
	centerTickPrice1To0 sdk.Dec,
	lowerReserve0 sdk.Int,
	upperReserve1 sdk.Int,
	amount0 sdk.Int,
	amount1 sdk.Int,
	totalShares sdk.Int,
) (trueAmount0 sdk.Int, trueAmount1 sdk.Int, sharesMinted sdk.Int) {
	lowerReserve0Dec := lowerReserve0.ToDec()
	upperReserve1Dec := upperReserve1.ToDec()
	amount0Dec := amount0.ToDec()
	amount1Dec := amount1.ToDec()

	// See spec: https://www.notion.so/dualityxyz/Autoswap-Spec-e856fa7b2438403c95147010d479b98c
	if upperReserve1Dec.GT(sdk.ZeroDec()) {
		trueAmount0 = sdk.MinDec(
			amount0Dec,
			amount1Dec.Mul(lowerReserve0Dec).Quo(upperReserve1Dec)).TruncateInt()
	} else {
		trueAmount0 = amount0
	}

	if lowerReserve0Dec.GT(sdk.ZeroDec()) {
		trueAmount1 = sdk.MinDec(
			amount1Dec,
			amount0Dec.Mul(upperReserve1Dec).Quo(lowerReserve0Dec)).TruncateInt()
	} else {
		trueAmount1 = amount1
	}

	valueMintedToken0 := CalcShares(trueAmount0, trueAmount1, centerTickPrice1To0)
	valueExistingToken0 := CalcShares(lowerReserve0, upperReserve1, centerTickPrice1To0)
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMinted = valueMintedToken0.Quo(valueExistingToken0).MulInt(totalShares).TruncateInt()
	} else {
		sharesMinted = valueMintedToken0.TruncateInt()
	}
	return
}
func IsTickOutOfRange(tickIndex int64) bool {
	absTickIndex := Abs(tickIndex)
	return absTickIndex > MaxTickExp
}

func MustCalcPrice0To1(tickIndex int64) sdk.Dec {
	price, err := CalcPrice0To1(tickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

// Calculates the price for a swap from token 0 to token 1 given a tick
// tickIndex refers to the index of a specified tick
func CalcPrice0To1(tickIndex int64) (sdk.Dec, error) {
	if IsTickOutOfRange(tickIndex) {
		return sdk.ZeroDec(), types.ErrTickOutsideRange
	}

	if 0 <= tickIndex {
		return sdk.OneDec().Quo(Pow(BasePrice(), uint64(tickIndex))), nil
	} else {
		return Pow(BasePrice(), uint64(-1*tickIndex)), nil
	}
}

func MustCalcPrice1To0(tickIndex int64) sdk.Dec {
	price, err := CalcPrice1To0(tickIndex)
	if err != nil {
		panic(err)
	}
	return price
}

// Calculates the price for a swap from token 1 to token 0 given a tick
// tickIndex refers to the index of a specified tick
func CalcPrice1To0(tickIndex int64) (sdk.Dec, error) {

	if IsTickOutOfRange(tickIndex) {
		return sdk.ZeroDec(), types.ErrTickOutsideRange
	}

	if 0 <= tickIndex {
		return Pow(BasePrice(), uint64(tickIndex)), nil
	} else {
		return sdk.OneDec().Quo(Pow(BasePrice(), uint64(-1*tickIndex))), nil
	}
}

func (k Keeper) ShouldDeinitAfterSwap(ctx sdk.Context, deinitedTick *types.Tick, is0To1 bool) bool {
	return deinitedTick != nil &&
		((is0To1 && !k.TickHasToken1(ctx, deinitedTick)) ||
			(!is0To1 && !k.TickHasToken0(ctx, deinitedTick)))
}

// Checks if a tick has reserves0 at any fee tier
func (k Keeper) TickHasToken0(ctx sdk.Context, tick *types.Tick) bool {
	// TODO: clean up tickdata proto
	for _, r0 := range tick.TickData.Reserve0 {
		if r0.GT(sdk.ZeroInt()) {
			return true
		}
	}

	for i := tick.LimitOrderTranche0To1.FillTrancheIndex; i <= tick.LimitOrderTranche0To1.PlaceTrancheIndex; i++ {
		if k.TickTrancheHasToken0(ctx, tick, i) {
			return true
		}
	}

	return false
}

func (k Keeper) TickTrancheHasToken0(ctx sdk.Context, tick *types.Tick, trancheIndex uint64) bool {
	token0, _ := PairToTokens(tick.PairId)
	tranche, found := k.GetLimitOrderTranche(
		ctx,
		tick.PairId,
		tick.TickIndex,
		token0,
		trancheIndex,
	)
	if found && tranche.ReservesTokenIn.GT(sdk.ZeroInt()) {
		return true
	} else {
		return false
	}
}

// Checks if a tick has reserve1 at any fee tier
func (k Keeper) TickHasToken1(ctx sdk.Context, tick *types.Tick) bool {
	for _, s := range tick.TickData.Reserve1 {
		if s.GT(sdk.ZeroInt()) {
			return true
		}
	}

	for i := tick.LimitOrderTranche1To0.FillTrancheIndex; i <= tick.LimitOrderTranche1To0.PlaceTrancheIndex; i++ {
		if k.TickTrancheHasToken1(ctx, tick, i) {
			return true
		}
	}

	return false
}

func (k Keeper) TickTrancheHasToken1(ctx sdk.Context, tick *types.Tick, trancheIndex uint64) bool {
	_, token1 := PairToTokens(tick.PairId)
	tranche, found := k.GetLimitOrderTranche(
		ctx,
		tick.PairId,
		tick.TickIndex,
		token1,
		trancheIndex,
	)
	return found && tranche.ReservesTokenIn.GT(sdk.ZeroInt())
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
