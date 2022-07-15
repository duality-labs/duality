/* eslint-disable */
import { Pool } from "../dex/pool";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Ticks {
  token0: string;
  token1: string;
  poolsZeroToOne: Pool | undefined;
  poolsOneToZero: Pool | undefined;
}

const baseTicks: object = { token0: "", token1: "" };

export const Ticks = {
  encode(message: Ticks, writer: Writer = Writer.create()): Writer {
    if (message.token0 !== "") {
      writer.uint32(10).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(18).string(message.token1);
    }
    if (message.poolsZeroToOne !== undefined) {
      Pool.encode(message.poolsZeroToOne, writer.uint32(26).fork()).ldelim();
    }
    if (message.poolsOneToZero !== undefined) {
      Pool.encode(message.poolsOneToZero, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Ticks {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTicks } as Ticks;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token0 = reader.string();
          break;
        case 2:
          message.token1 = reader.string();
          break;
        case 3:
          message.poolsZeroToOne = Pool.decode(reader, reader.uint32());
          break;
        case 4:
          message.poolsOneToZero = Pool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Ticks {
    const message = { ...baseTicks } as Ticks;
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = String(object.token0);
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = String(object.token1);
    } else {
      message.token1 = "";
    }
    if (object.poolsZeroToOne !== undefined && object.poolsZeroToOne !== null) {
      message.poolsZeroToOne = Pool.fromJSON(object.poolsZeroToOne);
    } else {
      message.poolsZeroToOne = undefined;
    }
    if (object.poolsOneToZero !== undefined && object.poolsOneToZero !== null) {
      message.poolsOneToZero = Pool.fromJSON(object.poolsOneToZero);
    } else {
      message.poolsOneToZero = undefined;
    }
    return message;
  },

  toJSON(message: Ticks): unknown {
    const obj: any = {};
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.poolsZeroToOne !== undefined &&
      (obj.poolsZeroToOne = message.poolsZeroToOne
        ? Pool.toJSON(message.poolsZeroToOne)
        : undefined);
    message.poolsOneToZero !== undefined &&
      (obj.poolsOneToZero = message.poolsOneToZero
        ? Pool.toJSON(message.poolsOneToZero)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Ticks>): Ticks {
    const message = { ...baseTicks } as Ticks;
    if (object.token0 !== undefined && object.token0 !== null) {
      message.token0 = object.token0;
    } else {
      message.token0 = "";
    }
    if (object.token1 !== undefined && object.token1 !== null) {
      message.token1 = object.token1;
    } else {
      message.token1 = "";
    }
    if (object.poolsZeroToOne !== undefined && object.poolsZeroToOne !== null) {
      message.poolsZeroToOne = Pool.fromPartial(object.poolsZeroToOne);
    } else {
      message.poolsZeroToOne = undefined;
    }
    if (object.poolsOneToZero !== undefined && object.poolsOneToZero !== null) {
      message.poolsOneToZero = Pool.fromPartial(object.poolsOneToZero);
    } else {
      message.poolsOneToZero = undefined;
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
