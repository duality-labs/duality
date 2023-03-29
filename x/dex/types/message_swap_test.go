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
				TokenIn:  "TokenA",
				TokenOut: "TokenB",
				AmountIn: sdk.OneInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgSwap{
				Creator:  sample.AccAddress(),
				Receiver: "invalid address",
				TokenIn:  "TokenA",
				TokenOut: "TokenB",
				AmountIn: sdk.OneInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid zero swap",
			msg: MsgSwap{
				Creator:  sample.AccAddress(),
				Receiver: sample.AccAddress(),
				TokenIn:  "TokenA",
				TokenOut: "TokenB",
				AmountIn: sdk.ZeroInt(),
			},
			err: ErrZeroSwap,
		},
		{
			name: "valid msg",
			msg: MsgSwap{
				Creator:  sample.AccAddress(),
				Receiver: sample.AccAddress(),
				TokenIn:  "TokenA",
				TokenOut: "TokenB",
				AmountIn: sdk.OneInt(),
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
