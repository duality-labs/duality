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

	amount0, err := sdk.NewDecFromStr(msg.AmountA)
	amount1, err := sdk.NewDecFromStr(msg.AmountB)

	if token0 != msg.TokenA {
		tmp := amount0
		amount0 = amount1
		amount1 = tmp
	}

	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
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
