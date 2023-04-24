package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
	"gotest.tools/assert"
)

func TestPriceMath(t *testing.T) {
	tick := 352437
	amount := sdk.MustNewDecFromStr("1000000000000000000000")
	basePrice := utils.BasePrice()
	expected := amount.Quo(basePrice.Power(uint64(tick))).TruncateInt()
	result := types.MustNewPrice(int64(tick)).Mul(amount).TruncateInt()
	assert.Equal(t, expected.Int64(), result.Int64())
}
