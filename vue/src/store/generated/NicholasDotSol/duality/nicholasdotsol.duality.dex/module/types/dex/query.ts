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
