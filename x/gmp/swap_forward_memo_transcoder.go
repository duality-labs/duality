package gmp

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	ibcswaptypes "github.com/duality-labs/duality/x/ibcswap/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

const SwapForwardABI = `[
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "creator",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "receiver",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "tokenIn",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "tokenOut",
        "type": "string"
      },
      {
        "internalType": "int64",
        "name": "tickIndex",
        "type": "int64"
      },
      {
        "internalType": "uint256",
        "name": "amountIn",
        "type": "uint256"
      },
      {
        "internalType": "enum Interface.LimitOrderType",
        "name": "orderType",
        "type": "uint8"
      },
      {
        "internalType": "uint256",
        "name": "expirationTime",
        "type": "uint256"
      },
      {
        "internalType": "bool",
        "name": "nonRefundable",
        "type": "bool"
      },
      {
        "internalType": "string",
        "name": "refundAddress",
        "type": "string"
      },
			{
        "internalType": "string",
        "name": "nextArgs",
        "type": "bytes"
      }
    ],
    "name": "swapAndForward",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`
const SwapForwardFunctionName = "swapAndForward"

type SwapForwardArgs struct {
	Creator   string
	Receiver  string
	TokenIn   string
	TokenOut  string
	TickIndex int64
	AmountIn  *big.Int
	OrderType uint8
	// expirationTime is only valid iff orderType == GOOD_TIL_TIME.
	ExpirationTime *big.Int
	NonRefundable  bool
	RefundAddress  string
	NextArgs       []byte
}

type SwapForwardMemoTranscoder struct {
	abi abi.ABI
}

func NewSwapForwardMemoTranscoder() SwapForwardMemoTranscoder {
	abi, err := abi.JSON(strings.NewReader(SwapForwardABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse contract ABI: %v", err))
	}
	return SwapForwardMemoTranscoder{
		abi: abi,
	}
}

func (g SwapForwardMemoTranscoder) Process(payload []byte) ([]byte, error) {
	args, err := g.UnmarshalABI(payload)
	if err != nil {
		return nil, err
	}

	memo, err := g.MapArgsToMemo(args)
	if err != nil {
		return nil, err
	}

	jsonBz, err := g.MarshallJSON(memo)
	if err != nil {
		return nil, err
	}
	return jsonBz, err
}

func (g SwapForwardMemoTranscoder) UnmarshalABI(payload []byte) (*SwapForwardArgs, error) {
	// trim first 4 bytes because these are the first 4 bytes of the hash of the function signature
	// maybe we should assert on these [152, 67, 63, 112]
	payload = payload[4:]

	// Unpack the ABI-encoded data into Go variables.
	inputArgs := g.abi.Methods[SwapForwardFunctionName].Inputs
	unpacked, err := inputArgs.Unpack(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack memo data: %v", err)
	}

	args := &SwapForwardArgs{}
	err = inputArgs.Copy(args, unpacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack memo data: %v", err)
	}

	return args, nil
}

func (g SwapForwardMemoTranscoder) MapArgsToMemo(args *SwapForwardArgs) (*ibcswaptypes.PacketMetadata, error) {
	// Convert big.Int to int64
	if !args.ExpirationTime.IsInt64() {
		return nil, fmt.Errorf("big.Int value is too large to be represented as an int64")
	}
	timestamp := args.ExpirationTime.Int64()

	// Convert Unix timestamp to time.Time
	t := time.Unix(timestamp, 0)

	var nextMemoJSON *ibcswaptypes.JSONObject
	if len(args.NextArgs) > 0 {
		nextMemoJSON = new(ibcswaptypes.JSONObject)
		err := nextMemoJSON.UnmarshalJSON(args.NextArgs)
		if err != nil {
			return nil, fmt.Errorf("could not decode nextMemo")
		}
	} else {
		nextMemoJSON = new(ibcswaptypes.JSONObject)
	}

	metadata := &ibcswaptypes.PacketMetadata{
		Swap: &ibcswaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &dextypes.MsgPlaceLimitOrder{
				Creator:        args.Creator,
				Receiver:       args.Receiver,
				TokenIn:        args.TokenIn,
				TokenOut:       args.TokenOut,
				TickIndex:      args.TickIndex,
				AmountIn:       sdk.NewIntFromBigInt(args.AmountIn),
				OrderType:      dextypes.LimitOrderType(args.OrderType),
				ExpirationTime: &t,
			},
			NonRefundable: args.NonRefundable,
			RefundAddress: args.RefundAddress,
			Next:          nextMemoJSON,
		},
	}
	return metadata, nil
}

func (g SwapForwardMemoTranscoder) MarshallJSON(metadata *ibcswaptypes.PacketMetadata) ([]byte, error) {
	memoBz, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	return memoBz, nil
}
