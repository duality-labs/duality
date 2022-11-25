/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TokenPairType {
  currentTick0To1: number;
  currentTick1To0: number;
}

function createBaseTokenPairType(): TokenPairType {
  return { currentTick0To1: 0, currentTick1To0: 0 };
}

export const TokenPairType = {
  encode(message: TokenPairType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.currentTick0To1 !== 0) {
      writer.uint32(8).int64(message.currentTick0To1);
    }
    if (message.currentTick1To0 !== 0) {
      writer.uint32(16).int64(message.currentTick1To0);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenPairType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTokenPairType();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.currentTick0To1 = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.currentTick1To0 = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenPairType {
    return {
      currentTick0To1: isSet(object.currentTick0To1) ? Number(object.currentTick0To1) : 0,
      currentTick1To0: isSet(object.currentTick1To0) ? Number(object.currentTick1To0) : 0,
    };
  },

  toJSON(message: TokenPairType): unknown {
    const obj: any = {};
    message.currentTick0To1 !== undefined && (obj.currentTick0To1 = Math.round(message.currentTick0To1));
    message.currentTick1To0 !== undefined && (obj.currentTick1To0 = Math.round(message.currentTick1To0));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TokenPairType>, I>>(object: I): TokenPairType {
    const message = createBaseTokenPairType();
    message.currentTick0To1 = object.currentTick0To1 ?? 0;
    message.currentTick1To0 = object.currentTick1To0 ?? 0;
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
