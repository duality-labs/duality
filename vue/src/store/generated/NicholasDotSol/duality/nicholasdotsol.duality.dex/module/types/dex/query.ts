/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { Ticks } from "../dex/ticks";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Pairs } from "../dex/pairs";
import { IndexQueue } from "../dex/index_queue";
import { Nodes } from "../dex/nodes";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetTicksRequest {
  token0: string;
  token1: string;
  price: string;
  fee: string;
  orderType: string;
}

export interface QueryGetTicksResponse {
  ticks: Ticks | undefined;
}

export interface QueryAllTicksRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTicksResponse {
  ticks: Ticks[];
  pagination: PageResponse | undefined;
}

export interface QueryGetPairsRequest {
  token0: string;
  token1: string;
}

export interface QueryGetPairsResponse {
  pairs: Pairs | undefined;
}

export interface QueryAllPairsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPairsResponse {
  pairs: Pairs[];
  pagination: PageResponse | undefined;
}

export interface QueryGetIndexQueueRequest {
  token0: string;
  token1: string;
  index: number;
}

export interface QueryGetIndexQueueResponse {
  indexQueue: IndexQueue | undefined;
}

export interface QueryAllIndexQueueRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllIndexQueueResponse {
  indexQueue: IndexQueue[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNodesRequest {
  node: string;
}

export interface QueryGetNodesResponse {
  nodes: Nodes | undefined;
}

export interface QueryAllNodesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNodesResponse {
  nodes: Nodes[];
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

const baseQueryGetTicksRequest: object = {
  token0: "",
  token1: "",
  price: "",
  fee: "",
  orderType: "",
};

export const QueryGetTicksRequest = {
  encode(
    message: QueryGetTicksRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.token0 !== "") {
      writer.uint32(10).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(18).string(message.token1);
    }
    if (message.price !== "") {
      writer.uint32(26).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(34).string(message.fee);
    }
    if (message.orderType !== "") {
      writer.uint32(42).string(message.orderType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTicksRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTicksRequest } as QueryGetTicksRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token0 = reader.string();
          break;
        case 2:
          message.token1 = reader.string();
          break;
        case 3:
          message.price = reader.string();
          break;
        case 4:
          message.fee = reader.string();
          break;
        case 5:
          message.orderType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTicksRequest {
    const message = { ...baseQueryGetTicksRequest } as QueryGetTicksRequest;
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = String(object.token0);
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = String(object.token1);
    } else {
      message.token1 = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    return message;
  },

  toJSON(message: QueryGetTicksRequest): unknown {
    const obj: any = {};
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetTicksRequest>): QueryGetTicksRequest {
    const message = { ...baseQueryGetTicksRequest } as QueryGetTicksRequest;
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = object.token0;
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = object.token1;
    } else {
      message.token1 = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    return message;
  },
};

const baseQueryGetTicksResponse: object = {};

export const QueryGetTicksResponse = {
  encode(
    message: QueryGetTicksResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ticks !== undefined) {
      Ticks.encode(message.ticks, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTicksResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTicksResponse } as QueryGetTicksResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ticks = Ticks.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTicksResponse {
    const message = { ...baseQueryGetTicksResponse } as QueryGetTicksResponse;
    if (object.ticks !== undefined && object.ticks !== null) {
      message.ticks = Ticks.fromJSON(object.ticks);
    } else {
      message.ticks = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTicksResponse): unknown {
    const obj: any = {};
    message.ticks !== undefined &&
      (obj.ticks = message.ticks ? Ticks.toJSON(message.ticks) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTicksResponse>
  ): QueryGetTicksResponse {
    const message = { ...baseQueryGetTicksResponse } as QueryGetTicksResponse;
    if (object.ticks !== undefined && object.ticks !== null) {
      message.ticks = Ticks.fromPartial(object.ticks);
    } else {
      message.ticks = undefined;
    }
    return message;
  },
};

const baseQueryAllTicksRequest: object = {};

export const QueryAllTicksRequest = {
  encode(
    message: QueryAllTicksRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTicksRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTicksRequest } as QueryAllTicksRequest;
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

  fromJSON(object: any): QueryAllTicksRequest {
    const message = { ...baseQueryAllTicksRequest } as QueryAllTicksRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTicksRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllTicksRequest>): QueryAllTicksRequest {
    const message = { ...baseQueryAllTicksRequest } as QueryAllTicksRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTicksResponse: object = {};

export const QueryAllTicksResponse = {
  encode(
    message: QueryAllTicksResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.ticks) {
      Ticks.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTicksResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTicksResponse } as QueryAllTicksResponse;
    message.ticks = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ticks.push(Ticks.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllTicksResponse {
    const message = { ...baseQueryAllTicksResponse } as QueryAllTicksResponse;
    message.ticks = [];
    if (object.ticks !== undefined && object.ticks !== null) {
      for (const e of object.ticks) {
        message.ticks.push(Ticks.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTicksResponse): unknown {
    const obj: any = {};
    if (message.ticks) {
      obj.ticks = message.ticks.map((e) => (e ? Ticks.toJSON(e) : undefined));
    } else {
      obj.ticks = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTicksResponse>
  ): QueryAllTicksResponse {
    const message = { ...baseQueryAllTicksResponse } as QueryAllTicksResponse;
    message.ticks = [];
    if (object.ticks !== undefined && object.ticks !== null) {
      for (const e of object.ticks) {
        message.ticks.push(Ticks.fromPartial(e));
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

const baseQueryGetPairsRequest: object = { token0: "", token1: "" };

export const QueryGetPairsRequest = {
  encode(
    message: QueryGetPairsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.token0 !== "") {
      writer.uint32(10).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(18).string(message.token1);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPairsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPairsRequest } as QueryGetPairsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token0 = reader.string();
          break;
        case 2:
          message.token1 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPairsRequest {
    const message = { ...baseQueryGetPairsRequest } as QueryGetPairsRequest;
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = String(object.token0);
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = String(object.token1);
    } else {
      message.token1 = "";
    }
    return message;
  },

  toJSON(message: QueryGetPairsRequest): unknown {
    const obj: any = {};
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetPairsRequest>): QueryGetPairsRequest {
    const message = { ...baseQueryGetPairsRequest } as QueryGetPairsRequest;
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = object.token0;
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = object.token1;
    } else {
      message.token1 = "";
    }
    return message;
  },
};

const baseQueryGetPairsResponse: object = {};

export const QueryGetPairsResponse = {
  encode(
    message: QueryGetPairsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pairs !== undefined) {
      Pairs.encode(message.pairs, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPairsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPairsResponse } as QueryGetPairsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairs = Pairs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPairsResponse {
    const message = { ...baseQueryGetPairsResponse } as QueryGetPairsResponse;
    if (object.pairs !== undefined && object.pairs !== null) {
      message.pairs = Pairs.fromJSON(object.pairs);
    } else {
      message.pairs = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPairsResponse): unknown {
    const obj: any = {};
    message.pairs !== undefined &&
      (obj.pairs = message.pairs ? Pairs.toJSON(message.pairs) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPairsResponse>
  ): QueryGetPairsResponse {
    const message = { ...baseQueryGetPairsResponse } as QueryGetPairsResponse;
    if (object.pairs !== undefined && object.pairs !== null) {
      message.pairs = Pairs.fromPartial(object.pairs);
    } else {
      message.pairs = undefined;
    }
    return message;
  },
};

const baseQueryAllPairsRequest: object = {};

export const QueryAllPairsRequest = {
  encode(
    message: QueryAllPairsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPairsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPairsRequest } as QueryAllPairsRequest;
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

  fromJSON(object: any): QueryAllPairsRequest {
    const message = { ...baseQueryAllPairsRequest } as QueryAllPairsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPairsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllPairsRequest>): QueryAllPairsRequest {
    const message = { ...baseQueryAllPairsRequest } as QueryAllPairsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPairsResponse: object = {};

export const QueryAllPairsResponse = {
  encode(
    message: QueryAllPairsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.pairs) {
      Pairs.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPairsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPairsResponse } as QueryAllPairsResponse;
    message.pairs = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairs.push(Pairs.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllPairsResponse {
    const message = { ...baseQueryAllPairsResponse } as QueryAllPairsResponse;
    message.pairs = [];
    if (object.pairs !== undefined && object.pairs !== null) {
      for (const e of object.pairs) {
        message.pairs.push(Pairs.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPairsResponse): unknown {
    const obj: any = {};
    if (message.pairs) {
      obj.pairs = message.pairs.map((e) => (e ? Pairs.toJSON(e) : undefined));
    } else {
      obj.pairs = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPairsResponse>
  ): QueryAllPairsResponse {
    const message = { ...baseQueryAllPairsResponse } as QueryAllPairsResponse;
    message.pairs = [];
    if (object.pairs !== undefined && object.pairs !== null) {
      for (const e of object.pairs) {
        message.pairs.push(Pairs.fromPartial(e));
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

const baseQueryGetIndexQueueRequest: object = {
  token0: "",
  token1: "",
  index: 0,
};

export const QueryGetIndexQueueRequest = {
  encode(
    message: QueryGetIndexQueueRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.token0 !== "") {
      writer.uint32(10).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(18).string(message.token1);
    }
    if (message.index !== 0) {
      writer.uint32(24).int32(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetIndexQueueRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetIndexQueueRequest,
    } as QueryGetIndexQueueRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token0 = reader.string();
          break;
        case 2:
          message.token1 = reader.string();
          break;
        case 3:
          message.index = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetIndexQueueRequest {
    const message = {
      ...baseQueryGetIndexQueueRequest,
    } as QueryGetIndexQueueRequest;
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = String(object.token0);
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = String(object.token1);
    } else {
      message.token1 = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = Number(object.index);
    } else {
      message.index = 0;
    }
    return message;
  },

  toJSON(message: QueryGetIndexQueueRequest): unknown {
    const obj: any = {};
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetIndexQueueRequest>
  ): QueryGetIndexQueueRequest {
    const message = {
      ...baseQueryGetIndexQueueRequest,
    } as QueryGetIndexQueueRequest;
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = object.token0;
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = object.token1;
    } else {
      message.token1 = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = 0;
    }
    return message;
  },
};

const baseQueryGetIndexQueueResponse: object = {};

export const QueryGetIndexQueueResponse = {
  encode(
    message: QueryGetIndexQueueResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.indexQueue !== undefined) {
      IndexQueue.encode(message.indexQueue, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetIndexQueueResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetIndexQueueResponse,
    } as QueryGetIndexQueueResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.indexQueue = IndexQueue.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetIndexQueueResponse {
    const message = {
      ...baseQueryGetIndexQueueResponse,
    } as QueryGetIndexQueueResponse;
    if (object.indexQueue !== undefined && object.indexQueue !== null) {
      message.indexQueue = IndexQueue.fromJSON(object.indexQueue);
    } else {
      message.indexQueue = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetIndexQueueResponse): unknown {
    const obj: any = {};
    message.indexQueue !== undefined &&
      (obj.indexQueue = message.indexQueue
        ? IndexQueue.toJSON(message.indexQueue)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetIndexQueueResponse>
  ): QueryGetIndexQueueResponse {
    const message = {
      ...baseQueryGetIndexQueueResponse,
    } as QueryGetIndexQueueResponse;
    if (object.indexQueue !== undefined && object.indexQueue !== null) {
      message.indexQueue = IndexQueue.fromPartial(object.indexQueue);
    } else {
      message.indexQueue = undefined;
    }
    return message;
  },
};

const baseQueryAllIndexQueueRequest: object = {};

export const QueryAllIndexQueueRequest = {
  encode(
    message: QueryAllIndexQueueRequest,
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
  ): QueryAllIndexQueueRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllIndexQueueRequest,
    } as QueryAllIndexQueueRequest;
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

  fromJSON(object: any): QueryAllIndexQueueRequest {
    const message = {
      ...baseQueryAllIndexQueueRequest,
    } as QueryAllIndexQueueRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllIndexQueueRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllIndexQueueRequest>
  ): QueryAllIndexQueueRequest {
    const message = {
      ...baseQueryAllIndexQueueRequest,
    } as QueryAllIndexQueueRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllIndexQueueResponse: object = {};

export const QueryAllIndexQueueResponse = {
  encode(
    message: QueryAllIndexQueueResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.indexQueue) {
      IndexQueue.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllIndexQueueResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllIndexQueueResponse,
    } as QueryAllIndexQueueResponse;
    message.indexQueue = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.indexQueue.push(IndexQueue.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllIndexQueueResponse {
    const message = {
      ...baseQueryAllIndexQueueResponse,
    } as QueryAllIndexQueueResponse;
    message.indexQueue = [];
    if (object.indexQueue !== undefined && object.indexQueue !== null) {
      for (const e of object.indexQueue) {
        message.indexQueue.push(IndexQueue.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllIndexQueueResponse): unknown {
    const obj: any = {};
    if (message.indexQueue) {
      obj.indexQueue = message.indexQueue.map((e) =>
        e ? IndexQueue.toJSON(e) : undefined
      );
    } else {
      obj.indexQueue = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllIndexQueueResponse>
  ): QueryAllIndexQueueResponse {
    const message = {
      ...baseQueryAllIndexQueueResponse,
    } as QueryAllIndexQueueResponse;
    message.indexQueue = [];
    if (object.indexQueue !== undefined && object.indexQueue !== null) {
      for (const e of object.indexQueue) {
        message.indexQueue.push(IndexQueue.fromPartial(e));
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

const baseQueryGetNodesRequest: object = { node: "" };

export const QueryGetNodesRequest = {
  encode(
    message: QueryGetNodesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.node !== "") {
      writer.uint32(10).string(message.node);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNodesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNodesRequest } as QueryGetNodesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.node = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNodesRequest {
    const message = { ...baseQueryGetNodesRequest } as QueryGetNodesRequest;
    if (object.node !== undefined && object.node !== null) {
      message.node = String(object.node);
    } else {
      message.node = "";
    }
    return message;
  },

  toJSON(message: QueryGetNodesRequest): unknown {
    const obj: any = {};
    message.node !== undefined && (obj.node = message.node);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetNodesRequest>): QueryGetNodesRequest {
    const message = { ...baseQueryGetNodesRequest } as QueryGetNodesRequest;
    if (object.node !== undefined && object.node !== null) {
      message.node = object.node;
    } else {
      message.node = "";
    }
    return message;
  },
};

const baseQueryGetNodesResponse: object = {};

export const QueryGetNodesResponse = {
  encode(
    message: QueryGetNodesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nodes !== undefined) {
      Nodes.encode(message.nodes, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNodesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNodesResponse } as QueryGetNodesResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nodes = Nodes.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNodesResponse {
    const message = { ...baseQueryGetNodesResponse } as QueryGetNodesResponse;
    if (object.nodes !== undefined && object.nodes !== null) {
      message.nodes = Nodes.fromJSON(object.nodes);
    } else {
      message.nodes = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNodesResponse): unknown {
    const obj: any = {};
    message.nodes !== undefined &&
      (obj.nodes = message.nodes ? Nodes.toJSON(message.nodes) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNodesResponse>
  ): QueryGetNodesResponse {
    const message = { ...baseQueryGetNodesResponse } as QueryGetNodesResponse;
    if (object.nodes !== undefined && object.nodes !== null) {
      message.nodes = Nodes.fromPartial(object.nodes);
    } else {
      message.nodes = undefined;
    }
    return message;
  },
};

const baseQueryAllNodesRequest: object = {};

export const QueryAllNodesRequest = {
  encode(
    message: QueryAllNodesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNodesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNodesRequest } as QueryAllNodesRequest;
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

  fromJSON(object: any): QueryAllNodesRequest {
    const message = { ...baseQueryAllNodesRequest } as QueryAllNodesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNodesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllNodesRequest>): QueryAllNodesRequest {
    const message = { ...baseQueryAllNodesRequest } as QueryAllNodesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNodesResponse: object = {};

export const QueryAllNodesResponse = {
  encode(
    message: QueryAllNodesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.nodes) {
      Nodes.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNodesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNodesResponse } as QueryAllNodesResponse;
    message.nodes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nodes.push(Nodes.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllNodesResponse {
    const message = { ...baseQueryAllNodesResponse } as QueryAllNodesResponse;
    message.nodes = [];
    if (object.nodes !== undefined && object.nodes !== null) {
      for (const e of object.nodes) {
        message.nodes.push(Nodes.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNodesResponse): unknown {
    const obj: any = {};
    if (message.nodes) {
      obj.nodes = message.nodes.map((e) => (e ? Nodes.toJSON(e) : undefined));
    } else {
      obj.nodes = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNodesResponse>
  ): QueryAllNodesResponse {
    const message = { ...baseQueryAllNodesResponse } as QueryAllNodesResponse;
    message.nodes = [];
    if (object.nodes !== undefined && object.nodes !== null) {
      for (const e of object.nodes) {
        message.nodes.push(Nodes.fromPartial(e));
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
  /** Queries a Ticks by index. */
  Ticks(request: QueryGetTicksRequest): Promise<QueryGetTicksResponse>;
  /** Queries a list of Ticks items. */
  TicksAll(request: QueryAllTicksRequest): Promise<QueryAllTicksResponse>;
  /** Queries a Pairs by index. */
  Pairs(request: QueryGetPairsRequest): Promise<QueryGetPairsResponse>;
  /** Queries a list of Pairs items. */
  PairsAll(request: QueryAllPairsRequest): Promise<QueryAllPairsResponse>;
  /** Queries a IndexQueue by index. */
  IndexQueue(
    request: QueryGetIndexQueueRequest
  ): Promise<QueryGetIndexQueueResponse>;
  /** Queries a list of IndexQueue items. */
  IndexQueueAll(
    request: QueryAllIndexQueueRequest
  ): Promise<QueryAllIndexQueueResponse>;
  /** Queries a Nodes by index. */
  Nodes(request: QueryGetNodesRequest): Promise<QueryGetNodesResponse>;
  /** Queries a list of Nodes items. */
  NodesAll(request: QueryAllNodesRequest): Promise<QueryAllNodesResponse>;
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

  Ticks(request: QueryGetTicksRequest): Promise<QueryGetTicksResponse> {
    const data = QueryGetTicksRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "Ticks",
      data
    );
    return promise.then((data) =>
      QueryGetTicksResponse.decode(new Reader(data))
    );
  }

  TicksAll(request: QueryAllTicksRequest): Promise<QueryAllTicksResponse> {
    const data = QueryAllTicksRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "TicksAll",
      data
    );
    return promise.then((data) =>
      QueryAllTicksResponse.decode(new Reader(data))
    );
  }

  Pairs(request: QueryGetPairsRequest): Promise<QueryGetPairsResponse> {
    const data = QueryGetPairsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "Pairs",
      data
    );
    return promise.then((data) =>
      QueryGetPairsResponse.decode(new Reader(data))
    );
  }

  PairsAll(request: QueryAllPairsRequest): Promise<QueryAllPairsResponse> {
    const data = QueryAllPairsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "PairsAll",
      data
    );
    return promise.then((data) =>
      QueryAllPairsResponse.decode(new Reader(data))
    );
  }

  IndexQueue(
    request: QueryGetIndexQueueRequest
  ): Promise<QueryGetIndexQueueResponse> {
    const data = QueryGetIndexQueueRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "IndexQueue",
      data
    );
    return promise.then((data) =>
      QueryGetIndexQueueResponse.decode(new Reader(data))
    );
  }

  IndexQueueAll(
    request: QueryAllIndexQueueRequest
  ): Promise<QueryAllIndexQueueResponse> {
    const data = QueryAllIndexQueueRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "IndexQueueAll",
      data
    );
    return promise.then((data) =>
      QueryAllIndexQueueResponse.decode(new Reader(data))
    );
  }

  Nodes(request: QueryGetNodesRequest): Promise<QueryGetNodesResponse> {
    const data = QueryGetNodesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "Nodes",
      data
    );
    return promise.then((data) =>
      QueryGetNodesResponse.decode(new Reader(data))
    );
  }

  NodesAll(request: QueryAllNodesRequest): Promise<QueryAllNodesResponse> {
    const data = QueryAllNodesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "NodesAll",
      data
    );
    return promise.then((data) =>
      QueryAllNodesResponse.decode(new Reader(data))
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
