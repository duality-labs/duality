package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/cwhooks/types"
)

func (k Keeper) EmitTrigger(ctx sdk.Context, triggerKey, triggerValue string) {
	// JCP TODO: validate key/val since we are storing them simply we need to make sure they follow regexp
	tStore := prefix.NewStore(
		ctx.TransientStore(k.tKey),
		types.KeyPrefix(types.TriggerKeyPrefix),
	)

	key := types.TriggerKey(triggerKey, triggerValue)

	tStore.Set(key, []byte{})
}

func (k Keeper) GetAllTriggers(ctx sdk.Context) (list []types.Trigger) {
	tStore := prefix.NewStore(
		ctx.TransientStore(k.tKey),
		types.KeyPrefix(types.TriggerKeyPrefix),
	)

	iterator := sdk.KVStorePrefixIterator(tStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val, err := types.ParseTriggerFromBytes(iterator.Key())
		if err != nil {
			panic("error while parsing trigger")
		}
		list = append(list, val)
	}

	return
}

func (k Keeper) GetAllTriggeredHooks(ctx sdk.Context) (list []types.Hook) {
	// If we do anything fancier here (ie. value-less hooks remember to de-dupe
	triggers := k.GetAllTriggers(ctx)

	for _, trigger := range triggers {
		hooks := k.GetAllHooksForKeyValue(ctx, trigger.Key, trigger.Value)
		list = append(list, hooks...)
	}

	return list
}

func (k Keeper) InvokeAllTriggeredHooks(ctx sdk.Context) {
	// JCP TODO: think about logic to prevent timing out block
	hooks := k.GetAllTriggeredHooks(ctx)
	for _, hook := range hooks {
		resp, err := k.ExecuteHook(ctx, hook)
		if err != nil {
			ctx.EventManager().EmitEvent(CreateHookFailedEvent(hook.Id, err))
		}
		ctx.EventManager().EmitEvent(CreateHookSuccessEvent(hook.Id, resp))
	}
}

// JCP TODO: Move these to dedicated types.events switc to typed event

func CreateHookFailedEvent(
	hookID uint64,
	err error,
) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "cwhooks"),
		sdk.NewAttribute(sdk.AttributeKeyAction, "HookFailed"),
		sdk.NewAttribute("TriggerID", strconv.FormatUint(hookID, 10)),
		sdk.NewAttribute("Error", err.Error()),
	}

	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}

func CreateHookSuccessEvent(
	hookID uint64,
	resp []byte,
) sdk.Event {
	attrs := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, "cwhooks"),
		sdk.NewAttribute(sdk.AttributeKeyAction, "HookFailed"),
		sdk.NewAttribute("TriggerID", strconv.FormatUint(hookID, 10)),
		sdk.NewAttribute("RESP", string(resp)),
	}

	return sdk.NewEvent(sdk.EventTypeMessage, attrs...)
}
