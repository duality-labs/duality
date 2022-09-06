/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TickDataType {
  reserve0AndShares: string;
  reserve1: string;
}

const baseTickDataType: object = { reserve0AndShares: "", reserve1: "" };

export const TickDataType = {
  encode(message: TickDataType, writer: Writer = Writer.create()): Writer {
    if (message.reserve0AndShares !== "") {
      writer.uint32(10).string(message.reserve0AndShares);
    }
    if (message.reserve1 !== "") {
      writer.uint32(18).string(message.reserve1);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TickDataType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTickDataType } as TickDataType;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.reserve0AndShares = reader.string();
          break;
        case 2:
          message.reserve1 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TickDataType {
    const message = { ...baseTickDataType } as TickDataType;
    if (
      object.reserve0AndShares !== undefined &&
      object.reserve0AndShares !== null
    ) {
      message.reserve0AndShares = String(object.reserve0AndShares);
    } else {
      message.reserve0AndShares = "";
    }
    if (object.reserve1 !== undefined && object.reserve1 !== null) {
      message.reserve1 = String(object.reserve1);
    } else {
      message.reserve1 = "";
    }
    return message;
  },

  toJSON(message: TickDataType): unknown {
    const obj: any = {};
    message.reserve0AndShares !== undefined &&
      (obj.reserve0AndShares = message.reserve0AndShares);
    message.reserve1 !== undefined && (obj.reserve1 = message.reserve1);
    return obj;
  },

  fromPartial(object: DeepPartial<TickDataType>): TickDataType {
    const message = { ...baseTickDataType } as TickDataType;
    if (
      object.reserve0AndShares !== undefined &&
      object.reserve0AndShares !== null
    ) {
      message.reserve0AndShares = object.reserve0AndShares;
    } else {
      message.reserve0AndShares = "";
    }
    if (object.reserve1 !== undefined && object.reserve1 !== null) {
      message.reserve1 = object.reserve1;
    } else {
      message.reserve1 = "";
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
