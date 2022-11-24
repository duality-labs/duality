/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

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

const baseLimitOrderTranche: object = {
  pairId: "",
  tokenIn: "",
  tickIndex: 0,
  trancheIndex: 0,
  reservesTokenIn: "",
  reservesTokenOut: "",
  totalTokenIn: "",
  totalTokenOut: "",
};

export const LimitOrderTranche = {
  encode(message: LimitOrderTranche, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): LimitOrderTranche {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLimitOrderTranche } as LimitOrderTranche;
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
    const message = { ...baseLimitOrderTranche } as LimitOrderTranche;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tokenIn !== undefined && object.tokenIn !== null) {
      message.tokenIn = String(object.tokenIn);
    } else {
      message.tokenIn = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.trancheIndex !== undefined && object.trancheIndex !== null) {
      message.trancheIndex = Number(object.trancheIndex);
    } else {
      message.trancheIndex = 0;
    }
    if (
      object.reservesTokenIn !== undefined &&
      object.reservesTokenIn !== null
    ) {
      message.reservesTokenIn = String(object.reservesTokenIn);
    } else {
      message.reservesTokenIn = "";
    }
    if (
      object.reservesTokenOut !== undefined &&
      object.reservesTokenOut !== null
    ) {
      message.reservesTokenOut = String(object.reservesTokenOut);
    } else {
      message.reservesTokenOut = "";
    }
    if (object.totalTokenIn !== undefined && object.totalTokenIn !== null) {
      message.totalTokenIn = String(object.totalTokenIn);
    } else {
      message.totalTokenIn = "";
    }
    if (object.totalTokenOut !== undefined && object.totalTokenOut !== null) {
      message.totalTokenOut = String(object.totalTokenOut);
    } else {
      message.totalTokenOut = "";
    }
    return message;
  },

  toJSON(message: LimitOrderTranche): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tokenIn !== undefined && (obj.tokenIn = message.tokenIn);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.trancheIndex !== undefined &&
      (obj.trancheIndex = message.trancheIndex);
    message.reservesTokenIn !== undefined &&
      (obj.reservesTokenIn = message.reservesTokenIn);
    message.reservesTokenOut !== undefined &&
      (obj.reservesTokenOut = message.reservesTokenOut);
    message.totalTokenIn !== undefined &&
      (obj.totalTokenIn = message.totalTokenIn);
    message.totalTokenOut !== undefined &&
      (obj.totalTokenOut = message.totalTokenOut);
    return obj;
  },

  fromPartial(object: DeepPartial<LimitOrderTranche>): LimitOrderTranche {
    const message = { ...baseLimitOrderTranche } as LimitOrderTranche;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tokenIn !== undefined && object.tokenIn !== null) {
      message.tokenIn = object.tokenIn;
    } else {
      message.tokenIn = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.trancheIndex !== undefined && object.trancheIndex !== null) {
      message.trancheIndex = object.trancheIndex;
    } else {
      message.trancheIndex = 0;
    }
    if (
      object.reservesTokenIn !== undefined &&
      object.reservesTokenIn !== null
    ) {
      message.reservesTokenIn = object.reservesTokenIn;
    } else {
      message.reservesTokenIn = "";
    }
    if (
      object.reservesTokenOut !== undefined &&
      object.reservesTokenOut !== null
    ) {
      message.reservesTokenOut = object.reservesTokenOut;
    } else {
      message.reservesTokenOut = "";
    }
    if (object.totalTokenIn !== undefined && object.totalTokenIn !== null) {
      message.totalTokenIn = object.totalTokenIn;
    } else {
      message.totalTokenIn = "";
    }
    if (object.totalTokenOut !== undefined && object.totalTokenOut !== null) {
      message.totalTokenOut = object.totalTokenOut;
    } else {
      message.totalTokenOut = "";
    }
    return message;
  },
};

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
