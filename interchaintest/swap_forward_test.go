package interchaintest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	"github.com/duality-labs/duality/x/dex/types"
	swaptypes "github.com/duality-labs/duality/x/ibcswap/types"
	"github.com/iancoleman/orderedmap"
	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	"github.com/strangelove-ventures/interchaintest/v4/relayer"
	"github.com/strangelove-ventures/interchaintest/v4/relayer/rly"
	"github.com/strangelove-ventures/interchaintest/v4/testreporter"
	"github.com/strangelove-ventures/interchaintest/v4/testutil"
	forwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v7/router/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestSwapAndForward_Success asserts that the swap and forward middleware stack works as intended with Duality running as a
// consumer chain connected to two other chains via IBC.
func TestSwapAndForward_Success(t *testing.T) {
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

	bz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

	nextJSON := new(swaptypes.JSONObject)
	err = json.Unmarshal(bz, nextJSON)
	require.NoError(t, err)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
				Creator:   chainBAddr,
				Receiver:  chainBAddr,
				TokenIn:   chainADenomTrace.IBCDenom(),
				TokenOut:  chainB.Config().Denom,
				AmountIn:  swapAmount,
				TickIndex: 2,
				OrderType: types.LimitOrderType_FILL_OR_KILL,
			},
			Next: nextJSON,
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
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

// TestSwapAndForward_MultiHopSuccess asserts that the swap and forward middleware stack works as intended in the case
// of a multi-hop forward after the swap.
func TestSwapAndForward_MultiHopSuccess(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// Create chain factory
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg},
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-c", GasPrices: "0.0uatom"}},
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-d", GasPrices: "0.0uatom"}},
	},
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB, chainC, chainD := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain), chains[2].(*cosmos.CosmosChain), chains[3].(*cosmos.CosmosChain)

	// Create relayer factory with the go-relayer
	ctx := context.Background()
	client, network := interchaintest.DockerSetup(t)

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
		AddChain(chainD).
		AddRelayer(r, "relayer").
		AddProviderConsumerLink(interchaintest.ProviderConsumerLink{
			Provider: chainA,
			Consumer: chainB,
			Relayer:  r,
			Path:     pathICS,
		}).
		AddLink(interchaintest.InterchainLink{
			Chain1:           chainB,
			Chain2:           chainC,
			Relayer:          r,
			Path:             pathChainBChainC,
			CreateClientOpts: ibc.CreateClientOptions{TrustingPeriod: "15m"},
		}).
		AddLink(interchaintest.InterchainLink{
			Chain1:           chainC,
			Chain2:           chainD,
			Relayer:          r,
			Path:             pathChainCChainD,
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
	require.NoError(t, r.StartRelayer(ctx, eRep, pathICS, pathChainBChainC, pathChainCChainD))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, chainA, chainB, chainC, chainD)

	// wait a couple blocks for wallets to be initialized and channels to be opened
	err = testutil.WaitForBlocks(ctx, 20, chainA, chainB, chainC, chainD)
	require.NoError(t, err)

	chainAKey, chainBKey, chainCKey, chainDKey := users[0], users[1], users[2], users[3]

	// Get our bech32 encoded user addresses
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainBAddr := chainBKey.Bech32Address(chainB.Config().Bech32Prefix)
	chainCAddr := chainCKey.Bech32Address(chainC.Config().Bech32Prefix)
	chainDAddr := chainDKey.Bech32Address(chainD.Config().Bech32Prefix)

	// Get the original acc balances for each chains respective native token
	chainAOrigBalNative, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainAOrigBalNative)

	chainBOrigBalNative, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBOrigBalNative)

	// Get channels between chainA-chainB, chainB-chainC, and chainC-chainD
	channels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)

	var abChannel ibc.ChannelOutput
	for _, channel := range channels {
		if channel.PortID == "transfer" {
			abChannel = channel
		}
	}

	chainCChannels, err := r.GetChannels(ctx, eRep, chainC.Config().ChainID)
	require.NoError(t, err)
	cbChannel := chainCChannels[1]

	chainDChannels, err := r.GetChannels(ctx, eRep, chainD.Config().ChainID)
	require.NoError(t, err)
	dcChannel := chainDChannels[0]

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
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+15, transferTx.Packet)
	require.NoError(t, err)

	// Get the IBC denom on chainB for the native token from chainA
	chainATokenDenom := transfertypes.GetPrefixedDenom(abChannel.Counterparty.PortID, abChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the IBC denom on chainC for the native token from chainB
	chainCTokenDenom := transfertypes.GetPrefixedDenom(cbChannel.PortID, cbChannel.ChannelID, chainB.Config().Denom)
	chainCDenomTrace := transfertypes.ParseDenomTrace(chainCTokenDenom)

	// Get the IBC denom on chainD for the native token from chainB that has travelled from chainB->chainC->chainD
	chainDTokenDenom := transfertypes.GetPrefixedDenom(dcChannel.PortID, dcChannel.ChannelID, chainCTokenDenom)
	chainDDenomTrace := transfertypes.ParseDenomTrace(chainDTokenDenom)

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

	expectedOut := sdktypes.NewInt(99_990)

	retries := uint8(0)
	nextForward := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainDAddr,
			Port:     dcChannel.Counterparty.PortID,
			Channel:  dcChannel.Counterparty.ChannelID,
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     nil,
		},
	}
	nextForwardBz, err := json.Marshal(nextForward)
	require.NoError(t, err)
	nextForwardJSON := forwardtypes.NewJSONObject(false, nextForwardBz, orderedmap.OrderedMap{})

	forwardMetadata := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainCAddr,
			Port:     cbChannel.Counterparty.PortID,
			Channel:  cbChannel.Counterparty.ChannelID,
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     nextForwardJSON,
		},
	}
	bz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

	nextJSON := new(swaptypes.JSONObject)
	err = json.Unmarshal(bz, nextJSON)
	require.NoError(t, err)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
				Creator:   chainBAddr,
				Receiver:  chainBAddr,
				TokenIn:   chainADenomTrace.IBCDenom(),
				TokenOut:  chainB.Config().Denom,
				AmountIn:  swapAmount,
				TickIndex: 2,
				OrderType: types.LimitOrderType_FILL_OR_KILL,
			},
			Next: nextJSON,
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap and forward is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on chainA
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABalTransfer-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on chainD
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)

	chainCBal, err := chainC.GetBalance(ctx, chainCAddr, chainCDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainCBal)

	chainDBal, err := chainD.GetBalance(ctx, chainDAddr, chainDDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, expectedOut.Int64(), chainDBal)
}

// TestSwapAndForward_UnwindIBCDenomSuccess asserts that the swap and forward middleware stack works as intended in the
// case that a native token from ChainB is sent to ChainA and then ChainA initiates a swap and forward with the token.
// This asserts that denom unwinding works as intended when going ChainB->ChainA->ChainB.
func TestSwapAndForward_UnwindIBCDenomSuccess(t *testing.T) {
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

	chainCChannels, err := r.GetChannels(ctx, eRep, chainC.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainCChannels))
	cbChannel := chainCChannels[0]

	// Compose details for an IBC transfer
	tmpTransferAmount := int64(1000000)
	transfer := ibc.WalletAmount{
		Address: chainBAddr,
		Denom:   chainA.Config().Denom,
		Amount:  tmpTransferAmount,
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

	// Get the IBC denom on chainC for the native token from chainA that travelled from chainA->chainB->chainC
	chainCTokenDenom := transfertypes.GetPrefixedDenom(cbChannel.PortID, cbChannel.ChannelID, chainATokenDenom)
	chainCDenomTrace := transfertypes.ParseDenomTrace(chainCTokenDenom)

	// Assert that the funds are gone from the acc on chainA and present in the acc on chainB
	chainABalTransfer, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainAOrigBalNative-tmpTransferAmount, chainABalTransfer)

	chainBBalTransfer, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, tmpTransferAmount, chainBBalTransfer)

	// Compose the deposit cmd for initializing a pool on Duality (chainB)
	depositAmount := sdktypes.NewInt(1000000)

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

	// Compose a transfer from ChainB->ChainA for ChainB's native token
	chainBTransfer := ibc.WalletAmount{
		Address: chainAAddr,
		Denom:   chainB.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	transferTx, err = chainB.SendIBCTransfer(ctx, abChannel.Counterparty.ChannelID, chainBAddr, chainBTransfer, ibc.TransferOptions{
		Timeout: nil,
		Memo:    "",
	})
	require.NoError(t, err)

	chainBHeight, err := chainB.Height(ctx)
	require.NoError(t, err)

	// Poll for the ack to know the transfer is complete
	_, err = testutil.PollForAck(ctx, chainB, chainBHeight, chainBHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Get the IBC denom on chainA for the native token from chainB
	chainBTokenDenom := transfertypes.GetPrefixedDenom(abChannel.PortID, abChannel.ChannelID, chainB.Config().Denom)
	chainBDenomTrace := transfertypes.ParseDenomTrace(chainBTokenDenom)

	// Assert that the funds are present in the acc on chainA
	chainABal, err := chainA.GetBalance(ctx, chainAAddr, chainBDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, ibcTransferAmount, chainABal)

	// Assert that the funds are gone from the acc on chainB
	chainBBal, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative-ibcTransferAmount, chainBBal)

	// Compose the IBC transfer memo metadata to be used in the swap and forward
	swapAmount := sdktypes.NewInt(100_000)
	expectedOut := sdktypes.NewInt(99_990)

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

	bz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

	nextJSON := new(swaptypes.JSONObject)
	err = json.Unmarshal(bz, nextJSON)
	require.NoError(t, err)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
				Creator:   chainBAddr,
				Receiver:  chainBAddr,
				TokenIn:   chainB.Config().Denom,
				TokenOut:  chainADenomTrace.IBCDenom(),
				AmountIn:  swapAmount,
				TickIndex: 2,
				OrderType: types.LimitOrderType_FILL_OR_KILL,
			},
			Next: nextJSON,
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	swapTransfer := ibc.WalletAmount{
		Address: chainBAddr,
		Denom:   chainBDenomTrace.IBCDenom(),
		Amount:  ibcTransferAmount,
	}

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, swapTransfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap and forward is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on chainA
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainBDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainABal-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on chainC
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBal, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)

	chainCBal, err := chainC.GetBalance(ctx, chainCAddr, chainCDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, expectedOut.Int64(), chainCBal)
}

// TestSwapAndForward_ForwardFails asserts that the swap and forward middleware stack works as intended in the case
// that an incoming IBC swap succeeds but the forward fails.
func TestSwapAndForward_ForwardFails(t *testing.T) {
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

	chainCChannels, err := r.GetChannels(ctx, eRep, chainC.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainCChannels))
	cbChannel := chainCChannels[0]

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
	expectedOut := sdktypes.NewInt(99_990)

	retries := uint8(0)
	forwardMetadata := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainCAddr,
			Port:     cbChannel.Counterparty.PortID,
			Channel:  "invalid-channel", // add an invalid channel identifier so the forward fails
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     nil,
		},
	}

	bz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

	nextJSON := new(swaptypes.JSONObject)
	err = json.Unmarshal(bz, nextJSON)
	require.NoError(t, err)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
				Creator:   chainBAddr,
				Receiver:  chainBAddr,
				TokenIn:   chainADenomTrace.IBCDenom(),
				TokenOut:  chainB.Config().Denom,
				AmountIn:  swapAmount,
				TickIndex: 2,
				OrderType: types.LimitOrderType_FILL_OR_KILL,
			},
			Next: nextJSON,
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap and forward is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on chainA
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABalTransfer-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on chainB
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative+expectedOut.Int64(), chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)

	chainCBal, err := chainC.GetBalance(ctx, chainCAddr, chainCDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainCBal)
}
