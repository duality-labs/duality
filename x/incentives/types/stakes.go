package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Stakes []*Stake

// JCP TODO: remove?
// func (stakes Stakes) CoinsByQueryCondition(distrTo QueryCondition) sdk.Coins {
// 	coins := sdk.Coins{}
// 	for _, stake := range stakes {
// 		coinsToAdd := stake.CoinsPassingQueryCondition(distrTo)
// 		if !coinsToAdd.Empty() {
// 			coins = coins.Add(coinsToAdd...)
// 		}
// 	}
// 	return coins
// }

func (stakes Stakes) GetCoins() sdk.Coins {
	coins := sdk.Coins{}
	for _, stake := range stakes {
		coinsToAdd := stake.GetCoins()
		if !coinsToAdd.Empty() {
			coins = coins.Add(coinsToAdd...)
		}
	}
	return coins
}
