package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) depositVerification(goCtx context.Context, msg types.MsgDeposit) (string, string, sdk.AccAddress, sdk.Dec, sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := k.SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	maxFee := k.GetFeeListCount(ctx)

	if msg.FeeIndex >= maxFee {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", msg.FeeIndex)
	}

	amount0, err := sdk.NewDecFromStr(msg.AmountA)
	amount1, err := sdk.NewDecFromStr(msg.AmountB)

	// Error checking for valid sdk.Dec
	if err != nil || (amount0.Equal(sdk.ZeroDec()) && amount1.Equal(sdk.ZeroDec())) {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid amount: %s", err)
	}

	if token0 != msg.TokenA {
		tmp := amount0
		amount0 = amount1
		amount1 = tmp
	}

	AccountsToken0Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, token0).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsToken0Balance.LT(amount0) {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	AccountsToken1Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, token1).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsToken1Balance.LT(amount1) {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, amount0, amount1, nil
}

func (k Keeper) withdrawlVerification(goCtx context.Context, msg types.MsgWithdrawl) (string, string, sdk.AccAddress, sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := k.SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	maxFee := k.GetFeeListCount(ctx)

	if msg.FeeIndex >= maxFee {
		return "", "", nil, sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", msg.FeeIndex)
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	sharesToRemove, err := sdk.NewDecFromStr(msg.SharesToRemove)

	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	pairId := k.CreatePairId(token0, token1)
	shares, sharesFound := k.GetShares(ctx, msg.Creator, pairId, msg.PriceIndex, msg.FeeIndex)

	if !sharesFound {
		return "", "", nil, sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	if shares.SharesOwned.LT(sharesToRemove) {
		return "", "", nil, sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
	}

	return token0, token1, callerAddr, sharesToRemove, nil
}

func (k Keeper) swapVerification(goCtx context.Context, msg types.MsgSwap) (string, string, sdk.AccAddress, sdk.Dec, sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := k.SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	amountIn, err := sdk.NewDecFromStr(msg.AmountIn)

	if msg.TokenIn != token0 && msg.TokenIn != token1 {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "TokenIn must be either Tokne0 or Token1")
	}
	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	minOut, err := sdk.NewDecFromStr(msg.MinOut)

	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	AccountsAmountInBalance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.TokenIn).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsAmountInBalance.LT(amountIn) {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, amountIn, minOut, nil
}
