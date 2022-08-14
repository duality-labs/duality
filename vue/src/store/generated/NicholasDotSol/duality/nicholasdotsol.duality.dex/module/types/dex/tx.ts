/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface MsgAddLiquidity {
  creator: string;
  tokenA: string;
  tokenB: string;
  tokenDirection: string;
  amount: string;
  price: string;
  fee: string;
  orderType: string;
}

export interface MsgAddLiquidityResponse {}

export interface MsgRemoveLiquidity {
  creator: string;
  tokenA: string;
  tokenB: string;
  tokenDirection: string;
  shares: string;
  price: string;
  fee: string;
  orderType: string;
}

export interface MsgRemoveLiquidityResponse {}

const baseMsgAddLiquidity: object = {
  creator: "",
  tokenA: "",
  tokenB: "",
  tokenDirection: "",
  amount: "",
  price: "",
  fee: "",
  orderType: "",
};

export const MsgAddLiquidity = {
  encode(message: MsgAddLiquidity, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenA !== "") {
      writer.uint32(18).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(26).string(message.tokenB);
    }
    if (message.tokenDirection !== "") {
      writer.uint32(34).string(message.tokenDirection);
    }
    if (message.amount !== "") {
      writer.uint32(42).string(message.amount);
    }
    if (message.price !== "") {
      writer.uint32(50).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(58).string(message.fee);
    }
    if (message.orderType !== "") {
      writer.uint32(66).string(message.orderType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddLiquidity {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddLiquidity } as MsgAddLiquidity;
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
          message.tokenDirection = reader.string();
          break;
        case 5:
          message.amount = reader.string();
          break;
        case 6:
          message.price = reader.string();
          break;
        case 7:
          message.fee = reader.string();
          break;
        case 8:
          message.orderType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddLiquidity {
    const message = { ...baseMsgAddLiquidity } as MsgAddLiquidity;
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
    if (object.tokenDirection !== undefined && object.tokenDirection !== null) {
      message.tokenDirection = String(object.tokenDirection);
    } else {
      message.tokenDirection = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    return message;
  },

  toJSON(message: MsgAddLiquidity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.tokenDirection !== undefined &&
      (obj.tokenDirection = message.tokenDirection);
    message.amount !== undefined && (obj.amount = message.amount);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddLiquidity>): MsgAddLiquidity {
    const message = { ...baseMsgAddLiquidity } as MsgAddLiquidity;
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
    if (object.tokenDirection !== undefined && object.tokenDirection !== null) {
      message.tokenDirection = object.tokenDirection;
    } else {
      message.tokenDirection = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    return message;
  },
};

const baseMsgAddLiquidityResponse: object = {};

export const MsgAddLiquidityResponse = {
  encode(_: MsgAddLiquidityResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddLiquidityResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddLiquidityResponse,
    } as MsgAddLiquidityResponse;
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

  fromJSON(_: any): MsgAddLiquidityResponse {
    const message = {
      ...baseMsgAddLiquidityResponse,
    } as MsgAddLiquidityResponse;
    return message;
  },

  toJSON(_: MsgAddLiquidityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddLiquidityResponse>
  ): MsgAddLiquidityResponse {
    const message = {
      ...baseMsgAddLiquidityResponse,
    } as MsgAddLiquidityResponse;
    return message;
  },
};

const baseMsgRemoveLiquidity: object = {
  creator: "",
  tokenA: "",
  tokenB: "",
  tokenDirection: "",
  shares: "",
  price: "",
  fee: "",
  orderType: "",
};

export const MsgRemoveLiquidity = {
  encode(
    message: MsgRemoveLiquidity,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenA !== "") {
      writer.uint32(18).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(26).string(message.tokenB);
    }
    if (message.tokenDirection !== "") {
      writer.uint32(34).string(message.tokenDirection);
    }
    if (message.shares !== "") {
      writer.uint32(42).string(message.shares);
    }
    if (message.price !== "") {
      writer.uint32(50).string(message.price);
    }
    if (message.fee !== "") {
      writer.uint32(58).string(message.fee);
    }
    if (message.orderType !== "") {
      writer.uint32(66).string(message.orderType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRemoveLiquidity {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRemoveLiquidity } as MsgRemoveLiquidity;
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
          message.tokenDirection = reader.string();
          break;
        case 5:
          message.shares = reader.string();
          break;
        case 6:
          message.price = reader.string();
          break;
        case 7:
          message.fee = reader.string();
          break;
        case 8:
          message.orderType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveLiquidity {
    const message = { ...baseMsgRemoveLiquidity } as MsgRemoveLiquidity;
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
    if (object.tokenDirection !== undefined && object.tokenDirection !== null) {
      message.tokenDirection = String(object.tokenDirection);
    } else {
      message.tokenDirection = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = String(object.shares);
    } else {
      message.shares = "";
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    return message;
  },

  toJSON(message: MsgRemoveLiquidity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.tokenDirection !== undefined &&
      (obj.tokenDirection = message.tokenDirection);
    message.shares !== undefined && (obj.shares = message.shares);
    message.price !== undefined && (obj.price = message.price);
    message.fee !== undefined && (obj.fee = message.fee);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRemoveLiquidity>): MsgRemoveLiquidity {
    const message = { ...baseMsgRemoveLiquidity } as MsgRemoveLiquidity;
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
    if (object.tokenDirection !== undefined && object.tokenDirection !== null) {
      message.tokenDirection = object.tokenDirection;
    } else {
      message.tokenDirection = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = object.shares;
    } else {
      message.shares = "";
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    return message;
  },
};

const baseMsgRemoveLiquidityResponse: object = {};

export const MsgRemoveLiquidityResponse = {
  encode(
    _: MsgRemoveLiquidityResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveLiquidityResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveLiquidityResponse,
    } as MsgRemoveLiquidityResponse;
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

  fromJSON(_: any): MsgRemoveLiquidityResponse {
    const message = {
      ...baseMsgRemoveLiquidityResponse,
    } as MsgRemoveLiquidityResponse;
    return message;
  },

  toJSON(_: MsgRemoveLiquidityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRemoveLiquidityResponse>
  ): MsgRemoveLiquidityResponse {
    const message = {
      ...baseMsgRemoveLiquidityResponse,
    } as MsgRemoveLiquidityResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddLiquidity(request: MsgAddLiquidity): Promise<MsgAddLiquidityResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RemoveLiquidity(
    request: MsgRemoveLiquidity
  ): Promise<MsgRemoveLiquidityResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AddLiquidity(request: MsgAddLiquidity): Promise<MsgAddLiquidityResponse> {
    const data = MsgAddLiquidity.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "AddLiquidity",
      data
    );
    return promise.then((data) =>
      MsgAddLiquidityResponse.decode(new Reader(data))
    );
  }

  RemoveLiquidity(
    request: MsgRemoveLiquidity
  ): Promise<MsgRemoveLiquidityResponse> {
    const data = MsgRemoveLiquidity.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "RemoveLiquidity",
      data
    );
    return promise.then((data) =>
      MsgRemoveLiquidityResponse.decode(new Reader(data))
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
