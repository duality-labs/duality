/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
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

export interface MsgWithdrawlResponse {
}

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

export interface MsgPlaceLimitOrderResponse {
}

export interface MsgWithdrawFilledLimitOrder {
  creator: string;
  receiver: string;
  tokenA: string;
  tokenB: string;
  tickIndex: number;
  keyToken: string;
  key: number;
}

export interface MsgWithdrawFilledLimitOrderResponse {
}

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

export interface MsgCancelLimitOrderResponse {
}

function createBaseMsgDeposit(): MsgDeposit {
  return {
    creator: "",
    receiver: "",
    tokenA: "",
    tokenB: "",
    amountsA: [],
    amountsB: [],
    tickIndexes: [],
    feeIndexes: [],
  };
}

export const MsgDeposit = {
  encode(message: MsgDeposit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeposit {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeposit();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      tokenA: isSet(object.tokenA) ? String(object.tokenA) : "",
      tokenB: isSet(object.tokenB) ? String(object.tokenB) : "",
      amountsA: Array.isArray(object?.amountsA) ? object.amountsA.map((e: any) => String(e)) : [],
      amountsB: Array.isArray(object?.amountsB) ? object.amountsB.map((e: any) => String(e)) : [],
      tickIndexes: Array.isArray(object?.tickIndexes) ? object.tickIndexes.map((e: any) => Number(e)) : [],
      feeIndexes: Array.isArray(object?.feeIndexes) ? object.feeIndexes.map((e: any) => Number(e)) : [],
    };
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
      obj.tickIndexes = message.tickIndexes.map((e) => Math.round(e));
    } else {
      obj.tickIndexes = [];
    }
    if (message.feeIndexes) {
      obj.feeIndexes = message.feeIndexes.map((e) => Math.round(e));
    } else {
      obj.feeIndexes = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeposit>, I>>(object: I): MsgDeposit {
    const message = createBaseMsgDeposit();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.tokenA = object.tokenA ?? "";
    message.tokenB = object.tokenB ?? "";
    message.amountsA = object.amountsA?.map((e) => e) || [];
    message.amountsB = object.amountsB?.map((e) => e) || [];
    message.tickIndexes = object.tickIndexes?.map((e) => e) || [];
    message.feeIndexes = object.feeIndexes?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgDepositResponse(): MsgDepositResponse {
  return { Reserve0Deposited: [], Reserve1Deposited: [] };
}

export const MsgDepositResponse = {
  encode(message: MsgDepositResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Reserve0Deposited) {
      writer.uint32(10).string(v!);
    }
    for (const v of message.Reserve1Deposited) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDepositResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDepositResponse();
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
    return {
      Reserve0Deposited: Array.isArray(object?.Reserve0Deposited)
        ? object.Reserve0Deposited.map((e: any) => String(e))
        : [],
      Reserve1Deposited: Array.isArray(object?.Reserve1Deposited)
        ? object.Reserve1Deposited.map((e: any) => String(e))
        : [],
    };
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

  fromPartial<I extends Exact<DeepPartial<MsgDepositResponse>, I>>(object: I): MsgDepositResponse {
    const message = createBaseMsgDepositResponse();
    message.Reserve0Deposited = object.Reserve0Deposited?.map((e) => e) || [];
    message.Reserve1Deposited = object.Reserve1Deposited?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgWithdrawl(): MsgWithdrawl {
  return { creator: "", receiver: "", tokenA: "", tokenB: "", sharesToRemove: [], tickIndexes: [], feeIndexes: [] };
}

export const MsgWithdrawl = {
  encode(message: MsgWithdrawl, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawl {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawl();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      tokenA: isSet(object.tokenA) ? String(object.tokenA) : "",
      tokenB: isSet(object.tokenB) ? String(object.tokenB) : "",
      sharesToRemove: Array.isArray(object?.sharesToRemove) ? object.sharesToRemove.map((e: any) => String(e)) : [],
      tickIndexes: Array.isArray(object?.tickIndexes) ? object.tickIndexes.map((e: any) => Number(e)) : [],
      feeIndexes: Array.isArray(object?.feeIndexes) ? object.feeIndexes.map((e: any) => Number(e)) : [],
    };
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
      obj.tickIndexes = message.tickIndexes.map((e) => Math.round(e));
    } else {
      obj.tickIndexes = [];
    }
    if (message.feeIndexes) {
      obj.feeIndexes = message.feeIndexes.map((e) => Math.round(e));
    } else {
      obj.feeIndexes = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawl>, I>>(object: I): MsgWithdrawl {
    const message = createBaseMsgWithdrawl();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.tokenA = object.tokenA ?? "";
    message.tokenB = object.tokenB ?? "";
    message.sharesToRemove = object.sharesToRemove?.map((e) => e) || [];
    message.tickIndexes = object.tickIndexes?.map((e) => e) || [];
    message.feeIndexes = object.feeIndexes?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgWithdrawlResponse(): MsgWithdrawlResponse {
  return {};
}

export const MsgWithdrawlResponse = {
  encode(_: MsgWithdrawlResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawlResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawlResponse();
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
    return {};
  },

  toJSON(_: MsgWithdrawlResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawlResponse>, I>>(_: I): MsgWithdrawlResponse {
    const message = createBaseMsgWithdrawlResponse();
    return message;
  },
};

function createBaseMsgSwap(): MsgSwap {
  return { creator: "", receiver: "", tokenA: "", tokenB: "", amountIn: "", tokenIn: "", minOut: "" };
}

export const MsgSwap = {
  encode(message: MsgSwap, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSwap {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSwap();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      tokenA: isSet(object.tokenA) ? String(object.tokenA) : "",
      tokenB: isSet(object.tokenB) ? String(object.tokenB) : "",
      amountIn: isSet(object.amountIn) ? String(object.amountIn) : "",
      tokenIn: isSet(object.tokenIn) ? String(object.tokenIn) : "",
      minOut: isSet(object.minOut) ? String(object.minOut) : "",
    };
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

  fromPartial<I extends Exact<DeepPartial<MsgSwap>, I>>(object: I): MsgSwap {
    const message = createBaseMsgSwap();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.tokenA = object.tokenA ?? "";
    message.tokenB = object.tokenB ?? "";
    message.amountIn = object.amountIn ?? "";
    message.tokenIn = object.tokenIn ?? "";
    message.minOut = object.minOut ?? "";
    return message;
  },
};

function createBaseMsgSwapResponse(): MsgSwapResponse {
  return { coinOut: undefined };
}

export const MsgSwapResponse = {
  encode(message: MsgSwapResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.coinOut !== undefined) {
      Coin.encode(message.coinOut, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSwapResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSwapResponse();
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
    return { coinOut: isSet(object.coinOut) ? Coin.fromJSON(object.coinOut) : undefined };
  },

  toJSON(message: MsgSwapResponse): unknown {
    const obj: any = {};
    message.coinOut !== undefined && (obj.coinOut = message.coinOut ? Coin.toJSON(message.coinOut) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSwapResponse>, I>>(object: I): MsgSwapResponse {
    const message = createBaseMsgSwapResponse();
    message.coinOut = (object.coinOut !== undefined && object.coinOut !== null)
      ? Coin.fromPartial(object.coinOut)
      : undefined;
    return message;
  },
};

function createBaseMsgPlaceLimitOrder(): MsgPlaceLimitOrder {
  return { creator: "", receiver: "", tokenA: "", tokenB: "", tickIndex: 0, tokenIn: "", amountIn: "" };
}

export const MsgPlaceLimitOrder = {
  encode(message: MsgPlaceLimitOrder, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlaceLimitOrder {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlaceLimitOrder();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      tokenA: isSet(object.tokenA) ? String(object.tokenA) : "",
      tokenB: isSet(object.tokenB) ? String(object.tokenB) : "",
      tickIndex: isSet(object.tickIndex) ? Number(object.tickIndex) : 0,
      tokenIn: isSet(object.tokenIn) ? String(object.tokenIn) : "",
      amountIn: isSet(object.amountIn) ? String(object.amountIn) : "",
    };
  },

  toJSON(message: MsgPlaceLimitOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.tickIndex !== undefined && (obj.tickIndex = Math.round(message.tickIndex));
    message.tokenIn !== undefined && (obj.tokenIn = message.tokenIn);
    message.amountIn !== undefined && (obj.amountIn = message.amountIn);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPlaceLimitOrder>, I>>(object: I): MsgPlaceLimitOrder {
    const message = createBaseMsgPlaceLimitOrder();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.tokenA = object.tokenA ?? "";
    message.tokenB = object.tokenB ?? "";
    message.tickIndex = object.tickIndex ?? 0;
    message.tokenIn = object.tokenIn ?? "";
    message.amountIn = object.amountIn ?? "";
    return message;
  },
};

function createBaseMsgPlaceLimitOrderResponse(): MsgPlaceLimitOrderResponse {
  return {};
}

export const MsgPlaceLimitOrderResponse = {
  encode(_: MsgPlaceLimitOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlaceLimitOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlaceLimitOrderResponse();
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
    return {};
  },

  toJSON(_: MsgPlaceLimitOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPlaceLimitOrderResponse>, I>>(_: I): MsgPlaceLimitOrderResponse {
    const message = createBaseMsgPlaceLimitOrderResponse();
    return message;
  },
};

function createBaseMsgWithdrawFilledLimitOrder(): MsgWithdrawFilledLimitOrder {
  return { creator: "", receiver: "", tokenA: "", tokenB: "", tickIndex: 0, keyToken: "", key: 0 };
}

export const MsgWithdrawFilledLimitOrder = {
  encode(message: MsgWithdrawFilledLimitOrder, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawFilledLimitOrder {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawFilledLimitOrder();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      tokenA: isSet(object.tokenA) ? String(object.tokenA) : "",
      tokenB: isSet(object.tokenB) ? String(object.tokenB) : "",
      tickIndex: isSet(object.tickIndex) ? Number(object.tickIndex) : 0,
      keyToken: isSet(object.keyToken) ? String(object.keyToken) : "",
      key: isSet(object.key) ? Number(object.key) : 0,
    };
  },

  toJSON(message: MsgWithdrawFilledLimitOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.tickIndex !== undefined && (obj.tickIndex = Math.round(message.tickIndex));
    message.keyToken !== undefined && (obj.keyToken = message.keyToken);
    message.key !== undefined && (obj.key = Math.round(message.key));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawFilledLimitOrder>, I>>(object: I): MsgWithdrawFilledLimitOrder {
    const message = createBaseMsgWithdrawFilledLimitOrder();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.tokenA = object.tokenA ?? "";
    message.tokenB = object.tokenB ?? "";
    message.tickIndex = object.tickIndex ?? 0;
    message.keyToken = object.keyToken ?? "";
    message.key = object.key ?? 0;
    return message;
  },
};

function createBaseMsgWithdrawFilledLimitOrderResponse(): MsgWithdrawFilledLimitOrderResponse {
  return {};
}

export const MsgWithdrawFilledLimitOrderResponse = {
  encode(_: MsgWithdrawFilledLimitOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawFilledLimitOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawFilledLimitOrderResponse();
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
    return {};
  },

  toJSON(_: MsgWithdrawFilledLimitOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawFilledLimitOrderResponse>, I>>(
    _: I,
  ): MsgWithdrawFilledLimitOrderResponse {
    const message = createBaseMsgWithdrawFilledLimitOrderResponse();
    return message;
  },
};

function createBaseMsgCancelLimitOrder(): MsgCancelLimitOrder {
  return { creator: "", receiver: "", tokenA: "", tokenB: "", tickIndex: 0, keyToken: "", key: 0, sharesOut: "" };
}

export const MsgCancelLimitOrder = {
  encode(message: MsgCancelLimitOrder, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelLimitOrder {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelLimitOrder();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      tokenA: isSet(object.tokenA) ? String(object.tokenA) : "",
      tokenB: isSet(object.tokenB) ? String(object.tokenB) : "",
      tickIndex: isSet(object.tickIndex) ? Number(object.tickIndex) : 0,
      keyToken: isSet(object.keyToken) ? String(object.keyToken) : "",
      key: isSet(object.key) ? Number(object.key) : 0,
      sharesOut: isSet(object.sharesOut) ? String(object.sharesOut) : "",
    };
  },

  toJSON(message: MsgCancelLimitOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.tickIndex !== undefined && (obj.tickIndex = Math.round(message.tickIndex));
    message.keyToken !== undefined && (obj.keyToken = message.keyToken);
    message.key !== undefined && (obj.key = Math.round(message.key));
    message.sharesOut !== undefined && (obj.sharesOut = message.sharesOut);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelLimitOrder>, I>>(object: I): MsgCancelLimitOrder {
    const message = createBaseMsgCancelLimitOrder();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.tokenA = object.tokenA ?? "";
    message.tokenB = object.tokenB ?? "";
    message.tickIndex = object.tickIndex ?? 0;
    message.keyToken = object.keyToken ?? "";
    message.key = object.key ?? 0;
    message.sharesOut = object.sharesOut ?? "";
    return message;
  },
};

function createBaseMsgCancelLimitOrderResponse(): MsgCancelLimitOrderResponse {
  return {};
}

export const MsgCancelLimitOrderResponse = {
  encode(_: MsgCancelLimitOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelLimitOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelLimitOrderResponse();
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
    return {};
  },

  toJSON(_: MsgCancelLimitOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelLimitOrderResponse>, I>>(_: I): MsgCancelLimitOrderResponse {
    const message = createBaseMsgCancelLimitOrderResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse>;
  Withdrawl(request: MsgWithdrawl): Promise<MsgWithdrawlResponse>;
  Swap(request: MsgSwap): Promise<MsgSwapResponse>;
  PlaceLimitOrder(request: MsgPlaceLimitOrder): Promise<MsgPlaceLimitOrderResponse>;
  WithdrawFilledLimitOrder(request: MsgWithdrawFilledLimitOrder): Promise<MsgWithdrawFilledLimitOrderResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CancelLimitOrder(request: MsgCancelLimitOrder): Promise<MsgCancelLimitOrderResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Deposit = this.Deposit.bind(this);
    this.Withdrawl = this.Withdrawl.bind(this);
    this.Swap = this.Swap.bind(this);
    this.PlaceLimitOrder = this.PlaceLimitOrder.bind(this);
    this.WithdrawFilledLimitOrder = this.WithdrawFilledLimitOrder.bind(this);
    this.CancelLimitOrder = this.CancelLimitOrder.bind(this);
  }
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse> {
    const data = MsgDeposit.encode(request).finish();
    const promise = this.rpc.request("nicholasdotsol.duality.dex.Msg", "Deposit", data);
    return promise.then((data) => MsgDepositResponse.decode(new _m0.Reader(data)));
  }

  Withdrawl(request: MsgWithdrawl): Promise<MsgWithdrawlResponse> {
    const data = MsgWithdrawl.encode(request).finish();
    const promise = this.rpc.request("nicholasdotsol.duality.dex.Msg", "Withdrawl", data);
    return promise.then((data) => MsgWithdrawlResponse.decode(new _m0.Reader(data)));
  }

  Swap(request: MsgSwap): Promise<MsgSwapResponse> {
    const data = MsgSwap.encode(request).finish();
    const promise = this.rpc.request("nicholasdotsol.duality.dex.Msg", "Swap", data);
    return promise.then((data) => MsgSwapResponse.decode(new _m0.Reader(data)));
  }

  PlaceLimitOrder(request: MsgPlaceLimitOrder): Promise<MsgPlaceLimitOrderResponse> {
    const data = MsgPlaceLimitOrder.encode(request).finish();
    const promise = this.rpc.request("nicholasdotsol.duality.dex.Msg", "PlaceLimitOrder", data);
    return promise.then((data) => MsgPlaceLimitOrderResponse.decode(new _m0.Reader(data)));
  }

  WithdrawFilledLimitOrder(request: MsgWithdrawFilledLimitOrder): Promise<MsgWithdrawFilledLimitOrderResponse> {
    const data = MsgWithdrawFilledLimitOrder.encode(request).finish();
    const promise = this.rpc.request("nicholasdotsol.duality.dex.Msg", "WithdrawFilledLimitOrder", data);
    return promise.then((data) => MsgWithdrawFilledLimitOrderResponse.decode(new _m0.Reader(data)));
  }

  CancelLimitOrder(request: MsgCancelLimitOrder): Promise<MsgCancelLimitOrderResponse> {
    const data = MsgCancelLimitOrder.encode(request).finish();
    const promise = this.rpc.request("nicholasdotsol.duality.dex.Msg", "CancelLimitOrder", data);
    return promise.then((data) => MsgCancelLimitOrderResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
