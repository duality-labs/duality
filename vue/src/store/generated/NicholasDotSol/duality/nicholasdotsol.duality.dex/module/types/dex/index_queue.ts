/* eslint-disable */
import { IndexQueueType } from "../dex/index_queue_type";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface IndexQueue {
  index: number;
  queue: IndexQueueType[];
}

const baseIndexQueue: object = { index: 0 };

export const IndexQueue = {
  encode(message: IndexQueue, writer: Writer = Writer.create()): Writer {
    if (message.index !== 0) {
      writer.uint32(8).int32(message.index);
    }
    for (const v of message.queue) {
      IndexQueueType.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): IndexQueue {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseIndexQueue } as IndexQueue;
    message.queue = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.int32();
          break;
        case 3:
          message.queue.push(IndexQueueType.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IndexQueue {
    const message = { ...baseIndexQueue } as IndexQueue;
    message.queue = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = Number(object.index);
    } else {
      message.index = 0;
    }
    if (object.queue !== undefined && object.queue !== null) {
      for (const e of object.queue) {
        message.queue.push(IndexQueueType.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: IndexQueue): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    if (message.queue) {
      obj.queue = message.queue.map((e) =>
        e ? IndexQueueType.toJSON(e) : undefined
      );
    } else {
      obj.queue = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<IndexQueue>): IndexQueue {
    const message = { ...baseIndexQueue } as IndexQueue;
    message.queue = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = 0;
    }
    if (object.queue !== undefined && object.queue !== null) {
      for (const e of object.queue) {
        message.queue.push(IndexQueueType.fromPartial(e));
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
