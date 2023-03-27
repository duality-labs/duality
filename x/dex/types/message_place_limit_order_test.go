package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgPlaceLimitOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPlaceLimitOrder
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgPlaceLimitOrder{
				Creator:   "invalid_address",
				Receiver:  sample.AccAddress(),
				TokenIn:   "TokenA",
				TokenOut:  "TokenB",
				TickIndex: 0,
				AmountIn:  sdk.OneInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgPlaceLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  "invalid_address",
				TokenIn:   "TokenA",
				TokenOut:  "TokenB",
				TickIndex: 0,
				AmountIn:  sdk.OneInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid zero limit order",
			msg: MsgPlaceLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  sample.AccAddress(),
				TokenIn:   "TokenA",
				TokenOut:  "TokenB",
				TickIndex: 0,
				AmountIn:  sdk.ZeroInt(),
			},
			err: ErrZeroLimitOrder,
		}, {
			name: "valid msg",
			msg: MsgPlaceLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  sample.AccAddress(),
				TokenIn:   "TokenA",
				TokenOut:  "TokenB",
				TickIndex: 0,
				AmountIn:  sdk.OneInt(),
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
