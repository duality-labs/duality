package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateIncentivePlan_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateIncentivePlan
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateIncentivePlan{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateIncentivePlan{
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

func TestMsgUpdateIncentivePlan_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateIncentivePlan
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateIncentivePlan{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateIncentivePlan{
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

func TestMsgDeleteIncentivePlan_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteIncentivePlan
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteIncentivePlan{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteIncentivePlan{
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
