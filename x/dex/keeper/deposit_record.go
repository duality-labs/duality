package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

func (k Keeper) GetAllDepositsForAddress(ctx sdk.Context, addr sdk.AccAddress) []*types.DepositRecord {
	var depositArr []*types.DepositRecord
	k.bankKeeper.IterateAccountBalances(ctx, addr,
		func(sharesMaybe sdk.Coin) bool {
			depositDenom, err := types.NewDepositDenomFromString(sharesMaybe.Denom)
			if err != nil {
				return false
			}

			depositRecord := &types.DepositRecord{
				PairID:          depositDenom.PairID,
				SharesOwned:     sharesMaybe.Amount,
				CenterTickIndex: depositDenom.Tick,
				LowerTickIndex:  depositDenom.Tick - utils.MustSafeUint64(depositDenom.Fee),
				UpperTickIndex:  depositDenom.Tick + utils.MustSafeUint64(depositDenom.Fee),
				Fee:             depositDenom.Fee,
			}
			depositArr = append(depositArr, depositRecord)

			return false
		},
	)

	return depositArr
}
