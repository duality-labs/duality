package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NOTE: Currently we are using TruncateInt in multiple places for converting Decs back into sdk.Ints.
// This may create some accounting anomalies but seems preferable to other alternatives.
// See full ADR here: https://www.notion.so/dualityxyz/A-Modest-Proposal-For-Truncating-696a919d59254876a617f82fb9567895

// Handles core logic for MsgDeposit, checking and initializing data structures (tick, pair), calculating
// shares based on amount deposited, and sending funds to moduleAddress.
func (k Keeper) DepositCore(
	goCtx context.Context,
	msg *types.MsgDeposit,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	amounts0 []sdk.Int,
	amounts1 []sdk.Int,
) (amounts0Deposit []sdk.Int, amounts1Deposit []sdk.Int, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pair := k.GetOrInitPair(
		goCtx,
		token0,
		token1,
	)
	pairId := pair.PairId
	totalAmountReserve0 := sdk.ZeroInt()
	totalAmountReserve1 := sdk.ZeroInt()
	passedDeposit := 0
	amounts0Deposited := make([]sdk.Int, len(amounts0))
	amounts1Deposited := make([]sdk.Int, len(amounts1))
	for i := 0; i < len(amounts0); i++ {
		amounts0Deposited[i] = sdk.ZeroInt()
		amounts1Deposited[i] = sdk.ZeroInt()
	}

	feeTiers := k.GetAllFeeTier(ctx)

	for i, amount0 := range amounts0 {
		amount1 := amounts1[i]
		tickIndex := msg.TickIndexes[i]
		feeIndex := msg.FeeIndexes[i]
		autoswap := msg.Options[i].Autoswap

		// check that feeIndex is a valid index of the fee tier
		if feeIndex >= uint64(len(feeTiers)) {
			return nil, nil, sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", feeIndex)
		}
		fee := feeTiers[feeIndex].Fee
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee

		// behind enemy lines checks
		// TODO: Allow user to deposit "behind enemy lines"
		if amount0.GT(sdk.ZeroInt()) && pair.CurrentTick0To1 <= lowerTickIndex {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}
		// TODO: Allow user to deposit "behind enemy lines"
		if amount1.GT(sdk.ZeroInt()) && upperTickIndex <= pair.CurrentTick1To0 {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		// check for non-zero deposit
		if amount0.Equal(sdk.ZeroInt()) && amount1.Equal(sdk.ZeroInt()) {
			return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Cannot deposit 0,0")
		}

		lowerTick, err := k.GetOrInitTick(goCtx, pairId, lowerTickIndex)
		if err != nil {
			return nil, nil, sdkerrors.Wrapf(err, "Invalid lower tick (%d)", lowerTickIndex)
		}
		upperTick, err := k.GetOrInitTick(goCtx, pairId, upperTickIndex)
		if err != nil {
			return nil, nil, sdkerrors.Wrapf(err, "Invalid upper tick (%d)", lowerTickIndex)
		}

		sharesId := CreateSharesId(token0, token1, tickIndex, feeIndex)
		totalShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		pool := NewPool(
			&pair,
			tickIndex,
			feeIndex,
			&lowerTick,
			&upperTick,
		)

		oldReserve0 := pool.GetLowerReserve0()
		oldReserve1 := pool.GetUpperReserve1()

		inAmount0, inAmount1, outShares := pool.Deposit(amount0, amount1, totalShares, autoswap)
		pool.Save(goCtx, k)
		if outShares.GT(sdk.ZeroInt()) { // update shares accounting
			if err := k.MintShares(ctx, receiverAddr, outShares, sharesId); err != nil {
				return nil, nil, err
			}
		}

		if inAmount0.Equal(sdk.ZeroInt()) && inAmount1.Equal(sdk.ZeroInt()) {
			ctx.EventManager().EmitEvent(types.CreateDepositFailedEvent(
				msg.Creator,
				msg.Receiver,
				token0,
				token1,
				fmt.Sprint(tickIndex),
				fmt.Sprint(tickIndex),
				oldReserve0.String(),
				oldReserve1.String(),
				amount0.String(),
				amount1.String(),
			))
			continue
		}

		k.UpdateTickPointersPostAddToken0(goCtx, &pair, &lowerTick)
		k.UpdateTickPointersPostAddToken1(goCtx, &pair, &upperTick)
		k.SetTradingPair(ctx, pair)

		amounts0Deposited[i] = inAmount0
		amounts1Deposited[i] = inAmount1
		totalAmountReserve0 = totalAmountReserve0.Add(inAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(inAmount1)

		passedDeposit++

		ctx.EventManager().EmitEvent(types.CreateDepositEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes[i]),
			fmt.Sprint(msg.FeeIndexes[i]),
			pool.GetLowerReserve0().Sub(inAmount0).String(),
			pool.GetUpperReserve1().Sub(inAmount1).String(),
			pool.GetLowerReserve0().String(),
			pool.GetUpperReserve1().String(),
			outShares.String(),
		))
	}

	if passedDeposit == 0 {
		return nil, nil, sdkerrors.Wrapf(types.ErrAllDepositsFailed, "All deposits failed")
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

	_ = goCtx
	return amounts0Deposited, amounts1Deposited, nil
}

// Handles core logic for MsgWithdrawl; calculating and withdrawing reserve0,reserve1 from a specified tick given a specfied number of shares to remove.
// Calculates the amount of reserve0, reserve1 to withdraw based on the percetange of the desired number of shares to remove compared to the total number of shares at the given tick
func (k Keeper) WithdrawCore(goCtx context.Context, msg *types.MsgWithdrawl, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {

	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := CreatePairId(token0, token1)
	pair, found := k.GetTradingPair(ctx, pairId)
	if !found {
		return types.ErrValidPairNotFound
	}
	totalReserve0ToRemove := sdk.ZeroInt()
	totalReserve1ToRemove := sdk.ZeroInt()
	feeTiers := k.GetAllFeeTier(ctx)

	for i, feeIndex := range msg.FeeIndexes {
		sharesToRemove := msg.SharesToRemove[i]
		tickIndex := msg.TickIndexes[i]

		// check that feeIndex is a valid index of the fee tier
		if feeIndex >= uint64(len(feeTiers)) {
			return sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", feeIndex)
		}

		fee := feeTiers[feeIndex].Fee
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee
		sharesId := CreateSharesId(token0, token1, tickIndex, feeIndex)
		totalShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		if totalShares.LT(sharesToRemove) {
			return sdkerrors.Wrapf(types.ErrNotEnoughShares, "Insufficient shares %s", sharesId)
		}

		lowerTick, lowerTickFound := k.GetTick(ctx, pairId, lowerTickIndex)
		upperTick, upperTickFound := k.GetTick(ctx, pairId, upperTickIndex)
		if !lowerTickFound || !upperTickFound {
			return types.ErrValidTickNotFound
		}

		pool := NewPool(
			&pair,
			tickIndex,
			feeIndex,
			&lowerTick,
			&upperTick,
		)
		outAmount0, outAmount1, err := pool.Withdraw(sharesToRemove, totalShares)
		if err != nil {
			return err
		}
		pool.Save(goCtx, k)
		if sharesToRemove.GT(sdk.ZeroInt()) { // update shares accounting
			if err := k.BurnShares(ctx, callerAddr, sharesToRemove, sharesId); err != nil {
				return err
			}
		}

		totalReserve0ToRemove = totalReserve0ToRemove.Add(outAmount0)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(outAmount1)

		if outAmount0.GT(sdk.ZeroInt()) {
			k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &lowerTick)
		}

		if outAmount1.GT(sdk.ZeroInt()) {
			k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &upperTick)
		}

		k.SetTradingPair(ctx, pair)

		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes[i]),
			fmt.Sprint(msg.FeeIndexes[i]),
			pool.LowerTick0.TickData.Reserve0[feeIndex].Add(outAmount0).String(),
			pool.UpperTick1.TickData.Reserve1[feeIndex].Add(outAmount1).String(),
			pool.LowerTick0.TickData.Reserve0[feeIndex].String(),
			pool.UpperTick1.TickData.Reserve1[feeIndex].String(),
			sharesToRemove.String(),
		))
	}
	k.SetTradingPair(ctx, pair)
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

	// sends totalReserve1ToRemove to msg.Receiver
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
	msg *types.MsgSwap,
	tokenIn string,
	tokenOut string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) (sdk.Coin, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	feeTiers := k.GetAllFeeTier(ctx)

	pair, err := k.GetDirectionalTradingPair(ctx, tokenIn, tokenOut)
	if err != nil {
		return sdk.Coin{}, err
	}

	remainingIn := msg.AmountIn
	totalOut := sdk.ZeroInt()

	// verify that amount left is not zero and that there are additional valid ticks to check
	// for !amount_left.Equal(sdk.ZeroInt()) && pair.TokenPair.CurrentTick0To1 <= pair.MaxTick {
	liqIter := NewLiquidityIterator(k, goCtx, pair, feeTiers)
	for liqIter.HasNext() && remainingIn.GT(sdk.ZeroInt()) {
		liq := liqIter.Next()

		// break as soon as we iterated past tickLimit
		if liq.Price().LT(msg.LimitPrice) {
			break
		}

		// price only gets worse as we iterate, so we can greedily abort
		// when the price is too low for minOut to be reached.
		idealOut := totalOut.Add(remainingIn.ToDec().Mul(liq.Price()).TruncateInt())
		if idealOut.LT(msg.MinOut) {
			return sdk.Coin{}, types.ErrNotEnoughLiquidity
		}

		inAmount, outAmount, initedTick, deinitedTick := liq.Swap(remainingIn)

		remainingIn = remainingIn.Sub(inAmount)
		totalOut = totalOut.Add(outAmount)
		// Saving all the time
		liq.Save(goCtx, k)

		if initedTick != nil {
			k.InitLiquidity(&pair, initedTick.TickIndex)
		}
		if k.ShouldDeinit(ctx, deinitedTick, pair) {
			k.DeinitLiquidity(goCtx, &pair, deinitedTick.TickIndex)
		}

	}
	k.SetTradingPair(ctx, pair.TradingPair)

	if totalOut.LT(msg.MinOut) || msg.AmountIn.Equal(remainingIn) {
		return sdk.Coin{}, types.ErrNotEnoughLiquidity
	}

	// TODO: Move this to a separate ExecuteSwap function. Ditto for all other *Core fns
	amountToDeposit := msg.AmountIn.Sub(remainingIn)
	coinIn := sdk.NewCoin(tokenIn, amountToDeposit)
	coinOut := sdk.NewCoin(tokenOut, totalOut)

	if amountToDeposit.GT(sdk.ZeroInt()) {

		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
			return sdk.Coin{}, err
		}
	} else {
		return sdk.Coin{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
	}

	if totalOut.GT(sdk.ZeroInt()) {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return sdk.Coin{}, err
		}
	}
	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		tokenIn, tokenOut, msg.AmountIn.String(), totalOut.String(), msg.MinOut.String(),
	))

	return coinOut, nil
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder, tokenIn string, tokenOut string, callerAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := SortTokens(ctx, tokenIn, tokenOut)
	if err != nil {
		return err
	}
	rawPair := k.GetOrInitPair(goCtx, token0, token1)
	pair := types.NewDirectionalTradingPair(rawPair, tokenIn, tokenOut)
	pairId := pair.PairId
	tick, err := k.GetOrInitTick(goCtx, pair.PairId, msg.TickIndex)
	if err != nil {
		return err
	}

	tickIndex := msg.TickIndex
	receiver := msg.Receiver

	var fillTrancheIndex *uint64
	var placeTrancheIndex *uint64

	if msg.TokenIn == token0 {
		if msg.TickIndex > pair.CurrentTick0To1 {
			return types.ErrPlaceLimitOrderBehindPairLiquidity
		}
		fillTrancheIndex = &tick.LimitOrderTranche0To1.FillTrancheIndex
		placeTrancheIndex = &tick.LimitOrderTranche0To1.PlaceTrancheIndex
	} else {
		if msg.TickIndex < pair.CurrentTick1To0 {
			return types.ErrPlaceLimitOrderBehindPairLiquidity
		}
		fillTrancheIndex = &tick.LimitOrderTranche1To0.FillTrancheIndex
		placeTrancheIndex = &tick.LimitOrderTranche1To0.PlaceTrancheIndex
	}

	tranche := k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, *placeTrancheIndex)
	trancheUser := k.GetOrInitLimitOrderTrancheUser(goCtx, pairId, tickIndex, tokenIn, *placeTrancheIndex, receiver)
	if tranche.ReservesTokenIn.LT(tranche.TotalTokenIn) {
		*placeTrancheIndex++
		k.SetTick(ctx, pairId, tick)
		tranche = k.GetOrInitLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, *placeTrancheIndex)
		trancheUser = k.GetOrInitLimitOrderTrancheUser(goCtx, pairId, tickIndex, tokenIn, *placeTrancheIndex, receiver)
	}
	tranche.ReservesTokenIn = tranche.ReservesTokenIn.Add(msg.AmountIn)
	tranche.TotalTokenIn = tranche.TotalTokenIn.Add(msg.AmountIn)
	trancheUser.SharesOwned = trancheUser.SharesOwned.Add(msg.AmountIn)

	k.SetLimitOrderTrancheUser(ctx, trancheUser)
	k.SetLimitOrderTranche(ctx, tranche)

	if msg.AmountIn.GT(sdk.ZeroInt()) {
		k.InitLiquidity(&pair, tick.TickIndex)
		k.SetTradingPair(ctx, pair.TradingPair)

		coin0 := sdk.NewCoin(tokenIn, msg.AmountIn)
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0})
		if err != nil {
			return err
		}
	}

	ctx.EventManager().EmitEvent(types.CreatePlaceLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), msg.AmountIn.String(), strconv.Itoa(int(*fillTrancheIndex)),
	))

	return nil
}

// Handles MsgCancelLimitOrder, removing a specifed number of shares from a limit order and returning the respective amount in terms of the reserve to the user
func (k Keeper) CancelLimitOrderCore(goCtx context.Context, msg *types.MsgCancelLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId := CreatePairId(token0, token1)

	tick, tickFound := k.GetTick(ctx, pairId, msg.TickIndex)
	if !tickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	trancheUser, found := k.GetLimitOrderTrancheUser(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	if !found {
		return types.ErrValidLimitOrderMapsNotFound
	}
	// checks that the user has some number of limit order shares wished to withdraw
	if trancheUser.SharesOwned.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	tranche, found := k.GetLimitOrderTranche(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
	if !found {
		return types.ErrValidLimitOrderMapsNotFound
	}

	var priceLimitOutToIn sdk.Dec
	var limitTrancheIndexes *types.LimitTrancheIndexes
	if msg.KeyToken == token0 {
		priceLimitOutToIn = sdk.OneDec().Quo(*tick.Price0To1)
		limitTrancheIndexes = tick.LimitOrderTranche0To1
	} else {
		priceLimitOutToIn = *tick.Price0To1
		limitTrancheIndexes = tick.LimitOrderTranche1To0
	}
	totalTokenInDec := sdk.NewDecFromInt(tranche.TotalTokenIn)
	totalTokenOutDec := sdk.NewDecFromInt(tranche.TotalTokenOut)
	filledAmount := priceLimitOutToIn.Mul(totalTokenOutDec)
	ratioNotFilled := totalTokenInDec.Sub(filledAmount).Quo(totalTokenInDec)
	amountToCancel := trancheUser.SharesOwned.ToDec().Mul(ratioNotFilled).TruncateInt()

	trancheUser.SharesCancelled = trancheUser.SharesCancelled.Add(amountToCancel)
	k.SetLimitOrderTrancheUser(ctx, trancheUser)
	tranche.ReservesTokenIn = tranche.ReservesTokenIn.Sub(amountToCancel)
	k.SetLimitOrderTranche(ctx, tranche)

	for tranche.ReservesTokenIn.Equal(sdk.ZeroInt()) && limitTrancheIndexes.FillTrancheIndex < limitTrancheIndexes.PlaceTrancheIndex {
		limitTrancheIndexes.FillTrancheIndex++
		tranche, found = k.GetLimitOrderTranche(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
		if !found {
			return types.ErrValidLimitOrderMapsNotFound
		}
	}
	k.SetTick(ctx, pairId, tick)

	if amountToCancel.GT(sdk.ZeroInt()) {
		// See top NOTE on rounding
		coinOut := sdk.NewCoin(msg.KeyToken, amountToCancel)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot cancel liqudity from this limit order at this time")
	}

	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountToCancel.String(),
	))

	tokenOut, tokenIn := GetInOutTokens(msg.KeyToken, token0, token1)
	pair, _ := k.GetDirectionalTradingPair(ctx, tokenIn, tokenOut)
	if k.ShouldDeinit(ctx, &tick, pair) {
		k.DeinitLiquidity(goCtx, &pair, tick.TickIndex)
		k.SetTradingPair(ctx, pair.TradingPair)
	}

	return nil
}

// Handles MsgWithdrawFilledLimitOrder, calculates and sends filled liqudity from module to user for a limit order based on amount wished to receive.
func (k Keeper) WithdrawFilledLimitOrderCore(
	goCtx context.Context,
	msg *types.MsgWithdrawFilledLimitOrder,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := CreatePairId(token0, token1)

	orderTokenIn := msg.KeyToken
	var orderTokenOut string
	if msg.KeyToken == token0 {
		orderTokenOut = token1
	} else {
		orderTokenOut = token0
	}
	trancheIndex := msg.Key
	tickIndex := msg.TickIndex

	tranche, found := k.GetLimitOrderTranche(ctx, pairId, tickIndex, orderTokenIn, trancheIndex)
	if !found {
		return types.ErrValidLimitOrderMapsNotFound
	}

	trancheUser, found := k.GetLimitOrderTrancheUser(
		ctx,
		pairId,
		tickIndex,
		orderTokenIn,
		trancheIndex,
		msg.Creator,
	)
	if !found {
		return types.ErrValidLimitOrderMapsNotFound
	}
	sharesToWithdraw := trancheUser.SharesOwned.Sub(trancheUser.SharesCancelled)

	// checks that the user has some number of limit order shares wished to withdraw
	if sharesToWithdraw.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	tick, found := k.GetTick(ctx, pairId, msg.TickIndex)
	if !found {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	var priceLimitInToOut sdk.Dec
	var priceLimitOutToIn sdk.Dec
	if orderTokenIn == token0 {
		priceLimitInToOut = *tick.Price0To1
	} else {
		priceLimitInToOut = sdk.OneDec().Quo(*tick.Price0To1)
	}
	priceLimitOutToIn = sdk.OneDec().Quo(priceLimitInToOut)

	reservesTokenOutDec := sdk.NewDecFromInt(tranche.ReservesTokenOut)
	amountFilled := priceLimitOutToIn.MulInt(tranche.TotalTokenOut)
	ratioFilled := amountFilled.QuoInt(tranche.TotalTokenIn)
	maxAllowedToWithdraw := sdk.MinInt(
		ratioFilled.MulInt(trancheUser.SharesOwned).TruncateInt(), // cannot withdraw more than what's been filled
		sharesToWithdraw,
	)
	amountOutTokenIn := maxAllowedToWithdraw.Sub(trancheUser.SharesWithdrawn)

	amountOutTokenOut := priceLimitInToOut.MulInt(amountOutTokenIn)

	trancheUser.SharesWithdrawn = maxAllowedToWithdraw
	k.SetLimitOrderTrancheUser(ctx, trancheUser)

	// See top NOTE on rounding
	tranche.ReservesTokenOut = reservesTokenOutDec.Sub(amountOutTokenOut).TruncateInt()
	k.SetLimitOrderTranche(ctx, tranche)

	if amountOutTokenOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(orderTokenOut, amountOutTokenOut.TruncateInt())
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot withdraw additional liqudity from this limit order at this time")
	}

	ctx.EventManager().EmitEvent(types.WithdrawFilledLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountOutTokenOut.String(),
	))

	return nil
}
