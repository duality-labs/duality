package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

// addStakeRefs adds appropriate reference keys preceded by a prefix.
// A prefix indicates whether the stake is unstaking or not.
func (k Keeper) addStakeRefs(ctx sdk.Context, stake *types.Stake) error {
	refKeys, err := getStakeRefKeys(stake)
	if err != nil {
		return err
	}
	for _, refKey := range refKeys {
		if err := k.addRefByKey(ctx, refKey, stake.ID); err != nil {
			return err
		}
	}
	return nil
}

// deleteStakeRefs deletes all the stake references of the stake with the given stake prefix.
func (k Keeper) deleteStakeRefs(ctx sdk.Context, stake *types.Stake) error {
	refKeys, err := getStakeRefKeys(stake)
	if err != nil {
		return err
	}
	for _, refKey := range refKeys {
		k.deleteRefByKey(ctx, refKey, stake.ID)
	}
	return nil
}

func getStakeRefKeys(stake *types.Stake) ([][]byte, error) {
	owner, err := sdk.AccAddressFromBech32(stake.Owner)
	if err != nil {
		return nil, err
	}

	refKeys := [][]byte{}
	refKeys = append(refKeys, types.KeyPrefixStakeIndex)
	refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixStakeIndexAccount, owner))

	for _, coin := range stake.Coins {
		depositDenom, err := dextypes.NewDepositDenomFromString(coin.Denom)
		if err != nil {
			panic("Only valid LP tokens should be staked")
		}
		denomBz := []byte(coin.Denom)
		pairIdBz := []byte(depositDenom.PairID.Stringify())
		tickBz := dextypes.TickIndexToBytes(depositDenom.Tick, depositDenom.PairID, depositDenom.PairID.Token1)
		refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixStakeIndexDenom, denomBz))
		refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixStakeIndexPairTick, pairIdBz, tickBz))
		refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixStakeIndexAccountDenom, owner, denomBz))
		refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixStakeIndexPairTimestamp, pairIdBz, types.GetTimeKey(stake.StartTime)))
	}

	return refKeys, nil
}
