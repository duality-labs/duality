package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type UserProfile struct {
	Address sdk.AccAddress
}

func NewUserProfile(address sdk.AccAddress) UserProfile {
	return UserProfile{Address: address}
}

func (u UserProfile) GetAllLimitOrders(ctx sdk.Context, k Keeper) []types.LimitOrderTrancheUser {
	return k.GetAllLimitOrderTrancheUserForAddress(ctx, u.Address)
}

func (u UserProfile) GetAllDeposits(ctx sdk.Context, k Keeper) []types.DepositRecord {
	var depositArr []types.DepositRecord
	k.bankKeeper.IterateAccountBalances(ctx, u.Address,
		func(sharesMaybe sdk.Coin) bool {
			depositDenom, err := types.NewDepositDenomFromString(sharesMaybe.Denom)
			if err != nil {
				return false
			}

			depositRecord := types.DepositRecord{
				PairID:          depositDenom.PairID,
				SharesOwned:     sharesMaybe.Amount,
				CenterTickIndex: depositDenom.Tick,
				LowerTickIndex:  depositDenom.Tick - int64(depositDenom.Fee),
				UpperTickIndex:  depositDenom.Tick + int64(depositDenom.Fee),
				Fee:             depositDenom.Fee,
			}
			depositArr = append(depositArr, depositRecord)

			return false
		},
	)

	return depositArr
}

func (u UserProfile) GetAllPositions(ctx sdk.Context, k Keeper) types.UserPositions {
	deposits := u.GetAllDeposits(ctx, k)
	limitOrders := u.GetAllLimitOrders(ctx, k)

	return types.UserPositions{
		PoolDeposits: deposits,
		LimitOrders:  limitOrders,
	}
}
