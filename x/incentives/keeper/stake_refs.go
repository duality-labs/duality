package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

// addStakeRefs adds appropriate reference keys preceded by a prefix.
// A prefix indicates whether the stake is unstaking or not.
func (k Keeper) addStakeRefs(ctx sdk.Context, stake *types.Stake) error {
	refKeys, err := k.getStakeRefKeys(ctx, stake)
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
	refKeys, err := k.getStakeRefKeys(ctx, stake)
	if err != nil {
		return err
	}
	for _, refKey := range refKeys {
		err = k.deleteRefByKey(ctx, refKey, stake.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) getStakeRefKeys(ctx sdk.Context, stake *types.Stake) ([][]byte, error) {
	owner, err := sdk.AccAddressFromBech32(stake.Owner)
	if err != nil {
		return nil, err
	}
	nKeys := 2 + 4*len(stake.Coins)
	refKeys := make([]string, 0, nKeys)
	refKeys = append(refKeys, string(types.KeyPrefixStakeIndex))
	refKeys = append(refKeys, string(types.CombineKeys(types.KeyPrefixStakeIndexAccount, owner)))

	for _, coin := range stake.Coins {
		poolMetadata, err := k.dk.GetPoolMetadataByDenom(ctx, coin.Denom)
		if err != nil {
			panic("Only valid LP tokens should be staked")
		}
		denomBz := []byte(coin.Denom)
		pairIDBz := []byte(depositDenom.PairID.Stringify())
		tickBz := dextypes.TickIndexToBytes(
			depositDenom.Tick,
			depositDenom.PairID,
			depositDenom.PairID.Token1,
		)
		refKeys = append(refKeys, string(types.CombineKeys(types.KeyPrefixStakeIndexDenom, denomBz)))
		refKeys = append(refKeys, string(types.CombineKeys(types.KeyPrefixStakeIndexPairTick, pairIDBz, tickBz)))
		refKeys = append(refKeys, string(types.CombineKeys(types.KeyPrefixStakeIndexAccountDenom, owner, denomBz)))
		refKeys = append(refKeys, string(types.CombineKeys(
			types.KeyPrefixStakeIndexPairTimestamp,
			pairIDBz,
			types.GetTimeKey(stake.StartTime),
		)))
	}

	// Since we might end up with duplicate refkeys we need to de-dupe the list
	uniqueRefKeyBytes := make([][]byte, 0, len(refKeys))
	seen := make(map[string]bool)
	for _, k := range refKeys {
		if !seen[k] {
			seen[k] = true
			uniqueRefKeyBytes = append(uniqueRefKeyBytes, []byte(k))
		}
	}
	return uniqueRefKeyBytes, nil
}
