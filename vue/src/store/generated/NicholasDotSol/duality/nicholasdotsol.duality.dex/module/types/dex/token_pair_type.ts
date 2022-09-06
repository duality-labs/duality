/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TokenPairType {
  currentTick0To1: string;
  currentTick1To0: string;
}

const baseTokenPairType: object = { currentTick0To1: "", currentTick1To0: "" };

export const TokenPairType = {
  encode(message: TokenPairType, writer: Writer = Writer.create()): Writer {
    if (message.currentTick0To1 !== "") {
      writer.uint32(10).string(message.currentTick0To1);
    }
    if (message.currentTick1To0 !== "") {
      writer.uint32(18).string(message.currentTick1To0);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TokenPairType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTokenPairType } as TokenPairType;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.currentTick0To1 = reader.string();
          break;
        case 2:
          message.currentTick1To0 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenPairType {
    const message = { ...baseTokenPairType } as TokenPairType;
    if (
      object.currentTick0To1 !== undefined &&
      object.currentTick0To1 !== null
    ) {
      message.currentTick0To1 = String(object.currentTick0To1);
    } else {
      message.currentTick0To1 = "";
    }
    if (
      object.currentTick1To0 !== undefined &&
      object.currentTick1To0 !== null
    ) {
      message.currentTick1To0 = String(object.currentTick1To0);
    } else {
      message.currentTick1To0 = "";
    }
    return message;
  },

  toJSON(message: TokenPairType): unknown {
    const obj: any = {};
    message.currentTick0To1 !== undefined &&
      (obj.currentTick0To1 = message.currentTick0To1);
    message.currentTick1To0 !== undefined &&
      (obj.currentTick1To0 = message.currentTick1To0);
    return obj;
  },

  fromPartial(object: DeepPartial<TokenPairType>): TokenPairType {
    const message = { ...baseTokenPairType } as TokenPairType;
    if (
      object.currentTick0To1 !== undefined &&
      object.currentTick0To1 !== null
    ) {
      message.currentTick0To1 = object.currentTick0To1;
    } else {
      message.currentTick0To1 = "";
    }
    if (
      object.currentTick1To0 !== undefined &&
      object.currentTick1To0 !== null
    ) {
      message.currentTick1To0 = object.currentTick1To0;
    } else {
      message.currentTick1To0 = "";
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
