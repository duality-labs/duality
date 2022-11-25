/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderTrancheUser {
  pairId: string;
  token: string;
  tickIndex: number;
  count: number;
  address: string;
  sharesOwned: string;
  sharesWithdrawn: string;
  sharesCancelled: string;
}

function createBaseLimitOrderTrancheUser(): LimitOrderTrancheUser {
  return {
    pairId: "",
    token: "",
    tickIndex: 0,
    count: 0,
    address: "",
    sharesOwned: "",
    sharesWithdrawn: "",
    sharesCancelled: "",
  };
}

export const LimitOrderTrancheUser = {
  encode(message: LimitOrderTrancheUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.token !== "") {
      writer.uint32(18).string(message.token);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(24).int64(message.tickIndex);
    }
    if (message.count !== 0) {
      writer.uint32(32).uint64(message.count);
    }
    if (message.address !== "") {
      writer.uint32(42).string(message.address);
    }
    if (message.sharesOwned !== "") {
      writer.uint32(50).string(message.sharesOwned);
    }
    if (message.sharesWithdrawn !== "") {
      writer.uint32(58).string(message.sharesWithdrawn);
    }
    if (message.sharesCancelled !== "") {
      writer.uint32(66).string(message.sharesCancelled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LimitOrderTrancheUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLimitOrderTrancheUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairId = reader.string();
          break;
        case 2:
          message.token = reader.string();
          break;
        case 3:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.address = reader.string();
          break;
        case 6:
          message.sharesOwned = reader.string();
          break;
        case 7:
          message.sharesWithdrawn = reader.string();
          break;
        case 8:
          message.sharesCancelled = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderTrancheUser {
    return {
      pairId: isSet(object.pairId) ? String(object.pairId) : "",
      token: isSet(object.token) ? String(object.token) : "",
      tickIndex: isSet(object.tickIndex) ? Number(object.tickIndex) : 0,
      count: isSet(object.count) ? Number(object.count) : 0,
      address: isSet(object.address) ? String(object.address) : "",
      sharesOwned: isSet(object.sharesOwned) ? String(object.sharesOwned) : "",
      sharesWithdrawn: isSet(object.sharesWithdrawn) ? String(object.sharesWithdrawn) : "",
      sharesCancelled: isSet(object.sharesCancelled) ? String(object.sharesCancelled) : "",
    };
  },

  toJSON(message: LimitOrderTrancheUser): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.token !== undefined && (obj.token = message.token);
    message.tickIndex !== undefined && (obj.tickIndex = Math.round(message.tickIndex));
    message.count !== undefined && (obj.count = Math.round(message.count));
    message.address !== undefined && (obj.address = message.address);
    message.sharesOwned !== undefined && (obj.sharesOwned = message.sharesOwned);
    message.sharesWithdrawn !== undefined && (obj.sharesWithdrawn = message.sharesWithdrawn);
    message.sharesCancelled !== undefined && (obj.sharesCancelled = message.sharesCancelled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LimitOrderTrancheUser>, I>>(object: I): LimitOrderTrancheUser {
    const message = createBaseLimitOrderTrancheUser();
    message.pairId = object.pairId ?? "";
    message.token = object.token ?? "";
    message.tickIndex = object.tickIndex ?? 0;
    message.count = object.count ?? 0;
    message.address = object.address ?? "";
    message.sharesOwned = object.sharesOwned ?? "";
    message.sharesWithdrawn = object.sharesWithdrawn ?? "";
    message.sharesCancelled = object.sharesCancelled ?? "";
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
