package types

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

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
			Next: "",
		},
	}
	_, err := json.Marshal(pm)
	require.NoError(t, err)
}

func TestPacketMetadata_Unmarshal(t *testing.T) {
	metadata := "{\n  \"swap\": {\n    \"creator\": \"test-1\",\n    \"receiver\": \"test-1\",\n    \"tokenA\": \"token-a\",\n    \"tokenB\": \"token-b\",\n    \"amountIn\": \"123\",\n    \"tokenIn\": \"token-in\",\n    \"minOut\": \"456\",\n    \"next\": \"\"\n  }\n}"
	pm := &PacketMetadata{}
	err := json.Unmarshal([]byte(metadata), pm)
	require.NoError(t, err)
}

func TestSwapMetadata_ValidatePass(t *testing.T) {
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
			Next: "",
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
			Next: "",
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
			Next: "",
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
			Next: "",
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
			Next: "",
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
			Next: "",
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
			Next: "",
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
			Next: "",
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
			Next: "",
		},
	}
	_, err = json.Marshal(pm)
	require.NoError(t, err)
	require.Error(t, pm.Swap.Validate())
}
