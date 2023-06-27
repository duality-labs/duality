package interchaintest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/duality-labs/duality/x/dex/types"
	swaptypes "github.com/duality-labs/duality/x/ibcswap/types"
	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	"github.com/strangelove-ventures/interchaintest/v4/relayer"
	"github.com/strangelove-ventures/interchaintest/v4/relayer/rly"
	"github.com/strangelove-ventures/interchaintest/v4/testreporter"
	"github.com/strangelove-ventures/interchaintest/v4/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestIBCSwapMiddleware_Success asserts that the IBC swap middleware works as intended with Duality running as a
// consumer chain connected to the Cosmos Hub.
func TestIBCSwapMiddleware_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// Create chain factory with Duality and Cosmos Hub
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg},
	},
	)

	// Get both chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

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
		AddRelayer(r, "relayer").
		AddProviderConsumerLink(interchaintest.ProviderConsumerLink{
			Provider:         chainA,
			Consumer:         chainB,
			Relayer:          r,
			Path:             pathICS,
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
	require.NoError(t, r.StartRelayer(ctx, eRep, pathICS))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, chainA, chainB)

	// wait a couple blocks for wallets to be initialized
	err = testutil.WaitForBlocks(ctx, 5, chainA, chainB)
	require.NoError(t, err)

	chainAKey, chainBKey := users[0], users[1]

	// Get our bech32 encoded user address
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainBAddr := chainBKey.Bech32Address(chainB.Config().Bech32Prefix)

	// Get the original acc balances on both chains for their native tokens
	chainAOrigBalNative, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainAOrigBalNative)

	chainBOrigBalNative, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBOrigBalNative)

	// Get channel between Gaia and Duality
	abChannel, err := ibc.GetTransferChannel(ctx, r, eRep, chainA.Config().ChainID, chainB.Config().ChainID)
	require.NoError(t, err)

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBAddr,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	// Send an IBC transfer from Gaia to Duality, so we can initialize a pool with the IBC denom token + native Duality token
	transferTx, err := chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{
		Timeout: nil,
		Memo:    "",
	})
	require.NoError(t, err)

	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	// Poll for the ack to know that the transfer is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Get the IBC denom for ATOM on Duality
	chainATokenDenom := transfertypes.GetPrefixedDenom(abChannel.Counterparty.PortID, abChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Assert that the funds are gone from the acc on Gaia and present in the acc on Duality
	chainABalTransfer, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainAOrigBalNative-ibcTransferAmount, chainABalTransfer)

	chainBBalIBCTransfer, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, ibcTransferAmount, chainBBalIBCTransfer)

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
	require.Equal(t, chainBBalIBCTransfer-depositAmount.Int64(), chainBBalIBC)

	chainBBalNative, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBOrigBalNative-depositAmount.Int64(), chainBBalNative)

	swapAmount := sdktypes.NewInt(100000)
	expectedOut := sdktypes.NewInt(99_990)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
				Creator:   chainBAddr,
				Receiver:  chainBAddr,
				TokenIn:   chainADenomTrace.IBCDenom(),
				TokenOut:  chainB.Config().Denom,
				AmountIn:  swapAmount,
				TickIndex: 10,
				OrderType: types.LimitOrderType_FILL_OR_KILL,
				// TODO: enable soon
				// MaxAmountOut: minOut,
			},
			Next: nil,
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from Gaia to Duality with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on Gaia
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABalTransfer-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on Duality
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative+expectedOut.Int64(), chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)
}

// TestIBCSwapMiddleware_FailRefund asserts that the IBC swap middleware works as intended with Duality running as a
// consumer chain connected to the Cosmos Hub. The swap should fail and a refund to the src chain should take place.
func TestIBCSwapMiddleware_FailRefund(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// Create chain factory with Gaia and Duality
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg},
	},
	)

	// Get both chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

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
		AddRelayer(r, "relayer").
		AddProviderConsumerLink(interchaintest.ProviderConsumerLink{
			Provider:         chainA,
			Consumer:         chainB,
			Relayer:          r,
			Path:             pathICS,
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
	require.NoError(t, r.StartRelayer(ctx, eRep, pathICS))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, chainA, chainB)

	// wait a couple blocks for wallets to be initialized
	err = testutil.WaitForBlocks(ctx, 10, chainA, chainB)
	require.NoError(t, err)

	chainAKey, chainBKey := users[0], users[1]

	// Get our bech32 encoded user address
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainBAddr := chainBKey.Bech32Address(chainB.Config().Bech32Prefix)

	// Get channel between Gaia and Duality
	abChannel, err := ibc.GetTransferChannel(ctx, r, eRep, chainA.Config().ChainID, chainB.Config().ChainID)
	require.NoError(t, err)

	// Get the IBC denom for ATOM on Duality
	chainATokenDenom := transfertypes.GetPrefixedDenom(abChannel.Counterparty.PortID, abChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the acc balances before the transfer and swap takes place
	chainABal, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainABal)

	chainBNativeBal, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBNativeBal)

	chainBIBCBal, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainBIBCBal)

	// Compose the swap metadata, this swap will fail because there is no pool initialized for this pair
	swapAmount := sdktypes.NewInt(100000)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
				Creator:   chainBAddr,
				Receiver:  chainBAddr,
				TokenIn:   chainADenomTrace.IBCDenom(),
				TokenOut:  chainB.Config().Denom,
				AmountIn:  swapAmount,
				TickIndex: 0,
				OrderType: types.LimitOrderType_FILL_OR_KILL,
				// TODO: enable soon
				// MaxAmountOut: minOut,
			},
			NonRefundable: false,
			Next:          nil,
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBAddr,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from Gaia to Duality with packet memo containing the swap metadata
	transferTx, err := chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap has failed
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+15, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are not present in the account on Duality
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBNativeBal, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBIBCBal, chainBBalIBCSwap)

	// Check that the refund takes place and the funds are moved back to the account on Gaia
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABal, chainABalAfterSwap)
}

// TestIBCSwapMiddleware_FailNoRefund asserts that the IBC swap middleware works as intended with Duality running as a
// consumer chain connected to the Cosmos Hub. The swap should fail and funds should remain on Duality.
func TestIBCSwapMiddleware_FailNoRefund(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// Create chain factory with Duality and Cosmos Hub
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg},
	},
	)

	// Get both chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

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
		AddRelayer(r, "relayer").
		AddProviderConsumerLink(interchaintest.ProviderConsumerLink{
			Provider:         chainA,
			Consumer:         chainB,
			Relayer:          r,
			Path:             pathICS,
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
	require.NoError(t, r.StartRelayer(ctx, eRep, pathICS))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, chainA, chainB)

	// wait a couple blocks for wallets to be initialized
	err = testutil.WaitForBlocks(ctx, 2, chainA, chainB)
	require.NoError(t, err)

	chainAKey, chainBKey := users[0], users[1]

	// Get our bech32 encoded user address
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainBAddr := chainBKey.Bech32Address(chainB.Config().Bech32Prefix)

	// Get channel between Gaia and Duality
	abChannel, err := ibc.GetTransferChannel(ctx, r, eRep, chainA.Config().ChainID, chainB.Config().ChainID)
	require.NoError(t, err)

	// Get the IBC denom for ATOM on Duality
	chainATokenDenom := transfertypes.GetPrefixedDenom(abChannel.Counterparty.PortID, abChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the acc balances before the transfer and swap takes place
	chainABal, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainABal)

	chainBNativeBal, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBNativeBal)

	chainBIBCBal, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainBIBCBal)

	// Compose the swap metadata, this swap will fail because there is no pool initialized for this pair
	swapAmount := sdktypes.NewInt(100000)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
				Creator:   chainBAddr,
				Receiver:  chainBAddr,
				TokenIn:   chainADenomTrace.IBCDenom(),
				TokenOut:  chainB.Config().Denom,
				AmountIn:  swapAmount,
				TickIndex: 1,
				OrderType: types.LimitOrderType_FILL_OR_KILL,
				// TODO: enable soon
				// MaxAmountOut: minOut,
			},
			NonRefundable: true,
			Next:          nil,
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBAddr,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from Gaia to Duality with packet memo containing the swap metadata
	transferTx, err := chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap has failed
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+15, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are present in the account on Duality
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBNativeBal, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBIBCBal+ibcTransferAmount, chainBBalIBCSwap)

	// Check that no refund takes place and the funds are not in the account on Gaia
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABal-ibcTransferAmount, chainABalAfterSwap)
}

// TestIBCSwapMiddleware_FailWithRefundAddr asserts that the IBC swap middleware works as intended with Duality running as a
// consumer chain connected to the Cosmos Hub. The swap should fail and funds should remain on Duality but be moved
// to the refund address.
func TestIBCSwapMiddleware_FailWithRefundAddr(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// Create chain factory with Duality and Cosmos Hub
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg},
	},
	)

	// Get both chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

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
		AddRelayer(r, "relayer").
		AddProviderConsumerLink(interchaintest.ProviderConsumerLink{
			Provider:         chainA,
			Consumer:         chainB,
			Relayer:          r,
			Path:             pathICS,
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
	require.NoError(t, r.StartRelayer(ctx, eRep, pathICS))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, chainA, chainB, chainB)

	// wait a couple blocks for wallets to be initialized
	err = testutil.WaitForBlocks(ctx, 2, chainA, chainB)
	require.NoError(t, err)

	chainAKey, chainBKey, chainBRefundKey := users[0], users[1], users[2]

	// Get our bech32 encoded user address
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainBAddr := chainBKey.Bech32Address(chainB.Config().Bech32Prefix)
	chainBRefundAddr := chainBRefundKey.Bech32Address(chainB.Config().Bech32Prefix)

	// Get channel between Gaia and Duality
	abChannel, err := ibc.GetTransferChannel(ctx, r, eRep, chainA.Config().ChainID, chainB.Config().ChainID)
	require.NoError(t, err)

	// Get the IBC denom for ATOM on Duality
	chainATokenDenom := transfertypes.GetPrefixedDenom(abChannel.Counterparty.PortID, abChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the acc balances before the transfer and swap takes place
	chainABal, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainABal)

	chainBNativeBal, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBNativeBal)

	chainBIBCBal, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainBIBCBal)

	chainBRefundBal, err := chainB.GetBalance(ctx, chainBRefundAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainBRefundBal)

	// Compose the swap metadata, this swap will fail because there is no pool initialized for this pair
	swapAmount := sdktypes.NewInt(100000)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgPlaceLimitOrder: &types.MsgPlaceLimitOrder{
				Creator:   chainBAddr,
				Receiver:  chainBAddr,
				TokenIn:   chainADenomTrace.IBCDenom(),
				TokenOut:  chainB.Config().Denom,
				AmountIn:  swapAmount,
				TickIndex: 1,
				OrderType: types.LimitOrderType_FILL_OR_KILL,
				// TODO: enable soon
				// MaxAmountOut: minOut,
			},
			NonRefundable: true,
			RefundAddress: chainBRefundAddr,
			Next:          nil,
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBAddr,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from Gaia to Duality with packet memo containing the swap metadata
	transferTx, err := chainA.SendIBCTransfer(ctx, abChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap has failed
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+15, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds have been moved to the refund address
	refundNewBal, err := chainB.GetBalance(ctx, chainBRefundAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBRefundBal+ibcTransferAmount, refundNewBal)

	// Check that the funds are present in the account on Duality
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBAddr, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBNativeBal, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBAddr, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBIBCBal, chainBBalIBCSwap)

	// Check that no refund takes place and the funds are not in the account on Gaia
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABal-ibcTransferAmount, chainABalAfterSwap)
}
