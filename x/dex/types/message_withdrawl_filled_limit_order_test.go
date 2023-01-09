package types

import (
	"testing"

	"github.com/NicholasDotSol/duality/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgWithdrawFilledLimitOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgWithdrawFilledLimitOrder
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgWithdrawFilledLimitOrder{
				Creator:   "invalid_address",
				Receiver:  sample.AccAddress(),
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				KeyToken:  "TokenA",
				Key:       0,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgWithdrawFilledLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  "invalid address",
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				KeyToken:  "TokenA",
				Key:       0,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid key token",
			msg: MsgWithdrawFilledLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  sample.AccAddress(),
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				KeyToken:  "TokenC",
				Key:       0,
			},
			err: ErrInvalidTradingPair,
		}, {
			name: "valid msg",
			msg: MsgWithdrawFilledLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  sample.AccAddress(),
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				KeyToken:  "TokenA",
				Key:       0,
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
