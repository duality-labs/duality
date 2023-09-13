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
			err := types.ValidatePoolDenom(sharesMaybe.Denom)
			if err != nil {
				return false
			}

			poolMetadata, err := k.GetPoolMetadataByDenom(ctx, sharesMaybe.Denom)
			if err != nil {
				panic("Can't get info for PoolDenom")
			}
			fee := utils.MustSafeUint64ToInt64(poolMetadata.Fee)
			depositRecord := &types.DepositRecord{
				PairID:          poolMetadata.PairID,
				SharesOwned:     sharesMaybe.Amount,
				CenterTickIndex: poolMetadata.Tick,
				LowerTickIndex:  poolMetadata.Tick - fee,
				UpperTickIndex:  poolMetadata.Tick + fee,
				Fee:             poolMetadata.Fee,
			}
			depositArr = append(depositArr, depositRecord)

			return false
		},
	)

	return depositArr
}
