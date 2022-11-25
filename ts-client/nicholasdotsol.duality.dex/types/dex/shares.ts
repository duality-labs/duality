/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Shares {
  address: string;
  pairId: string;
  tickIndex: number;
  feeIndex: number;
  sharesOwned: string;
}

function createBaseShares(): Shares {
  return { address: "", pairId: "", tickIndex: 0, feeIndex: 0, sharesOwned: "" };
}

export const Shares = {
  encode(message: Shares, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.pairId !== "") {
      writer.uint32(18).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(24).int64(message.tickIndex);
    }
    if (message.feeIndex !== 0) {
      writer.uint32(32).uint64(message.feeIndex);
    }
    if (message.sharesOwned !== "") {
      writer.uint32(42).string(message.sharesOwned);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Shares {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseShares();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.pairId = reader.string();
          break;
        case 3:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.feeIndex = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.sharesOwned = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Shares {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      pairId: isSet(object.pairId) ? String(object.pairId) : "",
      tickIndex: isSet(object.tickIndex) ? Number(object.tickIndex) : 0,
      feeIndex: isSet(object.feeIndex) ? Number(object.feeIndex) : 0,
      sharesOwned: isSet(object.sharesOwned) ? String(object.sharesOwned) : "",
    };
  },

  toJSON(message: Shares): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = Math.round(message.tickIndex));
    message.feeIndex !== undefined && (obj.feeIndex = Math.round(message.feeIndex));
    message.sharesOwned !== undefined && (obj.sharesOwned = message.sharesOwned);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Shares>, I>>(object: I): Shares {
    const message = createBaseShares();
    message.address = object.address ?? "";
    message.pairId = object.pairId ?? "";
    message.tickIndex = object.tickIndex ?? 0;
    message.feeIndex = object.feeIndex ?? 0;
    message.sharesOwned = object.sharesOwned ?? "";
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
