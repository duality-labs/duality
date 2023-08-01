package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

func (k Keeper) GetAllDepositsForAddress(
	ctx sdk.Context,
	addr sdk.AccAddress,
) []*types.DepositRecord {
	var depositArr []*types.DepositRecord
	k.bankKeeper.IterateAccountBalances(ctx, addr,
		func(sharesMaybe sdk.Coin) bool {
			poolID, err := types.ParsePoolIDFromDepositDenom(sharesMaybe.Denom)
			if err != nil {
				return false
			}

			metadata, err := k.GetPoolMetadataByID(ctx, poolID)
			if err != nil {
				return false
			}

			depositRecord := &types.DepositRecord{
				PairID:          metadata.PairId,
				SharesOwned:     sharesMaybe.Amount,
				CenterTickIndex: metadata.NormalizedCenterTickIndex,
				LowerTickIndex: metadata.NormalizedCenterTickIndex - utils.MustSafeUint64(
					metadata.Fee,
				),
				UpperTickIndex: metadata.NormalizedCenterTickIndex + utils.MustSafeUint64(
					metadata.Fee,
				),
				Fee: metadata.Fee,
			}
			depositArr = append(depositArr, depositRecord)

			return false
		},
	)

	return depositArr
}
