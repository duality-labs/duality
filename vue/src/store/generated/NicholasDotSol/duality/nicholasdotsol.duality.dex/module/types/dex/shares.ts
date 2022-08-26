/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Shares {
  address: string;
  price: string;
  fee: string;
  orderType: string;
  sharesOwned: string;
}

const baseShares: object = {
  address: "",
  price: "",
  fee: "",
  orderType: "",
  sharesOwned: "",
};

export const Shares = {
  encode(message: Shares, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.price !== "") {
      writer.uint32(18).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    if (message.orderType !== "") {
      writer.uint32(34).string(message.orderType);
    }
    if (message.sharesOwned !== "") {
      writer.uint32(42).string(message.sharesOwned);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Shares {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseShares } as Shares;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.price = reader.string();
          break;
        case 3:
          message.fee = reader.string();
          break;
        case 4:
          message.orderType = reader.string();
          break;
        case 5:
          message.sharesOwned = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Shares {
    const message = { ...baseShares } as Shares;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    if (object.sharesOwned !== undefined && object.sharesOwned !== null) {
      message.sharesOwned = String(object.sharesOwned);
    } else {
      message.sharesOwned = "";
    }
    return message;
  },

  toJSON(message: Shares): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.sharesOwned !== undefined &&
      (obj.sharesOwned = message.sharesOwned);
    return obj;
  },

  fromPartial(object: DeepPartial<Shares>): Shares {
    const message = { ...baseShares } as Shares;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    if (object.sharesOwned !== undefined && object.sharesOwned !== null) {
      message.sharesOwned = object.sharesOwned;
    } else {
      message.sharesOwned = "";
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
