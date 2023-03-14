package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

// NOTE: Currently we are using TruncateInt in multiple places for converting Decs back into sdk.Ints.
// This may create some accounting anomalies but seems preferable to other alternatives.
// See full ADR here: https://www.notion.so/dualityxyz/A-Modest-Proposal-For-Truncating-696a919d59254876a617f82fb9567895

// Handles core logic for MsgDeposit, checking and initializing data structures (tick, pair), calculating
// shares based on amount deposited, and sending funds to moduleAddress.
func (k Keeper) DepositCore(
	goCtx context.Context,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	amounts0 []sdk.Int,
	amounts1 []sdk.Int,
	tickIndices []int64,
	feeIndices []uint64,
	options []*types.DepositOptions,
) (amounts0Deposit []sdk.Int, amounts1Deposit []sdk.Int, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := CreatePairId(token0, token1)
	totalAmountReserve0 := sdk.ZeroInt()
	totalAmountReserve1 := sdk.ZeroInt()
	amounts0Deposited := make([]sdk.Int, len(amounts0))
	amounts1Deposited := make([]sdk.Int, len(amounts1))
	for i := 0; i < len(amounts0); i++ {
		amounts0Deposited[i] = sdk.ZeroInt()
		amounts1Deposited[i] = sdk.ZeroInt()
	}

	feeTiers := k.GetAllFeeTier(ctx)

	for i, amount0 := range amounts0 {
		amount1 := amounts1[i]
		tickIndex := tickIndices[i]
		feeIndex := feeIndices[i]
		autoswap := options[i].Autoswap

		// check that feeIndex is a valid index of the fee tier
		if feeIndex >= uint64(len(feeTiers)) {
			return nil, nil, sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "%d", feeIndex)
		}
		fee := feeTiers[feeIndex].Fee
		feeUInt := utils.MustSafeUint64(fee)
		lowerTickIndex := tickIndex - feeUInt
		upperTickIndex := tickIndex + feeUInt

		// behind enemy lines checks
		// TODO: Allow user to deposit "behind enemy lines"
		if amount0.GT(sdk.ZeroInt()) && k.IsBehindEnemyLines(ctx, pairId, pairId.Token0, lowerTickIndex) {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}
		// TODO: Allow user to deposit "behind enemy lines"
		if amount1.GT(sdk.ZeroInt()) && k.IsBehindEnemyLines(ctx, pairId, pairId.Token1, upperTickIndex) {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		sharesId := CreateSharesId(token0, token1, tickIndex, feeIndex)
		existingShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		pool, err := k.GetOrInitPool(
			ctx,
			pairId,
			tickIndex,
			feeTiers[feeIndex],
		)

		if err != nil {
			return nil, nil, err
		}

		inAmount0, inAmount1, outShares := pool.Deposit(amount0, amount1, existingShares, autoswap)

		k.SavePool(ctx, pool)

		if inAmount0.Equal(sdk.ZeroInt()) && inAmount1.Equal(sdk.ZeroInt()) {
			return nil, nil, types.ErrZeroTrueDeposit
		}

		if outShares.IsZero() {
			return nil, nil, types.ErrDepositShareUnderflow
		}

		if err := k.MintShares(ctx, receiverAddr, outShares, sharesId); err != nil {
			return nil, nil, err
		}

		amounts0Deposited[i] = inAmount0
		amounts1Deposited[i] = inAmount1
		totalAmountReserve0 = totalAmountReserve0.Add(inAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(inAmount1)

		ctx.EventManager().EmitEvent(types.CreateDepositEvent(
			callerAddr.String(),
			receiverAddr.String(),
			token0,
			token1,
			fmt.Sprint(tickIndices[i]),
			fmt.Sprint(feeIndices[i]),
			pool.GetLowerReserve0().Sub(inAmount0).String(),
			pool.GetUpperReserve1().Sub(inAmount1).String(),
			pool.GetLowerReserve0().String(),
			pool.GetUpperReserve1().String(),
			outShares.String(),
		))
	}

	if totalAmountReserve0.GT(sdk.ZeroInt()) {
		coin0 := sdk.NewCoin(token0, totalAmountReserve0)
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, nil, err
		}
	}

	if totalAmountReserve1.GT(sdk.ZeroInt()) {
		coin1 := sdk.NewCoin(token1, totalAmountReserve1)
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, nil, err
		}
	}

	return amounts0Deposited, amounts1Deposited, nil
}

// Handles core logic for MsgWithdrawl; calculating and withdrawing reserve0,reserve1 from a specified tick given a specfied number of shares to remove.
// Calculates the amount of reserve0, reserve1 to withdraw based on the percetange of the desired number of shares to remove compared to the total number of shares at the given tick
func (k Keeper) WithdrawCore(
	goCtx context.Context,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	sharesToRemoveList []sdk.Int,
	tickIndices []int64,
	feeIndices []uint64,
) error {

	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := CreatePairId(token0, token1)
	totalReserve0ToRemove := sdk.ZeroInt()
	totalReserve1ToRemove := sdk.ZeroInt()
	feeTiers := k.GetAllFeeTier(ctx)

	for i, feeIndex := range feeIndices {
		sharesToRemove := sharesToRemoveList[i]
		tickIndex := tickIndices[i]

		// check that feeIndex is a valid index of the fee tier
		if feeIndex >= uint64(len(feeTiers)) {
			return sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "%d", feeIndex)
		}

		feeTier := feeTiers[feeIndex]

		pool, err := k.GetOrInitPool(ctx, pairId, tickIndex, feeTier)
		if err != nil {
			return err
		}

		sharesId := CreateSharesId(token0, token1, tickIndex, feeIndex)
		totalShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		if totalShares.LT(sharesToRemove) {
			return sdkerrors.Wrapf(types.ErrInsufficientShares, "%s does not have %s shares of type %s", callerAddr, sharesToRemove, sharesId)
		}

		outAmount0, outAmount1 := pool.Withdraw(sharesToRemove, totalShares)
		k.SavePool(ctx, pool)
		if sharesToRemove.GT(sdk.ZeroInt()) { // update shares accounting
			if err := k.BurnShares(ctx, callerAddr, sharesToRemove, sharesId); err != nil {
				return err
			}
		}

		totalReserve0ToRemove = totalReserve0ToRemove.Add(outAmount0)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(outAmount1)

		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(
			callerAddr.String(),
			receiverAddr.String(),
			token0,
			token1,
			fmt.Sprint(tickIndices[i]),
			fmt.Sprint(feeIndices[i]),
			pool.LowerTick0.Reserves.Add(outAmount0).String(),
			pool.UpperTick1.Reserves.Add(outAmount1).String(),
			pool.LowerTick0.Reserves.String(),
			pool.UpperTick1.Reserves.String(),
			sharesToRemove.String(),
		))
	}

	if totalReserve0ToRemove.GT(sdk.ZeroInt()) {
		coin0 := sdk.NewCoin(token0, totalReserve0ToRemove)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			receiverAddr,
			sdk.Coins{coin0},
		)
		if err != nil {
			return err
		}
	}

	// sends totalReserve1ToRemove to receiverAddr
	if totalReserve1ToRemove.GT(sdk.ZeroInt()) {
		coin1 := sdk.NewCoin(token1, totalReserve1ToRemove)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			receiverAddr,
			sdk.Coins{coin1},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// Handles core logic for the asset 0 to asset1 direction of MsgSwap; faciliates swapping amount0 for some amount of amount1, given a specified pair (token0, token1)
func (k Keeper) SwapCore(goCtx context.Context,
	tokenIn string,
	tokenOut string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	amountIn sdk.Int,
	minOut sdk.Int,
	limitPrice *sdk.Dec,
) (sdk.Int, sdk.Int, sdk.Int, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	cacheCtx, writeCache := ctx.CacheContext()
	pairId, err := CreatePairIdFromUnsorted(tokenIn, tokenOut)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), err
	}
	pair := types.NewDirectionalTradingPair(pairId, tokenIn, tokenOut)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), err
	}

	remainingIn := amountIn
	totalOut := sdk.ZeroInt()

	// verify that amount left is not zero and that there are additional valid ticks to check
	liqIter := NewLiquidityIterator(k, ctx, pair)
	defer liqIter.Close()
	for remainingIn.GT(sdk.ZeroInt()) {
		liq := liqIter.Next()
		if liq == nil {
			break
		}

		// break as soon as we iterated past limitPrice
		if limitPrice != nil && liq.Price().ToDec().LT(*limitPrice) {
			break

		}

		// price only gets worse as we iterate, so we can greedily abort
		// when the price is too low for minOut to be reached.
		idealOut := totalOut.Add(liq.Price().MulInt(remainingIn).TruncateInt())
		if idealOut.LT(minOut) {
			return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), nil
		}

		inAmount, outAmount := liq.Swap(remainingIn)

		remainingIn = remainingIn.Sub(inAmount)
		totalOut = totalOut.Add(outAmount)

		k.SaveLiquidity(cacheCtx, liq)
	}

	writeCache()
	totalIn := amountIn.Sub(remainingIn)
	return totalIn, totalOut, remainingIn, nil
}

func (k Keeper) ExecuteSwap(goCtx context.Context,
	tokenIn string,
	tokenOut string,
	amountIn sdk.Int,
	amountOut sdk.Int,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	minOut sdk.Int,
) (sdk.Coin, error) {

	if amountOut.LT(minOut) || amountOut.IsZero() {
		return sdk.Coin{}, types.ErrSlippageLimitReached
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	coinIn := sdk.NewCoin(tokenIn, amountIn)
	coinOut := sdk.NewCoin(tokenOut, amountOut)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
		return sdk.Coin{}, err
	}

	if amountOut.GT(sdk.ZeroInt()) {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return sdk.Coin{}, err
		}
	}
	ctx.EventManager().EmitEvent(types.CreateSwapEvent(callerAddr.String(), receiverAddr.String(),
		tokenIn, tokenOut, amountIn.String(), amountOut.String(), minOut.String(),
	))

	return coinOut, nil
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(
	goCtx context.Context,
	tokenIn string,
	tokenOut string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	amountIn sdk.Int,
	tickIndex int64,
	orderType types.LimitOrderType,
	goodTill *time.Time,
) (trancheKey string, err error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	token0, token1, err := SortTokens(tokenIn, tokenOut)
	if err != nil {
		return "", err
	}
	pairId := CreatePairId(token0, token1)
	var placeTranche types.LimitOrderTranche
	// TODO: JCP maybe move this somewhere else to simplify logic here
	switch orderType {
	case types.LimitOrderType_JUST_IN_TIME:
		placeTranche, err = k.InitPlaceTrancheWithGoodtill(ctx, pairId, tokenIn, tickIndex, types.JITGoodTillTime)
	case types.LimitOrderType_GOOD_TILl_DATE:
		placeTranche, err = k.InitPlaceTrancheWithGoodtill(ctx, pairId, tokenIn, tickIndex, *goodTill)
	default:
		placeTranche, err = k.GetOrInitPlaceTranche(ctx, pairId, tokenIn, tickIndex)
	}
	if err != nil {
		return "", err
	}

	trancheKey = placeTranche.TrancheKey
	trancheUser := k.GetOrInitLimitOrderTrancheUser(ctx, pairId, tickIndex, tokenIn, trancheKey, receiverAddr.String())
	amountLeft := amountIn

	// For everything except just-in-time (JIT) orders try to execute as a swap first

	var totalIn sdk.Int
	if !orderType.IsJIT() {
		var amountInSwap, amountOutSwap sdk.Int
		limitPrice := placeTranche.PriceMakerToTaker().ToDec()
		amountInSwap, amountOutSwap, amountLeft, err = k.SwapCore(
			goCtx,
			tokenIn,
			tokenOut,
			callerAddr,
			receiverAddr,
			amountIn,
			// TODO: JCP verify that this is maximally efficient vs precomputing minOut
			sdk.ZeroInt(),
			&limitPrice,
		)
		if err != nil {
			return "", err
		}
		totalIn = amountInSwap
		trancheUser.ReservesFromSwap = amountOutSwap
	}

	if !amountLeft.IsZero() && orderType.IsFoK() {
		return "", types.ErrFOKLimitOrderNotFilled
	}

	sharesIssued := sdk.ZeroInt()
	// FOR GTC, JIT & GoodTill try to place a maker limitOrder with remaining Amount
	if !amountLeft.IsZero() && (orderType.IsGTC() || orderType.IsJIT() || orderType.IsGoodTill()) {
		// TODO: JCP confirm that we never need this check. If GTC they should have already traded through cheaper liq so it doesn't matter
		// if JIT we just kinda assume they know what they are doing and we won't stop them.
		// if we do want this we can save a calculation by doing this first and skipping swap step in not BEL

		// if k.IsBehindEnemyLines(ctx, placeTranche.PairId, placeTranche.TokenIn, placeTranche.TickIndex) {
		// 	return "", types.ErrPlaceLimitOrderBehindPairLiquidity
		// }
		placeTranche.PlaceMakerLimitOrder(ctx, amountLeft)
		trancheUser.SharesOwned = trancheUser.SharesOwned.Add(amountLeft)

		if orderType.IsJIT() || orderType.IsGoodTill() {
			goodTillRecord := NewGoodTillRecord(pairId, tokenIn, tickIndex, trancheKey, *goodTill)
			k.SetGoodTillRecord(ctx, goodTillRecord)
		}
		k.SaveTranche(ctx, placeTranche)
		totalIn = totalIn.Add(amountLeft)
		sharesIssued = amountLeft
	}

	k.SaveTrancheUser(ctx, trancheUser)

	coin0 := sdk.NewCoin(tokenIn, totalIn)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0})
	if err != nil {
		return "", err
	}
	ctx.EventManager().EmitEvent(types.CreatePlaceLimitOrderEvent(
		callerAddr.String(),
		receiverAddr.String(),
		token0,
		token1,
		tokenIn,
		totalIn.String(),
		sharesIssued.String(),
		trancheKey,
	))

	return trancheKey, nil
}

// Handles MsgCancelLimitOrder, removing a specifed number of shares from a limit order and returning the respective amount in terms of the reserve to the user
func (k Keeper) CancelLimitOrderCore(
	goCtx context.Context,
	token0 string,
	token1 string,
	keyToken string,
	callerAddr sdk.AccAddress,
	tickIndex int64,
	trancheKey string,
) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId := CreatePairId(token0, token1)

	tranche, foundTranche := k.GetLimitOrderTranche(ctx, pairId, keyToken, tickIndex, trancheKey)
	trancheUser, foundTrancheUser := k.GetLimitOrderTrancheUser(ctx, pairId, tickIndex, keyToken, trancheKey, callerAddr.String())

	if !foundTranche || !foundTrancheUser {
		return types.ErrActiveLimitOrderNotFound
	}

	amountToCancel := tranche.RemoveTokenIn(trancheUser)
	trancheUser.SharesCancelled = trancheUser.SharesCancelled.Add(amountToCancel)
	userReserves := trancheUser.WithdrawSwapReserves()

	amountOut := amountToCancel.Add(userReserves)
	if amountOut.GT(sdk.ZeroInt()) {
		coinOut := sdk.NewCoin(keyToken, amountOut)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
		k.SaveTrancheUser(ctx, trancheUser)
		k.SaveTranche(ctx, *tranche)

	} else {
		return sdkerrors.Wrapf(types.ErrCancelEmptyLimitOrder, "%s", tranche.TrancheKey)
	}

	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(
		callerAddr.String(),
		token0,
		token1,
		keyToken,
		trancheKey,
		amountToCancel.String(),
	))

	return nil
}

// Handles MsgWithdrawFilledLimitOrder, calculates and sends filled liqudity from module to user for a limit order based on amount wished to receive.
func (k Keeper) WithdrawFilledLimitOrderCore(
	goCtx context.Context,
	token0 string,
	token1 string,
	tokenIn string,
	callerAddr sdk.AccAddress,
	tickIndex int64,
	trancheKey string,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := CreatePairId(token0, token1)
	_, tokenOut := GetInOutTokens(tokenIn, token0, token1)

	trancheUser, foundTrancheUser := k.GetLimitOrderTrancheUser(
		ctx,
		pairId,
		tickIndex,
		tokenIn,
		trancheKey,
		callerAddr.String(),
	)
	if !foundTrancheUser {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderTrancheNotFound, "%s", trancheKey)
	}

	tranche, wasFilled, foundTranche := k.FindLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, trancheKey)

	amountOutTokenOut := sdk.ZeroDec()
	remainingTokenIn := sdk.ZeroInt()
	// It's possible that a TrancheUser exists but tranche does not if LO was filled entirely through a swap
	if foundTranche {

		var amountOutTokenIn sdk.Int
		amountOutTokenIn, amountOutTokenOut = tranche.Withdraw(trancheUser)

		trancheUser.SharesWithdrawn = trancheUser.SharesWithdrawn.Add(amountOutTokenIn)
		// TODO: this is a bit of a messy pattern

		if wasFilled {
			//This is only relevant for JIT and GoodTill limit orders where the arhived
			remainingTokenIn = tranche.RemoveTokenIn(trancheUser)
			// TODO: switch to k.SaveFilledLimitOrderTranche and delete empty tranches
			k.SetFilledLimitOrderTranche(ctx, tranche)
		} else {
			k.SetLimitOrderTranche(ctx, tranche)
		}
	}

	userReserves := trancheUser.WithdrawSwapReserves()
	amountOutTokenOut = amountOutTokenOut.Add(userReserves.ToDec())

	k.SaveTrancheUser(ctx, trancheUser)

	if !amountOutTokenOut.IsZero() || !remainingTokenIn.IsZero() {
		coinOut := sdk.NewCoin(tokenOut, amountOutTokenOut.TruncateInt())
		coinInRefund := sdk.NewCoin(tokenIn, remainingTokenIn)
		// TODO: NewCoins does a lot of unnecesary work to sanitize and validate coins, all we really is to remove zero coins
		coins := sdk.NewCoins(coinOut, coinInRefund)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, coins); err != nil {
			return err
		}
	} else {
		return types.ErrWithdrawEmptyLimitOrder
	}

	ctx.EventManager().EmitEvent(types.WithdrawFilledLimitOrderEvent(
		callerAddr.String(),
		token0,
		token1,
		tokenIn,
		trancheKey,
		amountOutTokenOut.String(),
	))

	return nil
}
