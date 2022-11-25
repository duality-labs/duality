/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import {
  InfractionType,
  infractionTypeFromJSON,
  infractionTypeToJSON,
} from "../../../../cosmos/staking/v1beta1/staking";
import { Any } from "../../../../google/protobuf/any";
import { Duration } from "../../../../google/protobuf/duration";
import { SlashPacketData } from "../../v1/ccv";

export const protobufPackage = "interchain_security.ccv.consumer.v1";

/** Params defines the parameters for CCV consumer module */
export interface Params {
  /**
   * TODO: Remove enabled flag and find a better way to setup e2e tests
   * See: https://github.com/cosmos/interchain-security/issues/339
   */
  enabled: boolean;
  /**
   * /////////////////////
   * Distribution Params
   * Number of blocks between ibc-token-transfers from the consumer chain to
   * the provider chain. Note that at this transmission event a fraction of
   * the accumulated tokens are divided and sent consumer redistribution
   * address.
   */
  blocksPerDistributionTransmission: number;
  /**
   * Channel, and provider-chain receiving address to send distribution token
   * transfers over. These parameters is auto-set during the consumer <->
   * provider handshake procedure.
   */
  distributionTransmissionChannel: string;
  providerFeePoolAddrStr: string;
  /** Sent CCV related IBC packets will timeout after this duration */
  ccvTimeoutPeriod:
    | Duration
    | undefined;
  /** Sent transfer related IBC packets will timeout after this duration */
  transferTimeoutPeriod:
    | Duration
    | undefined;
  /**
   * The fraction of tokens allocated to the consumer redistribution address
   * during distribution events. The fraction is a string representing a
   * decimal number. For example "0.75" would represent 75%.
   */
  consumerRedistributionFraction: string;
  /**
   * The number of historical info entries to persist in store.
   * This param is a part of the cosmos sdk staking module. In the case of
   * a ccv enabled consumer chain, the ccv module acts as the staking module.
   */
  historicalEntries: number;
  /**
   * Unbonding period for the consumer,
   * which should be smaller than that of the provider in general.
   */
  unbondingPeriod: Duration | undefined;
}

/**
 * LastTransmissionBlockHeight is the last time validator holding
 * pools were transmitted to the provider chain
 */
export interface LastTransmissionBlockHeight {
  height: number;
}

/** CrossChainValidator defines the validators for CCV consumer module */
export interface CrossChainValidator {
  address: Uint8Array;
  power: number;
  /** pubkey is the consensus public key of the validator, as a Protobuf Any. */
  pubkey: Any | undefined;
}

/** SlashRequest defines a slashing request for CCV consumer module */
export interface SlashRequest {
  packet: SlashPacketData | undefined;
  infraction: InfractionType;
}

/** SlashRequests is a list of slash requests for CCV consumer module */
export interface SlashRequests {
  requests: SlashRequest[];
}

function createBaseParams(): Params {
  return {
    enabled: false,
    blocksPerDistributionTransmission: 0,
    distributionTransmissionChannel: "",
    providerFeePoolAddrStr: "",
    ccvTimeoutPeriod: undefined,
    transferTimeoutPeriod: undefined,
    consumerRedistributionFraction: "",
    historicalEntries: 0,
    unbondingPeriod: undefined,
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.enabled === true) {
      writer.uint32(8).bool(message.enabled);
    }
    if (message.blocksPerDistributionTransmission !== 0) {
      writer.uint32(16).int64(message.blocksPerDistributionTransmission);
    }
    if (message.distributionTransmissionChannel !== "") {
      writer.uint32(26).string(message.distributionTransmissionChannel);
    }
    if (message.providerFeePoolAddrStr !== "") {
      writer.uint32(34).string(message.providerFeePoolAddrStr);
    }
    if (message.ccvTimeoutPeriod !== undefined) {
      Duration.encode(message.ccvTimeoutPeriod, writer.uint32(42).fork()).ldelim();
    }
    if (message.transferTimeoutPeriod !== undefined) {
      Duration.encode(message.transferTimeoutPeriod, writer.uint32(50).fork()).ldelim();
    }
    if (message.consumerRedistributionFraction !== "") {
      writer.uint32(58).string(message.consumerRedistributionFraction);
    }
    if (message.historicalEntries !== 0) {
      writer.uint32(64).int64(message.historicalEntries);
    }
    if (message.unbondingPeriod !== undefined) {
      Duration.encode(message.unbondingPeriod, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.enabled = reader.bool();
          break;
        case 2:
          message.blocksPerDistributionTransmission = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.distributionTransmissionChannel = reader.string();
          break;
        case 4:
          message.providerFeePoolAddrStr = reader.string();
          break;
        case 5:
          message.ccvTimeoutPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 6:
          message.transferTimeoutPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 7:
          message.consumerRedistributionFraction = reader.string();
          break;
        case 8:
          message.historicalEntries = longToNumber(reader.int64() as Long);
          break;
        case 9:
          message.unbondingPeriod = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      enabled: isSet(object.enabled) ? Boolean(object.enabled) : false,
      blocksPerDistributionTransmission: isSet(object.blocksPerDistributionTransmission)
        ? Number(object.blocksPerDistributionTransmission)
        : 0,
      distributionTransmissionChannel: isSet(object.distributionTransmissionChannel)
        ? String(object.distributionTransmissionChannel)
        : "",
      providerFeePoolAddrStr: isSet(object.providerFeePoolAddrStr) ? String(object.providerFeePoolAddrStr) : "",
      ccvTimeoutPeriod: isSet(object.ccvTimeoutPeriod) ? Duration.fromJSON(object.ccvTimeoutPeriod) : undefined,
      transferTimeoutPeriod: isSet(object.transferTimeoutPeriod)
        ? Duration.fromJSON(object.transferTimeoutPeriod)
        : undefined,
      consumerRedistributionFraction: isSet(object.consumerRedistributionFraction)
        ? String(object.consumerRedistributionFraction)
        : "",
      historicalEntries: isSet(object.historicalEntries) ? Number(object.historicalEntries) : 0,
      unbondingPeriod: isSet(object.unbondingPeriod) ? Duration.fromJSON(object.unbondingPeriod) : undefined,
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.blocksPerDistributionTransmission !== undefined
      && (obj.blocksPerDistributionTransmission = Math.round(message.blocksPerDistributionTransmission));
    message.distributionTransmissionChannel !== undefined
      && (obj.distributionTransmissionChannel = message.distributionTransmissionChannel);
    message.providerFeePoolAddrStr !== undefined && (obj.providerFeePoolAddrStr = message.providerFeePoolAddrStr);
    message.ccvTimeoutPeriod !== undefined
      && (obj.ccvTimeoutPeriod = message.ccvTimeoutPeriod ? Duration.toJSON(message.ccvTimeoutPeriod) : undefined);
    message.transferTimeoutPeriod !== undefined && (obj.transferTimeoutPeriod = message.transferTimeoutPeriod
      ? Duration.toJSON(message.transferTimeoutPeriod)
      : undefined);
    message.consumerRedistributionFraction !== undefined
      && (obj.consumerRedistributionFraction = message.consumerRedistributionFraction);
    message.historicalEntries !== undefined && (obj.historicalEntries = Math.round(message.historicalEntries));
    message.unbondingPeriod !== undefined
      && (obj.unbondingPeriod = message.unbondingPeriod ? Duration.toJSON(message.unbondingPeriod) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.enabled = object.enabled ?? false;
    message.blocksPerDistributionTransmission = object.blocksPerDistributionTransmission ?? 0;
    message.distributionTransmissionChannel = object.distributionTransmissionChannel ?? "";
    message.providerFeePoolAddrStr = object.providerFeePoolAddrStr ?? "";
    message.ccvTimeoutPeriod = (object.ccvTimeoutPeriod !== undefined && object.ccvTimeoutPeriod !== null)
      ? Duration.fromPartial(object.ccvTimeoutPeriod)
      : undefined;
    message.transferTimeoutPeriod =
      (object.transferTimeoutPeriod !== undefined && object.transferTimeoutPeriod !== null)
        ? Duration.fromPartial(object.transferTimeoutPeriod)
        : undefined;
    message.consumerRedistributionFraction = object.consumerRedistributionFraction ?? "";
    message.historicalEntries = object.historicalEntries ?? 0;
    message.unbondingPeriod = (object.unbondingPeriod !== undefined && object.unbondingPeriod !== null)
      ? Duration.fromPartial(object.unbondingPeriod)
      : undefined;
    return message;
  },
};

function createBaseLastTransmissionBlockHeight(): LastTransmissionBlockHeight {
  return { height: 0 };
}

export const LastTransmissionBlockHeight = {
  encode(message: LastTransmissionBlockHeight, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.height !== 0) {
      writer.uint32(8).int64(message.height);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LastTransmissionBlockHeight {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLastTransmissionBlockHeight();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.height = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LastTransmissionBlockHeight {
    return { height: isSet(object.height) ? Number(object.height) : 0 };
  },

  toJSON(message: LastTransmissionBlockHeight): unknown {
    const obj: any = {};
    message.height !== undefined && (obj.height = Math.round(message.height));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LastTransmissionBlockHeight>, I>>(object: I): LastTransmissionBlockHeight {
    const message = createBaseLastTransmissionBlockHeight();
    message.height = object.height ?? 0;
    return message;
  },
};

function createBaseCrossChainValidator(): CrossChainValidator {
  return { address: new Uint8Array(), power: 0, pubkey: undefined };
}

export const CrossChainValidator = {
  encode(message: CrossChainValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address.length !== 0) {
      writer.uint32(10).bytes(message.address);
    }
    if (message.power !== 0) {
      writer.uint32(16).int64(message.power);
    }
    if (message.pubkey !== undefined) {
      Any.encode(message.pubkey, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CrossChainValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCrossChainValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.bytes();
          break;
        case 2:
          message.power = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.pubkey = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CrossChainValidator {
    return {
      address: isSet(object.address) ? bytesFromBase64(object.address) : new Uint8Array(),
      power: isSet(object.power) ? Number(object.power) : 0,
      pubkey: isSet(object.pubkey) ? Any.fromJSON(object.pubkey) : undefined,
    };
  },

  toJSON(message: CrossChainValidator): unknown {
    const obj: any = {};
    message.address !== undefined
      && (obj.address = base64FromBytes(message.address !== undefined ? message.address : new Uint8Array()));
    message.power !== undefined && (obj.power = Math.round(message.power));
    message.pubkey !== undefined && (obj.pubkey = message.pubkey ? Any.toJSON(message.pubkey) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CrossChainValidator>, I>>(object: I): CrossChainValidator {
    const message = createBaseCrossChainValidator();
    message.address = object.address ?? new Uint8Array();
    message.power = object.power ?? 0;
    message.pubkey = (object.pubkey !== undefined && object.pubkey !== null)
      ? Any.fromPartial(object.pubkey)
      : undefined;
    return message;
  },
};

function createBaseSlashRequest(): SlashRequest {
  return { packet: undefined, infraction: 0 };
}

export const SlashRequest = {
  encode(message: SlashRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.packet !== undefined) {
      SlashPacketData.encode(message.packet, writer.uint32(10).fork()).ldelim();
    }
    if (message.infraction !== 0) {
      writer.uint32(16).int32(message.infraction);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SlashRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSlashRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.packet = SlashPacketData.decode(reader, reader.uint32());
          break;
        case 2:
          message.infraction = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SlashRequest {
    return {
      packet: isSet(object.packet) ? SlashPacketData.fromJSON(object.packet) : undefined,
      infraction: isSet(object.infraction) ? infractionTypeFromJSON(object.infraction) : 0,
    };
  },

  toJSON(message: SlashRequest): unknown {
    const obj: any = {};
    message.packet !== undefined && (obj.packet = message.packet ? SlashPacketData.toJSON(message.packet) : undefined);
    message.infraction !== undefined && (obj.infraction = infractionTypeToJSON(message.infraction));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SlashRequest>, I>>(object: I): SlashRequest {
    const message = createBaseSlashRequest();
    message.packet = (object.packet !== undefined && object.packet !== null)
      ? SlashPacketData.fromPartial(object.packet)
      : undefined;
    message.infraction = object.infraction ?? 0;
    return message;
  },
};

function createBaseSlashRequests(): SlashRequests {
  return { requests: [] };
}

export const SlashRequests = {
  encode(message: SlashRequests, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.requests) {
      SlashRequest.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SlashRequests {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSlashRequests();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.requests.push(SlashRequest.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SlashRequests {
    return {
      requests: Array.isArray(object?.requests) ? object.requests.map((e: any) => SlashRequest.fromJSON(e)) : [],
    };
  },

  toJSON(message: SlashRequests): unknown {
    const obj: any = {};
    if (message.requests) {
      obj.requests = message.requests.map((e) => e ? SlashRequest.toJSON(e) : undefined);
    } else {
      obj.requests = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SlashRequests>, I>>(object: I): SlashRequests {
    const message = createBaseSlashRequests();
    message.requests = object.requests?.map((e) => SlashRequest.fromPartial(e)) || [];
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

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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
