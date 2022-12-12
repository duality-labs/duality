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

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, autoswap, createrAddr, amount0, amount1, err := k.DepositVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	// TODO: DECIDE WHETHER WE SHOULD HAVE A SEPERATE MSG SERVER FUNCTION FOR DEPOSIT
	Amounts0Deposit, Amounts1Deposit, err := k.DepositCore(
		goCtx,
		msg,
		token0,
		token1,
		createrAddr,
		amount0,
		amount1,
		autoswap,
	)

	if err != nil {
		return nil, err
	}

	_ = ctx

	return &types.MsgDepositResponse{Amounts0Deposit, Amounts1Deposit}, nil
}

func (k msgServer) Withdrawl(goCtx context.Context, msg *types.MsgWithdrawl) (*types.MsgWithdrawlResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, receiverAddr, err := k.WithdrawlVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	err = k.WithdrawCore(goCtx, msg, token0, token1, createrAddr, receiverAddr)

	if err != nil {
		return nil, err
	}

	_ = ctx

	return &types.MsgWithdrawlResponse{}, nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, receiverAddr, err := k.SwapVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	var amount_out sdk.Int
	var amount_left sdk.Int
	var coinOut sdk.Coin
	if msg.TokenIn == token0 {
		amount_out, amount_left, err = k.Swap0to1(goCtx, msg, token0, token1, createrAddr)

		if err != nil {
			return nil, err
		}

		amountToDeposit := msg.AmountIn.Sub(amount_left)
		if amountToDeposit.GT(sdk.ZeroInt()) {
			coinIn := sdk.NewCoin(token0, amountToDeposit)
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		} else {
			return &types.MsgSwapResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
		}

		if amount_out.GT(sdk.ZeroInt()) {

			coinOut = sdk.NewCoin(token1, amount_out)
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		}

	} else {
		amount_out, amount_left, err = k.Swap1to0(goCtx, msg, token0, token1, createrAddr)

		if err != nil {
			return nil, err
		}

		amountToDeposit := msg.AmountIn.Sub(amount_left)
		if amountToDeposit.GT(sdk.ZeroInt()) {
			coinIn := sdk.NewCoin(token1, amountToDeposit)
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		} else {
			return &types.MsgSwapResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
		}

		if amount_out.GT(sdk.ZeroInt()) {

			coinOut = sdk.NewCoin(token0, amount_out)
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
				return &types.MsgSwapResponse{}, err
			}

		}

	}

	_ = ctx

	return &types.MsgSwapResponse{coinOut}, nil
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
