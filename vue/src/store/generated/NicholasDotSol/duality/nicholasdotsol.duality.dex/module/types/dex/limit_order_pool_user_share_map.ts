/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface LimitOrderPoolUserShareMap {
  count: string;
  address: string;
  sharesOwned: string;
}

const baseLimitOrderPoolUserShareMap: object = {
  count: "",
  address: "",
  sharesOwned: "",
};

export const LimitOrderPoolUserShareMap = {
  encode(
    message: LimitOrderPoolUserShareMap,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== "") {
      writer.uint32(10).string(message.count);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.sharesOwned !== "") {
      writer.uint32(26).string(message.sharesOwned);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): LimitOrderPoolUserShareMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseLimitOrderPoolUserShareMap,
    } as LimitOrderPoolUserShareMap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.sharesOwned = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LimitOrderPoolUserShareMap {
    const message = {
      ...baseLimitOrderPoolUserShareMap,
    } as LimitOrderPoolUserShareMap;
    if (object.count !== undefined && object.count !== null) {
      message.count = String(object.count);
    } else {
      message.count = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.sharesOwned !== undefined && object.sharesOwned !== null) {
      message.sharesOwned = String(object.sharesOwned);
    } else {
      message.sharesOwned = "";
    }
    return message;
  },

  toJSON(message: LimitOrderPoolUserShareMap): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.address !== undefined && (obj.address = message.address);
    message.sharesOwned !== undefined &&
      (obj.sharesOwned = message.sharesOwned);
    return obj;
  },

  fromPartial(
    object: DeepPartial<LimitOrderPoolUserShareMap>
  ): LimitOrderPoolUserShareMap {
    const message = {
      ...baseLimitOrderPoolUserShareMap,
    } as LimitOrderPoolUserShareMap;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.sharesOwned !== undefined && object.sharesOwned !== null) {
      message.sharesOwned = object.sharesOwned;
    } else {
      message.sharesOwned = "";
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
