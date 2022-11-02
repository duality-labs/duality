package keeper

import (
	"context"
	"fmt"

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

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, amount0, amount1, err := k.DepositVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	err = k.DepositCore(goCtx, msg, token0, token1, createrAddr, amount0, amount1)

	if err != nil {
		return nil, err
	}

	_ = ctx

	return &types.MsgDepositResponse{}, nil
}

func (k msgServer) Withdrawl(goCtx context.Context, msg *types.MsgWithdrawl) (*types.MsgWithdrawlResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, receiverAddr, err := k.WithdrawlVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	err = k.WithdrawCore(goCtx, msg, token0, token1, createrAddr, receiverAddr)
	_ = ctx

	return &types.MsgWithdrawlResponse{}, nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, receiverAddr, err := k.SwapVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	var amount_out sdk.Dec

	if msg.TokenIn == token0 {
		amount_out, err = k.Swap0to1(goCtx, msg, token0, token1, createrAddr)

		if err != nil {
			return nil, err
		}

		if msg.AmountIn.GT(sdk.ZeroDec()) {
			coinIn := sdk.NewCoin(token0, sdk.NewIntFromBigInt(msg.AmountIn.BigInt()))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		} else {
			return &types.MsgSwapResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
		}

		if amount_out.GT(sdk.ZeroDec()) {

			coinOut := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount_out.BigInt()))
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		}

		ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
			token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
		))

	} else {
		amount_out, err = k.Swap1to0(goCtx, msg, token0, token1, createrAddr)

		if err != nil {
			return nil, err
		}

		if msg.AmountIn.GT(sdk.ZeroDec()) {
			coinIn := sdk.NewCoin(token1, sdk.NewIntFromBigInt(msg.AmountIn.BigInt()))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		} else {
			return &types.MsgSwapResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
		}

		if amount_out.GT(sdk.ZeroDec()) {

			coinOut := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amount_out.BigInt()))
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
				return &types.MsgSwapResponse{}, err
			}

		}
		ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
			token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
		))

	}

	_ = ctx

	return &types.MsgSwapResponse{}, nil
}

// TODO: ADD CHECKS THAT ALL PARAMETERS ARE PASSED IN
func (k msgServer) Route(goCtx context.Context, msg *types.MsgRoute) (*types.MsgRouteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	createrAddr, receiverAddr, amountIn, err := k.routeVerification(goCtx, *msg)
	// Nil for now
	_ = receiverAddr

	if err != nil {
		return nil, err
	}

	var amount_out sdk.Dec

	// Number of potential chunks in route
	// TODO: Add this as argument to MsgRoute?
	var numChunks int64 = 10
	amount_out, err = k.DynamicRouteSwap(goCtx, createrAddr, receiverAddr, msg.TokenIn, msg.TokenOut, amountIn, msg.MinOut, numChunks)

	if err != nil {
		return nil, err
	}

	if amountIn.GT(sdk.ZeroDec()) {
		coinIn := sdk.NewCoin(msg.TokenIn, sdk.NewIntFromBigInt(amountIn.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
			return &types.MsgRouteResponse{}, err
		}
	} else {
		return &types.MsgRouteResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
	}

	if amount_out.GT(sdk.ZeroDec()) {

		coinOut := sdk.NewCoin(msg.TokenOut, sdk.NewIntFromBigInt(amount_out.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return &types.MsgRouteResponse{}, err
		}
	}

	_ = ctx
	fmt.Println("Amount Out from Swap", amount_out)
	return &types.MsgRouteResponse{}, nil
}

func (k msgServer) PlaceLimitOrder(goCtx context.Context, msg *types.MsgPlaceLimitOrder) (*types.MsgPlaceLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, err := k.PlaceLimitOrderVerification(goCtx, *msg)

	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}

	err = k.PlaceLimitOrderCore(goCtx, msg, token0, token1, createrAddr)

	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}
	_ = ctx

	return &types.MsgPlaceLimitOrderResponse{}, nil
}

func (k msgServer) WithdrawFilledLimitOrder(goCtx context.Context, msg *types.MsgWithdrawFilledLimitOrder) (*types.MsgWithdrawFilledLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, receiverAddr, err := k.WithdrawLimitOrderVerification(goCtx, *msg)

	if err != nil {
		return &types.MsgWithdrawFilledLimitOrderResponse{}, err
	}

	err = k.WithdrawFilledLimitOrderCore(goCtx, msg, token0, token1, createrAddr, receiverAddr)

	if err != nil {
		return &types.MsgWithdrawFilledLimitOrderResponse{}, err
	}

	_ = ctx

	return &types.MsgWithdrawFilledLimitOrderResponse{}, nil
}

func (k msgServer) CancelLimitOrder(goCtx context.Context, msg *types.MsgCancelLimitOrder) (*types.MsgCancelLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, receiverAddr, err := k.CancelLimitOrderVerification(goCtx, *msg)

	if err != nil {
		return &types.MsgCancelLimitOrderResponse{}, err
	}

	err = k.CancelLimitOrderCore(goCtx, msg, token0, token1, createrAddr, receiverAddr)

	if err != nil {
		return &types.MsgCancelLimitOrderResponse{}, err
	}
	_ = ctx

	return &types.MsgCancelLimitOrderResponse{}, nil
}
