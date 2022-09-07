/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface AdjacenyMatrix {
  edgeMatrix: boolean[];
}

const baseAdjacenyMatrix: object = { edgeMatrix: false };

export const AdjacenyMatrix = {
  encode(message: AdjacenyMatrix, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).fork();
    for (const v of message.edgeMatrix) {
      writer.bool(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AdjacenyMatrix {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAdjacenyMatrix } as AdjacenyMatrix;
    message.edgeMatrix = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.edgeMatrix.push(reader.bool());
            }
          } else {
            message.edgeMatrix.push(reader.bool());
          }
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
    message.edgeMatrix = [];
    if (object.edgeMatrix !== undefined && object.edgeMatrix !== null) {
      for (const e of object.edgeMatrix) {
        message.edgeMatrix.push(Boolean(e));
      }
    }
    return message;
  },

  toJSON(message: AdjacenyMatrix): unknown {
    const obj: any = {};
    if (message.edgeMatrix) {
      obj.edgeMatrix = message.edgeMatrix.map((e) => e);
    } else {
      obj.edgeMatrix = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AdjacenyMatrix>): AdjacenyMatrix {
    const message = { ...baseAdjacenyMatrix } as AdjacenyMatrix;
    message.edgeMatrix = [];
    if (object.edgeMatrix !== undefined && object.edgeMatrix !== null) {
      for (const e of object.edgeMatrix) {
        message.edgeMatrix.push(e);
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
