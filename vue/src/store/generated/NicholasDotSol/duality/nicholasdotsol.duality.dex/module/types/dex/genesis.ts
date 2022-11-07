/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { TickObject } from "../dex/tick_map";
import { PairMap } from "../dex/pair_map";
import { Tokens } from "../dex/tokens";
import { TokenObject } from "../dex/token_map";
import { Shares } from "../dex/shares";
import { FeeList } from "../dex/fee_list";
import { LimitOrderPoolUserShareMap } from "../dex/limit_order_pool_user_share_map";
import { LimitOrderPoolUserSharesWithdrawn } from "../dex/limit_order_pool_user_shares_withdrawn";
import { LimitOrderPoolTotalSharesMap } from "../dex/limit_order_pool_total_shares_map";
import { LimitOrderPoolReserveMap } from "../dex/limit_order_pool_reserve_map";
import { LimitOrderPoolFillMap } from "../dex/limit_order_pool_fill_map";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** GenesisState defines the dex module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  tickObjectList: TickObject[];
  pairMapList: PairMap[];
  tokensList: Tokens[];
  tokensCount: number;
  tokenObjectList: TokenObject[];
  sharesList: Shares[];
  feeListList: FeeList[];
  feeListCount: number;
  limitOrderPoolUserShareMapList: LimitOrderPoolUserShareMap[];
  limitOrderPoolUserSharesWithdrawnList: LimitOrderPoolUserSharesWithdrawn[];
  limitOrderPoolTotalSharesMapList: LimitOrderPoolTotalSharesMap[];
  limitOrderPoolReserveMapList: LimitOrderPoolReserveMap[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  limitOrderPoolFillMapList: LimitOrderPoolFillMap[];
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
    for (const v of message.pairMapList) {
      PairMap.encode(v!, writer.uint32(26).fork()).ldelim();
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
    for (const v of message.limitOrderPoolUserShareMapList) {
      LimitOrderPoolUserShareMap.encode(v!, writer.uint32(114).fork()).ldelim();
    }
    for (const v of message.limitOrderPoolUserSharesWithdrawnList) {
      LimitOrderPoolUserSharesWithdrawn.encode(
        v!,
        writer.uint32(122).fork()
      ).ldelim();
    }
    for (const v of message.limitOrderPoolTotalSharesMapList) {
      LimitOrderPoolTotalSharesMap.encode(
        v!,
        writer.uint32(130).fork()
      ).ldelim();
    }
    for (const v of message.limitOrderPoolReserveMapList) {
      LimitOrderPoolReserveMap.encode(v!, writer.uint32(138).fork()).ldelim();
    }
    for (const v of message.limitOrderPoolFillMapList) {
      LimitOrderPoolFillMap.encode(v!, writer.uint32(146).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.tickObjectList = [];
    message.pairMapList = [];
    message.tokensList = [];
    message.tokenObjectList = [];
    message.sharesList = [];
    message.feeListList = [];
    message.limitOrderPoolUserShareMapList = [];
    message.limitOrderPoolUserSharesWithdrawnList = [];
    message.limitOrderPoolTotalSharesMapList = [];
    message.limitOrderPoolReserveMapList = [];
    message.limitOrderPoolFillMapList = [];
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
          message.pairMapList.push(PairMap.decode(reader, reader.uint32()));
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
          message.limitOrderPoolUserShareMapList.push(
            LimitOrderPoolUserShareMap.decode(reader, reader.uint32())
          );
          break;
        case 15:
          message.limitOrderPoolUserSharesWithdrawnList.push(
            LimitOrderPoolUserSharesWithdrawn.decode(reader, reader.uint32())
          );
          break;
        case 16:
          message.limitOrderPoolTotalSharesMapList.push(
            LimitOrderPoolTotalSharesMap.decode(reader, reader.uint32())
          );
          break;
        case 17:
          message.limitOrderPoolReserveMapList.push(
            LimitOrderPoolReserveMap.decode(reader, reader.uint32())
          );
          break;
        case 18:
          message.limitOrderPoolFillMapList.push(
            LimitOrderPoolFillMap.decode(reader, reader.uint32())
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
    message.pairMapList = [];
    message.tokensList = [];
    message.tokenObjectList = [];
    message.sharesList = [];
    message.feeListList = [];
    message.limitOrderPoolUserShareMapList = [];
    message.limitOrderPoolUserSharesWithdrawnList = [];
    message.limitOrderPoolTotalSharesMapList = [];
    message.limitOrderPoolReserveMapList = [];
    message.limitOrderPoolFillMapList = [];
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
    if (object.pairMapList !== undefined && object.pairMapList !== null) {
      for (const e of object.pairMapList) {
        message.pairMapList.push(PairMap.fromJSON(e));
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
      object.limitOrderPoolUserShareMapList !== undefined &&
      object.limitOrderPoolUserShareMapList !== null
    ) {
      for (const e of object.limitOrderPoolUserShareMapList) {
        message.limitOrderPoolUserShareMapList.push(
          LimitOrderPoolUserShareMap.fromJSON(e)
        );
      }
    }
    if (
      object.limitOrderPoolUserSharesWithdrawnList !== undefined &&
      object.limitOrderPoolUserSharesWithdrawnList !== null
    ) {
      for (const e of object.limitOrderPoolUserSharesWithdrawnList) {
        message.limitOrderPoolUserSharesWithdrawnList.push(
          LimitOrderPoolUserSharesWithdrawn.fromJSON(e)
        );
      }
    }
    if (
      object.limitOrderPoolTotalSharesMapList !== undefined &&
      object.limitOrderPoolTotalSharesMapList !== null
    ) {
      for (const e of object.limitOrderPoolTotalSharesMapList) {
        message.limitOrderPoolTotalSharesMapList.push(
          LimitOrderPoolTotalSharesMap.fromJSON(e)
        );
      }
    }
    if (
      object.limitOrderPoolReserveMapList !== undefined &&
      object.limitOrderPoolReserveMapList !== null
    ) {
      for (const e of object.limitOrderPoolReserveMapList) {
        message.limitOrderPoolReserveMapList.push(
          LimitOrderPoolReserveMap.fromJSON(e)
        );
      }
    }
    if (
      object.limitOrderPoolFillMapList !== undefined &&
      object.limitOrderPoolFillMapList !== null
    ) {
      for (const e of object.limitOrderPoolFillMapList) {
        message.limitOrderPoolFillMapList.push(
          LimitOrderPoolFillMap.fromJSON(e)
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
    if (message.pairMapList) {
      obj.pairMapList = message.pairMapList.map((e) =>
        e ? PairMap.toJSON(e) : undefined
      );
    } else {
      obj.pairMapList = [];
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
    if (message.limitOrderPoolUserShareMapList) {
      obj.limitOrderPoolUserShareMapList = message.limitOrderPoolUserShareMapList.map(
        (e) => (e ? LimitOrderPoolUserShareMap.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolUserShareMapList = [];
    }
    if (message.limitOrderPoolUserSharesWithdrawnList) {
      obj.limitOrderPoolUserSharesWithdrawnList = message.limitOrderPoolUserSharesWithdrawnList.map(
        (e) => (e ? LimitOrderPoolUserSharesWithdrawn.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolUserSharesWithdrawnList = [];
    }
    if (message.limitOrderPoolTotalSharesMapList) {
      obj.limitOrderPoolTotalSharesMapList = message.limitOrderPoolTotalSharesMapList.map(
        (e) => (e ? LimitOrderPoolTotalSharesMap.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolTotalSharesMapList = [];
    }
    if (message.limitOrderPoolReserveMapList) {
      obj.limitOrderPoolReserveMapList = message.limitOrderPoolReserveMapList.map(
        (e) => (e ? LimitOrderPoolReserveMap.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolReserveMapList = [];
    }
    if (message.limitOrderPoolFillMapList) {
      obj.limitOrderPoolFillMapList = message.limitOrderPoolFillMapList.map(
        (e) => (e ? LimitOrderPoolFillMap.toJSON(e) : undefined)
      );
    } else {
      obj.limitOrderPoolFillMapList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.tickObjectList = [];
    message.pairMapList = [];
    message.tokensList = [];
    message.tokenObjectList = [];
    message.sharesList = [];
    message.feeListList = [];
    message.limitOrderPoolUserShareMapList = [];
    message.limitOrderPoolUserSharesWithdrawnList = [];
    message.limitOrderPoolTotalSharesMapList = [];
    message.limitOrderPoolReserveMapList = [];
    message.limitOrderPoolFillMapList = [];
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
    if (object.pairMapList !== undefined && object.pairMapList !== null) {
      for (const e of object.pairMapList) {
        message.pairMapList.push(PairMap.fromPartial(e));
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
      object.limitOrderPoolUserShareMapList !== undefined &&
      object.limitOrderPoolUserShareMapList !== null
    ) {
      for (const e of object.limitOrderPoolUserShareMapList) {
        message.limitOrderPoolUserShareMapList.push(
          LimitOrderPoolUserShareMap.fromPartial(e)
        );
      }
    }
    if (
      object.limitOrderPoolUserSharesWithdrawnList !== undefined &&
      object.limitOrderPoolUserSharesWithdrawnList !== null
    ) {
      for (const e of object.limitOrderPoolUserSharesWithdrawnList) {
        message.limitOrderPoolUserSharesWithdrawnList.push(
          LimitOrderPoolUserSharesWithdrawn.fromPartial(e)
        );
      }
    }
    if (
      object.limitOrderPoolTotalSharesMapList !== undefined &&
      object.limitOrderPoolTotalSharesMapList !== null
    ) {
      for (const e of object.limitOrderPoolTotalSharesMapList) {
        message.limitOrderPoolTotalSharesMapList.push(
          LimitOrderPoolTotalSharesMap.fromPartial(e)
        );
      }
    }
    if (
      object.limitOrderPoolReserveMapList !== undefined &&
      object.limitOrderPoolReserveMapList !== null
    ) {
      for (const e of object.limitOrderPoolReserveMapList) {
        message.limitOrderPoolReserveMapList.push(
          LimitOrderPoolReserveMap.fromPartial(e)
        );
      }
    }
    if (
      object.limitOrderPoolFillMapList !== undefined &&
      object.limitOrderPoolFillMapList !== null
    ) {
      for (const e of object.limitOrderPoolFillMapList) {
        message.limitOrderPoolFillMapList.push(
          LimitOrderPoolFillMap.fromPartial(e)
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
