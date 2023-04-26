package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	. "github.com/duality-labs/duality/x/mev/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSend
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSend{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSend{
				Creator: sample.AccAddress(),
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
