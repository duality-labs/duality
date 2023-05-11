package interchaintest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	gmp "github.com/duality-labs/duality/x/gmp"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	"github.com/strangelove-ventures/interchaintest/v4/relayer"
	"github.com/strangelove-ventures/interchaintest/v4/relayer/rly"
	"github.com/strangelove-ventures/interchaintest/v4/testreporter"
	"github.com/strangelove-ventures/interchaintest/v4/testutil"
	forwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestSwapAndForward_Success asserts that the swap and forward middleware stack works as intended with Duality running as a
// consumer chain connected to two other chains via IBC.
func TestGMPSwapAndForward_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// Create chain factory
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg},
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-c", GasPrices: "0.0uatom"}},
	},
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB, chainC := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain), chains[2].(*cosmos.CosmosChain)

	ctx := context.Background()
	client, network := interchaintest.DockerSetup(t)

	// Create relayer factory with the go-relayer
	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-paths_update", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	// Initialize the Interchain object which describes the chains, relayers, and paths between chains
	ic := interchaintest.NewInterchain().
		AddChain(chainA).
		AddChain(chainB).
		AddChain(chainC).
		AddRelayer(r, "relayer").
		AddProviderConsumerLink(interchaintest.ProviderConsumerLink{
			Provider:         chainA,
			Consumer:         chainB,
			Relayer:          r,
			Path:             pathICS,
			CreateClientOpts: ibc.CreateClientOptions{TrustingPeriod: "15m"},
		}).
		AddLink(interchaintest.InterchainLink{
			Chain1:           chainB,
			Chain2:           chainC,
			Relayer:          r,
			Path:             pathChainBChainC,
			CreateClientOpts: ibc.CreateClientOptions{TrustingPeriod: "15m"},
		})

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: false,
	}))

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathICS, pathChainBChainC))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, chainA, chainB, chainC)

	// wait a couple blocks for wallets to be initialized and channels to be opened
	err = testutil.WaitForBlocks(ctx, 10, chainA, chainB, chainC)
	require.NoError(t, err)

	chainAKey, chainBKey, chainCKey := users[0], users[1], users[2]

	// Get our bech32 encoded user addresses
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainBAddr := chainBKey.Bech32Address(chainB.Config().Bech32Prefix)
	chainCAddr := chainCKey.Bech32Address(chainC.Config().Bech32Prefix)

	// Get the original acc balances for each chains respective native token
	chainAOrigBalNative, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainAOrigBalNative)

	chainBOrigBalNative, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBOrigBalNative)

	// Get channels between chainA-chainB and chainB-chainC
	channels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)

	var abChannel ibc.ChannelOutput
	for _, channel := range channels {
		if channel.PortID == "transfer" {
			abChannel = channel
		}
	}

	channels, err = r.GetChannels(ctx, eRep, chainC.Config().ChainID)
	require.NoError(t, err)

	cbChannel := channels[0]

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBAddr,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	// Send an IBC transfer from chainA to chainB, so we can initialize a pool with the IBC denom token + native Duality token
	transferTx, err := chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{
		Timeout: nil,
		Memo:    "",
	})
	require.NoError(t, err)

	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	// Poll for the ack to know that the transfer is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Get the IBC denom on chainB for the native token from chainA
	chainATokenDenom := transfertypes.GetPrefixedDenom(abChannel.Counterparty.PortID, abChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the IBC denom on chainC for the native token from chainB
	chainCTokenDenom := transfertypes.GetPrefixedDenom(cbChannel.PortID, cbChannel.ChannelID, chainB.Config().Denom)
	chainCDenomTrace := transfertypes.ParseDenomTrace(chainCTokenDenom)

	// Assert that the funds are gone from the acc on chainA and present in the acc on chainB
	chainABalTransfer, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainAOrigBalNative-ibcTransferAmount, chainABalTransfer)

	chainBBalTransfer, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, ibcTransferAmount, chainBBalTransfer)

	// Compose the deposit cmd for initializing a pool on Duality (chainB)
	depositAmount := sdktypes.NewInt(100000)

	depositCmd := []string{
		chainB.Config().Bin, "tx", "dex", "deposit",
		chainBAddr,
		chainB.Config().Denom,
		chainADenomTrace.IBCDenom(),
		depositAmount.String(),
		depositAmount.String(),
		"0",
		"1",
		"false",
		"--chain-id", chainB.Config().ChainID,
		"--node", chainB.GetRPCAddress(),
		"--from", chainBKey.KeyName,
		"--keyring-backend", "test",
		"--gas", "auto",
		"--yes",
		"--home", chainB.HomeDir(),
	}

	// Execute the deposit cmd to initialize the pool on Duality
	_, _, err = chainB.Exec(ctx, depositCmd, nil)
	require.NoError(t, err)

	// Wait for the tx to be included in a block
	err = testutil.WaitForBlocks(ctx, 5, chainB)
	require.NoError(t, err)

	// Assert that the deposit was successful and the funds are moved out of the Duality user acc
	chainBBalIBC, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalTransfer-depositAmount.Int64(), chainBBalIBC)

	chainBBalNative, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBOrigBalNative-depositAmount.Int64(), chainBBalNative)

	// Compose the IBC transfer memo metadata to be used in the swap and forward
	swapAmount := sdktypes.NewInt(100000)
	expectedAmountOut := sdktypes.NewInt(99990)

	retries := uint8(0)
	forwardMetadata := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainCAddr,
			Port:     cbChannel.Counterparty.PortID,
			Channel:  cbChannel.Counterparty.ChannelID,
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     nil,
		},
	}

	forwardMetadataBz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

	// forwardNextJSON := new(swaptypes.JSONObject)
	// err = json.Unmarshal(forwardMetadataBz, forwardNextJSON)
	// require.NoError(t, err)

	// swapMetadata := swaptypes.PacketMetadata{
	// 	Swap: &swaptypes.SwapMetadata{
	// 		MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
	// 			Creator:   chainBAddr,
	// 			Receiver:  chainBAddr,
	// 			TokenIn:   chainADenomTrace.IBCDenom(),
	// 			TokenOut:  chainB.Config().Denom,
	// 			AmountIn:  swapAmount,
	// 			TickIndex: 2,
	// 			OrderType: types.LimitOrderType_FILL_OR_KILL,
	// 		},
	// 		Next: nextJSON,
	// 	},
	// }

	// swapMetadataBz, err := json.Marshal(swapMetadata)
	// require.NoError(t, err)

	abi, err := abi.JSON(strings.NewReader(gmp.SwapForwardABI))
	assert.NoError(t, err)

	// timeIn10Minutes := time.Now().Add(10 * time.Minute).Unix()
	args := gmp.SwapForwardArgs{
		Creator:   chainBAddr,
		Receiver:  chainBAddr,
		TokenIn:   chainADenomTrace.IBCDenom(),
		TokenOut:  chainB.Config().Denom,
		AmountIn:  swapAmount.BigInt(),
		TickIndex: 2,
		OrderType: uint8(dextypes.LimitOrderType_FILL_OR_KILL),
		// TODO: Test whether these are mandatory (they shouldn't be)
		// NonRefundable:  false,
		// RefundAddress:  "alice",
		NextArgs: forwardMetadataBz,
	}

	payloadBz, err := abi.Pack(
		gmp.SwapForwardFunctionName,
		args.Creator,
		args.Receiver,
		args.TokenIn,
		args.TokenOut,
		args.TickIndex,
		args.AmountIn,
		uint8(args.OrderType),
		args.ExpirationTime,
		args.NonRefundable,
		args.RefundAddress,
		args.NextArgs,
	)
	assert.NoError(t, err)

	gmpMetadata := gmp.Message{
		SourceChain:   "axelar", // TODO: double check this is a good test value
		SourceAddress: "alice",
		Payload:       payloadBz,
		Type:          gmp.TypeGeneralMessageWithToken, // should be 2
	}
	gmpMetadataBz, err := json.Marshal(gmpMetadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(gmpMetadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap and forward is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on chainA
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABalTransfer-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on chainC
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)

	chainCBal, err := chainC.GetBalance(ctx, chainCAddr, chainCDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, expectedAmountOut.Int64(), chainCBal)
}
