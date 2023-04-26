package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
)

// NewStake returns a new instance of period stake.
func NewStake(id uint64, owner sdk.AccAddress, coins sdk.Coins, startTime time.Time) *Stake {
	return &Stake{
		ID:        id,
		Owner:     owner.String(),
		Coins:     coins,
		StartTime: startTime,
	}
}

// OwnerAddress returns stakes owner address.
func (p Stake) OwnerAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(p.Owner)
	if err != nil {
		panic(err)
	}
	return addr
}

func (p Stake) SingleCoin() (sdk.Coin, error) {
	if len(p.Coins) != 1 {
		return sdk.Coin{}, fmt.Errorf("Stake %d has no single coin: %s", p.ID, p.Coins)
	}
	return p.Coins[0], nil
}

func (p Stake) ValidateBasic() error {
	for _, coin := range p.Coins {
		_, err := dextypes.NewDepositDenomFromString(coin.Denom)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p Stake) CoinsPassingQueryCondition(distrTo QueryCondition) sdk.Coins {
	coins := p.Coins
	switch len(p.Coins) {
	case 0:
		return nil

	case 1:
		coin := coins[0]
		if !distrTo.Test(coin.Denom) {
			return nil
		}
		return sdk.Coins{coin}

	default:
		// Binary search the amount of coins remaining
		denomPrefix := dextypes.DepositDenomPairIDPrefix(distrTo.PairID.Token0, distrTo.PairID.Token1)
		coins := sdk.Coins{}
		low := 0
		high := len(coins)
		for low < high {
			mid := low + ((high - low) / 2)
			coin := coins[mid]
			switch {
			case distrTo.Test(coin.Denom):
				coins = coins.Add(coin)

				midLeft := mid - 1
				for 0 <= midLeft {
					coin = coins[midLeft]
					if !distrTo.Test(coin.Denom) {
						break
					}
					coins = coins.Add(coin)
					midLeft--
				}

				midRight := mid + 1
				for midRight < len(coins) {
					coin = coins[midRight]
					if !distrTo.Test(coin.Denom) {
						break
					}
					coins = coins.Add(coin)
					midRight++
				}

				return coins
			case denomPrefix < coin.Denom:
				high = mid
			default:
				low = mid
			}
		}
		return nil
	}
}
