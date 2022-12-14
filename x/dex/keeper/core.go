package keeper

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NOTE: Currently we are using TruncateInt in multiple places for converting Decs back into sdk.Ints.
// This may create some accounting anomalies but seems preferable to other alternatives.
// See full ADR here: https://www.notion.so/dualityxyz/A-Modest-Proposal-For-Truncating-696a919d59254876a617f82fb9567895

// Handles core logic for MsgDeposit, checking and initializing data structures (tick, pair), calculating
// shares based on amount deposited, and sending funds to moduleAddress
func (k Keeper) DepositCore(
	goCtx context.Context,
	msg *types.MsgDeposit,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
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

	FeeTier := k.GetAllFeeTier(ctx)

	for i, amount0 := range amounts0 {
		amount1 := amounts1[i]
		tickIndex := msg.TickIndexes[i]
		feeIndex := msg.FeeIndexes[i]
		fee := FeeTier[feeIndex].Fee
		curTick0to1 := pair.CurrentTick0To1
		curTick1to0 := pair.CurrentTick1To0
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee

		// TODO: Allow user to deposit "behind enemy lines"
		if amount0.GT(sdk.ZeroInt()) && curTick0to1 <= lowerTickIndex {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		// TODO: Allow user to deposit "behind enemy lines"
		if amount1.GT(sdk.ZeroInt()) && upperTickIndex <= curTick1to0 {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		lowerTick := k.GetOrInitTick(goCtx, pairId, lowerTickIndex)
		upperTick := k.GetOrInitTick(goCtx, pairId, upperTickIndex)

		sharesId := k.CreateSharesId(token0, token1, tickIndex, fee)
		totalShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		pool, err := NewPool(
			pairId,
			tickIndex,
			feeIndex,
			fee,
			&lowerTick,
			&upperTick,
		)
		if err != nil {
			return amounts0Deposited, amounts1Deposited, err
		}
		oldReserve0 := pool.GetLowerReserve0()
		oldReserve1 := pool.GetUpperReserve1()

		inAmount0, inAmount1, outShares := pool.Deposit(amount0, amount1, totalShares)
		pool.Save(goCtx, k)
		if outShares.GT(sdk.ZeroInt()) { // update shares accounting
			if err := k.MintShares(ctx, callerAddr, outShares, sharesId); err != nil {
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

		k.SetTradingPair(ctx, pair)

		k.UpdateTickPointersPostAddToken0(goCtx, &pair, &lowerTick)
		k.UpdateTickPointersPostAddToken1(goCtx, &pair, &upperTick)

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
	pairId := k.CreatePairId(token0, token1)
	pair, found := k.GetTradingPair(ctx, pairId)
	if !found {
		return types.ErrValidPairNotFound
	}
	totalReserve0ToRemove := sdk.ZeroInt()
	totalReserve1ToRemove := sdk.ZeroInt()

	for i, feeIndex := range msg.FeeIndexes {
		sharesToRemove := msg.SharesToRemove[i]
		tickIndex := msg.TickIndexes[i]

		feeValue, found := k.GetFeeTier(ctx, feeIndex)
		if !found {
			return types.ErrValidFeeIndexNotFound
		}
		fee := feeValue.Fee
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee
		lowerTick, lowerTickFound := k.GetTick(ctx, pairId, lowerTickIndex)
		upperTick, upperTickFound := k.GetTick(ctx, pairId, upperTickIndex)
		if !lowerTickFound || !upperTickFound {
			return types.ErrValidTickNotFound
		}

		sharesId := k.CreateSharesId(token0, token1, tickIndex, fee)
		totalShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		pool, err := NewPool(
			pairId,
			tickIndex,
			feeIndex,
			fee,
			&lowerTick,
			&upperTick,
		)
		if err != nil {
			return err
		}
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
func (k Keeper) Swap0to1(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Int, sdk.Int, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	feeSize := k.GetFeeTierCount(ctx)
	FeeTier := k.GetAllFeeTier(ctx)
	pair, pairFound := k.GetTradingPair(ctx, pairId)
	if !pairFound {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}
	if pair.CurrentTick0To1 == math.MaxInt64 {
		return sdk.ZeroInt(), sdk.ZeroInt(), types.ErrNotEnoughLiquidity
	}

	remainingInAmount0 := msg.AmountIn
	totalOutAmount1 := sdk.ZeroInt()

	// verify that amount left is not zero and that there are additional valid ticks to check
	// for !amount_left.Equal(sdk.ZeroInt()) && pair.TokenPair.CurrentTick0To1 <= pair.MaxTick {
	for remainingInAmount0.GT(sdk.ZeroInt()) && pair.CurrentTick0To1 <= pair.MaxTick {
		Current1Data, Current1Found := k.GetTick(ctx, pairId, pair.CurrentTick0To1)
		if !Current1Found {
			pair.CurrentTick0To1++
			continue
		}

		var i uint64 = 0

		for i < feeSize && remainingInAmount0.GT(sdk.ZeroInt()) {
			fee := FeeTier[i].Fee
			Current0Data, found := k.GetTick(ctx, pairId, pair.CurrentTick0To1-2*fee)
			if !found {
				i++
				continue
			}

			pool, err := NewPool(
				pairId,
				pair.CurrentTick0To1-fee,
				i,
				fee,
				&Current0Data,
				&Current1Data,
			)
			if err != nil {
				return sdk.ZeroInt(), sdk.ZeroInt(), err
			}
			inAmount0, outAmount1 := pool.Swap0To1(remainingInAmount0)
			remainingInAmount0 = remainingInAmount0.Sub(inAmount0)
			totalOutAmount1 = totalOutAmount1.Add(outAmount1)
			pool.Save(goCtx, k)
			k.UpdateTickPointersPostAddToken0(goCtx, &pair, &Current0Data)
			i++
		}

		if i == feeSize && remainingInAmount0.GT(sdk.ZeroInt()) {
			var err error
			var remainingInAmount0Dec sdk.Dec
			remainingInAmount0Dec, totalOutAmount1, err = k.SwapLimitOrder0to1(
				goCtx,
				pairId,
				token1,
				totalOutAmount1,
				remainingInAmount0.ToDec(),
				pair.CurrentTick0To1,
			)
			remainingInAmount0 = remainingInAmount0Dec.TruncateInt()

			if err != nil {
				return sdk.ZeroInt(), sdk.ZeroInt(), err
			}
		}
		k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &Current1Data)
	}

	k.SetTradingPair(ctx, pair)

	// Check to see if amount_out meets the threshold of minOut
	if totalOutAmount1.LT(msg.MinOut) {
		return sdk.ZeroInt(), sdk.ZeroInt(), types.ErrNotEnoughLiquidity
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), totalOutAmount1.String(), msg.MinOut.String(),
	))

	// Returns amount_out to keeper/msg.server: Swap
	// @Dev token transfers happen in keeper/msg.server: Swap
	return totalOutAmount1, remainingInAmount0, nil
}

// Handles core logic for the asset 1 to asset 0 direction of MsgSwap; faciliates swapping amount1 for some amount of amount0, given a specified pair (token0, token1)
func (k Keeper) Swap1to0(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Int, sdk.Int, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	feeSize := k.GetFeeTierCount(ctx)
	FeeTier := k.GetAllFeeTier(ctx)
	pair, found := k.GetTradingPair(ctx, pairId)
	if !found {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}
	if pair.CurrentTick1To0 == math.MinInt64 {
		return sdk.ZeroInt(), sdk.ZeroInt(), types.ErrNotEnoughLiquidity
	}

	remainingInAmount1 := msg.AmountIn
	totalOutAmount0 := sdk.ZeroInt()
	for remainingInAmount1.GT(sdk.ZeroInt()) && pair.CurrentTick1To0 >= pair.MinTick {

		Current0Data, Current0Found := k.GetTick(ctx, pairId, pair.CurrentTick1To0)
		if !Current0Found {
			pair.CurrentTick1To0 = pair.CurrentTick1To0 - 1
			continue
		}

		var i uint64 = 0

		for i < feeSize && remainingInAmount1.GT(sdk.ZeroInt()) {
			fee := FeeTier[i].Fee
			Current1Data, found := k.GetTick(ctx, pairId, pair.CurrentTick1To0+2*fee)
			if !found {
				i++
				continue
			}

			pool, err := NewPool(
				pairId,
				pair.CurrentTick1To0+fee,
				i,
				fee,
				&Current0Data,
				&Current1Data,
			)
			if err != nil {
				return sdk.ZeroInt(), sdk.ZeroInt(), err
			}
			inAmount1, outAmount0 := pool.Swap1To0(remainingInAmount1)
			remainingInAmount1 = remainingInAmount1.Sub(inAmount1)
			totalOutAmount0 = totalOutAmount0.Add(outAmount0)
			pool.Save(goCtx, k)
			k.UpdateTickPointersPostAddToken1(goCtx, &pair, &Current1Data)
			i++
		}

		k.SetTick(ctx, pairId, Current0Data)

		if i == feeSize && remainingInAmount1.GT(sdk.ZeroInt()) {
			var err error
			var remainingInAmount1Dec sdk.Dec
			remainingInAmount1Dec, totalOutAmount0, err = k.SwapLimitOrder1to0(
				goCtx,
				pairId,
				token0,
				totalOutAmount0,
				remainingInAmount1.ToDec(),
				pair.CurrentTick1To0,
			)
			remainingInAmount1 = remainingInAmount1Dec.TruncateInt()

			if err != nil {
				return sdk.ZeroInt(), sdk.ZeroInt(), err
			}
		}
		k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &Current0Data)
	}

	k.SetTradingPair(ctx, pair)

	if totalOutAmount0.LT(msg.MinOut) {
		return sdk.ZeroInt(), sdk.ZeroInt(), types.ErrNotEnoughLiquidity
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), totalOutAmount0.String(), msg.MinOut.String(),
	))

	return totalOutAmount0, remainingInAmount1, nil
}

// Handles swapping asset 0 for asset 1 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error
func (k Keeper) SwapLimitOrder0to1(
	goCtx context.Context,
	pairId string,
	tokenOut string,
	amountOut sdk.Int,
	amountRemainingTokenIn sdk.Dec,
	tickIndex int64,
) (newAmountRemainingTokenIn sdk.Dec, newAmountOut sdk.Int, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	priceInToOut, err := CalcPrice0To1(tickIndex)
	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroInt(), err
	}
	priceOutToIn := sdk.OneDec().Quo(priceInToOut)

	tick, tickFound := k.GetTick(ctx, pairId, tickIndex)
	if !tickFound {
		return amountRemainingTokenIn, amountOut, nil
	}

	fillTranche := &tick.LimitOrderTranche1To0.FillTrancheIndex
	placeTranche := &tick.LimitOrderTranche1To0.PlaceTrancheIndex

	for amountRemainingTokenIn.GT(sdk.ZeroDec()) && *fillTranche < *placeTranche {
		amountRemainingTokenIn, amountOut, err = k.SwapLimitOrderTranche(
			goCtx,
			pairId,
			tokenOut,
			amountOut,
			amountRemainingTokenIn,
			tickIndex,
			*fillTranche,
			priceInToOut,
			priceOutToIn,
		)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroInt(), err
		}
		if !k.TickTrancheHasToken0(ctx, &tick, *fillTranche) {
			*fillTranche++
			k.SetTick(ctx, pairId, tick)
		}
	}

	if amountRemainingTokenIn.GT(sdk.ZeroDec()) {
		amountRemainingTokenIn, amountOut, err = k.SwapLimitOrderTranche(
			goCtx,
			pairId,
			tokenOut,
			amountOut,
			amountRemainingTokenIn,
			tickIndex,
			*fillTranche,
			priceInToOut,
			priceOutToIn,
		)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroInt(), err
		}
	}

	pair, _ := k.GetTradingPair(ctx, pairId)
	k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)

	return amountRemainingTokenIn, amountOut, nil
}

// Handles swapping asset 1 for asset 0 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error
func (k Keeper) SwapLimitOrder1to0(
	goCtx context.Context,
	pairId string,
	tokenOut string,
	amountOut sdk.Int,
	amountRemainingTokenIn sdk.Dec,
	tickIndex int64,
) (newAmountRemainingTokenIn sdk.Dec, newAmountOut sdk.Int, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	priceInToOut, err := CalcPrice1To0(tickIndex)
	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroInt(), err
	}
	priceOutToIn := sdk.OneDec().Quo(priceInToOut)

	tick, tickFound := k.GetTick(ctx, pairId, tickIndex)
	if !tickFound {
		return amountRemainingTokenIn, amountOut, nil
	}

	fillTranche := &tick.LimitOrderTranche0To1.FillTrancheIndex
	placeTranche := &tick.LimitOrderTranche0To1.PlaceTrancheIndex

	for amountRemainingTokenIn.GT(sdk.ZeroDec()) && *fillTranche < *placeTranche {
		amountRemainingTokenIn, amountOut, err = k.SwapLimitOrderTranche(
			goCtx,
			pairId,
			tokenOut,
			amountOut,
			amountRemainingTokenIn,
			tickIndex,
			*fillTranche,
			priceInToOut,
			priceOutToIn,
		)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroInt(), err
		}
		if !k.TickTrancheHasToken1(ctx, &tick, *fillTranche) {
			*fillTranche++
			k.SetTick(ctx, pairId, tick)
		}
	}

	if amountRemainingTokenIn.GT(sdk.ZeroDec()) {
		amountRemainingTokenIn, amountOut, err = k.SwapLimitOrderTranche(
			goCtx,
			pairId,
			tokenOut,
			amountOut,
			amountRemainingTokenIn,
			tickIndex,
			*fillTranche,
			priceInToOut,
			priceOutToIn,
		)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroInt(), err
		}
	}

	pair, _ := k.GetTradingPair(ctx, pairId)
	k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)

	return amountRemainingTokenIn, amountOut, nil
}

func (k Keeper) SwapLimitOrderTranche(
	goCtx context.Context,
	pairId string,
	tokenOut string,
	amountOut sdk.Int,
	amountRemainingTokenIn sdk.Dec,
	tickIndex int64,
	trancheIndex uint64,
	priceInToOut sdk.Dec,
	priceOutToIn sdk.Dec,
) (newAmountRemainingTokenIn sdk.Dec, newAmountOut sdk.Int, error error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tranche, found := k.GetLimitOrderTranche(ctx, pairId, tickIndex, tokenOut, trancheIndex)
	if !found {
		return amountRemainingTokenIn, amountOut, nil
	}
	reservesTokenOut := &tranche.ReservesTokenIn
	fillTokenIn := &tranche.ReservesTokenOut
	totalTokenIn := &tranche.TotalTokenOut
	// See top NOTE on rounding
	amountFilledTokenOut := priceInToOut.Mul(amountRemainingTokenIn).TruncateInt()

	if reservesTokenOut.LTE(amountFilledTokenOut) {
		amountOut = amountOut.Add(*reservesTokenOut)
		amountFilledTokenIn := priceOutToIn.MulInt(*reservesTokenOut)
		amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
		*reservesTokenOut = sdk.ZeroInt()
		*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn.TruncateInt())
		*totalTokenIn = totalTokenIn.Add(amountFilledTokenIn.TruncateInt())
	} else {
		amountOut = amountOut.Add(amountFilledTokenOut)
		*fillTokenIn = fillTokenIn.Add(amountRemainingTokenIn.TruncateInt())
		*totalTokenIn = totalTokenIn.Add(amountRemainingTokenIn.TruncateInt())
		*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
		amountRemainingTokenIn = sdk.ZeroDec()
	}
	k.SetLimitOrderTranche(ctx, tranche)

	return amountRemainingTokenIn, amountOut, nil
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pair := k.GetOrInitPair(goCtx, token0, token1)
	pairId := pair.PairId
	tick := k.GetOrInitTick(goCtx, pair.PairId, msg.TickIndex)

	fillTrancheIndex, placeTrancheIndex, err := GetFillAndPlaceIndexes(msg.TokenIn, token0, &pair, &tick)
	if err != nil {
		return err
	}

	tranche, err := k.GetCurrentPlaceTranche(ctx, msg.TokenIn, token0, pairId, &tick, placeTrancheIndex)
	if err != nil {
		return err
	}
	trancheUser := k.GetUserShares(goCtx, tranche, msg.Receiver)

	tranche.PlaceLimitOrder(msg.AmountIn, &trancheUser)

	tranche.Save(ctx, k)
	trancheUser.Save(ctx, k)

	k.SetTradingPair(ctx, pair)

	if msg.TokenIn == token0 {
		k.UpdateTickPointersPostAddToken0(goCtx, &pair, &tick)
	} else if msg.TokenIn == token1 {
		k.UpdateTickPointersPostAddToken1(goCtx, &pair, &tick)
	}

	if msg.AmountIn.GT(sdk.ZeroInt()) {
		coin0 := sdk.NewCoin(msg.TokenIn, msg.AmountIn)
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

	attemptedSharesOut := msg.SharesOut
	pairId := k.CreatePairId(token0, token1)

	tick, tickFound := k.GetTick(ctx, pairId, msg.TickIndex)
	if !tickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	// TODO: currently, msg.KeyToken == tokenOut in CancelLO but msg.KeyToken == tokenIn in PlaceLO. This should be standardized...
	tranche, err := NewLimitOrderTrancheFromKeeper(ctx, k, msg.KeyToken, token0, pairId, tick.TickIndex, msg.Key)
	if err != nil {
		return err
	}
	trancheUser := k.GetUserShares(goCtx, tranche, msg.Receiver)

	err = tranche.CancelLimitOrder(msg.SharesOut, &trancheUser)
	if err != nil {
		return err
	}

	tranche.Save(ctx, k)
	trancheUser.Save(ctx, k)

	if attemptedSharesOut.GT(sdk.ZeroInt()) {
		// See top NOTE on rounding
		coinOut := sdk.NewCoin(msg.KeyToken, attemptedSharesOut)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot cancel liqudity from this limit order at this time")
	}

	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), attemptedSharesOut.String(),
	))

	pair, _ := k.GetTradingPair(ctx, pairId)
	if msg.KeyToken == token0 {
		k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)
	} else {
		k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &tick)
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
	pairId := k.CreatePairId(token0, token1)

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

	tick, found := k.GetTick(ctx, pairId, msg.TickIndex)
	if !found {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	var priceLimitInToOut sdk.Dec
	var priceLimitOutToIn sdk.Dec
	var err error
	if orderTokenIn == token0 {
		priceLimitInToOut, err = CalcPrice0To1(tick.TickIndex)
		if err != nil {
			return err
		}
	} else {
		priceLimitInToOut, err = CalcPrice1To0(tick.TickIndex)
		if err != nil {
			return err
		}
	}
	priceLimitOutToIn = sdk.OneDec().Quo(priceLimitInToOut)

	// TODO: move to lo object
	reservesTokenOutDec := sdk.NewDecFromInt(tranche.ReservesTokenOut)
	amountFilled := priceLimitOutToIn.MulInt(tranche.TotalTokenOut)
	ratioFilled := amountFilled.QuoInt(tranche.TotalTokenIn)
	maxAllowedToWithdraw := sdk.MinInt(
		ratioFilled.MulInt(trancheUser.SharesOwned).TruncateInt(), // cannot withdraw more than what's been filled
		trancheUser.SharesOwned.Sub(trancheUser.SharesCancelled),  // cannot withdraw more than what you own
	)
	amountOutTokenIn := maxAllowedToWithdraw.Sub(trancheUser.SharesWithdrawn)

	amountOutTokenOut := priceLimitInToOut.MulInt(amountOutTokenIn)

	trancheUser.SharesWithdrawn = maxAllowedToWithdraw
	k.SetLimitOrderTrancheUser(ctx, trancheUser)

	// See top NOTE on rounding
	tranche.ReservesTokenOut = reservesTokenOutDec.Sub(amountOutTokenOut).TruncateInt()
	k.SetLimitOrderTranche(ctx, tranche)
	// end TODO

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
