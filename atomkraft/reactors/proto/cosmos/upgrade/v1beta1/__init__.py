# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: cosmos/upgrade/v1beta1/query.proto, cosmos/upgrade/v1beta1/tx.proto, cosmos/upgrade/v1beta1/upgrade.proto
# plugin: python-betterproto
import warnings
from dataclasses import dataclass
from datetime import datetime
from typing import Dict, List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


@dataclass(eq=False, repr=False)
class Plan(betterproto.Message):
    """
    Plan specifies information about a planned upgrade and when it should
    occur.
    """

    # Sets the name for the upgrade. This name will be used by the upgraded
    # version of the software to apply any special "on-upgrade" commands during
    # the first BeginBlock method after the upgrade is applied. It is also used
    # to detect whether a software version can handle a given upgrade. If no
    # upgrade handler with this name has been set in the software, it will be
    # assumed that the software is out-of-date when the upgrade Time or Height is
    # reached and the software will exit.
    name: str = betterproto.string_field(1)
    # Deprecated: Time based upgrades have been deprecated. Time based upgrade
    # logic has been removed from the SDK. If this field is not empty, an error
    # will be thrown.
    time: datetime = betterproto.message_field(2)
    # The height at which the upgrade must be performed. Only used if Time is not
    # set.
    height: int = betterproto.int64_field(3)
    # Any application specific upgrade info to be included on-chain such as a git
    # commit that validators could automatically upgrade to
    info: str = betterproto.string_field(4)
    # Deprecated: UpgradedClientState field has been deprecated. IBC upgrade
    # logic has been moved to the IBC module in the sub module 02-client. If this
    # field is not empty, an error will be thrown.
    upgraded_client_state: "betterproto_lib_google_protobuf.Any" = (
        betterproto.message_field(5)
    )

    def __post_init__(self) -> None:
        super().__post_init__()
        if self.time:
            warnings.warn("Plan.time is deprecated", DeprecationWarning)
        if self.upgraded_client_state:
            warnings.warn(
                "Plan.upgraded_client_state is deprecated", DeprecationWarning
            )


@dataclass(eq=False, repr=False)
class SoftwareUpgradeProposal(betterproto.Message):
    """
    SoftwareUpgradeProposal is a gov Content type for initiating a software
    upgrade. Deprecated: This legacy proposal is deprecated in favor of Msg-
    based gov proposals, see MsgSoftwareUpgrade.
    """

    title: str = betterproto.string_field(1)
    description: str = betterproto.string_field(2)
    plan: "Plan" = betterproto.message_field(3)

    def __post_init__(self) -> None:
        warnings.warn("SoftwareUpgradeProposal is deprecated", DeprecationWarning)
        super().__post_init__()


@dataclass(eq=False, repr=False)
class CancelSoftwareUpgradeProposal(betterproto.Message):
    """
    CancelSoftwareUpgradeProposal is a gov Content type for cancelling a
    software upgrade. Deprecated: This legacy proposal is deprecated in favor
    of Msg-based gov proposals, see MsgCancelUpgrade.
    """

    title: str = betterproto.string_field(1)
    description: str = betterproto.string_field(2)

    def __post_init__(self) -> None:
        warnings.warn("CancelSoftwareUpgradeProposal is deprecated", DeprecationWarning)
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ModuleVersion(betterproto.Message):
    """
    ModuleVersion specifies a module and its consensus version. Since: cosmos-
    sdk 0.43
    """

    # name of the app module
    name: str = betterproto.string_field(1)
    # consensus version of the app module
    version: int = betterproto.uint64_field(2)


@dataclass(eq=False, repr=False)
class MsgSoftwareUpgrade(betterproto.Message):
    """
    MsgSoftwareUpgrade is the Msg/SoftwareUpgrade request type. Since: cosmos-
    sdk 0.46
    """

    # authority is the address of the governance account.
    authority: str = betterproto.string_field(1)
    # plan is the upgrade plan.
    plan: "Plan" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class MsgSoftwareUpgradeResponse(betterproto.Message):
    """
    MsgSoftwareUpgradeResponse is the Msg/SoftwareUpgrade response type. Since:
    cosmos-sdk 0.46
    """

    pass


@dataclass(eq=False, repr=False)
class MsgCancelUpgrade(betterproto.Message):
    """
    MsgCancelUpgrade is the Msg/CancelUpgrade request type. Since: cosmos-sdk
    0.46
    """

    # authority is the address of the governance account.
    authority: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class MsgCancelUpgradeResponse(betterproto.Message):
    """
    MsgCancelUpgradeResponse is the Msg/CancelUpgrade response type. Since:
    cosmos-sdk 0.46
    """

    pass


@dataclass(eq=False, repr=False)
class QueryCurrentPlanRequest(betterproto.Message):
    """
    QueryCurrentPlanRequest is the request type for the Query/CurrentPlan RPC
    method.
    """

    pass


@dataclass(eq=False, repr=False)
class QueryCurrentPlanResponse(betterproto.Message):
    """
    QueryCurrentPlanResponse is the response type for the Query/CurrentPlan RPC
    method.
    """

    # plan is the current upgrade plan.
    plan: "Plan" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class QueryAppliedPlanRequest(betterproto.Message):
    """
    QueryCurrentPlanRequest is the request type for the Query/AppliedPlan RPC
    method.
    """

    # name is the name of the applied plan to query for.
    name: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class QueryAppliedPlanResponse(betterproto.Message):
    """
    QueryAppliedPlanResponse is the response type for the Query/AppliedPlan RPC
    method.
    """

    # height is the block height at which the plan was applied.
    height: int = betterproto.int64_field(1)


@dataclass(eq=False, repr=False)
class QueryUpgradedConsensusStateRequest(betterproto.Message):
    """
    QueryUpgradedConsensusStateRequest is the request type for the
    Query/UpgradedConsensusState RPC method.
    """

    # last height of the current chain must be sent in request as this is the
    # height under which next consensus state is stored
    last_height: int = betterproto.int64_field(1)

    def __post_init__(self) -> None:
        warnings.warn(
            "QueryUpgradedConsensusStateRequest is deprecated", DeprecationWarning
        )
        super().__post_init__()


@dataclass(eq=False, repr=False)
class QueryUpgradedConsensusStateResponse(betterproto.Message):
    """
    QueryUpgradedConsensusStateResponse is the response type for the
    Query/UpgradedConsensusState RPC method.
    """

    # Since: cosmos-sdk 0.43
    upgraded_consensus_state: bytes = betterproto.bytes_field(2)

    def __post_init__(self) -> None:
        warnings.warn(
            "QueryUpgradedConsensusStateResponse is deprecated", DeprecationWarning
        )
        super().__post_init__()


@dataclass(eq=False, repr=False)
class QueryModuleVersionsRequest(betterproto.Message):
    """
    QueryModuleVersionsRequest is the request type for the Query/ModuleVersions
    RPC method. Since: cosmos-sdk 0.43
    """

    # module_name is a field to query a specific module consensus version from
    # state. Leaving this empty will fetch the full list of module versions from
    # state
    module_name: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class QueryModuleVersionsResponse(betterproto.Message):
    """
    QueryModuleVersionsResponse is the response type for the
    Query/ModuleVersions RPC method. Since: cosmos-sdk 0.43
    """

    # module_versions is a list of module names with their consensus versions.
    module_versions: List["ModuleVersion"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class QueryAuthorityRequest(betterproto.Message):
    """
    QueryAuthorityRequest is the request type for Query/Authority Since:
    cosmos-sdk 0.46
    """

    pass


@dataclass(eq=False, repr=False)
class QueryAuthorityResponse(betterproto.Message):
    """
    QueryAuthorityResponse is the response type for Query/Authority Since:
    cosmos-sdk 0.46
    """

    address: str = betterproto.string_field(1)


class MsgStub(betterproto.ServiceStub):
    async def software_upgrade(
        self, *, authority: str = "", plan: "Plan" = None
    ) -> "MsgSoftwareUpgradeResponse":

        request = MsgSoftwareUpgrade()
        request.authority = authority
        if plan is not None:
            request.plan = plan

        return await self._unary_unary(
            "/cosmos.upgrade.v1beta1.Msg/SoftwareUpgrade",
            request,
            MsgSoftwareUpgradeResponse,
        )

    async def cancel_upgrade(
        self, *, authority: str = ""
    ) -> "MsgCancelUpgradeResponse":

        request = MsgCancelUpgrade()
        request.authority = authority

        return await self._unary_unary(
            "/cosmos.upgrade.v1beta1.Msg/CancelUpgrade",
            request,
            MsgCancelUpgradeResponse,
        )


class QueryStub(betterproto.ServiceStub):
    async def current_plan(self) -> "QueryCurrentPlanResponse":

        request = QueryCurrentPlanRequest()

        return await self._unary_unary(
            "/cosmos.upgrade.v1beta1.Query/CurrentPlan",
            request,
            QueryCurrentPlanResponse,
        )

    async def applied_plan(self, *, name: str = "") -> "QueryAppliedPlanResponse":

        request = QueryAppliedPlanRequest()
        request.name = name

        return await self._unary_unary(
            "/cosmos.upgrade.v1beta1.Query/AppliedPlan",
            request,
            QueryAppliedPlanResponse,
        )

    async def upgraded_consensus_state(
        self, *, last_height: int = 0
    ) -> "QueryUpgradedConsensusStateResponse":

        request = QueryUpgradedConsensusStateRequest()
        request.last_height = last_height

        return await self._unary_unary(
            "/cosmos.upgrade.v1beta1.Query/UpgradedConsensusState",
            request,
            QueryUpgradedConsensusStateResponse,
        )

    async def module_versions(
        self, *, module_name: str = ""
    ) -> "QueryModuleVersionsResponse":

        request = QueryModuleVersionsRequest()
        request.module_name = module_name

        return await self._unary_unary(
            "/cosmos.upgrade.v1beta1.Query/ModuleVersions",
            request,
            QueryModuleVersionsResponse,
        )

    async def authority(self) -> "QueryAuthorityResponse":

        request = QueryAuthorityRequest()

        return await self._unary_unary(
            "/cosmos.upgrade.v1beta1.Query/Authority", request, QueryAuthorityResponse
        )


class MsgBase(ServiceBase):
    async def software_upgrade(
        self, authority: str, plan: "Plan"
    ) -> "MsgSoftwareUpgradeResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def cancel_upgrade(self, authority: str) -> "MsgCancelUpgradeResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_software_upgrade(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "authority": request.authority,
            "plan": request.plan,
        }

        response = await self.software_upgrade(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_cancel_upgrade(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "authority": request.authority,
        }

        response = await self.cancel_upgrade(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/cosmos.upgrade.v1beta1.Msg/SoftwareUpgrade": grpclib.const.Handler(
                self.__rpc_software_upgrade,
                grpclib.const.Cardinality.UNARY_UNARY,
                MsgSoftwareUpgrade,
                MsgSoftwareUpgradeResponse,
            ),
            "/cosmos.upgrade.v1beta1.Msg/CancelUpgrade": grpclib.const.Handler(
                self.__rpc_cancel_upgrade,
                grpclib.const.Cardinality.UNARY_UNARY,
                MsgCancelUpgrade,
                MsgCancelUpgradeResponse,
            ),
        }


class QueryBase(ServiceBase):
    async def current_plan(self) -> "QueryCurrentPlanResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def applied_plan(self, name: str) -> "QueryAppliedPlanResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def upgraded_consensus_state(
        self, last_height: int
    ) -> "QueryUpgradedConsensusStateResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def module_versions(self, module_name: str) -> "QueryModuleVersionsResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def authority(self) -> "QueryAuthorityResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_current_plan(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.current_plan(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_applied_plan(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "name": request.name,
        }

        response = await self.applied_plan(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_upgraded_consensus_state(
        self, stream: grpclib.server.Stream
    ) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "last_height": request.last_height,
        }

        response = await self.upgraded_consensus_state(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_module_versions(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "module_name": request.module_name,
        }

        response = await self.module_versions(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_authority(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.authority(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/cosmos.upgrade.v1beta1.Query/CurrentPlan": grpclib.const.Handler(
                self.__rpc_current_plan,
                grpclib.const.Cardinality.UNARY_UNARY,
                QueryCurrentPlanRequest,
                QueryCurrentPlanResponse,
            ),
            "/cosmos.upgrade.v1beta1.Query/AppliedPlan": grpclib.const.Handler(
                self.__rpc_applied_plan,
                grpclib.const.Cardinality.UNARY_UNARY,
                QueryAppliedPlanRequest,
                QueryAppliedPlanResponse,
            ),
            "/cosmos.upgrade.v1beta1.Query/UpgradedConsensusState": grpclib.const.Handler(
                self.__rpc_upgraded_consensus_state,
                grpclib.const.Cardinality.UNARY_UNARY,
                QueryUpgradedConsensusStateRequest,
                QueryUpgradedConsensusStateResponse,
            ),
            "/cosmos.upgrade.v1beta1.Query/ModuleVersions": grpclib.const.Handler(
                self.__rpc_module_versions,
                grpclib.const.Cardinality.UNARY_UNARY,
                QueryModuleVersionsRequest,
                QueryModuleVersionsResponse,
            ),
            "/cosmos.upgrade.v1beta1.Query/Authority": grpclib.const.Handler(
                self.__rpc_authority,
                grpclib.const.Cardinality.UNARY_UNARY,
                QueryAuthorityRequest,
                QueryAuthorityResponse,
            ),
        }


import betterproto.lib.google.protobuf as betterproto_lib_google_protobuf
