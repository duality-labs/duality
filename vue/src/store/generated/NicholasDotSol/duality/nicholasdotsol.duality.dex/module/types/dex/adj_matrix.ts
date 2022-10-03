/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { EdgeRow } from "../dex/edge_row";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface AdjMatrix {
  id: number;
  edgeRow: EdgeRow | undefined;
}

const baseAdjMatrix: object = { id: 0 };

export const AdjMatrix = {
  encode(message: AdjMatrix, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.edgeRow !== undefined) {
      EdgeRow.encode(message.edgeRow, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AdjMatrix {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAdjMatrix } as AdjMatrix;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.edgeRow = EdgeRow.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AdjMatrix {
    const message = { ...baseAdjMatrix } as AdjMatrix;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.edgeRow !== undefined && object.edgeRow !== null) {
      message.edgeRow = EdgeRow.fromJSON(object.edgeRow);
    } else {
      message.edgeRow = undefined;
    }
    return message;
  },

  toJSON(message: AdjMatrix): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.edgeRow !== undefined &&
      (obj.edgeRow = message.edgeRow
        ? EdgeRow.toJSON(message.edgeRow)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<AdjMatrix>): AdjMatrix {
    const message = { ...baseAdjMatrix } as AdjMatrix;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.edgeRow !== undefined && object.edgeRow !== null) {
      message.edgeRow = EdgeRow.fromPartial(object.edgeRow);
    } else {
      message.edgeRow = undefined;
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
