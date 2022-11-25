/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderTranche {
  pairId: string;
  tokenIn: string;
  tickIndex: number;
  trancheIndex: number;
  reservesTokenIn: string;
  reservesTokenOut: string;
  totalTokenIn: string;
  totalTokenOut: string;
}

function createBaseLimitOrderTranche(): LimitOrderTranche {
  return {
    pairId: "",
    tokenIn: "",
    tickIndex: 0,
    trancheIndex: 0,
    reservesTokenIn: "",
    reservesTokenOut: "",
    totalTokenIn: "",
    totalTokenOut: "",
  };
}

export const LimitOrderTranche = {
  encode(message: LimitOrderTranche, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tokenIn !== "") {
      writer.uint32(18).string(message.tokenIn);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(24).int64(message.tickIndex);
    }
    if (message.trancheIndex !== 0) {
      writer.uint32(32).uint64(message.trancheIndex);
    }
    if (message.reservesTokenIn !== "") {
      writer.uint32(42).string(message.reservesTokenIn);
    }
    if (message.reservesTokenOut !== "") {
      writer.uint32(50).string(message.reservesTokenOut);
    }
    if (message.totalTokenIn !== "") {
      writer.uint32(58).string(message.totalTokenIn);
    }
    if (message.totalTokenOut !== "") {
      writer.uint32(66).string(message.totalTokenOut);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LimitOrderTranche {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLimitOrderTranche();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairId = reader.string();
          break;
        case 2:
          message.tokenIn = reader.string();
          break;
        case 3:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.trancheIndex = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.reservesTokenIn = reader.string();
          break;
        case 6:
          message.reservesTokenOut = reader.string();
          break;
        case 7:
          message.totalTokenIn = reader.string();
          break;
        case 8:
          message.totalTokenOut = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderTranche {
    return {
      pairId: isSet(object.pairId) ? String(object.pairId) : "",
      tokenIn: isSet(object.tokenIn) ? String(object.tokenIn) : "",
      tickIndex: isSet(object.tickIndex) ? Number(object.tickIndex) : 0,
      trancheIndex: isSet(object.trancheIndex) ? Number(object.trancheIndex) : 0,
      reservesTokenIn: isSet(object.reservesTokenIn) ? String(object.reservesTokenIn) : "",
      reservesTokenOut: isSet(object.reservesTokenOut) ? String(object.reservesTokenOut) : "",
      totalTokenIn: isSet(object.totalTokenIn) ? String(object.totalTokenIn) : "",
      totalTokenOut: isSet(object.totalTokenOut) ? String(object.totalTokenOut) : "",
    };
  },

  toJSON(message: LimitOrderTranche): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tokenIn !== undefined && (obj.tokenIn = message.tokenIn);
    message.tickIndex !== undefined && (obj.tickIndex = Math.round(message.tickIndex));
    message.trancheIndex !== undefined && (obj.trancheIndex = Math.round(message.trancheIndex));
    message.reservesTokenIn !== undefined && (obj.reservesTokenIn = message.reservesTokenIn);
    message.reservesTokenOut !== undefined && (obj.reservesTokenOut = message.reservesTokenOut);
    message.totalTokenIn !== undefined && (obj.totalTokenIn = message.totalTokenIn);
    message.totalTokenOut !== undefined && (obj.totalTokenOut = message.totalTokenOut);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LimitOrderTranche>, I>>(object: I): LimitOrderTranche {
    const message = createBaseLimitOrderTranche();
    message.pairId = object.pairId ?? "";
    message.tokenIn = object.tokenIn ?? "";
    message.tickIndex = object.tickIndex ?? 0;
    message.trancheIndex = object.trancheIndex ?? 0;
    message.reservesTokenIn = object.reservesTokenIn ?? "";
    message.reservesTokenOut = object.reservesTokenOut ?? "";
    message.totalTokenIn = object.totalTokenIn ?? "";
    message.totalTokenOut = object.totalTokenOut ?? "";
    return message;
  },
};

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
