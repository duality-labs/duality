/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { TickDataType } from "../dex/tick_data_type";
import { LimitOrderTrancheTrancheIndexes } from "../dex/limit_order_pool_tranche_indexes";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TickMap {
  pairId: string;
  tickIndex: number;
  tickData: TickDataType | undefined;
  LimitOrderTranche0to1: LimitOrderTrancheTrancheIndexes | undefined;
  LimitOrderTranche1to0: LimitOrderTrancheTrancheIndexes | undefined;
}

const baseTickMap: object = { pairId: "", tickIndex: 0 };

export const TickMap = {
  encode(message: TickMap, writer: Writer = Writer.create()): Writer {
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
      LimitOrderTrancheTrancheIndexes.encode(
        message.LimitOrderTranche0to1,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.LimitOrderTranche1to0 !== undefined) {
      LimitOrderTrancheTrancheIndexes.encode(
        message.LimitOrderTranche1to0,
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TickMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTickMap } as TickMap;
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
          message.LimitOrderTranche0to1 = LimitOrderTrancheTrancheIndexes.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.LimitOrderTranche1to0 = LimitOrderTrancheTrancheIndexes.decode(
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

  fromJSON(object: any): TickMap {
    const message = { ...baseTickMap } as TickMap;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.tickData !== undefined && object.tickData !== null) {
      message.tickData = TickDataType.fromJSON(object.tickData);
    } else {
      message.tickData = undefined;
    }
    if (
      object.LimitOrderTranche0to1 !== undefined &&
      object.LimitOrderTranche0to1 !== null
    ) {
      message.LimitOrderTranche0to1 = LimitOrderTrancheTrancheIndexes.fromJSON(
        object.LimitOrderTranche0to1
      );
    } else {
      message.LimitOrderTranche0to1 = undefined;
    }
    if (
      object.LimitOrderTranche1to0 !== undefined &&
      object.LimitOrderTranche1to0 !== null
    ) {
      message.LimitOrderTranche1to0 = LimitOrderTrancheTrancheIndexes.fromJSON(
        object.LimitOrderTranche1to0
      );
    } else {
      message.LimitOrderTranche1to0 = undefined;
    }
    return message;
  },

  toJSON(message: TickMap): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.tickData !== undefined &&
      (obj.tickData = message.tickData
        ? TickDataType.toJSON(message.tickData)
        : undefined);
    message.LimitOrderTranche0to1 !== undefined &&
      (obj.LimitOrderTranche0to1 = message.LimitOrderTranche0to1
        ? LimitOrderTrancheTrancheIndexes.toJSON(message.LimitOrderTranche0to1)
        : undefined);
    message.LimitOrderTranche1to0 !== undefined &&
      (obj.LimitOrderTranche1to0 = message.LimitOrderTranche1to0
        ? LimitOrderTrancheTrancheIndexes.toJSON(message.LimitOrderTranche1to0)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<TickMap>): TickMap {
    const message = { ...baseTickMap } as TickMap;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.tickData !== undefined && object.tickData !== null) {
      message.tickData = TickDataType.fromPartial(object.tickData);
    } else {
      message.tickData = undefined;
    }
    if (
      object.LimitOrderTranche0to1 !== undefined &&
      object.LimitOrderTranche0to1 !== null
    ) {
      message.LimitOrderTranche0to1 = LimitOrderTrancheTrancheIndexes.fromPartial(
        object.LimitOrderTranche0to1
      );
    } else {
      message.LimitOrderTranche0to1 = undefined;
    }
    if (
      object.LimitOrderTranche1to0 !== undefined &&
      object.LimitOrderTranche1to0 !== null
    ) {
      message.LimitOrderTranche1to0 = LimitOrderTrancheTrancheIndexes.fromPartial(
        object.LimitOrderTranche1to0
      );
    } else {
      message.LimitOrderTranche1to0 = undefined;
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
