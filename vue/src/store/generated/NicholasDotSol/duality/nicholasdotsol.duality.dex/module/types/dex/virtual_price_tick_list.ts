/* eslint-disable */
import { VirtualPriceTickQueue } from "../dex/virtual_price_tick_queue";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface VirtualPriceTickList {
  vPrice: string;
  direction: string;
  orderType: string;
  virtualTicks: VirtualPriceTickQueue | undefined;
}

const baseVirtualPriceTickList: object = {
  vPrice: "",
  direction: "",
  orderType: "",
};

export const VirtualPriceTickList = {
  encode(
    message: VirtualPriceTickList,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.vPrice !== "") {
      writer.uint32(10).string(message.vPrice);
    }
    if (message.direction !== "") {
      writer.uint32(18).string(message.direction);
    }
    if (message.orderType !== "") {
      writer.uint32(26).string(message.orderType);
    }
    if (message.virtualTicks !== undefined) {
      VirtualPriceTickQueue.encode(
        message.virtualTicks,
        writer.uint32(34).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VirtualPriceTickList {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVirtualPriceTickList } as VirtualPriceTickList;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vPrice = reader.string();
          break;
        case 2:
          message.direction = reader.string();
          break;
        case 3:
          message.orderType = reader.string();
          break;
        case 4:
          message.virtualTicks = VirtualPriceTickQueue.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VirtualPriceTickList {
    const message = { ...baseVirtualPriceTickList } as VirtualPriceTickList;
    if (object.vPrice !== undefined && object.vPrice !== null) {
      message.vPrice = String(object.vPrice);
    } else {
      message.vPrice = "";
    }
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = String(object.direction);
    } else {
      message.direction = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    if (object.virtualTicks !== undefined && object.virtualTicks !== null) {
      message.virtualTicks = VirtualPriceTickQueue.fromJSON(
        object.virtualTicks
      );
    } else {
      message.virtualTicks = undefined;
    }
    return message;
  },

  toJSON(message: VirtualPriceTickList): unknown {
    const obj: any = {};
    message.vPrice !== undefined && (obj.vPrice = message.vPrice);
    message.direction !== undefined && (obj.direction = message.direction);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.virtualTicks !== undefined &&
      (obj.virtualTicks = message.virtualTicks
        ? VirtualPriceTickQueue.toJSON(message.virtualTicks)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<VirtualPriceTickList>): VirtualPriceTickList {
    const message = { ...baseVirtualPriceTickList } as VirtualPriceTickList;
    if (object.vPrice !== undefined && object.vPrice !== null) {
      message.vPrice = object.vPrice;
    } else {
      message.vPrice = "";
    }
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = object.direction;
    } else {
      message.direction = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    if (object.virtualTicks !== undefined && object.virtualTicks !== null) {
      message.virtualTicks = VirtualPriceTickQueue.fromPartial(
        object.virtualTicks
      );
    } else {
      message.virtualTicks = undefined;
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
