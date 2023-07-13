//nolint:gochecknoglobals
package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/group"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	forwardmiddleware "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/router"
	forwardkeeper "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/router/keeper"
	forwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/router/types"
	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7"
	ibchookskeeper "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7/keeper"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7/types"
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	tendermint "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	ibctestingcore "github.com/cosmos/interchain-security/v3/legacy_ibc_testing/core"
	ibctesting "github.com/cosmos/interchain-security/v3/legacy_ibc_testing/testing"
	"github.com/spf13/cast"

	ccvconsumer "github.com/cosmos/interchain-security/v3/x/ccv/consumer"
	ccvconsumerkeeper "github.com/cosmos/interchain-security/v3/x/ccv/consumer/keeper"
	ccvconsumertypes "github.com/cosmos/interchain-security/v3/x/ccv/consumer/types"

	dexmodule "github.com/duality-labs/duality/x/dex"
	dexmodulekeeper "github.com/duality-labs/duality/x/dex/keeper"
	dexmoduletypes "github.com/duality-labs/duality/x/dex/types"

	testutil "github.com/cosmos/interchain-security/v3/testutil/integration"

	appparams "github.com/duality-labs/duality/app/params"
	epochsmodule "github.com/duality-labs/duality/x/epochs"
	epochsmodulekeeper "github.com/duality-labs/duality/x/epochs/keeper"
	epochsmoduletypes "github.com/duality-labs/duality/x/epochs/types"
	gmpmiddleware "github.com/duality-labs/duality/x/gmp"
	swapmiddleware "github.com/duality-labs/duality/x/ibcswap"
	swapkeeper "github.com/duality-labs/duality/x/ibcswap/keeper"
	swaptypes "github.com/duality-labs/duality/x/ibcswap/types"
	incentivesmodule "github.com/duality-labs/duality/x/incentives"
	incentivesmodulekeeper "github.com/duality-labs/duality/x/incentives/keeper"
	incentivesmoduletypes "github.com/duality-labs/duality/x/incentives/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	// this line is used by starport scaffolding # stargate/app/moduleImport
	buildermodule "github.com/skip-mev/pob/x/builder"
	builderkeeper "github.com/skip-mev/pob/x/builder/keeper"
	builderrewards "github.com/skip-mev/pob/x/builder/rewards_address_provider"
	buildertypes "github.com/skip-mev/pob/x/builder/types"

	pobabci "github.com/skip-mev/pob/abci"
	pobmempool "github.com/skip-mev/pob/mempool"
)

const (
	AccountAddressPrefix = "cosmos"
	Name                 = "duality"
)

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		ccvconsumer.AppModuleBasic{},
		dexmodule.AppModuleBasic{},
		forwardmiddleware.AppModuleBasic{},
		swapmiddleware.AppModuleBasic{},
		epochsmodule.AppModuleBasic{},
		incentivesmodule.AppModuleBasic{},
		tendermint.AppModuleBasic{},
		genutil.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		buildermodule.AppModuleBasic{},
		wasm.AppModuleBasic{},
		ibchooks.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:  nil,
		ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		dexmoduletypes.ModuleName: {
			authtypes.Minter,
			authtypes.Burner,
			authtypes.Staking,
		},
		ccvconsumertypes.ConsumerRedistributeName:     nil,
		ccvconsumertypes.ConsumerToSendToProviderName: nil,
		incentivesmoduletypes.ModuleName:              nil,
		buildertypes.ModuleName:                       nil,
		wasm.ModuleName:                               {authtypes.Burner},

		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	// This is the address of the admin multisig group, the first group policy configured in x/group.
	// You can rederive this by checking out the `multisig-setup` branch and looking at the README.md.
	AppAuthority = "cosmos1afk9zr2hn2jsac63h4hm60vl9z3e5u69gndzf7c99cqge3vzwjzsfwkgpd"
)

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
	_ ibctesting.TestingApp   = (*App)(nil)
	_ testutil.ConsumerApp    = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	AuthzKeeper           authzkeeper.Keeper
	BankKeeper            bankkeeper.Keeper
	BuildKeeper           builderkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	IBCKeeper             *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper        evidencekeeper.Keeper
	TransferKeeper        ibctransferkeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	ConsumerKeeper        ccvconsumerkeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper
	GroupKeeper           groupkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper         capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper    capabilitykeeper.ScopedKeeper
	ScopedCCVConsumerKeeper capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper        capabilitykeeper.ScopedKeeper

	DexKeeper     dexmodulekeeper.Keeper
	SwapKeeper    swapkeeper.Keeper
	ForwardKeeper *forwardkeeper.Keeper

	EpochsKeeper *epochsmodulekeeper.Keeper

	IncentivesKeeper *incentivesmodulekeeper.Keeper
	WasmKeeper       wasm.Keeper
	IBCHooksKeeper   ibchookskeeper.Keeper

	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// mm is the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator

	// Skip POB's custom checkTx handler
	checkTxHandler pobabci.CheckTx
}

func NewApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	appOpts servertypes.AppOptions,
	encConfig appparams.EncodingConfig,
	wasmOpts []wasm.Option,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encConfig.Marshaler
	cdc := encConfig.Amino
	interfaceRegistry := encConfig.InterfaceRegistry
	txConfig := encConfig.TxConfig

	bApp := baseapp.NewBaseApp(
		Name,
		logger,
		db,
		encConfig.TxConfig.TxDecoder(),
		baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		authz.ModuleName,
		banktypes.StoreKey,
		slashingtypes.StoreKey,
		paramstypes.StoreKey,
		ibcexported.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,
		dexmoduletypes.StoreKey,
		ccvconsumertypes.StoreKey,
		forwardtypes.StoreKey,
		epochsmoduletypes.StoreKey,
		incentivesmoduletypes.StoreKey,
		consensusparamtypes.StoreKey,
		crisistypes.StoreKey,
		group.StoreKey,
		buildertypes.StoreKey,
		wasm.StoreKey,
		ibchookstypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		txConfig:          txConfig,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	// set the BaseApp's parameter store
	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(
		appCodec,
		keys[upgradetypes.StoreKey],
		AppAuthority,
	)
	bApp.SetParamStore(&app.ConsensusParamsKeeper)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedCCVConsumerKeeper := app.CapabilityKeeper.ScopeToModule(ccvconsumertypes.ModuleName)
	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)
	app.CapabilityKeeper.Seal()
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		maccPerms,
		sdk.Bech32PrefixAccAddr,
		AppAuthority,
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.BlockedModuleAccountAddrs(),
		AppAuthority,
	)

	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		AppAuthority,
	)

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibcexported.StoreKey],
		app.GetSubspace(ibcexported.ModuleName),
		&app.ConsumerKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		encConfig.Amino,
		keys[slashingtypes.StoreKey],
		&app.ConsumerKeeper,
		AppAuthority,
	)

	app.ConsumerKeeper = ccvconsumerkeeper.NewKeeper(
		appCodec,
		keys[ccvconsumertypes.StoreKey],
		app.GetSubspace(ccvconsumertypes.ModuleName),
		scopedCCVConsumerKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.IBCKeeper.ConnectionKeeper,
		app.IBCKeeper.ClientKeeper,
		app.SlashingKeeper,
		app.BankKeeper,
		app.AccountKeeper,
		&app.TransferKeeper,
		app.IBCKeeper,
		authtypes.FeeCollectorName,
	)

	app.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		keys[crisistypes.StoreKey],
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		AppAuthority,
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)

	groupConfig := group.DefaultConfig()
	/*
		Example of setting group params:
		groupConfig.MaxMetadataLen = 1000
	*/
	app.GroupKeeper = groupkeeper.NewKeeper(
		keys[group.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
		groupConfig,
	)

	// ... other modules keepers

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.ConsumerKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	app.ConsumerKeeper = *app.ConsumerKeeper.SetHooks(app.SlashingKeeper.Hooks())
	consumerModule := ccvconsumer.NewAppModule(
		app.ConsumerKeeper,
		app.GetSubspace(ccvconsumertypes.ModuleName),
	)

	app.DexKeeper = *dexmodulekeeper.NewKeeper(
		appCodec,
		keys[dexmoduletypes.StoreKey],
		keys[dexmoduletypes.MemStoreKey],
		app.GetSubspace(dexmoduletypes.ModuleName),

		app.BankKeeper,
	)
	dexModule := dexmodule.NewAppModule(appCodec, app.DexKeeper, app.AccountKeeper, app.BankKeeper)

	// Create swap middleware keeper
	app.SwapKeeper = swapkeeper.NewKeeper(
		appCodec,
		app.MsgServiceRouter(),
		app.IBCKeeper.ChannelKeeper,
		app.BankKeeper,
	)
	swapModule := swapmiddleware.NewAppModule(app.SwapKeeper)

	// Create packet forward middleware keeper
	app.ForwardKeeper = forwardkeeper.NewKeeper(
		appCodec,
		keys[forwardtypes.StoreKey],
		app.GetSubspace(forwardtypes.ModuleName),
		app.TransferKeeper, // This is zero value because transfer keeper is not initialized yet
		app.IBCKeeper.ChannelKeeper,
		&NoopDistributionKeeper{}, // Use a no-op implementation of the Distribution Keeper to avoid NPE
		app.BankKeeper,
		&app.IBCKeeper.ChannelKeeper,
	)
	forwardModule := forwardmiddleware.NewAppModule(app.ForwardKeeper)

	app.EpochsKeeper = epochsmodulekeeper.NewKeeper(keys[epochsmoduletypes.StoreKey])

	app.IncentivesKeeper = incentivesmodulekeeper.NewKeeper(
		keys[incentivesmoduletypes.StoreKey],
		app.GetSubspace(incentivesmoduletypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.EpochsKeeper,
		app.DexKeeper,
		AppAuthority,
	)

	app.IncentivesKeeper.SetHooks(
		incentivesmoduletypes.NewMultiIncentiveHooks(
		// insert Incentives hooks receivers here
		),
	)

	incentivesModule := incentivesmodule.NewAppModule(
		app.IncentivesKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.EpochsKeeper,
	)

	app.EpochsKeeper.SetHooks(epochsmoduletypes.NewMultiEpochHooks(
		app.IncentivesKeeper.Hooks(),
	))

	// NB: This must be initialized AFTER app.EpochsKeeper.SetHooks() because otherwise
	// we dereference an out-of-date EpochsKeeper.
	epochsModule := epochsmodule.NewAppModule(*app.EpochsKeeper)

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// wasmOpts = append(
	// 	wasmbinding.RegisterCustomPlugins(
	// 		&app.InterchainTxsKeeper,
	// 		&app.InterchainQueriesKeeper,
	// 		app.TransferKeeper,
	// 		&app.AdminmoduleKeeper,
	// 		app.FeeBurnerKeeper,
	// 		app.FeeKeeper,
	// 		&app.BankKeeper,
	// 		app.TokenFactoryKeeper,
	// 		&app.CronKeeper,
	// 	),
	// 	wasmOpts...)

	app.WasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		nil,
		nil,
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		strings.Join(AllCapabilities(), ","),
		AppAuthority,
		wasmOpts...,
	)

	// 'ibc-hooks' module - depends on
	// 1. 'auth'
	// 2. 'bank'
	// 3. 'distr'
	app.IBCHooksKeeper = ibchookskeeper.NewKeeper(
		app.keys[ibchookstypes.StoreKey],
	)
	ics20WasmHooks := ibchooks.NewWasmHooks(
		&app.IBCHooksKeeper,
		&app.WasmKeeper,
		AccountAddressPrefix,
	) // The contract keeper needs to be set later
	hooksICS4Wrapper := ibchooks.NewICS4Middleware(
		app.IBCKeeper.ChannelKeeper,
		ics20WasmHooks,
	)

	// Create Transfer Module
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		hooksICS4Wrapper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)
	app.ForwardKeeper.SetTransferKeeper(app.TransferKeeper)

	// Set the initialized transfer keeper for forward middleware
	app.ForwardKeeper.SetTransferKeeper(app.TransferKeeper)

	moduleAddress := app.AccountKeeper.GetModuleAddress(buildertypes.ModuleName)
	rewardsAddressProvider := builderrewards.NewFixedAddressRewardsAddressProvider(moduleAddress)

	app.BuildKeeper = builderkeeper.NewKeeperWithRewardsAddressProvider(
		appCodec,
		app.keys[buildertypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		rewardsAddressProvider,
		AppAuthority,
	)

	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	// Create our IBC middleware stack from bottom to top.
	// The stack order is this:
	// -- channel.OnRecvPacket
	// -- gmp.OnRecvPacket
	// -- swap.OnRecvPacket
	// -- forward.OnRecvPacket
	// -- ibchooks.OnRecvPacket
	// -- transfer.OnRecvPacket
	// -- ibchookswrapper.OnRecvPacket
	//
	// The confusing thing is that everything flows down to transfer,
	// but then once it gets back up to swap, swap will call down to forward
	// again. Then forward has to know through the context not to call down to
	// transfer a second time. This is in my opinion a bad separation of concerns.
	var transferStack ibcporttypes.IBCModule
	transferStack = transfer.NewIBCModule(app.TransferKeeper)
	transferStack = ibchooks.NewIBCMiddleware(transferStack, &hooksICS4Wrapper)
	transferStack = forwardmiddleware.NewIBCMiddleware(
		transferStack,
		app.ForwardKeeper,
		0, // TODO explore changing default values for retries and timeouts
		forwardkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
		forwardkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
	)
	transferStack = swapmiddleware.NewIBCMiddleware(transferStack, app.SwapKeeper)
	transferStack = gmpmiddleware.NewIBCMiddleware(transferStack)

	wasmStack := wasm.NewIBCHandler(
		app.WasmKeeper,
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack)
	ibcRouter.AddRoute(ccvconsumertypes.ModuleName, consumerModule)
	ibcRouter.AddRoute(wasm.ModuleName, wasmStack)
	// this line is used by starport scaffolding # ibc/app/router
	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper,
			app.ConsumerKeeper,
			app.BaseApp.DeliverTx,
			encConfig.TxConfig,
		),
		auth.NewAppModule(
			appCodec,
			app.AccountKeeper,
			authsims.RandomGenesisAccounts,
			app.GetSubspace(authtypes.ModuleName),
		),
		authzmodule.NewAppModule(
			appCodec,
			app.AuthzKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(
			appCodec,
			app.BankKeeper,
			app.AccountKeeper,
			app.GetSubspace(banktypes.ModuleName),
		),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(
			appCodec,
			app.AccountKeeper,
			app.BankKeeper,
			app.FeeGrantKeeper,
			app.interfaceRegistry,
		),
		slashing.NewAppModule(
			appCodec,
			app.SlashingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.ConsumerKeeper,
			app.GetSubspace(slashingtypes.ModuleName),
		),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),

		transferModule,
		consumerModule,
		dexModule,
		forwardModule,
		swapModule,
		epochsModule,
		incentivesModule,
		ibchooks.NewAppModule(app.AccountKeeper),
		// this line is used by starport scaffolding # stargate/app/appModule

		// always be last to make sure that it checks for all invariants and not only part of them
		crisis.NewAppModule(
			app.CrisisKeeper,
			skipGenesisInvariants,
			app.GetSubspace(crisistypes.ModuleName),
		),
		groupmodule.NewAppModule(
			appCodec,
			app.GroupKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		buildermodule.NewAppModule(appCodec, app.BuildKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		epochsmoduletypes.ModuleName,
		capabilitytypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		vestingtypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		ccvconsumertypes.ModuleName,
		dexmoduletypes.ModuleName,
		forwardtypes.ModuleName,
		swaptypes.ModuleName,
		incentivesmoduletypes.ModuleName,
		group.ModuleName,
		buildertypes.ModuleName,
		ibchookstypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		vestingtypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		ibcexported.ModuleName,
		genutiltypes.ModuleName,
		ibctransfertypes.ModuleName,
		ccvconsumertypes.ModuleName,
		forwardtypes.ModuleName,
		swaptypes.ModuleName,
		epochsmoduletypes.ModuleName,
		incentivesmoduletypes.ModuleName,
		group.ModuleName,
		buildertypes.ModuleName,
		ibchookstypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers

		// NOTE: Because of the gas sensitivity of PurgeExpiredLimit order operations
		// dexmodule must be the last endBlock module to run
		dexmoduletypes.ModuleName,
	)

	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	// NOTE: The genutils module must occur after consumer so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		ibcexported.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		authz.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		ccvconsumertypes.ModuleName,
		genutiltypes.ModuleName,
		dexmoduletypes.ModuleName,
		forwardtypes.ModuleName,
		swaptypes.ModuleName,
		epochsmoduletypes.ModuleName,
		incentivesmoduletypes.ModuleName,
		group.ModuleName,
		buildertypes.ModuleName,
		ibchookstypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	app.mm.RegisterInvariants(app.CrisisKeeper)
	app.configurator = module.NewConfigurator(
		app.appCodec,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
	)
	app.mm.RegisterServices(
		module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()),
	)

	// create the simulation manager and define the order of the modules for deterministic simulations
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(
			app.appCodec,
			app.AccountKeeper,
			authsims.RandomGenesisAccounts,
			app.GetSubspace(authtypes.ModuleName),
		),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	factory := pobmempool.NewDefaultAuctionFactory(app.txConfig.TxDecoder())
	mempool := pobmempool.NewAuctionMempool(
		app.txConfig.TxDecoder(),
		app.txConfig.TxEncoder(),
		0,
		factory,
	)
	app.SetMempool(mempool)

	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: encConfig.TxConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCKeeper:         app.IBCKeeper,
			ConsumerKeeper:    app.ConsumerKeeper,
			WasmConfig:        &wasmConfig,
			TXCounterStoreKey: keys[wasm.StoreKey],
			TxEncoder:         app.txConfig.TxEncoder(),
			BuilderKeeper:     app.BuildKeeper,
			Mempool:           mempool,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %w", err))
	}

	proposalHandlers := pobabci.NewProposalHandler(
		mempool,
		app.Logger(),
		anteHandler,
		app.txConfig.TxEncoder(),
		app.txConfig.TxDecoder(),
	)
	app.SetPrepareProposal(proposalHandlers.PrepareProposalHandler())
	app.SetProcessProposal(proposalHandlers.ProcessProposalHandler())

	// Set the custom CheckTx handler on BaseApp.
	checkTxHandler := pobabci.NewCheckTxHandler(
		app.BaseApp,
		app.txConfig.TxDecoder(),
		mempool,
		anteHandler,
		app.ChainID(),
	)
	app.SetCheckTx(checkTxHandler.CheckTx())

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	// must be before Loading version
	// requires the snapshot store to be created and registered as a BaseAppOption
	// see cmd/wasmd/root.go: 206 - 214 approx
	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), &app.WasmKeeper),
		)
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extension: %s", err))
		}
	}

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.ScopedCCVConsumerKeeper = scopedCCVConsumerKeeper
	app.ScopedWasmKeeper = scopedWasmKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// GetBaseApp returns the base app of the application
func (app App) GetBaseApp() *baseapp.BaseApp { return app.BaseApp }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// CheckTx will check the transaction with the provided checkTxHandler. We override the default
// handler so that we can verify bid transactions before they are inserted into the mempool.
// With the POB CheckTx, we can verify the bid transaction and all of the bundled transactions
// before inserting the bid transaction into the mempool.
func (app *App) CheckTx(req abci.RequestCheckTx) abci.ResponseCheckTx {
	return app.checkTxHandler(req)
}

// SetCheckTx sets the checkTxHandler for the app.
func (app *App) SetCheckTx(handler pobabci.CheckTx) {
	app.checkTxHandler = handler
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedModuleAccountAddrs returns all the app's blocked module account
// addresses.
func (app *App) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()
	delete(modAccAddrs, AppAuthority)
	delete(
		modAccAddrs,
		authtypes.NewModuleAddress(ccvconsumertypes.ConsumerToSendToProviderName).String(),
	)

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns an app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns an InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, _ config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(
		app.BaseApp.GRPCQueryRouter(),
		clientCtx,
		app.BaseApp.Simulate,
		app.interfaceRegistry,
	)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

// RegisterNodeService implements the Application.RegisterNodeService method.
func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}

	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(
	appCodec codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	key,
	tkey storetypes.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibcexported.ModuleName)
	paramsKeeper.Subspace(ccvconsumertypes.ModuleName)
	paramsKeeper.Subspace(dexmoduletypes.ModuleName)
	paramsKeeper.Subspace(forwardtypes.ModuleName).WithKeyTable(forwardtypes.ParamKeyTable())
	paramsKeeper.Subspace(epochsmoduletypes.ModuleName)
	paramsKeeper.Subspace(incentivesmoduletypes.ModuleName)
	paramsKeeper.Subspace(buildertypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// ConsumerApp interface implementations for interchain-security/v3 tests

// GetTxConfig implements the TestingApp interface.
func (app *App) GetTxConfig() client.TxConfig {
	return app.txConfig
}

// GetIBCKeeper implements the TestingApp interface.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// // GetStakingKeeper implements the TestingApp interface.
//
//	func (app *App) GetStakingKeeper() ibcclienttypes.StakingKeeper {
//		return app.ConsumerKeeper
//	}
//
// GetStakingKeeper implements the TestingApp interface.
func (app *App) GetStakingKeeper() ibctestingcore.StakingKeeper {
	return app.ConsumerKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *App) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetConsumerKeeper implements the ConsumerApp interface.
func (app *App) GetConsumerKeeper() ccvconsumerkeeper.Keeper {
	return app.ConsumerKeeper
}

// GetTestBankKeeper implements the ConsumerApp interface.
func (app *App) GetTestBankKeeper() testutil.TestBankKeeper {
	return app.BankKeeper
}

// GetTestAccountKeeper implements the ConsumerApp interface.
func (app *App) GetTestAccountKeeper() testutil.TestAccountKeeper {
	return app.AccountKeeper
}

// GetTestSlashingKeeper implements the ConsumerApp interface.
func (app *App) GetTestSlashingKeeper() testutil.TestSlashingKeeper {
	return app.SlashingKeeper
}

// GetTestEvidenceKeeper implements the ConsumerApp interface.
func (app *App) GetTestEvidenceKeeper() testutil.TestEvidenceKeeper {
	return app.EvidenceKeeper
}

// NoopDistributionKeeper is a replacement for the distribution keeper that results in a no-op when its methods are called.
// This is needed because the forward middleware expects the distribution keeper for funding the community pool if
// a tax is set greater than 0. We only use this to avoid a possible nil pointer exception.
type NoopDistributionKeeper struct {
	// TODO explore other ways of funding community pool for consumer chains and remove this NoopDistributionKeeper
}

// FundCommunityPool is a no-op function that returns nil and logs that it was invoked.
// Realistically this should never be called.
func (k NoopDistributionKeeper) FundCommunityPool(
	ctx sdk.Context,
	amount sdk.Coins,
	sender sdk.AccAddress,
) error {
	ctx.Logger().Info("FundCommunityPool call was invoked from the no-op distribution keeper")
	return nil
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range GetMaccPerms() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	// allow the following addresses to receive funds
	delete(modAccAddrs, AppAuthority)

	return modAccAddrs
}

// ModuleManager returns the app ModuleManager
func (app *App) ModuleManager() *module.Manager {
	return app.mm
}

// ChainID gets chainID from private fields of BaseApp
// Should be removed once SDK 0.50.x will be adopted
func (app *App) ChainID() string {
	field := reflect.ValueOf(app.BaseApp).Elem().FieldByName("chainID")
	return field.String()
}
