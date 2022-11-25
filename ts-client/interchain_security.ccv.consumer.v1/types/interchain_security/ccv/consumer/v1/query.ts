/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "interchain_security.ccv.consumer.v1";

/** NextFeeDistributionEstimate holds information about next fee distribution */
export interface NextFeeDistributionEstimate {
  /** current block height at the time of querying */
  currentHeight: number;
  /** block height at which last distribution took place */
  lastHeight: number;
  /** block height at which next distribution will take place */
  nextHeight: number;
  /** ratio between consumer and provider fee distribution */
  distributionFraction: string;
  /** total accruead fees at the time of querying */
  total: string;
  /** amount distibuted to provider chain */
  toProvider: string;
  /** amount distributed (kept) by consumer chain */
  toConsumer: string;
}

export interface QueryNextFeeDistributionEstimateRequest {
}

export interface QueryNextFeeDistributionEstimateResponse {
  data: NextFeeDistributionEstimate | undefined;
}

function createBaseNextFeeDistributionEstimate(): NextFeeDistributionEstimate {
  return {
    currentHeight: 0,
    lastHeight: 0,
    nextHeight: 0,
    distributionFraction: "",
    total: "",
    toProvider: "",
    toConsumer: "",
  };
}

export const NextFeeDistributionEstimate = {
  encode(message: NextFeeDistributionEstimate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.currentHeight !== 0) {
      writer.uint32(8).int64(message.currentHeight);
    }
    if (message.lastHeight !== 0) {
      writer.uint32(16).int64(message.lastHeight);
    }
    if (message.nextHeight !== 0) {
      writer.uint32(24).int64(message.nextHeight);
    }
    if (message.distributionFraction !== "") {
      writer.uint32(34).string(message.distributionFraction);
    }
    if (message.total !== "") {
      writer.uint32(42).string(message.total);
    }
    if (message.toProvider !== "") {
      writer.uint32(50).string(message.toProvider);
    }
    if (message.toConsumer !== "") {
      writer.uint32(58).string(message.toConsumer);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NextFeeDistributionEstimate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNextFeeDistributionEstimate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.currentHeight = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.lastHeight = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.nextHeight = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.distributionFraction = reader.string();
          break;
        case 5:
          message.total = reader.string();
          break;
        case 6:
          message.toProvider = reader.string();
          break;
        case 7:
          message.toConsumer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NextFeeDistributionEstimate {
    return {
      currentHeight: isSet(object.currentHeight) ? Number(object.currentHeight) : 0,
      lastHeight: isSet(object.lastHeight) ? Number(object.lastHeight) : 0,
      nextHeight: isSet(object.nextHeight) ? Number(object.nextHeight) : 0,
      distributionFraction: isSet(object.distributionFraction) ? String(object.distributionFraction) : "",
      total: isSet(object.total) ? String(object.total) : "",
      toProvider: isSet(object.toProvider) ? String(object.toProvider) : "",
      toConsumer: isSet(object.toConsumer) ? String(object.toConsumer) : "",
    };
  },

  toJSON(message: NextFeeDistributionEstimate): unknown {
    const obj: any = {};
    message.currentHeight !== undefined && (obj.currentHeight = Math.round(message.currentHeight));
    message.lastHeight !== undefined && (obj.lastHeight = Math.round(message.lastHeight));
    message.nextHeight !== undefined && (obj.nextHeight = Math.round(message.nextHeight));
    message.distributionFraction !== undefined && (obj.distributionFraction = message.distributionFraction);
    message.total !== undefined && (obj.total = message.total);
    message.toProvider !== undefined && (obj.toProvider = message.toProvider);
    message.toConsumer !== undefined && (obj.toConsumer = message.toConsumer);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NextFeeDistributionEstimate>, I>>(object: I): NextFeeDistributionEstimate {
    const message = createBaseNextFeeDistributionEstimate();
    message.currentHeight = object.currentHeight ?? 0;
    message.lastHeight = object.lastHeight ?? 0;
    message.nextHeight = object.nextHeight ?? 0;
    message.distributionFraction = object.distributionFraction ?? "";
    message.total = object.total ?? "";
    message.toProvider = object.toProvider ?? "";
    message.toConsumer = object.toConsumer ?? "";
    return message;
  },
};

function createBaseQueryNextFeeDistributionEstimateRequest(): QueryNextFeeDistributionEstimateRequest {
  return {};
}

export const QueryNextFeeDistributionEstimateRequest = {
  encode(_: QueryNextFeeDistributionEstimateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryNextFeeDistributionEstimateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryNextFeeDistributionEstimateRequest();
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

  fromJSON(_: any): QueryNextFeeDistributionEstimateRequest {
    return {};
  },

  toJSON(_: QueryNextFeeDistributionEstimateRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryNextFeeDistributionEstimateRequest>, I>>(
    _: I,
  ): QueryNextFeeDistributionEstimateRequest {
    const message = createBaseQueryNextFeeDistributionEstimateRequest();
    return message;
  },
};

function createBaseQueryNextFeeDistributionEstimateResponse(): QueryNextFeeDistributionEstimateResponse {
  return { data: undefined };
}

export const QueryNextFeeDistributionEstimateResponse = {
  encode(message: QueryNextFeeDistributionEstimateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.data !== undefined) {
      NextFeeDistributionEstimate.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryNextFeeDistributionEstimateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryNextFeeDistributionEstimateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = NextFeeDistributionEstimate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryNextFeeDistributionEstimateResponse {
    return { data: isSet(object.data) ? NextFeeDistributionEstimate.fromJSON(object.data) : undefined };
  },

  toJSON(message: QueryNextFeeDistributionEstimateResponse): unknown {
    const obj: any = {};
    message.data !== undefined
      && (obj.data = message.data ? NextFeeDistributionEstimate.toJSON(message.data) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryNextFeeDistributionEstimateResponse>, I>>(
    object: I,
  ): QueryNextFeeDistributionEstimateResponse {
    const message = createBaseQueryNextFeeDistributionEstimateResponse();
    message.data = (object.data !== undefined && object.data !== null)
      ? NextFeeDistributionEstimate.fromPartial(object.data)
      : undefined;
    return message;
  },
};

export interface Query {
  /**
   * ConsumerGenesis queries the genesis state needed to start a consumer chain
   * whose proposal has been accepted
   */
  QueryNextFeeDistribution(
    request: QueryNextFeeDistributionEstimateRequest,
  ): Promise<QueryNextFeeDistributionEstimateResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.QueryNextFeeDistribution = this.QueryNextFeeDistribution.bind(this);
  }
  QueryNextFeeDistribution(
    request: QueryNextFeeDistributionEstimateRequest,
  ): Promise<QueryNextFeeDistributionEstimateResponse> {
    const data = QueryNextFeeDistributionEstimateRequest.encode(request).finish();
    const promise = this.rpc.request("interchain_security.ccv.consumer.v1.Query", "QueryNextFeeDistribution", data);
    return promise.then((data) => QueryNextFeeDistributionEstimateResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
