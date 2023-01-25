package ibctest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/icza/dyno"
	"github.com/strangelove-ventures/ibctest/v3"
	"github.com/strangelove-ventures/ibctest/v3/chain/cosmos"
	"github.com/strangelove-ventures/ibctest/v3/ibc"

	"github.com/strangelove-ventures/ibctest/v3/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

const (
	aliceKeyName     = "alice"
	rlyChainBKeyName = "relayer-duality"

	rlyChainAMnemonic = "turkey sustain spoil ostrich false cradle tackle silent collect maple walnut brave rookie melody busy float monkey large drama romance rib search ride diary"
	rlyChainCMnemonic = "south document myth salon ribbon behave galaxy annual drama poem crime trick belt naive wedding open crunch ritual wrap clutch lumber capital cruel say"
	rlyChainDMnemonic = "toss door brown piece please foot regular globe depend symbol secret valve human window permit canvas frequent volume face clump cook canyon smooth unknown"

	chainAUserMnemonic = "obscure reform almost timber anxiety wave use shield choose icon crack visual bunker mountain wild range child cross wedding organ make tube oxygen talent"
	chainCUserMnemonic = "foil slogan drift judge scorpion hundred bomb tube praise attend south comic own update physical oil afford crime cage soon private found clip oak"
	chainDUserMnemonic = "boring paper simple peasant near virtual appear visit crowd orchard slide easy profit remind jacket donor argue guard ask proof special blush trade cross"

	pathChainAChainB = "chainA-chainB"
	pathChainBChainC = "chainB-chainC"
	pathChainCChainD = "chainC-chainD"

	genesisWalletAmount = int64(100_000_000)
	ibcTransferAmount   = int64(100_000)

	heighlinerUserString = "1025:1025"

	cosmosCoinType = "118"
)

var (
	chainCfg = ibc.ChainConfig{
		Type:    "cosmos",
		Name:    "duality",
		ChainID: "chain-b",
		Images: []ibc.DockerImage{{
			Repository: "duality",
			Version:    "local",
			UidGid:     heighlinerUserString,
		}},
		Bin:                 "dualityd",
		Bech32Prefix:        "cosmos",
		Denom:               "stake",
		CoinType:            cosmosCoinType,
		GasPrices:           "0.0stake",
		GasAdjustment:       1.2,
		TrustingPeriod:      "336h",
		NoHostMount:         false,
		ModifyGenesis:       nil,
		ConfigFileOverrides: nil,
		EncodingConfig:      dualityEncoding(),
	}
)

// TestDualityConsumerChainStart asserts that the chain can be properly spun up as a standalone consumer chain.
func TestDualityConsumerChainStart(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Number of full nodes and validators in the network
	nv := 1
	nf := 0

	// Create chain factory with Duality
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{Name: "duality", ChainConfig: chainCfg, NumValidators: &nv, NumFullNodes: &nf}},
	)

	// Get chain from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	duality := chains[0].(*cosmos.CosmosChain)

	ctx := context.Background()
	client, network := ibctest.DockerSetup(t)

	// Initialize the Duality nodes
	err = duality.Initialize(ctx, t.Name(), client, network)
	require.NoError(t, err, "failed to initialize duality chain")

	dualityValidator := duality.Validators[0]

	// Initialize the Duality node files, create genesis wallets, and start the chain
	kr := keyring.NewInMemory()

	dualityWallets, err := initDuality(ctx, dualityValidator, kr, []string{aliceKeyName})
	require.NoError(t, err)

	t.Cleanup(func() {
		err = dualityValidator.StopContainer(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to stop duality validator container: %w", err))
		}
	})

	// Wait a block to ensure the chain is up and running
	err = testutil.WaitForBlocks(ctx, 1, duality)
	require.NoError(t, err)

	// Assert that the genesis wallet contains the specified balance from initialization
	// Mostly just here to ensure we can now query state from the chain
	bal, err := duality.GetBalance(ctx, dualityWallets[0].Address, duality.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, bal)
}

// dualityEncoding registers the Duality dex modules custom types, so we can see them in the block database.
func dualityEncoding() *simappparams.EncodingConfig {
	cfg := cosmos.DefaultEncoding()
	dextypes.RegisterInterfaces(cfg.InterfaceRegistry)
	return &cfg
}

// initDuality creates and funds the genesis wallets, initializes the nodes, adds the standalone consumer chain
// data to the genesis file and starts the chain.
func initDuality(
	ctx context.Context,
	dualityValidator *cosmos.ChainNode,
	kr keyring.Keyring,
	keys []string,
) ([]ibc.Wallet, error) {
	userWallets := make([]ibc.Wallet, len(keys))

	// Generate wallet mnemonics and add to the keyring
	for i, key := range keys {
		wallet := ibctest.BuildWallet(kr, key, chainCfg)
		wallet.KeyName = key
		userWallets[i] = wallet

		err := dualityValidator.RecoverKey(ctx, key, wallet.Mnemonic)
		if err != nil {
			return nil, fmt.Errorf("failed to restore key for %s: %w", key, err)
		}
	}

	// Initialize the nodes on Duality
	err := dualityValidator.InitFullNodeFiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize duality validator config")
	}

	// Generate initial wallet balances to be used at genesis
	genesisWallets := make([]ibc.WalletAmount, len(userWallets))
	for i, wallet := range userWallets {
		genesisWallets[i] = ibc.WalletAmount{
			Address: wallet.Address,
			Denom:   chainCfg.Denom,
			Amount:  genesisWalletAmount,
		}
	}

	// Add genesis accounts for each wallet
	for _, wallet := range genesisWallets {
		err = dualityValidator.AddGenesisAccount(ctx, wallet.Address, []sdktypes.Coin{sdktypes.NewCoin(wallet.Denom, sdktypes.NewIntFromUint64(uint64(wallet.Amount)))})
		if err != nil {
			return nil, fmt.Errorf("failed to add genesis account for %s: %w", wallet.Address, err)
		}
	}

	// Read the genesis file, modify it, and then overwrite the genesis file on the node
	genBz, err := dualityValidator.GenesisFileContent(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to read duality genesis file: %w", err)
	}

	feeList := Fees{FeeList: []FeeTier{
		{0, 1},
		{Id: 1, Fee: 0},
		{Id: 2, Fee: 5},
		{Id: 3, Fee: 10},
	}}

	genBz, err = modifyGenesisDuality(genBz, feeList)
	if err != nil {
		return nil, fmt.Errorf("failed to modify duality genesis file: %w", err)
	}

	err = dualityValidator.OverwriteGenesisFile(ctx, genBz)
	if err != nil {
		return nil, fmt.Errorf("failed to write duality genesis file: %w", err)
	}

	// Execute the command to alter the genesis file for Duality to run as a standalone consumer chain
	_, _, err = dualityValidator.ExecBin(ctx, "add-consumer-section")
	if err != nil {
		return nil, fmt.Errorf("failed to add consumer section to duality validator genesis file: %w", err)
	}

	// Create and start the container for the single validator on Duality
	err = dualityValidator.CreateNodeContainer(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create duality validator container: %w", err)
	}

	err = dualityValidator.StartContainer(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create duality validator container: %w", err)
	}

	return userWallets, nil
}

func modifyGenesisDuality(genbz []byte, feeList Fees) ([]byte, error) {
	g := make(map[string]interface{})
	if err := json.Unmarshal(genbz, &g); err != nil {
		return nil, fmt.Errorf("failed to unmarshal genesis file: %w", err)
	}
	if err := dyno.Set(g, feeList.FeeList, "app_state", "dex", "FeeTierList"); err != nil {
		return nil, fmt.Errorf("failed to set fee list in genesis json: %w", err)
	}
	if err := dyno.Set(g, len(feeList.FeeList), "app_state", "dex", "FeeTierCount"); err != nil {
		return nil, fmt.Errorf("failed set fee list count in genesis json")
	}

	out, err := json.Marshal(g)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal genesis bytes to json: %w", err)
	}
	return out, nil

}

type Fees struct {
	FeeList []FeeTier `yaml:"feeListList"`
}

type FeeTier struct {
	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fee int64  `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
}
