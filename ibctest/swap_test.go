package ibctest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/duality-labs/duality/x/dex/types"
	swaptypes "github.com/duality-labs/duality/x/ibcswap/types"
	"github.com/strangelove-ventures/ibctest/v3"
	"github.com/strangelove-ventures/ibctest/v3/chain/cosmos"
	"github.com/strangelove-ventures/ibctest/v3/ibc"
	"github.com/strangelove-ventures/ibctest/v3/relayer"
	"github.com/strangelove-ventures/ibctest/v3/relayer/rly"
	"github.com/strangelove-ventures/ibctest/v3/testreporter"
	"github.com/strangelove-ventures/ibctest/v3/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestIBCSwapMiddleware_Success asserts that the IBC swap middleware works as intended with Duality running as a
// standalone consumer chain connected to the Cosmos Hub.
func TestIBCSwapMiddleware_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Number of full nodes and validators in the network
	nv := 1
	nf := 0

	// Create chain factory with Gaia and Duality
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "cosmoshub-4", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg, NumValidators: &nv, NumFullNodes: &nf}},
	)

	// Get both chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	ctx := context.Background()
	client, network := ibctest.DockerSetup(t)

	// Create relayer factory with the go-relayer
	// TODO the custom docker image can be removed here once ICS query fix is merged into main in the relayer
	r := ibctest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-ics_consumer_unbonding_period_query", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	// Initialize the Interchain object which describes the chains, relayers, and paths between chains
	// We only use this for spinning up Gaia and initializing the relayer config because there is no ICS support for Duality.
	ic := ibctest.NewInterchain().
		AddChain(chainA).
		AddRelayer(r, "relayer")

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	require.NoError(t, ic.Build(ctx, eRep, ibctest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: true,
	}))

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Initialize the Duality nodes
	err = chainB.Initialize(ctx, t.Name(), client, network)
	require.NoError(t, err, "failed to initialize duality chain")

	chainBValidator := chainB.Validators[0]

	// Initialize the Duality node files, create genesis wallets, and start the chain
	kr := keyring.NewInMemory()

	chainBWallets, err := initDuality(ctx, chainBValidator, kr, []string{aliceKeyName, rlyChainBKeyName})
	require.NoError(t, err)

	chainBKey, rlyChainBKey := chainBWallets[0], chainBWallets[1]

	t.Cleanup(func() {
		err = chainBValidator.StopContainer(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to stop duality validator container: %w", err))
		}
	})

	// Create and fund a wallet on Gaia for the relayer and a user acc
	chainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainAUserMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	chainAKey.Mnemonic = chainAUserMnemonic

	rlyChainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainAMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	rlyChainAKey.Mnemonic = rlyChainAMnemonic

	// Wait a few blocks to ensure the wallets are created and funded
	err = testutil.WaitForBlocks(ctx, 5, chainA)
	require.NoError(t, err)

	// Get our bech32 encoded user address
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)

	// Get the original acc balances on both chains for their native tokens
	chainAOrigBalNative, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainAOrigBalNative)

	chainBOrigBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBOrigBalNative)

	// Add chain configs to the relayer for both chains
	err = r.AddChainConfiguration(ctx, eRep, chainA.Config(), rlyChainAKey.KeyName, chainA.GetRPCAddress(), chainA.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainB.Config(), rlyChainBKey.KeyName, chainB.GetRPCAddress(), chainB.GetGRPCAddress())
	require.NoError(t, err)

	// Configure keys for the relayer to use for both chains
	err = r.RestoreKey(ctx, eRep, chainA.Config().ChainID, rlyChainAKey.KeyName, cosmosCoinType, rlyChainAKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainB.Config().ChainID, rlyChainBKey.KeyName, cosmosCoinType, rlyChainBKey.Mnemonic)
	require.NoError(t, err)

	// Create a new path in the relayer config for the Gaia<>Duality path
	err = r.GeneratePath(ctx, eRep, chainA.Config().ChainID, chainB.Config().ChainID, pathChainAChainB)
	require.NoError(t, err)

	// Link the path between Gaia and Duality
	err = r.LinkPath(ctx, eRep, pathChainAChainB, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathChainAChainB))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	// Get channel between Gaia and Duality
	channels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(channels))
	chainAChannel := channels[0]

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBKey.Address,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	// Send an IBC transfer from Gaia to Duality, so we can initialize a pool with the IBC denom token + native Duality token
	transferTx, err := chainA.SendIBCTransfer(ctx, chainAChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{
		Timeout: nil,
		Memo:    "",
	})
	require.NoError(t, err)

	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	// Poll for the ack to know that the transfer is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Get the IBC denom for ATOM on Duality
	chainATokenDenom := transfertypes.GetPrefixedDenom(chainAChannel.Counterparty.PortID, chainAChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Assert that the funds are gone from the acc on Gaia and present in the acc on Duality
	chainABalTransfer, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainAOrigBalNative-ibcTransferAmount, chainABalTransfer)

	chainBBalIBCTransfer, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, ibcTransferAmount, chainBBalIBCTransfer)

	depositAmount := sdktypes.NewInt(100000)

	depositCmd := []string{
		chainB.Config().Bin, "tx", "dex", "deposit",
		chainBKey.Address,
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
	chainBBalIBC, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBCTransfer-depositAmount.Int64(), chainBBalIBC)

	chainBBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBOrigBalNative-depositAmount.Int64(), chainBBalNative)

	swapAmount := sdktypes.NewInt(100000)
	minOut := sdktypes.NewInt(100000)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  chainBKey.Address,
				Receiver: chainBKey.Address,
				TokenA:   chainADenomTrace.IBCDenom(),
				TokenB:   chainB.Config().Denom,
				AmountIn: swapAmount,
				TokenIn:  chainADenomTrace.IBCDenom(),
				MinOut:   minOut,
			},
			Next: "",
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from Gaia to Duality with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, chainAChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on Gaia
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABalTransfer-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on Duality
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative+minOut.Int64(), chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)
}

// TestIBCSwapMiddleware_FailRefund asserts that the IBC swap middleware works as intended with Duality running as a
// standalone consumer chain connected to the Cosmos Hub. The swap should fail and a refund to the src chain should take place.
func TestIBCSwapMiddleware_FailRefund(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Number of full nodes and validators in the network
	nv := 1
	nf := 0

	// Create chain factory with Gaia and Duality
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "cosmoshub-4", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg, NumValidators: &nv, NumFullNodes: &nf}},
	)

	// Get both chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	ctx := context.Background()
	client, network := ibctest.DockerSetup(t)

	// Create relayer factory with the go-relayer
	// TODO the custom docker image can be removed here once ICS query fix is merged into main in the relayer
	r := ibctest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-ics_consumer_unbonding_period_query", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	// Initialize the Interchain object which describes the chains, relayers, and paths between chains
	// We only use this for spinning up Gaia and initializing the relayer config because there is no ICS support for Duality.
	ic := ibctest.NewInterchain().
		AddChain(chainA).
		AddRelayer(r, "relayer")

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	require.NoError(t, ic.Build(ctx, eRep, ibctest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: true,
	}))

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Initialize the Duality nodes
	err = chainB.Initialize(ctx, t.Name(), client, network)
	require.NoError(t, err, "failed to initialize duality chain")

	chainBValidator := chainB.Validators[0]

	// Initialize the Duality node files, create genesis wallets, and start the chain
	kr := keyring.NewInMemory()

	chainBWallets, err := initDuality(ctx, chainBValidator, kr, []string{aliceKeyName, rlyChainBKeyName})
	require.NoError(t, err)

	chainBKey, rlyChainBKey := chainBWallets[0], chainBWallets[1]

	t.Cleanup(func() {
		err = chainBValidator.StopContainer(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to stop duality validator container: %w", err))
		}
	})

	// Create and fund a wallet on Gaia for the relayer and a user acc
	chainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainAUserMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	chainAKey.Mnemonic = chainAUserMnemonic

	rlyChainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainAMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	rlyChainAKey.Mnemonic = rlyChainAMnemonic

	// Wait a few blocks to ensure the wallets are created and funded
	err = testutil.WaitForBlocks(ctx, 5, chainA)
	require.NoError(t, err)

	// Get our bech32 encoded user address
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)

	// Add chain configs to the relayer for both chains
	err = r.AddChainConfiguration(ctx, eRep, chainA.Config(), rlyChainAKey.KeyName, chainA.GetRPCAddress(), chainA.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainB.Config(), rlyChainBKey.KeyName, chainB.GetRPCAddress(), chainB.GetGRPCAddress())
	require.NoError(t, err)

	// Configure keys for the relayer to use for both chains
	err = r.RestoreKey(ctx, eRep, chainA.Config().ChainID, rlyChainAKey.KeyName, cosmosCoinType, rlyChainAKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainB.Config().ChainID, rlyChainBKey.KeyName, cosmosCoinType, rlyChainBKey.Mnemonic)
	require.NoError(t, err)

	// Create a new path in the relayer config for the Gaia<>Duality path
	err = r.GeneratePath(ctx, eRep, chainA.Config().ChainID, chainB.Config().ChainID, pathChainAChainB)
	require.NoError(t, err)

	// Link the path between Gaia and Duality
	err = r.LinkPath(ctx, eRep, pathChainAChainB, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathChainAChainB))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	// Get channel between Gaia and Duality
	channels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(channels))
	chainAChannel := channels[0]

	// Get the IBC denom for ATOM on Duality
	chainATokenDenom := transfertypes.GetPrefixedDenom(chainAChannel.Counterparty.PortID, chainAChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the acc balances before the transfer and swap takes place
	chainABal, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainABal)

	chainBNativeBal, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBNativeBal)

	chainBIBCBal, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainBIBCBal)

	// Compose the swap metadata, this swap will fail because there is no pool initialized for this pair
	swapAmount := sdktypes.NewInt(100000)
	minOut := sdktypes.NewInt(100000)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  chainBKey.Address,
				Receiver: chainBKey.Address,
				TokenA:   chainADenomTrace.IBCDenom(),
				TokenB:   chainB.Config().Denom,
				AmountIn: swapAmount,
				TokenIn:  chainADenomTrace.IBCDenom(),
				MinOut:   minOut,
			},
			NonRefundable: false,
			Next:          "",
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBKey.Address,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from Gaia to Duality with packet memo containing the swap metadata
	transferTx, err := chainA.SendIBCTransfer(ctx, chainAChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap has failed
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are not present in the account on Duality
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBNativeBal, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBIBCBal, chainBBalIBCSwap)

	// Check that the refund takes place and the funds are moved back to the account on Gaia
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABal, chainABalAfterSwap)
}

// TestIBCSwapMiddleware_FailNoRefund asserts that the IBC swap middleware works as intended with Duality running as a
// standalone consumer chain connected to the Cosmos Hub. The swap should fail and funds should remain on Duality.
func TestIBCSwapMiddleware_FailNoRefund(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Number of full nodes and validators in the network
	nv := 1
	nf := 0

	// Create chain factory with Gaia and Duality
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "cosmoshub-4", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg, NumValidators: &nv, NumFullNodes: &nf}},
	)

	// Get both chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	ctx := context.Background()
	client, network := ibctest.DockerSetup(t)

	// Create relayer factory with the go-relayer
	// TODO the custom docker image can be removed here once ICS query fix is merged into main in the relayer
	r := ibctest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-ics_consumer_unbonding_period_query", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	// Initialize the Interchain object which describes the chains, relayers, and paths between chains
	// We only use this for spinning up Gaia and initializing the relayer config because there is no ICS support for Duality.
	ic := ibctest.NewInterchain().
		AddChain(chainA).
		AddRelayer(r, "relayer")

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	require.NoError(t, ic.Build(ctx, eRep, ibctest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: true,
	}))

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Initialize the Duality nodes
	err = chainB.Initialize(ctx, t.Name(), client, network)
	require.NoError(t, err, "failed to initialize duality chain")

	chainBValidator := chainB.Validators[0]

	// Initialize the Duality node files, create genesis wallets, and start the chain
	kr := keyring.NewInMemory()

	chainBWallets, err := initDuality(ctx, chainBValidator, kr, []string{aliceKeyName, rlyChainBKeyName})
	require.NoError(t, err)

	chainBKey, rlyChainBKey := chainBWallets[0], chainBWallets[1]

	t.Cleanup(func() {
		err = chainBValidator.StopContainer(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to stop duality validator container: %w", err))
		}
	})

	// Create and fund a wallet on Gaia for the relayer and a user acc
	chainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainAUserMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	chainAKey.Mnemonic = chainAUserMnemonic

	rlyChainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainAMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	rlyChainAKey.Mnemonic = rlyChainAMnemonic

	// Wait a few blocks to ensure the wallets are created and funded
	err = testutil.WaitForBlocks(ctx, 5, chainA)
	require.NoError(t, err)

	// Get our bech32 encoded user address
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)

	// Add chain configs to the relayer for both chains
	err = r.AddChainConfiguration(ctx, eRep, chainA.Config(), rlyChainAKey.KeyName, chainA.GetRPCAddress(), chainA.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainB.Config(), rlyChainBKey.KeyName, chainB.GetRPCAddress(), chainB.GetGRPCAddress())
	require.NoError(t, err)

	// Configure keys for the relayer to use for both chains
	err = r.RestoreKey(ctx, eRep, chainA.Config().ChainID, rlyChainAKey.KeyName, cosmosCoinType, rlyChainAKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainB.Config().ChainID, rlyChainBKey.KeyName, cosmosCoinType, rlyChainBKey.Mnemonic)
	require.NoError(t, err)

	// Create a new path in the relayer config for the Gaia<>Duality path
	err = r.GeneratePath(ctx, eRep, chainA.Config().ChainID, chainB.Config().ChainID, pathChainAChainB)
	require.NoError(t, err)

	// Link the path between Gaia and Duality
	err = r.LinkPath(ctx, eRep, pathChainAChainB, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathChainAChainB))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	// Get channel between Gaia and Duality
	channels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(channels))
	chainAChannel := channels[0]

	// Get the IBC denom for ATOM on Duality
	chainATokenDenom := transfertypes.GetPrefixedDenom(chainAChannel.Counterparty.PortID, chainAChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the acc balances before the transfer and swap takes place
	chainABal, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainABal)

	chainBNativeBal, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBNativeBal)

	chainBIBCBal, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainBIBCBal)

	// Compose the swap metadata, this swap will fail because there is no pool initialized for this pair
	swapAmount := sdktypes.NewInt(100000)
	minOut := sdktypes.NewInt(100000)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  chainBKey.Address,
				Receiver: chainBKey.Address,
				TokenA:   chainADenomTrace.IBCDenom(),
				TokenB:   chainB.Config().Denom,
				AmountIn: swapAmount,
				TokenIn:  chainADenomTrace.IBCDenom(),
				MinOut:   minOut,
			},
			NonRefundable: true,
			Next:          "",
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBKey.Address,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from Gaia to Duality with packet memo containing the swap metadata
	transferTx, err := chainA.SendIBCTransfer(ctx, chainAChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap has failed
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are present in the account on Duality
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBNativeBal, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBIBCBal+ibcTransferAmount, chainBBalIBCSwap)

	// Check that no refund takes place and the funds are not in the account on Gaia
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABal-ibcTransferAmount, chainABalAfterSwap)
}
