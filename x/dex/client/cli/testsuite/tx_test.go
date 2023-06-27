package cli_test

import (
	"fmt"
	"io"
	"regexp"
	"testing"

	"cosmossdk.io/api/tendermint/abci"
	rpcclientmock "github.com/cometbft/cometbft/rpc/client/mock"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutilmod "github.com/cosmos/cosmos-sdk/types/module/testutil"
	bankclient "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/duality-labs/duality/x/dex"
	dexClient "github.com/duality-labs/duality/x/dex/client/cli"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TxTestSuite struct {
	suite.Suite

	kr        keyring.Keyring
	encCfg    testutilmod.TestEncodingConfig
	baseCtx   client.Context
	clientCtx client.Context
	addrs     []sdk.AccAddress

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
	s.encCfg = testutilmod.MakeTestEncodingConfig(dex.AppModuleBasic{})
	s.kr = keyring.NewInMemory(s.encCfg.Codec)
	s.baseCtx = client.Context{}.
		WithKeyring(s.kr).
		WithTxConfig(s.encCfg.TxConfig).
		WithCodec(s.encCfg.Codec).
		WithClient(clitestutil.MockCometRPC{Client: rpcclientmock.Client{}}).
		WithAccountRetriever(client.MockAccountRetriever{}).
		WithOutput(io.Discard).
		WithChainID("test-chain")

	ctxGen := func() client.Context {
		bz, _ := s.encCfg.Codec.Marshal(&sdk.TxResponse{})
		c := clitestutil.NewMockCometRPC(abci.ResponseQuery{
			Value: bz,
		})
		return s.baseCtx.WithClient(c)
	}
	s.clientCtx = ctxGen()

	s.addrs = make([]sdk.AccAddress, 0)
	for i := 0; i < 3; i++ {
		k, _, err := s.clientCtx.Keyring.NewMnemonic("NewValidator", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
		s.Require().NoError(err)

		pub, err := k.GetPubKey()
		s.Require().NoError(err)

		newAddr := sdk.AccAddress(pub.Address())
		s.addrs = append(s.addrs, newAddr)
	}
}

func (s *TxTestSuite) fundAccount(clientCtx client.Context, from, to sdk.AccAddress, coins sdk.Coins) {
	require := s.Require()

	commonFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
	}
	cmd := bankclient.Cmd()
	args := append(tc.args, commonFlags...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)

	out, err := clitestutil.MsgSendExec(
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
			// "withdrawal [receiver] [token-a] [token-b] [list of shares-to-remove] [list of tick-index] [list of fee indexes] ",
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
			cmd := dexClient.CmdWithdrawal()
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
			// "place-limit-order [receiver] [token-in] [token-out] [tick-index] [amount-in] ?[order-type] ?[expirationTime] ?(--max-amout-out)"
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
		{
			name:     "valid with maxAmountOut",
			args:     []string{s.addr1.String(), "TokenB", "TokenA", "[2]", "10", "FILL_OR_KILL", "--max-amount-out=10"},
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

	argsSwap := append([]string{s.addr1.String(), "TokenA", "TokenB", "0", "30", "IMMEDIATE_OR_CANCEL"}, commonFlags...)
	cmd = dexClient.CmdPlaceLimitOrder()
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
