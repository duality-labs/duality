/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderPoolFillMap {
  count: number;
  fill: string;
}

const baseLimitOrderPoolFillMap: object = { count: 0, fill: "" };

export const LimitOrderPoolFillMap = {
  encode(
    message: LimitOrderPoolFillMap,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.fill !== "") {
      writer.uint32(18).string(message.fill);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LimitOrderPoolFillMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLimitOrderPoolFillMap } as LimitOrderPoolFillMap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.fill = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderPoolFillMap {
    const message = { ...baseLimitOrderPoolFillMap } as LimitOrderPoolFillMap;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.fill !== undefined && object.fill !== null) {
      message.fill = String(object.fill);
    } else {
      message.fill = "";
    }
    return message;
  },

  toJSON(message: LimitOrderPoolFillMap): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.fill !== undefined && (obj.fill = message.fill);
    return obj;
  },

  fromPartial(
    object: DeepPartial<LimitOrderPoolFillMap>
  ): LimitOrderPoolFillMap {
    const message = { ...baseLimitOrderPoolFillMap } as LimitOrderPoolFillMap;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.fill !== undefined && object.fill !== null) {
      message.fill = object.fill;
    } else {
      message.fill = "";
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
