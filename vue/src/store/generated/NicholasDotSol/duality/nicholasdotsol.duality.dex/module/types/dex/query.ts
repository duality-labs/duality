/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { TickMap } from "../dex/tick_map";
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

export interface QueryGetTickMapRequest {
  tickIndex: string;
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

const baseQueryGetTickMapRequest: object = { tickIndex: "" };

export const QueryGetTickMapRequest = {
  encode(
    message: QueryGetTickMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tickIndex !== "") {
      writer.uint32(10).string(message.tickIndex);
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
          message.tickIndex = reader.string();
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
      message.tickIndex = String(object.tickIndex);
    } else {
      message.tickIndex = "";
    }
    return message;
  },

  toJSON(message: QueryGetTickMapRequest): unknown {
    const obj: any = {};
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTickMapRequest>
  ): QueryGetTickMapRequest {
    const message = { ...baseQueryGetTickMapRequest } as QueryGetTickMapRequest;
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = "";
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a TickMap by index. */
  TickMap(request: QueryGetTickMapRequest): Promise<QueryGetTickMapResponse>;
  /** Queries a list of TickMap items. */
  TickMapAll(request: QueryAllTickMapRequest): Promise<QueryAllTickMapResponse>;
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
