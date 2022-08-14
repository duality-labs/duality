/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Pairs {
  token0: string;
  token1: string;
  tickSpacing: string;
  currentIndex: string;
  bitArray: string;
  tickmap: string;
  virtualPricemap: string;
}

const basePairs: object = {
  token0: "",
  token1: "",
  tickSpacing: "",
  currentIndex: "",
  bitArray: "",
  tickmap: "",
  virtualPricemap: "",
};

export const Pairs = {
  encode(message: Pairs, writer: Writer = Writer.create()): Writer {
    if (message.token0 !== "") {
      writer.uint32(10).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(18).string(message.token1);
    }
    if (message.tickSpacing !== "") {
      writer.uint32(26).string(message.tickSpacing);
    }
    if (message.currentIndex !== "") {
      writer.uint32(34).string(message.currentIndex);
    }
    if (message.bitArray !== "") {
      writer.uint32(42).string(message.bitArray);
    }
    if (message.tickmap !== "") {
      writer.uint32(50).string(message.tickmap);
    }
    if (message.virtualPricemap !== "") {
      writer.uint32(58).string(message.virtualPricemap);
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
          message.tickSpacing = reader.string();
          break;
        case 4:
          message.currentIndex = reader.string();
          break;
        case 5:
          message.bitArray = reader.string();
          break;
        case 6:
          message.tickmap = reader.string();
          break;
        case 7:
          message.virtualPricemap = reader.string();
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
      message.tickSpacing = String(object.tickSpacing);
    } else {
      message.tickSpacing = "";
    }
    if (object.currentIndex !== undefined && object.currentIndex !== null) {
      message.currentIndex = String(object.currentIndex);
    } else {
      message.currentIndex = "";
    }
    if (object.bitArray !== undefined && object.bitArray !== null) {
      message.bitArray = String(object.bitArray);
    } else {
      message.bitArray = "";
    }
    if (object.tickmap !== undefined && object.tickmap !== null) {
      message.tickmap = String(object.tickmap);
    } else {
      message.tickmap = "";
    }
    if (
      object.virtualPricemap !== undefined &&
      object.virtualPricemap !== null
    ) {
      message.virtualPricemap = String(object.virtualPricemap);
    } else {
      message.virtualPricemap = "";
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
    message.bitArray !== undefined && (obj.bitArray = message.bitArray);
    message.tickmap !== undefined && (obj.tickmap = message.tickmap);
    message.virtualPricemap !== undefined &&
      (obj.virtualPricemap = message.virtualPricemap);
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
      message.tickSpacing = "";
    }
    if (object.currentIndex !== undefined && object.currentIndex !== null) {
      message.currentIndex = object.currentIndex;
    } else {
      message.currentIndex = "";
    }
    if (object.bitArray !== undefined && object.bitArray !== null) {
      message.bitArray = object.bitArray;
    } else {
      message.bitArray = "";
    }
    if (object.tickmap !== undefined && object.tickmap !== null) {
      message.tickmap = object.tickmap;
    } else {
      message.tickmap = "";
    }
    if (
      object.virtualPricemap !== undefined &&
      object.virtualPricemap !== null
    ) {
      message.virtualPricemap = object.virtualPricemap;
    } else {
      message.virtualPricemap = "";
    }
    return message;
  },
};

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
