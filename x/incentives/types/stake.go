package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
)

// NewStake returns a new instance of period stake.
func NewStake(
	id uint64,
	owner sdk.AccAddress,
	coins sdk.Coins,
	startTime time.Time,
	startDistEpoch int64,
) *Stake {
	coins = coins.Sort()
	return &Stake{
		ID:             id,
		Owner:          owner.String(),
		Coins:          coins,
		StartTime:      startTime,
		StartDistEpoch: startDistEpoch,
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
		err := dextypes.ValidatePoolDenom(coin.Denom)
		if err != nil {
			return err
		}
	}
	return nil
}
