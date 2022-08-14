/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Node {
  token: string;
  outgoingEdges: string[];
}

const baseNode: object = { token: "", outgoingEdges: "" };

export const Node = {
  encode(message: Node, writer: Writer = Writer.create()): Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    for (const v of message.outgoingEdges) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Node {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNode } as Node;
    message.outgoingEdges = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
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

  fromJSON(object: any): Node {
    const message = { ...baseNode } as Node;
    message.outgoingEdges = [];
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.outgoingEdges !== undefined && object.outgoingEdges !== null) {
      for (const e of object.outgoingEdges) {
        message.outgoingEdges.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Node): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    if (message.outgoingEdges) {
      obj.outgoingEdges = message.outgoingEdges.map((e) => e);
    } else {
      obj.outgoingEdges = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Node>): Node {
    const message = { ...baseNode } as Node;
    message.outgoingEdges = [];
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
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
