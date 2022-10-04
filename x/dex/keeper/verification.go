package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) depositVerification(goCtx context.Context, msg types.MsgDeposit) (string, string, sdk.AccAddress, []sdk.Dec, []sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := k.SortTokens(msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, nil, nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return "", "", nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	feeCount := k.GetFeeListCount(ctx)

	for i, _ := range msg.FeeIndexes {
		if msg.FeeIndexes[i] >= feeCount {
			return "", "", nil, nil, nil, sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", msg.FeeIndexes[i])
		}
	}

	amounts0 := msg.AmountsA
	amounts1 := msg.AmountsB

	if token0 != msg.TokenA {
		tmp := msg.AmountsA
		amounts0 = msg.AmountsB
		amounts1 = tmp
	}

	for i, _ := range amounts0 {
		// Error checking for valid sdk.Dec
		if err != nil || (amounts0[i].Equal(sdk.ZeroDec()) && amounts1[i].Equal(sdk.ZeroDec())) {
			return "", "", nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid amount: %s", err)
		}

		AccountsToken0Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, token0).Amount)

		// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
		if AccountsToken0Balance.LT(amounts0[i]) {
			return "", "", nil, nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
		}

		AccountsToken1Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, token1).Amount)

		// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
		if AccountsToken1Balance.LT(amounts1[i]) {
			return "", "", nil, nil, nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
		}

	}

	return token0, token1, callerAddr, amounts0, amounts1, nil
}

func (k Keeper) withdrawlVerification(goCtx context.Context, msg types.MsgWithdrawl) (string, string, sdk.AccAddress, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := k.SortTokens(msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	feeCount := k.GetFeeListCount(ctx)

	if len(msg.SharesToRemove) != len(msg.TickIndexes) || len(msg.SharesToRemove) != len(msg.FeeIndexes) {
		return "", "", nil, sdkerrors.Wrapf(types.ErrUnbalancedTxArray, "Input Arrays are not of the same length")
	}

	for i, _ := range msg.FeeIndexes {
		if msg.FeeIndexes[i] >= feeCount {
			return "", "", nil, sdkerrors.Wrapf(types.ErrValidFeeIndexNotFound, "(%d) does not correspond to a valid fee", msg.FeeIndexes[i])
		}
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	pairId := k.CreatePairId(token0, token1)

	for i, shareToRemove := range msg.SharesToRemove {
		shares, sharesFound := k.GetShares(ctx, msg.Creator, pairId, msg.TickIndexes[i], msg.FeeIndexes[i])

		if !sharesFound {
			return "", "", nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
		}

		if shares.SharesOwned.LT(shareToRemove) {
			return "", "", nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares were found")
		}
	}

	return token0, token1, callerAddr, nil
}
func (k Keeper) routeVerification(goCtx context.Context, msg types.MsgRoute) (sdk.AccAddress, sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	return k.tradeVerification(ctx, msg.Creator, msg.Receiver, msg.AmountIn, msg.TokenIn, msg.MinOut)

}

func (k Keeper) swapVerification(goCtx context.Context, msg types.MsgSwap) (string, string, sdk.AccAddress, error) {

	token0, token1, err := k.SortTokens(msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	if msg.TokenIn != token0 && msg.TokenIn != token1 {
		return "", "", nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "TokenIn must be either Tokne0 or Token1")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	callerAddr, _, _, err := k.tradeVerification(ctx, msg.Creator, msg.Receiver, msg.AmountIn, msg.TokenIn, msg.MinOut)
	return token0, token1, callerAddr, err
}

func (k Keeper) tradeVerification(ctx sdk.Context, Creator string, Receiver string, UncheckedAmountIn sdk.Dec, TokenIn string, UncheckedMinOut sdk.Dec) (sdk.AccAddress, sdk.Dec, sdk.Dec, error) {

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(Creator)
	// Error checking for the calling address
	if err != nil {
		return nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(Receiver)
	// Error Checking for receiver address
	// Note we do not actually need to save the sdk.AccAddress here but we do want the address to be checked to determine if it valid
	if err != nil {
		return nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	AccountsAmountInBalance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, TokenIn).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsAmountInBalance.LT(UncheckedAmountIn) {
		return "", "", nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, nil
}
