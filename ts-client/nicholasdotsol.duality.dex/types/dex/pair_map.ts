/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { TokenPairType } from "./token_pair_type";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface PairMap {
  pairId: string;
  tokenPair: TokenPairType | undefined;
  maxTick: number;
  minTick: number;
}

function createBasePairMap(): PairMap {
  return { pairId: "", tokenPair: undefined, maxTick: 0, minTick: 0 };
}

export const PairMap = {
  encode(message: PairMap, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tokenPair !== undefined) {
      TokenPairType.encode(message.tokenPair, writer.uint32(18).fork()).ldelim();
    }
    if (message.maxTick !== 0) {
      writer.uint32(24).int64(message.maxTick);
    }
    if (message.minTick !== 0) {
      writer.uint32(32).int64(message.minTick);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PairMap {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePairMap();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairId = reader.string();
          break;
        case 2:
          message.tokenPair = TokenPairType.decode(reader, reader.uint32());
          break;
        case 3:
          message.maxTick = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.minTick = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PairMap {
    return {
      pairId: isSet(object.pairId) ? String(object.pairId) : "",
      tokenPair: isSet(object.tokenPair) ? TokenPairType.fromJSON(object.tokenPair) : undefined,
      maxTick: isSet(object.maxTick) ? Number(object.maxTick) : 0,
      minTick: isSet(object.minTick) ? Number(object.minTick) : 0,
    };
  },

  toJSON(message: PairMap): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tokenPair !== undefined
      && (obj.tokenPair = message.tokenPair ? TokenPairType.toJSON(message.tokenPair) : undefined);
    message.maxTick !== undefined && (obj.maxTick = Math.round(message.maxTick));
    message.minTick !== undefined && (obj.minTick = Math.round(message.minTick));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PairMap>, I>>(object: I): PairMap {
    const message = createBasePairMap();
    message.pairId = object.pairId ?? "";
    message.tokenPair = (object.tokenPair !== undefined && object.tokenPair !== null)
      ? TokenPairType.fromPartial(object.tokenPair)
      : undefined;
    message.maxTick = object.maxTick ?? 0;
    message.minTick = object.minTick ?? 0;
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
