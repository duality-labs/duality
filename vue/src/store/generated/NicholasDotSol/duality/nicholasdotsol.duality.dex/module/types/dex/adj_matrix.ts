/* eslint-disable */
import { EdgeRow } from "./edge_row";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface AdjacenyMatrix {
  EdgeMatrix: EdgeRow[];
}

const baseAdjacenyMatrix: object = {};

export const AdjacenyMatrix = {
  encode(message: AdjacenyMatrix, writer: Writer = Writer.create()): Writer {
    for (const v of message.EdgeMatrix) {
      EdgeRow.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AdjacenyMatrix {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAdjacenyMatrix } as AdjacenyMatrix;
    message.EdgeMatrix = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.EdgeMatrix.push(EdgeRow.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AdjacenyMatrix {
    const message = { ...baseAdjacenyMatrix } as AdjacenyMatrix;
    message.EdgeMatrix = [];
    if (object.EdgeMatrix !== undefined && object.EdgeMatrix !== null) {
      for (const e of object.EdgeMatrix) {
        message.EdgeMatrix.push(EdgeRow.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: AdjacenyMatrix): unknown {
    const obj: any = {};
    if (message.EdgeMatrix) {
      obj.EdgeMatrix = message.EdgeMatrix.map((e) =>
        e ? EdgeRow.toJSON(e) : undefined
      );
    } else {
      obj.EdgeMatrix = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AdjacenyMatrix>): AdjacenyMatrix {
    const message = { ...baseAdjacenyMatrix } as AdjacenyMatrix;
    message.EdgeMatrix = [];
    if (object.EdgeMatrix !== undefined && object.EdgeMatrix !== null) {
      for (const e of object.EdgeMatrix) {
        message.EdgeMatrix.push(EdgeRow.fromPartial(e));
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
