/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Share {
  owner: string;
  token0: string;
  token1: string;
  price: string;
  fee: string;
  shareAmount: string;
}

const baseShare: object = {
  owner: "",
  token0: "",
  token1: "",
  price: "",
  fee: "",
  shareAmount: "",
};

export const Share = {
  encode(message: Share, writer: Writer = Writer.create()): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.token0 !== "") {
      writer.uint32(18).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(26).string(message.token1);
    }
    if (message.price !== "") {
      writer.uint32(34).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(42).string(message.fee);
    }
    if (message.shareAmount !== "") {
      writer.uint32(50).string(message.shareAmount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Share {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseShare } as Share;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.token0 = reader.string();
          break;
        case 3:
          message.token1 = reader.string();
          break;
        case 4:
          message.price = reader.string();
          break;
        case 5:
          message.fee = reader.string();
          break;
        case 6:
          message.shareAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Share {
    const message = { ...baseShare } as Share;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = String(object.token0);
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = String(object.token1);
    } else {
      message.token1 = "";
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
    if (object.shareAmount !== undefined && object.shareAmount !== null) {
      message.shareAmount = String(object.shareAmount);
    } else {
      message.shareAmount = "";
    }
    return message;
  },

  toJSON(message: Share): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.shareAmount !== undefined &&
      (obj.shareAmount = message.shareAmount);
    return obj;
  },

  fromPartial(object: DeepPartial<Share>): Share {
    const message = { ...baseShare } as Share;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = object.token0;
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = object.token1;
    } else {
      message.token1 = "";
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
    if (object.shareAmount !== undefined && object.shareAmount !== null) {
      message.shareAmount = object.shareAmount;
    } else {
      message.shareAmount = "";
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
