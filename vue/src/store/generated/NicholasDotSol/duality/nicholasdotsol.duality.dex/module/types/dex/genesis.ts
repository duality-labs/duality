/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { Nodes } from "../dex/nodes";
import { VirtualPriceTickQueue } from "../dex/virtual_price_tick_queue";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** GenesisState defines the dex module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  nodesList: Nodes[];
  nodesCount: number;
  virtualPriceTickQueueList: VirtualPriceTickQueue[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  virtualPriceTickQueueCount: number;
}

const baseGenesisState: object = {
  nodesCount: 0,
  virtualPriceTickQueueCount: 0,
};

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
    for (const v of message.virtualPriceTickQueueList) {
      VirtualPriceTickQueue.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.virtualPriceTickQueueCount !== 0) {
      writer.uint32(40).uint64(message.virtualPriceTickQueueCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.nodesList = [];
    message.virtualPriceTickQueueList = [];
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
          message.virtualPriceTickQueueList.push(
            VirtualPriceTickQueue.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.virtualPriceTickQueueCount = longToNumber(
            reader.uint64() as Long
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
    message.virtualPriceTickQueueList = [];
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
    if (
      object.virtualPriceTickQueueList !== undefined &&
      object.virtualPriceTickQueueList !== null
    ) {
      for (const e of object.virtualPriceTickQueueList) {
        message.virtualPriceTickQueueList.push(
          VirtualPriceTickQueue.fromJSON(e)
        );
      }
    }
    if (
      object.virtualPriceTickQueueCount !== undefined &&
      object.virtualPriceTickQueueCount !== null
    ) {
      message.virtualPriceTickQueueCount = Number(
        object.virtualPriceTickQueueCount
      );
    } else {
      message.virtualPriceTickQueueCount = 0;
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
    if (message.virtualPriceTickQueueList) {
      obj.virtualPriceTickQueueList = message.virtualPriceTickQueueList.map(
        (e) => (e ? VirtualPriceTickQueue.toJSON(e) : undefined)
      );
    } else {
      obj.virtualPriceTickQueueList = [];
    }
    message.virtualPriceTickQueueCount !== undefined &&
      (obj.virtualPriceTickQueueCount = message.virtualPriceTickQueueCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.nodesList = [];
    message.virtualPriceTickQueueList = [];
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
    if (
      object.virtualPriceTickQueueList !== undefined &&
      object.virtualPriceTickQueueList !== null
    ) {
      for (const e of object.virtualPriceTickQueueList) {
        message.virtualPriceTickQueueList.push(
          VirtualPriceTickQueue.fromPartial(e)
        );
      }
    }
    if (
      object.virtualPriceTickQueueCount !== undefined &&
      object.virtualPriceTickQueueCount !== null
    ) {
      message.virtualPriceTickQueueCount = object.virtualPriceTickQueueCount;
    } else {
      message.virtualPriceTickQueueCount = 0;
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
