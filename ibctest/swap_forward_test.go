package ibctest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/NicholasDotSol/duality/x/dex/types"
	swaptypes "github.com/NicholasDotSol/duality/x/ibc-swap/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/strangelove-ventures/ibctest/v3"
	"github.com/strangelove-ventures/ibctest/v3/chain/cosmos"
	"github.com/strangelove-ventures/ibctest/v3/ibc"
	"github.com/strangelove-ventures/ibctest/v3/relayer"
	"github.com/strangelove-ventures/ibctest/v3/relayer/rly"
	"github.com/strangelove-ventures/ibctest/v3/testreporter"
	"github.com/strangelove-ventures/ibctest/v3/testutil"
	forwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v3/router/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestSwapAndForward_Success asserts that the swap and forward middleware stack works as intended with Duality running as a
// standalone consumer chain connected to two other chains via IBC.
func TestSwapAndForward_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Number of full nodes and validators in the network
	nv := 1
	nf := 0

	// Create chain factory
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg, NumValidators: &nv, NumFullNodes: &nf},
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "chain-c", GasPrices: "0.0uatom"}}},
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB, chainC := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain), chains[2].(*cosmos.CosmosChain)

	// Create relayer factory with the go-relayer
	// TODO the custom docker image can be removed here once ICS query fix is merged into default relayer version
	ctx := context.Background()
	client, network := ibctest.DockerSetup(t)

	r := ibctest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-ics_consumer_unbonding_period_query", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	// Initialize the Interchain object which describes the chains, relayers, and paths between chains
	// We use this for spinning up chainA and chainC and initializing the relayer config.
	ic := ibctest.NewInterchain().
		AddChain(chainA).
		AddChain(chainC).
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

	// Create and fund user and relayer wallets on ChainA and ChainC
	chainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainAUserMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	chainAKey.Mnemonic = chainAUserMnemonic

	rlyChainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainAMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	rlyChainAKey.Mnemonic = rlyChainAMnemonic

	chainCKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainCUserMnemonic, genesisWalletAmount, chainC)
	require.NoError(t, err)
	chainCKey.Mnemonic = chainCUserMnemonic

	rlyChainCKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainCMnemonic, genesisWalletAmount, chainC)
	require.NoError(t, err)
	rlyChainCKey.Mnemonic = rlyChainCMnemonic

	// Wait a few blocks to ensure the wallets are created and funded
	err = testutil.WaitForBlocks(ctx, 5, chainA, chainC)
	require.NoError(t, err)

	// Get our bech32 encoded user addresses
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainCAddr := chainCKey.Bech32Address(chainC.Config().Bech32Prefix)

	// Get the original acc balances for each chains respective native token
	chainAOrigBalNative, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainAOrigBalNative)

	chainBOrigBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBOrigBalNative)

	// Add chain configs to the relayer for all three chains
	err = r.AddChainConfiguration(ctx, eRep, chainA.Config(), rlyChainAKey.KeyName, chainA.GetRPCAddress(), chainA.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainB.Config(), rlyChainBKey.KeyName, chainB.GetRPCAddress(), chainB.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainC.Config(), rlyChainCKey.KeyName, chainC.GetRPCAddress(), chainC.GetGRPCAddress())
	require.NoError(t, err)

	// Configure keys for the relayer to use for all three chains
	err = r.RestoreKey(ctx, eRep, chainA.Config().ChainID, rlyChainAKey.KeyName, cosmosCoinType, rlyChainAKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainB.Config().ChainID, rlyChainBKey.KeyName, cosmosCoinType, rlyChainBKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainC.Config().ChainID, rlyChainCKey.KeyName, cosmosCoinType, rlyChainCKey.Mnemonic)
	require.NoError(t, err)

	// Create new paths in the relayer config
	err = r.GeneratePath(ctx, eRep, chainA.Config().ChainID, chainB.Config().ChainID, pathChainAChainB)
	require.NoError(t, err)

	err = r.GeneratePath(ctx, eRep, chainC.Config().ChainID, chainB.Config().ChainID, pathChainBChainC)
	require.NoError(t, err)

	// Link the paths
	err = r.LinkPath(ctx, eRep, pathChainAChainB, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	err = r.LinkPath(ctx, eRep, pathChainBChainC, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathChainAChainB, pathChainBChainC))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	// Get channels between chainA-chainB and chainB-chainC
	chainAChannels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainAChannels))
	chainAChannel := chainAChannels[0]

	chainCChannels, err := r.GetChannels(ctx, eRep, chainC.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainCChannels))
	chainCChannel := chainCChannels[0]

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBKey.Address,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	// Send an IBC transfer from chainA to chainB, so we can initialize a pool with the IBC denom token + native Duality token
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

	// Get the IBC denom on chainB for the native token from chainA
	chainATokenDenom := transfertypes.GetPrefixedDenom(chainAChannel.Counterparty.PortID, chainAChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the IBC denom on chainC for the native token from chainB
	chainCTokenDenom := transfertypes.GetPrefixedDenom(chainCChannel.PortID, chainCChannel.ChannelID, chainB.Config().Denom)
	chainCDenomTrace := transfertypes.ParseDenomTrace(chainCTokenDenom)

	// Assert that the funds are gone from the acc on chainA and present in the acc on chainB
	chainABalTransfer, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainAOrigBalNative-ibcTransferAmount, chainABalTransfer)

	chainBBalTransfer, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, ibcTransferAmount, chainBBalTransfer)

	// Compose the deposit cmd for initializing a pool on Duality (chainB)
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
	require.Equal(t, chainBBalTransfer-depositAmount.Int64(), chainBBalIBC)

	chainBBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBOrigBalNative-depositAmount.Int64(), chainBBalNative)

	// Compose the IBC transfer memo metadata to be used in the swap and forward
	swapAmount := sdktypes.NewInt(100000)
	minOut := sdktypes.NewInt(100000)

	retries := uint8(0)
	forwardMetadata := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainCAddr,
			Port:     chainCChannel.Counterparty.PortID,
			Channel:  chainCChannel.Counterparty.ChannelID,
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     nil,
		}}

	bz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

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
			Next: string(bz),
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, chainAChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap and forward is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on chainA
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABalTransfer-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on chainC
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)

	chainCBal, err := chainC.GetBalance(ctx, chainCAddr, chainCDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, minOut.Int64(), chainCBal)
}

// TestSwapAndForward_MultiHopSuccess asserts that the swap and forward middleware stack works as intended in the case
// pf a multi-hop forward after the swap.
func TestSwapAndForward_MultiHopSuccess(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Number of full nodes and validators in the network
	nv := 1
	nf := 0

	// Create chain factory
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{Name: "gaia", Version: "justin-forward_middleware_v3_refactor", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg, NumValidators: &nv, NumFullNodes: &nf},
		{Name: "gaia", Version: "justin-forward_middleware_v3_refactor", ChainConfig: ibc.ChainConfig{ChainID: "chain-c", GasPrices: "0.0uatom"}},
		{Name: "gaia", Version: "justin-forward_middleware_v3_refactor", ChainConfig: ibc.ChainConfig{ChainID: "chain-d", GasPrices: "0.0uatom"}}},
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB, chainC, chainD := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain), chains[2].(*cosmos.CosmosChain), chains[3].(*cosmos.CosmosChain)

	// Create relayer factory with the go-relayer
	// TODO the custom docker image can be removed here once ICS query fix is merged into default relayer version
	ctx := context.Background()
	client, network := ibctest.DockerSetup(t)

	r := ibctest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-ics_consumer_unbonding_period_query", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	// Initialize the Interchain object which describes the chains, relayers, and paths between chains
	// We use this for spinning up chainA, chainC, and chainD and initializing the relayer config.
	ic := ibctest.NewInterchain().
		AddChain(chainA).
		AddChain(chainC).
		AddChain(chainD).
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

	// Create and fund user and relayer wallets on ChainA, ChainC, and ChainD
	chainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainAUserMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	chainAKey.Mnemonic = chainAUserMnemonic

	rlyChainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainAMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	rlyChainAKey.Mnemonic = rlyChainAMnemonic

	chainCKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainCUserMnemonic, genesisWalletAmount, chainC)
	require.NoError(t, err)
	chainCKey.Mnemonic = chainCUserMnemonic

	rlyChainCKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainCMnemonic, genesisWalletAmount, chainC)
	require.NoError(t, err)
	rlyChainCKey.Mnemonic = rlyChainCMnemonic

	chainDKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainDUserMnemonic, genesisWalletAmount, chainD)
	require.NoError(t, err)
	chainDKey.Mnemonic = chainDUserMnemonic

	rlyChainDKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainDMnemonic, genesisWalletAmount, chainD)
	require.NoError(t, err)
	rlyChainDKey.Mnemonic = rlyChainDMnemonic

	// Wait a few blocks to ensure the wallets are created and funded
	err = testutil.WaitForBlocks(ctx, 5, chainA, chainC)
	require.NoError(t, err)

	// Get our bech32 encoded user addresses
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainCAddr := chainCKey.Bech32Address(chainC.Config().Bech32Prefix)
	chainDAddr := chainDKey.Bech32Address(chainD.Config().Bech32Prefix)

	// Get the original acc balances for each chains respective native token
	chainAOrigBalNative, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainAOrigBalNative)

	chainBOrigBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBOrigBalNative)

	// Add chain configs to the relayer for all chains
	err = r.AddChainConfiguration(ctx, eRep, chainA.Config(), rlyChainAKey.KeyName, chainA.GetRPCAddress(), chainA.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainB.Config(), rlyChainBKey.KeyName, chainB.GetRPCAddress(), chainB.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainC.Config(), rlyChainCKey.KeyName, chainC.GetRPCAddress(), chainC.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainD.Config(), rlyChainDKey.KeyName, chainD.GetRPCAddress(), chainD.GetGRPCAddress())
	require.NoError(t, err)

	// Configure keys for the relayer to use for all chains
	err = r.RestoreKey(ctx, eRep, chainA.Config().ChainID, rlyChainAKey.KeyName, cosmosCoinType, rlyChainAKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainB.Config().ChainID, rlyChainBKey.KeyName, cosmosCoinType, rlyChainBKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainC.Config().ChainID, rlyChainCKey.KeyName, cosmosCoinType, rlyChainCKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainD.Config().ChainID, rlyChainDKey.KeyName, cosmosCoinType, rlyChainDKey.Mnemonic)
	require.NoError(t, err)

	// Create new paths in the relayer config
	err = r.GeneratePath(ctx, eRep, chainA.Config().ChainID, chainB.Config().ChainID, pathChainAChainB)
	require.NoError(t, err)

	err = r.GeneratePath(ctx, eRep, chainC.Config().ChainID, chainB.Config().ChainID, pathChainBChainC)
	require.NoError(t, err)

	err = r.GeneratePath(ctx, eRep, chainC.Config().ChainID, chainD.Config().ChainID, pathChainCChainD)
	require.NoError(t, err)

	// Link the paths
	err = r.LinkPath(ctx, eRep, pathChainAChainB, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	err = r.LinkPath(ctx, eRep, pathChainBChainC, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	err = r.LinkPath(ctx, eRep, pathChainCChainD, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathChainAChainB, pathChainBChainC, pathChainCChainD))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	// Get channels between chainA-chainB, chainB-chainC, and chainC-chainD
	chainAChannels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainAChannels))
	chainAChannel := chainAChannels[0]

	chainCChannels, err := r.GetChannels(ctx, eRep, chainC.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 2, len(chainCChannels))
	chainCChannel := chainCChannels[0]

	chainDChannels, err := r.GetChannels(ctx, eRep, chainD.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainDChannels))
	chainDChannel := chainDChannels[0]

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBKey.Address,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	// Send an IBC transfer from chainA to chainB, so we can initialize a pool with the IBC denom token + native Duality token
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

	// Get the IBC denom on chainB for the native token from chainA
	chainATokenDenom := transfertypes.GetPrefixedDenom(chainAChannel.Counterparty.PortID, chainAChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the IBC denom on chainC for the native token from chainB
	chainCTokenDenom := transfertypes.GetPrefixedDenom(chainCChannel.PortID, chainCChannel.ChannelID, chainB.Config().Denom)
	chainCDenomTrace := transfertypes.ParseDenomTrace(chainCTokenDenom)

	// Get the IBC denom on chainD for the native token from chainB that has travelled from chainB->chainC->chainD
	chainDTokenDenom := transfertypes.GetPrefixedDenom(chainDChannel.PortID, chainDChannel.ChannelID, chainCTokenDenom)
	chainDDenomTrace := transfertypes.ParseDenomTrace(chainDTokenDenom)

	// Assert that the funds are gone from the acc on chainA and present in the acc on chainB
	chainABalTransfer, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainAOrigBalNative-ibcTransferAmount, chainABalTransfer)

	chainBBalTransfer, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, ibcTransferAmount, chainBBalTransfer)

	// Compose the deposit cmd for initializing a pool on Duality (chainB)
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
	require.Equal(t, chainBBalTransfer-depositAmount.Int64(), chainBBalIBC)

	chainBBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBOrigBalNative-depositAmount.Int64(), chainBBalNative)

	// Compose the IBC transfer memo metadata to be used in the swap and forward
	swapAmount := sdktypes.NewInt(100000)
	minOut := sdktypes.NewInt(100000)

	retries := uint8(0)
	nextForward := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainDAddr,
			Port:     chainDChannel.Counterparty.PortID,
			Channel:  chainDChannel.Counterparty.ChannelID,
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     nil,
		},
	}
	nextForwardBz, err := json.Marshal(nextForward)
	require.NoError(t, err)
	nextForwardStr := string(nextForwardBz)

	forwardMetadata := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainCAddr,
			Port:     chainCChannel.Counterparty.PortID,
			Channel:  chainCChannel.Counterparty.ChannelID,
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     &nextForwardStr,
		}}
	bz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

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
			Next: string(bz),
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, chainAChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap and forward is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on chainA
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABalTransfer-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on chainD
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)

	chainCBal, err := chainC.GetBalance(ctx, chainCAddr, chainCDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainCBal)

	chainDBal, err := chainD.GetBalance(ctx, chainDAddr, chainDDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, minOut.Int64(), chainDBal)
}

// TestSwapAndForward_UnwindIBCDenomSuccess asserts that the swap and forward middleware stack works as intended in the
// case that a native token from ChainB is sent to ChainA and then ChainA initiates a swap and forward with the token.
// This asserts that denom unwinding works as intended when going ChainB->ChainA->ChainB.
func TestSwapAndForward_UnwindIBCDenomSuccess(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Number of full nodes and validators in the network
	nv := 1
	nf := 0

	// Create chain factory
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg, NumValidators: &nv, NumFullNodes: &nf},
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "chain-c", GasPrices: "0.0uatom"}}},
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB, chainC := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain), chains[2].(*cosmos.CosmosChain)

	// Create relayer factory with the go-relayer
	// TODO the custom docker image can be removed here once ICS query fix is merged into default relayer version
	ctx := context.Background()
	client, network := ibctest.DockerSetup(t)

	r := ibctest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-ics_consumer_unbonding_period_query", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	// Initialize the Interchain object which describes the chains, relayers, and paths between chains
	// We use this for spinning up chainA and chainC and initializing the relayer config.
	ic := ibctest.NewInterchain().
		AddChain(chainA).
		AddChain(chainC).
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

	// Create and fund user and relayer wallets on ChainA and ChainC
	chainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainAUserMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	chainAKey.Mnemonic = chainAUserMnemonic

	rlyChainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainAMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	rlyChainAKey.Mnemonic = rlyChainAMnemonic

	chainCKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainCUserMnemonic, genesisWalletAmount, chainC)
	require.NoError(t, err)
	chainCKey.Mnemonic = chainCUserMnemonic

	rlyChainCKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainCMnemonic, genesisWalletAmount, chainC)
	require.NoError(t, err)
	rlyChainCKey.Mnemonic = rlyChainCMnemonic

	// Wait a few blocks to ensure the wallets are created and funded
	err = testutil.WaitForBlocks(ctx, 5, chainA, chainC)
	require.NoError(t, err)

	// Get our bech32 encoded user addresses
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainCAddr := chainCKey.Bech32Address(chainC.Config().Bech32Prefix)

	// Get the original acc balances for each chains respective native token
	chainAOrigBalNative, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainAOrigBalNative)

	chainBOrigBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBOrigBalNative)

	// Add chain configs to the relayer for all three chains
	err = r.AddChainConfiguration(ctx, eRep, chainA.Config(), rlyChainAKey.KeyName, chainA.GetRPCAddress(), chainA.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainB.Config(), rlyChainBKey.KeyName, chainB.GetRPCAddress(), chainB.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainC.Config(), rlyChainCKey.KeyName, chainC.GetRPCAddress(), chainC.GetGRPCAddress())
	require.NoError(t, err)

	// Configure keys for the relayer to use for all three chains
	err = r.RestoreKey(ctx, eRep, chainA.Config().ChainID, rlyChainAKey.KeyName, cosmosCoinType, rlyChainAKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainB.Config().ChainID, rlyChainBKey.KeyName, cosmosCoinType, rlyChainBKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainC.Config().ChainID, rlyChainCKey.KeyName, cosmosCoinType, rlyChainCKey.Mnemonic)
	require.NoError(t, err)

	// Create new paths in the relayer config
	err = r.GeneratePath(ctx, eRep, chainA.Config().ChainID, chainB.Config().ChainID, pathChainAChainB)
	require.NoError(t, err)

	err = r.GeneratePath(ctx, eRep, chainC.Config().ChainID, chainB.Config().ChainID, pathChainBChainC)
	require.NoError(t, err)

	// Link the paths
	err = r.LinkPath(ctx, eRep, pathChainAChainB, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	err = r.LinkPath(ctx, eRep, pathChainBChainC, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathChainAChainB, pathChainBChainC))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	// Get channels between chainA-chainB and chainB-chainC
	chainAChannels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainAChannels))
	chainAChannel := chainAChannels[0]

	chainCChannels, err := r.GetChannels(ctx, eRep, chainC.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainCChannels))
	chainCChannel := chainCChannels[0]

	// Compose details for an IBC transfer
	tmpTransferAmount := int64(1000000)
	transfer := ibc.WalletAmount{
		Address: chainBKey.Address,
		Denom:   chainA.Config().Denom,
		Amount:  tmpTransferAmount,
	}

	// Send an IBC transfer from chainA to chainB, so we can initialize a pool with the IBC denom token + native Duality token
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

	// Get the IBC denom on chainB for the native token from chainA
	chainATokenDenom := transfertypes.GetPrefixedDenom(chainAChannel.Counterparty.PortID, chainAChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the IBC denom on chainC for the native token from chainA that travelled from chainA->chainB->chainC
	chainCTokenDenom := transfertypes.GetPrefixedDenom(chainCChannel.PortID, chainCChannel.ChannelID, chainATokenDenom)
	chainCDenomTrace := transfertypes.ParseDenomTrace(chainCTokenDenom)

	// Get the IBC denom on chainC for the native token from chainB
	//chainCTokenDenom := transfertypes.GetPrefixedDenom(chainCChannel.PortID, chainCChannel.ChannelID, chainB.Config().Denom)
	//chainCDenomTrace := transfertypes.ParseDenomTrace(chainCTokenDenom)

	// Assert that the funds are gone from the acc on chainA and present in the acc on chainB
	chainABalTransfer, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainAOrigBalNative-tmpTransferAmount, chainABalTransfer)

	chainBBalTransfer, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, tmpTransferAmount, chainBBalTransfer)

	// Compose the deposit cmd for initializing a pool on Duality (chainB)
	depositAmount := sdktypes.NewInt(1000000)

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
	require.Equal(t, chainBBalTransfer-depositAmount.Int64(), chainBBalIBC)

	chainBBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBOrigBalNative-depositAmount.Int64(), chainBBalNative)

	// Compose a transfer from ChainB->ChainA for ChainB's native token
	chainBTransfer := ibc.WalletAmount{
		Address: chainAAddr,
		Denom:   chainB.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	transferTx, err = chainB.SendIBCTransfer(ctx, chainAChannel.Counterparty.ChannelID, chainBKey.Address, chainBTransfer, ibc.TransferOptions{
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
	chainBTokenDenom := transfertypes.GetPrefixedDenom(chainAChannel.PortID, chainAChannel.ChannelID, chainB.Config().Denom)
	chainBDenomTrace := transfertypes.ParseDenomTrace(chainBTokenDenom)

	// Assert that the funds are present in the acc on chainA
	chainABal, err := chainA.GetBalance(ctx, chainAAddr, chainBDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, ibcTransferAmount, chainABal)

	// Assert that the funds are gone from the acc on chainB
	chainBBal, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative-ibcTransferAmount, chainBBal)

	// Compose the IBC transfer memo metadata to be used in the swap and forward
	swapAmount := sdktypes.NewInt(100_000)
	minOut := sdktypes.NewInt(100_000)

	retries := uint8(0)
	forwardMetadata := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainCAddr,
			Port:     chainCChannel.Counterparty.PortID,
			Channel:  chainCChannel.Counterparty.ChannelID,
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     nil,
		}}

	bz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

	metadata := swaptypes.PacketMetadata{
		Swap: &swaptypes.SwapMetadata{
			MsgSwap: &types.MsgSwap{
				Creator:  chainBKey.Address,
				Receiver: chainBKey.Address,
				TokenA:   chainB.Config().Denom,
				TokenB:   chainADenomTrace.IBCDenom(),
				AmountIn: swapAmount,
				TokenIn:  chainB.Config().Denom,
				MinOut:   minOut,
			},
			Next: string(bz),
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	swapTransfer := ibc.WalletAmount{
		Address: chainBKey.Address,
		Denom:   chainBDenomTrace.IBCDenom(),
		Amount:  ibcTransferAmount,
	}

	t.Logf("CHAIN A DENOM: %s \n", chainADenomTrace.IBCDenom())
	t.Logf("CHAIN B DENOM: %s \n", chainBDenomTrace.IBCDenom())

	/*
	   swap_forward_test.go:1033: CHAIN A DENOM: ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2
	   swap_forward_test.go:1034: CHAIN B DENOM: ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878
	*/

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, chainAChannel.ChannelID, chainAAddr, swapTransfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap and forward is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on chainA
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainBDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainABal-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on chainC
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBal, chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)

	chainCBal, err := chainC.GetBalance(ctx, chainCAddr, chainCDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, minOut.Int64(), chainCBal)
}

// TestSwapAndForward_ForwardFails asserts that the swap and forward middleware stack works as intended in the case
// that an incoming IBC swap succeeds but the forward fails.
func TestSwapAndForward_ForwardFails(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Number of full nodes and validators in the network
	nv := 1
	nf := 0

	// Create chain factory
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}},
		{Name: "duality", ChainConfig: chainCfg, NumValidators: &nv, NumFullNodes: &nf},
		{Name: "gaia", Version: "v8.0.0-rc3", ChainConfig: ibc.ChainConfig{ChainID: "chain-c", GasPrices: "0.0uatom"}}},
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB, chainC := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain), chains[2].(*cosmos.CosmosChain)

	// Create relayer factory with the go-relayer
	// TODO the custom docker image can be removed here once ICS query fix is merged into default relayer version
	ctx := context.Background()
	client, network := ibctest.DockerSetup(t)

	r := ibctest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-ics_consumer_unbonding_period_query", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	// Initialize the Interchain object which describes the chains, relayers, and paths between chains
	// We use this for spinning up chainA and chainC and initializing the relayer config.
	ic := ibctest.NewInterchain().
		AddChain(chainA).
		AddChain(chainC).
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

	// Create and fund user and relayer wallets on ChainA and ChainC
	chainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainAUserMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	chainAKey.Mnemonic = chainAUserMnemonic

	rlyChainAKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainAMnemonic, genesisWalletAmount, chainA)
	require.NoError(t, err)
	rlyChainAKey.Mnemonic = rlyChainAMnemonic

	chainCKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), chainCUserMnemonic, genesisWalletAmount, chainC)
	require.NoError(t, err)
	chainCKey.Mnemonic = chainCUserMnemonic

	rlyChainCKey, err := ibctest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), rlyChainCMnemonic, genesisWalletAmount, chainC)
	require.NoError(t, err)
	rlyChainCKey.Mnemonic = rlyChainCMnemonic

	// Wait a few blocks to ensure the wallets are created and funded
	err = testutil.WaitForBlocks(ctx, 5, chainA, chainC)
	require.NoError(t, err)

	// Get our bech32 encoded user addresses
	chainAAddr := chainAKey.Bech32Address(chainA.Config().Bech32Prefix)
	chainCAddr := chainCKey.Bech32Address(chainC.Config().Bech32Prefix)

	// Get the original acc balances for each chains respective native token
	chainAOrigBalNative, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainAOrigBalNative)

	chainBOrigBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, chainBOrigBalNative)

	// Add chain configs to the relayer for all three chains
	err = r.AddChainConfiguration(ctx, eRep, chainA.Config(), rlyChainAKey.KeyName, chainA.GetRPCAddress(), chainA.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainB.Config(), rlyChainBKey.KeyName, chainB.GetRPCAddress(), chainB.GetGRPCAddress())
	require.NoError(t, err)

	err = r.AddChainConfiguration(ctx, eRep, chainC.Config(), rlyChainCKey.KeyName, chainC.GetRPCAddress(), chainC.GetGRPCAddress())
	require.NoError(t, err)

	// Configure keys for the relayer to use for all three chains
	err = r.RestoreKey(ctx, eRep, chainA.Config().ChainID, rlyChainAKey.KeyName, cosmosCoinType, rlyChainAKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainB.Config().ChainID, rlyChainBKey.KeyName, cosmosCoinType, rlyChainBKey.Mnemonic)
	require.NoError(t, err)

	err = r.RestoreKey(ctx, eRep, chainC.Config().ChainID, rlyChainCKey.KeyName, cosmosCoinType, rlyChainCKey.Mnemonic)
	require.NoError(t, err)

	// Create new paths in the relayer config
	err = r.GeneratePath(ctx, eRep, chainA.Config().ChainID, chainB.Config().ChainID, pathChainAChainB)
	require.NoError(t, err)

	err = r.GeneratePath(ctx, eRep, chainC.Config().ChainID, chainB.Config().ChainID, pathChainBChainC)
	require.NoError(t, err)

	// Link the paths
	err = r.LinkPath(ctx, eRep, pathChainAChainB, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	err = r.LinkPath(ctx, eRep, pathChainBChainC, ibc.DefaultChannelOpts(), ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathChainAChainB, pathChainBChainC))

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occured while stopping the relayer: %s", err))
			}
		},
	)

	// Get channels between chainA-chainB and chainB-chainC
	chainAChannels, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainAChannels))
	chainAChannel := chainAChannels[0]

	chainCChannels, err := r.GetChannels(ctx, eRep, chainC.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(chainCChannels))
	chainCChannel := chainCChannels[0]

	// Compose details for an IBC transfer
	transfer := ibc.WalletAmount{
		Address: chainBKey.Address,
		Denom:   chainA.Config().Denom,
		Amount:  ibcTransferAmount,
	}

	// Send an IBC transfer from chainA to chainB, so we can initialize a pool with the IBC denom token + native Duality token
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

	// Get the IBC denom on chainB for the native token from chainA
	chainATokenDenom := transfertypes.GetPrefixedDenom(chainAChannel.Counterparty.PortID, chainAChannel.Counterparty.ChannelID, chainA.Config().Denom)
	chainADenomTrace := transfertypes.ParseDenomTrace(chainATokenDenom)

	// Get the IBC denom on chainC for the native token from chainB
	chainCTokenDenom := transfertypes.GetPrefixedDenom(chainCChannel.PortID, chainCChannel.ChannelID, chainB.Config().Denom)
	chainCDenomTrace := transfertypes.ParseDenomTrace(chainCTokenDenom)

	// Assert that the funds are gone from the acc on chainA and present in the acc on chainB
	chainABalTransfer, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainAOrigBalNative-ibcTransferAmount, chainABalTransfer)

	chainBBalTransfer, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, ibcTransferAmount, chainBBalTransfer)

	// Compose the deposit cmd for initializing a pool on Duality (chainB)
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
	require.Equal(t, chainBBalTransfer-depositAmount.Int64(), chainBBalIBC)

	chainBBalNative, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBOrigBalNative-depositAmount.Int64(), chainBBalNative)

	// Compose the IBC transfer memo metadata to be used in the swap and forward
	swapAmount := sdktypes.NewInt(100000)
	minOut := sdktypes.NewInt(100000)

	retries := uint8(0)
	forwardMetadata := forwardtypes.PacketMetadata{
		Forward: &forwardtypes.ForwardMetadata{
			Receiver: chainCAddr,
			Port:     chainCChannel.Counterparty.PortID,
			Channel:  "invalid-channel", // add an invalid channel identifier so the forward fails
			Timeout:  5 * time.Minute,
			Retries:  &retries,
			Next:     nil,
		}}

	bz, err := json.Marshal(forwardMetadata)
	require.NoError(t, err)

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
			Next: string(bz),
		},
	}

	metadataBz, err := json.Marshal(metadata)
	require.NoError(t, err)

	chainAHeight, err = chainA.Height(ctx)
	require.NoError(t, err)

	// Send an IBC transfer from chainA to chainB with packet memo containing the swap metadata
	transferTx, err = chainA.SendIBCTransfer(ctx, chainAChannel.ChannelID, chainAAddr, transfer, ibc.TransferOptions{Memo: string(metadataBz)})
	require.NoError(t, err)

	// Poll for the ack to know that the swap and forward is complete
	_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+20, transferTx.Packet)
	require.NoError(t, err)

	// Check that the funds are moved out of the acc on chainA
	chainABalAfterSwap, err := chainA.GetBalance(ctx, chainAAddr, chainA.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainABalTransfer-ibcTransferAmount, chainABalAfterSwap)

	// Check that the funds are now present in the acc on chainB
	chainBBalNativeSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainB.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, chainBBalNative+minOut.Int64(), chainBBalNativeSwap)

	chainBBalIBCSwap, err := chainB.GetBalance(ctx, chainBKey.Address, chainADenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, chainBBalIBC, chainBBalIBCSwap)

	chainCBal, err := chainC.GetBalance(ctx, chainCAddr, chainCDenomTrace.IBCDenom())
	require.NoError(t, err)
	require.Equal(t, int64(0), chainCBal)
}
