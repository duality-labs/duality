package ibcswap

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/duality-labs/duality/x/ibcswap/keeper"
	"github.com/duality-labs/duality/x/ibcswap/types"
	forwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v3/router/types"
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks for the swap middleware given the
// swap keeper and the underlying application.
type IBCMiddleware struct {
	app    porttypes.IBCModule
	keeper keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application.
func NewIBCMiddleware(app porttypes.IBCModule, k keeper.Keeper) IBCMiddleware {
	return IBCMiddleware{
		app:    app,
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	return im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
}

// OnChanOpenTry implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID, channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (version string, err error) {
	return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, counterpartyVersion)
}

// OnChanOpenAck implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID, channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCModule interface.
func (im IBCMiddleware) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCModule interface.
func (im IBCMiddleware) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket checks the memo field on this packet and if the metadata inside's root key indicates this packet
// should be handled by the swap middleware it attempts to perform a swap. If the swap is successful
// the underlying application's OnRecvPacket callback is invoked.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return channeltypes.NewErrorAcknowledgement(err.Error())
	}

	m := &types.PacketMetadata{}
	err := json.Unmarshal([]byte(data.Memo), m)
	if err != nil || m.Swap == nil {
		// Not a packet that should be handled by the swap middleware
		return im.app.OnRecvPacket(ctx, packet, relayer)
	}

	metadata := m.Swap
	if err := metadata.Validate(); err != nil {
		return channeltypes.NewErrorAcknowledgement(err.Error())
	}

	// Check if a middleware before this has already processed this packet and passed the processed key via context
	var processed bool
	goCtx := ctx.Context()
	p := goCtx.Value(types.ProcessedKey{})

	if p != nil {
		if pb, ok := p.(bool); ok {
			processed = pb
		}
	}

	// Compose our context with values that will be used to pass through to the forward middleware
	ctxWithForwardFlags := context.WithValue(ctx.Context(), forwardtypes.ProcessedKey{}, true)
	ctxWithForwardFlags = context.WithValue(ctxWithForwardFlags, forwardtypes.NonrefundableKey{}, true)
	ctxWithForwardFlags = context.WithValue(ctxWithForwardFlags, forwardtypes.DisableDenomCompositionKey{}, true)
	wrappedSdkCtx := ctx.WithContext(ctxWithForwardFlags)

	// If this packet has been handled by another middleware in the stack there is no need to call into the
	// underlying app, otherwise the transfer module's OnRecvPacket callback could be invoked more than once
	// which would mint/burn vouchers more than once.
	var ack ibcexported.Acknowledgement
	if !processed {
		ack = im.app.OnRecvPacket(wrappedSdkCtx, packet, relayer)
		if ack == nil || !ack.Success() {
			return ack
		}
	}

	// Attempt to perform a swap since this packets memo included swap metadata.
	res, err := im.keeper.Swap(ctx, metadata.MsgSwap)
	if err != nil {
		swapErr := sdkerrors.Wrap(types.ErrSwapFailed, err.Error())

		// If a refund to the counterparty chain is not being issued we return a successful ack containing the error
		// This ensures that a refund is not issued on the counterparty chain
		// see https://github.com/cosmos/ibc-go/blob/3ecc7dd3aef5790ec5d906936a297b34adf1ee41/modules/apps/transfer/keeper/relay.go#L320
		if metadata.NonRefundable {
			return channeltypes.NewResultAcknowledgement([]byte(swapErr.Error()))
		}

		// We need to get the denom for this token on this chain before issuing a refund on this side
		denomOnThisChain := getDenomForThisChain(
			packet.DestinationPort, packet.DestinationChannel,
			packet.SourcePort, packet.SourceChannel,
			data.Denom,
		)

		data.Denom = denomOnThisChain

		// We called into the transfer modules OnRecvPacket callback to mint or unescrow the funds on this side
		// so if the swap fails we need to explicitly refund to handle the bookkeeping properly
		err = im.keeper.RefundPacketToken(ctx, packet, data)
		if err != nil {
			// If the refund fails on this side we want to make sure that the refund does not happen on the counterparty,
			// so we return a successful ack containing the error
			return channeltypes.NewResultAcknowledgement([]byte(err.Error()))
		}

		return channeltypes.NewErrorAcknowledgement(swapErr.Error())
	}

	// If there is no next field set in the metadata return ack
	if metadata.Next == "" {
		return ack
	}

	// Set the new packet data to include the token denom and amount that was received from the swap.
	data.Denom = res.CoinOut.Denom
	data.Amount = res.CoinOut.Amount.String()

	// Swaps can come into Duality over IBC where the swap creator could be a module/contract swapping on behalf of the user.
	// Then the swap's receiver field will be the user controlled address where funds are deposited afterwards.
	// Before passing to the forward middleware we need to override the packet receiver field to now point to the
	// user controlled address that will be initiating the forward since this is where the funds are after the swap.
	data.Receiver = m.Swap.Receiver

	// We need to reset the packets memo field so that the root key in the metadata is the
	// next field from the current metadata.
	data.Memo = m.Swap.Next

	dataBz, err := transfertypes.ModuleCdc.MarshalJSON(&data)
	if err != nil {
		return ack
	}

	packet.Data = dataBz

	// The forward middleware should return a nil ack if the forward is initiated properly.
	// If not an error occurred, and we return the original ack.
	newAck := im.app.OnRecvPacket(wrappedSdkCtx, packet, relayer)
	if newAck != nil {
		return ack
	}

	return nil
}

// OnAcknowledgementPacket implements the IBCModule interface.
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCMiddleware) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}

// SendPacket implements the ICS4 Wrapper interface.
func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
) error {
	return im.keeper.SendPacket(ctx, chanCap, packet)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface.
func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	return im.keeper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

// getDenomForThisChain composes a new token denom by either unwinding or prefixing the specified token denom appropriately.
// This is necessary because the token denom in the packet data is from the perspective of the counterparty chain.
func getDenomForThisChain(port, channel, counterpartyPort, counterpartyChannel, denom string) string {
	counterpartyPrefix := transfertypes.GetDenomPrefix(counterpartyPort, counterpartyChannel)
	if strings.HasPrefix(denom, counterpartyPrefix) {
		// unwind denom
		unwoundDenom := denom[len(counterpartyPrefix):]
		denomTrace := transfertypes.ParseDenomTrace(unwoundDenom)
		if denomTrace.Path == "" {
			// denom is now unwound back to native denom
			return unwoundDenom
		}
		// denom is still IBC denom
		return denomTrace.IBCDenom()
	}
	// append port and channel from this chain to denom
	prefixedDenom := transfertypes.GetDenomPrefix(port, channel) + denom
	return transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
}
