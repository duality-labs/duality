/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../dex/params";
import { TickObject } from "../dex/tick_map";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { PairObject } from "../dex/pair_map";
import { Tokens } from "../dex/tokens";
import { TokenObject } from "../dex/token_map";
import { Shares } from "../dex/shares";
import { FeeList } from "../dex/fee_list";
import { LimitOrderPoolUserShareObject } from "../dex/limit_order_pool_user_share_map";
import { LimitOrderPoolUserSharesWithdrawnObject } from "../dex/limit_order_pool_user_shares_withdrawn";
import { LimitOrderPoolTotalSharesObject } from "../dex/limit_order_pool_total_shares_map";
import { LimitOrderPoolReserveObject } from "../dex/limit_order_pool_reserve_map";
import { LimitOrderPoolFillObject } from "../dex/limit_order_pool_fill_map";

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

export interface QueryGetPairObjectRequest {
  pairId: string;
}

export interface QueryGetPairObjectResponse {
  pairObject: PairObject | undefined;
}

export interface QueryAllPairObjectRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPairObjectResponse {
  pairObject: PairObject[];
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

export interface QueryGetLimitOrderPoolUserShareObjectRequest {
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
  address: string;
}

export interface QueryGetLimitOrderPoolUserShareObjectResponse {
  limitOrderPoolUserShareObject: LimitOrderPoolUserShareObject | undefined;
}

export interface QueryAllLimitOrderPoolUserShareObjectRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolUserShareObjectResponse {
  limitOrderPoolUserShareObject: LimitOrderPoolUserShareObject[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest {
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
  address: string;
}

export interface QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse {
  limitOrderPoolUserSharesWithdrawnObject:
    | LimitOrderPoolUserSharesWithdrawnObject
    | undefined;
}

export interface QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse {
  limitOrderPoolUserSharesWithdrawnObject: LimitOrderPoolUserSharesWithdrawnObject[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolTotalSharesObjectRequest {
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
}

export interface QueryGetLimitOrderPoolTotalSharesObjectResponse {
  limitOrderPoolTotalSharesObject: LimitOrderPoolTotalSharesObject | undefined;
}

export interface QueryAllLimitOrderPoolTotalSharesObjectRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolTotalSharesObjectResponse {
  limitOrderPoolTotalSharesObject: LimitOrderPoolTotalSharesObject[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolReserveObjectRequest {
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
}

export interface QueryGetLimitOrderPoolReserveObjectResponse {
  limitOrderPoolReserveObject: LimitOrderPoolReserveObject | undefined;
}

export interface QueryAllLimitOrderPoolReserveObjectRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolReserveObjectResponse {
  limitOrderPoolReserveObject: LimitOrderPoolReserveObject[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLimitOrderPoolFillObjectRequest {
  pairId: string;
  tickIndex: number;
  token: string;
  count: number;
}

export interface QueryGetLimitOrderPoolFillObjectResponse {
  limitOrderPoolFillObject: LimitOrderPoolFillObject | undefined;
}

export interface QueryAllLimitOrderPoolFillObjectRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLimitOrderPoolFillObjectResponse {
  limitOrderPoolFillObject: LimitOrderPoolFillObject[];
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

const baseQueryGetPairObjectRequest: object = { pairId: "" };

export const QueryGetPairObjectRequest = {
  encode(
    message: QueryGetPairObjectRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPairObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPairObjectRequest } as QueryGetPairObjectRequest;
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

  fromJSON(object: any): QueryGetPairObjectRequest {
    const message = { ...baseQueryGetPairObjectRequest } as QueryGetPairObjectRequest;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    return message;
  },

  toJSON(message: QueryGetPairObjectRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPairObjectRequest>
  ): QueryGetPairObjectRequest {
    const message = { ...baseQueryGetPairObjectRequest } as QueryGetPairObjectRequest;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    return message;
  },
};

const baseQueryGetPairObjectResponse: object = {};

export const QueryGetPairObjectResponse = {
  encode(
    message: QueryGetPairObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairObject !== undefined) {
      PairObject.encode(message.pairObject, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPairObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPairObjectResponse,
    } as QueryGetPairObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairObject = PairObject.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPairObjectResponse {
    const message = {
      ...baseQueryGetPairObjectResponse,
    } as QueryGetPairObjectResponse;
    if (object.pairObject !== undefined && object.pairObject !== null) {
      message.pairObject = PairObject.fromJSON(object.pairObject);
    } else {
      message.pairObject = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPairObjectResponse): unknown {
    const obj: any = {};
    message.pairObject !== undefined &&
      (obj.pairObject = message.pairObject
        ? PairObject.toJSON(message.pairObject)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPairObjectResponse>
  ): QueryGetPairObjectResponse {
    const message = {
      ...baseQueryGetPairObjectResponse,
    } as QueryGetPairObjectResponse;
    if (object.pairObject !== undefined && object.pairObject !== null) {
      message.pairObject = PairObject.fromPartial(object.pairObject);
    } else {
      message.pairObject = undefined;
    }
    return message;
  },
};

const baseQueryAllPairObjectRequest: object = {};

export const QueryAllPairObjectRequest = {
  encode(
    message: QueryAllPairObjectRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPairObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPairObjectRequest } as QueryAllPairObjectRequest;
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

  fromJSON(object: any): QueryAllPairObjectRequest {
    const message = { ...baseQueryAllPairObjectRequest } as QueryAllPairObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPairObjectRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPairObjectRequest>
  ): QueryAllPairObjectRequest {
    const message = { ...baseQueryAllPairObjectRequest } as QueryAllPairObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPairObjectResponse: object = {};

export const QueryAllPairObjectResponse = {
  encode(
    message: QueryAllPairObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.pairObject) {
      PairObject.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPairObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPairObjectResponse,
    } as QueryAllPairObjectResponse;
    message.pairObject = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairObject.push(PairObject.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllPairObjectResponse {
    const message = {
      ...baseQueryAllPairObjectResponse,
    } as QueryAllPairObjectResponse;
    message.pairObject = [];
    if (object.pairObject !== undefined && object.pairObject !== null) {
      for (const e of object.pairObject) {
        message.pairObject.push(PairObject.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPairObjectResponse): unknown {
    const obj: any = {};
    if (message.pairObject) {
      obj.pairObject = message.pairObject.map((e) =>
        e ? PairObject.toJSON(e) : undefined
      );
    } else {
      obj.pairObject = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPairObjectResponse>
  ): QueryAllPairObjectResponse {
    const message = {
      ...baseQueryAllPairObjectResponse,
    } as QueryAllPairObjectResponse;
    message.pairObject = [];
    if (object.pairObject !== undefined && object.pairObject !== null) {
      for (const e of object.pairObject) {
        message.pairObject.push(PairObject.fromPartial(e));
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

const baseQueryGetLimitOrderPoolUserShareObjectRequest: object = {
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
  address: "",
};

export const QueryGetLimitOrderPoolUserShareObjectRequest = {
  encode(
    message: QueryGetLimitOrderPoolUserShareObjectRequest,
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
  ): QueryGetLimitOrderPoolUserShareObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareObjectRequest,
    } as QueryGetLimitOrderPoolUserShareObjectRequest;
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

  fromJSON(object: any): QueryGetLimitOrderPoolUserShareObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareObjectRequest,
    } as QueryGetLimitOrderPoolUserShareObjectRequest;
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

  toJSON(message: QueryGetLimitOrderPoolUserShareObjectRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserShareObjectRequest>
  ): QueryGetLimitOrderPoolUserShareObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareObjectRequest,
    } as QueryGetLimitOrderPoolUserShareObjectRequest;
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

const baseQueryGetLimitOrderPoolUserShareObjectResponse: object = {};

export const QueryGetLimitOrderPoolUserShareObjectResponse = {
  encode(
    message: QueryGetLimitOrderPoolUserShareObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolUserShareObject !== undefined) {
      LimitOrderPoolUserShareObject.encode(
        message.limitOrderPoolUserShareObject,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolUserShareObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareObjectResponse,
    } as QueryGetLimitOrderPoolUserShareObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolUserShareObject = LimitOrderPoolUserShareObject.decode(
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

  fromJSON(object: any): QueryGetLimitOrderPoolUserShareObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareObjectResponse,
    } as QueryGetLimitOrderPoolUserShareObjectResponse;
    if (
      object.limitOrderPoolUserShareObject !== undefined &&
      object.limitOrderPoolUserShareObject !== null
    ) {
      message.limitOrderPoolUserShareObject = LimitOrderPoolUserShareObject.fromJSON(
        object.limitOrderPoolUserShareObject
      );
    } else {
      message.limitOrderPoolUserShareObject = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolUserShareObjectResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolUserShareObject !== undefined &&
      (obj.limitOrderPoolUserShareObject = message.limitOrderPoolUserShareObject
        ? LimitOrderPoolUserShareObject.toJSON(message.limitOrderPoolUserShareObject)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserShareObjectResponse>
  ): QueryGetLimitOrderPoolUserShareObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolUserShareObjectResponse,
    } as QueryGetLimitOrderPoolUserShareObjectResponse;
    if (
      object.limitOrderPoolUserShareObject !== undefined &&
      object.limitOrderPoolUserShareObject !== null
    ) {
      message.limitOrderPoolUserShareObject = LimitOrderPoolUserShareObject.fromPartial(
        object.limitOrderPoolUserShareObject
      );
    } else {
      message.limitOrderPoolUserShareObject = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolUserShareObjectRequest: object = {};

export const QueryAllLimitOrderPoolUserShareObjectRequest = {
  encode(
    message: QueryAllLimitOrderPoolUserShareObjectRequest,
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
  ): QueryAllLimitOrderPoolUserShareObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareObjectRequest,
    } as QueryAllLimitOrderPoolUserShareObjectRequest;
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

  fromJSON(object: any): QueryAllLimitOrderPoolUserShareObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareObjectRequest,
    } as QueryAllLimitOrderPoolUserShareObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolUserShareObjectRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolUserShareObjectRequest>
  ): QueryAllLimitOrderPoolUserShareObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareObjectRequest,
    } as QueryAllLimitOrderPoolUserShareObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolUserShareObjectResponse: object = {};

export const QueryAllLimitOrderPoolUserShareObjectResponse = {
  encode(
    message: QueryAllLimitOrderPoolUserShareObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolUserShareObject) {
      LimitOrderPoolUserShareObject.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllLimitOrderPoolUserShareObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareObjectResponse,
    } as QueryAllLimitOrderPoolUserShareObjectResponse;
    message.limitOrderPoolUserShareObject = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolUserShareObject.push(
            LimitOrderPoolUserShareObject.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllLimitOrderPoolUserShareObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareObjectResponse,
    } as QueryAllLimitOrderPoolUserShareObjectResponse;
    message.limitOrderPoolUserShareObject = [];
    if (
      object.limitOrderPoolUserShareObject !== undefined &&
      object.limitOrderPoolUserShareObject !== null
    ) {
      for (const e of object.limitOrderPoolUserShareObject) {
        message.limitOrderPoolUserShareObject.push(
          LimitOrderPoolUserShareObject.fromJSON(e)
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

  toJSON(message: QueryAllLimitOrderPoolUserShareObjectResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolUserShareObject) {
      obj.limitOrderPoolUserShareObject = message.limitOrderPoolUserShareObject.map(
        (e) => (e ? LimitOrderPoolUserShareObject.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolUserShareObject = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolUserShareObjectResponse>
  ): QueryAllLimitOrderPoolUserShareObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolUserShareObjectResponse,
    } as QueryAllLimitOrderPoolUserShareObjectResponse;
    message.limitOrderPoolUserShareObject = [];
    if (
      object.limitOrderPoolUserShareObject !== undefined &&
      object.limitOrderPoolUserShareObject !== null
    ) {
      for (const e of object.limitOrderPoolUserShareObject) {
        message.limitOrderPoolUserShareObject.push(
          LimitOrderPoolUserShareObject.fromPartial(e)
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

const baseQueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest: object = {
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
  address: "",
};

export const QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest = {
  encode(
    message: QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest,
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
  ): QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest;
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

  fromJSON(object: any): QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest;
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

  toJSON(message: QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest>
  ): QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest;
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

const baseQueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse: object = {};

export const QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse = {
  encode(
    message: QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolUserSharesWithdrawnObject !== undefined) {
      LimitOrderPoolUserSharesWithdrawnObject.encode(
        message.limitOrderPoolUserSharesWithdrawnObject,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolUserSharesWithdrawnObject = LimitOrderPoolUserSharesWithdrawnObject.decode(
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

  fromJSON(object: any): QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse;
    if (
      object.limitOrderPoolUserSharesWithdrawnObject !== undefined &&
      object.limitOrderPoolUserSharesWithdrawnObject !== null
    ) {
      message.limitOrderPoolUserSharesWithdrawnObject = LimitOrderPoolUserSharesWithdrawnObject.fromJSON(
        object.limitOrderPoolUserSharesWithdrawnObject
      );
    } else {
      message.limitOrderPoolUserSharesWithdrawnObject = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolUserSharesWithdrawnObject !== undefined &&
      (obj.limitOrderPoolUserSharesWithdrawnObject = message.limitOrderPoolUserSharesWithdrawnObject
        ? LimitOrderPoolUserSharesWithdrawnObject.toJSON(
            message.limitOrderPoolUserSharesWithdrawnObject
          )
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse>
  ): QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse,
    } as QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse;
    if (
      object.limitOrderPoolUserSharesWithdrawnObject !== undefined &&
      object.limitOrderPoolUserSharesWithdrawnObject !== null
    ) {
      message.limitOrderPoolUserSharesWithdrawnObject = LimitOrderPoolUserSharesWithdrawnObject.fromPartial(
        object.limitOrderPoolUserSharesWithdrawnObject
      );
    } else {
      message.limitOrderPoolUserSharesWithdrawnObject = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest: object = {};

export const QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest = {
  encode(
    message: QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest,
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
  ): QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest;
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

  fromJSON(object: any): QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest>
  ): QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse: object = {};

export const QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse = {
  encode(
    message: QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolUserSharesWithdrawnObject) {
      LimitOrderPoolUserSharesWithdrawnObject.encode(
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
  ): QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse;
    message.limitOrderPoolUserSharesWithdrawnObject = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolUserSharesWithdrawnObject.push(
            LimitOrderPoolUserSharesWithdrawnObject.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse;
    message.limitOrderPoolUserSharesWithdrawnObject = [];
    if (
      object.limitOrderPoolUserSharesWithdrawnObject !== undefined &&
      object.limitOrderPoolUserSharesWithdrawnObject !== null
    ) {
      for (const e of object.limitOrderPoolUserSharesWithdrawnObject) {
        message.limitOrderPoolUserSharesWithdrawnObject.push(
          LimitOrderPoolUserSharesWithdrawnObject.fromJSON(e)
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

  toJSON(message: QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolUserSharesWithdrawnObject) {
      obj.limitOrderPoolUserSharesWithdrawnObject = message.limitOrderPoolUserSharesWithdrawnObject.map(
        (e) => (e ? LimitOrderPoolUserSharesWithdrawnObject.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolUserSharesWithdrawnObject = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse>
  ): QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse,
    } as QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse;
    message.limitOrderPoolUserSharesWithdrawnObject = [];
    if (
      object.limitOrderPoolUserSharesWithdrawnObject !== undefined &&
      object.limitOrderPoolUserSharesWithdrawnObject !== null
    ) {
      for (const e of object.limitOrderPoolUserSharesWithdrawnObject) {
        message.limitOrderPoolUserSharesWithdrawnObject.push(
          LimitOrderPoolUserSharesWithdrawnObject.fromPartial(e)
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

const baseQueryGetLimitOrderPoolTotalSharesObjectRequest: object = {
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
};

export const QueryGetLimitOrderPoolTotalSharesObjectRequest = {
  encode(
    message: QueryGetLimitOrderPoolTotalSharesObjectRequest,
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
  ): QueryGetLimitOrderPoolTotalSharesObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesObjectRequest,
    } as QueryGetLimitOrderPoolTotalSharesObjectRequest;
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

  fromJSON(object: any): QueryGetLimitOrderPoolTotalSharesObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesObjectRequest,
    } as QueryGetLimitOrderPoolTotalSharesObjectRequest;
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

  toJSON(message: QueryGetLimitOrderPoolTotalSharesObjectRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolTotalSharesObjectRequest>
  ): QueryGetLimitOrderPoolTotalSharesObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesObjectRequest,
    } as QueryGetLimitOrderPoolTotalSharesObjectRequest;
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

const baseQueryGetLimitOrderPoolTotalSharesObjectResponse: object = {};

export const QueryGetLimitOrderPoolTotalSharesObjectResponse = {
  encode(
    message: QueryGetLimitOrderPoolTotalSharesObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolTotalSharesObject !== undefined) {
      LimitOrderPoolTotalSharesObject.encode(
        message.limitOrderPoolTotalSharesObject,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolTotalSharesObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesObjectResponse,
    } as QueryGetLimitOrderPoolTotalSharesObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolTotalSharesObject = LimitOrderPoolTotalSharesObject.decode(
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

  fromJSON(object: any): QueryGetLimitOrderPoolTotalSharesObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesObjectResponse,
    } as QueryGetLimitOrderPoolTotalSharesObjectResponse;
    if (
      object.limitOrderPoolTotalSharesObject !== undefined &&
      object.limitOrderPoolTotalSharesObject !== null
    ) {
      message.limitOrderPoolTotalSharesObject = LimitOrderPoolTotalSharesObject.fromJSON(
        object.limitOrderPoolTotalSharesObject
      );
    } else {
      message.limitOrderPoolTotalSharesObject = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolTotalSharesObjectResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolTotalSharesObject !== undefined &&
      (obj.limitOrderPoolTotalSharesObject = message.limitOrderPoolTotalSharesObject
        ? LimitOrderPoolTotalSharesObject.toJSON(
            message.limitOrderPoolTotalSharesObject
          )
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolTotalSharesObjectResponse>
  ): QueryGetLimitOrderPoolTotalSharesObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolTotalSharesObjectResponse,
    } as QueryGetLimitOrderPoolTotalSharesObjectResponse;
    if (
      object.limitOrderPoolTotalSharesObject !== undefined &&
      object.limitOrderPoolTotalSharesObject !== null
    ) {
      message.limitOrderPoolTotalSharesObject = LimitOrderPoolTotalSharesObject.fromPartial(
        object.limitOrderPoolTotalSharesObject
      );
    } else {
      message.limitOrderPoolTotalSharesObject = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolTotalSharesObjectRequest: object = {};

export const QueryAllLimitOrderPoolTotalSharesObjectRequest = {
  encode(
    message: QueryAllLimitOrderPoolTotalSharesObjectRequest,
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
  ): QueryAllLimitOrderPoolTotalSharesObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesObjectRequest,
    } as QueryAllLimitOrderPoolTotalSharesObjectRequest;
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

  fromJSON(object: any): QueryAllLimitOrderPoolTotalSharesObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesObjectRequest,
    } as QueryAllLimitOrderPoolTotalSharesObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolTotalSharesObjectRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolTotalSharesObjectRequest>
  ): QueryAllLimitOrderPoolTotalSharesObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesObjectRequest,
    } as QueryAllLimitOrderPoolTotalSharesObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolTotalSharesObjectResponse: object = {};

export const QueryAllLimitOrderPoolTotalSharesObjectResponse = {
  encode(
    message: QueryAllLimitOrderPoolTotalSharesObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolTotalSharesObject) {
      LimitOrderPoolTotalSharesObject.encode(
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
  ): QueryAllLimitOrderPoolTotalSharesObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesObjectResponse,
    } as QueryAllLimitOrderPoolTotalSharesObjectResponse;
    message.limitOrderPoolTotalSharesObject = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolTotalSharesObject.push(
            LimitOrderPoolTotalSharesObject.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllLimitOrderPoolTotalSharesObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesObjectResponse,
    } as QueryAllLimitOrderPoolTotalSharesObjectResponse;
    message.limitOrderPoolTotalSharesObject = [];
    if (
      object.limitOrderPoolTotalSharesObject !== undefined &&
      object.limitOrderPoolTotalSharesObject !== null
    ) {
      for (const e of object.limitOrderPoolTotalSharesObject) {
        message.limitOrderPoolTotalSharesObject.push(
          LimitOrderPoolTotalSharesObject.fromJSON(e)
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

  toJSON(message: QueryAllLimitOrderPoolTotalSharesObjectResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolTotalSharesObject) {
      obj.limitOrderPoolTotalSharesObject = message.limitOrderPoolTotalSharesObject.map(
        (e) => (e ? LimitOrderPoolTotalSharesObject.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolTotalSharesObject = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolTotalSharesObjectResponse>
  ): QueryAllLimitOrderPoolTotalSharesObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolTotalSharesObjectResponse,
    } as QueryAllLimitOrderPoolTotalSharesObjectResponse;
    message.limitOrderPoolTotalSharesObject = [];
    if (
      object.limitOrderPoolTotalSharesObject !== undefined &&
      object.limitOrderPoolTotalSharesObject !== null
    ) {
      for (const e of object.limitOrderPoolTotalSharesObject) {
        message.limitOrderPoolTotalSharesObject.push(
          LimitOrderPoolTotalSharesObject.fromPartial(e)
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

const baseQueryGetLimitOrderPoolReserveObjectRequest: object = {
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
};

export const QueryGetLimitOrderPoolReserveObjectRequest = {
  encode(
    message: QueryGetLimitOrderPoolReserveObjectRequest,
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
  ): QueryGetLimitOrderPoolReserveObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolReserveObjectRequest,
    } as QueryGetLimitOrderPoolReserveObjectRequest;
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

  fromJSON(object: any): QueryGetLimitOrderPoolReserveObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveObjectRequest,
    } as QueryGetLimitOrderPoolReserveObjectRequest;
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

  toJSON(message: QueryGetLimitOrderPoolReserveObjectRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolReserveObjectRequest>
  ): QueryGetLimitOrderPoolReserveObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveObjectRequest,
    } as QueryGetLimitOrderPoolReserveObjectRequest;
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

const baseQueryGetLimitOrderPoolReserveObjectResponse: object = {};

export const QueryGetLimitOrderPoolReserveObjectResponse = {
  encode(
    message: QueryGetLimitOrderPoolReserveObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolReserveObject !== undefined) {
      LimitOrderPoolReserveObject.encode(
        message.limitOrderPoolReserveObject,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolReserveObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolReserveObjectResponse,
    } as QueryGetLimitOrderPoolReserveObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolReserveObject = LimitOrderPoolReserveObject.decode(
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

  fromJSON(object: any): QueryGetLimitOrderPoolReserveObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveObjectResponse,
    } as QueryGetLimitOrderPoolReserveObjectResponse;
    if (
      object.limitOrderPoolReserveObject !== undefined &&
      object.limitOrderPoolReserveObject !== null
    ) {
      message.limitOrderPoolReserveObject = LimitOrderPoolReserveObject.fromJSON(
        object.limitOrderPoolReserveObject
      );
    } else {
      message.limitOrderPoolReserveObject = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolReserveObjectResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolReserveObject !== undefined &&
      (obj.limitOrderPoolReserveObject = message.limitOrderPoolReserveObject
        ? LimitOrderPoolReserveObject.toJSON(message.limitOrderPoolReserveObject)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolReserveObjectResponse>
  ): QueryGetLimitOrderPoolReserveObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolReserveObjectResponse,
    } as QueryGetLimitOrderPoolReserveObjectResponse;
    if (
      object.limitOrderPoolReserveObject !== undefined &&
      object.limitOrderPoolReserveObject !== null
    ) {
      message.limitOrderPoolReserveObject = LimitOrderPoolReserveObject.fromPartial(
        object.limitOrderPoolReserveObject
      );
    } else {
      message.limitOrderPoolReserveObject = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolReserveObjectRequest: object = {};

export const QueryAllLimitOrderPoolReserveObjectRequest = {
  encode(
    message: QueryAllLimitOrderPoolReserveObjectRequest,
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
  ): QueryAllLimitOrderPoolReserveObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolReserveObjectRequest,
    } as QueryAllLimitOrderPoolReserveObjectRequest;
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

  fromJSON(object: any): QueryAllLimitOrderPoolReserveObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolReserveObjectRequest,
    } as QueryAllLimitOrderPoolReserveObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolReserveObjectRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolReserveObjectRequest>
  ): QueryAllLimitOrderPoolReserveObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolReserveObjectRequest,
    } as QueryAllLimitOrderPoolReserveObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolReserveObjectResponse: object = {};

export const QueryAllLimitOrderPoolReserveObjectResponse = {
  encode(
    message: QueryAllLimitOrderPoolReserveObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolReserveObject) {
      LimitOrderPoolReserveObject.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllLimitOrderPoolReserveObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolReserveObjectResponse,
    } as QueryAllLimitOrderPoolReserveObjectResponse;
    message.limitOrderPoolReserveObject = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolReserveObject.push(
            LimitOrderPoolReserveObject.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllLimitOrderPoolReserveObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolReserveObjectResponse,
    } as QueryAllLimitOrderPoolReserveObjectResponse;
    message.limitOrderPoolReserveObject = [];
    if (
      object.limitOrderPoolReserveObject !== undefined &&
      object.limitOrderPoolReserveObject !== null
    ) {
      for (const e of object.limitOrderPoolReserveObject) {
        message.limitOrderPoolReserveObject.push(
          LimitOrderPoolReserveObject.fromJSON(e)
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

  toJSON(message: QueryAllLimitOrderPoolReserveObjectResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolReserveObject) {
      obj.limitOrderPoolReserveObject = message.limitOrderPoolReserveObject.map((e) =>
        e ? LimitOrderPoolReserveObject.toJSON(e) : undefined
      );
    } else {
      obj.limitOrderPoolReserveObject = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolReserveObjectResponse>
  ): QueryAllLimitOrderPoolReserveObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolReserveObjectResponse,
    } as QueryAllLimitOrderPoolReserveObjectResponse;
    message.limitOrderPoolReserveObject = [];
    if (
      object.limitOrderPoolReserveObject !== undefined &&
      object.limitOrderPoolReserveObject !== null
    ) {
      for (const e of object.limitOrderPoolReserveObject) {
        message.limitOrderPoolReserveObject.push(
          LimitOrderPoolReserveObject.fromPartial(e)
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

const baseQueryGetLimitOrderPoolFillObjectRequest: object = {
  pairId: "",
  tickIndex: 0,
  token: "",
  count: 0,
};

export const QueryGetLimitOrderPoolFillObjectRequest = {
  encode(
    message: QueryGetLimitOrderPoolFillObjectRequest,
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
  ): QueryGetLimitOrderPoolFillObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolFillObjectRequest,
    } as QueryGetLimitOrderPoolFillObjectRequest;
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

  fromJSON(object: any): QueryGetLimitOrderPoolFillObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolFillObjectRequest,
    } as QueryGetLimitOrderPoolFillObjectRequest;
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

  toJSON(message: QueryGetLimitOrderPoolFillObjectRequest): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.token !== undefined && (obj.token = message.token);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolFillObjectRequest>
  ): QueryGetLimitOrderPoolFillObjectRequest {
    const message = {
      ...baseQueryGetLimitOrderPoolFillObjectRequest,
    } as QueryGetLimitOrderPoolFillObjectRequest;
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

const baseQueryGetLimitOrderPoolFillObjectResponse: object = {};

export const QueryGetLimitOrderPoolFillObjectResponse = {
  encode(
    message: QueryGetLimitOrderPoolFillObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.limitOrderPoolFillObject !== undefined) {
      LimitOrderPoolFillObject.encode(
        message.limitOrderPoolFillObject,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLimitOrderPoolFillObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLimitOrderPoolFillObjectResponse,
    } as QueryGetLimitOrderPoolFillObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolFillObject = LimitOrderPoolFillObject.decode(
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

  fromJSON(object: any): QueryGetLimitOrderPoolFillObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolFillObjectResponse,
    } as QueryGetLimitOrderPoolFillObjectResponse;
    if (
      object.limitOrderPoolFillObject !== undefined &&
      object.limitOrderPoolFillObject !== null
    ) {
      message.limitOrderPoolFillObject = LimitOrderPoolFillObject.fromJSON(
        object.limitOrderPoolFillObject
      );
    } else {
      message.limitOrderPoolFillObject = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLimitOrderPoolFillObjectResponse): unknown {
    const obj: any = {};
    message.limitOrderPoolFillObject !== undefined &&
      (obj.limitOrderPoolFillObject = message.limitOrderPoolFillObject
        ? LimitOrderPoolFillObject.toJSON(message.limitOrderPoolFillObject)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLimitOrderPoolFillObjectResponse>
  ): QueryGetLimitOrderPoolFillObjectResponse {
    const message = {
      ...baseQueryGetLimitOrderPoolFillObjectResponse,
    } as QueryGetLimitOrderPoolFillObjectResponse;
    if (
      object.limitOrderPoolFillObject !== undefined &&
      object.limitOrderPoolFillObject !== null
    ) {
      message.limitOrderPoolFillObject = LimitOrderPoolFillObject.fromPartial(
        object.limitOrderPoolFillObject
      );
    } else {
      message.limitOrderPoolFillObject = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolFillObjectRequest: object = {};

export const QueryAllLimitOrderPoolFillObjectRequest = {
  encode(
    message: QueryAllLimitOrderPoolFillObjectRequest,
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
  ): QueryAllLimitOrderPoolFillObjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolFillObjectRequest,
    } as QueryAllLimitOrderPoolFillObjectRequest;
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

  fromJSON(object: any): QueryAllLimitOrderPoolFillObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolFillObjectRequest,
    } as QueryAllLimitOrderPoolFillObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolFillObjectRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolFillObjectRequest>
  ): QueryAllLimitOrderPoolFillObjectRequest {
    const message = {
      ...baseQueryAllLimitOrderPoolFillObjectRequest,
    } as QueryAllLimitOrderPoolFillObjectRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLimitOrderPoolFillObjectResponse: object = {};

export const QueryAllLimitOrderPoolFillObjectResponse = {
  encode(
    message: QueryAllLimitOrderPoolFillObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.limitOrderPoolFillObject) {
      LimitOrderPoolFillObject.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllLimitOrderPoolFillObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLimitOrderPoolFillObjectResponse,
    } as QueryAllLimitOrderPoolFillObjectResponse;
    message.limitOrderPoolFillObject = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limitOrderPoolFillObject.push(
            LimitOrderPoolFillObject.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllLimitOrderPoolFillObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolFillObjectResponse,
    } as QueryAllLimitOrderPoolFillObjectResponse;
    message.limitOrderPoolFillObject = [];
    if (
      object.limitOrderPoolFillObject !== undefined &&
      object.limitOrderPoolFillObject !== null
    ) {
      for (const e of object.limitOrderPoolFillObject) {
        message.limitOrderPoolFillObject.push(LimitOrderPoolFillObject.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLimitOrderPoolFillObjectResponse): unknown {
    const obj: any = {};
    if (message.limitOrderPoolFillObject) {
      obj.limitOrderPoolFillObject = message.limitOrderPoolFillObject.map((e) =>
        e ? LimitOrderPoolFillObject.toJSON(e) : undefined
      );
    } else {
      obj.limitOrderPoolFillObject = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLimitOrderPoolFillObjectResponse>
  ): QueryAllLimitOrderPoolFillObjectResponse {
    const message = {
      ...baseQueryAllLimitOrderPoolFillObjectResponse,
    } as QueryAllLimitOrderPoolFillObjectResponse;
    message.limitOrderPoolFillObject = [];
    if (
      object.limitOrderPoolFillObject !== undefined &&
      object.limitOrderPoolFillObject !== null
    ) {
      for (const e of object.limitOrderPoolFillObject) {
        message.limitOrderPoolFillObject.push(
          LimitOrderPoolFillObject.fromPartial(e)
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
  /** Queries a PairObject by index. */
  PairObject(request: QueryGetPairObjectRequest): Promise<QueryGetPairObjectResponse>;
  /** Queries a list of PairObject items. */
  PairObjectAll(request: QueryAllPairObjectRequest): Promise<QueryAllPairObjectResponse>;
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
  /** Queries a LimitOrderPoolUserShareObject by index. */
  LimitOrderPoolUserShareObject(
    request: QueryGetLimitOrderPoolUserShareObjectRequest
  ): Promise<QueryGetLimitOrderPoolUserShareObjectResponse>;
  /** Queries a list of LimitOrderPoolUserShareObject items. */
  LimitOrderPoolUserShareObjectAll(
    request: QueryAllLimitOrderPoolUserShareObjectRequest
  ): Promise<QueryAllLimitOrderPoolUserShareObjectResponse>;
  /** Queries a LimitOrderPoolUserSharesWithdrawnObject by index. */
  LimitOrderPoolUserSharesWithdrawnObject(
    request: QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest
  ): Promise<QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse>;
  /** Queries a list of LimitOrderPoolUserSharesWithdrawnObject items. */
  LimitOrderPoolUserSharesWithdrawnObjectAll(
    request: QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest
  ): Promise<QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse>;
  /** Queries a LimitOrderPoolTotalSharesObject by index. */
  LimitOrderPoolTotalSharesObject(
    request: QueryGetLimitOrderPoolTotalSharesObjectRequest
  ): Promise<QueryGetLimitOrderPoolTotalSharesObjectResponse>;
  /** Queries a list of LimitOrderPoolTotalSharesObject items. */
  LimitOrderPoolTotalSharesObjectAll(
    request: QueryAllLimitOrderPoolTotalSharesObjectRequest
  ): Promise<QueryAllLimitOrderPoolTotalSharesObjectResponse>;
  /** Queries a LimitOrderPoolReserveObject by index. */
  LimitOrderPoolReserveObject(
    request: QueryGetLimitOrderPoolReserveObjectRequest
  ): Promise<QueryGetLimitOrderPoolReserveObjectResponse>;
  /** Queries a list of LimitOrderPoolReserveObject items. */
  LimitOrderPoolReserveObjectAll(
    request: QueryAllLimitOrderPoolReserveObjectRequest
  ): Promise<QueryAllLimitOrderPoolReserveObjectResponse>;
  /** Queries a LimitOrderPoolFillObject by index. */
  LimitOrderPoolFillObject(
    request: QueryGetLimitOrderPoolFillObjectRequest
  ): Promise<QueryGetLimitOrderPoolFillObjectResponse>;
  /** Queries a list of LimitOrderPoolFillObject items. */
  LimitOrderPoolFillObjectAll(
    request: QueryAllLimitOrderPoolFillObjectRequest
  ): Promise<QueryAllLimitOrderPoolFillObjectResponse>;
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

  PairObject(request: QueryGetPairObjectRequest): Promise<QueryGetPairObjectResponse> {
    const data = QueryGetPairObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "PairObject",
      data
    );
    return promise.then((data) =>
      QueryGetPairObjectResponse.decode(new Reader(data))
    );
  }

  PairObjectAll(
    request: QueryAllPairObjectRequest
  ): Promise<QueryAllPairObjectResponse> {
    const data = QueryAllPairObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "PairObjectAll",
      data
    );
    return promise.then((data) =>
      QueryAllPairObjectResponse.decode(new Reader(data))
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

  LimitOrderPoolUserShareObject(
    request: QueryGetLimitOrderPoolUserShareObjectRequest
  ): Promise<QueryGetLimitOrderPoolUserShareObjectResponse> {
    const data = QueryGetLimitOrderPoolUserShareObjectRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolUserShareObject",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolUserShareObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolUserShareObjectAll(
    request: QueryAllLimitOrderPoolUserShareObjectRequest
  ): Promise<QueryAllLimitOrderPoolUserShareObjectResponse> {
    const data = QueryAllLimitOrderPoolUserShareObjectRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolUserShareObjectAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolUserShareObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolUserSharesWithdrawnObject(
    request: QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest
  ): Promise<QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse> {
    const data = QueryGetLimitOrderPoolUserSharesWithdrawnObjectRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolUserSharesWithdrawnObject",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolUserSharesWithdrawnObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolUserSharesWithdrawnObjectAll(
    request: QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest
  ): Promise<QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse> {
    const data = QueryAllLimitOrderPoolUserSharesWithdrawnObjectRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolUserSharesWithdrawnObjectAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolUserSharesWithdrawnObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolTotalSharesObject(
    request: QueryGetLimitOrderPoolTotalSharesObjectRequest
  ): Promise<QueryGetLimitOrderPoolTotalSharesObjectResponse> {
    const data = QueryGetLimitOrderPoolTotalSharesObjectRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolTotalSharesObject",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolTotalSharesObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolTotalSharesObjectAll(
    request: QueryAllLimitOrderPoolTotalSharesObjectRequest
  ): Promise<QueryAllLimitOrderPoolTotalSharesObjectResponse> {
    const data = QueryAllLimitOrderPoolTotalSharesObjectRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolTotalSharesObjectAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolTotalSharesObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolReserveObject(
    request: QueryGetLimitOrderPoolReserveObjectRequest
  ): Promise<QueryGetLimitOrderPoolReserveObjectResponse> {
    const data = QueryGetLimitOrderPoolReserveObjectRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolReserveObject",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolReserveObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolReserveObjectAll(
    request: QueryAllLimitOrderPoolReserveObjectRequest
  ): Promise<QueryAllLimitOrderPoolReserveObjectResponse> {
    const data = QueryAllLimitOrderPoolReserveObjectRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolReserveObjectAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolReserveObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolFillObject(
    request: QueryGetLimitOrderPoolFillObjectRequest
  ): Promise<QueryGetLimitOrderPoolFillObjectResponse> {
    const data = QueryGetLimitOrderPoolFillObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolFillObject",
      data
    );
    return promise.then((data) =>
      QueryGetLimitOrderPoolFillObjectResponse.decode(new Reader(data))
    );
  }

  LimitOrderPoolFillObjectAll(
    request: QueryAllLimitOrderPoolFillObjectRequest
  ): Promise<QueryAllLimitOrderPoolFillObjectResponse> {
    const data = QueryAllLimitOrderPoolFillObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "LimitOrderPoolFillObjectAll",
      data
    );
    return promise.then((data) =>
      QueryAllLimitOrderPoolFillObjectResponse.decode(new Reader(data))
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
