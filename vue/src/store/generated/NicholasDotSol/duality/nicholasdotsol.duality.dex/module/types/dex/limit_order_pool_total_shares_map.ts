/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderPoolTotalSharesMap {
  count: string;
  totalShares: string;
}

const baseLimitOrderPoolTotalSharesMap: object = { count: "", totalShares: "" };

export const LimitOrderPoolTotalSharesMap = {
  encode(
    message: LimitOrderPoolTotalSharesMap,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== "") {
      writer.uint32(10).string(message.count);
    }
    if (message.totalShares !== "") {
      writer.uint32(18).string(message.totalShares);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): LimitOrderPoolTotalSharesMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseLimitOrderPoolTotalSharesMap,
    } as LimitOrderPoolTotalSharesMap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = reader.string();
          break;
        case 2:
          message.totalShares = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderPoolTotalSharesMap {
    const message = {
      ...baseLimitOrderPoolTotalSharesMap,
    } as LimitOrderPoolTotalSharesMap;
    if (object.count !== undefined && object.count !== null) {
      message.count = String(object.count);
    } else {
      message.count = "";
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = String(object.totalShares);
    } else {
      message.totalShares = "";
    }
    return message;
  },

  toJSON(message: LimitOrderPoolTotalSharesMap): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.totalShares !== undefined &&
      (obj.totalShares = message.totalShares);
    return obj;
  },

  fromPartial(
    object: DeepPartial<LimitOrderPoolTotalSharesMap>
  ): LimitOrderPoolTotalSharesMap {
    const message = {
      ...baseLimitOrderPoolTotalSharesMap,
    } as LimitOrderPoolTotalSharesMap;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = "";
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = object.totalShares;
    } else {
      message.totalShares = "";
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
