/* eslint-disable */
import { Params } from "../dex/params";
import { Share } from "../dex/share";
import { Tick } from "../dex/tick";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** GenesisState defines the dex module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  shareList: Share[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  tickList: Tick[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.shareList) {
      Share.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.tickList) {
      Tick.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.shareList = [];
    message.tickList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.shareList.push(Share.decode(reader, reader.uint32()));
          break;
        case 3:
          message.tickList.push(Tick.decode(reader, reader.uint32()));
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
    message.shareList = [];
    message.tickList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.shareList !== undefined && object.shareList !== null) {
      for (const e of object.shareList) {
        message.shareList.push(Share.fromJSON(e));
      }
    }
    if (object.tickList !== undefined && object.tickList !== null) {
      for (const e of object.tickList) {
        message.tickList.push(Tick.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.shareList) {
      obj.shareList = message.shareList.map((e) =>
        e ? Share.toJSON(e) : undefined
      );
    } else {
      obj.shareList = [];
    }
    if (message.tickList) {
      obj.tickList = message.tickList.map((e) =>
        e ? Tick.toJSON(e) : undefined
      );
    } else {
      obj.tickList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.shareList = [];
    message.tickList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.shareList !== undefined && object.shareList !== null) {
      for (const e of object.shareList) {
        message.shareList.push(Share.fromPartial(e));
      }
    }
    if (object.tickList !== undefined && object.tickList !== null) {
      for (const e of object.tickList) {
        message.tickList.push(Tick.fromPartial(e));
      }
    }
    return message;
  },
};

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
