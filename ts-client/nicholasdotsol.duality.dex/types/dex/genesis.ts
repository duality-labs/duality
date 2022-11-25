/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { AdjanceyMatrix } from "./adjancey_matrix";
import { EdgeRow } from "./edge_row";
import { FeeList } from "./fee_list";
import { LimitOrderTranche } from "./limit_order_tranche";
import { LimitOrderTrancheUser } from "./limit_order_tranche_user";
import { PairMap } from "./pair_map";
import { Params } from "./params";
import { Shares } from "./shares";
import { TickMap } from "./tick_map";
import { TokenMap } from "./token_map";
import { Tokens } from "./tokens";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** GenesisState defines the dex module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  tickMapList: TickMap[];
  pairMapList: PairMap[];
  tokensList: Tokens[];
  tokensCount: number;
  tokenMapList: TokenMap[];
  sharesList: Shares[];
  feeListList: FeeList[];
  feeListCount: number;
  edgeRowList: EdgeRow[];
  edgeRowCount: number;
  adjanceyMatrixList: AdjanceyMatrix[];
  adjanceyMatrixCount: number;
  LimitOrderTrancheUserList: LimitOrderTrancheUser[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  LimitOrderTrancheList: LimitOrderTranche[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    tickMapList: [],
    pairMapList: [],
    tokensList: [],
    tokensCount: 0,
    tokenMapList: [],
    sharesList: [],
    feeListList: [],
    feeListCount: 0,
    edgeRowList: [],
    edgeRowCount: 0,
    adjanceyMatrixList: [],
    adjanceyMatrixCount: 0,
    LimitOrderTrancheUserList: [],
    LimitOrderTrancheList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.tickMapList) {
      TickMap.encode(v!, writer.uint32(18).fork()).ldelim();
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
    for (const v of message.tokenMapList) {
      TokenMap.encode(v!, writer.uint32(50).fork()).ldelim();
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
    for (const v of message.edgeRowList) {
      EdgeRow.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    if (message.edgeRowCount !== 0) {
      writer.uint32(88).uint64(message.edgeRowCount);
    }
    for (const v of message.adjanceyMatrixList) {
      AdjanceyMatrix.encode(v!, writer.uint32(98).fork()).ldelim();
    }
    if (message.adjanceyMatrixCount !== 0) {
      writer.uint32(104).uint64(message.adjanceyMatrixCount);
    }
    for (const v of message.LimitOrderTrancheUserList) {
      LimitOrderTrancheUser.encode(v!, writer.uint32(114).fork()).ldelim();
    }
    for (const v of message.LimitOrderTrancheList) {
      LimitOrderTranche.encode(v!, writer.uint32(122).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.tickMapList.push(TickMap.decode(reader, reader.uint32()));
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
          message.tokenMapList.push(TokenMap.decode(reader, reader.uint32()));
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
        case 10:
          message.edgeRowList.push(EdgeRow.decode(reader, reader.uint32()));
          break;
        case 11:
          message.edgeRowCount = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.adjanceyMatrixList.push(AdjanceyMatrix.decode(reader, reader.uint32()));
          break;
        case 13:
          message.adjanceyMatrixCount = longToNumber(reader.uint64() as Long);
          break;
        case 14:
          message.LimitOrderTrancheUserList.push(LimitOrderTrancheUser.decode(reader, reader.uint32()));
          break;
        case 15:
          message.LimitOrderTrancheList.push(LimitOrderTranche.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      tickMapList: Array.isArray(object?.tickMapList) ? object.tickMapList.map((e: any) => TickMap.fromJSON(e)) : [],
      pairMapList: Array.isArray(object?.pairMapList) ? object.pairMapList.map((e: any) => PairMap.fromJSON(e)) : [],
      tokensList: Array.isArray(object?.tokensList) ? object.tokensList.map((e: any) => Tokens.fromJSON(e)) : [],
      tokensCount: isSet(object.tokensCount) ? Number(object.tokensCount) : 0,
      tokenMapList: Array.isArray(object?.tokenMapList)
        ? object.tokenMapList.map((e: any) => TokenMap.fromJSON(e))
        : [],
      sharesList: Array.isArray(object?.sharesList) ? object.sharesList.map((e: any) => Shares.fromJSON(e)) : [],
      feeListList: Array.isArray(object?.feeListList) ? object.feeListList.map((e: any) => FeeList.fromJSON(e)) : [],
      feeListCount: isSet(object.feeListCount) ? Number(object.feeListCount) : 0,
      edgeRowList: Array.isArray(object?.edgeRowList) ? object.edgeRowList.map((e: any) => EdgeRow.fromJSON(e)) : [],
      edgeRowCount: isSet(object.edgeRowCount) ? Number(object.edgeRowCount) : 0,
      adjanceyMatrixList: Array.isArray(object?.adjanceyMatrixList)
        ? object.adjanceyMatrixList.map((e: any) => AdjanceyMatrix.fromJSON(e))
        : [],
      adjanceyMatrixCount: isSet(object.adjanceyMatrixCount) ? Number(object.adjanceyMatrixCount) : 0,
      LimitOrderTrancheUserList: Array.isArray(object?.LimitOrderTrancheUserList)
        ? object.LimitOrderTrancheUserList.map((e: any) => LimitOrderTrancheUser.fromJSON(e))
        : [],
      LimitOrderTrancheList: Array.isArray(object?.LimitOrderTrancheList)
        ? object.LimitOrderTrancheList.map((e: any) => LimitOrderTranche.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.tickMapList) {
      obj.tickMapList = message.tickMapList.map((e) => e ? TickMap.toJSON(e) : undefined);
    } else {
      obj.tickMapList = [];
    }
    if (message.pairMapList) {
      obj.pairMapList = message.pairMapList.map((e) => e ? PairMap.toJSON(e) : undefined);
    } else {
      obj.pairMapList = [];
    }
    if (message.tokensList) {
      obj.tokensList = message.tokensList.map((e) => e ? Tokens.toJSON(e) : undefined);
    } else {
      obj.tokensList = [];
    }
    message.tokensCount !== undefined && (obj.tokensCount = Math.round(message.tokensCount));
    if (message.tokenMapList) {
      obj.tokenMapList = message.tokenMapList.map((e) => e ? TokenMap.toJSON(e) : undefined);
    } else {
      obj.tokenMapList = [];
    }
    if (message.sharesList) {
      obj.sharesList = message.sharesList.map((e) => e ? Shares.toJSON(e) : undefined);
    } else {
      obj.sharesList = [];
    }
    if (message.feeListList) {
      obj.feeListList = message.feeListList.map((e) => e ? FeeList.toJSON(e) : undefined);
    } else {
      obj.feeListList = [];
    }
    message.feeListCount !== undefined && (obj.feeListCount = Math.round(message.feeListCount));
    if (message.edgeRowList) {
      obj.edgeRowList = message.edgeRowList.map((e) => e ? EdgeRow.toJSON(e) : undefined);
    } else {
      obj.edgeRowList = [];
    }
    message.edgeRowCount !== undefined && (obj.edgeRowCount = Math.round(message.edgeRowCount));
    if (message.adjanceyMatrixList) {
      obj.adjanceyMatrixList = message.adjanceyMatrixList.map((e) => e ? AdjanceyMatrix.toJSON(e) : undefined);
    } else {
      obj.adjanceyMatrixList = [];
    }
    message.adjanceyMatrixCount !== undefined && (obj.adjanceyMatrixCount = Math.round(message.adjanceyMatrixCount));
    if (message.LimitOrderTrancheUserList) {
      obj.LimitOrderTrancheUserList = message.LimitOrderTrancheUserList.map((e) =>
        e ? LimitOrderTrancheUser.toJSON(e) : undefined
      );
    } else {
      obj.LimitOrderTrancheUserList = [];
    }
    if (message.LimitOrderTrancheList) {
      obj.LimitOrderTrancheList = message.LimitOrderTrancheList.map((e) => e ? LimitOrderTranche.toJSON(e) : undefined);
    } else {
      obj.LimitOrderTrancheList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.tickMapList = object.tickMapList?.map((e) => TickMap.fromPartial(e)) || [];
    message.pairMapList = object.pairMapList?.map((e) => PairMap.fromPartial(e)) || [];
    message.tokensList = object.tokensList?.map((e) => Tokens.fromPartial(e)) || [];
    message.tokensCount = object.tokensCount ?? 0;
    message.tokenMapList = object.tokenMapList?.map((e) => TokenMap.fromPartial(e)) || [];
    message.sharesList = object.sharesList?.map((e) => Shares.fromPartial(e)) || [];
    message.feeListList = object.feeListList?.map((e) => FeeList.fromPartial(e)) || [];
    message.feeListCount = object.feeListCount ?? 0;
    message.edgeRowList = object.edgeRowList?.map((e) => EdgeRow.fromPartial(e)) || [];
    message.edgeRowCount = object.edgeRowCount ?? 0;
    message.adjanceyMatrixList = object.adjanceyMatrixList?.map((e) => AdjanceyMatrix.fromPartial(e)) || [];
    message.adjanceyMatrixCount = object.adjanceyMatrixCount ?? 0;
    message.LimitOrderTrancheUserList =
      object.LimitOrderTrancheUserList?.map((e) => LimitOrderTrancheUser.fromPartial(e)) || [];
    message.LimitOrderTrancheList = object.LimitOrderTrancheList?.map((e) => LimitOrderTranche.fromPartial(e)) || [];
    return message;
  },
};

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
