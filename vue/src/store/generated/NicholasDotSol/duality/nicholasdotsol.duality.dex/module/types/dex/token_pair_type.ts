/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TokenPairType {
  currentTick0To1: number;
  currentTick1To0: number;
}

const baseTokenPairType: object = { currentTick0To1: 0, currentTick1To0: 0 };

export const TokenPairType = {
  encode(message: TokenPairType, writer: Writer = Writer.create()): Writer {
    if (message.currentTick0To1 !== 0) {
      writer.uint32(8).int64(message.currentTick0To1);
    }
    if (message.currentTick1To0 !== 0) {
      writer.uint32(16).int64(message.currentTick1To0);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TokenPairType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTokenPairType } as TokenPairType;
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
    const message = { ...baseTokenPairType } as TokenPairType;
    if (
      object.currentTick0To1 !== undefined &&
      object.currentTick0To1 !== null
    ) {
      message.currentTick0To1 = Number(object.currentTick0To1);
    } else {
      message.currentTick0To1 = 0;
    }
    if (
      object.currentTick1To0 !== undefined &&
      object.currentTick1To0 !== null
    ) {
      message.currentTick1To0 = Number(object.currentTick1To0);
    } else {
      message.currentTick1To0 = 0;
    }
    return message;
  },

  toJSON(message: TokenPairType): unknown {
    const obj: any = {};
    message.currentTick0To1 !== undefined &&
      (obj.currentTick0To1 = message.currentTick0To1);
    message.currentTick1To0 !== undefined &&
      (obj.currentTick1To0 = message.currentTick1To0);
    return obj;
  },

  fromPartial(object: DeepPartial<TokenPairType>): TokenPairType {
    const message = { ...baseTokenPairType } as TokenPairType;
    if (
      object.currentTick0To1 !== undefined &&
      object.currentTick0To1 !== null
    ) {
      message.currentTick0To1 = object.currentTick0To1;
    } else {
      message.currentTick0To1 = 0;
    }
    if (
      object.currentTick1To0 !== undefined &&
      object.currentTick1To0 !== null
    ) {
      message.currentTick1To0 = object.currentTick1To0;
    } else {
      message.currentTick1To0 = 0;
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
