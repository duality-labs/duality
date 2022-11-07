# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: osmosis/mint/v1beta1/genesis.proto, osmosis/mint/v1beta1/mint.proto, osmosis/mint/v1beta1/query.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import Dict, List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


@dataclass(eq=False, repr=False)
class Minter(betterproto.Message):
    """Minter represents the minting state."""

    # epoch_provisions represent rewards for the current epoch.
    epoch_provisions: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class WeightedAddress(betterproto.Message):
    """
    WeightedAddress represents an address with a weight assigned to it. The
    weight is used to determine the proportion of the total minted tokens to be
    minted to the address.
    """

    address: str = betterproto.string_field(1)
    weight: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class DistributionProportions(betterproto.Message):
    """
    DistributionProportions defines the distribution proportions of the minted
    denom. In other words, defines which stakeholders will receive the minted
    denoms and how much.
    """

    # staking defines the proportion of the minted mint_denom that is to be
    # allocated as staking rewards.
    staking: str = betterproto.string_field(1)
    # pool_incentives defines the proportion of the minted mint_denom that is to
    # be allocated as pool incentives.
    pool_incentives: str = betterproto.string_field(2)
    # developer_rewards defines the proportion of the minted mint_denom that is
    # to be allocated to developer rewards address.
    developer_rewards: str = betterproto.string_field(3)
    # community_pool defines the proportion of the minted mint_denom that is to
    # be allocated to the community pool.
    community_pool: str = betterproto.string_field(4)


@dataclass(eq=False, repr=False)
class Params(betterproto.Message):
    """Params holds parameters for the x/mint module."""

    # mint_denom is the denom of the coin to mint.
    mint_denom: str = betterproto.string_field(1)
    # genesis_epoch_provisions epoch provisions from the first epoch.
    genesis_epoch_provisions: str = betterproto.string_field(2)
    # epoch_identifier mint epoch identifier e.g. (day, week).
    epoch_identifier: str = betterproto.string_field(3)
    # reduction_period_in_epochs the number of epochs it takes to reduce the
    # rewards.
    reduction_period_in_epochs: int = betterproto.int64_field(4)
    # reduction_factor is the reduction multiplier to execute at the end of each
    # period set by reduction_period_in_epochs.
    reduction_factor: str = betterproto.string_field(5)
    # distribution_proportions defines the distribution proportions of the minted
    # denom. In other words, defines which stakeholders will receive the minted
    # denoms and how much.
    distribution_proportions: "DistributionProportions" = betterproto.message_field(6)
    # weighted_developer_rewards_receivers is the address to receive developer
    # rewards with weights assignedt to each address. The final amount that each
    # address receives is: epoch_provisions *
    # distribution_proportions.developer_rewards * Address's Weight.
    weighted_developer_rewards_receivers: List[
        "WeightedAddress"
    ] = betterproto.message_field(7)
    # minting_rewards_distribution_start_epoch start epoch to distribute minting
    # rewards
    minting_rewards_distribution_start_epoch: int = betterproto.int64_field(8)


@dataclass(eq=False, repr=False)
class GenesisState(betterproto.Message):
    """GenesisState defines the mint module's genesis state."""

    # minter is an abstraction for holding current rewards information.
    minter: "Minter" = betterproto.message_field(1)
    # params defines all the paramaters of the mint module.
    params: "Params" = betterproto.message_field(2)
    # reduction_started_epoch is the first epoch in which the reduction of mint
    # begins.
    reduction_started_epoch: int = betterproto.int64_field(3)


@dataclass(eq=False, repr=False)
class QueryParamsRequest(betterproto.Message):
    """
    QueryParamsRequest is the request type for the Query/Params RPC method.
    """

    pass


@dataclass(eq=False, repr=False)
class QueryParamsResponse(betterproto.Message):
    """
    QueryParamsResponse is the response type for the Query/Params RPC method.
    """

    # params defines the parameters of the module.
    params: "Params" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class QueryEpochProvisionsRequest(betterproto.Message):
    """
    QueryEpochProvisionsRequest is the request type for the
    Query/EpochProvisions RPC method.
    """

    pass


@dataclass(eq=False, repr=False)
class QueryEpochProvisionsResponse(betterproto.Message):
    """
    QueryEpochProvisionsResponse is the response type for the
    Query/EpochProvisions RPC method.
    """

    # epoch_provisions is the current minting per epoch provisions value.
    epoch_provisions: bytes = betterproto.bytes_field(1)


class QueryStub(betterproto.ServiceStub):
    async def params(self) -> "QueryParamsResponse":

        request = QueryParamsRequest()

        return await self._unary_unary(
            "/osmosis.mint.v1beta1.Query/Params", request, QueryParamsResponse
        )

    async def epoch_provisions(self) -> "QueryEpochProvisionsResponse":

        request = QueryEpochProvisionsRequest()

        return await self._unary_unary(
            "/osmosis.mint.v1beta1.Query/EpochProvisions",
            request,
            QueryEpochProvisionsResponse,
        )


class QueryBase(ServiceBase):
    async def params(self) -> "QueryParamsResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def epoch_provisions(self) -> "QueryEpochProvisionsResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_params(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.params(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_epoch_provisions(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.epoch_provisions(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/osmosis.mint.v1beta1.Query/Params": grpclib.const.Handler(
                self.__rpc_params,
                grpclib.const.Cardinality.UNARY_UNARY,
                QueryParamsRequest,
                QueryParamsResponse,
            ),
            "/osmosis.mint.v1beta1.Query/EpochProvisions": grpclib.const.Handler(
                self.__rpc_epoch_provisions,
                grpclib.const.Cardinality.UNARY_UNARY,
                QueryEpochProvisionsRequest,
                QueryEpochProvisionsResponse,
            ),
        }
