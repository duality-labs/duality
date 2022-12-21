package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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
