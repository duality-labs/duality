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

	token0, token1, createrAddr, amount0, amount1, err := k.depositVerification(goCtx, *msg)

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

	token0, token1, createrAddr, err := k.withdrawlVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	err = k.WithdrawCore(goCtx, msg, token0, token1, createrAddr)
	_ = ctx

	return &types.MsgWithdrawlResponse{}, nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, amountIn, minOut, err := k.swapVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	var amount_out sdk.Dec

	if msg.TokenIn == token0 {
		amount_out, err = k.Swap0to1(goCtx, msg, token0, token1, createrAddr, amountIn, minOut)

		if err != nil {
			return nil, err
		}

		if amountIn.GT(sdk.ZeroDec()) {
			coinIn := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amountIn.BigInt()))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		} else {
			return &types.MsgSwapResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
		}

		if amount_out.GT(sdk.ZeroDec()) {

			coinOut := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount_out.BigInt()))
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(msg.Receiver), sdk.Coins{coinOut}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		}

	} else {
		amount_out, err = k.Swap1to0(goCtx, msg, token0, token1, createrAddr, amountIn, minOut)

		if err != nil {
			return nil, err
		}

		if amountIn.GT(sdk.ZeroDec()) {
			coinIn := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amountIn.BigInt()))
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, createrAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		} else {
			return &types.MsgSwapResponse{}, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
		}

		if amount_out.GT(sdk.ZeroDec()) {

			coinOut := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amount_out.BigInt()))
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(msg.Receiver), sdk.Coins{coinOut}); err != nil {
				return &types.MsgSwapResponse{}, err
			}
		}

	}

	_ = ctx

	return &types.MsgSwapResponse{}, nil
}
