package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Handles logic for MsgDeposit
func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Calls DepositVerification to run pre-checks on input paramters and msg.Creator
	//return sorted token0 and token1 and respective amount0 and amount1 (sdk.Decs)
	token0, token1, createrAddr, amount0, amount1, err := k.DepositVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	// Calls DeppsitCore which handles all logic / bankkeeping for MsgDeposit
	err = k.DepositCore(goCtx, msg, token0, token1, createrAddr, amount0, amount1)

	if err != nil {
		return nil, err
	}

	_ = ctx

	return &types.MsgDepositResponse{}, nil
}

// Handles logic for MsgWithdrawl
func (k msgServer) Withdrawl(goCtx context.Context, msg *types.MsgWithdrawl) (*types.MsgWithdrawlResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Calls WithdrawVerification to run pre-checks on input paramters and msg.Creator
	//return sorted token0 and token1 and creator and receiver Address (sdk.AccAddress)
	token0, token1, createrAddr, receiverAddr, err := k.WithdrawlVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	// Calls WithdrawCore (in keeer/core.go) which handles all logic / bankkeeping for MsgWithdrawl
	err = k.WithdrawCore(goCtx, msg, token0, token1, createrAddr, receiverAddr)
	_ = ctx

	return &types.MsgWithdrawlResponse{}, nil
}

// Handles logiv for MsgSwap
func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Calls SwapVerification to run pre-checks on input paramters and msg.Creator
	//return sorted token0 and token1 and creator and receiver Address (sdk.AccAddress)
	token0, token1, createrAddr, receiverAddr, err := k.SwapVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	var amount_out sdk.Dec

	// check input direction (0 or 1)
	if msg.TokenIn == token0 {
		// Handles core logic,  calculating amount_out given some input amount
		amount_out, err = k.Swap0to1(goCtx, msg, token0, token1, createrAddr)

		if err != nil {
			return nil, err
		}

		// Sends amountIn from account to module
		if msg.AmountIn.GT(sdk.ZeroDec()) {
			coinIn := sdk.NewCoin(token0, sdk.NewIntFromBigInt(msg.AmountIn.BigInt()))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		} else {
			return &types.MsgSwapResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
		}

		// Sends amountOut from module to account
		if amount_out.GT(sdk.ZeroDec()) {

			coinOut := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount_out.BigInt()))
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		}

	} else {
		//If tokenIn == token1

		// Handles core logic,  calculating amount_out given some input amount
		amount_out, err = k.Swap1to0(goCtx, msg, token0, token1, createrAddr)

		if err != nil {
			return nil, err
		}

		// Sends amountIn from account to Module
		if msg.AmountIn.GT(sdk.ZeroDec()) {
			coinIn := sdk.NewCoin(token1, sdk.NewIntFromBigInt(msg.AmountIn.BigInt()))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		} else {
			return &types.MsgSwapResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
		}

		// Sends amount_out from module to account
		if amount_out.GT(sdk.ZeroDec()) {

			coinOut := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amount_out.BigInt()))
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
				return &types.MsgSwapResponse{}, err
			}

		}

	}

	_ = ctx

	return &types.MsgSwapResponse{}, nil
}

// Handles logic for MsgPlaceLimitOrder
func (k msgServer) PlaceLimitOrder(goCtx context.Context, msg *types.MsgPlaceLimitOrder) (*types.MsgPlaceLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Calls PlaceLimitOrderVerification to run pre-checks on input paramters and msg.Creator
	//return sorted token0 and token1 and creatorAddr (sdk.AccAddress)
	token0, token1, createrAddr, err := k.PlaceLimitOrderVerification(goCtx, *msg)

	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}

	// Calls PlaceLimitOrderCore (in keeper/core.go) which handles all logic / bankkeeping for MsgPlaceLimitOrder
	err = k.PlaceLimitOrderCore(goCtx, msg, token0, token1, createrAddr)

	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}
	_ = ctx

	return &types.MsgPlaceLimitOrderResponse{}, nil
}

// Handles logic for MsgWithdrawFilledLimitOrder
func (k msgServer) WithdrawFilledLimitOrder(goCtx context.Context, msg *types.MsgWithdrawFilledLimitOrder) (*types.MsgWithdrawFilledLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Calls WithdrawFilledLimitOrderVerification to run pre-checks on input paramters and msg.Creator
	//return sorted token0 and token1 and creatorAddr and receiverAddr (sdk.AccAddress)
	token0, token1, createrAddr, receiverAddr, err := k.WithdrawLimitOrderVerification(goCtx, *msg)

	if err != nil {
		return &types.MsgWithdrawFilledLimitOrderResponse{}, err
	}

	// Calls WithdrawFilled LimitOrder(in keeper/core.go) which handles all logic / bankkeeping for MsgPlaceLimitOrder
	err = k.WithdrawFilledLimitOrderCore(goCtx, msg, token0, token1, createrAddr, receiverAddr)

	if err != nil {
		return &types.MsgWithdrawFilledLimitOrderResponse{}, err
	}

	_ = ctx

	return &types.MsgWithdrawFilledLimitOrderResponse{}, nil
}

func (k msgServer) CancelLimitOrder(goCtx context.Context, msg *types.MsgCancelLimitOrder) (*types.MsgCancelLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Calls CancelLimitOrderVerification to run pre-checks on input paramters and msg.Creator
	//return sorted token0 and token1 and creatorAddr and receiverAddr (sdk.AccAddress)
	token0, token1, createrAddr, receiverAddr, err := k.CancelLimitOrderVerification(goCtx, *msg)

	if err != nil {
		return &types.MsgCancelLimitOrderResponse{}, err
	}

	// Calls CancelLimitOrder(in keeper/core.go) which handles all logic / bankkeeping for MsgCancelLimitOrder
	err = k.CancelLimitOrderCore(goCtx, msg, token0, token1, createrAddr, receiverAddr)

	if err != nil {
		return &types.MsgCancelLimitOrderResponse{}, err
	}
	_ = ctx

	return &types.MsgCancelLimitOrderResponse{}, nil
}
