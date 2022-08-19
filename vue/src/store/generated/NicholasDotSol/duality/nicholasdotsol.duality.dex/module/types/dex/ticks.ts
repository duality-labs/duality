/* eslint-disable */
import { OrderParams } from "../dex/order_params";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Ticks {
  price: string;
  fee: string;
  orderType: string;
  reserve0: string;
  reserve1: string;
  pairPrice: string;
  pairFee: string;
  totalShares: string;
  orderparams: OrderParams[];
}

const baseTicks: object = {
  price: "",
  fee: "",
  orderType: "",
  reserve0: "",
  reserve1: "",
  pairPrice: "",
  pairFee: "",
  totalShares: "",
};

export const Ticks = {
  encode(message: Ticks, writer: Writer = Writer.create()): Writer {
    if (message.price !== "") {
      writer.uint32(10).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(18).string(message.fee);
    }
    if (message.orderType !== "") {
      writer.uint32(26).string(message.orderType);
    }
    if (message.reserve0 !== "") {
      writer.uint32(34).string(message.reserve0);
    }
    if (message.reserve1 !== "") {
      writer.uint32(42).string(message.reserve1);
    }
    if (message.pairPrice !== "") {
      writer.uint32(50).string(message.pairPrice);
    }
    if (message.pairFee !== "") {
      writer.uint32(58).string(message.pairFee);
    }
    if (message.totalShares !== "") {
      writer.uint32(66).string(message.totalShares);
    }
    for (const v of message.orderparams) {
      OrderParams.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Ticks {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTicks } as Ticks;
    message.orderparams = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.price = reader.string();
          break;
        case 2:
          message.fee = reader.string();
          break;
        case 3:
          message.orderType = reader.string();
          break;
        case 4:
          message.reserve0 = reader.string();
          break;
        case 5:
          message.reserve1 = reader.string();
          break;
        case 6:
          message.pairPrice = reader.string();
          break;
        case 7:
          message.pairFee = reader.string();
          break;
        case 8:
          message.totalShares = reader.string();
          break;
        case 9:
          message.orderparams.push(OrderParams.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Ticks {
    const message = { ...baseTicks } as Ticks;
    message.orderparams = [];
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
    if (object.reserve0 !== undefined && object.reserve0 !== null) {
      message.reserve0 = String(object.reserve0);
    } else {
      message.reserve0 = "";
    }
    if (object.reserve1 !== undefined && object.reserve1 !== null) {
      message.reserve1 = String(object.reserve1);
    } else {
      message.reserve1 = "";
    }
    if (object.pairPrice !== undefined && object.pairPrice !== null) {
      message.pairPrice = String(object.pairPrice);
    } else {
      message.pairPrice = "";
    }
    if (object.pairFee !== undefined && object.pairFee !== null) {
      message.pairFee = String(object.pairFee);
    } else {
      message.pairFee = "";
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = String(object.totalShares);
    } else {
      message.totalShares = "";
    }
    if (object.orderparams !== undefined && object.orderparams !== null) {
      for (const e of object.orderparams) {
        message.orderparams.push(OrderParams.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Ticks): unknown {
    const obj: any = {};
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.reserve0 !== undefined && (obj.reserve0 = message.reserve0);
    message.reserve1 !== undefined && (obj.reserve1 = message.reserve1);
    message.pairPrice !== undefined && (obj.pairPrice = message.pairPrice);
    message.pairFee !== undefined && (obj.pairFee = message.pairFee);
    message.totalShares !== undefined &&
      (obj.totalShares = message.totalShares);
    if (message.orderparams) {
      obj.orderparams = message.orderparams.map((e) =>
        e ? OrderParams.toJSON(e) : undefined
      );
    } else {
      obj.orderparams = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Ticks>): Ticks {
    const message = { ...baseTicks } as Ticks;
    message.orderparams = [];
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
    if (object.reserve0 !== undefined && object.reserve0 !== null) {
      message.reserve0 = object.reserve0;
    } else {
      message.reserve0 = "";
    }
    if (object.reserve1 !== undefined && object.reserve1 !== null) {
      message.reserve1 = object.reserve1;
    } else {
      message.reserve1 = "";
    }
    if (object.pairPrice !== undefined && object.pairPrice !== null) {
      message.pairPrice = object.pairPrice;
    } else {
      message.pairPrice = "";
    }
    if (object.pairFee !== undefined && object.pairFee !== null) {
      message.pairFee = object.pairFee;
    } else {
      message.pairFee = "";
    }
    if (object.totalShares !== undefined && object.totalShares !== null) {
      message.totalShares = object.totalShares;
    } else {
      message.totalShares = "";
    }
    if (object.orderparams !== undefined && object.orderparams !== null) {
      for (const e of object.orderparams) {
        message.orderparams.push(OrderParams.fromPartial(e));
      }
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
