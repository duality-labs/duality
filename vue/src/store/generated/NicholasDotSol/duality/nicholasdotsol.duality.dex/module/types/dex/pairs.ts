/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Ticks } from "../dex/ticks";
import { VirtualPriceQueue } from "../dex/virtual_price_queue";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Pairs {
  token0: string;
  token1: string;
  tickSpacing: number;
  currentPrice: string;
  bitArray: Uint8Array[];
  tickmap: Ticks | undefined;
  virtualPriceMap: VirtualPriceQueue | undefined;
}

const basePairs: object = {
  token0: "",
  token1: "",
  tickSpacing: 0,
  currentPrice: "",
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
      writer.uint32(24).uint64(message.tickSpacing);
    }
    if (message.currentPrice !== "") {
      writer.uint32(34).string(message.currentPrice);
    }
    for (const v of message.bitArray) {
      writer.uint32(42).bytes(v!);
    }
    if (message.tickmap !== undefined) {
      Ticks.encode(message.tickmap, writer.uint32(50).fork()).ldelim();
    }
    if (message.virtualPriceMap !== undefined) {
      VirtualPriceQueue.encode(
        message.virtualPriceMap,
        writer.uint32(58).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pairs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePairs } as Pairs;
    message.bitArray = [];
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
          message.tickSpacing = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.currentPrice = reader.string();
          break;
        case 5:
          message.bitArray.push(reader.bytes());
          break;
        case 6:
          message.tickmap = Ticks.decode(reader, reader.uint32());
          break;
        case 7:
          message.virtualPriceMap = VirtualPriceQueue.decode(
            reader,
            reader.uint32()
          );
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
    message.bitArray = [];
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
    if (object.currentPrice !== undefined && object.currentPrice !== null) {
      message.currentPrice = String(object.currentPrice);
    } else {
      message.currentPrice = "";
    }
    if (object.bitArray !== undefined && object.bitArray !== null) {
      for (const e of object.bitArray) {
        message.bitArray.push(bytesFromBase64(e));
      }
    }
    if (object.tickmap !== undefined && object.tickmap !== null) {
      message.tickmap = Ticks.fromJSON(object.tickmap);
    } else {
      message.tickmap = undefined;
    }
    if (
      object.virtualPriceMap !== undefined &&
      object.virtualPriceMap !== null
    ) {
      message.virtualPriceMap = VirtualPriceQueue.fromJSON(
        object.virtualPriceMap
      );
    } else {
      message.virtualPriceMap = undefined;
    }
    return message;
  },

  toJSON(message: Pairs): unknown {
    const obj: any = {};
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.tickSpacing !== undefined &&
      (obj.tickSpacing = message.tickSpacing);
    message.currentPrice !== undefined &&
      (obj.currentPrice = message.currentPrice);
    if (message.bitArray) {
      obj.bitArray = message.bitArray.map((e) =>
        base64FromBytes(e !== undefined ? e : new Uint8Array())
      );
    } else {
      obj.bitArray = [];
    }
    message.tickmap !== undefined &&
      (obj.tickmap = message.tickmap
        ? Ticks.toJSON(message.tickmap)
        : undefined);
    message.virtualPriceMap !== undefined &&
      (obj.virtualPriceMap = message.virtualPriceMap
        ? VirtualPriceQueue.toJSON(message.virtualPriceMap)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Pairs>): Pairs {
    const message = { ...basePairs } as Pairs;
    message.bitArray = [];
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
    if (object.currentPrice !== undefined && object.currentPrice !== null) {
      message.currentPrice = object.currentPrice;
    } else {
      message.currentPrice = "";
    }
    if (object.bitArray !== undefined && object.bitArray !== null) {
      for (const e of object.bitArray) {
        message.bitArray.push(e);
      }
    }
    if (object.tickmap !== undefined && object.tickmap !== null) {
      message.tickmap = Ticks.fromPartial(object.tickmap);
    } else {
      message.tickmap = undefined;
    }
    if (
      object.virtualPriceMap !== undefined &&
      object.virtualPriceMap !== null
    ) {
      message.virtualPriceMap = VirtualPriceQueue.fromPartial(
        object.virtualPriceMap
      );
    } else {
      message.virtualPriceMap = undefined;
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

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

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
