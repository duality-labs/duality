/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { TickObject } from "../dex/tick_map";
import { PairObject } from "../dex/pair_map";
import { Tokens } from "../dex/tokens";
import { TokenObject } from "../dex/token_map";
import { Shares } from "../dex/shares";
import { FeeList } from "../dex/fee_list";
import { LimitOrderPoolUserShareObject } from "../dex/limit_order_pool_user_share_map";
import { LimitOrderPoolUserSharesWithdrawnObject } from "../dex/limit_order_pool_user_shares_withdrawn";
import { LimitOrderPoolTotalSharesObject } from "../dex/limit_order_pool_total_shares_map";
import { LimitOrderPoolReserveObject } from "../dex/limit_order_pool_reserve_map";
import { LimitOrderPoolFillObject } from "../dex/limit_order_pool_fill_map";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** GenesisState defines the dex module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  tickObjectList: TickObject[];
  pairObjectList: PairObject[];
  tokensList: Tokens[];
  tokensCount: number;
  tokenObjectList: TokenObject[];
  sharesList: Shares[];
  feeListList: FeeList[];
  feeListCount: number;
  limitOrderPoolUserShareObjectList: LimitOrderPoolUserShareObject[];
  limitOrderPoolUserSharesWithdrawnObjectList: LimitOrderPoolUserSharesWithdrawnObject[];
  limitOrderPoolTotalSharesObjectList: LimitOrderPoolTotalSharesObject[];
  limitOrderPoolReserveObjectList: LimitOrderPoolReserveObject[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  limitOrderPoolFillObjectList: LimitOrderPoolFillObject[];
}

const baseGenesisState: object = { tokensCount: 0, feeListCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.tickObjectList) {
      TickObject.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.pairObjectList) {
      PairObject.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.tokensList) {
      Tokens.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.tokensCount !== 0) {
      writer.uint32(40).uint64(message.tokensCount);
    }
    for (const v of message.tokenObjectList) {
      TokenObject.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.sharesList) {
      Shares.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.feeListList) {
      FeeList.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.feeListCount !== 0) {
      writer.uint32(72).uint64(message.feeListCount);
    }
    for (const v of message.limitOrderPoolUserShareObjectList) {
      LimitOrderPoolUserShareObject.encode(v!, writer.uint32(114).fork()).ldelim();
    }
    for (const v of message.limitOrderPoolUserSharesWithdrawnObjectList) {
      LimitOrderPoolUserSharesWithdrawnObject.encode(
        v!,
        writer.uint32(122).fork()
      ).ldelim();
    }
    for (const v of message.limitOrderPoolTotalSharesObjectList) {
      LimitOrderPoolTotalSharesObject.encode(
        v!,
        writer.uint32(130).fork()
      ).ldelim();
    }
    for (const v of message.limitOrderPoolReserveObjectList) {
      LimitOrderPoolReserveObject.encode(v!, writer.uint32(138).fork()).ldelim();
    }
    for (const v of message.limitOrderPoolFillObjectList) {
      LimitOrderPoolFillObject.encode(v!, writer.uint32(146).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.tickObjectList = [];
    message.pairObjectList = [];
    message.tokensList = [];
    message.tokenObjectList = [];
    message.sharesList = [];
    message.feeListList = [];
    message.limitOrderPoolUserShareObjectList = [];
    message.limitOrderPoolUserSharesWithdrawnObjectList = [];
    message.limitOrderPoolTotalSharesObjectList = [];
    message.limitOrderPoolReserveObjectList = [];
    message.limitOrderPoolFillObjectList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.tickObjectList.push(TickObject.decode(reader, reader.uint32()));
          break;
        case 3:
          message.pairObjectList.push(PairObject.decode(reader, reader.uint32()));
          break;
        case 4:
          message.tokensList.push(Tokens.decode(reader, reader.uint32()));
          break;
        case 5:
          message.tokensCount = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.tokenObjectList.push(TokenObject.decode(reader, reader.uint32()));
          break;
        case 7:
          message.sharesList.push(Shares.decode(reader, reader.uint32()));
          break;
        case 8:
          message.feeListList.push(FeeList.decode(reader, reader.uint32()));
          break;
        case 9:
          message.feeListCount = longToNumber(reader.uint64() as Long);
          break;
        case 14:
          message.limitOrderPoolUserShareObjectList.push(
            LimitOrderPoolUserShareObject.decode(reader, reader.uint32())
          );
          break;
        case 15:
          message.limitOrderPoolUserSharesWithdrawnObjectList.push(
            LimitOrderPoolUserSharesWithdrawnObject.decode(reader, reader.uint32())
          );
          break;
        case 16:
          message.limitOrderPoolTotalSharesObjectList.push(
            LimitOrderPoolTotalSharesObject.decode(reader, reader.uint32())
          );
          break;
        case 17:
          message.limitOrderPoolReserveObjectList.push(
            LimitOrderPoolReserveObject.decode(reader, reader.uint32())
          );
          break;
        case 18:
          message.limitOrderPoolFillObjectList.push(
            LimitOrderPoolFillObject.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.tickObjectList = [];
    message.pairObjectList = [];
    message.tokensList = [];
    message.tokenObjectList = [];
    message.sharesList = [];
    message.feeListList = [];
    message.limitOrderPoolUserShareObjectList = [];
    message.limitOrderPoolUserSharesWithdrawnObjectList = [];
    message.limitOrderPoolTotalSharesObjectList = [];
    message.limitOrderPoolReserveObjectList = [];
    message.limitOrderPoolFillObjectList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.tickObjectList !== undefined && object.tickObjectList !== null) {
      for (const e of object.tickObjectList) {
        message.tickObjectList.push(TickObject.fromJSON(e));
      }
    }
    if (object.pairObjectList !== undefined && object.pairObjectList !== null) {
      for (const e of object.pairObjectList) {
        message.pairObjectList.push(PairObject.fromJSON(e));
      }
    }
    if (object.tokensList !== undefined && object.tokensList !== null) {
      for (const e of object.tokensList) {
        message.tokensList.push(Tokens.fromJSON(e));
      }
    }
    if (object.tokensCount !== undefined && object.tokensCount !== null) {
      message.tokensCount = Number(object.tokensCount);
    } else {
      message.tokensCount = 0;
    }
    if (object.tokenObjectList !== undefined && object.tokenObjectList !== null) {
      for (const e of object.tokenObjectList) {
        message.tokenObjectList.push(TokenObject.fromJSON(e));
      }
    }
    if (object.sharesList !== undefined && object.sharesList !== null) {
      for (const e of object.sharesList) {
        message.sharesList.push(Shares.fromJSON(e));
      }
    }
    if (object.feeListList !== undefined && object.feeListList !== null) {
      for (const e of object.feeListList) {
        message.feeListList.push(FeeList.fromJSON(e));
      }
    }
    if (object.feeListCount !== undefined && object.feeListCount !== null) {
      message.feeListCount = Number(object.feeListCount);
    } else {
      message.feeListCount = 0;
    }
    if (
      object.limitOrderPoolUserShareObjectList !== undefined &&
      object.limitOrderPoolUserShareObjectList !== null
    ) {
      for (const e of object.limitOrderPoolUserShareObjectList) {
        message.limitOrderPoolUserShareObjectList.push(
          LimitOrderPoolUserShareObject.fromJSON(e)
        );
      }
    }
    if (
      object.limitOrderPoolUserSharesWithdrawnObjectList !== undefined &&
      object.limitOrderPoolUserSharesWithdrawnObjectList !== null
    ) {
      for (const e of object.limitOrderPoolUserSharesWithdrawnObjectList) {
        message.limitOrderPoolUserSharesWithdrawnObjectList.push(
          LimitOrderPoolUserSharesWithdrawnObject.fromJSON(e)
        );
      }
    }
    if (
      object.limitOrderPoolTotalSharesObjectList !== undefined &&
      object.limitOrderPoolTotalSharesObjectList !== null
    ) {
      for (const e of object.limitOrderPoolTotalSharesObjectList) {
        message.limitOrderPoolTotalSharesObjectList.push(
          LimitOrderPoolTotalSharesObject.fromJSON(e)
        );
      }
    }
    if (
      object.limitOrderPoolReserveObjectList !== undefined &&
      object.limitOrderPoolReserveObjectList !== null
    ) {
      for (const e of object.limitOrderPoolReserveObjectList) {
        message.limitOrderPoolReserveObjectList.push(
          LimitOrderPoolReserveObject.fromJSON(e)
        );
      }
    }
    if (
      object.limitOrderPoolFillObjectList !== undefined &&
      object.limitOrderPoolFillObjectList !== null
    ) {
      for (const e of object.limitOrderPoolFillObjectList) {
        message.limitOrderPoolFillObjectList.push(
          LimitOrderPoolFillObject.fromJSON(e)
        );
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.tickObjectList) {
      obj.tickObjectList = message.tickObjectList.map((e) =>
        e ? TickObject.toJSON(e) : undefined
      );
    } else {
      obj.tickObjectList = [];
    }
    if (message.pairObjectList) {
      obj.pairObjectList = message.pairObjectList.map((e) =>
        e ? PairObject.toJSON(e) : undefined
      );
    } else {
      obj.pairObjectList = [];
    }
    if (message.tokensList) {
      obj.tokensList = message.tokensList.map((e) =>
        e ? Tokens.toJSON(e) : undefined
      );
    } else {
      obj.tokensList = [];
    }
    message.tokensCount !== undefined &&
      (obj.tokensCount = message.tokensCount);
    if (message.tokenObjectList) {
      obj.tokenObjectList = message.tokenObjectList.map((e) =>
        e ? TokenObject.toJSON(e) : undefined
      );
    } else {
      obj.tokenObjectList = [];
    }
    if (message.sharesList) {
      obj.sharesList = message.sharesList.map((e) =>
        e ? Shares.toJSON(e) : undefined
      );
    } else {
      obj.sharesList = [];
    }
    if (message.feeListList) {
      obj.feeListList = message.feeListList.map((e) =>
        e ? FeeList.toJSON(e) : undefined
      );
    } else {
      obj.feeListList = [];
    }
    message.feeListCount !== undefined &&
      (obj.feeListCount = message.feeListCount);
    if (message.limitOrderPoolUserShareObjectList) {
      obj.limitOrderPoolUserShareObjectList = message.limitOrderPoolUserShareObjectList.map(
        (e) => (e ? LimitOrderPoolUserShareObject.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolUserShareObjectList = [];
    }
    if (message.limitOrderPoolUserSharesWithdrawnObjectList) {
      obj.limitOrderPoolUserSharesWithdrawnObjectList = message.limitOrderPoolUserSharesWithdrawnObjectList.map(
        (e) => (e ? LimitOrderPoolUserSharesWithdrawnObject.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolUserSharesWithdrawnObjectList = [];
    }
    if (message.limitOrderPoolTotalSharesObjectList) {
      obj.limitOrderPoolTotalSharesObjectList = message.limitOrderPoolTotalSharesObjectList.map(
        (e) => (e ? LimitOrderPoolTotalSharesObject.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolTotalSharesObjectList = [];
    }
    if (message.limitOrderPoolReserveObjectList) {
      obj.limitOrderPoolReserveObjectList = message.limitOrderPoolReserveObjectList.map(
        (e) => (e ? LimitOrderPoolReserveObject.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolReserveObjectList = [];
    }
    if (message.limitOrderPoolFillObjectList) {
      obj.limitOrderPoolFillObjectList = message.limitOrderPoolFillObjectList.map(
        (e) => (e ? LimitOrderPoolFillObject.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolFillObjectList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.tickObjectList = [];
    message.pairObjectList = [];
    message.tokensList = [];
    message.tokenObjectList = [];
    message.sharesList = [];
    message.feeListList = [];
    message.limitOrderPoolUserShareObjectList = [];
    message.limitOrderPoolUserSharesWithdrawnObjectList = [];
    message.limitOrderPoolTotalSharesObjectList = [];
    message.limitOrderPoolReserveObjectList = [];
    message.limitOrderPoolFillObjectList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.tickObjectList !== undefined && object.tickObjectList !== null) {
      for (const e of object.tickObjectList) {
        message.tickObjectList.push(TickObject.fromPartial(e));
      }
    }
    if (object.pairObjectList !== undefined && object.pairObjectList !== null) {
      for (const e of object.pairObjectList) {
        message.pairObjectList.push(PairObject.fromPartial(e));
      }
    }
    if (object.tokensList !== undefined && object.tokensList !== null) {
      for (const e of object.tokensList) {
        message.tokensList.push(Tokens.fromPartial(e));
      }
    }
    if (object.tokensCount !== undefined && object.tokensCount !== null) {
      message.tokensCount = object.tokensCount;
    } else {
      message.tokensCount = 0;
    }
    if (object.tokenObjectList !== undefined && object.tokenObjectList !== null) {
      for (const e of object.tokenObjectList) {
        message.tokenObjectList.push(TokenObject.fromPartial(e));
      }
    }
    if (object.sharesList !== undefined && object.sharesList !== null) {
      for (const e of object.sharesList) {
        message.sharesList.push(Shares.fromPartial(e));
      }
    }
    if (object.feeListList !== undefined && object.feeListList !== null) {
      for (const e of object.feeListList) {
        message.feeListList.push(FeeList.fromPartial(e));
      }
    }
    if (object.feeListCount !== undefined && object.feeListCount !== null) {
      message.feeListCount = object.feeListCount;
    } else {
      message.feeListCount = 0;
    }
    if (
      object.limitOrderPoolUserShareObjectList !== undefined &&
      object.limitOrderPoolUserShareObjectList !== null
    ) {
      for (const e of object.limitOrderPoolUserShareObjectList) {
        message.limitOrderPoolUserShareObjectList.push(
          LimitOrderPoolUserShareObject.fromPartial(e)
        );
      }
    }
    if (
      object.limitOrderPoolUserSharesWithdrawnObjectList !== undefined &&
      object.limitOrderPoolUserSharesWithdrawnObjectList !== null
    ) {
      for (const e of object.limitOrderPoolUserSharesWithdrawnObjectList) {
        message.limitOrderPoolUserSharesWithdrawnObjectList.push(
          LimitOrderPoolUserSharesWithdrawnObject.fromPartial(e)
        );
      }
    }
    if (
      object.limitOrderPoolTotalSharesObjectList !== undefined &&
      object.limitOrderPoolTotalSharesObjectList !== null
    ) {
      for (const e of object.limitOrderPoolTotalSharesObjectList) {
        message.limitOrderPoolTotalSharesObjectList.push(
          LimitOrderPoolTotalSharesObject.fromPartial(e)
        );
      }
    }
    if (
      object.limitOrderPoolReserveObjectList !== undefined &&
      object.limitOrderPoolReserveObjectList !== null
    ) {
      for (const e of object.limitOrderPoolReserveObjectList) {
        message.limitOrderPoolReserveObjectList.push(
          LimitOrderPoolReserveObject.fromPartial(e)
        );
      }
    }
    if (
      object.limitOrderPoolFillObjectList !== undefined &&
      object.limitOrderPoolFillObjectList !== null
    ) {
      for (const e of object.limitOrderPoolFillObjectList) {
        message.limitOrderPoolFillObjectList.push(
          LimitOrderPoolFillObject.fromPartial(e)
        );
      }
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
