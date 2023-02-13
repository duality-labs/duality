package types

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/iancoleman/orderedmap"
	forwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v3/router/types"
	"github.com/stretchr/testify/require"
)

// TestPacketMetadata_Marshal asserts that the marshaling of the swap metadata works as intended.
func TestPacketMetadata_Marshal(t *testing.T) {
	pm := PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "test-1",
				Receiver: "test-1",
				TokenA:   "token-a",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(123),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(456),
			},
			Next: nil,
		},
	}
	_, err := json.Marshal(pm)
	require.NoError(t, err)
}

// TestPacketMetadata_MarshalWithNext asserts that the marshaling of the swap metadata works as intended with next field initialized.
func TestPacketMetadata_MarshalWithNext(t *testing.T) {
	forwardMedata := &forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: "cosmos14zde8usc4ur04y3aqnufzzmv2uqdpwwttr5uwv",
			Port:     "transfer",
			Channel:  "channel-0",
			Timeout:  0,
			Retries:  nil,
			Next:     nil,
		}}
	nextBz, err := json.Marshal(forwardMedata)
	require.NoError(t, err)

	pm := PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "test-1",
				Receiver: "test-1",
				TokenA:   "token-a",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(123),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(456),
			},
			Next: NewJSONObject(false, nextBz, orderedmap.OrderedMap{}),
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
}

// TestPacketMetadata_Unmarshal asserts that unmarshaling works as intended.
func TestPacketMetadata_Unmarshal(t *testing.T) {
	metadata := "{\n  \"swap\": {\n    \"creator\": \"test-1\",\n    \"receiver\": \"test-1\",\n    \"tokenA\": \"token-a\",\n    \"tokenB\": \"token-b\",\n    \"amountIn\": \"123\",\n    \"tokenIn\": \"token-in\",\n    \"minOut\": \"456\",\n    \"next\": \"\"\n  }\n}"
	pm := &PacketMetadata{}
	err := json.Unmarshal([]byte(metadata), pm)
	require.NoError(t, err)
}

// TestPacketMetadata_UnmarshalStringNext asserts that unmarshaling works as intended when next is escaped json string.
func TestPacketMetadata_UnmarshalStringNext(t *testing.T) {
	metadata := "{\n  \"swap\": {\n    \"creator\": \"test-1\",\n    \"receiver\": \"test-1\",\n    \"tokenA\": \"token-a\",\n    \"tokenB\": \"token-b\",\n    \"amountIn\": \"123\",\n    \"tokenIn\": \"token-in\",\n    \"minOut\": \"456\",\n    \"next\": \" {\\\"forward\\\":{\\\"receiver\\\":\\\"cosmos1f4cur2krsua2th9kkp7n0zje4stea4p9tu70u8\\\",\\\"port\\\":\\\"transfer\\\",\\\"channel\\\":\\\"channel-0\\\",\\\"timeout\\\":0,\\\"next\\\":{\\\"forward\\\":{\\\"receiver\\\":\\\"cosmos1l505zhahp24v5jsmps9vs5asah759fdce06sfp\\\",\\\"port\\\":\\\"transfer\\\",\\\"channel\\\":\\\"channel-0\\\",\\\"timeout\\\":0}}}}\"\n  }\n}"
	pm := &PacketMetadata{}
	err := json.Unmarshal([]byte(metadata), pm)
	require.NoError(t, err)
}

// TestPacketMetadata_UnmarshalJSONNext asserts that unmarshaling works as intended when next is a raw json object.
func TestPacketMetadata_UnmarshalJSONNext(t *testing.T) {
	metadata := "{\"swap\":{\"creator\":\"test-1\",\"receiver\":\"test-1\",\"tokenA\":\"token-a\",\"tokenB\":\"token-b\",\"amountIn\":\"123\",\"tokenIn\":\"token-in\",\"minOut\":\"456\",\"tickLimit\":\"0\",\"next\":{\"forward\":{\"receiver\":\"cosmos14zde8usc4ur04y3aqnufzzmv2uqdpwwttr5uwv\",\"port\":\"transfer\",\"channel\":\"channel-0\"}}}}"
	pm := &PacketMetadata{}
	err := json.Unmarshal([]byte(metadata), pm)
	require.NoError(t, err)
}

func TestSwapMetadata_ValidatePass(t *testing.T) {
	pm := PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:    "cosmos14zde8usc4ur04y3aqnufzzmv2uqdpwwttr5uwv",
				Receiver:   "cosmos14zde8usc4ur04y3aqnufzzmv2uqdpwwttr5uwv",
				TokenA:     "token-a",
				TokenB:     "token-b",
				AmountIn:   sdk.NewInt(123),
				TokenIn:    "token-a",
				MinOut:     sdk.NewInt(456),
				LimitPrice: sdk.NewDec(789),
			},
			Next: nil,
		},
	}
	_, err := json.Marshal(pm)
	require.NoError(t, err)

	require.NoError(t, pm.Swap.Validate())
}

func TestSwapMetadata_ValidateFail(t *testing.T) {
	pm := PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "",
				Receiver: "test-1",
				TokenA:   "token-a",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(123),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(456),
			},
			Next: nil,
		},
	}
	_, err := json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())

	pm = PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "creator",
				Receiver: "",
				TokenA:   "token-a",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(123),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(456),
			},
			Next: nil,
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())

	pm = PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "creator",
				Receiver: "test-1",
				TokenA:   "",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(123),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(456),
			},
			Next: nil,
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())

	pm = PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "creator",
				Receiver: "receiver",
				TokenA:   "token-a",
				TokenB:   "",
				AmountIn: sdk.NewInt(123),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(456),
			},
			Next: nil,
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())

	pm = PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "creator",
				Receiver: "receiver",
				TokenA:   "token-a",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(0),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(456),
			},
			Next: nil,
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())

	pm = PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "creator",
				Receiver: "receiver",
				TokenA:   "token-a",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(-1),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(456),
			},
			Next: nil,
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())

	pm = PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "creator",
				Receiver: "receiver",
				TokenA:   "token-a",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(123),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(0),
			},
			Next: nil,
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())

	pm = PacketMetadata{
		&SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  "creator",
				Receiver: "receiver",
				TokenA:   "token-a",
				TokenB:   "token-b",
				AmountIn: sdk.NewInt(123),
				TokenIn:  "token-in",
				MinOut:   sdk.NewInt(-1),
			},
			Next: nil,
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())
}