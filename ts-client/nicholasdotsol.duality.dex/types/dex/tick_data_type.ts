/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Reserve0AndSharesType } from "./reserve_0_and_shares_type";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface TickDataType {
  reserve0AndShares: Reserve0AndSharesType[];
  reserve1: string[];
}

function createBaseTickDataType(): TickDataType {
  return { reserve0AndShares: [], reserve1: [] };
}

export const TickDataType = {
  encode(message: TickDataType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.reserve0AndShares) {
      Reserve0AndSharesType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.reserve1) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TickDataType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTickDataType();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.reserve0AndShares.push(Reserve0AndSharesType.decode(reader, reader.uint32()));
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
    return {
      reserve0AndShares: Array.isArray(object?.reserve0AndShares)
        ? object.reserve0AndShares.map((e: any) => Reserve0AndSharesType.fromJSON(e))
        : [],
      reserve1: Array.isArray(object?.reserve1) ? object.reserve1.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: TickDataType): unknown {
    const obj: any = {};
    if (message.reserve0AndShares) {
      obj.reserve0AndShares = message.reserve0AndShares.map((e) => e ? Reserve0AndSharesType.toJSON(e) : undefined);
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

  fromPartial<I extends Exact<DeepPartial<TickDataType>, I>>(object: I): TickDataType {
    const message = createBaseTickDataType();
    message.reserve0AndShares = object.reserve0AndShares?.map((e) => Reserve0AndSharesType.fromPartial(e)) || [];
    message.reserve1 = object.reserve1?.map((e) => e) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };
