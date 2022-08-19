/* eslint-disable */
import { VirtualPriceQueueType } from "../dex/virtual_price_queue_type";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface VirtualPriceQueue {
  vPrice: string;
  direction: string;
  orderType: string;
  queue: VirtualPriceQueueType[];
}

const baseVirtualPriceQueue: object = {
  vPrice: "",
  direction: "",
  orderType: "",
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
    for (const v of message.queue) {
      VirtualPriceQueueType.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VirtualPriceQueue {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVirtualPriceQueue } as VirtualPriceQueue;
    message.queue = [];
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
          message.queue.push(
            VirtualPriceQueueType.decode(reader, reader.uint32())
          );
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
    message.queue = [];
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
    if (object.queue !== undefined && object.queue !== null) {
      for (const e of object.queue) {
        message.queue.push(VirtualPriceQueueType.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: VirtualPriceQueue): unknown {
    const obj: any = {};
    message.vPrice !== undefined && (obj.vPrice = message.vPrice);
    message.direction !== undefined && (obj.direction = message.direction);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    if (message.queue) {
      obj.queue = message.queue.map((e) =>
        e ? VirtualPriceQueueType.toJSON(e) : undefined
      );
    } else {
      obj.queue = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<VirtualPriceQueue>): VirtualPriceQueue {
    const message = { ...baseVirtualPriceQueue } as VirtualPriceQueue;
    message.queue = [];
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
    if (object.queue !== undefined && object.queue !== null) {
      for (const e of object.queue) {
        message.queue.push(VirtualPriceQueueType.fromPartial(e));
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
