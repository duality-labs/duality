/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface MsgDeposit {
  creator: string;
  receiver: string;
  tokenA: string;
  tokenB: string;
  amountsA: string[];
  amountsB: string[];
  tickIndexes: number[];
  feeIndexes: number[];
}

export interface MsgDepositResponse {
  Reserve0Deposited: string[];
  Reserve1Deposited: string[];
}

export interface MsgWithdrawl {
  creator: string;
  receiver: string;
  tokenA: string;
  tokenB: string;
  sharesToRemove: string[];
  tickIndexes: number[];
  feeIndexes: number[];
}

export interface MsgWithdrawlResponse {}

export interface MsgSwap {
  creator: string;
  receiver: string;
  tokenA: string;
  tokenB: string;
  amountIn: string;
  tokenIn: string;
  minOut: string;
}

export interface MsgSwapResponse {
  coinOut: Coin | undefined;
}

export interface MsgPlaceLimitOrder {
  creator: string;
  receiver: string;
  tokenA: string;
  tokenB: string;
  tickIndex: number;
  tokenIn: string;
  amountIn: string;
}

export interface MsgPlaceLimitOrderResponse {}

export interface MsgWithdrawFilledLimitOrder {
  creator: string;
  receiver: string;
  tokenA: string;
  tokenB: string;
  tickIndex: number;
  keyToken: string;
  key: number;
}

export interface MsgWithdrawFilledLimitOrderResponse {}

export interface MsgCancelLimitOrder {
  creator: string;
  receiver: string;
  tokenA: string;
  tokenB: string;
  tickIndex: number;
  keyToken: string;
  key: number;
  sharesOut: string;
}

export interface MsgCancelLimitOrderResponse {}

const baseMsgDeposit: object = {
  creator: "",
  receiver: "",
  tokenA: "",
  tokenB: "",
  amountsA: "",
  amountsB: "",
  tickIndexes: 0,
  feeIndexes: 0,
};

export const MsgDeposit = {
  encode(message: MsgDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.tokenA !== "") {
      writer.uint32(26).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(34).string(message.tokenB);
    }
    for (const v of message.amountsA) {
      writer.uint32(42).string(v!);
    }
    for (const v of message.amountsB) {
      writer.uint32(50).string(v!);
    }
    writer.uint32(58).fork();
    for (const v of message.tickIndexes) {
      writer.int64(v);
    }
    writer.ldelim();
    writer.uint32(66).fork();
    for (const v of message.feeIndexes) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeposit } as MsgDeposit;
    message.amountsA = [];
    message.amountsB = [];
    message.tickIndexes = [];
    message.feeIndexes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.tokenA = reader.string();
          break;
        case 4:
          message.tokenB = reader.string();
          break;
        case 5:
          message.amountsA.push(reader.string());
          break;
        case 6:
          message.amountsB.push(reader.string());
          break;
        case 7:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.tickIndexes.push(longToNumber(reader.int64() as Long));
            }
          } else {
            message.tickIndexes.push(longToNumber(reader.int64() as Long));
          }
          break;
        case 8:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.feeIndexes.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.feeIndexes.push(longToNumber(reader.uint64() as Long));
          }
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
    message.amountsA = [];
    message.amountsB = [];
    message.tickIndexes = [];
    message.feeIndexes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
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
    if (object.amountsA !== undefined && object.amountsA !== null) {
      for (const e of object.amountsA) {
        message.amountsA.push(String(e));
      }
    }
    if (object.amountsB !== undefined && object.amountsB !== null) {
      for (const e of object.amountsB) {
        message.amountsB.push(String(e));
      }
    }
    if (object.tickIndexes !== undefined && object.tickIndexes !== null) {
      for (const e of object.tickIndexes) {
        message.tickIndexes.push(Number(e));
      }
    }
    if (object.feeIndexes !== undefined && object.feeIndexes !== null) {
      for (const e of object.feeIndexes) {
        message.feeIndexes.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: MsgDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    if (message.amountsA) {
      obj.amountsA = message.amountsA.map((e) => e);
    } else {
      obj.amountsA = [];
    }
    if (message.amountsB) {
      obj.amountsB = message.amountsB.map((e) => e);
    } else {
      obj.amountsB = [];
    }
    if (message.tickIndexes) {
      obj.tickIndexes = message.tickIndexes.map((e) => e);
    } else {
      obj.tickIndexes = [];
    }
    if (message.feeIndexes) {
      obj.feeIndexes = message.feeIndexes.map((e) => e);
    } else {
      obj.feeIndexes = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeposit>): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    message.amountsA = [];
    message.amountsB = [];
    message.tickIndexes = [];
    message.feeIndexes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
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
    if (object.amountsA !== undefined && object.amountsA !== null) {
      for (const e of object.amountsA) {
        message.amountsA.push(e);
      }
    }
    if (object.amountsB !== undefined && object.amountsB !== null) {
      for (const e of object.amountsB) {
        message.amountsB.push(e);
      }
    }
    if (object.tickIndexes !== undefined && object.tickIndexes !== null) {
      for (const e of object.tickIndexes) {
        message.tickIndexes.push(e);
      }
    }
    if (object.feeIndexes !== undefined && object.feeIndexes !== null) {
      for (const e of object.feeIndexes) {
        message.feeIndexes.push(e);
      }
    }
    return message;
  },
};

const baseMsgDepositResponse: object = {
  Reserve0Deposited: "",
  Reserve1Deposited: "",
};

export const MsgDepositResponse = {
  encode(
    message: MsgDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Reserve0Deposited) {
      writer.uint32(10).string(v!);
    }
    for (const v of message.Reserve1Deposited) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    message.Reserve0Deposited = [];
    message.Reserve1Deposited = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Reserve0Deposited.push(reader.string());
          break;
        case 2:
          message.Reserve1Deposited.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    message.Reserve0Deposited = [];
    message.Reserve1Deposited = [];
    if (
      object.Reserve0Deposited !== undefined &&
      object.Reserve0Deposited !== null
    ) {
      for (const e of object.Reserve0Deposited) {
        message.Reserve0Deposited.push(String(e));
      }
    }
    if (
      object.Reserve1Deposited !== undefined &&
      object.Reserve1Deposited !== null
    ) {
      for (const e of object.Reserve1Deposited) {
        message.Reserve1Deposited.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgDepositResponse): unknown {
    const obj: any = {};
    if (message.Reserve0Deposited) {
      obj.Reserve0Deposited = message.Reserve0Deposited.map((e) => e);
    } else {
      obj.Reserve0Deposited = [];
    }
    if (message.Reserve1Deposited) {
      obj.Reserve1Deposited = message.Reserve1Deposited.map((e) => e);
    } else {
      obj.Reserve1Deposited = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDepositResponse>): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    message.Reserve0Deposited = [];
    message.Reserve1Deposited = [];
    if (
      object.Reserve0Deposited !== undefined &&
      object.Reserve0Deposited !== null
    ) {
      for (const e of object.Reserve0Deposited) {
        message.Reserve0Deposited.push(e);
      }
    }
    if (
      object.Reserve1Deposited !== undefined &&
      object.Reserve1Deposited !== null
    ) {
      for (const e of object.Reserve1Deposited) {
        message.Reserve1Deposited.push(e);
      }
    }
    return message;
  },
};

const baseMsgWithdrawl: object = {
  creator: "",
  receiver: "",
  tokenA: "",
  tokenB: "",
  sharesToRemove: "",
  tickIndexes: 0,
  feeIndexes: 0,
};

export const MsgWithdrawl = {
  encode(message: MsgWithdrawl, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.tokenA !== "") {
      writer.uint32(26).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(34).string(message.tokenB);
    }
    for (const v of message.sharesToRemove) {
      writer.uint32(42).string(v!);
    }
    writer.uint32(50).fork();
    for (const v of message.tickIndexes) {
      writer.int64(v);
    }
    writer.ldelim();
    writer.uint32(58).fork();
    for (const v of message.feeIndexes) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawl {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawl } as MsgWithdrawl;
    message.sharesToRemove = [];
    message.tickIndexes = [];
    message.feeIndexes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.tokenA = reader.string();
          break;
        case 4:
          message.tokenB = reader.string();
          break;
        case 5:
          message.sharesToRemove.push(reader.string());
          break;
        case 6:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.tickIndexes.push(longToNumber(reader.int64() as Long));
            }
          } else {
            message.tickIndexes.push(longToNumber(reader.int64() as Long));
          }
          break;
        case 7:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.feeIndexes.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.feeIndexes.push(longToNumber(reader.uint64() as Long));
          }
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
    message.sharesToRemove = [];
    message.tickIndexes = [];
    message.feeIndexes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
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
      for (const e of object.sharesToRemove) {
        message.sharesToRemove.push(String(e));
      }
    }
    if (object.tickIndexes !== undefined && object.tickIndexes !== null) {
      for (const e of object.tickIndexes) {
        message.tickIndexes.push(Number(e));
      }
    }
    if (object.feeIndexes !== undefined && object.feeIndexes !== null) {
      for (const e of object.feeIndexes) {
        message.feeIndexes.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: MsgWithdrawl): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    if (message.sharesToRemove) {
      obj.sharesToRemove = message.sharesToRemove.map((e) => e);
    } else {
      obj.sharesToRemove = [];
    }
    if (message.tickIndexes) {
      obj.tickIndexes = message.tickIndexes.map((e) => e);
    } else {
      obj.tickIndexes = [];
    }
    if (message.feeIndexes) {
      obj.feeIndexes = message.feeIndexes.map((e) => e);
    } else {
      obj.feeIndexes = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWithdrawl>): MsgWithdrawl {
    const message = { ...baseMsgWithdrawl } as MsgWithdrawl;
    message.sharesToRemove = [];
    message.tickIndexes = [];
    message.feeIndexes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
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
      for (const e of object.sharesToRemove) {
        message.sharesToRemove.push(e);
      }
    }
    if (object.tickIndexes !== undefined && object.tickIndexes !== null) {
      for (const e of object.tickIndexes) {
        message.tickIndexes.push(e);
      }
    }
    if (object.feeIndexes !== undefined && object.feeIndexes !== null) {
      for (const e of object.feeIndexes) {
        message.feeIndexes.push(e);
      }
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
  receiver: "",
  tokenA: "",
  tokenB: "",
  amountIn: "",
  tokenIn: "",
  minOut: "",
};

export const MsgSwap = {
  encode(message: MsgSwap, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.tokenA !== "") {
      writer.uint32(26).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(34).string(message.tokenB);
    }
    if (message.amountIn !== "") {
      writer.uint32(42).string(message.amountIn);
    }
    if (message.tokenIn !== "") {
      writer.uint32(50).string(message.tokenIn);
    }
    if (message.minOut !== "") {
      writer.uint32(58).string(message.minOut);
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
          message.receiver = reader.string();
          break;
        case 3:
          message.tokenA = reader.string();
          break;
        case 4:
          message.tokenB = reader.string();
          break;
        case 5:
          message.amountIn = reader.string();
          break;
        case 6:
          message.tokenIn = reader.string();
          break;
        case 7:
          message.minOut = reader.string();
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
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
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
    if (object.minOut !== undefined && object.minOut !== null) {
      message.minOut = String(object.minOut);
    } else {
      message.minOut = "";
    }
    return message;
  },

  toJSON(message: MsgSwap): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.amountIn !== undefined && (obj.amountIn = message.amountIn);
    message.tokenIn !== undefined && (obj.tokenIn = message.tokenIn);
    message.minOut !== undefined && (obj.minOut = message.minOut);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwap>): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
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
    if (object.minOut !== undefined && object.minOut !== null) {
      message.minOut = object.minOut;
    } else {
      message.minOut = "";
    }
    return message;
  },
};

const baseMsgSwapResponse: object = {};

export const MsgSwapResponse = {
  encode(message: MsgSwapResponse, writer: Writer = Writer.create()): Writer {
    if (message.coinOut !== undefined) {
      Coin.encode(message.coinOut, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.coinOut = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    if (object.coinOut !== undefined && object.coinOut !== null) {
      message.coinOut = Coin.fromJSON(object.coinOut);
    } else {
      message.coinOut = undefined;
    }
    return message;
  },

  toJSON(message: MsgSwapResponse): unknown {
    const obj: any = {};
    message.coinOut !== undefined &&
      (obj.coinOut = message.coinOut
        ? Coin.toJSON(message.coinOut)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwapResponse>): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    if (object.coinOut !== undefined && object.coinOut !== null) {
      message.coinOut = Coin.fromPartial(object.coinOut);
    } else {
      message.coinOut = undefined;
    }
    return message;
  },
};

const baseMsgPlaceLimitOrder: object = {
  creator: "",
  receiver: "",
  tokenA: "",
  tokenB: "",
  tickIndex: 0,
  tokenIn: "",
  amountIn: "",
};

export const MsgPlaceLimitOrder = {
  encode(
    message: MsgPlaceLimitOrder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.tokenA !== "") {
      writer.uint32(26).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(34).string(message.tokenB);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(40).int64(message.tickIndex);
    }
    if (message.tokenIn !== "") {
      writer.uint32(50).string(message.tokenIn);
    }
    if (message.amountIn !== "") {
      writer.uint32(58).string(message.amountIn);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPlaceLimitOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPlaceLimitOrder } as MsgPlaceLimitOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.tokenA = reader.string();
          break;
        case 4:
          message.tokenB = reader.string();
          break;
        case 5:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.tokenIn = reader.string();
          break;
        case 7:
          message.amountIn = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlaceLimitOrder {
    const message = { ...baseMsgPlaceLimitOrder } as MsgPlaceLimitOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
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
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.tokenIn !== undefined && object.tokenIn !== null) {
      message.tokenIn = String(object.tokenIn);
    } else {
      message.tokenIn = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = String(object.amountIn);
    } else {
      message.amountIn = "";
    }
    return message;
  },

  toJSON(message: MsgPlaceLimitOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.tokenIn !== undefined && (obj.tokenIn = message.tokenIn);
    message.amountIn !== undefined && (obj.amountIn = message.amountIn);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPlaceLimitOrder>): MsgPlaceLimitOrder {
    const message = { ...baseMsgPlaceLimitOrder } as MsgPlaceLimitOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
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
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.tokenIn !== undefined && object.tokenIn !== null) {
      message.tokenIn = object.tokenIn;
    } else {
      message.tokenIn = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = object.amountIn;
    } else {
      message.amountIn = "";
    }
    return message;
  },
};

const baseMsgPlaceLimitOrderResponse: object = {};

export const MsgPlaceLimitOrderResponse = {
  encode(
    _: MsgPlaceLimitOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgPlaceLimitOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgPlaceLimitOrderResponse,
    } as MsgPlaceLimitOrderResponse;
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

  fromJSON(_: any): MsgPlaceLimitOrderResponse {
    const message = {
      ...baseMsgPlaceLimitOrderResponse,
    } as MsgPlaceLimitOrderResponse;
    return message;
  },

  toJSON(_: MsgPlaceLimitOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgPlaceLimitOrderResponse>
  ): MsgPlaceLimitOrderResponse {
    const message = {
      ...baseMsgPlaceLimitOrderResponse,
    } as MsgPlaceLimitOrderResponse;
    return message;
  },
};

const baseMsgWithdrawFilledLimitOrder: object = {
  creator: "",
  receiver: "",
  tokenA: "",
  tokenB: "",
  tickIndex: 0,
  keyToken: "",
  key: 0,
};

export const MsgWithdrawFilledLimitOrder = {
  encode(
    message: MsgWithdrawFilledLimitOrder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.tokenA !== "") {
      writer.uint32(26).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(34).string(message.tokenB);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(40).int64(message.tickIndex);
    }
    if (message.keyToken !== "") {
      writer.uint32(50).string(message.keyToken);
    }
    if (message.key !== 0) {
      writer.uint32(56).uint64(message.key);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWithdrawFilledLimitOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawFilledLimitOrder,
    } as MsgWithdrawFilledLimitOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.tokenA = reader.string();
          break;
        case 4:
          message.tokenB = reader.string();
          break;
        case 5:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.keyToken = reader.string();
          break;
        case 7:
          message.key = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawFilledLimitOrder {
    const message = {
      ...baseMsgWithdrawFilledLimitOrder,
    } as MsgWithdrawFilledLimitOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
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
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.keyToken !== undefined && object.keyToken !== null) {
      message.keyToken = String(object.keyToken);
    } else {
      message.keyToken = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = Number(object.key);
    } else {
      message.key = 0;
    }
    return message;
  },

  toJSON(message: MsgWithdrawFilledLimitOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.keyToken !== undefined && (obj.keyToken = message.keyToken);
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgWithdrawFilledLimitOrder>
  ): MsgWithdrawFilledLimitOrder {
    const message = {
      ...baseMsgWithdrawFilledLimitOrder,
    } as MsgWithdrawFilledLimitOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
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
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.keyToken !== undefined && object.keyToken !== null) {
      message.keyToken = object.keyToken;
    } else {
      message.keyToken = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = 0;
    }
    return message;
  },
};

const baseMsgWithdrawFilledLimitOrderResponse: object = {};

export const MsgWithdrawFilledLimitOrderResponse = {
  encode(
    _: MsgWithdrawFilledLimitOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWithdrawFilledLimitOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawFilledLimitOrderResponse,
    } as MsgWithdrawFilledLimitOrderResponse;
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

  fromJSON(_: any): MsgWithdrawFilledLimitOrderResponse {
    const message = {
      ...baseMsgWithdrawFilledLimitOrderResponse,
    } as MsgWithdrawFilledLimitOrderResponse;
    return message;
  },

  toJSON(_: MsgWithdrawFilledLimitOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgWithdrawFilledLimitOrderResponse>
  ): MsgWithdrawFilledLimitOrderResponse {
    const message = {
      ...baseMsgWithdrawFilledLimitOrderResponse,
    } as MsgWithdrawFilledLimitOrderResponse;
    return message;
  },
};

const baseMsgCancelLimitOrder: object = {
  creator: "",
  receiver: "",
  tokenA: "",
  tokenB: "",
  tickIndex: 0,
  keyToken: "",
  key: 0,
  sharesOut: "",
};

export const MsgCancelLimitOrder = {
  encode(
    message: MsgCancelLimitOrder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.tokenA !== "") {
      writer.uint32(26).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(34).string(message.tokenB);
    }
    if (message.tickIndex !== 0) {
      writer.uint32(40).int64(message.tickIndex);
    }
    if (message.keyToken !== "") {
      writer.uint32(50).string(message.keyToken);
    }
    if (message.key !== 0) {
      writer.uint32(56).uint64(message.key);
    }
    if (message.sharesOut !== "") {
      writer.uint32(66).string(message.sharesOut);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelLimitOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelLimitOrder } as MsgCancelLimitOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.tokenA = reader.string();
          break;
        case 4:
          message.tokenB = reader.string();
          break;
        case 5:
          message.tickIndex = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.keyToken = reader.string();
          break;
        case 7:
          message.key = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.sharesOut = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelLimitOrder {
    const message = { ...baseMsgCancelLimitOrder } as MsgCancelLimitOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
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
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = Number(object.tickIndex);
    } else {
      message.tickIndex = 0;
    }
    if (object.keyToken !== undefined && object.keyToken !== null) {
      message.keyToken = String(object.keyToken);
    } else {
      message.keyToken = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = Number(object.key);
    } else {
      message.key = 0;
    }
    if (object.sharesOut !== undefined && object.sharesOut !== null) {
      message.sharesOut = String(object.sharesOut);
    } else {
      message.sharesOut = "";
    }
    return message;
  },

  toJSON(message: MsgCancelLimitOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.tickIndex !== undefined && (obj.tickIndex = message.tickIndex);
    message.keyToken !== undefined && (obj.keyToken = message.keyToken);
    message.key !== undefined && (obj.key = message.key);
    message.sharesOut !== undefined && (obj.sharesOut = message.sharesOut);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCancelLimitOrder>): MsgCancelLimitOrder {
    const message = { ...baseMsgCancelLimitOrder } as MsgCancelLimitOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
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
    if (object.tickIndex !== undefined && object.tickIndex !== null) {
      message.tickIndex = object.tickIndex;
    } else {
      message.tickIndex = 0;
    }
    if (object.keyToken !== undefined && object.keyToken !== null) {
      message.keyToken = object.keyToken;
    } else {
      message.keyToken = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = 0;
    }
    if (object.sharesOut !== undefined && object.sharesOut !== null) {
      message.sharesOut = object.sharesOut;
    } else {
      message.sharesOut = "";
    }
    return message;
  },
};

const baseMsgCancelLimitOrderResponse: object = {};

export const MsgCancelLimitOrderResponse = {
  encode(
    _: MsgCancelLimitOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCancelLimitOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCancelLimitOrderResponse,
    } as MsgCancelLimitOrderResponse;
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

  fromJSON(_: any): MsgCancelLimitOrderResponse {
    const message = {
      ...baseMsgCancelLimitOrderResponse,
    } as MsgCancelLimitOrderResponse;
    return message;
  },

  toJSON(_: MsgCancelLimitOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCancelLimitOrderResponse>
  ): MsgCancelLimitOrderResponse {
    const message = {
      ...baseMsgCancelLimitOrderResponse,
    } as MsgCancelLimitOrderResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse>;
  Withdrawl(request: MsgWithdrawl): Promise<MsgWithdrawlResponse>;
  Swap(request: MsgSwap): Promise<MsgSwapResponse>;
  PlaceLimitOrder(
    request: MsgPlaceLimitOrder
  ): Promise<MsgPlaceLimitOrderResponse>;
  WithdrawFilledLimitOrder(
    request: MsgWithdrawFilledLimitOrder
  ): Promise<MsgWithdrawFilledLimitOrderResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CancelLimitOrder(
    request: MsgCancelLimitOrder
  ): Promise<MsgCancelLimitOrderResponse>;
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

  PlaceLimitOrder(
    request: MsgPlaceLimitOrder
  ): Promise<MsgPlaceLimitOrderResponse> {
    const data = MsgPlaceLimitOrder.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "PlaceLimitOrder",
      data
    );
    return promise.then((data) =>
      MsgPlaceLimitOrderResponse.decode(new Reader(data))
    );
  }

  WithdrawFilledLimitOrder(
    request: MsgWithdrawFilledLimitOrder
  ): Promise<MsgWithdrawFilledLimitOrderResponse> {
    const data = MsgWithdrawFilledLimitOrder.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "WithdrawFilledLimitOrder",
      data
    );
    return promise.then((data) =>
      MsgWithdrawFilledLimitOrderResponse.decode(new Reader(data))
    );
  }

  CancelLimitOrder(
    request: MsgCancelLimitOrder
  ): Promise<MsgCancelLimitOrderResponse> {
    const data = MsgCancelLimitOrder.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "CancelLimitOrder",
      data
    );
    return promise.then((data) =>
      MsgCancelLimitOrderResponse.decode(new Reader(data))
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
