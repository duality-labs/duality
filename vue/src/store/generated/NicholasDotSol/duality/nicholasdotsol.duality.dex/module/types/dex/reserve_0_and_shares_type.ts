/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Reserve0AndSharesType {
  reserve0: string;
  totalShares: string;
}

const baseReserve0AndSharesType: object = { reserve0: "", totalShares: "" };

export const Reserve0AndSharesType = {
  encode(
    message: Reserve0AndSharesType,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.reserve0 !== "") {
      writer.uint32(10).string(message.reserve0);
    }
    if (message.totalShares !== "") {
      writer.uint32(18).string(message.totalShares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Reserve0AndSharesType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseReserve0AndSharesType } as Reserve0AndSharesType;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.reserve0 = reader.string();
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

  fromJSON(object: any): Reserve0AndSharesType {
    const message = { ...baseReserve0AndSharesType } as Reserve0AndSharesType;
    if (object.reserve0 !== undefined && object.reserve0 !== null) {
      message.reserve0 = String(object.reserve0);
    } else {
      message.reserve0 = "";
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = String(object.totalShares);
    } else {
      message.totalShares = "";
    }
    return message;
  },

  toJSON(message: Reserve0AndSharesType): unknown {
    const obj: any = {};
    message.reserve0 !== undefined && (obj.reserve0 = message.reserve0);
    message.totalShares !== undefined &&
      (obj.totalShares = message.totalShares);
    return obj;
  },

  fromPartial(
    object: DeepPartial<Reserve0AndSharesType>
  ): Reserve0AndSharesType {
    const message = { ...baseReserve0AndSharesType } as Reserve0AndSharesType;
    if (object.reserve0 !== undefined && object.reserve0 !== null) {
      message.reserve0 = object.reserve0;
    } else {
      message.reserve0 = "";
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
