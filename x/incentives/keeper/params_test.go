package keeper_test

import (
	"testing"

	testkeeper "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/x/incentives/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IncentivesKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
