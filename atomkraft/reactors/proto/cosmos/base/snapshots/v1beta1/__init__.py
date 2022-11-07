# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: cosmos/base/snapshots/v1beta1/snapshot.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase


@dataclass(eq=False, repr=False)
class Snapshot(betterproto.Message):
    """Snapshot contains Tendermint state sync snapshot info."""

    height: int = betterproto.uint64_field(1)
    format: int = betterproto.uint32_field(2)
    chunks: int = betterproto.uint32_field(3)
    hash: bytes = betterproto.bytes_field(4)
    metadata: "Metadata" = betterproto.message_field(5)


@dataclass(eq=False, repr=False)
class Metadata(betterproto.Message):
    """Metadata contains SDK-specific snapshot metadata."""

    chunk_hashes: List[bytes] = betterproto.bytes_field(1)


@dataclass(eq=False, repr=False)
class SnapshotItem(betterproto.Message):
    """SnapshotItem is an item contained in a rootmulti.Store snapshot."""

    store: "SnapshotStoreItem" = betterproto.message_field(1, group="item")
    iavl: "SnapshotIavlItem" = betterproto.message_field(2, group="item")
    extension: "SnapshotExtensionMeta" = betterproto.message_field(3, group="item")
    extension_payload: "SnapshotExtensionPayload" = betterproto.message_field(
        4, group="item"
    )
    kv: "SnapshotKvItem" = betterproto.message_field(5, group="item")
    schema: "SnapshotSchema" = betterproto.message_field(6, group="item")


@dataclass(eq=False, repr=False)
class SnapshotStoreItem(betterproto.Message):
    """SnapshotStoreItem contains metadata about a snapshotted store."""

    name: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class SnapshotIavlItem(betterproto.Message):
    """SnapshotIAVLItem is an exported IAVL node."""

    key: bytes = betterproto.bytes_field(1)
    value: bytes = betterproto.bytes_field(2)
    # version is block height
    version: int = betterproto.int64_field(3)
    # height is depth of the tree.
    height: int = betterproto.int32_field(4)


@dataclass(eq=False, repr=False)
class SnapshotExtensionMeta(betterproto.Message):
    """
    SnapshotExtensionMeta contains metadata about an external snapshotter.
    """

    name: str = betterproto.string_field(1)
    format: int = betterproto.uint32_field(2)


@dataclass(eq=False, repr=False)
class SnapshotExtensionPayload(betterproto.Message):
    """
    SnapshotExtensionPayload contains payloads of an external snapshotter.
    """

    payload: bytes = betterproto.bytes_field(1)


@dataclass(eq=False, repr=False)
class SnapshotKvItem(betterproto.Message):
    """SnapshotKVItem is an exported Key/Value Pair"""

    key: bytes = betterproto.bytes_field(1)
    value: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class SnapshotSchema(betterproto.Message):
    """SnapshotSchema is an exported schema of smt store"""

    keys: List[bytes] = betterproto.bytes_field(1)
