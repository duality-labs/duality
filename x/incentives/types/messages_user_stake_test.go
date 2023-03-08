package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateUserStake_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateUserStake
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateUserStake{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateUserStake{
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

func TestMsgUpdateUserStake_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateUserStake
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateUserStake{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateUserStake{
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

func TestMsgDeleteUserStake_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteUserStake
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteUserStake{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteUserStake{
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
