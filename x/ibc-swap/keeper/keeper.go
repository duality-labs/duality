package keeper

import (
	"fmt"

	dextypes "github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/NicholasDotSol/duality/x/ibc-swap/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/log"
)

// Middleware must implement types.ChannelKeeper and types.PortKeeper, expected interfaces
// so that it can wrap IBC channel and port logic for underlying application.
var (
	_ types.ChannelKeeper = Keeper{}
	_ types.PortKeeper    = Keeper{}
)

// Keeper defines the swap middleware keeper.
type Keeper struct {
	cdc              codec.BinaryCodec
	msgServiceRouter *baseapp.MsgServiceRouter

	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	bankKeeper    types.BankKeeper
}

// NewKeeper creates a new swap Keeper instance.
func NewKeeper(
	cdc codec.BinaryCodec,
	msgServiceRouter *baseapp.MsgServiceRouter,
	ics4Wrapper porttypes.ICS4Wrapper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	return Keeper{
		cdc:              cdc,
		msgServiceRouter: msgServiceRouter,

		ics4Wrapper:   ics4Wrapper,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		bankKeeper:    bankKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+host.ModuleName+"-"+types.ModuleName)
}

// Swap calls into the base app's msg service router so that the appropriate handler is called when sending the swap msg.
func (k Keeper) Swap(ctx sdk.Context, msg *dextypes.MsgSwap) (*dextypes.MsgSwapResponse, error) {
	swapHandler := k.msgServiceRouter.Handler(msg)
	if swapHandler == nil {
		return nil, sdkerrors.Wrap(types.ErrMsgHandlerInvalid, fmt.Sprintf("could not find the handler for %T", msg))
	}

	res, err := swapHandler(ctx, msg)
	if err != nil {
		return nil, err
	}

	msgSwapRes := &dextypes.MsgSwapResponse{}
	if err := proto.Unmarshal(res.Data, msgSwapRes); err != nil {
		return nil, err
	}

	return msgSwapRes, nil
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function.
func (k Keeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	return k.portKeeper.BindPort(ctx, portID)
}

// GetChannel wraps IBC ChannelKeeper's GetChannel function.
func (k Keeper) GetChannel(ctx sdk.Context, portID, channelID string) (channeltypes.Channel, bool) {
	return k.channelKeeper.GetChannel(ctx, portID, channelID)
}

// GetPacketCommitment wraps IBC ChannelKeeper's GetPacketCommitment function.
func (k Keeper) GetPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) []byte {
	return k.channelKeeper.GetPacketCommitment(ctx, portID, channelID, sequence)
}

// GetNextSequenceSend wraps IBC ChannelKeeper's GetNextSequenceSend function.
func (k Keeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return k.channelKeeper.GetNextSequenceSend(ctx, portID, channelID)
}

// SendPacket wraps IBC ChannelKeeper's SendPacket function
func (k Keeper) SendPacket(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	return k.ics4Wrapper.SendPacket(ctx, chanCap, packet)
}

// WriteAcknowledgement wraps IBC ChannelKeeper's WriteAcknowledgement function.
// ICS29 WriteAcknowledgement is used for asynchronous acknowledgements.
func (k Keeper) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI, acknowledgement ibcexported.Acknowledgement) error {
	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, acknowledgement)
}

// RefundPacketToken handles the burning/minting of vouchers when an asset should be refunded, this comes straight from ics20.
// This is only used in the case where we call into the transfer modules OnRecvPacket callback but then the swap fails.
func (k Keeper) RefundPacketToken(ctx sdk.Context, packet channeltypes.Packet, data transfertypes.FungibleTokenPacketData) error {
	// parse the denomination from the full denom path
	trace := transfertypes.ParseDenomTrace(data.Denom)

	// parse the transfer amount
	transferAmount, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		return sdkerrors.Wrapf(transfertypes.ErrInvalidAmount, "unable to parse transfer amount (%s) into math.Int", data.Amount)
	}
	token := sdk.NewCoin(trace.IBCDenom(), transferAmount)

	// decode the sender address
	sender, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return err
	}

	if transfertypes.SenderChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), data.Denom) {
		// unescrow tokens back to sender
		escrowAddress := transfertypes.GetEscrowAddress(packet.GetSourcePort(), packet.GetSourceChannel())
		if err := k.bankKeeper.SendCoins(ctx, escrowAddress, sender, sdk.NewCoins(token)); err != nil {
			// NOTE: this error is only expected to occur given an unexpected bug or a malicious
			// counterparty module. The bug may occur in bank or any part of the code that allows
			// the escrow address to be drained. A malicious counterparty module could drain the
			// escrow address by allowing more tokens to be sent back then were escrowed.
			return sdkerrors.Wrap(err, "unable to unescrow tokens, this may be caused by a malicious counterparty module or a bug: please open an issue on counterparty module")
		}

		return nil
	}

	// mint vouchers back to sender
	if err := k.bankKeeper.MintCoins(
		ctx, transfertypes.ModuleName, sdk.NewCoins(token),
	); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, transfertypes.ModuleName, sender, sdk.NewCoins(token)); err != nil {
		panic(fmt.Sprintf("unable to send coins from module to account despite previously minting coins to module account: %v", err))
	}

	return nil
}
