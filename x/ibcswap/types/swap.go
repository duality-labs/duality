package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dextypes "github.com/duality-labs/duality/x/dex/types"
)

/*
An example JSON blob that would be marshaled into PacketMetadata where the next field can contain any arbitrary data.

{
  "swap": {
    "creator": "test-1",
    "receiver": "test-1",
    "tokenA": "token-a",
    "tokenB": "token-b",
    "amountIn": "123.000000000000000000",
    "tokenIn": "token-in",
    "minOut": "456.000000000000000000",
	"non-refundable": false,
    "next": ""
  }
}
*/

// PacketMetadata wraps the SwapMetadata. The root key in the incoming ICS20 transfer packet's memo needs to be set to the same
// value as the json tag in order for the swap middleware to process the swap.
type PacketMetadata struct {
	Swap *SwapMetadata `json:"swap"`
}

// SwapMetadata defines the parameters necessary to perform a swap utilizing the memo field from an incoming ICS20
// transfer packet. The next field is a string so that you can nest any arbitrary metadata to be handled
// further in the middleware stack or on the counterparty.
type SwapMetadata struct {
	*dextypes.MsgSwap
	NonRefundable bool   `json:"non-refundable,omitempty"`
	Next          string `json:"next,omitempty"`
}

// Validate ensures that all the required fields are present in the SwapMetadata and contain valid values.
func (sm SwapMetadata) Validate() error {
	if err := sm.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(ErrInvalidSwapMetadata, err.Error())
	}
	if sm.TokenA == "" {
		return sdkerrors.Wrap(ErrInvalidSwapMetadata, "swap tokenA cannot be an empty string")
	}
	if sm.TokenB == "" {
		return sdkerrors.Wrap(ErrInvalidSwapMetadata, "swap tokenB cannot be an empty string")
	}
	if sm.TokenIn == "" {
		return sdkerrors.Wrap(ErrInvalidSwapMetadata, "swap tokenIn cannot be an empty string")
	}
	if sm.AmountIn.IsZero() || sm.AmountIn.IsNil() {
		return sdkerrors.Wrap(ErrInvalidSwapMetadata, "swap amountIn cannot be 0 or nil")
	}
	if sm.AmountIn.IsNegative() {
		return sdkerrors.Wrap(ErrInvalidSwapMetadata, "swap amountIn cannot be negative")
	}
	if sm.MinOut.IsZero() || sm.MinOut.IsNil() {
		return sdkerrors.Wrap(ErrInvalidSwapMetadata, "swap minOut cannot be 0 or nil")
	}
	if sm.MinOut.IsNegative() {
		return sdkerrors.Wrap(ErrInvalidSwapMetadata, "swap minOut cannot be negative")
	}
	return nil
}
