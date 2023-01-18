package types_test

import (
	"testing"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

var ZeroInt sdk.Int = sdk.NewInt(0)
var TenInt sdk.Int = sdk.NewInt(10)

func TestHasTokenEmptyReserves(t *testing.T) {

	// WHEN has no reserves
	tick := types.TickLiquidity{LPReserve: &ZeroInt}
	assert.False(t, tick.HasToken())
}

func TestHasTokenEmptyLO(t *testing.T) {

	// WHEN has no limits orders
	tick := types.TickLiquidity{
		LimitOrderTranche: &types.LimitOrderTranche{
			ReservesTokenIn: sdk.NewInt(0),
		},
	}
	assert.False(t, tick.HasToken())
}

func TestHasToken0HasReserves(t *testing.T) {

	// WHEN tick has Reserves
	ten := sdk.NewInt(10)
	tick := types.TickLiquidity{LPReserve: &ten}

	assert.True(t, tick.HasToken())
}

func TestHasTokenHasLO(t *testing.T) {

	// WHEN has limit ordeers
	tick := types.TickLiquidity{
		LimitOrderTranche: &types.LimitOrderTranche{
			ReservesTokenIn: sdk.NewInt(10),
		},
	}
	assert.True(t, tick.HasToken())
}
