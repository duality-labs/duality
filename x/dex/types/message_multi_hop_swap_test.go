package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	. "github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestMsgMultiHopSwap_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMultiHopSwap
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgMultiHopSwap{
				Creator:  "invalid_address",
				Receiver: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid receiver address",
			msg: MsgMultiHopSwap{
				Creator:  sample.AccAddress(),
				Receiver: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "missing route",
			msg: MsgMultiHopSwap{
				Creator:  sample.AccAddress(),
				Receiver: sample.AccAddress(),
			},
			err: ErrMissingMultihopRoute,
		},
		{
			name: "invalid exit tokens",
			msg: MsgMultiHopSwap{
				Creator:  sample.AccAddress(),
				Receiver: sample.AccAddress(),
				Routes: []*MultiHopRoute{
					{Hops: []string{"A", "B", "C"}},
					{Hops: []string{"A", "B", "Z"}},
				},
			},
			err: ErrMultihopExitTokensMismatch,
		},
		{
			name: "invalid amountIn",
			msg: MsgMultiHopSwap{
				Creator:  sample.AccAddress(),
				Receiver: sample.AccAddress(),
				Routes:   []*MultiHopRoute{{Hops: []string{"A", "B", "C"}}},
				AmountIn: sdk.NewInt(-1),
			},
			err: ErrZeroSwap,
		},
		{
			name: "valid",
			msg: MsgMultiHopSwap{
				Routes:   []*MultiHopRoute{{Hops: []string{"A", "B", "C"}}},
				Creator:  sample.AccAddress(),
				Receiver: sample.AccAddress(),
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
