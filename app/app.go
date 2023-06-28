//nolint:gochecknoglobals
package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	// adminmodulemodule "github.com/cosmos/admin-module/x/adminmodule"
	// adminmodulecli "github.com/cosmos/admin-module/x/adminmodule/client/cli"
	// adminmodulemodulekeeper "github.com/cosmos/admin-module/x/adminmodule/keeper"
	// adminmodulemoduletypes "github.com/cosmos/admin-module/x/adminmodule/types"

	"cosmossdk.io/simapp"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
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
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	ibctestingcore "github.com/cosmos/interchain-security/v3/legacy_ibc_testing/core"
	"github.com/ignite/cli/docs"
	"github.com/ignite/cli/ignite/pkg/openapiconsole"
	"github.com/spf13/cast"
	forwardmiddleware "github.com/strangelove-ventures/packet-forward-middleware/v7/router"
	forwardkeeper "github.com/strangelove-ventures/packet-forward-middleware/v7/router/keeper"
	forwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v7/router/types"
	"github.com/tendermint/spm/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	ccvconsumer "github.com/cosmos/interchain-security/v3/x/ccv/consumer"
	ccvconsumerkeeper "github.com/cosmos/interchain-security/v3/x/ccv/consumer/keeper"
	ccvconsumertypes "github.com/cosmos/interchain-security/v3/x/ccv/consumer/types"

	dexmodule "github.com/duality-labs/duality/x/dex"
	dexmodulekeeper "github.com/duality-labs/duality/x/dex/keeper"
	dexmoduletypes "github.com/duality-labs/duality/x/dex/types"

	icstestintegration "github.com/cosmos/interchain-security/v3/testutil/integration"

	epochsmodule "github.com/duality-labs/duality/x/epochs"
	epochsmodulekeeper "github.com/duality-labs/duality/x/epochs/keeper"
	epochsmoduletypes "github.com/duality-labs/duality/x/epochs/types"
	swapmiddleware "github.com/duality-labs/duality/x/ibcswap"
	swapkeeper "github.com/duality-labs/duality/x/ibcswap/keeper"
	swaptypes "github.com/duality-labs/duality/x/ibcswap/types"

	incentivesmodule "github.com/duality-labs/duality/x/incentives"
	incentivesmodulekeeper "github.com/duality-labs/duality/x/incentives/keeper"
	incentivesmoduletypes "github.com/duality-labs/duality/x/incentives/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport
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
		// adminmodulemodule.NewAppModuleBasic(
		// 	govclient.NewProposalHandler(
		// 		adminmodulecli.NewSubmitParamChangeProposalTxCmd,
		// 	),
		// 	govclient.NewProposalHandler(
		// 		adminmodulecli.NewCmdSubmitUpgradeProposal,
		// 	),
		// 	govclient.NewProposalHandler(
		// 		adminmodulecli.NewCmdSubmitCancelUpgradeProposal,
		// 	),
		// ),
		dexmodule.AppModuleBasic{},
		forwardmiddleware.AppModuleBasic{},
		swapmiddleware.AppModuleBasic{},
		epochsmodule.AppModuleBasic{},
		incentivesmodule.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:                    nil,
		ibctransfertypes.ModuleName:                   {authtypes.Minter, authtypes.Burner},
		dexmoduletypes.ModuleName:                     {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		ccvconsumertypes.ConsumerRedistributeName:     nil,
		ccvconsumertypes.ConsumerToSendToProviderName: nil,
		incentivesmoduletypes.ModuleName:              nil,

		// this line is used by starport scaffolding # stargate/app/maccPerms
	}
)

var (
	_ cosmoscmd.App           = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
	_ simapp.App              = (*App)(nil)
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
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	AuthzKeeper      authzkeeper.Keeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	ConsumerKeeper   ccvconsumerkeeper.Keeper
	// AdminmoduleKeeper adminmodulemodulekeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper         capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper    capabilitykeeper.ScopedKeeper
	ScopedCCVConsumerKeeper capabilitykeeper.ScopedKeeper

	DexKeeper     dexmodulekeeper.Keeper
	SwapKeeper    swapkeeper.Keeper
	ForwardKeeper *forwardkeeper.Keeper

	EpochsKeeper epochsmodulekeeper.Keeper

	IncentivesKeeper incentivesmodulekeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// mm is the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm *module.SimulationManager
}

// New returns a reference to an initialized blockchain app
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig cosmoscmd.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) cosmoscmd.App {
	return NewApp(
		logger,
		db,
		traceStore,
		loadLatest,
		skipUpgradeHeights,
		homePath,
		invCheckPeriod,
		encodingConfig,
		appOpts,
		baseAppOptions...)
}

func NewApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig cosmoscmd.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, authz.ModuleName, banktypes.StoreKey,
		slashingtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		dexmoduletypes.StoreKey, ccvconsumertypes.StoreKey,
		// adminmodulemoduletypes.StoreKey,
		forwardtypes.StoreKey,
		epochsmoduletypes.StoreKey, incentivesmoduletypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedCCVConsumerKeeper := app.CapabilityKeeper.ScopeToModule(ccvconsumertypes.ModuleName)
	app.CapabilityKeeper.Seal()
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authz.ModuleName], appCodec, app.MsgServiceRouter(),
	)

	// Remove the fee-pool from the group of blocked recipient addresses in bank
	// this is required for the consumer chain to be able to send tokens to the provider chain
	// RE: https://github.com/cosmos/interchain-security/v3/blob/main/app/consumer/app.go#L272
	bankBlockedAddrs := app.ModuleAccountAddrs()
	delete(bankBlockedAddrs, authtypes.NewModuleAddress(
		ccvconsumertypes.ConsumerToSendToProviderName).String())
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), bankBlockedAddrs,
	)

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &app.ConsumerKeeper, app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
	)

	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibchost.StoreKey],
		app.GetSubspace(ibchost.ModuleName),
		&app.ConsumerKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.ConsumerKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

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

	app.ConsumerKeeper = *app.ConsumerKeeper.SetHooks(app.SlashingKeeper.Hooks())
	consumerModule := ccvconsumer.NewAppModule(app.ConsumerKeeper)

	// adminRouter := govtypes.NewRouter()
	// adminRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
	// 	AddRoute(proposaltypes.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
	// 	AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
	// 	AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	// app.AdminmoduleKeeper = *adminmodulemodulekeeper.NewKeeper(
	// 	appCodec,
	// 	keys[adminmodulemoduletypes.StoreKey],
	// 	keys[adminmodulemoduletypes.MemStoreKey],
	// 	adminRouter,
	// 	IsProposalWhitelisted,
	// )
	// adminModule := adminmodulemodule.NewAppModule(appCodec, app.AdminmoduleKeeper)

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

	app.EpochsKeeper = *epochsmodulekeeper.NewKeeper(keys[epochsmoduletypes.StoreKey])
	epochsModule := epochsmodule.NewAppModule(app.EpochsKeeper)
	app.EpochsKeeper.SetHooks(
		epochsmoduletypes.NewMultiEpochHooks(
			app.IncentivesKeeper.Hooks(),
		),
	)

	// Set the initialized transfer keeper for forward middleware
	app.ForwardKeeper.SetTransferKeeper(app.TransferKeeper)

	app.IncentivesKeeper = *incentivesmodulekeeper.NewKeeper(
		keys[incentivesmoduletypes.StoreKey],
		app.GetSubspace(incentivesmoduletypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.EpochsKeeper,
		app.DexKeeper,
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

	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	// Create our IBC middleware stack from bottom to top.
	// The stack from top to bottom will look like this:
	// -- channel.OnRecvPacket
	// -- swap.OnRecvPacket
	// -- forward.OnRecvPacket
	// -- transfer.OnRecvPacket
	// see: https://github.com/cosmos/ibc-go/blob/main/docs/middleware/ics29-fee/integration.md#configuring-an-application-stack-with-fee-middleware
	var ibcStack ibcporttypes.IBCModule
	ibcStack = transfer.NewIBCModule(app.TransferKeeper)
	ibcStack = forwardmiddleware.NewIBCMiddleware(
		ibcStack,
		app.ForwardKeeper,
		0, // TODO explore changing default values for retries and timeouts
		forwardkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
		forwardkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
	)
	ibcStack = swapmiddleware.NewIBCMiddleware(ibcStack, app.SwapKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, ibcStack)
	ibcRouter.AddRoute(ccvconsumertypes.ModuleName, consumerModule)
	// this line is used by starport scaffolding # ibc/app/router
	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.ConsumerKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),

		transferModule,
		consumerModule,
		adminModule,
		dexModule,
		forwardModule,
		swapModule,
		epochsModule,
		incentivesModule,
		// this line is used by starport scaffolding # stargate/app/appModule
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
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		crisistypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		ccvconsumertypes.ModuleName,
		adminmodulemoduletypes.ModuleName,
		dexmoduletypes.ModuleName,
		forwardtypes.ModuleName,
		swaptypes.ModuleName,
		incentivesmoduletypes.ModuleName,
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
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		ccvconsumertypes.ModuleName,
		adminmodulemoduletypes.ModuleName,
		forwardtypes.ModuleName,
		swaptypes.ModuleName,
		epochsmoduletypes.ModuleName,
		incentivesmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers

		// NOTE: Because of the gas sensitivity of PurgeExpiredLimit order operations
		// dexmodule must be the last endBlock module to run
		dexmoduletypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		authz.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		ccvconsumertypes.ModuleName,
		adminmodulemoduletypes.ModuleName,
		dexmoduletypes.ModuleName,
		forwardtypes.ModuleName,
		swaptypes.ModuleName,
		epochsmoduletypes.ModuleName,
		incentivesmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()))

	// create the simulation manager and define the order of the modules for deterministic simulations
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.ConsumerKeeper),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
		dexModule,
		forwardModule,
		swapModule,
		// incentivesModule,
		// TODO: Enable lockupModule simulation testing

		// this line is used by starport scaffolding # stargate/app/appModule
	)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCKeeper:      app.IBCKeeper,
			ConsumerKeeper: app.ConsumerKeeper,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %w", err))
	}

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.ScopedCCVConsumerKeeper = scopedCCVConsumerKeeper
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
func (app *App) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
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
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register app's OpenAPI routes.
	apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
	apiSvr.Router.HandleFunc("/", openapiconsole.Handler(Name, "/static/openapi.yml"))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
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
	tkey sdk.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(ccvconsumertypes.ModuleName)
	paramsKeeper.Subspace(dexmoduletypes.ModuleName)
	paramsKeeper.Subspace(forwardtypes.ModuleName).WithKeyTable(forwardtypes.ParamKeyTable())
	paramsKeeper.Subspace(epochsmoduletypes.ModuleName)
	paramsKeeper.Subspace(incentivesmoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// ConsumerApp interface implementations for e2e tests

// GetTxConfig implements the TestingApp interface.
func (app *App) GetTxConfig() client.TxConfig {
	return cosmoscmd.MakeEncodingConfig(ModuleBasics).TxConfig
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

// GetE2eBankKeeper implements the ConsumerApp interface.
func (app *App) GetE2eBankKeeper() icstestintegration.TestBankKeeper {
	return app.BankKeeper
}

// GetE2eAccountKeeper implements the ConsumerApp interface.
func (app *App) GetE2eAccountKeeper() icstestintegration.TestAccountKeeper {
	return app.AccountKeeper
}

// GetE2eSlashingKeeper implements the ConsumerApp interface.
func (app *App) GetE2eSlashingKeeper() icstestintegration.TestSlashingKeeper {
	return app.SlashingKeeper
}

// GetE2eEvidenceKeeper implements the ConsumerApp interface.
func (app *App) GetE2eEvidenceKeeper() icstestintegration.TestEvidenceKeeper {
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
func (k NoopDistributionKeeper) FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error {
	ctx.Logger().Info("FundCommunityPool call was invoked from the no-op distribution keeper")
	return nil
}
