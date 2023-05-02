package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/ibcswap/types"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper defines the swap middleware keeper.
type Keeper struct {
	cdc              codec.BinaryCodec
	msgServiceRouter *baseapp.MsgServiceRouter

	ics4Wrapper porttypes.ICS4Wrapper
	bankKeeper  types.BankKeeper
}

// NewKeeper creates a new swap Keeper instance.
func NewKeeper(
	cdc codec.BinaryCodec,
	msgServiceRouter *baseapp.MsgServiceRouter,
	ics4Wrapper porttypes.ICS4Wrapper,
	bankKeeper types.BankKeeper,
) Keeper {
	return Keeper{
		cdc:              cdc,
		msgServiceRouter: msgServiceRouter,

		ics4Wrapper: ics4Wrapper,
		bankKeeper:  bankKeeper,
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

// SendPacket wraps IBC ChannelKeeper's SendPacket function.
func (k Keeper) SendPacket(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	return k.ics4Wrapper.SendPacket(ctx, chanCap, packet)
}

// WriteAcknowledgement wraps IBC ChannelKeeper's WriteAcknowledgement function.
func (k Keeper) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI, acknowledgement ibcexported.Acknowledgement) error {
	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, acknowledgement)
}

// RefundPacketToken handles the burning or escrow lock up of vouchers when an asset should be refunded.
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

	// decode the receiver address
	receiver, err := sdk.AccAddressFromBech32(data.Receiver)
	if err != nil {
		return err
	}

	// if the sender chain is source that means a voucher was minted on Duality when the ics20 transfer took place
	if transfertypes.SenderChainIsSource(packet.SourcePort, packet.SourceChannel, data.Denom) {
		// transfer coins from user account to transfer module
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, receiver, transfertypes.ModuleName, sdk.NewCoins(token))
		if err != nil {
			return err
		}

		// burn the coins
		err = k.bankKeeper.BurnCoins(ctx, transfertypes.ModuleName, sdk.NewCoins(token))
		if err != nil {
			return err
		}

		return nil
	}

	// transfer coins from user account to escrow address
	escrowAddress := transfertypes.GetEscrowAddress(packet.GetSourcePort(), packet.GetSourceChannel())
	err = k.bankKeeper.SendCoins(ctx, receiver, escrowAddress, sdk.NewCoins(token))
	if err != nil {
		return err
	}

	return nil
}

// SendCoins wraps the BankKeepers SendCoins function so it can be invoked from the middleware.
func (k Keeper) SendCoins(ctx sdk.Context, fromAddr string, toAddr string, amt sdk.Coins) error {
	from, err := sdk.AccAddressFromBech32(fromAddr)
	if err != nil {
		return err
	}

	to, err := sdk.AccAddressFromBech32(toAddr)
	if err != nil {
		return err
	}

	return k.bankKeeper.SendCoins(ctx, from, to, amt)
}

func (k Keeper) GetAppVersion(ctx sdk.Context, portID string, channelID string) (string, bool) {
	return k.ics4Wrapper.GetAppVersion(ctx, portID, channelID)
}
