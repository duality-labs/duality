/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../dex/params";
import { TickMap } from "../dex/tick_map";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { PairMap } from "../dex/pair_map";
import { Tokens } from "../dex/tokens";
import { TokenMap } from "../dex/token_map";
import { Shares } from "../dex/shares";
import { FeeList } from "../dex/fee_list";
import { EdgeRow } from "../dex/edge_row";
import { AdjanceyMatrix } from "../dex/adjancey_matrix";
import { LimitOrderPoolUserShareMap } from "../dex/limit_order_pool_user_share_map";
import { LimitOrderPoolUserSharesWithdrawn } from "../dex/limit_order_pool_user_shares_withdrawn";
import { LimitOrderPoolTotalSharesMap } from "../dex/limit_order_pool_total_shares_map";
import { LimitOrderPoolReserveMap } from "../dex/limit_order_pool_reserve_map";
import { LimitOrderPoolFillMap } from "../dex/limit_order_pool_fill_map";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetTickMapRequest {
  tickIndex: number;
  pairId: string;
}

export interface QueryGetTickMapResponse {
  tickMap: TickMap | undefined;
}

export interface QueryAllTickMapRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTickMapResponse {
  tickMap: TickMap[];
  pagination: PageResponse | undefined;
}

export interface QueryGetPairMapRequest {
  pairId: string;
}

export interface QueryGetPairMapResponse {
  pairMap: PairMap | undefined;
}

export interface QueryAllPairMapRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPairMapResponse {
  pairMap: PairMap[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTokensRequest {
  id: number;
}

export interface QueryGetTokensResponse {
  Tokens: Tokens | undefined;
}

export interface QueryAllTokensRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTokensResponse {
  Tokens: Tokens[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTokenMapRequest {
  address: string;
}

export interface QueryGetTokenMapResponse {
  tokenMap: TokenMap | undefined;
}

export interface QueryAllTokenMapRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTokenMapResponse {
  tokenMap: TokenMap[];
  pagination: PageResponse | undefined;
}

export interface QueryGetSharesRequest {
  address: string;
  pairId: string;
  tickIndex: number;
  fee: number;
}

export interface QueryGetSharesResponse {
  shares: Shares | undefined;
}

export interface QueryAllSharesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSharesResponse {
  shares: Shares[];
  pagination: PageResponse | undefined;
}

export interface QueryGetFeeListRequest {
  id: number;
}

export interface QueryGetFeeListResponse {
  FeeList: FeeList | undefined;
}

export interface QueryAllFeeListRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllFeeListResponse {
  FeeList: FeeList[];
  pagination: PageResponse | undefined;
}

export interface QueryGetEdgeRowRequest {
  id: number;
}

export interface QueryGetEdgeRowResponse {
  EdgeRow: EdgeRow | undefined;
}

export interface QueryAllEdgeRowRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEdgeRowResponse {
  EdgeRow: EdgeRow[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAdjanceyMatrixRequest {
  id: number;
}

export interface QueryGetAdjanceyMatrixResponse {
  AdjanceyMatrix: AdjanceyMatrix | undefined;
}

export interface QueryAllAdjanceyMatrixRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAdjanceyMatrixResponse {
  AdjanceyMatrix: AdjanceyMatrix[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolUserShareMapRequest {
  count: number;
  address: string;
  pairId: string;
  tickIndex: number;
  token: string;
}

export interface QueryGetLimitOrderPoolUserShareMapResponse {
  limitOrderPoolUserShareMap: LimitOrderPoolUserShareMap | undefined;
}

export interface QueryAllLimitOrderPoolUserShareMapRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolUserShareMapResponse {
  limitOrderPoolUserShareMap: LimitOrderPoolUserShareMap[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolUserSharesWithdrawnRequest {
  count: number;
  address: string;
  pairId: string;
  tickIndex: number;
  token: string;
}

export interface QueryGetLimitOrderPoolUserSharesWithdrawnResponse {
  limitOrderPoolUserSharesWithdrawn: LimitOrderPoolUserSharesWithdrawn | undefined;
}

export interface QueryAllLimitOrderPoolUserSharesWithdrawnRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolUserSharesWithdrawnResponse {
  limitOrderPoolUserSharesWithdrawn: LimitOrderPoolUserSharesWithdrawn[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolTotalSharesMapRequest {
  count: number;
  pairId: string;
  tickIndex: number;
  token: string;
}

export interface QueryGetLimitOrderPoolTotalSharesMapResponse {
  limitOrderPoolTotalSharesMap: LimitOrderPoolTotalSharesMap | undefined;
}

export interface QueryAllLimitOrderPoolTotalSharesMapRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolTotalSharesMapResponse {
  limitOrderPoolTotalSharesMap: LimitOrderPoolTotalSharesMap[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolReserveMapRequest {
  count: number;
  pairId: string;
  tickIndex: number;
  token: string;
}

export interface QueryGetLimitOrderPoolReserveMapResponse {
  limitOrderPoolReserveMap: LimitOrderPoolReserveMap | undefined;
}

export interface QueryAllLimitOrderPoolReserveMapRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolReserveMapResponse {
  limitOrderPoolReserveMap: LimitOrderPoolReserveMap[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolFillMapRequest {
  count: number;
  pairId: string;
  tickIndex: number;
  token: string;
}

export interface QueryGetLimitOrderPoolFillMapResponse {
  limitOrderPoolFillMap: LimitOrderPoolFillMap | undefined;
}

export interface QueryAllLimitOrderPoolFillMapRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolFillMapResponse {
  limitOrderPoolFillMap: LimitOrderPoolFillMap[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetTickMapRequest: object = { tickIndex: 0, pairId: "" };

export const QueryGetTickMapRequest = {
  encode(
    message: QueryGetTickMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tickIndex !== 0) {
      writer.uint32(8).int64(message.tickIndex);
    }
    if (message.pairId !== "") {
      writer.uint32(18).string(message.pairId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTickMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTickMapRequest } as QueryGetTickMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.pairId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTickMapRequest {
    const message = { ...baseQueryGetTickMapRequest } as QueryGetTickMapRequest;
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    return message;
  },

  toJSON(message: QueryGetTickMapRequest): unknown {
    const obj: any = {};
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTickMapRequest>
  ): QueryGetTickMapRequest {
    const message = { ...baseQueryGetTickMapRequest } as QueryGetTickMapRequest;
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    return message;
  },
};

const baseQueryGetTickMapResponse: object = {};

export const QueryGetTickMapResponse = {
  encode(
    message: QueryGetTickMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tickMap !== undefined) {
      TickMap.encode(message.tickMap, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTickMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTickMapResponse,
    } as QueryGetTickMapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tickMap = TickMap.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTickMapResponse {
    const message = {
      ...baseQueryGetTickMapResponse,
    } as QueryGetTickMapResponse;
    if (object.tickMap !== undefined && object.tickMap !== null) {
      message.tickMap = TickMap.fromJSON(object.tickMap);
    } else {
      message.tickMap = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTickMapResponse): unknown {
    const obj: any = {};
    message.tickMap !== undefined &&
      (obj.tickMap = message.tickMap
        ? TickMap.toJSON(message.tickMap)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTickMapResponse>
  ): QueryGetTickMapResponse {
    const message = {
      ...baseQueryGetTickMapResponse,
    } as QueryGetTickMapResponse;
    if (object.tickMap !== undefined && object.tickMap !== null) {
      message.tickMap = TickMap.fromPartial(object.tickMap);
    } else {
      message.tickMap = undefined;
    }
    return message;
  },
};

const baseQueryAllTickMapRequest: object = {};

export const QueryAllTickMapRequest = {
  encode(
    message: QueryAllTickMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTickMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTickMapRequest } as QueryAllTickMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTickMapRequest {
    const message = { ...baseQueryAllTickMapRequest } as QueryAllTickMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTickMapRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTickMapRequest>
  ): QueryAllTickMapRequest {
    const message = { ...baseQueryAllTickMapRequest } as QueryAllTickMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTickMapResponse: object = {};

export const QueryAllTickMapResponse = {
  encode(
    message: QueryAllTickMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.tickMap) {
      TickMap.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTickMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTickMapResponse,
    } as QueryAllTickMapResponse;
    message.tickMap = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tickMap.push(TickMap.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTickMapResponse {
    const message = {
      ...baseQueryAllTickMapResponse,
    } as QueryAllTickMapResponse;
    message.tickMap = [];
    if (object.tickMap !== undefined && object.tickMap !== null) {
      for (const e of object.tickMap) {
        message.tickMap.push(TickMap.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTickMapResponse): unknown {
    const obj: any = {};
    if (message.tickMap) {
      obj.tickMap = message.tickMap.map((e) =>
        e ? TickMap.toJSON(e) : undefined
      );
    } else {
      obj.tickMap = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTickMapResponse>
  ): QueryAllTickMapResponse {
    const message = {
      ...baseQueryAllTickMapResponse,
    } as QueryAllTickMapResponse;
    message.tickMap = [];
    if (object.tickMap !== undefined && object.tickMap !== null) {
      for (const e of object.tickMap) {
        message.tickMap.push(TickMap.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetPairMapRequest: object = { pairId: "" };

export const QueryGetPairMapRequest = {
  encode(
    message: QueryGetPairMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPairMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPairMapRequest } as QueryGetPairMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPairMapRequest {
    const message = { ...baseQueryGetPairMapRequest } as QueryGetPairMapRequest;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    return message;
  },

  toJSON(message: QueryGetPairMapRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPairMapRequest>
  ): QueryGetPairMapRequest {
    const message = { ...baseQueryGetPairMapRequest } as QueryGetPairMapRequest;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    return message;
  },
};

const baseQueryGetPairMapResponse: object = {};

export const QueryGetPairMapResponse = {
  encode(
    message: QueryGetPairMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairMap !== undefined) {
      PairMap.encode(message.pairMap, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPairMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPairMapResponse,
    } as QueryGetPairMapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairMap = PairMap.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPairMapResponse {
    const message = {
      ...baseQueryGetPairMapResponse,
    } as QueryGetPairMapResponse;
    if (object.pairMap !== undefined && object.pairMap !== null) {
      message.pairMap = PairMap.fromJSON(object.pairMap);
    } else {
      message.pairMap = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPairMapResponse): unknown {
    const obj: any = {};
    message.pairMap !== undefined &&
      (obj.pairMap = message.pairMap
        ? PairMap.toJSON(message.pairMap)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPairMapResponse>
  ): QueryGetPairMapResponse {
    const message = {
      ...baseQueryGetPairMapResponse,
    } as QueryGetPairMapResponse;
    if (object.pairMap !== undefined && object.pairMap !== null) {
      message.pairMap = PairMap.fromPartial(object.pairMap);
    } else {
      message.pairMap = undefined;
    }
    return message;
  },
};

const baseQueryAllPairMapRequest: object = {};

export const QueryAllPairMapRequest = {
  encode(
    message: QueryAllPairMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPairMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPairMapRequest } as QueryAllPairMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPairMapRequest {
    const message = { ...baseQueryAllPairMapRequest } as QueryAllPairMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPairMapRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPairMapRequest>
  ): QueryAllPairMapRequest {
    const message = { ...baseQueryAllPairMapRequest } as QueryAllPairMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPairMapResponse: object = {};

export const QueryAllPairMapResponse = {
  encode(
    message: QueryAllPairMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.pairMap) {
      PairMap.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPairMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPairMapResponse,
    } as QueryAllPairMapResponse;
    message.pairMap = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairMap.push(PairMap.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPairMapResponse {
    const message = {
      ...baseQueryAllPairMapResponse,
    } as QueryAllPairMapResponse;
    message.pairMap = [];
    if (object.pairMap !== undefined && object.pairMap !== null) {
      for (const e of object.pairMap) {
        message.pairMap.push(PairMap.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPairMapResponse): unknown {
    const obj: any = {};
    if (message.pairMap) {
      obj.pairMap = message.pairMap.map((e) =>
        e ? PairMap.toJSON(e) : undefined
      );
    } else {
      obj.pairMap = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPairMapResponse>
  ): QueryAllPairMapResponse {
    const message = {
      ...baseQueryAllPairMapResponse,
    } as QueryAllPairMapResponse;
    message.pairMap = [];
    if (object.pairMap !== undefined && object.pairMap !== null) {
      for (const e of object.pairMap) {
        message.pairMap.push(PairMap.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetTokensRequest: object = { id: 0 };

export const QueryGetTokensRequest = {
  encode(
    message: QueryGetTokensRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTokensRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTokensRequest } as QueryGetTokensRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokensRequest {
    const message = { ...baseQueryGetTokensRequest } as QueryGetTokensRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetTokensRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokensRequest>
  ): QueryGetTokensRequest {
    const message = { ...baseQueryGetTokensRequest } as QueryGetTokensRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetTokensResponse: object = {};

export const QueryGetTokensResponse = {
  encode(
    message: QueryGetTokensResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Tokens !== undefined) {
      Tokens.encode(message.Tokens, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTokensResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTokensResponse } as QueryGetTokensResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Tokens = Tokens.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokensResponse {
    const message = { ...baseQueryGetTokensResponse } as QueryGetTokensResponse;
    if (object.Tokens !== undefined && object.Tokens !== null) {
      message.Tokens = Tokens.fromJSON(object.Tokens);
    } else {
      message.Tokens = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTokensResponse): unknown {
    const obj: any = {};
    message.Tokens !== undefined &&
      (obj.Tokens = message.Tokens ? Tokens.toJSON(message.Tokens) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokensResponse>
  ): QueryGetTokensResponse {
    const message = { ...baseQueryGetTokensResponse } as QueryGetTokensResponse;
    if (object.Tokens !== undefined && object.Tokens !== null) {
      message.Tokens = Tokens.fromPartial(object.Tokens);
    } else {
      message.Tokens = undefined;
    }
    return message;
  },
};

const baseQueryAllTokensRequest: object = {};

export const QueryAllTokensRequest = {
  encode(
    message: QueryAllTokensRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTokensRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTokensRequest } as QueryAllTokensRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTokensRequest {
    const message = { ...baseQueryAllTokensRequest } as QueryAllTokensRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokensRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokensRequest>
  ): QueryAllTokensRequest {
    const message = { ...baseQueryAllTokensRequest } as QueryAllTokensRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTokensResponse: object = {};

export const QueryAllTokensResponse = {
  encode(
    message: QueryAllTokensResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Tokens) {
      Tokens.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTokensResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTokensResponse } as QueryAllTokensResponse;
    message.Tokens = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Tokens.push(Tokens.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTokensResponse {
    const message = { ...baseQueryAllTokensResponse } as QueryAllTokensResponse;
    message.Tokens = [];
    if (object.Tokens !== undefined && object.Tokens !== null) {
      for (const e of object.Tokens) {
        message.Tokens.push(Tokens.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokensResponse): unknown {
    const obj: any = {};
    if (message.Tokens) {
      obj.Tokens = message.Tokens.map((e) =>
        e ? Tokens.toJSON(e) : undefined
      );
    } else {
      obj.Tokens = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokensResponse>
  ): QueryAllTokensResponse {
    const message = { ...baseQueryAllTokensResponse } as QueryAllTokensResponse;
    message.Tokens = [];
    if (object.Tokens !== undefined && object.Tokens !== null) {
      for (const e of object.Tokens) {
        message.Tokens.push(Tokens.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetTokenMapRequest: object = { address: "" };

export const QueryGetTokenMapRequest = {
  encode(
    message: QueryGetTokenMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTokenMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTokenMapRequest,
    } as QueryGetTokenMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokenMapRequest {
    const message = {
      ...baseQueryGetTokenMapRequest,
    } as QueryGetTokenMapRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetTokenMapRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokenMapRequest>
  ): QueryGetTokenMapRequest {
    const message = {
      ...baseQueryGetTokenMapRequest,
    } as QueryGetTokenMapRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetTokenMapResponse: object = {};

export const QueryGetTokenMapResponse = {
  encode(
    message: QueryGetTokenMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tokenMap !== undefined) {
      TokenMap.encode(message.tokenMap, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTokenMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTokenMapResponse,
    } as QueryGetTokenMapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenMap = TokenMap.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokenMapResponse {
    const message = {
      ...baseQueryGetTokenMapResponse,
    } as QueryGetTokenMapResponse;
    if (object.tokenMap !== undefined && object.tokenMap !== null) {
      message.tokenMap = TokenMap.fromJSON(object.tokenMap);
    } else {
      message.tokenMap = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTokenMapResponse): unknown {
    const obj: any = {};
    message.tokenMap !== undefined &&
      (obj.tokenMap = message.tokenMap
        ? TokenMap.toJSON(message.tokenMap)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokenMapResponse>
  ): QueryGetTokenMapResponse {
    const message = {
      ...baseQueryGetTokenMapResponse,
    } as QueryGetTokenMapResponse;
    if (object.tokenMap !== undefined && object.tokenMap !== null) {
      message.tokenMap = TokenMap.fromPartial(object.tokenMap);
    } else {
      message.tokenMap = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenMapRequest: object = {};

export const QueryAllTokenMapRequest = {
  encode(
    message: QueryAllTokenMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTokenMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTokenMapRequest,
    } as QueryAllTokenMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTokenMapRequest {
    const message = {
      ...baseQueryAllTokenMapRequest,
    } as QueryAllTokenMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenMapRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokenMapRequest>
  ): QueryAllTokenMapRequest {
    const message = {
      ...baseQueryAllTokenMapRequest,
    } as QueryAllTokenMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenMapResponse: object = {};

export const QueryAllTokenMapResponse = {
  encode(
    message: QueryAllTokenMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.tokenMap) {
      TokenMap.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllTokenMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTokenMapResponse,
    } as QueryAllTokenMapResponse;
    message.tokenMap = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenMap.push(TokenMap.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTokenMapResponse {
    const message = {
      ...baseQueryAllTokenMapResponse,
    } as QueryAllTokenMapResponse;
    message.tokenMap = [];
    if (object.tokenMap !== undefined && object.tokenMap !== null) {
      for (const e of object.tokenMap) {
        message.tokenMap.push(TokenMap.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenMapResponse): unknown {
    const obj: any = {};
    if (message.tokenMap) {
      obj.tokenMap = message.tokenMap.map((e) =>
        e ? TokenMap.toJSON(e) : undefined
      );
    } else {
      obj.tokenMap = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokenMapResponse>
  ): QueryAllTokenMapResponse {
    const message = {
      ...baseQueryAllTokenMapResponse,
    } as QueryAllTokenMapResponse;
    message.tokenMap = [];
    if (object.tokenMap !== undefined && object.tokenMap !== null) {
      for (const e of object.tokenMap) {
        message.tokenMap.push(TokenMap.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetSharesRequest: object = {
  address: "",
  pairId: "",
  tickIndex: 0,
  fee: 0,
};

export const QueryGetSharesRequest = {
  encode(
    message: QueryGetSharesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.pairId !== "") {
      writer.uint32(18).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(24).int64(message.tickIndex);
    }
    if (message.fee !== 0) {
      writer.uint32(32).uint64(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetSharesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetSharesRequest } as QueryGetSharesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.pairId = reader.string();
          break;
        case 3:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.fee = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSharesRequest {
    const message = { ...baseQueryGetSharesRequest } as QueryGetSharesRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = Number(object.fee);
    } else {
      message.fee = 0;
    }
    return message;
  },

  toJSON(message: QueryGetSharesRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSharesRequest>
  ): QueryGetSharesRequest {
    const message = { ...baseQueryGetSharesRequest } as QueryGetSharesRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = 0;
    }
    return message;
  },
};

const baseQueryGetSharesResponse: object = {};

export const QueryGetSharesResponse = {
  encode(
    message: QueryGetSharesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.shares !== undefined) {
      Shares.encode(message.shares, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetSharesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetSharesResponse } as QueryGetSharesResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.shares = Shares.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSharesResponse {
    const message = { ...baseQueryGetSharesResponse } as QueryGetSharesResponse;
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Shares.fromJSON(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSharesResponse): unknown {
    const obj: any = {};
    message.shares !== undefined &&
      (obj.shares = message.shares ? Shares.toJSON(message.shares) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSharesResponse>
  ): QueryGetSharesResponse {
    const message = { ...baseQueryGetSharesResponse } as QueryGetSharesResponse;
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Shares.fromPartial(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },
};

const baseQueryAllSharesRequest: object = {};

export const QueryAllSharesRequest = {
  encode(
    message: QueryAllSharesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllSharesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllSharesRequest } as QueryAllSharesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSharesRequest {
    const message = { ...baseQueryAllSharesRequest } as QueryAllSharesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSharesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSharesRequest>
  ): QueryAllSharesRequest {
    const message = { ...baseQueryAllSharesRequest } as QueryAllSharesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSharesResponse: object = {};

export const QueryAllSharesResponse = {
  encode(
    message: QueryAllSharesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.shares) {
      Shares.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllSharesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllSharesResponse } as QueryAllSharesResponse;
    message.shares = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.shares.push(Shares.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSharesResponse {
    const message = { ...baseQueryAllSharesResponse } as QueryAllSharesResponse;
    message.shares = [];
    if (object.shares !== undefined && object.shares !== null) {
      for (const e of object.shares) {
        message.shares.push(Shares.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSharesResponse): unknown {
    const obj: any = {};
    if (message.shares) {
      obj.shares = message.shares.map((e) =>
        e ? Shares.toJSON(e) : undefined
      );
    } else {
      obj.shares = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSharesResponse>
  ): QueryAllSharesResponse {
    const message = { ...baseQueryAllSharesResponse } as QueryAllSharesResponse;
    message.shares = [];
    if (object.shares !== undefined && object.shares !== null) {
      for (const e of object.shares) {
        message.shares.push(Shares.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetFeeListRequest: object = { id: 0 };

export const QueryGetFeeListRequest = {
  encode(
    message: QueryGetFeeListRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFeeListRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFeeListRequest } as QueryGetFeeListRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeListRequest {
    const message = { ...baseQueryGetFeeListRequest } as QueryGetFeeListRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetFeeListRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFeeListRequest>
  ): QueryGetFeeListRequest {
    const message = { ...baseQueryGetFeeListRequest } as QueryGetFeeListRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetFeeListResponse: object = {};

export const QueryGetFeeListResponse = {
  encode(
    message: QueryGetFeeListResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.FeeList !== undefined) {
      FeeList.encode(message.FeeList, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFeeListResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFeeListResponse,
    } as QueryGetFeeListResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.FeeList = FeeList.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeListResponse {
    const message = {
      ...baseQueryGetFeeListResponse,
    } as QueryGetFeeListResponse;
    if (object.FeeList !== undefined && object.FeeList !== null) {
      message.FeeList = FeeList.fromJSON(object.FeeList);
    } else {
      message.FeeList = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFeeListResponse): unknown {
    const obj: any = {};
    message.FeeList !== undefined &&
      (obj.FeeList = message.FeeList
        ? FeeList.toJSON(message.FeeList)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFeeListResponse>
  ): QueryGetFeeListResponse {
    const message = {
      ...baseQueryGetFeeListResponse,
    } as QueryGetFeeListResponse;
    if (object.FeeList !== undefined && object.FeeList !== null) {
      message.FeeList = FeeList.fromPartial(object.FeeList);
    } else {
      message.FeeList = undefined;
    }
    return message;
  },
};

const baseQueryAllFeeListRequest: object = {};

export const QueryAllFeeListRequest = {
  encode(
    message: QueryAllFeeListRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFeeListRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFeeListRequest } as QueryAllFeeListRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFeeListRequest {
    const message = { ...baseQueryAllFeeListRequest } as QueryAllFeeListRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFeeListRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFeeListRequest>
  ): QueryAllFeeListRequest {
    const message = { ...baseQueryAllFeeListRequest } as QueryAllFeeListRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllFeeListResponse: object = {};

export const QueryAllFeeListResponse = {
  encode(
    message: QueryAllFeeListResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.FeeList) {
      FeeList.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFeeListResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllFeeListResponse,
    } as QueryAllFeeListResponse;
    message.FeeList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.FeeList.push(FeeList.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFeeListResponse {
    const message = {
      ...baseQueryAllFeeListResponse,
    } as QueryAllFeeListResponse;
    message.FeeList = [];
    if (object.FeeList !== undefined && object.FeeList !== null) {
      for (const e of object.FeeList) {
        message.FeeList.push(FeeList.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFeeListResponse): unknown {
    const obj: any = {};
    if (message.FeeList) {
      obj.FeeList = message.FeeList.map((e) =>
        e ? FeeList.toJSON(e) : undefined
      );
    } else {
      obj.FeeList = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFeeListResponse>
  ): QueryAllFeeListResponse {
    const message = {
      ...baseQueryAllFeeListResponse,
    } as QueryAllFeeListResponse;
    message.FeeList = [];
    if (object.FeeList !== undefined && object.FeeList !== null) {
      for (const e of object.FeeList) {
        message.FeeList.push(FeeList.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetEdgeRowRequest: object = { id: 0 };

export const QueryGetEdgeRowRequest = {
  encode(
    message: QueryGetEdgeRowRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEdgeRowRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetEdgeRowRequest } as QueryGetEdgeRowRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEdgeRowRequest {
    const message = { ...baseQueryGetEdgeRowRequest } as QueryGetEdgeRowRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetEdgeRowRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetEdgeRowRequest>
  ): QueryGetEdgeRowRequest {
    const message = { ...baseQueryGetEdgeRowRequest } as QueryGetEdgeRowRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetEdgeRowResponse: object = {};

export const QueryGetEdgeRowResponse = {
  encode(
    message: QueryGetEdgeRowResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.EdgeRow !== undefined) {
      EdgeRow.encode(message.EdgeRow, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEdgeRowResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetEdgeRowResponse,
    } as QueryGetEdgeRowResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.EdgeRow = EdgeRow.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEdgeRowResponse {
    const message = {
      ...baseQueryGetEdgeRowResponse,
    } as QueryGetEdgeRowResponse;
    if (object.EdgeRow !== undefined && object.EdgeRow !== null) {
      message.EdgeRow = EdgeRow.fromJSON(object.EdgeRow);
    } else {
      message.EdgeRow = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetEdgeRowResponse): unknown {
    const obj: any = {};
    message.EdgeRow !== undefined &&
      (obj.EdgeRow = message.EdgeRow
        ? EdgeRow.toJSON(message.EdgeRow)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetEdgeRowResponse>
  ): QueryGetEdgeRowResponse {
    const message = {
      ...baseQueryGetEdgeRowResponse,
    } as QueryGetEdgeRowResponse;
    if (object.EdgeRow !== undefined && object.EdgeRow !== null) {
      message.EdgeRow = EdgeRow.fromPartial(object.EdgeRow);
    } else {
      message.EdgeRow = undefined;
    }
    return message;
  },
};

const baseQueryAllEdgeRowRequest: object = {};

export const QueryAllEdgeRowRequest = {
  encode(
    message: QueryAllEdgeRowRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllEdgeRowRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllEdgeRowRequest } as QueryAllEdgeRowRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEdgeRowRequest {
    const message = { ...baseQueryAllEdgeRowRequest } as QueryAllEdgeRowRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllEdgeRowRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllEdgeRowRequest>
  ): QueryAllEdgeRowRequest {
    const message = { ...baseQueryAllEdgeRowRequest } as QueryAllEdgeRowRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllEdgeRowResponse: object = {};

export const QueryAllEdgeRowResponse = {
  encode(
    message: QueryAllEdgeRowResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.EdgeRow) {
      EdgeRow.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllEdgeRowResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllEdgeRowResponse,
    } as QueryAllEdgeRowResponse;
    message.EdgeRow = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.EdgeRow.push(EdgeRow.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEdgeRowResponse {
    const message = {
      ...baseQueryAllEdgeRowResponse,
    } as QueryAllEdgeRowResponse;
    message.EdgeRow = [];
    if (object.EdgeRow !== undefined && object.EdgeRow !== null) {
      for (const e of object.EdgeRow) {
        message.EdgeRow.push(EdgeRow.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllEdgeRowResponse): unknown {
    const obj: any = {};
    if (message.EdgeRow) {
      obj.EdgeRow = message.EdgeRow.map((e) =>
        e ? EdgeRow.toJSON(e) : undefined
      );
    } else {
      obj.EdgeRow = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllEdgeRowResponse>
  ): QueryAllEdgeRowResponse {
    const message = {
      ...baseQueryAllEdgeRowResponse,
    } as QueryAllEdgeRowResponse;
    message.EdgeRow = [];
    if (object.EdgeRow !== undefined && object.EdgeRow !== null) {
      for (const e of object.EdgeRow) {
        message.EdgeRow.push(EdgeRow.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetAdjanceyMatrixRequest: object = { id: 0 };

export const QueryGetAdjanceyMatrixRequest = {
  encode(
    message: QueryGetAdjanceyMatrixRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAdjanceyMatrixRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAdjanceyMatrixRequest,
    } as QueryGetAdjanceyMatrixRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAdjanceyMatrixRequest {
    const message = {
      ...baseQueryGetAdjanceyMatrixRequest,
    } as QueryGetAdjanceyMatrixRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetAdjanceyMatrixRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAdjanceyMatrixRequest>
  ): QueryGetAdjanceyMatrixRequest {
    const message = {
      ...baseQueryGetAdjanceyMatrixRequest,
    } as QueryGetAdjanceyMatrixRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetAdjanceyMatrixResponse: object = {};

export const QueryGetAdjanceyMatrixResponse = {
  encode(
    message: QueryGetAdjanceyMatrixResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.AdjanceyMatrix !== undefined) {
      AdjanceyMatrix.encode(
        message.AdjanceyMatrix,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAdjanceyMatrixResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAdjanceyMatrixResponse,
    } as QueryGetAdjanceyMatrixResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AdjanceyMatrix = AdjanceyMatrix.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAdjanceyMatrixResponse {
    const message = {
      ...baseQueryGetAdjanceyMatrixResponse,
    } as QueryGetAdjanceyMatrixResponse;
    if (object.AdjanceyMatrix !== undefined && object.AdjanceyMatrix !== null) {
      message.AdjanceyMatrix = AdjanceyMatrix.fromJSON(object.AdjanceyMatrix);
    } else {
      message.AdjanceyMatrix = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAdjanceyMatrixResponse): unknown {
    const obj: any = {};
    message.AdjanceyMatrix !== undefined &&
      (obj.AdjanceyMatrix = message.AdjanceyMatrix
        ? AdjanceyMatrix.toJSON(message.AdjanceyMatrix)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAdjanceyMatrixResponse>
  ): QueryGetAdjanceyMatrixResponse {
    const message = {
      ...baseQueryGetAdjanceyMatrixResponse,
    } as QueryGetAdjanceyMatrixResponse;
    if (object.AdjanceyMatrix !== undefined && object.AdjanceyMatrix !== null) {
      message.AdjanceyMatrix = AdjanceyMatrix.fromPartial(
        object.AdjanceyMatrix
      );
    } else {
      message.AdjanceyMatrix = undefined;
    }
    return message;
  },
};

const baseQueryAllAdjanceyMatrixRequest: object = {};

export const QueryAllAdjanceyMatrixRequest = {
  encode(
    message: QueryAllAdjanceyMatrixRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllAdjanceyMatrixRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAdjanceyMatrixRequest,
    } as QueryAllAdjanceyMatrixRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAdjanceyMatrixRequest {
    const message = {
      ...baseQueryAllAdjanceyMatrixRequest,
    } as QueryAllAdjanceyMatrixRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAdjanceyMatrixRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAdjanceyMatrixRequest>
  ): QueryAllAdjanceyMatrixRequest {
    const message = {
      ...baseQueryAllAdjanceyMatrixRequest,
    } as QueryAllAdjanceyMatrixRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllAdjanceyMatrixResponse: object = {};

export const QueryAllAdjanceyMatrixResponse = {
  encode(
    message: QueryAllAdjanceyMatrixResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.AdjanceyMatrix) {
      AdjanceyMatrix.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllAdjanceyMatrixResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAdjanceyMatrixResponse,
    } as QueryAllAdjanceyMatrixResponse;
    message.AdjanceyMatrix = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AdjanceyMatrix.push(
            AdjanceyMatrix.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAdjanceyMatrixResponse {
    const message = {
      ...baseQueryAllAdjanceyMatrixResponse,
    } as QueryAllAdjanceyMatrixResponse;
    message.AdjanceyMatrix = [];
    if (object.AdjanceyMatrix !== undefined && object.AdjanceyMatrix !== null) {
      for (const e of object.AdjanceyMatrix) {
        message.AdjanceyMatrix.push(AdjanceyMatrix.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAdjanceyMatrixResponse): unknown {
    const obj: any = {};
    if (message.AdjanceyMatrix) {
      obj.AdjanceyMatrix = message.AdjanceyMatrix.map((e) =>
        e ? AdjanceyMatrix.toJSON(e) : undefined
      );
    } else {
      obj.AdjanceyMatrix = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAdjanceyMatrixResponse>
  ): QueryAllAdjanceyMatrixResponse {
    const message = {
      ...baseQueryAllAdjanceyMatrixResponse,
    } as QueryAllAdjanceyMatrixResponse;
    message.AdjanceyMatrix = [];
    if (object.AdjanceyMatrix !== undefined && object.AdjanceyMatrix !== null) {
      for (const e of object.AdjanceyMatrix) {
        message.AdjanceyMatrix.push(AdjanceyMatrix.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolUserShareMapRequest: object = {
  count: 0,
  address: "",
  pairId: "",
  tickIndex: 0,
  token: "",
};

export const QueryGetLimitOrderPoolUserShareMapRequest = {
  encode(
    message: QueryGetLimitOrderPoolUserShareMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.pairId !== "") {
      writer.uint32(26).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(32).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(42).string(message.token);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolUserShareMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareMapRequest,
    } as QueryGetLimitOrderPoolUserShareMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.pairId = reader.string();
          break;
        case 4:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolUserShareMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareMapRequest,
    } as QueryGetLimitOrderPoolUserShareMapRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolUserShareMapRequest): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.address !== undefined && (obj.address = message.address);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserShareMapRequest>
  ): QueryGetLimitOrderPoolUserShareMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareMapRequest,
    } as QueryGetLimitOrderPoolUserShareMapRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolUserShareMapResponse: object = {};

export const QueryGetLimitOrderPoolUserShareMapResponse = {
  encode(
    message: QueryGetLimitOrderPoolUserShareMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolUserShareMap !== undefined) {
      LimitOrderPoolUserShareMap.encode(
        message.limitOrderPoolUserShareMap,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolUserShareMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareMapResponse,
    } as QueryGetLimitOrderPoolUserShareMapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolUserShareMap = LimitOrderPoolUserShareMap.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolUserShareMapResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareMapResponse,
    } as QueryGetLimitOrderPoolUserShareMapResponse;
    if (
      object.limitOrderPoolUserShareMap !== undefined &&
      object.limitOrderPoolUserShareMap !== null
    ) {
      message.limitOrderPoolUserShareMap = LimitOrderPoolUserShareMap.fromJSON(
        object.limitOrderPoolUserShareMap
      );
    } else {
      message.limitOrderPoolUserShareMap = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolUserShareMapResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolUserShareMap !== undefined &&
      (obj.limitOrderPoolUserShareMap = message.limitOrderPoolUserShareMap
        ? LimitOrderPoolUserShareMap.toJSON(message.limitOrderPoolUserShareMap)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserShareMapResponse>
  ): QueryGetLimitOrderPoolUserShareMapResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareMapResponse,
    } as QueryGetLimitOrderPoolUserShareMapResponse;
    if (
      object.limitOrderPoolUserShareMap !== undefined &&
      object.limitOrderPoolUserShareMap !== null
    ) {
      message.limitOrderPoolUserShareMap = LimitOrderPoolUserShareMap.fromPartial(
        object.limitOrderPoolUserShareMap
      );
    } else {
      message.limitOrderPoolUserShareMap = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolUserShareMapRequest: object = {};

export const QueryAllLimitOrderPoolUserShareMapRequest = {
  encode(
    message: QueryAllLimitOrderPoolUserShareMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolUserShareMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareMapRequest,
    } as QueryAllLimitOrderPoolUserShareMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolUserShareMapRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareMapRequest,
    } as QueryAllLimitOrderPoolUserShareMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolUserShareMapRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolUserShareMapRequest>
  ): QueryAllLimitOrderPoolUserShareMapRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareMapRequest,
    } as QueryAllLimitOrderPoolUserShareMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolUserShareMapResponse: object = {};

export const QueryAllLimitOrderPoolUserShareMapResponse = {
  encode(
    message: QueryAllLimitOrderPoolUserShareMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolUserShareMap) {
      LimitOrderPoolUserShareMap.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolUserShareMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareMapResponse,
    } as QueryAllLimitOrderPoolUserShareMapResponse;
    message.limitOrderPoolUserShareMap = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolUserShareMap.push(
            LimitOrderPoolUserShareMap.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolUserShareMapResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareMapResponse,
    } as QueryAllLimitOrderPoolUserShareMapResponse;
    message.limitOrderPoolUserShareMap = [];
    if (
      object.limitOrderPoolUserShareMap !== undefined &&
      object.limitOrderPoolUserShareMap !== null
    ) {
      for (const e of object.limitOrderPoolUserShareMap) {
        message.limitOrderPoolUserShareMap.push(
          LimitOrderPoolUserShareMap.fromJSON(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolUserShareMapResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolUserShareMap) {
      obj.limitOrderPoolUserShareMap = message.limitOrderPoolUserShareMap.map(
        (e) => (e ? LimitOrderPoolUserShareMap.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolUserShareMap = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolUserShareMapResponse>
  ): QueryAllLimitOrderPoolUserShareMapResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareMapResponse,
    } as QueryAllLimitOrderPoolUserShareMapResponse;
    message.limitOrderPoolUserShareMap = [];
    if (
      object.limitOrderPoolUserShareMap !== undefined &&
      object.limitOrderPoolUserShareMap !== null
    ) {
      for (const e of object.limitOrderPoolUserShareMap) {
        message.limitOrderPoolUserShareMap.push(
          LimitOrderPoolUserShareMap.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolUserSharesWithdrawnRequest: object = {
  count: 0,
  address: "",
  pairId: "",
  tickIndex: 0,
  token: "",
};

export const QueryGetLimitOrderPoolUserSharesWithdrawnRequest = {
  encode(
    message: QueryGetLimitOrderPoolUserSharesWithdrawnRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.pairId !== "") {
      writer.uint32(26).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(32).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(42).string(message.token);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolUserSharesWithdrawnRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnRequest,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.pairId = reader.string();
          break;
        case 4:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolUserSharesWithdrawnRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnRequest,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolUserSharesWithdrawnRequest): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.address !== undefined && (obj.address = message.address);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserSharesWithdrawnRequest>
  ): QueryGetLimitOrderPoolUserSharesWithdrawnRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnRequest,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolUserSharesWithdrawnResponse: object = {};

export const QueryGetLimitOrderPoolUserSharesWithdrawnResponse = {
  encode(
    message: QueryGetLimitOrderPoolUserSharesWithdrawnResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolUserSharesWithdrawn !== undefined) {
      LimitOrderPoolUserSharesWithdrawn.encode(
        message.limitOrderPoolUserSharesWithdrawn,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolUserSharesWithdrawnResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnResponse,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolUserSharesWithdrawn = LimitOrderPoolUserSharesWithdrawn.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolUserSharesWithdrawnResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnResponse,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnResponse;
    if (
      object.limitOrderPoolUserSharesWithdrawn !== undefined &&
      object.limitOrderPoolUserSharesWithdrawn !== null
    ) {
      message.limitOrderPoolUserSharesWithdrawn = LimitOrderPoolUserSharesWithdrawn.fromJSON(
        object.limitOrderPoolUserSharesWithdrawn
      );
    } else {
      message.limitOrderPoolUserSharesWithdrawn = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolUserSharesWithdrawnResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolUserSharesWithdrawn !== undefined &&
      (obj.limitOrderPoolUserSharesWithdrawn = message.limitOrderPoolUserSharesWithdrawn
        ? LimitOrderPoolUserSharesWithdrawn.toJSON(
            message.limitOrderPoolUserSharesWithdrawn
          )
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserSharesWithdrawnResponse>
  ): QueryGetLimitOrderPoolUserSharesWithdrawnResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnResponse,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnResponse;
    if (
      object.limitOrderPoolUserSharesWithdrawn !== undefined &&
      object.limitOrderPoolUserSharesWithdrawn !== null
    ) {
      message.limitOrderPoolUserSharesWithdrawn = LimitOrderPoolUserSharesWithdrawn.fromPartial(
        object.limitOrderPoolUserSharesWithdrawn
      );
    } else {
      message.limitOrderPoolUserSharesWithdrawn = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolUserSharesWithdrawnRequest: object = {};

export const QueryAllLimitOrderPoolUserSharesWithdrawnRequest = {
  encode(
    message: QueryAllLimitOrderPoolUserSharesWithdrawnRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolUserSharesWithdrawnRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnRequest,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolUserSharesWithdrawnRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnRequest,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolUserSharesWithdrawnRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolUserSharesWithdrawnRequest>
  ): QueryAllLimitOrderPoolUserSharesWithdrawnRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnRequest,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolUserSharesWithdrawnResponse: object = {};

export const QueryAllLimitOrderPoolUserSharesWithdrawnResponse = {
  encode(
    message: QueryAllLimitOrderPoolUserSharesWithdrawnResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolUserSharesWithdrawn) {
      LimitOrderPoolUserSharesWithdrawn.encode(
        v!,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolUserSharesWithdrawnResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnResponse,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnResponse;
    message.limitOrderPoolUserSharesWithdrawn = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolUserSharesWithdrawn.push(
            LimitOrderPoolUserSharesWithdrawn.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolUserSharesWithdrawnResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnResponse,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnResponse;
    message.limitOrderPoolUserSharesWithdrawn = [];
    if (
      object.limitOrderPoolUserSharesWithdrawn !== undefined &&
      object.limitOrderPoolUserSharesWithdrawn !== null
    ) {
      for (const e of object.limitOrderPoolUserSharesWithdrawn) {
        message.limitOrderPoolUserSharesWithdrawn.push(
          LimitOrderPoolUserSharesWithdrawn.fromJSON(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolUserSharesWithdrawnResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolUserSharesWithdrawn) {
      obj.limitOrderPoolUserSharesWithdrawn = message.limitOrderPoolUserSharesWithdrawn.map(
        (e) => (e ? LimitOrderPoolUserSharesWithdrawn.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolUserSharesWithdrawn = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolUserSharesWithdrawnResponse>
  ): QueryAllLimitOrderPoolUserSharesWithdrawnResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnResponse,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnResponse;
    message.limitOrderPoolUserSharesWithdrawn = [];
    if (
      object.limitOrderPoolUserSharesWithdrawn !== undefined &&
      object.limitOrderPoolUserSharesWithdrawn !== null
    ) {
      for (const e of object.limitOrderPoolUserSharesWithdrawn) {
        message.limitOrderPoolUserSharesWithdrawn.push(
          LimitOrderPoolUserSharesWithdrawn.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolTotalSharesMapRequest: object = {
  count: 0,
  pairId: "",
  tickIndex: 0,
  token: "",
};

export const QueryGetLimitOrderPoolTotalSharesMapRequest = {
  encode(
    message: QueryGetLimitOrderPoolTotalSharesMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.pairId !== "") {
      writer.uint32(18).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(24).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(34).string(message.token);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolTotalSharesMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesMapRequest,
    } as QueryGetLimitOrderPoolTotalSharesMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.pairId = reader.string();
          break;
        case 3:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolTotalSharesMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesMapRequest,
    } as QueryGetLimitOrderPoolTotalSharesMapRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolTotalSharesMapRequest): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolTotalSharesMapRequest>
  ): QueryGetLimitOrderPoolTotalSharesMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesMapRequest,
    } as QueryGetLimitOrderPoolTotalSharesMapRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolTotalSharesMapResponse: object = {};

export const QueryGetLimitOrderPoolTotalSharesMapResponse = {
  encode(
    message: QueryGetLimitOrderPoolTotalSharesMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolTotalSharesMap !== undefined) {
      LimitOrderPoolTotalSharesMap.encode(
        message.limitOrderPoolTotalSharesMap,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolTotalSharesMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesMapResponse,
    } as QueryGetLimitOrderPoolTotalSharesMapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolTotalSharesMap = LimitOrderPoolTotalSharesMap.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolTotalSharesMapResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesMapResponse,
    } as QueryGetLimitOrderPoolTotalSharesMapResponse;
    if (
      object.limitOrderPoolTotalSharesMap !== undefined &&
      object.limitOrderPoolTotalSharesMap !== null
    ) {
      message.limitOrderPoolTotalSharesMap = LimitOrderPoolTotalSharesMap.fromJSON(
        object.limitOrderPoolTotalSharesMap
      );
    } else {
      message.limitOrderPoolTotalSharesMap = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolTotalSharesMapResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolTotalSharesMap !== undefined &&
      (obj.limitOrderPoolTotalSharesMap = message.limitOrderPoolTotalSharesMap
        ? LimitOrderPoolTotalSharesMap.toJSON(
            message.limitOrderPoolTotalSharesMap
          )
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolTotalSharesMapResponse>
  ): QueryGetLimitOrderPoolTotalSharesMapResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesMapResponse,
    } as QueryGetLimitOrderPoolTotalSharesMapResponse;
    if (
      object.limitOrderPoolTotalSharesMap !== undefined &&
      object.limitOrderPoolTotalSharesMap !== null
    ) {
      message.limitOrderPoolTotalSharesMap = LimitOrderPoolTotalSharesMap.fromPartial(
        object.limitOrderPoolTotalSharesMap
      );
    } else {
      message.limitOrderPoolTotalSharesMap = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolTotalSharesMapRequest: object = {};

export const QueryAllLimitOrderPoolTotalSharesMapRequest = {
  encode(
    message: QueryAllLimitOrderPoolTotalSharesMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolTotalSharesMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesMapRequest,
    } as QueryAllLimitOrderPoolTotalSharesMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolTotalSharesMapRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesMapRequest,
    } as QueryAllLimitOrderPoolTotalSharesMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolTotalSharesMapRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolTotalSharesMapRequest>
  ): QueryAllLimitOrderPoolTotalSharesMapRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesMapRequest,
    } as QueryAllLimitOrderPoolTotalSharesMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolTotalSharesMapResponse: object = {};

export const QueryAllLimitOrderPoolTotalSharesMapResponse = {
  encode(
    message: QueryAllLimitOrderPoolTotalSharesMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolTotalSharesMap) {
      LimitOrderPoolTotalSharesMap.encode(
        v!,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolTotalSharesMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesMapResponse,
    } as QueryAllLimitOrderPoolTotalSharesMapResponse;
    message.limitOrderPoolTotalSharesMap = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolTotalSharesMap.push(
            LimitOrderPoolTotalSharesMap.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolTotalSharesMapResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesMapResponse,
    } as QueryAllLimitOrderPoolTotalSharesMapResponse;
    message.limitOrderPoolTotalSharesMap = [];
    if (
      object.limitOrderPoolTotalSharesMap !== undefined &&
      object.limitOrderPoolTotalSharesMap !== null
    ) {
      for (const e of object.limitOrderPoolTotalSharesMap) {
        message.limitOrderPoolTotalSharesMap.push(
          LimitOrderPoolTotalSharesMap.fromJSON(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolTotalSharesMapResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolTotalSharesMap) {
      obj.limitOrderPoolTotalSharesMap = message.limitOrderPoolTotalSharesMap.map(
        (e) => (e ? LimitOrderPoolTotalSharesMap.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolTotalSharesMap = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolTotalSharesMapResponse>
  ): QueryAllLimitOrderPoolTotalSharesMapResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesMapResponse,
    } as QueryAllLimitOrderPoolTotalSharesMapResponse;
    message.limitOrderPoolTotalSharesMap = [];
    if (
      object.limitOrderPoolTotalSharesMap !== undefined &&
      object.limitOrderPoolTotalSharesMap !== null
    ) {
      for (const e of object.limitOrderPoolTotalSharesMap) {
        message.limitOrderPoolTotalSharesMap.push(
          LimitOrderPoolTotalSharesMap.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolReserveMapRequest: object = {
  count: 0,
  pairId: "",
  tickIndex: 0,
  token: "",
};

export const QueryGetLimitOrderPoolReserveMapRequest = {
  encode(
    message: QueryGetLimitOrderPoolReserveMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.pairId !== "") {
      writer.uint32(18).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(24).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(34).string(message.token);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolReserveMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolReserveMapRequest,
    } as QueryGetLimitOrderPoolReserveMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.pairId = reader.string();
          break;
        case 3:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolReserveMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveMapRequest,
    } as QueryGetLimitOrderPoolReserveMapRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolReserveMapRequest): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolReserveMapRequest>
  ): QueryGetLimitOrderPoolReserveMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveMapRequest,
    } as QueryGetLimitOrderPoolReserveMapRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolReserveMapResponse: object = {};

export const QueryGetLimitOrderPoolReserveMapResponse = {
  encode(
    message: QueryGetLimitOrderPoolReserveMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolReserveMap !== undefined) {
      LimitOrderPoolReserveMap.encode(
        message.limitOrderPoolReserveMap,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolReserveMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolReserveMapResponse,
    } as QueryGetLimitOrderPoolReserveMapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolReserveMap = LimitOrderPoolReserveMap.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolReserveMapResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveMapResponse,
    } as QueryGetLimitOrderPoolReserveMapResponse;
    if (
      object.limitOrderPoolReserveMap !== undefined &&
      object.limitOrderPoolReserveMap !== null
    ) {
      message.limitOrderPoolReserveMap = LimitOrderPoolReserveMap.fromJSON(
        object.limitOrderPoolReserveMap
      );
    } else {
      message.limitOrderPoolReserveMap = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolReserveMapResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolReserveMap !== undefined &&
      (obj.limitOrderPoolReserveMap = message.limitOrderPoolReserveMap
        ? LimitOrderPoolReserveMap.toJSON(message.limitOrderPoolReserveMap)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolReserveMapResponse>
  ): QueryGetLimitOrderPoolReserveMapResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveMapResponse,
    } as QueryGetLimitOrderPoolReserveMapResponse;
    if (
      object.limitOrderPoolReserveMap !== undefined &&
      object.limitOrderPoolReserveMap !== null
    ) {
      message.limitOrderPoolReserveMap = LimitOrderPoolReserveMap.fromPartial(
        object.limitOrderPoolReserveMap
      );
    } else {
      message.limitOrderPoolReserveMap = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolReserveMapRequest: object = {};

export const QueryAllLimitOrderPoolReserveMapRequest = {
  encode(
    message: QueryAllLimitOrderPoolReserveMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolReserveMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolReserveMapRequest,
    } as QueryAllLimitOrderPoolReserveMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolReserveMapRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolReserveMapRequest,
    } as QueryAllLimitOrderPoolReserveMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolReserveMapRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolReserveMapRequest>
  ): QueryAllLimitOrderPoolReserveMapRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolReserveMapRequest,
    } as QueryAllLimitOrderPoolReserveMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolReserveMapResponse: object = {};

export const QueryAllLimitOrderPoolReserveMapResponse = {
  encode(
    message: QueryAllLimitOrderPoolReserveMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolReserveMap) {
      LimitOrderPoolReserveMap.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolReserveMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolReserveMapResponse,
    } as QueryAllLimitOrderPoolReserveMapResponse;
    message.limitOrderPoolReserveMap = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolReserveMap.push(
            LimitOrderPoolReserveMap.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolReserveMapResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolReserveMapResponse,
    } as QueryAllLimitOrderPoolReserveMapResponse;
    message.limitOrderPoolReserveMap = [];
    if (
      object.limitOrderPoolReserveMap !== undefined &&
      object.limitOrderPoolReserveMap !== null
    ) {
      for (const e of object.limitOrderPoolReserveMap) {
        message.limitOrderPoolReserveMap.push(
          LimitOrderPoolReserveMap.fromJSON(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolReserveMapResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolReserveMap) {
      obj.limitOrderPoolReserveMap = message.limitOrderPoolReserveMap.map((e) =>
        e ? LimitOrderPoolReserveMap.toJSON(e) : undefined
      );
    } else {
      obj.limitOrderPoolReserveMap = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolReserveMapResponse>
  ): QueryAllLimitOrderPoolReserveMapResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolReserveMapResponse,
    } as QueryAllLimitOrderPoolReserveMapResponse;
    message.limitOrderPoolReserveMap = [];
    if (
      object.limitOrderPoolReserveMap !== undefined &&
      object.limitOrderPoolReserveMap !== null
    ) {
      for (const e of object.limitOrderPoolReserveMap) {
        message.limitOrderPoolReserveMap.push(
          LimitOrderPoolReserveMap.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolFillMapRequest: object = {
  count: 0,
  pairId: "",
  tickIndex: 0,
  token: "",
};

export const QueryGetLimitOrderPoolFillMapRequest = {
  encode(
    message: QueryGetLimitOrderPoolFillMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.pairId !== "") {
      writer.uint32(18).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(24).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(34).string(message.token);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolFillMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolFillMapRequest,
    } as QueryGetLimitOrderPoolFillMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.pairId = reader.string();
          break;
        case 3:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolFillMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolFillMapRequest,
    } as QueryGetLimitOrderPoolFillMapRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolFillMapRequest): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolFillMapRequest>
  ): QueryGetLimitOrderPoolFillMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolFillMapRequest,
    } as QueryGetLimitOrderPoolFillMapRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    return message;
  },
};

const baseQueryGetLimitOrderPoolFillMapResponse: object = {};

export const QueryGetLimitOrderPoolFillMapResponse = {
  encode(
    message: QueryGetLimitOrderPoolFillMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolFillMap !== undefined) {
      LimitOrderPoolFillMap.encode(
        message.limitOrderPoolFillMap,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolFillMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolFillMapResponse,
    } as QueryGetLimitOrderPoolFillMapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolFillMap = LimitOrderPoolFillMap.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLimitOrderPoolFillMapResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolFillMapResponse,
    } as QueryGetLimitOrderPoolFillMapResponse;
    if (
      object.limitOrderPoolFillMap !== undefined &&
      object.limitOrderPoolFillMap !== null
    ) {
      message.limitOrderPoolFillMap = LimitOrderPoolFillMap.fromJSON(
        object.limitOrderPoolFillMap
      );
    } else {
      message.limitOrderPoolFillMap = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolFillMapResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolFillMap !== undefined &&
      (obj.limitOrderPoolFillMap = message.limitOrderPoolFillMap
        ? LimitOrderPoolFillMap.toJSON(message.limitOrderPoolFillMap)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolFillMapResponse>
  ): QueryGetLimitOrderPoolFillMapResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolFillMapResponse,
    } as QueryGetLimitOrderPoolFillMapResponse;
    if (
      object.limitOrderPoolFillMap !== undefined &&
      object.limitOrderPoolFillMap !== null
    ) {
      message.limitOrderPoolFillMap = LimitOrderPoolFillMap.fromPartial(
        object.limitOrderPoolFillMap
      );
    } else {
      message.limitOrderPoolFillMap = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolFillMapRequest: object = {};

export const QueryAllLimitOrderPoolFillMapRequest = {
  encode(
    message: QueryAllLimitOrderPoolFillMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolFillMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolFillMapRequest,
    } as QueryAllLimitOrderPoolFillMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolFillMapRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolFillMapRequest,
    } as QueryAllLimitOrderPoolFillMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolFillMapRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolFillMapRequest>
  ): QueryAllLimitOrderPoolFillMapRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolFillMapRequest,
    } as QueryAllLimitOrderPoolFillMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolFillMapResponse: object = {};

export const QueryAllLimitOrderPoolFillMapResponse = {
  encode(
    message: QueryAllLimitOrderPoolFillMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolFillMap) {
      LimitOrderPoolFillMap.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLimitOrderPoolFillMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolFillMapResponse,
    } as QueryAllLimitOrderPoolFillMapResponse;
    message.limitOrderPoolFillMap = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolFillMap.push(
            LimitOrderPoolFillMap.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLimitOrderPoolFillMapResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolFillMapResponse,
    } as QueryAllLimitOrderPoolFillMapResponse;
    message.limitOrderPoolFillMap = [];
    if (
      object.limitOrderPoolFillMap !== undefined &&
      object.limitOrderPoolFillMap !== null
    ) {
      for (const e of object.limitOrderPoolFillMap) {
        message.limitOrderPoolFillMap.push(LimitOrderPoolFillMap.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolFillMapResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolFillMap) {
      obj.limitOrderPoolFillMap = message.limitOrderPoolFillMap.map((e) =>
        e ? LimitOrderPoolFillMap.toJSON(e) : undefined
      );
    } else {
      obj.limitOrderPoolFillMap = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolFillMapResponse>
  ): QueryAllLimitOrderPoolFillMapResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolFillMapResponse,
    } as QueryAllLimitOrderPoolFillMapResponse;
    message.limitOrderPoolFillMap = [];
    if (
      object.limitOrderPoolFillMap !== undefined &&
      object.limitOrderPoolFillMap !== null
    ) {
      for (const e of object.limitOrderPoolFillMap) {
        message.limitOrderPoolFillMap.push(
          LimitOrderPoolFillMap.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a TickMap by index. */
  TickMap(request: QueryGetTickMapRequest): Promise<QueryGetTickMapResponse>;
  /** Queries a list of TickMap items. */
  TickMapAll(request: QueryAllTickMapRequest): Promise<QueryAllTickMapResponse>;
  /** Queries a PairMap by index. */
  PairMap(request: QueryGetPairMapRequest): Promise<QueryGetPairMapResponse>;
  /** Queries a list of PairMap items. */
  PairMapAll(request: QueryAllPairMapRequest): Promise<QueryAllPairMapResponse>;
  /** Queries a Tokens by id. */
  Tokens(request: QueryGetTokensRequest): Promise<QueryGetTokensResponse>;
  /** Queries a list of Tokens items. */
  TokensAll(request: QueryAllTokensRequest): Promise<QueryAllTokensResponse>;
  /** Queries a TokenMap by index. */
  TokenMap(request: QueryGetTokenMapRequest): Promise<QueryGetTokenMapResponse>;
  /** Queries a list of TokenMap items. */
  TokenMapAll(
    request: QueryAllTokenMapRequest
  ): Promise<QueryAllTokenMapResponse>;
  /** Queries a Shares by index. */
  Shares(request: QueryGetSharesRequest): Promise<QueryGetSharesResponse>;
  /** Queries a list of Shares items. */
  SharesAll(request: QueryAllSharesRequest): Promise<QueryAllSharesResponse>;
  /** Queries a FeeList by id. */
  FeeList(request: QueryGetFeeListRequest): Promise<QueryGetFeeListResponse>;
  /** Queries a list of FeeList items. */
  FeeListAll(request: QueryAllFeeListRequest): Promise<QueryAllFeeListResponse>;
  /** Queries a EdgeRow by id. */
  EdgeRow(request: QueryGetEdgeRowRequest): Promise<QueryGetEdgeRowResponse>;
  /** Queries a list of EdgeRow items. */
  EdgeRowAll(request: QueryAllEdgeRowRequest): Promise<QueryAllEdgeRowResponse>;
  /** Queries a AdjanceyMatrix by id. */
  AdjanceyMatrix(
    request: QueryGetAdjanceyMatrixRequest
  ): Promise<QueryGetAdjanceyMatrixResponse>;
  /** Queries a list of AdjanceyMatrix items. */
  AdjanceyMatrixAll(
    request: QueryAllAdjanceyMatrixRequest
  ): Promise<QueryAllAdjanceyMatrixResponse>;
  /** Queries a LimitOrderPoolUserShareMap by index. */
  LimitOrderPoolUserShareMap(
    request: QueryGetLimitOrderPoolUserShareMapRequest
  ): Promise<QueryGetLimitOrderPoolUserShareMapResponse>;
  /** Queries a list of LimitOrderPoolUserShareMap items. */
  LimitOrderPoolUserShareMapAll(
    request: QueryAllLimitOrderPoolUserShareMapRequest
  ): Promise<QueryAllLimitOrderPoolUserShareMapResponse>;
  /** Queries a LimitOrderPoolUserSharesWithdrawn by index. */
  LimitOrderPoolUserSharesWithdrawn(
    request: QueryGetLimitOrderPoolUserSharesWithdrawnRequest
  ): Promise<QueryGetLimitOrderPoolUserSharesWithdrawnResponse>;
  /** Queries a list of LimitOrderPoolUserSharesWithdrawn items. */
  LimitOrderPoolUserSharesWithdrawnAll(
    request: QueryAllLimitOrderPoolUserSharesWithdrawnRequest
  ): Promise<QueryAllLimitOrderPoolUserSharesWithdrawnResponse>;
  /** Queries a LimitOrderPoolTotalSharesMap by index. */
  LimitOrderPoolTotalSharesMap(
    request: QueryGetLimitOrderPoolTotalSharesMapRequest
  ): Promise<QueryGetLimitOrderPoolTotalSharesMapResponse>;
  /** Queries a list of LimitOrderPoolTotalSharesMap items. */
  LimitOrderPoolTotalSharesMapAll(
    request: QueryAllLimitOrderPoolTotalSharesMapRequest
  ): Promise<QueryAllLimitOrderPoolTotalSharesMapResponse>;
  /** Queries a LimitOrderPoolReserveMap by index. */
  LimitOrderPoolReserveMap(
    request: QueryGetLimitOrderPoolReserveMapRequest
  ): Promise<QueryGetLimitOrderPoolReserveMapResponse>;
  /** Queries a list of LimitOrderPoolReserveMap items. */
  LimitOrderPoolReserveMapAll(
    request: QueryAllLimitOrderPoolReserveMapRequest
  ): Promise<QueryAllLimitOrderPoolReserveMapResponse>;
  /** Queries a LimitOrderPoolFillMap by index. */
  LimitOrderPoolFillMap(
    request: QueryGetLimitOrderPoolFillMapRequest
  ): Promise<QueryGetLimitOrderPoolFillMapResponse>;
  /** Queries a list of LimitOrderPoolFillMap items. */
  LimitOrderPoolFillMapAll(
    request: QueryAllLimitOrderPoolFillMapRequest
  ): Promise<QueryAllLimitOrderPoolFillMapResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  TickMap(request: QueryGetTickMapRequest): Promise<QueryGetTickMapResponse> {
    const data = QueryGetTickMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TickMap",
      data
    );
    return promise.then((data) =>
      QueryGetTickMapResponse.decode(new Reader(data))
    );
  }

  TickMapAll(
    request: QueryAllTickMapRequest
  ): Promise<QueryAllTickMapResponse> {
    const data = QueryAllTickMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TickMapAll",
      data
    );
    return promise.then((data) =>
      QueryAllTickMapResponse.decode(new Reader(data))
    );
  }

  PairMap(request: QueryGetPairMapRequest): Promise<QueryGetPairMapResponse> {
    const data = QueryGetPairMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "PairMap",
      data
    );
    return promise.then((data) =>
      QueryGetPairMapResponse.decode(new Reader(data))
    );
  }

  PairMapAll(
    request: QueryAllPairMapRequest
  ): Promise<QueryAllPairMapResponse> {
    const data = QueryAllPairMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "PairMapAll",
      data
    );
    return promise.then((data) =>
      QueryAllPairMapResponse.decode(new Reader(data))
    );
  }

  Tokens(request: QueryGetTokensRequest): Promise<QueryGetTokensResponse> {
    const data = QueryGetTokensRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "Tokens",
      data
    );
    return promise.then((data) =>
      QueryGetTokensResponse.decode(new Reader(data))
    );
  }

  TokensAll(request: QueryAllTokensRequest): Promise<QueryAllTokensResponse> {
    const data = QueryAllTokensRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TokensAll",
      data
    );
    return promise.then((data) =>
      QueryAllTokensResponse.decode(new Reader(data))
    );
  }

  TokenMap(
    request: QueryGetTokenMapRequest
  ): Promise<QueryGetTokenMapResponse> {
    const data = QueryGetTokenMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TokenMap",
      data
    );
    return promise.then((data) =>
      QueryGetTokenMapResponse.decode(new Reader(data))
    );
  }

  TokenMapAll(
    request: QueryAllTokenMapRequest
  ): Promise<QueryAllTokenMapResponse> {
    const data = QueryAllTokenMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TokenMapAll",
      data
    );
    return promise.then((data) =>
      QueryAllTokenMapResponse.decode(new Reader(data))
    );
  }

  Shares(request: QueryGetSharesRequest): Promise<QueryGetSharesResponse> {
    const data = QueryGetSharesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "Shares",
      data
    );
    return promise.then((data) =>
      QueryGetSharesResponse.decode(new Reader(data))
    );
  }

  SharesAll(request: QueryAllSharesRequest): Promise<QueryAllSharesResponse> {
    const data = QueryAllSharesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "SharesAll",
      data
    );
    return promise.then((data) =>
      QueryAllSharesResponse.decode(new Reader(data))
    );
  }

  FeeList(request: QueryGetFeeListRequest): Promise<QueryGetFeeListResponse> {
    const data = QueryGetFeeListRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "FeeList",
      data
    );
    return promise.then((data) =>
      QueryGetFeeListResponse.decode(new Reader(data))
    );
  }

  FeeListAll(
    request: QueryAllFeeListRequest
  ): Promise<QueryAllFeeListResponse> {
    const data = QueryAllFeeListRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "FeeListAll",
      data
    );
    return promise.then((data) =>
      QueryAllFeeListResponse.decode(new Reader(data))
    );
  }

  EdgeRow(request: QueryGetEdgeRowRequest): Promise<QueryGetEdgeRowResponse> {
    const data = QueryGetEdgeRowRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "EdgeRow",
      data
    );
    return promise.then((data) =>
      QueryGetEdgeRowResponse.decode(new Reader(data))
    );
  }

  EdgeRowAll(
    request: QueryAllEdgeRowRequest
  ): Promise<QueryAllEdgeRowResponse> {
    const data = QueryAllEdgeRowRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "EdgeRowAll",
      data
    );
    return promise.then((data) =>
      QueryAllEdgeRowResponse.decode(new Reader(data))
    );
  }

  AdjanceyMatrix(
    request: QueryGetAdjanceyMatrixRequest
  ): Promise<QueryGetAdjanceyMatrixResponse> {
    const data = QueryGetAdjanceyMatrixRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "AdjanceyMatrix",
      data
    );
    return promise.then((data) =>
      QueryGetAdjanceyMatrixResponse.decode(new Reader(data))
    );
  }

  AdjanceyMatrixAll(
    request: QueryAllAdjanceyMatrixRequest
  ): Promise<QueryAllAdjanceyMatrixResponse> {
    const data = QueryAllAdjanceyMatrixRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "AdjanceyMatrixAll",
      data
    );
    return promise.then((data) =>
      QueryAllAdjanceyMatrixResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolUserShareMap(
    request: QueryGetLimitOrderPoolUserShareMapRequest
  ): Promise<QueryGetLimitOrderPoolUserShareMapResponse> {
    const data = QueryGetLimitOrderPoolUserShareMapRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolUserShareMap",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolUserShareMapResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolUserShareMapAll(
    request: QueryAllLimitOrderPoolUserShareMapRequest
  ): Promise<QueryAllLimitOrderPoolUserShareMapResponse> {
    const data = QueryAllLimitOrderPoolUserShareMapRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolUserShareMapAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolUserShareMapResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolUserSharesWithdrawn(
    request: QueryGetLimitOrderPoolUserSharesWithdrawnRequest
  ): Promise<QueryGetLimitOrderPoolUserSharesWithdrawnResponse> {
    const data = QueryGetLimitOrderPoolUserSharesWithdrawnRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolUserSharesWithdrawn",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolUserSharesWithdrawnResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolUserSharesWithdrawnAll(
    request: QueryAllLimitOrderPoolUserSharesWithdrawnRequest
  ): Promise<QueryAllLimitOrderPoolUserSharesWithdrawnResponse> {
    const data = QueryAllLimitOrderPoolUserSharesWithdrawnRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolUserSharesWithdrawnAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolUserSharesWithdrawnResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolTotalSharesMap(
    request: QueryGetLimitOrderPoolTotalSharesMapRequest
  ): Promise<QueryGetLimitOrderPoolTotalSharesMapResponse> {
    const data = QueryGetLimitOrderPoolTotalSharesMapRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolTotalSharesMap",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolTotalSharesMapResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolTotalSharesMapAll(
    request: QueryAllLimitOrderPoolTotalSharesMapRequest
  ): Promise<QueryAllLimitOrderPoolTotalSharesMapResponse> {
    const data = QueryAllLimitOrderPoolTotalSharesMapRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolTotalSharesMapAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolTotalSharesMapResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolReserveMap(
    request: QueryGetLimitOrderPoolReserveMapRequest
  ): Promise<QueryGetLimitOrderPoolReserveMapResponse> {
    const data = QueryGetLimitOrderPoolReserveMapRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolReserveMap",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolReserveMapResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolReserveMapAll(
    request: QueryAllLimitOrderPoolReserveMapRequest
  ): Promise<QueryAllLimitOrderPoolReserveMapResponse> {
    const data = QueryAllLimitOrderPoolReserveMapRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolReserveMapAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolReserveMapResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolFillMap(
    request: QueryGetLimitOrderPoolFillMapRequest
  ): Promise<QueryGetLimitOrderPoolFillMapResponse> {
    const data = QueryGetLimitOrderPoolFillMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolFillMap",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolFillMapResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolFillMapAll(
    request: QueryAllLimitOrderPoolFillMapRequest
  ): Promise<QueryAllLimitOrderPoolFillMapResponse> {
    const data = QueryAllLimitOrderPoolFillMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolFillMapAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolFillMapResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
