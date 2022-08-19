/* eslint-disable */
import { OrderParams } from "../dex/order_params";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface IndexQueueType {
  price: string;
  fee: string;
  orderparams: OrderParams | undefined;
}

const baseIndexQueueType: object = { price: "", fee: "" };

export const IndexQueueType = {
  encode(message: IndexQueueType, writer: Writer = Writer.create()): Writer {
    if (message.price !== "") {
      writer.uint32(10).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(18).string(message.fee);
    }
    if (message.orderparams !== undefined) {
      OrderParams.encode(
        message.orderparams,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): IndexQueueType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseIndexQueueType } as IndexQueueType;
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
          message.orderparams = OrderParams.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IndexQueueType {
    const message = { ...baseIndexQueueType } as IndexQueueType;
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
      message.orderparams = OrderParams.fromJSON(object.orderparams);
    } else {
      message.orderparams = undefined;
    }
    return message;
  },

  toJSON(message: IndexQueueType): unknown {
    const obj: any = {};
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.orderparams !== undefined &&
      (obj.orderparams = message.orderparams
        ? OrderParams.toJSON(message.orderparams)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<IndexQueueType>): IndexQueueType {
    const message = { ...baseIndexQueueType } as IndexQueueType;
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
      message.orderparams = OrderParams.fromPartial(object.orderparams);
    } else {
      message.orderparams = undefined;
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
