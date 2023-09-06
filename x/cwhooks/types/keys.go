package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "cwhooks"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cwhooks"

	// TransientStoreKey defines the transient store key
	TransientStoreKey = "trans_cwhooks"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	HookKeyPrefix    = "Hook/value/"
	HookCountKey     = "Hook/count/"
	HookIDKeyPrefix  = "HookID/ref/"
	TriggerKeyPrefix = "Trigger/value/"
)

func HookKey(triggerKey, triggerValue string, id uint64) []byte {
	var key []byte

	key = append(key, []byte(triggerKey)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(triggerValue)...)
	key = append(key, []byte("/")...)

	key = append(key, sdk.Uint64ToBigEndian(id)...)
	key = append(key, []byte("/")...)

	return key
}

func HookKVPrefix(triggerKey, triggerValue string) []byte {
	key := KeyPrefix(HookKeyPrefix)

	key = append(key, []byte(triggerKey)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(triggerValue)...)
	key = append(key, []byte("/")...)

	return key
}

func TriggerKey(triggerKey, triggerValue string) []byte {
	var key []byte

	key = append(key, []byte(triggerKey)...)
	key = append(key, []byte(TriggerSep)...)

	key = append(key, []byte(triggerValue)...)

	return key
}
