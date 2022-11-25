/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface Reserve0AndSharesType {
  reserve0: string;
  totalShares: string;
}

function createBaseReserve0AndSharesType(): Reserve0AndSharesType {
  return { reserve0: "", totalShares: "" };
}

export const Reserve0AndSharesType = {
  encode(message: Reserve0AndSharesType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.reserve0 !== "") {
      writer.uint32(10).string(message.reserve0);
    }
    if (message.totalShares !== "") {
      writer.uint32(18).string(message.totalShares);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Reserve0AndSharesType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReserve0AndSharesType();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.reserve0 = reader.string();
          break;
        case 2:
          message.totalShares = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Reserve0AndSharesType {
    return {
      reserve0: isSet(object.reserve0) ? String(object.reserve0) : "",
      totalShares: isSet(object.totalShares) ? String(object.totalShares) : "",
    };
  },

  toJSON(message: Reserve0AndSharesType): unknown {
    const obj: any = {};
    message.reserve0 !== undefined && (obj.reserve0 = message.reserve0);
    message.totalShares !== undefined && (obj.totalShares = message.totalShares);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Reserve0AndSharesType>, I>>(object: I): Reserve0AndSharesType {
    const message = createBaseReserve0AndSharesType();
    message.reserve0 = object.reserve0 ?? "";
    message.totalShares = object.totalShares ?? "";
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
