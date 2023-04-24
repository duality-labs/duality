package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	. "github.com/duality-labs/duality/x/dex/types"
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
				Creator:     "invalid_address",
				Receiver:    sample.AccAddress(),
				TokenIn:     "TokenA",
				TokenOut:    "TokenB",
				MaxAmountIn: sdk.OneInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid receiver",
			msg: MsgSwap{
				Creator:     sample.AccAddress(),
				Receiver:    "invalid address",
				TokenIn:     "TokenA",
				TokenOut:    "TokenB",
				MaxAmountIn: sdk.OneInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid zero swap",
			msg: MsgSwap{
				Creator:     sample.AccAddress(),
				Receiver:    sample.AccAddress(),
				TokenIn:     "TokenA",
				TokenOut:    "TokenB",
				MaxAmountIn: sdk.ZeroInt(),
			},
			err: ErrZeroSwap,
		},
		{
			name: "invalid negative maxAmountOut",
			msg: MsgSwap{
				Creator:      sample.AccAddress(),
				Receiver:     sample.AccAddress(),
				TokenIn:      "TokenA",
				TokenOut:     "TokenB",
				MaxAmountIn:  sdk.OneInt(),
				MaxAmountOut: sdk.NewInt(-1),
			},
			err: ErrNegativeMaxAmountOut,
		},
		{
			name: "valid msg",
			msg: MsgSwap{
				Creator:     sample.AccAddress(),
				Receiver:    sample.AccAddress(),
				TokenIn:     "TokenA",
				TokenOut:    "TokenB",
				MaxAmountIn: sdk.OneInt(),
			},
		},
		{
			name: "valid msg with maxAmountOut",
			msg: MsgSwap{
				Creator:      sample.AccAddress(),
				Receiver:     sample.AccAddress(),
				TokenIn:      "TokenA",
				TokenOut:     "TokenB",
				MaxAmountIn:  sdk.OneInt(),
				MaxAmountOut: sdk.OneInt(),
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
