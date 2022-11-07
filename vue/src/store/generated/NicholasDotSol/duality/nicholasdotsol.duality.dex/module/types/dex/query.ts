/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../dex/params";
import { TickObject } from "../dex/tick_map";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { PairMap } from "../dex/pair_map";
import { Tokens } from "../dex/tokens";
import { TokenObject } from "../dex/token_map";
import { Shares } from "../dex/shares";
import { FeeList } from "../dex/fee_list";
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

export interface QueryGetTickObjectRequest {
  tickIndex: number;
  pairId: string;
}

export interface QueryGetTickObjectResponse {
  tickObject: TickObject | undefined;
}

export interface QueryAllTickObjectRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTickObjectResponse {
  tickObject: TickObject[];
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

export interface QueryGetTokenObjectRequest {
  address: string;
}

export interface QueryGetTokenObjectResponse {
  tokenObject: TokenObject | undefined;
}

export interface QueryAllTokenObjectRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTokenObjectResponse {
  tokenObject: TokenObject[];
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

export interface QueryGetLimitOrderPoolUserShareMapRequest {
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
  address: string;
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
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
  address: string;
}

export interface QueryGetLimitOrderPoolUserSharesWithdrawnResponse {
  limitOrderPoolUserSharesWithdrawn:
    | LimitOrderPoolUserSharesWithdrawn
    | undefined;
}

export interface QueryAllLimitOrderPoolUserSharesWithdrawnRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolUserSharesWithdrawnResponse {
  limitOrderPoolUserSharesWithdrawn: LimitOrderPoolUserSharesWithdrawn[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolTotalSharesMapRequest {
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
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
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
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
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
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

const baseQueryGetTickObjectRequest: object = { tickIndex: 0, pairId: "" };

export const QueryGetTickObjectRequest = {
  encode(
    message: QueryGetTickObjectRequest,
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

  decode(input: Reader | Uint8Array, length?: number): QueryGetTickObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTickObjectRequest } as QueryGetTickObjectRequest;
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

  fromJSON(object: any): QueryGetTickObjectRequest {
    const message = { ...baseQueryGetTickObjectRequest } as QueryGetTickObjectRequest;
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

  toJSON(message: QueryGetTickObjectRequest): unknown {
    const obj: any = {};
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTickObjectRequest>
  ): QueryGetTickObjectRequest {
    const message = { ...baseQueryGetTickObjectRequest } as QueryGetTickObjectRequest;
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

const baseQueryGetTickObjectResponse: object = {};

export const QueryGetTickObjectResponse = {
  encode(
    message: QueryGetTickObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tickObject !== undefined) {
      TickObject.encode(message.tickObject, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTickObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTickObjectResponse,
    } as QueryGetTickObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tickObject = TickObject.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTickObjectResponse {
    const message = {
      ...baseQueryGetTickObjectResponse,
    } as QueryGetTickObjectResponse;
    if (object.tickObject !== undefined && object.tickObject !== null) {
      message.tickObject = TickObject.fromJSON(object.tickObject);
    } else {
      message.tickObject = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTickObjectResponse): unknown {
    const obj: any = {};
    message.tickObject !== undefined &&
      (obj.tickObject = message.tickObject
        ? TickObject.toJSON(message.tickObject)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTickObjectResponse>
  ): QueryGetTickObjectResponse {
    const message = {
      ...baseQueryGetTickObjectResponse,
    } as QueryGetTickObjectResponse;
    if (object.tickObject !== undefined && object.tickObject !== null) {
      message.tickObject = TickObject.fromPartial(object.tickObject);
    } else {
      message.tickObject = undefined;
    }
    return message;
  },
};

const baseQueryAllTickObjectRequest: object = {};

export const QueryAllTickObjectRequest = {
  encode(
    message: QueryAllTickObjectRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTickObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTickObjectRequest } as QueryAllTickObjectRequest;
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

  fromJSON(object: any): QueryAllTickObjectRequest {
    const message = { ...baseQueryAllTickObjectRequest } as QueryAllTickObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTickObjectRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTickObjectRequest>
  ): QueryAllTickObjectRequest {
    const message = { ...baseQueryAllTickObjectRequest } as QueryAllTickObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTickObjectResponse: object = {};

export const QueryAllTickObjectResponse = {
  encode(
    message: QueryAllTickObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.tickObject) {
      TickObject.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTickObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTickObjectResponse,
    } as QueryAllTickObjectResponse;
    message.tickObject = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tickObject.push(TickObject.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllTickObjectResponse {
    const message = {
      ...baseQueryAllTickObjectResponse,
    } as QueryAllTickObjectResponse;
    message.tickObject = [];
    if (object.tickObject !== undefined && object.tickObject !== null) {
      for (const e of object.tickObject) {
        message.tickObject.push(TickObject.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTickObjectResponse): unknown {
    const obj: any = {};
    if (message.tickObject) {
      obj.tickObject = message.tickObject.map((e) =>
        e ? TickObject.toJSON(e) : undefined
      );
    } else {
      obj.tickObject = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTickObjectResponse>
  ): QueryAllTickObjectResponse {
    const message = {
      ...baseQueryAllTickObjectResponse,
    } as QueryAllTickObjectResponse;
    message.tickObject = [];
    if (object.tickObject !== undefined && object.tickObject !== null) {
      for (const e of object.tickObject) {
        message.tickObject.push(TickObject.fromPartial(e));
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

const baseQueryGetTokenObjectRequest: object = { address: "" };

export const QueryGetTokenObjectRequest = {
  encode(
    message: QueryGetTokenObjectRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTokenObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTokenObjectRequest,
    } as QueryGetTokenObjectRequest;
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

  fromJSON(object: any): QueryGetTokenObjectRequest {
    const message = {
      ...baseQueryGetTokenObjectRequest,
    } as QueryGetTokenObjectRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetTokenObjectRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokenObjectRequest>
  ): QueryGetTokenObjectRequest {
    const message = {
      ...baseQueryGetTokenObjectRequest,
    } as QueryGetTokenObjectRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetTokenObjectResponse: object = {};

export const QueryGetTokenObjectResponse = {
  encode(
    message: QueryGetTokenObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tokenObject !== undefined) {
      TokenObject.encode(message.tokenObject, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTokenObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTokenObjectResponse,
    } as QueryGetTokenObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenObject = TokenObject.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokenObjectResponse {
    const message = {
      ...baseQueryGetTokenObjectResponse,
    } as QueryGetTokenObjectResponse;
    if (object.tokenObject !== undefined && object.tokenObject !== null) {
      message.tokenObject = TokenObject.fromJSON(object.tokenObject);
    } else {
      message.tokenObject = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTokenObjectResponse): unknown {
    const obj: any = {};
    message.tokenObject !== undefined &&
      (obj.tokenObject = message.tokenObject
        ? TokenObject.toJSON(message.tokenObject)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokenObjectResponse>
  ): QueryGetTokenObjectResponse {
    const message = {
      ...baseQueryGetTokenObjectResponse,
    } as QueryGetTokenObjectResponse;
    if (object.tokenObject !== undefined && object.tokenObject !== null) {
      message.tokenObject = TokenObject.fromPartial(object.tokenObject);
    } else {
      message.tokenObject = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenObjectRequest: object = {};

export const QueryAllTokenObjectRequest = {
  encode(
    message: QueryAllTokenObjectRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTokenObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTokenObjectRequest,
    } as QueryAllTokenObjectRequest;
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

  fromJSON(object: any): QueryAllTokenObjectRequest {
    const message = {
      ...baseQueryAllTokenObjectRequest,
    } as QueryAllTokenObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenObjectRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokenObjectRequest>
  ): QueryAllTokenObjectRequest {
    const message = {
      ...baseQueryAllTokenObjectRequest,
    } as QueryAllTokenObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenObjectResponse: object = {};

export const QueryAllTokenObjectResponse = {
  encode(
    message: QueryAllTokenObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.tokenObject) {
      TokenObject.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllTokenObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTokenObjectResponse,
    } as QueryAllTokenObjectResponse;
    message.tokenObject = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenObject.push(TokenObject.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllTokenObjectResponse {
    const message = {
      ...baseQueryAllTokenObjectResponse,
    } as QueryAllTokenObjectResponse;
    message.tokenObject = [];
    if (object.tokenObject !== undefined && object.tokenObject !== null) {
      for (const e of object.tokenObject) {
        message.tokenObject.push(TokenObject.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenObjectResponse): unknown {
    const obj: any = {};
    if (message.tokenObject) {
      obj.tokenObject = message.tokenObject.map((e) =>
        e ? TokenObject.toJSON(e) : undefined
      );
    } else {
      obj.tokenObject = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokenObjectResponse>
  ): QueryAllTokenObjectResponse {
    const message = {
      ...baseQueryAllTokenObjectResponse,
    } as QueryAllTokenObjectResponse;
    message.tokenObject = [];
    if (object.tokenObject !== undefined && object.tokenObject !== null) {
      for (const e of object.tokenObject) {
        message.tokenObject.push(TokenObject.fromPartial(e));
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

const baseQueryGetLimitOrderPoolUserShareMapRequest: object = {
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
  address: "",
};

export const QueryGetLimitOrderPoolUserShareMapRequest = {
  encode(
    message: QueryGetLimitOrderPoolUserShareMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(16).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.count !== 0) {
      writer.uint32(32).uint64(message.count);
    }
    if (message.address !== "") {
      writer.uint32(42).string(message.address);
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
          message.pairId = reader.string();
          break;
        case 2:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.address = reader.string();
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
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolUserShareMapRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserShareMapRequest>
  ): QueryGetLimitOrderPoolUserShareMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareMapRequest,
    } as QueryGetLimitOrderPoolUserShareMapRequest;
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
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
  address: "",
};

export const QueryGetLimitOrderPoolUserSharesWithdrawnRequest = {
  encode(
    message: QueryGetLimitOrderPoolUserSharesWithdrawnRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(16).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.count !== 0) {
      writer.uint32(32).uint64(message.count);
    }
    if (message.address !== "") {
      writer.uint32(42).string(message.address);
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
          message.pairId = reader.string();
          break;
        case 2:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.address = reader.string();
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
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolUserSharesWithdrawnRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserSharesWithdrawnRequest>
  ): QueryGetLimitOrderPoolUserSharesWithdrawnRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnRequest,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnRequest;
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
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
};

export const QueryGetLimitOrderPoolTotalSharesMapRequest = {
  encode(
    message: QueryGetLimitOrderPoolTotalSharesMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(16).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.count !== 0) {
      writer.uint32(32).uint64(message.count);
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
          message.pairId = reader.string();
          break;
        case 2:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.count = longToNumber(reader.uint64() as Long);
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
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolTotalSharesMapRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolTotalSharesMapRequest>
  ): QueryGetLimitOrderPoolTotalSharesMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesMapRequest,
    } as QueryGetLimitOrderPoolTotalSharesMapRequest;
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
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
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
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
};

export const QueryGetLimitOrderPoolReserveMapRequest = {
  encode(
    message: QueryGetLimitOrderPoolReserveMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(16).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.count !== 0) {
      writer.uint32(32).uint64(message.count);
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
          message.pairId = reader.string();
          break;
        case 2:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.count = longToNumber(reader.uint64() as Long);
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
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolReserveMapRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolReserveMapRequest>
  ): QueryGetLimitOrderPoolReserveMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveMapRequest,
    } as QueryGetLimitOrderPoolReserveMapRequest;
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
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
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
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
};

export const QueryGetLimitOrderPoolFillMapRequest = {
  encode(
    message: QueryGetLimitOrderPoolFillMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(16).int64(message.tickIndex);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.count !== 0) {
      writer.uint32(32).uint64(message.count);
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
          message.pairId = reader.string();
          break;
        case 2:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.count = longToNumber(reader.uint64() as Long);
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
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolFillMapRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolFillMapRequest>
  ): QueryGetLimitOrderPoolFillMapRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolFillMapRequest,
    } as QueryGetLimitOrderPoolFillMapRequest;
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
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
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
  /** Queries a TickObject by index. */
  TickObject(request: QueryGetTickObjectRequest): Promise<QueryGetTickObjectResponse>;
  /** Queries a list of TickObject items. */
  TickObjectAll(request: QueryAllTickObjectRequest): Promise<QueryAllTickObjectResponse>;
  /** Queries a PairMap by index. */
  PairMap(request: QueryGetPairMapRequest): Promise<QueryGetPairMapResponse>;
  /** Queries a list of PairMap items. */
  PairMapAll(request: QueryAllPairMapRequest): Promise<QueryAllPairMapResponse>;
  /** Queries a Tokens by id. */
  Tokens(request: QueryGetTokensRequest): Promise<QueryGetTokensResponse>;
  /** Queries a list of Tokens items. */
  TokensAll(request: QueryAllTokensRequest): Promise<QueryAllTokensResponse>;
  /** Queries a TokenObject by index. */
  TokenObject(request: QueryGetTokenObjectRequest): Promise<QueryGetTokenObjectResponse>;
  /** Queries a list of TokenObject items. */
  TokenObjectAll(
    request: QueryAllTokenObjectRequest
  ): Promise<QueryAllTokenObjectResponse>;
  /** Queries a Shares by index. */
  Shares(request: QueryGetSharesRequest): Promise<QueryGetSharesResponse>;
  /** Queries a list of Shares items. */
  SharesAll(request: QueryAllSharesRequest): Promise<QueryAllSharesResponse>;
  /** Queries a FeeList by id. */
  FeeList(request: QueryGetFeeListRequest): Promise<QueryGetFeeListResponse>;
  /** Queries a list of FeeList items. */
  FeeListAll(request: QueryAllFeeListRequest): Promise<QueryAllFeeListResponse>;
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

  TickObject(request: QueryGetTickObjectRequest): Promise<QueryGetTickObjectResponse> {
    const data = QueryGetTickObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TickObject",
      data
    );
    return promise.then((data) =>
      QueryGetTickObjectResponse.decode(new Reader(data))
    );
  }

  TickObjectAll(
    request: QueryAllTickObjectRequest
  ): Promise<QueryAllTickObjectResponse> {
    const data = QueryAllTickObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TickObjectAll",
      data
    );
    return promise.then((data) =>
      QueryAllTickObjectResponse.decode(new Reader(data))
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

  TokenObject(
    request: QueryGetTokenObjectRequest
  ): Promise<QueryGetTokenObjectResponse> {
    const data = QueryGetTokenObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TokenObject",
      data
    );
    return promise.then((data) =>
      QueryGetTokenObjectResponse.decode(new Reader(data))
    );
  }

  TokenObjectAll(
    request: QueryAllTokenObjectRequest
  ): Promise<QueryAllTokenObjectResponse> {
    const data = QueryAllTokenObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TokenObjectAll",
      data
    );
    return promise.then((data) =>
      QueryAllTokenObjectResponse.decode(new Reader(data))
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
