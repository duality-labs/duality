package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type UserProfile struct {
	Address sdk.AccAddress
}

func NewUserProfile(address sdk.AccAddress) UserProfile {
	return UserProfile{Address: address}
}

func (u UserProfile) GetAllLimitOrders(goCtx context.Context, k Keeper) []types.LimitOrderTrancheUser {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return k.GetAllLimitOrderTrancheUserForAddress(ctx, u.Address)
}

func (u UserProfile) GetAllDeposits(goCtx context.Context, k Keeper) []types.DepositRecord {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var depositArr []types.DepositRecord
	feeTiers := k.GetAllFeeTier(ctx)
	k.bankKeeper.IterateAccountBalances(ctx, u.Address,
		func(sharesMaybe sdk.Coin) bool {
			depositRecord, err := DepositSharesToData(sharesMaybe, feeTiers)
			if err == nil {
				depositArr = append(depositArr, depositRecord)
			}

			return false

		},
	)
	return depositArr
}

func (u UserProfile) GetAllPositions(goCtx context.Context, k Keeper) types.UserPositions {
	deposits := u.GetAllDeposits(goCtx, k)
	limitOrders := u.GetAllLimitOrders(goCtx, k)

	return types.UserPositions{
		PoolDeposits: deposits,
		LimitOrders:  limitOrders,
	}
}
