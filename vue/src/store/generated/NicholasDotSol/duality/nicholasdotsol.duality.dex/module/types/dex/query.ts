/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { Ticks } from "../dex/ticks";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

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

const baseQueryGetTicksRequest: object = { token0: "", token1: "" };

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
    return message;
  },

  toJSON(message: QueryGetTicksRequest): unknown {
    const obj: any = {};
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Ticks by index. */
  Ticks(request: QueryGetTicksRequest): Promise<QueryGetTicksResponse>;
  /** Queries a list of Ticks items. */
  TicksAll(request: QueryAllTicksRequest): Promise<QueryAllTicksResponse>;
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
