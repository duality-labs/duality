/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Shares {
  address: string;
  pairId: string;
  priceIndex: string;
  fee: string;
  sharesOwned: string;
}

const baseShares: object = {
  address: "",
  pairId: "",
  priceIndex: "",
  fee: "",
  sharesOwned: "",
};

export const Shares = {
  encode(message: Shares, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.pairId !== "") {
      writer.uint32(18).string(message.pairId);
    }
    if (message.priceIndex !== "") {
      writer.uint32(26).string(message.priceIndex);
    }
    if (message.fee !== "") {
      writer.uint32(34).string(message.fee);
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
          message.pairId = reader.string();
          break;
        case 3:
          message.priceIndex = reader.string();
          break;
        case 4:
          message.fee = reader.string();
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
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = String(object.priceIndex);
    } else {
      message.priceIndex = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
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
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.priceIndex !== undefined && (obj.priceIndex = message.priceIndex);
    message.fee !== undefined && (obj.fee = message.fee);
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
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = object.priceIndex;
    } else {
      message.priceIndex = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
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
