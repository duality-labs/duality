/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderPoolReserveMap {
  count: string;
  reserves: string;
}

const baseLimitOrderPoolReserveMap: object = { count: "", reserves: "" };

export const LimitOrderPoolReserveMap = {
  encode(
    message: LimitOrderPoolReserveMap,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== "") {
      writer.uint32(10).string(message.count);
    }
    if (message.reserves !== "") {
      writer.uint32(18).string(message.reserves);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): LimitOrderPoolReserveMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseLimitOrderPoolReserveMap,
    } as LimitOrderPoolReserveMap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = reader.string();
          break;
        case 2:
          message.reserves = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderPoolReserveMap {
    const message = {
      ...baseLimitOrderPoolReserveMap,
    } as LimitOrderPoolReserveMap;
    if (object.count !== undefined && object.count !== null) {
      message.count = String(object.count);
    } else {
      message.count = "";
    }
    if (object.reserves !== undefined && object.reserves !== null) {
      message.reserves = String(object.reserves);
    } else {
      message.reserves = "";
    }
    return message;
  },

  toJSON(message: LimitOrderPoolReserveMap): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.reserves !== undefined && (obj.reserves = message.reserves);
    return obj;
  },

  fromPartial(
    object: DeepPartial<LimitOrderPoolReserveMap>
  ): LimitOrderPoolReserveMap {
    const message = {
      ...baseLimitOrderPoolReserveMap,
    } as LimitOrderPoolReserveMap;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = "";
    }
    if (object.reserves !== undefined && object.reserves !== null) {
      message.reserves = object.reserves;
    } else {
      message.reserves = "";
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
