/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface MsgSingleDeposit {
  creator: string;
  token0: string;
  token1: string;
  price: string;
  fee: string;
  amounts0: string;
  amounts1: string;
  receiver: string;
}

export interface MsgSingleDepositResponse {
  sharesMinted: string;
}

export interface MsgSingleWithdraw {
  creator: string;
  token0: string;
  token1: string;
  price: string;
  fee: string;
  sharesRemoving: string;
  receiver: string;
}

export interface MsgSingleWithdrawResponse {
  amounts0: string;
  amounts1: string;
}

const baseMsgSingleDeposit: object = {
  creator: "",
  token0: "",
  token1: "",
  price: "",
  fee: "",
  amounts0: "",
  amounts1: "",
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
    if (message.fee !== "") {
      writer.uint32(42).string(message.fee);
    }
    if (message.amounts0 !== "") {
      writer.uint32(50).string(message.amounts0);
    }
    if (message.amounts1 !== "") {
      writer.uint32(58).string(message.amounts1);
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
          message.fee = reader.string();
          break;
        case 6:
          message.amounts0 = reader.string();
          break;
        case 7:
          message.amounts1 = reader.string();
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
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.amounts0 !== undefined && object.amounts0 !== null) {
      message.amounts0 = String(object.amounts0);
    } else {
      message.amounts0 = "";
    }
    if (object.amounts1 !== undefined && object.amounts1 !== null) {
      message.amounts1 = String(object.amounts1);
    } else {
      message.amounts1 = "";
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
      message.fee = "";
    }
    if (object.amounts0 !== undefined && object.amounts0 !== null) {
      message.amounts0 = object.amounts0;
    } else {
      message.amounts0 = "";
    }
    if (object.amounts1 !== undefined && object.amounts1 !== null) {
      message.amounts1 = object.amounts1;
    } else {
      message.amounts1 = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    return message;
  },
};

const baseMsgSingleDepositResponse: object = { sharesMinted: "" };

export const MsgSingleDepositResponse = {
  encode(
    message: MsgSingleDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.sharesMinted !== "") {
      writer.uint32(10).string(message.sharesMinted);
    }
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
        case 1:
          message.sharesMinted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSingleDepositResponse {
    const message = {
      ...baseMsgSingleDepositResponse,
    } as MsgSingleDepositResponse;
    if (object.sharesMinted !== undefined && object.sharesMinted !== null) {
      message.sharesMinted = String(object.sharesMinted);
    } else {
      message.sharesMinted = "";
    }
    return message;
  },

  toJSON(message: MsgSingleDepositResponse): unknown {
    const obj: any = {};
    message.sharesMinted !== undefined &&
      (obj.sharesMinted = message.sharesMinted);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSingleDepositResponse>
  ): MsgSingleDepositResponse {
    const message = {
      ...baseMsgSingleDepositResponse,
    } as MsgSingleDepositResponse;
    if (object.sharesMinted !== undefined && object.sharesMinted !== null) {
      message.sharesMinted = object.sharesMinted;
    } else {
      message.sharesMinted = "";
    }
    return message;
  },
};

const baseMsgSingleWithdraw: object = {
  creator: "",
  token0: "",
  token1: "",
  price: "",
  fee: "",
  sharesRemoving: "",
  receiver: "",
};

export const MsgSingleWithdraw = {
  encode(message: MsgSingleWithdraw, writer: Writer = Writer.create()): Writer {
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
    if (message.fee !== "") {
      writer.uint32(42).string(message.fee);
    }
    if (message.sharesRemoving !== "") {
      writer.uint32(50).string(message.sharesRemoving);
    }
    if (message.receiver !== "") {
      writer.uint32(58).string(message.receiver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSingleWithdraw {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSingleWithdraw } as MsgSingleWithdraw;
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
          message.fee = reader.string();
          break;
        case 6:
          message.sharesRemoving = reader.string();
          break;
        case 7:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSingleWithdraw {
    const message = { ...baseMsgSingleWithdraw } as MsgSingleWithdraw;
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
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.sharesRemoving !== undefined && object.sharesRemoving !== null) {
      message.sharesRemoving = String(object.sharesRemoving);
    } else {
      message.sharesRemoving = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    return message;
  },

  toJSON(message: MsgSingleWithdraw): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.token0 !== undefined && (obj.token0 = message.token0);
    message.token1 !== undefined && (obj.token1 = message.token1);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.sharesRemoving !== undefined &&
      (obj.sharesRemoving = message.sharesRemoving);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSingleWithdraw>): MsgSingleWithdraw {
    const message = { ...baseMsgSingleWithdraw } as MsgSingleWithdraw;
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
      message.fee = "";
    }
    if (object.sharesRemoving !== undefined && object.sharesRemoving !== null) {
      message.sharesRemoving = object.sharesRemoving;
    } else {
      message.sharesRemoving = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    return message;
  },
};

const baseMsgSingleWithdrawResponse: object = { amounts0: "", amounts1: "" };

export const MsgSingleWithdrawResponse = {
  encode(
    message: MsgSingleWithdrawResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.amounts0 !== "") {
      writer.uint32(10).string(message.amounts0);
    }
    if (message.amounts1 !== "") {
      writer.uint32(18).string(message.amounts1);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSingleWithdrawResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSingleWithdrawResponse,
    } as MsgSingleWithdrawResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amounts0 = reader.string();
          break;
        case 2:
          message.amounts1 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSingleWithdrawResponse {
    const message = {
      ...baseMsgSingleWithdrawResponse,
    } as MsgSingleWithdrawResponse;
    if (object.amounts0 !== undefined && object.amounts0 !== null) {
      message.amounts0 = String(object.amounts0);
    } else {
      message.amounts0 = "";
    }
    if (object.amounts1 !== undefined && object.amounts1 !== null) {
      message.amounts1 = String(object.amounts1);
    } else {
      message.amounts1 = "";
    }
    return message;
  },

  toJSON(message: MsgSingleWithdrawResponse): unknown {
    const obj: any = {};
    message.amounts0 !== undefined && (obj.amounts0 = message.amounts0);
    message.amounts1 !== undefined && (obj.amounts1 = message.amounts1);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSingleWithdrawResponse>
  ): MsgSingleWithdrawResponse {
    const message = {
      ...baseMsgSingleWithdrawResponse,
    } as MsgSingleWithdrawResponse;
    if (object.amounts0 !== undefined && object.amounts0 !== null) {
      message.amounts0 = object.amounts0;
    } else {
      message.amounts0 = "";
    }
    if (object.amounts1 !== undefined && object.amounts1 !== null) {
      message.amounts1 = object.amounts1;
    } else {
      message.amounts1 = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SingleDeposit(request: MsgSingleDeposit): Promise<MsgSingleDepositResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SingleWithdraw(
    request: MsgSingleWithdraw
  ): Promise<MsgSingleWithdrawResponse>;
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

  SingleWithdraw(
    request: MsgSingleWithdraw
  ): Promise<MsgSingleWithdrawResponse> {
    const data = MsgSingleWithdraw.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "SingleWithdraw",
      data
    );
    return promise.then((data) =>
      MsgSingleWithdrawResponse.decode(new Reader(data))
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
