/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface EdgeRow {
  edge: boolean[];
}

const baseEdgeRow: object = { edge: false };

export const EdgeRow = {
  encode(message: EdgeRow, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).fork();
    for (const v of message.edge) {
      writer.bool(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EdgeRow {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEdgeRow } as EdgeRow;
    message.edge = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.edge.push(reader.bool());
            }
          } else {
            message.edge.push(reader.bool());
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EdgeRow {
    const message = { ...baseEdgeRow } as EdgeRow;
    message.edge = [];
    if (object.edge !== undefined && object.edge !== null) {
      for (const e of object.edge) {
        message.edge.push(Boolean(e));
      }
    }
    return message;
  },

  toJSON(message: EdgeRow): unknown {
    const obj: any = {};
    if (message.edge) {
      obj.edge = message.edge.map((e) => e);
    } else {
      obj.edge = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<EdgeRow>): EdgeRow {
    const message = { ...baseEdgeRow } as EdgeRow;
    message.edge = [];
    if (object.edge !== undefined && object.edge !== null) {
      for (const e of object.edge) {
        message.edge.push(e);
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
