package keeper_test

import	(
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"github.com/NicholasDotSol/duality/x/dex/types"
	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestValidateTokensInOrder(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	token0, token1, err := keeper.ValidateTokens("TokenA", "TokenZ")
	require.Nil(t, err)
	assert.Equal(t, "TokenA", token0)
	assert.Equal(t, "TokenZ", token1)
}

func TestValidateTokensNotInOrder(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	token0, token1, err := keeper.ValidateTokens("TokenZ", "TokenA")
	require.Nil(t, err)
	assert.Equal(t, "TokenA", token0)
	assert.Equal(t, "TokenZ", token1)
}

func TestValidateTokensEqual(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	_, _, err := keeper.ValidateTokens("TokenA", "TokenA")
	require.ErrorIs(t, err, types.ErrInvalidTokenPair)
}

func TestValidateAddressValid(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	addr := "cosmos1qpa3plxxqytksex6scg7lxaxa5aytyrwmd0wd0"
	newAddr, err := keeper.ValidateAddress(addr, "creator")
	require.Nil(t, err)
	assert.Equal(t, addr, newAddr.String())
}

func TestValidateAddressInvalid(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	_, err := keeper.ValidateAddress("NotAddr", "creator")
	require.ErrorIs(t, sdkerrors.ErrInvalidAddress, err)
}

func TestValidateCore(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	addr := "cosmos1qpa3plxxqytksex6scg7lxaxa5aytyrwmd0wd0"
	message := &types.MsgWithdrawl{
		Creator:        addr,
		Receiver:       addr,
		TokenA:         "TokenB",
		TokenB:         "TokenA",
	}
	token0, token1, creatorAddr, receiverAddr, err := keeper.ValidateCore(message)
	require.Nil(t, err)
	assert.Equal(t, "TokenA", token0)
	assert.Equal(t, "TokenB", token1)
	assert.Equal(t, addr, creatorAddr.String())
	assert.Equal(t, addr, receiverAddr.String())
}

func TestValidateCoreTokensEqual(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	addr := "cosmos1qpa3plxxqytksex6scg7lxaxa5aytyrwmd0wd0"
	message := &types.MsgWithdrawl{
		Creator:        addr,
		Receiver:       addr,
		TokenA:         "TokenA",
		TokenB:         "TokenA",
	}
	_, _, _, _, err := keeper.ValidateCore(message)
	require.ErrorIs(t, err, types.ErrInvalidTokenPair)
}

func TestValidateCoreBadCreator(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	addr := "cosmos1qpa3plxxqytksex6scg7lxaxa5aytyrwmd0wd0"
	message := &types.MsgWithdrawl{
		Creator:        "Bad",
		Receiver:       addr,
		TokenA:         "TokenA",
		TokenB:         "TokenB",
	}
	_, _, _, _, err := keeper.ValidateCore(message)
	require.ErrorIs(t, sdkerrors.ErrInvalidAddress, err)
}

func TestValidateCoreBadReceiver(t *testing.T){
	keeper, _ := keepertest.DexKeeper(t)
	addr := "cosmos1qpa3plxxqytksex6scg7lxaxa5aytyrwmd0wd0"
	message := &types.MsgWithdrawl{
		Creator:        addr,
		Receiver:       "bad",
		TokenA:         "TokenA",
		TokenB:         "TokenB",
	}
	_, _, _, _, err := keeper.ValidateCore(message)
	require.ErrorIs(t, sdkerrors.ErrInvalidAddress, err)
}
