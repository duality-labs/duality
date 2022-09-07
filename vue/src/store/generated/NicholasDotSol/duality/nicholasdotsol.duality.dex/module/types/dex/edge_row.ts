/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface EdgeRow {
  edge: string;
}

const baseEdgeRow: object = { edge: "" };

export const EdgeRow = {
  encode(message: EdgeRow, writer: Writer = Writer.create()): Writer {
    if (message.edge !== "") {
      writer.uint32(10).string(message.edge);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EdgeRow {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEdgeRow } as EdgeRow;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.edge = reader.string();
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
    if (object.edge !== undefined && object.edge !== null) {
      message.edge = String(object.edge);
    } else {
      message.edge = "";
    }
    return message;
  },

  toJSON(message: EdgeRow): unknown {
    const obj: any = {};
    message.edge !== undefined && (obj.edge = message.edge);
    return obj;
  },

  fromPartial(object: DeepPartial<EdgeRow>): EdgeRow {
    const message = { ...baseEdgeRow } as EdgeRow;
    if (object.edge !== undefined && object.edge !== null) {
      message.edge = object.edge;
    } else {
      message.edge = "";
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
