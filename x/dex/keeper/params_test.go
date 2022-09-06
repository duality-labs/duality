package keeper_test

import (
	"testing"

	testkeeper "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DexKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
