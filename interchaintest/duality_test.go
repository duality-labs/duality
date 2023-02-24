package interchaintest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/icza/dyno"
	"github.com/strangelove-ventures/interchaintest/v3"
	"github.com/strangelove-ventures/interchaintest/v3/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v3/ibc"
	"github.com/strangelove-ventures/interchaintest/v3/relayer"
	"github.com/strangelove-ventures/interchaintest/v3/relayer/rly"
	"github.com/strangelove-ventures/interchaintest/v3/testreporter"

	"github.com/strangelove-ventures/interchaintest/v3/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

const (
	pathICS          = "provider-consumer"
	pathChainBChainC = "chainB-chainC"
	pathChainCChainD = "chainC-chainD"

	genesisWalletAmount = int64(100_000_000)
	ibcTransferAmount   = int64(100_000)

	heighlinerUserString = "1025:1025"

	cosmosCoinType = "118"
)

var (
	feeList = Fees{
		FeeList: []FeeTier{
			{0, 1},
			{Id: 1, Fee: 0},
			{Id: 2, Fee: 5},
			{Id: 3, Fee: 10},
		}}

	chainCfg = ibc.ChainConfig{
		Type:    "cosmos",
		Name:    "duality",
		ChainID: "chain-b",
		//Images: []ibc.DockerImage{{
		//	Repository: "ghcr.io/strangelove-ventures/heighliner/duality",
		//	Version:    "justin-ibc-swap",
		//	UidGid:     heighlinerUserString,
		//}},
		Images: []ibc.DockerImage{{
			Repository: "duality",
			Version:    "local",
			UidGid:     heighlinerUserString,
		}},
		Bin:            "dualityd",
		Bech32Prefix:   "cosmos",
		Denom:          "stake",
		CoinType:       cosmosCoinType,
		GasPrices:      "0.0stake",
		GasAdjustment:  1.2,
		TrustingPeriod: "336h",
		NoHostMount:    false,
		ModifyGenesis: func(config ibc.ChainConfig, bytes []byte) ([]byte, error) {
			return modifyGenesisDuality(bytes, feeList)
		},
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

	// Create chain factory with Duality and Cosmos Hub
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{Name: "duality", ChainConfig: chainCfg},
		{Name: "gaia", Version: "v9.0.0-rc1", ChainConfig: ibc.ChainConfig{ChainID: "chain-a", GasPrices: "0.0uatom"}}},
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	duality, gaia := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	ctx := context.Background()
	client, network := interchaintest.DockerSetup(t)

	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.CustomDockerImage("ghcr.io/cosmos/relayer", "andrew-paths_update", rly.RlyDefaultUidGid),
	).Build(t, client, network)

	ic := interchaintest.NewInterchain().
		AddChain(duality).
		AddChain(gaia).
		AddRelayer(r, "relayer").
		AddProviderConsumerLink(interchaintest.ProviderConsumerLink{
			Provider: gaia,
			Consumer: duality,
			Relayer:  r,
			Path:     pathICS,
		})

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: false,
	}))

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, duality)

	// Wait a block to ensure the chain is up and running
	err = testutil.WaitForBlocks(ctx, 1, duality)
	require.NoError(t, err)

	// Assert that the genesis wallet contains the specified balance from initialization
	bal, err := duality.GetBalance(ctx, users[0].Bech32Address(duality.Config().Bech32Prefix), duality.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, bal)
}

// dualityEncoding registers the Duality dex modules custom types, so we can see them in the block database.
func dualityEncoding() *simappparams.EncodingConfig {
	cfg := cosmos.DefaultEncoding()
	dextypes.RegisterInterfaces(cfg.InterfaceRegistry)
	return &cfg
}

func modifyGenesisDuality(genbz []byte, feeList Fees) ([]byte, error) {
	g := make(map[string]interface{})
	if err := json.Unmarshal(genbz, &g); err != nil {
		return nil, fmt.Errorf("failed to unmarshal genesis file: %w", err)
	}
	if err := dyno.Set(g, feeList.FeeList, "app_state", "dex", "feeTierList"); err != nil {
		return nil, fmt.Errorf("failed to set fee list in genesis json: %w", err)
	}
	if err := dyno.Set(g, len(feeList.FeeList), "app_state", "dex", "feeTierCount"); err != nil {
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
