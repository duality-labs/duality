/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface VirtualPriceQueue {
  vPrice: string;
  direction: string;
  orderType: string;
  price: string;
  fee: string;
  orderparams: string;
}

const baseVirtualPriceQueue: object = {
  vPrice: "",
  direction: "",
  orderType: "",
  price: "",
  fee: "",
  orderparams: "",
};

export const VirtualPriceQueue = {
  encode(message: VirtualPriceQueue, writer: Writer = Writer.create()): Writer {
    if (message.vPrice !== "") {
      writer.uint32(10).string(message.vPrice);
    }
    if (message.direction !== "") {
      writer.uint32(18).string(message.direction);
    }
    if (message.orderType !== "") {
      writer.uint32(26).string(message.orderType);
    }
    if (message.price !== "") {
      writer.uint32(34).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(42).string(message.fee);
    }
    if (message.orderparams !== "") {
      writer.uint32(50).string(message.orderparams);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VirtualPriceQueue {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVirtualPriceQueue } as VirtualPriceQueue;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vPrice = reader.string();
          break;
        case 2:
          message.direction = reader.string();
          break;
        case 3:
          message.orderType = reader.string();
          break;
        case 4:
          message.price = reader.string();
          break;
        case 5:
          message.fee = reader.string();
          break;
        case 6:
          message.orderparams = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VirtualPriceQueue {
    const message = { ...baseVirtualPriceQueue } as VirtualPriceQueue;
    if (object.vPrice !== undefined && object.vPrice !== null) {
      message.vPrice = String(object.vPrice);
    } else {
      message.vPrice = "";
    }
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = String(object.direction);
    } else {
      message.direction = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
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
    if (object.orderparams !== undefined && object.orderparams !== null) {
      message.orderparams = String(object.orderparams);
    } else {
      message.orderparams = "";
    }
    return message;
  },

  toJSON(message: VirtualPriceQueue): unknown {
    const obj: any = {};
    message.vPrice !== undefined && (obj.vPrice = message.vPrice);
    message.direction !== undefined && (obj.direction = message.direction);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.orderparams !== undefined &&
      (obj.orderparams = message.orderparams);
    return obj;
  },

  fromPartial(object: DeepPartial<VirtualPriceQueue>): VirtualPriceQueue {
    const message = { ...baseVirtualPriceQueue } as VirtualPriceQueue;
    if (object.vPrice !== undefined && object.vPrice !== null) {
      message.vPrice = object.vPrice;
    } else {
      message.vPrice = "";
    }
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = object.direction;
    } else {
      message.direction = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
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
    if (object.orderparams !== undefined && object.orderparams !== null) {
      message.orderparams = object.orderparams;
    } else {
      message.orderparams = "";
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
