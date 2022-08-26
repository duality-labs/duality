/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Pairs {
  token0: string;
  token1: string;
  tickSpacing: number;
  currentIndex: number;
}

const basePairs: object = {
  token0: "",
  token1: "",
  tickSpacing: 0,
  currentIndex: 0,
};

export const Pairs = {
  encode(message: Pairs, writer: Writer = Writer.create()): Writer {
    if (message.token0 !== "") {
      writer.uint32(10).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(18).string(message.token1);
    }
    if (message.tickSpacing !== 0) {
      writer.uint32(24).int64(message.tickSpacing);
    }
    if (message.currentIndex !== 0) {
      writer.uint32(32).int32(message.currentIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pairs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePairs } as Pairs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token0 = reader.string();
          break;
        case 2:
          message.token1 = reader.string();
          break;
        case 3:
          message.tickSpacing = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.currentIndex = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pairs {
    const message = { ...basePairs } as Pairs;
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
    if (object.tickSpacing !== undefined && object.tickSpacing !== null) {
      message.tickSpacing = Number(object.tickSpacing);
    } else {
      message.tickSpacing = 0;
    }
    if (object.currentIndex !== undefined && object.currentIndex !== null) {
      message.currentIndex = Number(object.currentIndex);
    } else {
      message.currentIndex = 0;
    }
    return message;
  },

  toJSON(message: Pairs): unknown {
    const obj: any = {};
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.tickSpacing !== undefined &&
      (obj.tickSpacing = message.tickSpacing);
    message.currentIndex !== undefined &&
      (obj.currentIndex = message.currentIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<Pairs>): Pairs {
    const message = { ...basePairs } as Pairs;
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
    if (object.tickSpacing !== undefined && object.tickSpacing !== null) {
      message.tickSpacing = object.tickSpacing;
    } else {
      message.tickSpacing = 0;
    }
    if (object.currentIndex !== undefined && object.currentIndex !== null) {
      message.currentIndex = object.currentIndex;
    } else {
      message.currentIndex = 0;
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
