/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { Nodes } from "../dex/nodes";
import { Ticks } from "../dex/ticks";
import { Pairs } from "../dex/pairs";
import { IndexQueue } from "../dex/index_queue";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** GenesisState defines the dex module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  nodesList: Nodes[];
  nodesCount: number;
  ticksList: Ticks[];
  pairsList: Pairs[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  indexQueueList: IndexQueue[];
}

const baseGenesisState: object = { nodesCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.nodesList) {
      Nodes.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.nodesCount !== 0) {
      writer.uint32(24).uint64(message.nodesCount);
    }
    for (const v of message.ticksList) {
      Ticks.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.pairsList) {
      Pairs.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.indexQueueList) {
      IndexQueue.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.nodesList = [];
    message.ticksList = [];
    message.pairsList = [];
    message.indexQueueList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.nodesList.push(Nodes.decode(reader, reader.uint32()));
          break;
        case 3:
          message.nodesCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.ticksList.push(Ticks.decode(reader, reader.uint32()));
          break;
        case 5:
          message.pairsList.push(Pairs.decode(reader, reader.uint32()));
          break;
        case 6:
          message.indexQueueList.push(
            IndexQueue.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.nodesList = [];
    message.ticksList = [];
    message.pairsList = [];
    message.indexQueueList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nodesList !== undefined && object.nodesList !== null) {
      for (const e of object.nodesList) {
        message.nodesList.push(Nodes.fromJSON(e));
      }
    }
    if (object.nodesCount !== undefined && object.nodesCount !== null) {
      message.nodesCount = Number(object.nodesCount);
    } else {
      message.nodesCount = 0;
    }
    if (object.ticksList !== undefined && object.ticksList !== null) {
      for (const e of object.ticksList) {
        message.ticksList.push(Ticks.fromJSON(e));
      }
    }
    if (object.pairsList !== undefined && object.pairsList !== null) {
      for (const e of object.pairsList) {
        message.pairsList.push(Pairs.fromJSON(e));
      }
    }
    if (object.indexQueueList !== undefined && object.indexQueueList !== null) {
      for (const e of object.indexQueueList) {
        message.indexQueueList.push(IndexQueue.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.nodesList) {
      obj.nodesList = message.nodesList.map((e) =>
        e ? Nodes.toJSON(e) : undefined
      );
    } else {
      obj.nodesList = [];
    }
    message.nodesCount !== undefined && (obj.nodesCount = message.nodesCount);
    if (message.ticksList) {
      obj.ticksList = message.ticksList.map((e) =>
        e ? Ticks.toJSON(e) : undefined
      );
    } else {
      obj.ticksList = [];
    }
    if (message.pairsList) {
      obj.pairsList = message.pairsList.map((e) =>
        e ? Pairs.toJSON(e) : undefined
      );
    } else {
      obj.pairsList = [];
    }
    if (message.indexQueueList) {
      obj.indexQueueList = message.indexQueueList.map((e) =>
        e ? IndexQueue.toJSON(e) : undefined
      );
    } else {
      obj.indexQueueList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.nodesList = [];
    message.ticksList = [];
    message.pairsList = [];
    message.indexQueueList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nodesList !== undefined && object.nodesList !== null) {
      for (const e of object.nodesList) {
        message.nodesList.push(Nodes.fromPartial(e));
      }
    }
    if (object.nodesCount !== undefined && object.nodesCount !== null) {
      message.nodesCount = object.nodesCount;
    } else {
      message.nodesCount = 0;
    }
    if (object.ticksList !== undefined && object.ticksList !== null) {
      for (const e of object.ticksList) {
        message.ticksList.push(Ticks.fromPartial(e));
      }
    }
    if (object.pairsList !== undefined && object.pairsList !== null) {
      for (const e of object.pairsList) {
        message.pairsList.push(Pairs.fromPartial(e));
      }
    }
    if (object.indexQueueList !== undefined && object.indexQueueList !== null) {
      for (const e of object.indexQueueList) {
        message.indexQueueList.push(IndexQueue.fromPartial(e));
      }
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
