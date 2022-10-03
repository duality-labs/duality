/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderPoolFillMap {
  count: string;
  fill: string;
}

const baseLimitOrderPoolFillMap: object = { count: "", fill: "" };

export const LimitOrderPoolFillMap = {
  encode(
    message: LimitOrderPoolFillMap,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== "") {
      writer.uint32(10).string(message.count);
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
          message.count = reader.string();
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
      message.count = String(object.count);
    } else {
      message.count = "";
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
      message.count = "";
    }
    if (object.fill !== undefined && object.fill !== null) {
      message.fill = object.fill;
    } else {
      message.fill = "";
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
