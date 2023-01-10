/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderTrancheTrancheIndexes {
  fillTrancheIndex: number;
  placeTrancheIndex: number;
}

const baseLimitOrderTrancheTrancheIndexes: object = {
  fillTrancheIndex: 0,
  placeTrancheIndex: 0,
};

export const LimitOrderTrancheTrancheIndexes = {
  encode(
    message: LimitOrderTrancheTrancheIndexes,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fillTrancheIndex !== 0) {
      writer.uint32(8).uint64(message.fillTrancheIndex);
    }
    if (message.placeTrancheIndex !== 0) {
      writer.uint32(16).uint64(message.placeTrancheIndex);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): LimitOrderTrancheTrancheIndexes {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseLimitOrderTrancheTrancheIndexes,
    } as LimitOrderTrancheTrancheIndexes;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fillTrancheIndex = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.placeTrancheIndex = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderTrancheTrancheIndexes {
    const message = {
      ...baseLimitOrderTrancheTrancheIndexes,
    } as LimitOrderTrancheTrancheIndexes;
    if (
      object.fillTrancheIndex !== undefined &&
      object.fillTrancheIndex !== null
    ) {
      message.fillTrancheIndex = Number(object.fillTrancheIndex);
    } else {
      message.fillTrancheIndex = 0;
    }
    if (
      object.placeTrancheIndex !== undefined &&
      object.placeTrancheIndex !== null
    ) {
      message.placeTrancheIndex = Number(object.placeTrancheIndex);
    } else {
      message.placeTrancheIndex = 0;
    }
    return message;
  },

  toJSON(message: LimitOrderTrancheTrancheIndexes): unknown {
    const obj: any = {};
    message.fillTrancheIndex !== undefined &&
      (obj.fillTrancheIndex = message.fillTrancheIndex);
    message.placeTrancheIndex !== undefined &&
      (obj.placeTrancheIndex = message.placeTrancheIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<LimitOrderTrancheTrancheIndexes>
  ): LimitOrderTrancheTrancheIndexes {
    const message = {
      ...baseLimitOrderTrancheTrancheIndexes,
    } as LimitOrderTrancheTrancheIndexes;
    if (
      object.fillTrancheIndex !== undefined &&
      object.fillTrancheIndex !== null
    ) {
      message.fillTrancheIndex = object.fillTrancheIndex;
    } else {
      message.fillTrancheIndex = 0;
    }
    if (
      object.placeTrancheIndex !== undefined &&
      object.placeTrancheIndex !== null
    ) {
      message.placeTrancheIndex = object.placeTrancheIndex;
    } else {
      message.placeTrancheIndex = 0;
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
