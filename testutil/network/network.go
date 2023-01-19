package network

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	dextypes "github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutil "github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ignite-hq/cli/ignite/pkg/cosmoscmd"
	types1 "github.com/tendermint/tendermint/abci/types"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmtypes "github.com/tendermint/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/testutil"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ccvconsumertypes "github.com/cosmos/interchain-security/x/ccv/consumer/types"
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
	net := network.New(t, cfg)
	t.Cleanup(net.Cleanup)
	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfig() network.Config {
	// app doesn't have this modules anymore, but we need them for test setup, which uses gentx and MsgCreateValidator
	app.ModuleBasics[genutiltypes.ModuleName] = genutil.AppModuleBasic{}
	app.ModuleBasics[stakingtypes.ModuleName] = staking.AppModuleBasic{}

	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	return network.Config{
		Codec:             encoding.Marshaler,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.Validator) servertypes.Application {
			err := modifyConsumerGenesis(val)
			if err != nil {
				panic(err)
			}
			return app.New(
				val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0,
				encoding,
				simapp.EmptyAppOptions{},
				baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
				baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
			)
		},
		GenesisState:  app.ModuleBasics.DefaultGenesis(encoding.Marshaler),
		TimeoutCommit: 2 * time.Second,
		ChainID:       "chain-" + tmrand.NewRand().Str(6),
		// Some changes are introduced to make the tests run as if Duality is a standalone chain.
		// This will only work if NumValidators is set to 1.
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: storetypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
	}
}

func modifyConsumerGenesis(val network.Validator) error {
	genFile := val.Ctx.Config.GenesisFile()
	appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to read genesis from the file")
	}

	tmProtoPublicKey, err := cryptocodec.ToTmProtoPublicKey(val.PubKey)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid public key")
	}

	initialValset := []types1.ValidatorUpdate{{PubKey: tmProtoPublicKey, Power: 100}}
	vals, err := tmtypes.PB2TM.ValidatorUpdates(initialValset)
	if err != nil {
		return sdkerrors.Wrap(err, "could not convert val updates to validator set")
	}

	consumerGenesisState := testutil.CreateMinimalConsumerTestGenesis()
	consumerGenesisState.InitialValSet = initialValset
	consumerGenesisState.ProviderConsensusState.NextValidatorsHash = tmtypes.NewValidatorSet(vals).Hash()

	if err := consumerGenesisState.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid consumer genesis")
	}

	consumerGenStateBz, err := val.ClientCtx.Codec.MarshalJSON(consumerGenesisState)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to marshal consumer genesis state into JSON")
	}

	appState[ccvconsumertypes.ModuleName] = consumerGenStateBz
	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to marshal application genesis state into JSON")
	}

	genDoc.AppState = appStateJSON
	err = genutil.ExportGenesisFile(genDoc, genFile)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to export genesis state")
	}

	return nil
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
	net := network.New(t, cfg)
	t.Cleanup(net.Cleanup)
	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfigCLITest() network.Config {
	// app doesn't have this modules anymore, but we need them for test setup, which uses gentx and MsgCreateValidator
	app.ModuleBasics[genutiltypes.ModuleName] = genutil.AppModuleBasic{}
	app.ModuleBasics[stakingtypes.ModuleName] = staking.AppModuleBasic{}

	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	return network.Config{
		Codec:             encoding.Marshaler,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.Validator) servertypes.Application {
			err := modifyConsumerGenesis(val)
			if err != nil {
				panic(err)
			}

			err = modifyConsumerGenesisCLITestSetup(val)
			if err != nil {
				panic(err)
			}

			return app.New(
				val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0,
				encoding,
				simapp.EmptyAppOptions{},
				baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
				baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
			)
		},
		GenesisState:  app.ModuleBasics.DefaultGenesis(encoding.Marshaler),
		TimeoutCommit: 2 * time.Second,
		ChainID:       "chain-" + tmrand.NewRand().Str(6),
		// Some changes are introduced to make the tests run as if Duality is a standalone chain.
		// This will only work if NumValidators is set to 1.
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000000006%s", sdk.DefaultBondDenom),
		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: storetypes.PruningOptionNothing,
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

	Index3Price0to1 := sdk.MustNewDecFromStr("0.999700059990001500")
	IndexNeg3Price0to1 := sdk.MustNewDecFromStr("1.000300030001000000")
	Index20Price0to1 := sdk.MustNewDecFromStr("0.998002098460885074")
	dexGenesisState.TickList = []dextypes.Tick{
		dextypes.Tick{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TickIndex: -3,
			TickData: &dextypes.TickDataType{
				Reserve0: []sdk.Int{sdk.ZeroInt(), sdk.NewInt(10), sdk.ZeroInt(), sdk.ZeroInt()},
				Reserve1: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
			},
			LimitOrderTranche0To1: &dextypes.LimitTrancheIndexes{
				PlaceTrancheIndex: 0,
				FillTrancheIndex:  0,
			},
			LimitOrderTranche1To0: &dextypes.LimitTrancheIndexes{
				PlaceTrancheIndex: 0,
				FillTrancheIndex:  0,
			},
			Price0To1: &IndexNeg3Price0to1,
		},
		dextypes.Tick{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TickIndex: 3,
			TickData: &dextypes.TickDataType{
				Reserve0: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
				Reserve1: []sdk.Int{sdk.ZeroInt(), sdk.NewInt(10), sdk.ZeroInt(), sdk.ZeroInt()},
			},
			LimitOrderTranche0To1: &dextypes.LimitTrancheIndexes{
				PlaceTrancheIndex: 0,
				FillTrancheIndex:  0,
			},
			LimitOrderTranche1To0: &dextypes.LimitTrancheIndexes{
				PlaceTrancheIndex: 0,
				FillTrancheIndex:  0,
			},
			Price0To1: &Index3Price0to1,
		},
		dextypes.Tick{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TickIndex: 20,
			TickData: &dextypes.TickDataType{
				Reserve0: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
				Reserve1: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
			},
			LimitOrderTranche0To1: &dextypes.LimitTrancheIndexes{
				PlaceTrancheIndex: 0,
				FillTrancheIndex:  0,
			},
			LimitOrderTranche1To0: &dextypes.LimitTrancheIndexes{
				PlaceTrancheIndex: 0,
				FillTrancheIndex:  0,
			},
			Price0To1: &Index20Price0to1,
		},
	}

	dexGenesisState.TradingPairList = []dextypes.TradingPair{
		dextypes.TradingPair{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			CurrentTick0To1: 3,
			CurrentTick1To0: -3,
			MaxTick:         20,
			MinTick:         -3,
		},
	}

	dexGenesisState.TokensList = []dextypes.Tokens{
		dextypes.Tokens{
			Id:      0,
			Address: "TokenA",
		},
		dextypes.Tokens{
			Id:      1,
			Address: "TokenB",
		},
	}

	dexGenesisState.LimitOrderTrancheUserList = []dextypes.LimitOrderTrancheUser{
		dextypes.LimitOrderTrancheUser{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			Token:           "TokenB",
			TickIndex:       20,
			Count:           0,
			Address:         val.Address.String(),
			SharesOwned:     sdk.NewInt(10),
			SharesWithdrawn: sdk.NewInt(0),
			SharesCancelled: sdk.NewInt(0),
		},
	}

	dexGenesisState.LimitOrderTrancheList = []dextypes.LimitOrderTranche{
		dextypes.LimitOrderTranche{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TokenIn:          "TokenA",
			TickIndex:        -3,
			TrancheIndex:     0,
			ReservesTokenIn:  sdk.NewInt(0),
			ReservesTokenOut: sdk.NewInt(0),
			TotalTokenIn:     sdk.NewInt(0),
			TotalTokenOut:    sdk.NewInt(0),
		},
		dextypes.LimitOrderTranche{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TokenIn:          "TokenB",
			TickIndex:        -3,
			TrancheIndex:     0,
			ReservesTokenIn:  sdk.NewInt(0),
			ReservesTokenOut: sdk.NewInt(0),
			TotalTokenIn:     sdk.NewInt(0),
			TotalTokenOut:    sdk.NewInt(0),
		},
		dextypes.LimitOrderTranche{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TokenIn:          "TokenA",
			TickIndex:        20,
			TrancheIndex:     0,
			ReservesTokenIn:  sdk.NewInt(0),
			ReservesTokenOut: sdk.NewInt(0),
			TotalTokenIn:     sdk.NewInt(0),
			TotalTokenOut:    sdk.NewInt(0),
		},
		dextypes.LimitOrderTranche{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TokenIn:          "TokenB",
			TickIndex:        20,
			TrancheIndex:     0,
			ReservesTokenIn:  sdk.NewInt(10),
			ReservesTokenOut: sdk.NewInt(0),
			TotalTokenIn:     sdk.NewInt(10),
			TotalTokenOut:    sdk.NewInt(0),
		},
		dextypes.LimitOrderTranche{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TokenIn:          "TokenA",
			TickIndex:        3,
			TrancheIndex:     0,
			ReservesTokenIn:  sdk.NewInt(0),
			ReservesTokenOut: sdk.NewInt(0),
			TotalTokenIn:     sdk.NewInt(0),
			TotalTokenOut:    sdk.NewInt(0),
		},
		dextypes.LimitOrderTranche{
			PairId: &dextypes.PairId{
				Token0: "TokenA",
				Token1: "TokenB",
			},
			TokenIn:          "TokenB",
			TickIndex:        3,
			TrancheIndex:     0,
			ReservesTokenIn:  sdk.NewInt(0),
			ReservesTokenOut: sdk.NewInt(0),
			TotalTokenIn:     sdk.NewInt(0),
			TotalTokenOut:    sdk.NewInt(0),
		},
	}
	// // dexGenesisState.
	newRawJSON, _ := json.Marshal(dexGenesisState)
	appState[dextypes.ModuleName] = newRawJSON

	bankData := appState[banktypes.ModuleName]
	var bankGenesisState banktypes.GenesisState
	json.Unmarshal(bankData, &bankGenesisState)

	bankGenesisState.Balances = []banktypes.Balance{
		banktypes.Balance{
			Address: val.Address.String(),
			Coins:   sdk.Coins{sdk.Coin{"DualityPoolShares-TokenA-TokenB-t0-f1", sdk.NewInt(20)}, sdk.Coin{"TokenA", sdk.NewInt(100000000)}, sdk.Coin{"TokenB", sdk.NewInt(100000000)}, sdk.Coin{sdk.DefaultBondDenom, sdk.NewInt(10000)}},
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
