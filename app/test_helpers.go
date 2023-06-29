package app

import (
	"encoding/json"
	"testing"
	"time"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

var DefaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			secp256k1.KeyType,
		},
	},
}

func setup(withGenesis bool, invCheckPeriod uint, chainID string) (*App, GenesisState) {
	db := dbm.NewMemDB()
	encCdc := MakeTestEncodingConfig()
	app := NewApp(log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		EmptyAppOptions{},
		baseapp.SetChainID(chainID),
	)
	if withGenesis {
		return app, NewDefaultGenesisState(encCdc.Marshaler)
	}

	return app, GenesisState{}
}

// Setup initializes a new SimApp. A Nop logger is set in SimApp.
func Setup(t *testing.T, isCheckTx bool) *App {
	// Give bank module minter permissions for testing (ie. KeeperTestHelper.FundAcc)
	// maccPerms[banktypes.ModuleName] = []string{authtypes.Minter}
	app, _ := setup(false, 5, "duality")
	if !isCheckTx {
		genesisState, _ := GenesisStateWithValSet(app)

		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		app.InitChain(
			abci.RequestInitChain{
				ChainId:         "duality",
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func Setup2(isCheckTx bool, chainID string) *App {
	// Give bank module minter permissions for testing (ie. KeeperTestHelper.FundAcc)
	// maccPerms[banktypes.ModuleName] = []string{authtypes.Minter}
	app, _ := setup(false, 5, chainID)
	if !isCheckTx {
		genesisState, valSet := GenesisStateWithValSet(app)

		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}
		valUpdate := make([]abci.ValidatorUpdate, 0, len(valSet.Validators))
		for _, v := range valSet.Validators {
			pubKey, err := cryptocodec.FromTmPubKeyInterface(v.PubKey)
			if err != nil {
				panic("cannot cooerce pk")
			}
			tmPubkey, _ := cryptocodec.ToTmProtoPublicKey(pubKey)
			if err != nil {
				panic(err)
			}
			valUpdate = append(valUpdate, abci.ValidatorUpdate{PubKey: tmPubkey, Power: 100000})
		}

		app.InitChain(
			abci.RequestInitChain{
				ChainId:         chainID,
				Validators:      valUpdate,
				ConsensusParams: DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)

		// ctx := app.NewContext(false, tmproto.Header{})
		// app.ConsumerKeeper.SetSmallestNonOptOutPower(ctx, 1)
		// app.Commit()

		// app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
		// 	ChainID:            chainID,
		// 	Height:             app.LastBlockHeight() + 1,
		// 	AppHash:            app.LastCommitID().Hash,
		// 	ValidatorsHash:     valSet.Hash(),
		// 	NextValidatorsHash: valSet.Hash(),
		// }})
	}

	return app
}

// GenesisStateWithValSet returns a new genesis state with the validator set
func SetupWithGenesisValSet(
	t *testing.T,
	valSet *tmtypes.ValidatorSet,
	genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) *App {
	t.Helper()

	app, genesisState := setup(true, 5, "dualitysdsd")
	genesisState, _ = GenesisStateWithValSet(app)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	app.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	// commit genesis changes
	app.Commit()
	app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
		Height:             app.LastBlockHeight() + 1,
		AppHash:            app.LastCommitID().Hash,
		ValidatorsHash:     valSet.Hash(),
		NextValidatorsHash: valSet.Hash(),
	}})

	return app
}

func GenesisStateWithValSet(app *App) (GenesisState, *tmtypes.ValidatorSet) {
	privVal := mock.NewPV()
	pubKey, _ := privVal.GetPubKey()
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccountWithAddress(senderPrivKey.PubKey().Address().Bytes())
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
	}

	//////////////////////
	balances := []banktypes.Balance{balance}
	genesisState := NewDefaultGenesisState(app.appCodec)
	genAccs := []authtypes.GenesisAccount{acc}
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, _ := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		pkAny, _ := codectypes.NewAnyWithValue(pk)
		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdk.OneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			MinSelfDelegation: sdk.ZeroInt(),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))

	}
	// set validators and delegations
	stakingGenesis := stakingtypes.NewGenesisState(stakingtypes.DefaultParams(), validators, delegations)

	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	for range delegations {
		// add delegated tokens to total supply
		totalSupply = totalSupply.Add(sdk.NewCoin(sdk.DefaultBondDenom, bondAmt))
	}

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, bondAmt)},
	})

	// update total supply
	bankGenesis := banktypes.NewGenesisState(
		banktypes.DefaultGenesisState().Params,
		balances,
		totalSupply,
		[]banktypes.Metadata{},
		[]banktypes.SendEnabled{},
	)
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)
	return genesisState, valSet
}

type EmptyAppOptions struct {
	servertypes.AppOptions
}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(_ string) interface{} {
	return nil
}

func FundAccount(
	bankKeeper bankkeeper.Keeper,
	ctx sdk.Context,
	addr sdk.AccAddress,
	amounts sdk.Coins,
) error {
	if err := bankKeeper.MintCoins(ctx, banktypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, banktypes.ModuleName, addr, amounts)
}
