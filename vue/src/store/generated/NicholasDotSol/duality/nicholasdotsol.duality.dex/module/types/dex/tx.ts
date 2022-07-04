/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface MsgSingleDeposit {
  creator: string;
  token0: string;
  token1: string;
  price: string;
  fee: number;
  amounts0: number;
  amounts1: number;
  receiver: string;
}

export interface MsgSingleDepositResponse {}

const baseMsgSingleDeposit: object = {
  creator: "",
  token0: "",
  token1: "",
  price: "",
  fee: 0,
  amounts0: 0,
  amounts1: 0,
  receiver: "",
};

export const MsgSingleDeposit = {
  encode(message: MsgSingleDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.token0 !== "") {
      writer.uint32(18).string(message.token0);
    }
    if (message.token1 !== "") {
      writer.uint32(26).string(message.token1);
    }
    if (message.price !== "") {
      writer.uint32(34).string(message.price);
    }
    if (message.fee !== 0) {
      writer.uint32(40).uint64(message.fee);
    }
    if (message.amounts0 !== 0) {
      writer.uint32(48).uint64(message.amounts0);
    }
    if (message.amounts1 !== 0) {
      writer.uint32(56).uint64(message.amounts1);
    }
    if (message.receiver !== "") {
      writer.uint32(66).string(message.receiver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSingleDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSingleDeposit } as MsgSingleDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.token0 = reader.string();
          break;
        case 3:
          message.token1 = reader.string();
          break;
        case 4:
          message.price = reader.string();
          break;
        case 5:
          message.fee = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.amounts0 = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.amounts1 = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSingleDeposit {
    const message = { ...baseMsgSingleDeposit } as MsgSingleDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = Number(object.fee);
    } else {
      message.fee = 0;
    }
    if (object.amounts0 !== undefined && object.amounts0 !== null) {
      message.amounts0 = Number(object.amounts0);
    } else {
      message.amounts0 = 0;
    }
    if (object.amounts1 !== undefined && object.amounts1 !== null) {
      message.amounts1 = Number(object.amounts1);
    } else {
      message.amounts1 = 0;
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    return message;
  },

  toJSON(message: MsgSingleDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.amounts0 !== undefined && (obj.amounts0 = message.amounts0);
    message.amounts1 !== undefined && (obj.amounts1 = message.amounts1);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSingleDeposit>): MsgSingleDeposit {
    const message = { ...baseMsgSingleDeposit } as MsgSingleDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = 0;
    }
    if (object.amounts0 !== undefined && object.amounts0 !== null) {
      message.amounts0 = object.amounts0;
    } else {
      message.amounts0 = 0;
    }
    if (object.amounts1 !== undefined && object.amounts1 !== null) {
      message.amounts1 = object.amounts1;
    } else {
      message.amounts1 = 0;
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    return message;
  },
};

const baseMsgSingleDepositResponse: object = {};

export const MsgSingleDepositResponse = {
  encode(
    _: MsgSingleDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSingleDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSingleDepositResponse,
    } as MsgSingleDepositResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSingleDepositResponse {
    const message = {
      ...baseMsgSingleDepositResponse,
    } as MsgSingleDepositResponse;
    return message;
  },

  toJSON(_: MsgSingleDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSingleDepositResponse>
  ): MsgSingleDepositResponse {
    const message = {
      ...baseMsgSingleDepositResponse,
    } as MsgSingleDepositResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SingleDeposit(request: MsgSingleDeposit): Promise<MsgSingleDepositResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SingleDeposit(request: MsgSingleDeposit): Promise<MsgSingleDepositResponse> {
    const data = MsgSingleDeposit.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "SingleDeposit",
      data
    );
    return promise.then((data) =>
      MsgSingleDepositResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
