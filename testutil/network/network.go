package network

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	pruningtypes "github.com/cosmos/cosmos-sdk/pruning/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutil "github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmdb "github.com/tendermint/tm-db"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/app"
	"github.com/stretchr/testify/require"
)

type (
	Network = network.Network
	Config  = network.Config
)

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func New(t *testing.T, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = DefaultConfig()
	} else {
		cfg = configs[0]
	}
	net, err := network.New(t, t.TempDir(), cfg)
	require.NoError(t, err)
	t.Cleanup(net.Cleanup)
	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfig() network.Config {

	encoding := app.MakeEncodingConfig()
	return network.Config{
		Codec:             encoding.Marshaler,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.Validator) servertypes.Application {
			return app.New(
				val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0,
				encoding,
				simapp.EmptyAppOptions{},
				baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
				baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
			)
		},
		GenesisState:    app.ModuleBasics.DefaultGenesis(encoding.Marshaler),
		TimeoutCommit:   2 * time.Second,
		ChainID:         "chain-" + tmrand.NewRand().Str(6),
		NumValidators:   2,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: pruningtypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
	}
}

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func NewCLITest(t *testing.T, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = DefaultConfigCLITest()
	} else {
		cfg = configs[0]
	}
	net, err := network.New(t, t.TempDir(), cfg)
	require.NoError(t, err)
	t.Cleanup(net.Cleanup)
	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfigCLITest() network.Config {

	encoding := app.MakeEncodingConfig()
	return network.Config{
		Codec:             encoding.Marshaler,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.Validator) servertypes.Application {
			return app.New(
				val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0,
				encoding,
				simapp.EmptyAppOptions{},
				baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
				baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
			)
		},
		GenesisState:    app.ModuleBasics.DefaultGenesis(encoding.Marshaler),
		TimeoutCommit:   2 * time.Second,
		ChainID:         "chain-" + tmrand.NewRand().Str(6),
		NumValidators:   2,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: pruningtypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
	}
}

func modifyConsumerGenesisCLITestSetup(val network.Validator) error {
	genFile := val.Ctx.Config.GenesisFile()
	appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
	if err != nil {
		return err
	}

	// Modify the data structure here
	dexData := appState[dextypes.ModuleName]
	var dexGenesisState dextypes.GenesisState
	json.Unmarshal(dexData, &dexGenesisState)

	dexGenesisState.FeeTierList = []dextypes.FeeTier{
		dextypes.FeeTier{
			Id:  0,
			Fee: 1,
		},
		dextypes.FeeTier{
			Id:  1,
			Fee: 3,
		},
		dextypes.FeeTier{
			Id:  2,
			Fee: 5,
		},
		dextypes.FeeTier{
			Id:  3,
			Fee: 10,
		},
	}
	dexGenesisState.FeeTierCount = 4

	newRawJSON, _ := json.Marshal(dexGenesisState)
	appState[dextypes.ModuleName] = newRawJSON

	bankData := appState[banktypes.ModuleName]
	var bankGenesisState banktypes.GenesisState
	json.Unmarshal(bankData, &bankGenesisState)

	bankGenesisState.Balances = []banktypes.Balance{
		banktypes.Balance{
			Address: val.Address.String(),
			Coins:   sdk.Coins{sdk.Coin{"TokenA", sdk.NewInt(100000000)}, sdk.Coin{"TokenB", sdk.NewInt(100000000)}, sdk.Coin{sdk.DefaultBondDenom, sdk.NewInt(10000)}},
		},
	}
	newRawJSON, _ = json.Marshal(bankGenesisState)
	appState[banktypes.ModuleName] = newRawJSON

	appStateJSON, _ := json.Marshal(appState)
	// Write the modified app state to the genesis file
	genDoc.AppState = json.RawMessage(appStateJSON)

	err = genutil.ExportGenesisFile(genDoc, genFile)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to export genesis state")
	}
	return nil
}
