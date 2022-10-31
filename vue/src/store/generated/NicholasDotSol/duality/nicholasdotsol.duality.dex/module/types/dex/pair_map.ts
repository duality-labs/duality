/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { TokenPairType } from "../dex/token_pair_type";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface PairMap {
  pairId: string;
  tokenPair: TokenPairType | undefined;
  totalTickCount: number;
}

const basePairMap: object = { pairId: "", totalTickCount: 0 };

export const PairMap = {
  encode(message: PairMap, writer: Writer = Writer.create()): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tokenPair !== undefined) {
      TokenPairType.encode(
        message.tokenPair,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.totalTickCount !== 0) {
      writer.uint32(24).int64(message.totalTickCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PairMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePairMap } as PairMap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairId = reader.string();
          break;
        case 2:
          message.tokenPair = TokenPairType.decode(reader, reader.uint32());
          break;
        case 3:
          message.totalTickCount = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PairMap {
    const message = { ...basePairMap } as PairMap;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tokenPair !== undefined && object.tokenPair !== null) {
      message.tokenPair = TokenPairType.fromJSON(object.tokenPair);
    } else {
      message.tokenPair = undefined;
    }
    if (object.totalTickCount !== undefined && object.totalTickCount !== null) {
      message.totalTickCount = Number(object.totalTickCount);
    } else {
      message.totalTickCount = 0;
    }
    return message;
  },

  toJSON(message: PairMap): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tokenPair !== undefined &&
      (obj.tokenPair = message.tokenPair
        ? TokenPairType.toJSON(message.tokenPair)
        : undefined);
    message.totalTickCount !== undefined &&
      (obj.totalTickCount = message.totalTickCount);
    return obj;
  },

  fromPartial(object: DeepPartial<PairMap>): PairMap {
    const message = { ...basePairMap } as PairMap;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tokenPair !== undefined && object.tokenPair !== null) {
      message.tokenPair = TokenPairType.fromPartial(object.tokenPair);
    } else {
      message.tokenPair = undefined;
    }
    if (object.totalTickCount !== undefined && object.totalTickCount !== null) {
      message.totalTickCount = object.totalTickCount;
    } else {
      message.totalTickCount = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
