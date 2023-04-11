package types_test

import (
	"testing"
	time "time"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/duality-labs/duality/app/apptesting"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
	incentivestypes "github.com/duality-labs/duality/x/incentives/types"
)

// TestMsgCreatePool tests if valid/invalid create pool messages are properly validated/invalidated
func TestMsgCreatePool(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())

	// make a proper createPool message
	createMsg := func(after func(msg incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge {
		distributeTo := types.QueryCondition{
			PairID: &dextypes.PairID{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			StartTick: -10,
			EndTick:   10,
		}

		properMsg := *incentivestypes.NewMsgCreateGauge(
			false,
			addr1,
			distributeTo,
			sdk.Coins{},
			time.Now(),
			2,
		)

		return after(properMsg)
	}

	// validate createPool message was created as intended
	msg := createMsg(func(msg incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge {
		return msg
	})
	require.Equal(t, msg.Route(), incentivestypes.RouterKey)
	require.Equal(t, msg.Type(), "create_gauge")
	signers := msg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1.String())

	tests := []struct {
		name       string
		msg        incentivestypes.MsgCreateGauge
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: createMsg(func(msg incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge {
				return msg
			}),
			expectPass: true,
		},
		{
			name: "empty owner",
			msg: createMsg(func(msg incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge {
				msg.Owner = ""
				return msg
			}),
			expectPass: false,
		},
		// TODO
		{
			name: "invalid distribution start time",
			msg: createMsg(func(msg incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge {
				msg.StartTime = time.Time{}
				return msg
			}),
			expectPass: false,
		},
		{
			name: "invalid num epochs paid over",
			msg: createMsg(func(msg incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge {
				msg.NumEpochsPaidOver = 0
				return msg
			}),
			expectPass: false,
		},
		{
			name: "invalid num epochs paid over for perpetual gauge",
			msg: createMsg(func(msg incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge {
				msg.NumEpochsPaidOver = 2
				msg.IsPerpetual = true
				return msg
			}),
			expectPass: false,
		},
		{
			name: "valid num epochs paid over for perpetual gauge",
			msg: createMsg(func(msg incentivestypes.MsgCreateGauge) incentivestypes.MsgCreateGauge {
				msg.NumEpochsPaidOver = 1
				msg.IsPerpetual = true
				return msg
			}),
			expectPass: true,
		},
	}

	for _, test := range tests {
		if test.expectPass {
			require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
		} else {
			require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
		}
	}
}

// TestMsgAddToGauge tests if valid/invalid add to gauge messages are properly validated/invalidated
func TestMsgAddToGauge(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())

	// make a proper addToGauge message
	createMsg := func(after func(msg incentivestypes.MsgAddToGauge) incentivestypes.MsgAddToGauge) incentivestypes.MsgAddToGauge {
		properMsg := *incentivestypes.NewMsgAddToGauge(
			addr1,
			1,
			sdk.Coins{sdk.NewInt64Coin("stake", 10)},
		)

		return after(properMsg)
	}

	// validate addToGauge message was created as intended
	msg := createMsg(func(msg incentivestypes.MsgAddToGauge) incentivestypes.MsgAddToGauge {
		return msg
	})
	require.Equal(t, msg.Route(), incentivestypes.RouterKey)
	require.Equal(t, msg.Type(), "add_to_gauge")
	signers := msg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1.String())

	tests := []struct {
		name       string
		msg        incentivestypes.MsgAddToGauge
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: createMsg(func(msg incentivestypes.MsgAddToGauge) incentivestypes.MsgAddToGauge {
				return msg
			}),
			expectPass: true,
		},
		{
			name: "empty owner",
			msg: createMsg(func(msg incentivestypes.MsgAddToGauge) incentivestypes.MsgAddToGauge {
				msg.Owner = ""
				return msg
			}),
			expectPass: false,
		},
		{
			name: "empty rewards",
			msg: createMsg(func(msg incentivestypes.MsgAddToGauge) incentivestypes.MsgAddToGauge {
				msg.Rewards = sdk.Coins{}
				return msg
			}),
			expectPass: false,
		},
	}

	for _, test := range tests {
		if test.expectPass {
			require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
		} else {
			require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
		}
	}
}

func TestMsgSetupLock(t *testing.T) {
	addr1, invalidAddr := apptesting.GenerateTestAddrs()

	tests := []struct {
		name       string
		msg        types.MsgLockTokens
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: types.MsgLockTokens{
				Owner: addr1,
				Coins: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			},
			expectPass: true,
		},
		{
			name: "invalid owner",
			msg: types.MsgLockTokens{
				Owner: invalidAddr,
				Coins: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			},
		},
		{
			name: "invalid coin length",
			msg: types.MsgLockTokens{
				Owner: addr1,
				Coins: sdk.NewCoins(sdk.NewCoin("test1", sdk.NewInt(100000)), sdk.NewCoin("test2", sdk.NewInt(100000))),
			},
		},
		{
			name: "zero token amount",
			msg: types.MsgLockTokens{
				Owner: addr1,
				Coins: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(0))),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expectPass {
				require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
				require.Equal(t, test.msg.Route(), types.RouterKey)
				require.Equal(t, test.msg.Type(), "lock_tokens")
				signers := test.msg.GetSigners()
				require.Equal(t, len(signers), 1)
				require.Equal(t, signers[0].String(), addr1)
			} else {
				require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
			}
		})
	}
}

func TestMsgBeginUnlockingAll(t *testing.T) {
	addr1, invalidAddr := apptesting.GenerateTestAddrs()

	tests := []struct {
		name       string
		msg        types.MsgBeginUnlockingAll
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: types.MsgBeginUnlockingAll{
				Owner: addr1,
			},
			expectPass: true,
		},
		{
			name: "invalid owner",
			msg: types.MsgBeginUnlockingAll{
				Owner: invalidAddr,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expectPass {
				require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
				require.Equal(t, test.msg.Route(), types.RouterKey)
				require.Equal(t, test.msg.Type(), "begin_unlocking_all")
				signers := test.msg.GetSigners()
				require.Equal(t, len(signers), 1)
				require.Equal(t, signers[0].String(), addr1)
			} else {
				require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
			}
		})
	}
}

func TestMsgBeginUnlocking(t *testing.T) {
	addr1, invalidAddr := apptesting.GenerateTestAddrs()

	tests := []struct {
		name       string
		msg        types.MsgBeginUnlocking
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: types.MsgBeginUnlocking{
				Owner: addr1,
				ID:    1,
				Coins: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			},
			expectPass: true,
		},
		{
			name: "invalid owner",
			msg: types.MsgBeginUnlocking{
				Owner: invalidAddr,
				ID:    1,
				Coins: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			},
		},
		{
			name: "invalid lockup ID",
			msg: types.MsgBeginUnlocking{
				Owner: addr1,
				ID:    0,
				Coins: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			},
		},
		{
			name: "invalid coins length",
			msg: types.MsgBeginUnlocking{
				Owner: addr1,
				ID:    1,
				Coins: sdk.NewCoins(sdk.NewCoin("test1", sdk.NewInt(100000)), sdk.NewCoin("test2", sdk.NewInt(100000))),
			},
		},
		{
			name: "zero coins (same as nil)",
			msg: types.MsgBeginUnlocking{
				Owner: addr1,
				ID:    1,
				Coins: sdk.NewCoins(sdk.NewCoin("test1", sdk.NewInt(0))),
			},
			expectPass: true,
		},
		{
			name: "nil coins (unlock by ID)",
			msg: types.MsgBeginUnlocking{
				Owner: addr1,
				ID:    1,
				Coins: sdk.NewCoins(),
			},
			expectPass: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expectPass {
				require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
				require.Equal(t, test.msg.Route(), types.RouterKey)
				require.Equal(t, test.msg.Type(), "begin_unlocking")
				signers := test.msg.GetSigners()
				require.Equal(t, len(signers), 1)
				require.Equal(t, signers[0].String(), addr1)
			} else {
				require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
			}
		})
	}
}