package keeper_test

import	(
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
)

func TestValidateTokensInOrder(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	token0, token1, err := keeper.validateTokens("TokenA", "TokenZ")
	require.Nil(err)
	assert.Equal(t, "TokenA", token0)
	assert.Equal(t, "TokenB", token0)

}
