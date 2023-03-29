package cli_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	network "github.com/duality-labs/duality/testutil/network"
	dexClient "github.com/duality-labs/duality/x/dex/client/cli"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TxTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	addr1      sdk.AccAddress
	addr2      sdk.AccAddress
	trancheKey string
}

func TestTxTestSuite(t *testing.T) {
	suite.Run(t, new(TxTestSuite))
}

func findTrancheKeyInTx(tx string) string {
	re := regexp.MustCompile(`TrancheKey.*?:\\"([a-z0-9]+)`)
	return re.FindStringSubmatch(tx)[1]
}

func (s *TxTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")
	nw := network.NewCLITest(s.T())
	s.network = nw

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	info1, _, err := s.network.Validators[0].ClientCtx.Keyring.NewMnemonic("acc1", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	info2, _, err := s.network.Validators[0].ClientCtx.Keyring.NewMnemonic("acc2", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	pk := info1.GetPubKey()
	s.addr1 = sdk.AccAddress(pk.Address())
	pk = info2.GetPubKey()
	s.addr2 = sdk.AccAddress(pk.Address())

	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	args := append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "[0]", "1", "false"}, commonFlags...)
	cmd := dexClient.CmdDeposit()
	_, err = cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(s.T(), err)

	args = append([]string{s.network.Validators[0].Address.String(), "TokenB", "TokenA", "[-20]", "10"}, commonFlags...)
	cmd = dexClient.CmdPlaceLimitOrder()
	txBuff, err := cli.ExecTestCLICmd(clientCtx, cmd, args)

	require.NoError(s.T(), err)
	s.trancheKey = findTrancheKeyInTx(txBuff.String())

	s.fundAccount(clientCtx, s.network.Validators[0].Address, s.addr1, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)), sdk.NewCoin("TokenA", sdk.NewInt(100000))))

	s.fundAccount(clientCtx, s.network.Validators[0].Address, s.addr2, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)), sdk.NewCoin("TokenA", sdk.NewInt(100000))))
}

func (s *TxTestSuite) fundAccount(clientCtx client.Context, from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) {
	require := s.Require()

	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
	}

	out, err := banktestutil.MsgSendExec(
		clientCtx,
		from,
		to,
		coins,
		commonFlags...,
	)
	require.NoError(err)
	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)
}

func (s *TxTestSuite) TestTxCmdDeposit() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			name:      "missing arguments",
			args:      []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "[0]", "false"},
			expErr:    true,
			expErrMsg: "Error: accepts 8 arg(s), received 7",
		},
		{
			name:      "too many arguments",
			args:      []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "[0]", "1", "false", s.addr1.String()},
			expErr:    true,
			expErrMsg: "Error: accepts 8 arg(s), received 9",
		},
		{
			name:     "valid",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "[0]", "1", "false"},
			errInRes: false,
		},
		{
			name:     "valid: multiple case",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "0,0", "10,10", "[25,25]", "1,1", "false,false"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdDeposit()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}
			}
		})
	}
}

func (s *TxTestSuite) TestTx2CmdWithdraw() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	// Deposit Funds
	args := append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "[0]", "0", "false"}, commonFlags...)
	cmd := dexClient.CmdDeposit()
	_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(s.T(), err)

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			// "withdrawl [receiver] [token-a] [token-b] [list of shares-to-remove] [list of tick-index] [list of fee indexes] ",
			name:      "missing arguments",
			args:      []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "[10]", "0"},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 5",
		},
		{
			name:      "too many arguments",
			args:      []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "[0]", "1", s.addr1.String()},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 7",
		},
		{
			name:     "valid",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "[0]", "1"},
			errInRes: false,
		},
		{
			name:     "valid: multiple case",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "2,2", "[0,0]", "0,1"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdWithdrawl()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}
			}
		})
	}
}

func (s *TxTestSuite) TestTx3CmdSwap() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			// "swap [receiver] [amount-in] [token-in] [token-out]",
			name:      "missing arguments",
			args:      []string{s.addr1.String(), "5", "TokenA"},
			expErr:    true,
			expErrMsg: "Error: accepts 4 arg(s), received 3",
		},
		{
			name:      "too many arguments",
			args:      []string{s.addr1.String(), "5", "TokenA", "TokenB", "BADARG"},
			expErr:    true,
			expErrMsg: "Error: accepts 4 arg(s), received 5",
		},
		{
			name:     "valid",
			args:     []string{s.addr1.String(), "2", "TokenA", "TokenB"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdSwap()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}
			}
		})
	}
}

func (s *TxTestSuite) TestTx4Cmd4PlaceLimitOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			// "place-limit-order [receiver] [token-in] [token-out] [tick-index] [amount-in] ?[order-type] ?[expirationTime]"
			name:      "missing arguments",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "[0]"},
			expErr:    true,
			expErrMsg: "Error: accepts between 5 and 7 arg(s), received 4",
		},
		{
			name:      "too many arguments",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "[0]", "10", "1", "1", "BAD"},
			expErr:    true,
			expErrMsg: "Error: accepts between 5 and 7 arg(s), received 8",
		},
		{
			name:      "invalid orderType",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "[0]", "10", "JUST_SEND_IT"},
			expErr:    true,
			expErrMsg: types.ErrInvalidOrderType.Error(),
		},
		{
			name:      "invalid goodTil",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "[0]", "10", "GOOD_TIL_TIME", "january"},
			expErr:    true,
			expErrMsg: types.ErrInvalidTimeString.Error(),
		},
		{
			name:     "valid",
			args:     []string{s.addr1.String(), "TokenB", "TokenA", "[0]", "10"},
			errInRes: false,
		},
		{
			name:     "valid goodTil",
			args:     []string{s.addr1.String(), "TokenB", "TokenA", "[0]", "10", "GOOD_TIL_TIME", "06/15/2025 02:00:00"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdPlaceLimitOrder()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}
			}
		})
	}
}

func (s *TxTestSuite) TestTx5CmdCancelLimitOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			//  "cancel-limit-order [tranche-key]"
			name:      "missing arguments",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many arguments",
			args:      []string{"trancheKey123", "extraarg"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:     "valid",
			args:     []string{s.trancheKey},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdCancelLimitOrder()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}
			}
		})
	}
}

func (s *TxTestSuite) TestTx6CmdWithdrawFilledLimitOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	// Place Limit Order
	args := append([]string{s.network.Validators[0].Address.String(), "TokenB", "TokenA", "[0]", "10"}, commonFlags...)
	cmd := dexClient.CmdPlaceLimitOrder()
	txBuff, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(s.T(), err)
	trancheKey := findTrancheKeyInTx(txBuff.String())

	argsSwap := append([]string{s.network.Validators[0].Address.String(), "30", "TokenA", "TokenB"}, commonFlags...)
	cmd = dexClient.CmdSwap()
	_, err = cli.ExecTestCLICmd(clientCtx, cmd, argsSwap)
	require.NoError(s.T(), err)

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			//  "withdraw-filled-limit-order [tranche-key]"
			name:      "missing arguments",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many arguments",
			args:      []string{"trancheKey123", "EXTRA-ARG"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:     "valid",
			args:     []string{trancheKey},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdWithdrawFilledLimitOrder()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}
			}
		})
	}
}
