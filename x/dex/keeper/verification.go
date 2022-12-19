package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) DepositVerification(goCtx context.Context, msg types.MsgDeposit) (string, string, sdk.AccAddress, []sdk.Int, []sdk.Int, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate msg
	if err := msg.ValidateBasic(); err != nil {
		return "", "", nil, nil, nil, err
	}
	callerAddr := sdk.MustAccAddressFromBech32(msg.Receiver)

	// lexographically sort token0, token1
	token0, token1, err := SortTokens(ctx, msg.TokenA, msg.TokenB)
	if err != nil {
		return "", "", nil, nil, nil, err
	}
	amounts0, amounts1 := SortAmounts(msg.TokenA, token0, msg.AmountsA, msg.AmountsB)

	return token0, token1, callerAddr, amounts0, amounts1, nil
}

func (k Keeper) WithdrawlVerification(goCtx context.Context, msg types.MsgWithdrawl) (string, string, sdk.AccAddress, sdk.AccAddress, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// lexographically sort token0, token1
	token0, token1, err := SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrInvalidTradingPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// gets total number of fee tiers
	feeCount := k.GetFeeTierCount(ctx)

	// makes sure that there is the same number of sharesToRemove as ticks specfied
	if len(msg.SharesToRemove) != len(msg.TickIndexes) || len(msg.SharesToRemove) != len(msg.FeeIndexes) {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrUnbalancedTxArray, "Input Arrays are not of the same length")
	}

	// make sure that all feeIndexes (fee list index) is a valid index of the fee tier
	for i, _ := range msg.FeeIndexes {
		if msg.FeeIndexes[i] >= feeCount {
			return "", "", nil, nil, sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", msg.FeeIndexes[i])
		}
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	// Error checking for valid sdk.Int
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	FeeTier := k.GetAllFeeTier(ctx)
	// count total shares to remove from each pair:tickIndex,feeIndex
	totalSharesIn := make(map[string]sdk.Int)
	for i, sharesToRemove := range msg.SharesToRemove {
		fee := FeeTier[msg.FeeIndexes[i]].Fee
		sharesId := k.CreateSharesId(token0, token1, msg.TickIndexes[i], fee)
		if accum, ok := totalSharesIn[sharesId]; !ok {
			totalSharesIn[sharesId] = sharesToRemove
		} else {
			totalSharesIn[sharesId] = accum.Add(sharesToRemove)
		}
	}

	// verify sufficient shares for each
	for sharesDenom, amountToRemove := range totalSharesIn {
		sharesBalance := k.bankKeeper.GetBalance(ctx, callerAddr, sharesDenom)
		if sharesBalance.Amount.LT(amountToRemove) {
			return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
		}
	}

	return token0, token1, callerAddr, receiverAddr, nil
}

func (k Keeper) SwapVerification(goCtx context.Context, msg types.MsgSwap) (string, string, sdk.AccAddress, sdk.AccAddress, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// lexographically sort token0, token1
	token0, token1, err := SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrInvalidTradingPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	if msg.TokenIn != token0 && msg.TokenIn != token1 {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrInvalidTradingPair, "TokenIn must be either Tokne0 or Token1")
	}
	// Error checking for valid sdk.Int
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	AccountsAmountInBalance := k.bankKeeper.GetBalance(ctx, callerAddr, msg.TokenIn).Amount

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsAmountInBalance.LT(msg.AmountIn) {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, receiverAddr, nil
}

func (k Keeper) PlaceLimitOrderVerification(goCtx context.Context, msg types.MsgPlaceLimitOrder) (string, string, sdk.AccAddress, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// lexographically sort token0, token1
	token0, token1, err := SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(types.ErrInvalidTradingPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	//NOTE: We do not use the sdk.AccAddress of Receiver in PlaceLimitOrder and thus do not need to save it
	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	if msg.TokenIn != token0 && msg.TokenIn != token1 {
		return "", "", nil, sdkerrors.Wrapf(types.ErrInvalidTradingPair, "TokenIn must be either Tokne0 or Token1")
	}
	// Error checking for valid sdk.Int
	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	AccountsAmountInBalance := k.bankKeeper.GetBalance(ctx, callerAddr, msg.TokenIn).Amount

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsAmountInBalance.LT(msg.AmountIn) {
		return "", "", nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, nil
}

func (k Keeper) WithdrawLimitOrderVerification(goCtx context.Context, msg types.MsgWithdrawFilledLimitOrder) (string, string, sdk.AccAddress, sdk.AccAddress, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrInvalidTradingPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	ReceiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	pairId := k.CreatePairId(token0, token1)

	shares, sharesFound := k.GetLimitOrderTrancheUser(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Receiver)
	if !sharesFound {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	// checks that the user has some number of limit order shares wished to withdraw
	if shares.SharesOwned.LTE(sdk.ZeroInt()) {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	_ = ctx
	return token0, token1, callerAddr, ReceiverAddr, nil
}

func (k Keeper) CancelLimitOrderVerification(goCtx context.Context, msg types.MsgCancelLimitOrder) (string, string, sdk.AccAddress, sdk.AccAddress, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// lexographically sort token0, token1
	token0, token1, err := SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrInvalidTradingPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return "", "", nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	// createPairId (token0/ token1)
	pairId := k.CreatePairId(token0, token1)

	shares, sharesFound := k.GetLimitOrderTrancheUser(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)

	if !sharesFound {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	// checks that the user has some number of limit order shares wished to withdraw
	if shares.SharesOwned.LTE(sdk.ZeroInt()) {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	_ = ctx
	return token0, token1, callerAddr, receiverAddr, nil
}
