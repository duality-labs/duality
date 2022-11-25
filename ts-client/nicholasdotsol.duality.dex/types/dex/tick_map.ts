/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { LimitOrderTrancheTrancheIndexes } from "./limit_order_pool_tranche_indexes";
import { TickDataType } from "./tick_data_type";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TickMap {
  pairId: string;
  tickIndex: number;
  tickData: TickDataType | undefined;
  LimitOrderTranche0to1: LimitOrderTrancheTrancheIndexes | undefined;
  LimitOrderTranche1to0: LimitOrderTrancheTrancheIndexes | undefined;
}

function createBaseTickMap(): TickMap {
  return {
    pairId: "",
    tickIndex: 0,
    tickData: undefined,
    LimitOrderTranche0to1: undefined,
    LimitOrderTranche1to0: undefined,
  };
}

export const TickMap = {
  encode(message: TickMap, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(16).int64(message.tickIndex);
    }
    if (message.tickData !== undefined) {
      TickDataType.encode(message.tickData, writer.uint32(26).fork()).ldelim();
    }
    if (message.LimitOrderTranche0to1 !== undefined) {
      LimitOrderTrancheTrancheIndexes.encode(message.LimitOrderTranche0to1, writer.uint32(34).fork()).ldelim();
    }
    if (message.LimitOrderTranche1to0 !== undefined) {
      LimitOrderTrancheTrancheIndexes.encode(message.LimitOrderTranche1to0, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TickMap {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTickMap();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairId = reader.string();
          break;
        case 2:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.tickData = TickDataType.decode(reader, reader.uint32());
          break;
        case 4:
          message.LimitOrderTranche0to1 = LimitOrderTrancheTrancheIndexes.decode(reader, reader.uint32());
          break;
        case 5:
          message.LimitOrderTranche1to0 = LimitOrderTrancheTrancheIndexes.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TickMap {
    return {
      pairId: isSet(object.pairId) ? String(object.pairId) : "",
      tickIndex: isSet(object.tickIndex) ? Number(object.tickIndex) : 0,
      tickData: isSet(object.tickData) ? TickDataType.fromJSON(object.tickData) : undefined,
      LimitOrderTranche0to1: isSet(object.LimitOrderTranche0to1)
        ? LimitOrderTrancheTrancheIndexes.fromJSON(object.LimitOrderTranche0to1)
        : undefined,
      LimitOrderTranche1to0: isSet(object.LimitOrderTranche1to0)
        ? LimitOrderTrancheTrancheIndexes.fromJSON(object.LimitOrderTranche1to0)
        : undefined,
    };
  },

  toJSON(message: TickMap): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = Math.round(message.tickIndex));
    message.tickData !== undefined
      && (obj.tickData = message.tickData ? TickDataType.toJSON(message.tickData) : undefined);
    message.LimitOrderTranche0to1 !== undefined && (obj.LimitOrderTranche0to1 = message.LimitOrderTranche0to1
      ? LimitOrderTrancheTrancheIndexes.toJSON(message.LimitOrderTranche0to1)
      : undefined);
    message.LimitOrderTranche1to0 !== undefined && (obj.LimitOrderTranche1to0 = message.LimitOrderTranche1to0
      ? LimitOrderTrancheTrancheIndexes.toJSON(message.LimitOrderTranche1to0)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TickMap>, I>>(object: I): TickMap {
    const message = createBaseTickMap();
    message.pairId = object.pairId ?? "";
    message.tickIndex = object.tickIndex ?? 0;
    message.tickData = (object.tickData !== undefined && object.tickData !== null)
      ? TickDataType.fromPartial(object.tickData)
      : undefined;
    message.LimitOrderTranche0to1 =
      (object.LimitOrderTranche0to1 !== undefined && object.LimitOrderTranche0to1 !== null)
        ? LimitOrderTrancheTrancheIndexes.fromPartial(object.LimitOrderTranche0to1)
        : undefined;
    message.LimitOrderTranche1to0 =
      (object.LimitOrderTranche1to0 !== undefined && object.LimitOrderTranche1to0 !== null)
        ? LimitOrderTrancheTrancheIndexes.fromPartial(object.LimitOrderTranche1to0)
        : undefined;
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
