package gmp_test

import (
	// "strings"

	"encoding/hex"
	"math/big"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	. "github.com/duality-labs/duality/x/gmp"
	ibcswaptypes "github.com/duality-labs/duality/x/ibcswap/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iancoleman/orderedmap"
	"github.com/stretchr/testify/assert"
)

func TestSwapForwardMemoTranscoder_Basic(t *testing.T) {
	// prepare input
	abi, err := abi.JSON(strings.NewReader(SwapForwardABI))
	assert.NoError(t, err)

	timeIn10Minutes := time.Now().Add(10 * time.Minute).Unix()
	args := SwapForwardArgs{
		Creator:        "alice",
		Receiver:       "bob",
		TokenIn:        "foo",
		TokenOut:       "bar",
		TickIndex:      0,
		AmountIn:       big.NewInt(1),
		OrderType:      uint8(dextypes.LimitOrderType_IMMEDIATE_OR_CANCEL),
		ExpirationTime: big.NewInt(timeIn10Minutes),
		NonRefundable:  false,
		RefundAddress:  "alice",
		NextArgs:       []byte("{}"),
	}

	input, err := abi.Pack(
		SwapForwardFunctionName,
		args.Creator,
		args.Receiver,
		args.TokenIn,
		args.TokenOut,
		args.TickIndex,
		args.AmountIn,
		uint8(args.OrderType),
		args.ExpirationTime,
		args.NonRefundable,
		args.RefundAddress,
		args.NextArgs,
	)
	assert.NoError(t, err)

	// prepare expected output
	expectedTime := time.Unix(timeIn10Minutes, 0)
	expectedJSONObject := new(ibcswaptypes.JSONObject)
	expectedJSONObject.UnmarshalJSON([]byte("{}"))
	expected := &ibcswaptypes.PacketMetadata{
		Swap: &ibcswaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &dextypes.MsgPlaceLimitOrder{
				Creator:        "alice",
				Receiver:       "bob",
				TokenIn:        "foo",
				TokenOut:       "bar",
				TickIndex:      0,
				AmountIn:       sdk.NewIntFromUint64(1),
				OrderType:      dextypes.LimitOrderType_IMMEDIATE_OR_CANCEL,
				ExpirationTime: &expectedTime,
			},
			NonRefundable: false,
			RefundAddress: "alice",
			// Next:          ibcswaptypes.NewJSONObject(true, nil, *orderedmap.New()),
			Next: expectedJSONObject,
		},
	}

	// setup subject
	transcoder := NewSwapForwardMemoTranscoder()

	// test for actual
	var argsOut *SwapForwardArgs
	argsOut, err = transcoder.UnmarshalABI(input)
	assert.NoError(t, err)

	actual, err := transcoder.MapArgsToMemo(argsOut)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)

	_, err = transcoder.MarshallJSON(actual)
	assert.NoError(t, err)
}

func TestSwapForwardMemoTranscoder_GroundTruthInput(t *testing.T) {
	// prepare input
	inputBz := `0x98433f70000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026000000000000000000000000000000000000000000000000000000000000002a00000000000000000000000000000000000000000000000000000000000000005616c6963650000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003626f6200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003666f6f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000362617200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005616c69636500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000027b7d000000000000000000000000000000000000000000000000000000000000`
	// trim leading "0x" when decoding
	input, err := hex.DecodeString(inputBz[2:])
	assert.NoError(t, err)

	// prepare expected output
	// timeIn10Minutes := time.Now().Add(10 * time.Minute).Unix()
	expectedTime := time.Unix(0, 0)

	expected := &ibcswaptypes.PacketMetadata{
		Swap: &ibcswaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &dextypes.MsgPlaceLimitOrder{
				Creator:        "alice",
				Receiver:       "bob",
				TokenIn:        "foo",
				TokenOut:       "bar",
				TickIndex:      0,
				AmountIn:       sdk.NewIntFromUint64(100),
				OrderType:      dextypes.LimitOrderType_GOOD_TIL_CANCELLED,
				ExpirationTime: &expectedTime,
			},
			NonRefundable: false,
			RefundAddress: "alice",
			Next:          ibcswaptypes.NewJSONObject(true, []byte{}, *orderedmap.New()),
		},
	}

	// setup subject
	transcoder := NewSwapForwardMemoTranscoder()

	// test for actual
	var argsOut *SwapForwardArgs
	argsOut, err = transcoder.UnmarshalABI(input)
	assert.NoError(t, err)

	actual, err := transcoder.MapArgsToMemo(argsOut)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)

	_, err = transcoder.MarshallJSON(actual)
	assert.NoError(t, err)
}
