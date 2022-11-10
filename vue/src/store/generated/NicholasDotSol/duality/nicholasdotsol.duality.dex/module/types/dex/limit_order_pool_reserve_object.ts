/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderPoolReserveObject {
  pairId: string;
  token: string;
  tickIndex: number;
  count: number;
  reserves: string;
}

const baseLimitOrderPoolReserveObject: object = {
  pairId: "",
  token: "",
  tickIndex: 0,
  count: 0,
  reserves: "",
};

export const LimitOrderPoolReserveObject = {
  encode(
    message: LimitOrderPoolReserveObject,
    writer: Writer = Writer.create()
  ): Writer {
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
    if (message.reserves !== "") {
      writer.uint32(42).string(message.reserves);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): LimitOrderPoolReserveObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseLimitOrderPoolReserveObject,
    } as LimitOrderPoolReserveObject;
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
          message.reserves = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderPoolReserveObject {
    const message = {
      ...baseLimitOrderPoolReserveObject,
    } as LimitOrderPoolReserveObject;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.reserves !== undefined && object.reserves !== null) {
      message.reserves = String(object.reserves);
    } else {
      message.reserves = "";
    }
    return message;
  },

  toJSON(message: LimitOrderPoolReserveObject): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.token !== undefined && (obj.token = message.token);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.count !== undefined && (obj.count = message.count);
    message.reserves !== undefined && (obj.reserves = message.reserves);
    return obj;
  },

  fromPartial(
    object: DeepPartial<LimitOrderPoolReserveObject>
  ): LimitOrderPoolReserveObject {
    const message = {
      ...baseLimitOrderPoolReserveObject,
    } as LimitOrderPoolReserveObject;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.reserves !== undefined && object.reserves !== null) {
      message.reserves = object.reserves;
    } else {
      message.reserves = "";
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
