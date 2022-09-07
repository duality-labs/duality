/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface MsgDeposit {
  creator: string;
  tokenA: string;
  tokenB: string;
  amount0: string;
  amount1: string;
  priceIndex: string;
  fee: string;
}

export interface MsgDepositResponse {}

export interface MsgWithdrawl {
  creator: string;
  tokenA: string;
  tokenB: string;
  sharesToRemove: string;
  priceIndex: string;
  fee: string;
  receiver: string;
}

export interface MsgWithdrawlResponse {}

export interface MsgSwap {
  creator: string;
  amountIn: string;
  tokenIn: string;
  slippageTolerance: string;
}

export interface MsgSwapResponse {}

const baseMsgDeposit: object = {
  creator: "",
  tokenA: "",
  tokenB: "",
  amount0: "",
  amount1: "",
  priceIndex: "",
  fee: "",
};

export const MsgDeposit = {
  encode(message: MsgDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenA !== "") {
      writer.uint32(18).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(26).string(message.tokenB);
    }
    if (message.amount0 !== "") {
      writer.uint32(34).string(message.amount0);
    }
    if (message.amount1 !== "") {
      writer.uint32(42).string(message.amount1);
    }
    if (message.priceIndex !== "") {
      writer.uint32(50).string(message.priceIndex);
    }
    if (message.fee !== "") {
      writer.uint32(58).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeposit } as MsgDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenA = reader.string();
          break;
        case 3:
          message.tokenB = reader.string();
          break;
        case 4:
          message.amount0 = reader.string();
          break;
        case 5:
          message.amount1 = reader.string();
          break;
        case 6:
          message.priceIndex = reader.string();
          break;
        case 7:
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = String(object.tokenA);
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = String(object.tokenB);
    } else {
      message.tokenB = "";
    }
    if (object.amount0 !== undefined && object.amount0 !== null) {
      message.amount0 = String(object.amount0);
    } else {
      message.amount0 = "";
    }
    if (object.amount1 !== undefined && object.amount1 !== null) {
      message.amount1 = String(object.amount1);
    } else {
      message.amount1 = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = String(object.priceIndex);
    } else {
      message.priceIndex = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: MsgDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.amount0 !== undefined && (obj.amount0 = message.amount0);
    message.amount1 !== undefined && (obj.amount1 = message.amount1);
    message.priceIndex !== undefined && (obj.priceIndex = message.priceIndex);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeposit>): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = object.tokenA;
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = object.tokenB;
    } else {
      message.tokenB = "";
    }
    if (object.amount0 !== undefined && object.amount0 !== null) {
      message.amount0 = object.amount0;
    } else {
      message.amount0 = "";
    }
    if (object.amount1 !== undefined && object.amount1 !== null) {
      message.amount1 = object.amount1;
    } else {
      message.amount1 = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = object.priceIndex;
    } else {
      message.priceIndex = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

const baseMsgDepositResponse: object = {};

export const MsgDepositResponse = {
  encode(_: MsgDepositResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
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

  fromJSON(_: any): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    return message;
  },

  toJSON(_: MsgDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDepositResponse>): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    return message;
  },
};

const baseMsgWithdrawl: object = {
  creator: "",
  tokenA: "",
  tokenB: "",
  sharesToRemove: "",
  priceIndex: "",
  fee: "",
  receiver: "",
};

export const MsgWithdrawl = {
  encode(message: MsgWithdrawl, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenA !== "") {
      writer.uint32(18).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(26).string(message.tokenB);
    }
    if (message.sharesToRemove !== "") {
      writer.uint32(34).string(message.sharesToRemove);
    }
    if (message.priceIndex !== "") {
      writer.uint32(42).string(message.priceIndex);
    }
    if (message.fee !== "") {
      writer.uint32(50).string(message.fee);
    }
    if (message.receiver !== "") {
      writer.uint32(58).string(message.receiver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawl {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawl } as MsgWithdrawl;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenA = reader.string();
          break;
        case 3:
          message.tokenB = reader.string();
          break;
        case 4:
          message.sharesToRemove = reader.string();
          break;
        case 5:
          message.priceIndex = reader.string();
          break;
        case 6:
          message.fee = reader.string();
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

  fromJSON(object: any): MsgWithdrawl {
    const message = { ...baseMsgWithdrawl } as MsgWithdrawl;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = String(object.tokenA);
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = String(object.tokenB);
    } else {
      message.tokenB = "";
    }
    if (object.sharesToRemove !== undefined && object.sharesToRemove !== null) {
      message.sharesToRemove = String(object.sharesToRemove);
    } else {
      message.sharesToRemove = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = String(object.priceIndex);
    } else {
      message.priceIndex = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    return message;
  },

  toJSON(message: MsgWithdrawl): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.sharesToRemove !== undefined &&
      (obj.sharesToRemove = message.sharesToRemove);
    message.priceIndex !== undefined && (obj.priceIndex = message.priceIndex);
    message.fee !== undefined && (obj.fee = message.fee);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWithdrawl>): MsgWithdrawl {
    const message = { ...baseMsgWithdrawl } as MsgWithdrawl;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = object.tokenA;
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = object.tokenB;
    } else {
      message.tokenB = "";
    }
    if (object.sharesToRemove !== undefined && object.sharesToRemove !== null) {
      message.sharesToRemove = object.sharesToRemove;
    } else {
      message.sharesToRemove = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = object.priceIndex;
    } else {
      message.priceIndex = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    return message;
  },
};

const baseMsgWithdrawlResponse: object = {};

export const MsgWithdrawlResponse = {
  encode(_: MsgWithdrawlResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawlResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawlResponse } as MsgWithdrawlResponse;
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

  fromJSON(_: any): MsgWithdrawlResponse {
    const message = { ...baseMsgWithdrawlResponse } as MsgWithdrawlResponse;
    return message;
  },

  toJSON(_: MsgWithdrawlResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgWithdrawlResponse>): MsgWithdrawlResponse {
    const message = { ...baseMsgWithdrawlResponse } as MsgWithdrawlResponse;
    return message;
  },
};

const baseMsgSwap: object = {
  creator: "",
  amountIn: "",
  tokenIn: "",
  slippageTolerance: "",
};

export const MsgSwap = {
  encode(message: MsgSwap, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amountIn !== "") {
      writer.uint32(18).string(message.amountIn);
    }
    if (message.tokenIn !== "") {
      writer.uint32(26).string(message.tokenIn);
    }
    if (message.slippageTolerance !== "") {
      writer.uint32(34).string(message.slippageTolerance);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwap } as MsgSwap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amountIn = reader.string();
          break;
        case 3:
          message.tokenIn = reader.string();
          break;
        case 4:
          message.slippageTolerance = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = String(object.amountIn);
    } else {
      message.amountIn = "";
    }
    if (object.tokenIn !== undefined && object.tokenIn !== null) {
      message.tokenIn = String(object.tokenIn);
    } else {
      message.tokenIn = "";
    }
    if (
      object.slippageTolerance !== undefined &&
      object.slippageTolerance !== null
    ) {
      message.slippageTolerance = String(object.slippageTolerance);
    } else {
      message.slippageTolerance = "";
    }
    return message;
  },

  toJSON(message: MsgSwap): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amountIn !== undefined && (obj.amountIn = message.amountIn);
    message.tokenIn !== undefined && (obj.tokenIn = message.tokenIn);
    message.slippageTolerance !== undefined &&
      (obj.slippageTolerance = message.slippageTolerance);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwap>): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = object.amountIn;
    } else {
      message.amountIn = "";
    }
    if (object.tokenIn !== undefined && object.tokenIn !== null) {
      message.tokenIn = object.tokenIn;
    } else {
      message.tokenIn = "";
    }
    if (
      object.slippageTolerance !== undefined &&
      object.slippageTolerance !== null
    ) {
      message.slippageTolerance = object.slippageTolerance;
    } else {
      message.slippageTolerance = "";
    }
    return message;
  },
};

const baseMsgSwapResponse: object = {};

export const MsgSwapResponse = {
  encode(_: MsgSwapResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
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

  fromJSON(_: any): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    return message;
  },

  toJSON(_: MsgSwapResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSwapResponse>): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse>;
  Withdrawl(request: MsgWithdrawl): Promise<MsgWithdrawlResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Swap(request: MsgSwap): Promise<MsgSwapResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse> {
    const data = MsgDeposit.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "Deposit",
      data
    );
    return promise.then((data) => MsgDepositResponse.decode(new Reader(data)));
  }

  Withdrawl(request: MsgWithdrawl): Promise<MsgWithdrawlResponse> {
    const data = MsgWithdrawl.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "Withdrawl",
      data
    );
    return promise.then((data) =>
      MsgWithdrawlResponse.decode(new Reader(data))
    );
  }

  Swap(request: MsgSwap): Promise<MsgSwapResponse> {
    const data = MsgSwap.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "Swap",
      data
    );
    return promise.then((data) => MsgSwapResponse.decode(new Reader(data)));
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
