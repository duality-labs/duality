package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgSwap_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSwap
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgSwap{
				Creator:  "invalid_address",
				Receiver: sample.AccAddress(),
				TokenA:   "TokenA",
				TokenB:   "TokenB",
				AmountIn: sdk.OneInt(),
				TokenIn:  "TokenA",
				MinOut:   sdk.ZeroInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgSwap{
				Creator:  sample.AccAddress(),
				Receiver: "invalid address",
				TokenA:   "TokenA",
				TokenB:   "TokenB",
				AmountIn: sdk.OneInt(),
				TokenIn:  "TokenA",
				MinOut:   sdk.ZeroInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid token in",
			msg: MsgSwap{
				Creator:    sample.AccAddress(),
				Receiver:   sample.AccAddress(),
				TokenA:     "TokenA",
				TokenB:     "TokenB",
				AmountIn:   sdk.OneInt(),
				TokenIn:    "TokenC",
				MinOut:     sdk.ZeroInt(),
				LimitPrice: sdk.ZeroDec(),
			},
			err: ErrInvalidTokenIn,
		},
		{
			name: "invalid zero swap",
			msg: MsgSwap{
				Creator:    sample.AccAddress(),
				Receiver:   sample.AccAddress(),
				TokenA:     "TokenA",
				TokenB:     "TokenB",
				AmountIn:   sdk.ZeroInt(),
				TokenIn:    "TokenA",
				MinOut:     sdk.ZeroInt(),
				LimitPrice: sdk.ZeroDec(),
			},
			err: ErrZeroSwap,
		},
		{
			name: "invalid negative minOut",
			msg: MsgSwap{
				Creator:    sample.AccAddress(),
				Receiver:   sample.AccAddress(),
				TokenA:     "TokenA",
				TokenB:     "TokenB",
				AmountIn:   sdk.OneInt(),
				TokenIn:    "TokenA",
				MinOut:     sdk.NewInt(-1),
				LimitPrice: sdk.ZeroDec(),
			},
			err: ErrNegativeMinOut,
		},
		{
			name: "invalid zero limitPrice",
			msg: MsgSwap{
				Creator:    sample.AccAddress(),
				Receiver:   sample.AccAddress(),
				TokenA:     "TokenA",
				TokenB:     "TokenB",
				AmountIn:   sdk.OneInt(),
				TokenIn:    "TokenA",
				MinOut:     sdk.ZeroInt(),
				LimitPrice: sdk.NewDec(-1),
			},
			err: ErrNegativeLimitPrice,
		},
		{
			name: "valid msg",
			msg: MsgSwap{
				Creator:    sample.AccAddress(),
				Receiver:   sample.AccAddress(),
				TokenA:     "TokenA",
				TokenB:     "TokenB",
				AmountIn:   sdk.OneInt(),
				TokenIn:    "TokenA",
				MinOut:     sdk.ZeroInt(),
				LimitPrice: sdk.ZeroDec(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
