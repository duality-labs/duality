/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface OrderParams {
  orderRule: string;
  orderType: string;
  orderShares: string;
}

const baseOrderParams: object = {
  orderRule: "",
  orderType: "",
  orderShares: "",
};

export const OrderParams = {
  encode(message: OrderParams, writer: Writer = Writer.create()): Writer {
    if (message.orderRule !== "") {
      writer.uint32(10).string(message.orderRule);
    }
    if (message.orderType !== "") {
      writer.uint32(18).string(message.orderType);
    }
    if (message.orderShares !== "") {
      writer.uint32(26).string(message.orderShares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OrderParams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrderParams } as OrderParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orderRule = reader.string();
          break;
        case 2:
          message.orderType = reader.string();
          break;
        case 3:
          message.orderShares = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OrderParams {
    const message = { ...baseOrderParams } as OrderParams;
    if (object.orderRule !== undefined && object.orderRule !== null) {
      message.orderRule = String(object.orderRule);
    } else {
      message.orderRule = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    if (object.orderShares !== undefined && object.orderShares !== null) {
      message.orderShares = String(object.orderShares);
    } else {
      message.orderShares = "";
    }
    return message;
  },

  toJSON(message: OrderParams): unknown {
    const obj: any = {};
    message.orderRule !== undefined && (obj.orderRule = message.orderRule);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.orderShares !== undefined &&
      (obj.orderShares = message.orderShares);
    return obj;
  },

  fromPartial(object: DeepPartial<OrderParams>): OrderParams {
    const message = { ...baseOrderParams } as OrderParams;
    if (object.orderRule !== undefined && object.orderRule !== null) {
      message.orderRule = object.orderRule;
    } else {
      message.orderRule = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    if (object.orderShares !== undefined && object.orderShares !== null) {
      message.orderShares = object.orderShares;
    } else {
      message.orderShares = "";
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
