# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: cosmos/base/abci/v1beta1/abci.proto
# plugin: python-betterproto
import warnings
from dataclasses import dataclass
from typing import List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase


@dataclass(eq=False, repr=False)
class TxResponse(betterproto.Message):
    """
    TxResponse defines a structure containing relevant tx data and metadata.
    The tags are stringified and the log is JSON decoded.
    """

    # The block height
    height: int = betterproto.int64_field(1)
    # The transaction hash.
    txhash: str = betterproto.string_field(2)
    # Namespace for the Code
    codespace: str = betterproto.string_field(3)
    # Response code.
    code: int = betterproto.uint32_field(4)
    # Result bytes, if any.
    data: str = betterproto.string_field(5)
    # The output of the application's logger (raw string). May be non-
    # deterministic.
    raw_log: str = betterproto.string_field(6)
    # The output of the application's logger (typed). May be non-deterministic.
    logs: List["AbciMessageLog"] = betterproto.message_field(7)
    # Additional information. May be non-deterministic.
    info: str = betterproto.string_field(8)
    # Amount of gas requested for transaction.
    gas_wanted: int = betterproto.int64_field(9)
    # Amount of gas consumed by transaction.
    gas_used: int = betterproto.int64_field(10)
    # The request transaction bytes.
    tx: "betterproto_lib_google_protobuf.Any" = betterproto.message_field(11)
    # Time of the previous block. For heights > 1, it's the weighted median of
    # the timestamps of the valid votes in the block.LastCommit. For height == 1,
    # it's genesis time.
    timestamp: str = betterproto.string_field(12)
    # Events defines all the events emitted by processing a transaction. Note,
    # these events include those emitted by processing all the messages and those
    # emitted from the ante handler. Whereas Logs contains the events, with
    # additional metadata, emitted only by processing the messages. Since:
    # cosmos-sdk 0.42.11, 0.44.5, 0.45
    events: List["____tendermint_abci__.Event"] = betterproto.message_field(13)


@dataclass(eq=False, repr=False)
class AbciMessageLog(betterproto.Message):
    """
    ABCIMessageLog defines a structure containing an indexed tx ABCI message
    log.
    """

    msg_index: int = betterproto.uint32_field(1)
    log: str = betterproto.string_field(2)
    # Events contains a slice of Event objects that were emitted during some
    # execution.
    events: List["StringEvent"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class StringEvent(betterproto.Message):
    """
    StringEvent defines en Event object wrapper where all the attributes
    contain key/value pairs that are strings instead of raw bytes.
    """

    type: str = betterproto.string_field(1)
    attributes: List["Attribute"] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class Attribute(betterproto.Message):
    """
    Attribute defines an attribute wrapper where the key and value are strings
    instead of raw bytes.
    """

    key: str = betterproto.string_field(1)
    value: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class GasInfo(betterproto.Message):
    """GasInfo defines tx execution gas context."""

    # GasWanted is the maximum units of work we allow this tx to perform.
    gas_wanted: int = betterproto.uint64_field(1)
    # GasUsed is the amount of gas actually consumed.
    gas_used: int = betterproto.uint64_field(2)


@dataclass(eq=False, repr=False)
class Result(betterproto.Message):
    """Result is the union of ResponseFormat and ResponseCheckTx."""

    # Data is any data returned from message or handler execution. It MUST be
    # length prefixed in order to separate data from multiple message executions.
    # Deprecated. This field is still populated, but prefer msg_response instead
    # because it also contains the Msg response typeURL.
    data: bytes = betterproto.bytes_field(1)
    # Log contains the log information from message or handler execution.
    log: str = betterproto.string_field(2)
    # Events contains a slice of Event objects that were emitted during message
    # or handler execution.
    events: List["____tendermint_abci__.Event"] = betterproto.message_field(3)
    # msg_responses contains the Msg handler responses type packed in Anys.
    # Since: cosmos-sdk 0.46
    msg_responses: List[
        "betterproto_lib_google_protobuf.Any"
    ] = betterproto.message_field(4)

    def __post_init__(self) -> None:
        super().__post_init__()
        if self.data:
            warnings.warn("Result.data is deprecated", DeprecationWarning)


@dataclass(eq=False, repr=False)
class SimulationResponse(betterproto.Message):
    """
    SimulationResponse defines the response generated when a transaction is
    successfully simulated.
    """

    gas_info: "GasInfo" = betterproto.message_field(1)
    result: "Result" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class MsgData(betterproto.Message):
    """
    MsgData defines the data returned in a Result object during message
    execution.
    """

    msg_type: str = betterproto.string_field(1)
    data: bytes = betterproto.bytes_field(2)

    def __post_init__(self) -> None:
        warnings.warn("MsgData is deprecated", DeprecationWarning)
        super().__post_init__()


@dataclass(eq=False, repr=False)
class TxMsgData(betterproto.Message):
    """
    TxMsgData defines a list of MsgData. A transaction will have a MsgData
    object for each message.
    """

    # data field is deprecated and not populated.
    data: List["MsgData"] = betterproto.message_field(1)
    # msg_responses contains the Msg handler responses packed into Anys. Since:
    # cosmos-sdk 0.46
    msg_responses: List[
        "betterproto_lib_google_protobuf.Any"
    ] = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()
        if self.data:
            warnings.warn("TxMsgData.data is deprecated", DeprecationWarning)


@dataclass(eq=False, repr=False)
class SearchTxsResult(betterproto.Message):
    """SearchTxsResult defines a structure for querying txs pageable"""

    # Count of all txs
    total_count: int = betterproto.uint64_field(1)
    # Count of txs in current page
    count: int = betterproto.uint64_field(2)
    # Index of current page, start from 1
    page_number: int = betterproto.uint64_field(3)
    # Count of total pages
    page_total: int = betterproto.uint64_field(4)
    # Max count txs per page
    limit: int = betterproto.uint64_field(5)
    # List of txs in current page
    txs: List["TxResponse"] = betterproto.message_field(6)


from .....tendermint import abci as ____tendermint_abci__
import betterproto.lib.google.protobuf as betterproto_lib_google_protobuf
