/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Proposal } from "../cosmos/gov/v1beta1/gov";

export const protobufPackage = "cosmos.adminmodule.adminmodule";

/** this line is used by starport scaffolding # 3 */
export interface QueryAdminsRequest {
}

export interface QueryAdminsResponse {
  admins: string[];
}

export interface QueryArchivedProposalsRequest {
}

export interface QueryArchivedProposalsResponse {
  proposals: Proposal[];
}

function createBaseQueryAdminsRequest(): QueryAdminsRequest {
  return {};
}

export const QueryAdminsRequest = {
  encode(_: QueryAdminsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAdminsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAdminsRequest();
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

  fromJSON(_: any): QueryAdminsRequest {
    return {};
  },

  toJSON(_: QueryAdminsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAdminsRequest>, I>>(_: I): QueryAdminsRequest {
    const message = createBaseQueryAdminsRequest();
    return message;
  },
};

function createBaseQueryAdminsResponse(): QueryAdminsResponse {
  return { admins: [] };
}

export const QueryAdminsResponse = {
  encode(message: QueryAdminsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.admins) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAdminsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAdminsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.admins.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAdminsResponse {
    return { admins: Array.isArray(object?.admins) ? object.admins.map((e: any) => String(e)) : [] };
  },

  toJSON(message: QueryAdminsResponse): unknown {
    const obj: any = {};
    if (message.admins) {
      obj.admins = message.admins.map((e) => e);
    } else {
      obj.admins = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAdminsResponse>, I>>(object: I): QueryAdminsResponse {
    const message = createBaseQueryAdminsResponse();
    message.admins = object.admins?.map((e) => e) || [];
    return message;
  },
};

function createBaseQueryArchivedProposalsRequest(): QueryArchivedProposalsRequest {
  return {};
}

export const QueryArchivedProposalsRequest = {
  encode(_: QueryArchivedProposalsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryArchivedProposalsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryArchivedProposalsRequest();
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

  fromJSON(_: any): QueryArchivedProposalsRequest {
    return {};
  },

  toJSON(_: QueryArchivedProposalsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryArchivedProposalsRequest>, I>>(_: I): QueryArchivedProposalsRequest {
    const message = createBaseQueryArchivedProposalsRequest();
    return message;
  },
};

function createBaseQueryArchivedProposalsResponse(): QueryArchivedProposalsResponse {
  return { proposals: [] };
}

export const QueryArchivedProposalsResponse = {
  encode(message: QueryArchivedProposalsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.proposals) {
      Proposal.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryArchivedProposalsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryArchivedProposalsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposals.push(Proposal.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryArchivedProposalsResponse {
    return {
      proposals: Array.isArray(object?.proposals) ? object.proposals.map((e: any) => Proposal.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryArchivedProposalsResponse): unknown {
    const obj: any = {};
    if (message.proposals) {
      obj.proposals = message.proposals.map((e) => e ? Proposal.toJSON(e) : undefined);
    } else {
      obj.proposals = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryArchivedProposalsResponse>, I>>(
    object: I,
  ): QueryArchivedProposalsResponse {
    const message = createBaseQueryArchivedProposalsResponse();
    message.proposals = object.proposals?.map((e) => Proposal.fromPartial(e)) || [];
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a list of admins items. */
  Admins(request: QueryAdminsRequest): Promise<QueryAdminsResponse>;
  /** Queries a list of archived proposals. */
  ArchivedProposals(request: QueryArchivedProposalsRequest): Promise<QueryArchivedProposalsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Admins = this.Admins.bind(this);
    this.ArchivedProposals = this.ArchivedProposals.bind(this);
  }
  Admins(request: QueryAdminsRequest): Promise<QueryAdminsResponse> {
    const data = QueryAdminsRequest.encode(request).finish();
    const promise = this.rpc.request("cosmos.adminmodule.adminmodule.Query", "Admins", data);
    return promise.then((data) => QueryAdminsResponse.decode(new _m0.Reader(data)));
  }

  ArchivedProposals(request: QueryArchivedProposalsRequest): Promise<QueryArchivedProposalsResponse> {
    const data = QueryArchivedProposalsRequest.encode(request).finish();
    const promise = this.rpc.request("cosmos.adminmodule.adminmodule.Query", "ArchivedProposals", data);
    return promise.then((data) => QueryArchivedProposalsResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };
