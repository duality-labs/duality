package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCancelLimitOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelLimitOrder
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgCancelLimitOrder{
				Creator:    "invalid_address",
				TokenIn:    "TokenA",
				TokenOut:   "TokenB",
				TickIndex:  0,
				TrancheKey: "0",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid msg",
			msg: MsgCancelLimitOrder{
				Creator:    sample.AccAddress(),
				TokenIn:    "TokenA",
				TokenOut:   "TokenB",
				TickIndex:  0,
				TrancheKey: "0",
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
