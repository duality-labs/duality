package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
	dextypes "github.com/duality-labs/duality/x/dex/keeper"
)

// NewLock returns a new instance of period lock.
func NewLock(ID uint64, owner sdk.AccAddress, duration time.Duration, endTime time.Time, coins sdk.Coins) *Lock {
	return &Lock{
		ID:       ID,
		Owner:    owner.String(),
		Duration: duration,
		EndTime:  endTime,
		Coins:    coins,
	}
}

// IsUnlocking returns lock started unlocking already.
func (p Lock) IsUnlocking() bool {
	return !p.EndTime.Equal(time.Time{})
}

// OwnerAddress returns locks owner address.
func (p Lock) OwnerAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(p.Owner)
	if err != nil {
		panic(err)
	}
	return addr
}

func (p Lock) SingleCoin() (sdk.Coin, error) {
	if len(p.Coins) != 1 {
		return sdk.Coin{}, fmt.Errorf("Lock %d has no single coin: %s", p.ID, p.Coins)
	}
	return p.Coins[0], nil
}

func (p Lock) ValidateBasic() error {
	for _, coin := range p.Coins {
		_, err := dextypes.NewDepositDenomFromString(coin.Denom)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p Lock) CoinsPassingQueryCondition(distrTo QueryCondition) sdk.Coins {
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
		denomPrefix := dexkeeper.DepositDenomPairIDPrefix(distrTo.PairID.Token0, distrTo.PairID.Token1)
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
					midLeft -= 1
				}

				midRight := mid + 1
				for midRight < len(coins) {
					coin = coins[midRight]
					if !distrTo.Test(coin.Denom) {
						break
					}
					coins = coins.Add(coin)
					midRight += 1
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
