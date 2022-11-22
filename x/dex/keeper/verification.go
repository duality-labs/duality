package keeper

import (
	"context"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type Message interface {
	GetTokenA() string
	GetTokenB() string
	GetCreator() string
	GetReceiver() string
}

func (k Keeper) ValidateTokens(tokenA string, tokenB string) (string, string, error){
	// lexographically sort token0, token1
	token0, token1, err := k.SortTokens(tokenA, tokenB)
	if err != nil {
		return "", "", err
	}
	return token0, token1, nil
}

func (k Keeper) ValidateAddress(address string, addrType string)(sdk.AccAddress, error){
	// Converts input address (string) to sdk.AccAddress
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil,  sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid %s address (%s)", addrType, err)
	}
	return addr, nil
}

func (k Keeper) ValidateFees(ctx sdk.Context, feeIndexes []uint64) error{

	feeCount := k.GetFeeListCount(ctx)

	// make sure that all feeIndexes (fee list index) is a valid index of the fee tier
	for _, feeIndex := range feeIndexes {
		if feeIndex >= feeCount {
			return  sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", feeIndex)
		}
	}
	return nil
}

func (k Keeper) ValidateCore(msg Message)(tokenA string, tokenB string, callerAddre sdk.AccAddress, receiverAddr sdk.AccAddress, err error){


	token0, token1, err := k.ValidateTokens(msg.GetTokenA(), msg.GetTokenB())
	if err != nil {
		return "", "", nil, nil, err
	}

	// NOTE: technically I don't think we have to handle an error here since it will be caught upstream by ValidateBasic
	callerAddr, err := k.ValidateAddress(msg.GetCreator(), "creator")
	if err != nil {
		return "", "", nil, nil, err
	}

	receiverAddr, err = k.ValidateAddress(msg.GetReceiver(), "receiver")
	if err != nil {
		return "", "", nil, nil, err
	}


	return token0, token1, callerAddr, receiverAddr, nil
}


func (k Keeper) DepositVerification(goCtx context.Context, msg types.MsgDeposit) (string, string, sdk.AccAddress, []sdk.Dec, []sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, callerAddr, _, err := k.ValidateCore(&msg)
	if err != nil {
		return "", "", nil, nil, nil, err
	}

	err = k.ValidateFees(ctx, msg.FeeIndexes)
	if err != nil {
		return "", "", nil, nil, nil, err
	}

	amounts0 := msg.AmountsA
	amounts1 := msg.AmountsB

	if token0 != msg.TokenA {
		amounts0, amounts1 = msg.AmountsB, msg.AmountsA
	}

	totalAmount0ToDeposit := sdk.ZeroDec()
	totalAmount1ToDeposit := sdk.ZeroDec()
	// checks that amount0, amount1 are both not zero, and that the user has the balances they wish to deposit
	for i, _ := range amounts0 {
		// Error checking for valid sdk.Dec
		if err != nil || (amounts0[i].Equal(sdk.ZeroDec()) && amounts1[i].Equal(sdk.ZeroDec())) {
			return "", "", nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid amount: %s", err)
		}

		totalAmount0ToDeposit = totalAmount0ToDeposit.Add(amounts0[i])
		totalAmount1ToDeposit = totalAmount1ToDeposit.Add(amounts1[i])
	}

	AccountToken0Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, token0).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts

	if AccountToken0Balance.LT(totalAmount0ToDeposit) {
		return "", "", nil, nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	AccountsToken1Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, token1).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts

	if AccountsToken1Balance.LT(totalAmount1ToDeposit) {
		return "", "", nil, nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, amounts0, amounts1, nil
}

func (k Keeper) WithdrawlVerification(goCtx context.Context, msg types.MsgWithdrawl) (string, string, sdk.AccAddress, sdk.AccAddress, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, callerAddr, receiverAddr, err := k.ValidateCore(&msg)
	if err != nil {
		return "", "", nil, nil, err
	}

	// makes sure that there is the same number of sharesToRemove as ticks specfied
	if len(msg.SharesToRemove) != len(msg.TickIndexes) || len(msg.SharesToRemove) != len(msg.FeeIndexes) {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrUnbalancedTxArray, "Input Arrays are not of the same length")
	}

	err = k.ValidateFees(ctx, msg.FeeIndexes)
	if err != nil {
		return "", "", nil, nil, err
	}

	pairId := k.CreatePairId(token0, token1)

	// checks that the user has the specified number of shares they wish to withdraw
	for i, shareToRemove := range msg.SharesToRemove {
		shares, sharesFound := k.GetShares(ctx, msg.Creator, pairId, msg.TickIndexes[i], msg.FeeIndexes[i])

		if !sharesFound {
			return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
		}

		if shares.SharesOwned.LT(shareToRemove) {
			return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
		}
	}

	return token0, token1, callerAddr, receiverAddr, nil
}

func (k Keeper) SwapVerification(goCtx context.Context, msg types.MsgSwap) (string, string, sdk.AccAddress, sdk.AccAddress, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, callerAddr, receiverAddr, err := k.ValidateCore(&msg)
	if err != nil {
		return "", "", nil, nil, err
	}

	if msg.TokenIn != token0 && msg.TokenIn != token1 {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "TokenIn must be either Token0 or Token1")
	}

	AccountsAmountInBalance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.TokenIn).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsAmountInBalance.LT(msg.AmountIn) {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, receiverAddr, nil
}

func (k Keeper) PlaceLimitOrderVerification(goCtx context.Context, msg types.MsgPlaceLimitOrder) (string, string, sdk.AccAddress, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	token0, token1, callerAddr, _, err := k.ValidateCore(&msg)
	if err != nil {
		return "", "", nil, err
	}

	if msg.TokenIn != token0 && msg.TokenIn != token1 {
		return "", "", nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "TokenIn must be either Tokne0 or Token1")
	}

	AccountsAmountInBalance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.TokenIn).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsAmountInBalance.LT(msg.AmountIn) {
		return "", "", nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, nil
}

func (k Keeper) WithdrawLimitOrderVerification(goCtx context.Context, msg types.MsgWithdrawFilledLimitOrder) (string, string, sdk.AccAddress, sdk.AccAddress, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	token0, token1, callerAddr, receiverAddr, err := k.ValidateCore(&msg)
	if err != nil {
		return "", "", nil, nil, err
	}

	pairId := k.CreatePairId(token0, token1)

	shares, sharesFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Receiver)
	if !sharesFound {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	// checks that the user has some number of limit order shares wished to withdraw
	if shares.SharesOwned.LTE(sdk.ZeroDec()) {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	_ = ctx
	return token0, token1, callerAddr, receiverAddr, nil
}

func (k Keeper) CancelLimitOrderVerification(goCtx context.Context, msg types.MsgCancelLimitOrder) (string, string, sdk.AccAddress, sdk.AccAddress, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, callerAddr, receiverAddr, err := k.ValidateCore(&msg)
	if err != nil {
		return "", "", nil, nil, err
	}

	// createPairId (token0/ token1)
	pairId := k.CreatePairId(token0, token1)

	shares, sharesFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)

	if !sharesFound {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	// checks that the user has some number of limit order shares wished to withdraw
	if shares.SharesOwned.LTE(sdk.ZeroDec()) {
		return "", "", nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	_ = ctx
	return token0, token1, callerAddr, receiverAddr, nil
}
