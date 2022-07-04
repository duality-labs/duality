/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Tick {
  token0: string;
  token1: string;
  price: string;
  fee: number;
  reserves0: number;
  reserves1: number;
  totalShares: number;
}

const baseTick: object = {
  token0: "",
  token1: "",
  price: "",
  fee: 0,
  reserves0: 0,
  reserves1: 0,
  totalShares: 0,
};

export const Tick = {
  encode(message: Tick, writer: Writer = Writer.create()): Writer {
    if (message.token0 !== "") {
      writer.uint32(10).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(18).string(message.token1);
    }
    if (message.price !== "") {
      writer.uint32(26).string(message.price);
    }
    if (message.fee !== 0) {
      writer.uint32(32).uint64(message.fee);
    }
    if (message.reserves0 !== 0) {
      writer.uint32(40).uint64(message.reserves0);
    }
    if (message.reserves1 !== 0) {
      writer.uint32(48).uint64(message.reserves1);
    }
    if (message.totalShares !== 0) {
      writer.uint32(56).uint64(message.totalShares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Tick {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTick } as Tick;
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
          message.price = reader.string();
          break;
        case 4:
          message.fee = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.reserves0 = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.reserves1 = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.totalShares = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Tick {
    const message = { ...baseTick } as Tick;
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
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = Number(object.fee);
    } else {
      message.fee = 0;
    }
    if (object.reserves0 !== undefined && object.reserves0 !== null) {
      message.reserves0 = Number(object.reserves0);
    } else {
      message.reserves0 = 0;
    }
    if (object.reserves1 !== undefined && object.reserves1 !== null) {
      message.reserves1 = Number(object.reserves1);
    } else {
      message.reserves1 = 0;
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = Number(object.totalShares);
    } else {
      message.totalShares = 0;
    }
    return message;
  },

  toJSON(message: Tick): unknown {
    const obj: any = {};
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.reserves0 !== undefined && (obj.reserves0 = message.reserves0);
    message.reserves1 !== undefined && (obj.reserves1 = message.reserves1);
    message.totalShares !== undefined &&
      (obj.totalShares = message.totalShares);
    return obj;
  },

  fromPartial(object: DeepPartial<Tick>): Tick {
    const message = { ...baseTick } as Tick;
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
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = 0;
    }
    if (object.reserves0 !== undefined && object.reserves0 !== null) {
      message.reserves0 = object.reserves0;
    } else {
      message.reserves0 = 0;
    }
    if (object.reserves1 !== undefined && object.reserves1 !== null) {
      message.reserves1 = object.reserves1;
    } else {
      message.reserves1 = 0;
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = object.totalShares;
    } else {
      message.totalShares = 0;
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
