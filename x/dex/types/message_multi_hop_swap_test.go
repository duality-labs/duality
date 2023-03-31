package types_test

import (
	"testing"

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
			name: "valid addresses",
			msg: MsgMultiHopSwap{
				Creator:  sample.AccAddress(),
				Receiver: sample.AccAddress(),
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
