/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Pool {
  reserveA: string;
  reserveB: string;
  price: string;
  fee: string;
  totalShares: string;
  index: number;
}

const basePool: object = {
  reserveA: "",
  reserveB: "",
  price: "",
  fee: "",
  totalShares: "",
  index: 0,
};

export const Pool = {
  encode(message: Pool, writer: Writer = Writer.create()): Writer {
    if (message.reserveA !== "") {
      writer.uint32(10).string(message.reserveA);
    }
    if (message.reserveB !== "") {
      writer.uint32(18).string(message.reserveB);
    }
    if (message.price !== "") {
      writer.uint32(26).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(34).string(message.fee);
    }
    if (message.totalShares !== "") {
      writer.uint32(42).string(message.totalShares);
    }
    if (message.index !== 0) {
      writer.uint32(48).int32(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePool } as Pool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.reserveA = reader.string();
          break;
        case 2:
          message.reserveB = reader.string();
          break;
        case 3:
          message.price = reader.string();
          break;
        case 4:
          message.fee = reader.string();
          break;
        case 5:
          message.totalShares = reader.string();
          break;
        case 6:
          message.index = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pool {
    const message = { ...basePool } as Pool;
    if (object.reserveA !== undefined && object.reserveA !== null) {
      message.reserveA = String(object.reserveA);
    } else {
      message.reserveA = "";
    }
    if (object.reserveB !== undefined && object.reserveB !== null) {
      message.reserveB = String(object.reserveB);
    } else {
      message.reserveB = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = String(object.totalShares);
    } else {
      message.totalShares = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = Number(object.index);
    } else {
      message.index = 0;
    }
    return message;
  },

  toJSON(message: Pool): unknown {
    const obj: any = {};
    message.reserveA !== undefined && (obj.reserveA = message.reserveA);
    message.reserveB !== undefined && (obj.reserveB = message.reserveB);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.totalShares !== undefined &&
      (obj.totalShares = message.totalShares);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<Pool>): Pool {
    const message = { ...basePool } as Pool;
    if (object.reserveA !== undefined && object.reserveA !== null) {
      message.reserveA = object.reserveA;
    } else {
      message.reserveA = "";
    }
    if (object.reserveB !== undefined && object.reserveB !== null) {
      message.reserveB = object.reserveB;
    } else {
      message.reserveB = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = object.totalShares;
    } else {
      message.totalShares = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = 0;
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
