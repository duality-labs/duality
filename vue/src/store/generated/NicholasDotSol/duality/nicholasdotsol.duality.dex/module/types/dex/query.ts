/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../dex/params";
import { Nodes } from "../dex/nodes";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { VirtualPriceTickQueue } from "../dex/virtual_price_tick_queue";
import { Ticks } from "../dex/ticks";
import { VirtualPriceTickList } from "../dex/virtual_price_tick_list";
import { BitArr } from "../dex/bit_arr";
import { Pairs } from "../dex/pairs";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetNodesRequest {
  id: number;
}

export interface QueryGetNodesResponse {
  Nodes: Nodes | undefined;
}

export interface QueryAllNodesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNodesResponse {
  Nodes: Nodes[];
  pagination: PageResponse | undefined;
}

export interface QueryGetVirtualPriceTickQueueRequest {
  id: number;
}

export interface QueryGetVirtualPriceTickQueueResponse {
  VirtualPriceTickQueue: VirtualPriceTickQueue | undefined;
}

export interface QueryAllVirtualPriceTickQueueRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllVirtualPriceTickQueueResponse {
  VirtualPriceTickQueue: VirtualPriceTickQueue[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTicksRequest {
  price: string;
  fee: string;
  direction: string;
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

export interface QueryGetVirtualPriceTickListRequest {
  vPrice: string;
  direction: string;
  orderType: string;
}

export interface QueryGetVirtualPriceTickListResponse {
  virtualPriceTickList: VirtualPriceTickList | undefined;
}

export interface QueryAllVirtualPriceTickListRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllVirtualPriceTickListResponse {
  virtualPriceTickList: VirtualPriceTickList[];
  pagination: PageResponse | undefined;
}

export interface QueryGetBitArrRequest {
  id: number;
}

export interface QueryGetBitArrResponse {
  BitArr: BitArr | undefined;
}

export interface QueryAllBitArrRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBitArrResponse {
  BitArr: BitArr[];
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

const baseQueryGetNodesRequest: object = { id: 0 };

export const QueryGetNodesRequest = {
  encode(
    message: QueryGetNodesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
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
          message.id = longToNumber(reader.uint64() as Long);
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
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetNodesRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetNodesRequest>): QueryGetNodesRequest {
    const message = { ...baseQueryGetNodesRequest } as QueryGetNodesRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
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
    if (message.Nodes !== undefined) {
      Nodes.encode(message.Nodes, writer.uint32(10).fork()).ldelim();
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
          message.Nodes = Nodes.decode(reader, reader.uint32());
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
    if (object.Nodes !== undefined && object.Nodes !== null) {
      message.Nodes = Nodes.fromJSON(object.Nodes);
    } else {
      message.Nodes = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNodesResponse): unknown {
    const obj: any = {};
    message.Nodes !== undefined &&
      (obj.Nodes = message.Nodes ? Nodes.toJSON(message.Nodes) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNodesResponse>
  ): QueryGetNodesResponse {
    const message = { ...baseQueryGetNodesResponse } as QueryGetNodesResponse;
    if (object.Nodes !== undefined && object.Nodes !== null) {
      message.Nodes = Nodes.fromPartial(object.Nodes);
    } else {
      message.Nodes = undefined;
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
    for (const v of message.Nodes) {
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
    message.Nodes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Nodes.push(Nodes.decode(reader, reader.uint32()));
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
    message.Nodes = [];
    if (object.Nodes !== undefined && object.Nodes !== null) {
      for (const e of object.Nodes) {
        message.Nodes.push(Nodes.fromJSON(e));
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
    if (message.Nodes) {
      obj.Nodes = message.Nodes.map((e) => (e ? Nodes.toJSON(e) : undefined));
    } else {
      obj.Nodes = [];
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
    message.Nodes = [];
    if (object.Nodes !== undefined && object.Nodes !== null) {
      for (const e of object.Nodes) {
        message.Nodes.push(Nodes.fromPartial(e));
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

const baseQueryGetVirtualPriceTickQueueRequest: object = { id: 0 };

export const QueryGetVirtualPriceTickQueueRequest = {
  encode(
    message: QueryGetVirtualPriceTickQueueRequest,
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
  ): QueryGetVirtualPriceTickQueueRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualPriceTickQueueRequest,
    } as QueryGetVirtualPriceTickQueueRequest;
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

  fromJSON(object: any): QueryGetVirtualPriceTickQueueRequest {
    const message = {
      ...baseQueryGetVirtualPriceTickQueueRequest,
    } as QueryGetVirtualPriceTickQueueRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetVirtualPriceTickQueueRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualPriceTickQueueRequest>
  ): QueryGetVirtualPriceTickQueueRequest {
    const message = {
      ...baseQueryGetVirtualPriceTickQueueRequest,
    } as QueryGetVirtualPriceTickQueueRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetVirtualPriceTickQueueResponse: object = {};

export const QueryGetVirtualPriceTickQueueResponse = {
  encode(
    message: QueryGetVirtualPriceTickQueueResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.VirtualPriceTickQueue !== undefined) {
      VirtualPriceTickQueue.encode(
        message.VirtualPriceTickQueue,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualPriceTickQueueResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualPriceTickQueueResponse,
    } as QueryGetVirtualPriceTickQueueResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.VirtualPriceTickQueue = VirtualPriceTickQueue.decode(
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

  fromJSON(object: any): QueryGetVirtualPriceTickQueueResponse {
    const message = {
      ...baseQueryGetVirtualPriceTickQueueResponse,
    } as QueryGetVirtualPriceTickQueueResponse;
    if (
      object.VirtualPriceTickQueue !== undefined &&
      object.VirtualPriceTickQueue !== null
    ) {
      message.VirtualPriceTickQueue = VirtualPriceTickQueue.fromJSON(
        object.VirtualPriceTickQueue
      );
    } else {
      message.VirtualPriceTickQueue = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetVirtualPriceTickQueueResponse): unknown {
    const obj: any = {};
    message.VirtualPriceTickQueue !== undefined &&
      (obj.VirtualPriceTickQueue = message.VirtualPriceTickQueue
        ? VirtualPriceTickQueue.toJSON(message.VirtualPriceTickQueue)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualPriceTickQueueResponse>
  ): QueryGetVirtualPriceTickQueueResponse {
    const message = {
      ...baseQueryGetVirtualPriceTickQueueResponse,
    } as QueryGetVirtualPriceTickQueueResponse;
    if (
      object.VirtualPriceTickQueue !== undefined &&
      object.VirtualPriceTickQueue !== null
    ) {
      message.VirtualPriceTickQueue = VirtualPriceTickQueue.fromPartial(
        object.VirtualPriceTickQueue
      );
    } else {
      message.VirtualPriceTickQueue = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualPriceTickQueueRequest: object = {};

export const QueryAllVirtualPriceTickQueueRequest = {
  encode(
    message: QueryAllVirtualPriceTickQueueRequest,
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
  ): QueryAllVirtualPriceTickQueueRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualPriceTickQueueRequest,
    } as QueryAllVirtualPriceTickQueueRequest;
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

  fromJSON(object: any): QueryAllVirtualPriceTickQueueRequest {
    const message = {
      ...baseQueryAllVirtualPriceTickQueueRequest,
    } as QueryAllVirtualPriceTickQueueRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualPriceTickQueueRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualPriceTickQueueRequest>
  ): QueryAllVirtualPriceTickQueueRequest {
    const message = {
      ...baseQueryAllVirtualPriceTickQueueRequest,
    } as QueryAllVirtualPriceTickQueueRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualPriceTickQueueResponse: object = {};

export const QueryAllVirtualPriceTickQueueResponse = {
  encode(
    message: QueryAllVirtualPriceTickQueueResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.VirtualPriceTickQueue) {
      VirtualPriceTickQueue.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllVirtualPriceTickQueueResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualPriceTickQueueResponse,
    } as QueryAllVirtualPriceTickQueueResponse;
    message.VirtualPriceTickQueue = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.VirtualPriceTickQueue.push(
            VirtualPriceTickQueue.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllVirtualPriceTickQueueResponse {
    const message = {
      ...baseQueryAllVirtualPriceTickQueueResponse,
    } as QueryAllVirtualPriceTickQueueResponse;
    message.VirtualPriceTickQueue = [];
    if (
      object.VirtualPriceTickQueue !== undefined &&
      object.VirtualPriceTickQueue !== null
    ) {
      for (const e of object.VirtualPriceTickQueue) {
        message.VirtualPriceTickQueue.push(VirtualPriceTickQueue.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualPriceTickQueueResponse): unknown {
    const obj: any = {};
    if (message.VirtualPriceTickQueue) {
      obj.VirtualPriceTickQueue = message.VirtualPriceTickQueue.map((e) =>
        e ? VirtualPriceTickQueue.toJSON(e) : undefined
      );
    } else {
      obj.VirtualPriceTickQueue = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualPriceTickQueueResponse>
  ): QueryAllVirtualPriceTickQueueResponse {
    const message = {
      ...baseQueryAllVirtualPriceTickQueueResponse,
    } as QueryAllVirtualPriceTickQueueResponse;
    message.VirtualPriceTickQueue = [];
    if (
      object.VirtualPriceTickQueue !== undefined &&
      object.VirtualPriceTickQueue !== null
    ) {
      for (const e of object.VirtualPriceTickQueue) {
        message.VirtualPriceTickQueue.push(
          VirtualPriceTickQueue.fromPartial(e)
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

const baseQueryGetTicksRequest: object = {
  price: "",
  fee: "",
  direction: "",
  orderType: "",
};

export const QueryGetTicksRequest = {
  encode(
    message: QueryGetTicksRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.price !== "") {
      writer.uint32(10).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(18).string(message.fee);
    }
    if (message.direction !== "") {
      writer.uint32(26).string(message.direction);
    }
    if (message.orderType !== "") {
      writer.uint32(34).string(message.orderType);
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
          message.price = reader.string();
          break;
        case 2:
          message.fee = reader.string();
          break;
        case 3:
          message.direction = reader.string();
          break;
        case 4:
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
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = String(object.direction);
    } else {
      message.direction = "";
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
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.direction !== undefined && (obj.direction = message.direction);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetTicksRequest>): QueryGetTicksRequest {
    const message = { ...baseQueryGetTicksRequest } as QueryGetTicksRequest;
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
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = object.direction;
    } else {
      message.direction = "";
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

const baseQueryGetVirtualPriceTickListRequest: object = {
  vPrice: "",
  direction: "",
  orderType: "",
};

export const QueryGetVirtualPriceTickListRequest = {
  encode(
    message: QueryGetVirtualPriceTickListRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.vPrice !== "") {
      writer.uint32(10).string(message.vPrice);
    }
    if (message.direction !== "") {
      writer.uint32(18).string(message.direction);
    }
    if (message.orderType !== "") {
      writer.uint32(26).string(message.orderType);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualPriceTickListRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualPriceTickListRequest,
    } as QueryGetVirtualPriceTickListRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vPrice = reader.string();
          break;
        case 2:
          message.direction = reader.string();
          break;
        case 3:
          message.orderType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVirtualPriceTickListRequest {
    const message = {
      ...baseQueryGetVirtualPriceTickListRequest,
    } as QueryGetVirtualPriceTickListRequest;
    if (object.vPrice !== undefined && object.vPrice !== null) {
      message.vPrice = String(object.vPrice);
    } else {
      message.vPrice = "";
    }
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = String(object.direction);
    } else {
      message.direction = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    return message;
  },

  toJSON(message: QueryGetVirtualPriceTickListRequest): unknown {
    const obj: any = {};
    message.vPrice !== undefined && (obj.vPrice = message.vPrice);
    message.direction !== undefined && (obj.direction = message.direction);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualPriceTickListRequest>
  ): QueryGetVirtualPriceTickListRequest {
    const message = {
      ...baseQueryGetVirtualPriceTickListRequest,
    } as QueryGetVirtualPriceTickListRequest;
    if (object.vPrice !== undefined && object.vPrice !== null) {
      message.vPrice = object.vPrice;
    } else {
      message.vPrice = "";
    }
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = object.direction;
    } else {
      message.direction = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    return message;
  },
};

const baseQueryGetVirtualPriceTickListResponse: object = {};

export const QueryGetVirtualPriceTickListResponse = {
  encode(
    message: QueryGetVirtualPriceTickListResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.virtualPriceTickList !== undefined) {
      VirtualPriceTickList.encode(
        message.virtualPriceTickList,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVirtualPriceTickListResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVirtualPriceTickListResponse,
    } as QueryGetVirtualPriceTickListResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualPriceTickList = VirtualPriceTickList.decode(
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

  fromJSON(object: any): QueryGetVirtualPriceTickListResponse {
    const message = {
      ...baseQueryGetVirtualPriceTickListResponse,
    } as QueryGetVirtualPriceTickListResponse;
    if (
      object.virtualPriceTickList !== undefined &&
      object.virtualPriceTickList !== null
    ) {
      message.virtualPriceTickList = VirtualPriceTickList.fromJSON(
        object.virtualPriceTickList
      );
    } else {
      message.virtualPriceTickList = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetVirtualPriceTickListResponse): unknown {
    const obj: any = {};
    message.virtualPriceTickList !== undefined &&
      (obj.virtualPriceTickList = message.virtualPriceTickList
        ? VirtualPriceTickList.toJSON(message.virtualPriceTickList)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVirtualPriceTickListResponse>
  ): QueryGetVirtualPriceTickListResponse {
    const message = {
      ...baseQueryGetVirtualPriceTickListResponse,
    } as QueryGetVirtualPriceTickListResponse;
    if (
      object.virtualPriceTickList !== undefined &&
      object.virtualPriceTickList !== null
    ) {
      message.virtualPriceTickList = VirtualPriceTickList.fromPartial(
        object.virtualPriceTickList
      );
    } else {
      message.virtualPriceTickList = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualPriceTickListRequest: object = {};

export const QueryAllVirtualPriceTickListRequest = {
  encode(
    message: QueryAllVirtualPriceTickListRequest,
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
  ): QueryAllVirtualPriceTickListRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualPriceTickListRequest,
    } as QueryAllVirtualPriceTickListRequest;
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

  fromJSON(object: any): QueryAllVirtualPriceTickListRequest {
    const message = {
      ...baseQueryAllVirtualPriceTickListRequest,
    } as QueryAllVirtualPriceTickListRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualPriceTickListRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualPriceTickListRequest>
  ): QueryAllVirtualPriceTickListRequest {
    const message = {
      ...baseQueryAllVirtualPriceTickListRequest,
    } as QueryAllVirtualPriceTickListRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllVirtualPriceTickListResponse: object = {};

export const QueryAllVirtualPriceTickListResponse = {
  encode(
    message: QueryAllVirtualPriceTickListResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.virtualPriceTickList) {
      VirtualPriceTickList.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllVirtualPriceTickListResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVirtualPriceTickListResponse,
    } as QueryAllVirtualPriceTickListResponse;
    message.virtualPriceTickList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.virtualPriceTickList.push(
            VirtualPriceTickList.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllVirtualPriceTickListResponse {
    const message = {
      ...baseQueryAllVirtualPriceTickListResponse,
    } as QueryAllVirtualPriceTickListResponse;
    message.virtualPriceTickList = [];
    if (
      object.virtualPriceTickList !== undefined &&
      object.virtualPriceTickList !== null
    ) {
      for (const e of object.virtualPriceTickList) {
        message.virtualPriceTickList.push(VirtualPriceTickList.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVirtualPriceTickListResponse): unknown {
    const obj: any = {};
    if (message.virtualPriceTickList) {
      obj.virtualPriceTickList = message.virtualPriceTickList.map((e) =>
        e ? VirtualPriceTickList.toJSON(e) : undefined
      );
    } else {
      obj.virtualPriceTickList = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVirtualPriceTickListResponse>
  ): QueryAllVirtualPriceTickListResponse {
    const message = {
      ...baseQueryAllVirtualPriceTickListResponse,
    } as QueryAllVirtualPriceTickListResponse;
    message.virtualPriceTickList = [];
    if (
      object.virtualPriceTickList !== undefined &&
      object.virtualPriceTickList !== null
    ) {
      for (const e of object.virtualPriceTickList) {
        message.virtualPriceTickList.push(VirtualPriceTickList.fromPartial(e));
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

const baseQueryGetBitArrRequest: object = { id: 0 };

export const QueryGetBitArrRequest = {
  encode(
    message: QueryGetBitArrRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBitArrRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBitArrRequest } as QueryGetBitArrRequest;
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

  fromJSON(object: any): QueryGetBitArrRequest {
    const message = { ...baseQueryGetBitArrRequest } as QueryGetBitArrRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetBitArrRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBitArrRequest>
  ): QueryGetBitArrRequest {
    const message = { ...baseQueryGetBitArrRequest } as QueryGetBitArrRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetBitArrResponse: object = {};

export const QueryGetBitArrResponse = {
  encode(
    message: QueryGetBitArrResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.BitArr !== undefined) {
      BitArr.encode(message.BitArr, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBitArrResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBitArrResponse } as QueryGetBitArrResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.BitArr = BitArr.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBitArrResponse {
    const message = { ...baseQueryGetBitArrResponse } as QueryGetBitArrResponse;
    if (object.BitArr !== undefined && object.BitArr !== null) {
      message.BitArr = BitArr.fromJSON(object.BitArr);
    } else {
      message.BitArr = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBitArrResponse): unknown {
    const obj: any = {};
    message.BitArr !== undefined &&
      (obj.BitArr = message.BitArr ? BitArr.toJSON(message.BitArr) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBitArrResponse>
  ): QueryGetBitArrResponse {
    const message = { ...baseQueryGetBitArrResponse } as QueryGetBitArrResponse;
    if (object.BitArr !== undefined && object.BitArr !== null) {
      message.BitArr = BitArr.fromPartial(object.BitArr);
    } else {
      message.BitArr = undefined;
    }
    return message;
  },
};

const baseQueryAllBitArrRequest: object = {};

export const QueryAllBitArrRequest = {
  encode(
    message: QueryAllBitArrRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBitArrRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBitArrRequest } as QueryAllBitArrRequest;
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

  fromJSON(object: any): QueryAllBitArrRequest {
    const message = { ...baseQueryAllBitArrRequest } as QueryAllBitArrRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBitArrRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBitArrRequest>
  ): QueryAllBitArrRequest {
    const message = { ...baseQueryAllBitArrRequest } as QueryAllBitArrRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBitArrResponse: object = {};

export const QueryAllBitArrResponse = {
  encode(
    message: QueryAllBitArrResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.BitArr) {
      BitArr.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBitArrResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBitArrResponse } as QueryAllBitArrResponse;
    message.BitArr = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.BitArr.push(BitArr.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllBitArrResponse {
    const message = { ...baseQueryAllBitArrResponse } as QueryAllBitArrResponse;
    message.BitArr = [];
    if (object.BitArr !== undefined && object.BitArr !== null) {
      for (const e of object.BitArr) {
        message.BitArr.push(BitArr.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBitArrResponse): unknown {
    const obj: any = {};
    if (message.BitArr) {
      obj.BitArr = message.BitArr.map((e) =>
        e ? BitArr.toJSON(e) : undefined
      );
    } else {
      obj.BitArr = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBitArrResponse>
  ): QueryAllBitArrResponse {
    const message = { ...baseQueryAllBitArrResponse } as QueryAllBitArrResponse;
    message.BitArr = [];
    if (object.BitArr !== undefined && object.BitArr !== null) {
      for (const e of object.BitArr) {
        message.BitArr.push(BitArr.fromPartial(e));
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Nodes by id. */
  Nodes(request: QueryGetNodesRequest): Promise<QueryGetNodesResponse>;
  /** Queries a list of Nodes items. */
  NodesAll(request: QueryAllNodesRequest): Promise<QueryAllNodesResponse>;
  /** Queries a VirtualPriceTickQueue by id. */
  VirtualPriceTickQueue(
    request: QueryGetVirtualPriceTickQueueRequest
  ): Promise<QueryGetVirtualPriceTickQueueResponse>;
  /** Queries a list of VirtualPriceTickQueue items. */
  VirtualPriceTickQueueAll(
    request: QueryAllVirtualPriceTickQueueRequest
  ): Promise<QueryAllVirtualPriceTickQueueResponse>;
  /** Queries a Ticks by index. */
  Ticks(request: QueryGetTicksRequest): Promise<QueryGetTicksResponse>;
  /** Queries a list of Ticks items. */
  TicksAll(request: QueryAllTicksRequest): Promise<QueryAllTicksResponse>;
  /** Queries a VirtualPriceTickList by index. */
  VirtualPriceTickList(
    request: QueryGetVirtualPriceTickListRequest
  ): Promise<QueryGetVirtualPriceTickListResponse>;
  /** Queries a list of VirtualPriceTickList items. */
  VirtualPriceTickListAll(
    request: QueryAllVirtualPriceTickListRequest
  ): Promise<QueryAllVirtualPriceTickListResponse>;
  /** Queries a BitArr by id. */
  BitArr(request: QueryGetBitArrRequest): Promise<QueryGetBitArrResponse>;
  /** Queries a list of BitArr items. */
  BitArrAll(request: QueryAllBitArrRequest): Promise<QueryAllBitArrResponse>;
  /** Queries a Pairs by index. */
  Pairs(request: QueryGetPairsRequest): Promise<QueryGetPairsResponse>;
  /** Queries a list of Pairs items. */
  PairsAll(request: QueryAllPairsRequest): Promise<QueryAllPairsResponse>;
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

  VirtualPriceTickQueue(
    request: QueryGetVirtualPriceTickQueueRequest
  ): Promise<QueryGetVirtualPriceTickQueueResponse> {
    const data = QueryGetVirtualPriceTickQueueRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "VirtualPriceTickQueue",
      data
    );
    return promise.then((data) =>
      QueryGetVirtualPriceTickQueueResponse.decode(new Reader(data))
    );
  }

  VirtualPriceTickQueueAll(
    request: QueryAllVirtualPriceTickQueueRequest
  ): Promise<QueryAllVirtualPriceTickQueueResponse> {
    const data = QueryAllVirtualPriceTickQueueRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "VirtualPriceTickQueueAll",
      data
    );
    return promise.then((data) =>
      QueryAllVirtualPriceTickQueueResponse.decode(new Reader(data))
    );
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

  VirtualPriceTickList(
    request: QueryGetVirtualPriceTickListRequest
  ): Promise<QueryGetVirtualPriceTickListResponse> {
    const data = QueryGetVirtualPriceTickListRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "VirtualPriceTickList",
      data
    );
    return promise.then((data) =>
      QueryGetVirtualPriceTickListResponse.decode(new Reader(data))
    );
  }

  VirtualPriceTickListAll(
    request: QueryAllVirtualPriceTickListRequest
  ): Promise<QueryAllVirtualPriceTickListResponse> {
    const data = QueryAllVirtualPriceTickListRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "VirtualPriceTickListAll",
      data
    );
    return promise.then((data) =>
      QueryAllVirtualPriceTickListResponse.decode(new Reader(data))
    );
  }

  BitArr(request: QueryGetBitArrRequest): Promise<QueryGetBitArrResponse> {
    const data = QueryGetBitArrRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "BitArr",
      data
    );
    return promise.then((data) =>
      QueryGetBitArrResponse.decode(new Reader(data))
    );
  }

  BitArrAll(request: QueryAllBitArrRequest): Promise<QueryAllBitArrResponse> {
    const data = QueryAllBitArrRequest.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Query",
      "BitArrAll",
      data
    );
    return promise.then((data) =>
      QueryAllBitArrResponse.decode(new Reader(data))
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
