/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TickMap {
  tickIndex: string;
  tickData: string;
}

const baseTickMap: object = { tickIndex: "", tickData: "" };

export const TickMap = {
  encode(message: TickMap, writer: Writer = Writer.create()): Writer {
    if (message.tickIndex !== "") {
      writer.uint32(10).string(message.tickIndex);
    }
    if (message.tickData !== "") {
      writer.uint32(18).string(message.tickData);
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
          message.tickIndex = reader.string();
          break;
        case 2:
          message.tickData = reader.string();
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
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = String(object.tickIndex);
    } else {
      message.tickIndex = "";
    }
    if (object.tickData !== undefined && object.tickData !== null) {
      message.tickData = String(object.tickData);
    } else {
      message.tickData = "";
    }
    return message;
  },

  toJSON(message: TickMap): unknown {
    const obj: any = {};
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.tickData !== undefined && (obj.tickData = message.tickData);
    return obj;
  },

  fromPartial(object: DeepPartial<TickMap>): TickMap {
    const message = { ...baseTickMap } as TickMap;
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = "";
    }
    if (object.tickData !== undefined && object.tickData !== null) {
      message.tickData = object.tickData;
    } else {
      message.tickData = "";
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
