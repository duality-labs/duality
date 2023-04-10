package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Locks []*Lock

func (locks Locks) CoinsByQueryCondition(distrTo QueryCondition) sdk.Coins {
	coins := sdk.Coins{}
	for _, lock := range locks {
		coinsToAdd := lock.CoinsPassingQueryCondition(distrTo)
		if !coinsToAdd.Empty() {
			coins = coins.Add(coinsToAdd...)
		}
	}
	return coins
}

func (locks Locks) GetCoins() sdk.Coins {
	coins := sdk.Coins{}
	for _, lock := range locks {
		coinsToAdd := lock.GetCoins()
		if !coinsToAdd.Empty() {
			coins = coins.Add(coinsToAdd...)
		}
	}
	return coins
}
