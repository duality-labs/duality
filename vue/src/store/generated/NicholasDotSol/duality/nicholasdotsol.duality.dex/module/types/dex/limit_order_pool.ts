/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderPool {
  count: string;
  currentLimitOrderKey: string;
}

const baseLimitOrderPool: object = { count: "", currentLimitOrderKey: "" };

export const LimitOrderPool = {
  encode(message: LimitOrderPool, writer: Writer = Writer.create()): Writer {
    if (message.count !== "") {
      writer.uint32(10).string(message.count);
    }
    if (message.currentLimitOrderKey !== "") {
      writer.uint32(18).string(message.currentLimitOrderKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LimitOrderPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLimitOrderPool } as LimitOrderPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = reader.string();
          break;
        case 2:
          message.currentLimitOrderKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderPool {
    const message = { ...baseLimitOrderPool } as LimitOrderPool;
    if (object.count !== undefined && object.count !== null) {
      message.count = String(object.count);
    } else {
      message.count = "";
    }
    if (
      object.currentLimitOrderKey !== undefined &&
      object.currentLimitOrderKey !== null
    ) {
      message.currentLimitOrderKey = String(object.currentLimitOrderKey);
    } else {
      message.currentLimitOrderKey = "";
    }
    return message;
  },

  toJSON(message: LimitOrderPool): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.currentLimitOrderKey !== undefined &&
      (obj.currentLimitOrderKey = message.currentLimitOrderKey);
    return obj;
  },

  fromPartial(object: DeepPartial<LimitOrderPool>): LimitOrderPool {
    const message = { ...baseLimitOrderPool } as LimitOrderPool;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = "";
    }
    if (
      object.currentLimitOrderKey !== undefined &&
      object.currentLimitOrderKey !== null
    ) {
      message.currentLimitOrderKey = object.currentLimitOrderKey;
    } else {
      message.currentLimitOrderKey = "";
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
