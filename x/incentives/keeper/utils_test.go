package keeper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/duality-labs/duality/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCombineKeys(t *testing.T) {
	// create three keys, each different byte arrays
	key1 := []byte{0x11}
	key2 := []byte{0x12}
	key3 := []byte{0x13}

	// combine the three keys into a single key
	key := types.CombineKeys(key1, key2, key3)

	// three keys plus two separators is equal to a length of 5
	require.Len(t, key, 3+2)

	// ensure the newly created key is made up of the three previous keys (and the two key index separators)
	require.Equal(t, key[0], key1[0])
	require.Equal(t, key[1], types.KeyIndexSeparator[0])
	require.Equal(t, key[2], key2[0])
	require.Equal(t, key[3], types.KeyIndexSeparator[0])
	require.Equal(t, key[4], key3[0])
}

func TestFindIndex(t *testing.T) {
	// create an array of 5 IDs
	IDs := []uint64{1, 2, 3, 4, 5}

	// use the findIndex function to find the index of the respective IDs
	// if it doesn't exist, return -1
	require.Equal(t, findIndex(IDs, 1), 0)
	require.Equal(t, findIndex(IDs, 3), 2)
	require.Equal(t, findIndex(IDs, 5), 4)
	require.Equal(t, findIndex(IDs, 6), -1)
}

func TestRemoveValue(t *testing.T) {
	// create an array of 5 IDs
	IDs := []uint64{1, 2, 3, 4, 5}

	// remove an ID
	// ensure if ID exists, the length is reduced by one and the index of the removed ID is returned
	IDs, index1 := removeValue(IDs, 5)
	require.Len(t, IDs, 4)
	require.Equal(t, index1, 4)
	IDs, index2 := removeValue(IDs, 3)
	require.Len(t, IDs, 3)
	require.Equal(t, index2, 2)
	IDs, index3 := removeValue(IDs, 1)
	require.Len(t, IDs, 2)
	require.Equal(t, index3, 0)
	IDs, index4 := removeValue(IDs, 6)
	require.Len(t, IDs, 2)
	require.Equal(t, index4, -1)
}

func TestLockRefKeys(t *testing.T) {
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	// empty address and 1 coin
	lock1 := types.NewLock(1, sdk.AccAddress{}, time.Second, time.Now(), sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	_, err := unlockingLockRefKeys(lock1)
	require.Error(t, err)

	// empty address and 2 coins
	lock2 := types.NewLock(1, sdk.AccAddress{}, time.Second, time.Now(), sdk.Coins{sdk.NewInt64Coin("stake", 10), sdk.NewInt64Coin("atom", 1)})
	_, err = unlockingLockRefKeys(lock2)
	require.Error(t, err)

	// not empty address and 1 coin
	lock3 := types.NewLock(1, addr1, time.Second, time.Now(), sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	keys3, err := unlockingLockRefKeys(lock3)
	require.Len(t, keys3, 4)

	// not empty address and empty coin
	lock4 := types.NewLock(1, addr1, time.Second, time.Now(), sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	keys4, err := unlockingLockRefKeys(lock4)
	require.Len(t, keys4, 4)
	// not empty address and 2 coins
	lock5 := types.NewLock(1, addr1, time.Second, time.Now(), sdk.Coins{sdk.NewInt64Coin("stake", 10), sdk.NewInt64Coin("atom", 1)})
	keys5, err := unlockingLockRefKeys(lock5)
	require.Len(t, keys5, 6)
}
