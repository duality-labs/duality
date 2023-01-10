package types

import (
	"testing"

	"github.com/NicholasDotSol/duality/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				TokenIn:   "TokenA",
				AmountIn:  sdk.OneInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgPlaceLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  "invalid_address",
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				TokenIn:   "TokenA",
				AmountIn:  sdk.OneInt(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid tokenIn",
			msg: MsgPlaceLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  sample.AccAddress(),
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				TokenIn:   "TokenC",
				AmountIn:  sdk.OneInt(),
			},
			err: ErrInvalidTradingPair,
		}, {
			name: "valid msg",
			msg: MsgPlaceLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  sample.AccAddress(),
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				TokenIn:   "TokenA",
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
