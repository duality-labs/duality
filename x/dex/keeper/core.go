package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
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
		tickIndex := msg.TickIndexes[i]
		feeIndex := msg.FeeIndexes[i]
		autoswap := msg.Options[i].Autoswap

		// check that feeIndex is a valid index of the fee tier
		if feeIndex >= uint64(len(feeTiers)) {
			return nil, nil, sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "%d", feeIndex)
		}
		fee := feeTiers[feeIndex].Fee
		lowerTickIndex := tickIndex - int64(fee)
		upperTickIndex := tickIndex + int64(fee)

		// behind enemy lines checks
		// TODO: Allow user to deposit "behind enemy lines"
		if amount0.GT(sdk.ZeroInt()) && k.IsBehindEnemyLines(ctx, pairId, pairId.Token0, lowerTickIndex) {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}
		// TODO: Allow user to deposit "behind enemy lines"
		if amount1.GT(sdk.ZeroInt()) && k.IsBehindEnemyLines(ctx, pairId, pairId.Token1, upperTickIndex) {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		// check for non-zero deposit
		if amount0.Equal(sdk.ZeroInt()) && amount1.Equal(sdk.ZeroInt()) {
			return nil, nil, types.ErrZeroDeposit
		}

		sharesId := CreateSharesId(token0, token1, tickIndex, feeIndex)
		totalShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		pool, err := k.GetOrInitPool(
			ctx,
			pairId,
			tickIndex,
			feeTiers[feeIndex],
		)

		if err != nil {
			return nil, nil, err
		}

		inAmount0, inAmount1, outShares := pool.Deposit(amount0, amount1, totalShares, autoswap)

		pool.Save(ctx, k)

		if outShares.GT(sdk.ZeroInt()) { // update shares accounting
			if err := k.MintShares(ctx, receiverAddr, outShares, sharesId); err != nil {
				return nil, nil, err
			}
		}

		if inAmount0.Equal(sdk.ZeroInt()) && inAmount1.Equal(sdk.ZeroInt()) {
			return nil, nil, types.ErrZeroTrueDeposit
		}

		amounts0Deposited[i] = inAmount0
		amounts1Deposited[i] = inAmount1
		totalAmountReserve0 = totalAmountReserve0.Add(inAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(inAmount1)

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
	totalReserve0ToRemove := sdk.ZeroInt()
	totalReserve1ToRemove := sdk.ZeroInt()
	feeTiers := k.GetAllFeeTier(ctx)

	for i, feeIndex := range msg.FeeIndexes {
		sharesToRemove := msg.SharesToRemove[i]
		tickIndex := msg.TickIndexes[i]

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
			return sdkerrors.Wrapf(types.ErrInsufficientShares, "%s does not have %s shares of type %s", msg.Creator, sharesToRemove, sharesId)
		}

		outAmount0, outAmount1 := pool.Withdraw(sharesToRemove, totalShares)
		pool.Save(ctx, k)
		if sharesToRemove.GT(sdk.ZeroInt()) { // update shares accounting
			if err := k.BurnShares(ctx, callerAddr, sharesToRemove, sharesId); err != nil {
				return err
			}
		}

		totalReserve0ToRemove = totalReserve0ToRemove.Add(outAmount0)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(outAmount1)

		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes[i]),
			fmt.Sprint(msg.FeeIndexes[i]),
			pool.LowerTick0.LPReserve.Add(outAmount0).String(),
			pool.UpperTick1.LPReserve.Add(outAmount1).String(),
			pool.LowerTick0.LPReserve.String(),
			pool.UpperTick1.LPReserve.String(),
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
	cacheCtx, writeCache := ctx.CacheContext()
	pairId, err := CreatePairIdFromUnsorted(tokenIn, tokenOut)
	if err != nil {
		return sdk.Coin{}, err
	}
	pair := types.NewDirectionalTradingPair(pairId, tokenIn, tokenOut)
	if err != nil {
		return sdk.Coin{}, err
	}

	remainingIn := msg.AmountIn
	totalOut := sdk.ZeroInt()

	// verify that amount left is not zero and that there are additional valid ticks to check
	liqIter := NewLiquidityIterator(k, ctx, pair)
	defer liqIter.Close()
	for remainingIn.GT(sdk.ZeroInt()) {
		liq := liqIter.Next()
		if liq == nil {
			break
		}

		// break as soon as we iterated past tickLimit
		if liq.Price().LT(msg.LimitPrice) {
			break
		}

		// price only gets worse as we iterate, so we can greedily abort
		// when the price is too low for minOut to be reached.
		idealOut := totalOut.Add(remainingIn.ToDec().Mul(liq.Price()).TruncateInt())
		if idealOut.LT(msg.MinOut) {
			return sdk.Coin{}, types.ErrSlippageLimitReached
		}

		inAmount, outAmount := liq.Swap(remainingIn)

		remainingIn = remainingIn.Sub(inAmount)
		totalOut = totalOut.Add(outAmount)

		liq.Save(cacheCtx, k)

	}

	if totalOut.LT(msg.MinOut) || msg.AmountIn.Equal(remainingIn) {
		return sdk.Coin{}, types.ErrSlippageLimitReached
	}

	// TODO: Move this to a separate ExecuteSwap function. Ditto for all other *Core fns
	amountToDeposit := msg.AmountIn.Sub(remainingIn)
	coinIn := sdk.NewCoin(tokenIn, amountToDeposit)
	coinOut := sdk.NewCoin(tokenOut, totalOut)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
		return sdk.Coin{}, err
	}

	if totalOut.GT(sdk.ZeroInt()) {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return sdk.Coin{}, err
		}
	}
	writeCache()
	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		tokenIn, tokenOut, msg.AmountIn.String(), totalOut.String(), msg.MinOut.String(),
	))

	return coinOut, nil
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder, tokenIn string, tokenOut string, callerAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := SortTokens(tokenIn, tokenOut)
	if err != nil {
		return err
	}
	pairId := CreatePairId(token0, token1)
	tickIndex := msg.TickIndex
	receiver := msg.Receiver

	var placeTrancheTick types.TickLiquidity
	placeTrancheTick, found := k.GetPlaceTrancheTick(ctx, pairId, tokenIn, tickIndex)

	if !found {
		placeTrancheTick, err = k.InitPlaceTrancheTick(ctx, pairId, tokenIn, tickIndex)
		if err != nil {
			return err
		}
	}

	if k.IsBehindEnemyLines(ctx, pairId, msg.TokenIn, tickIndex) {
		return types.ErrPlaceLimitOrderBehindPairLiquidity
	}

	tranche := placeTrancheTick.LimitOrderTranche
	placeTrancheIndex := tranche.TrancheIndex

	trancheUser := k.GetOrInitLimitOrderTrancheUser(goCtx, pairId, tickIndex, tokenIn, placeTrancheIndex, receiver)

	tranche.ReservesTokenIn = tranche.ReservesTokenIn.Add(msg.AmountIn)
	tranche.TotalTokenIn = tranche.TotalTokenIn.Add(msg.AmountIn)
	trancheUser.SharesOwned = trancheUser.SharesOwned.Add(msg.AmountIn)

	if msg.AmountIn.GT(sdk.ZeroInt()) {
		k.SetTickLiquidity(ctx, placeTrancheTick)
		k.SetLimitOrderTrancheUser(ctx, trancheUser)

		coin0 := sdk.NewCoin(tokenIn, msg.AmountIn)
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0})
		if err != nil {
			return err
		}
	}

	ctx.EventManager().EmitEvent(types.CreatePlaceLimitOrderEvent(msg.Creator,
		msg.Receiver,
		token0,
		token1,
		msg.TokenIn,
		msg.AmountIn.String(),
		msg.AmountIn.String(),
		strconv.FormatUint(placeTrancheIndex, 10),
	))

	return nil
}

// Handles MsgCancelLimitOrder, removing a specifed number of shares from a limit order and returning the respective amount in terms of the reserve to the user
func (k Keeper) CancelLimitOrderCore(goCtx context.Context, msg *types.MsgCancelLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId := CreatePairId(token0, token1)

	tick, tickFound := k.GetTickLiquidityLO(ctx, pairId, msg.KeyToken, msg.TickIndex, msg.Key)
	if !tickFound {
		return types.ErrActiveLimitOrderNotFound
	}

	trancheUser, found := k.GetLimitOrderTrancheUser(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	if !found {
		return types.ErrIntOverflowLimitOrderTrancheUser
	}
	// checks that the user has some number of limit order shares wished to withdraw
	if trancheUser.SharesOwned.LTE(sdk.ZeroInt()) {
		return types.ErrNotEnoughLimitOrderShares
	}

	tranche := NewLimitOrderTranche(&tick)

	amountToCancel := tranche.Cancel(trancheUser)
	trancheUser.SharesCancelled = trancheUser.SharesCancelled.Add(amountToCancel)

	if amountToCancel.GT(sdk.ZeroInt()) {
		// See top NOTE on rounding
		coinOut := sdk.NewCoin(msg.KeyToken, amountToCancel)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
		k.SetLimitOrderTrancheUser(ctx, trancheUser)
		tranche.Save(ctx, k)

	} else {
		return sdkerrors.Wrapf(types.ErrCancelEmptyLimitOrder, "%d", tranche.TrancheIndex)
	}

	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountToCancel.String(),
	))

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

	trancheUser, found := k.GetLimitOrderTrancheUser(
		ctx,
		pairId,
		tickIndex,
		orderTokenIn,
		trancheIndex,
		msg.Creator,
	)
	if !found {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderTrancheUserNotFound, "tranche %d, user %s", trancheIndex, msg.Creator)
	}

	sharesToWithdraw := trancheUser.SharesOwned.Sub(trancheUser.SharesCancelled)

	// checks that the user has some number of limit order shares wished to withdraw
	if sharesToWithdraw.LTE(sdk.ZeroInt()) {
		return types.ErrNotEnoughLimitOrderShares
	}

	tranche, wasFilled, found := k.GetLimitOrderTranche(ctx, pairId, tickIndex, msg.KeyToken, trancheIndex)
	if !found {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderTrancheNotFound, "%d", trancheIndex)
	}

	var priceLimitInToOut sdk.Dec
	var priceLimitOutToIn sdk.Dec
	if orderTokenIn == token0 {
		priceLimitInToOut = MustCalcPrice0To1(tranche.TickIndex)
	} else {
		priceLimitInToOut = MustCalcPrice1To0(tranche.TickIndex)
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

	// TODO: this is a bit of a messy pattern
	if wasFilled {
		k.SetFilledLimitOrderTranche(ctx, tranche.CreateFilledTranche())
	} else {
		k.SetTickLiquidity(ctx, types.TickLiquidity{
			PairId:            pairId,
			TokenIn:           msg.KeyToken,
			TickIndex:         tickIndex,
			LiquidityType:     types.LiquidityTypeLO,
			LiquidityIndex:    trancheIndex,
			LimitOrderTranche: &tranche,
		})
	}

	if amountOutTokenOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(orderTokenOut, amountOutTokenOut.TruncateInt())
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return types.ErrWithdrawEmptyLimitOrder
	}

	ctx.EventManager().EmitEvent(types.WithdrawFilledLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountOutTokenOut.String(),
	))

	return nil
}
