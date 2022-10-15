/* eslint-disable */
import { Reserve0AndSharesType } from "../dex/reserve_0_and_shares_type";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TickDataType {
  reserve0AndShares: Reserve0AndSharesType[];
  reserve1: string[];
}

const baseTickDataType: object = { reserve1: "" };

export const TickDataType = {
  encode(message: TickDataType, writer: Writer = Writer.create()): Writer {
    for (const v of message.reserve0AndShares) {
      Reserve0AndSharesType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.reserve1) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TickDataType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTickDataType } as TickDataType;
    message.reserve0AndShares = [];
    message.reserve1 = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.reserve0AndShares.push(
            Reserve0AndSharesType.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.reserve1.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TickDataType {
    const message = { ...baseTickDataType } as TickDataType;
    message.reserve0AndShares = [];
    message.reserve1 = [];
    if (
      object.reserve0AndShares !== undefined &&
      object.reserve0AndShares !== null
    ) {
      for (const e of object.reserve0AndShares) {
        message.reserve0AndShares.push(Reserve0AndSharesType.fromJSON(e));
      }
    }
    if (object.reserve1 !== undefined && object.reserve1 !== null) {
      for (const e of object.reserve1) {
        message.reserve1.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: TickDataType): unknown {
    const obj: any = {};
    if (message.reserve0AndShares) {
      obj.reserve0AndShares = message.reserve0AndShares.map((e) =>
        e ? Reserve0AndSharesType.toJSON(e) : undefined
      );
    } else {
      obj.reserve0AndShares = [];
    }
    if (message.reserve1) {
      obj.reserve1 = message.reserve1.map((e) => e);
    } else {
      obj.reserve1 = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<TickDataType>): TickDataType {
    const message = { ...baseTickDataType } as TickDataType;
    message.reserve0AndShares = [];
    message.reserve1 = [];
    if (
      object.reserve0AndShares !== undefined &&
      object.reserve0AndShares !== null
    ) {
      for (const e of object.reserve0AndShares) {
        message.reserve0AndShares.push(Reserve0AndSharesType.fromPartial(e));
      }
    }
    if (object.reserve1 !== undefined && object.reserve1 !== null) {
      for (const e of object.reserve1) {
        message.reserve1.push(e);
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
