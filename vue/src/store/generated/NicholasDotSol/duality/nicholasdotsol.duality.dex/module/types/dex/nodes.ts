/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Nodes {
  node: string;
  outgoingEdges: string[];
}

const baseNodes: object = { node: "", outgoingEdges: "" };

export const Nodes = {
  encode(message: Nodes, writer: Writer = Writer.create()): Writer {
    if (message.node !== "") {
      writer.uint32(10).string(message.node);
    }
    for (const v of message.outgoingEdges) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Nodes {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNodes } as Nodes;
    message.outgoingEdges = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.node = reader.string();
          break;
        case 2:
          message.outgoingEdges.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Nodes {
    const message = { ...baseNodes } as Nodes;
    message.outgoingEdges = [];
    if (object.node !== undefined && object.node !== null) {
      message.node = String(object.node);
    } else {
      message.node = "";
    }
    if (object.outgoingEdges !== undefined && object.outgoingEdges !== null) {
      for (const e of object.outgoingEdges) {
        message.outgoingEdges.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Nodes): unknown {
    const obj: any = {};
    message.node !== undefined && (obj.node = message.node);
    if (message.outgoingEdges) {
      obj.outgoingEdges = message.outgoingEdges.map((e) => e);
    } else {
      obj.outgoingEdges = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Nodes>): Nodes {
    const message = { ...baseNodes } as Nodes;
    message.outgoingEdges = [];
    if (object.node !== undefined && object.node !== null) {
      message.node = object.node;
    } else {
      message.node = "";
    }
    if (object.outgoingEdges !== undefined && object.outgoingEdges !== null) {
      for (const e of object.outgoingEdges) {
        message.outgoingEdges.push(e);
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
