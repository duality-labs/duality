/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { InfractionType, infractionTypeFromJSON, infractionTypeToJSON } from "../../../cosmos/staking/v1beta1/staking";
import { Validator, ValidatorUpdate } from "../../../tendermint/abci/types";

export const protobufPackage = "interchain_security.ccv.v1";

/**
 * This packet is sent from provider chain to consumer chain if the validator
 * set for consumer chain changes (due to new bonding/unbonding messages or
 * slashing events) A VSCMatured packet from consumer chain will be sent
 * asynchronously once unbonding period is over, and this will function as
 * `UnbondingOver` message for this packet.
 */
export interface ValidatorSetChangePacketData {
  validatorUpdates: ValidatorUpdate[];
  valsetUpdateId: number;
  /**
   * consensus address of consumer chain validators
   * successfully slashed on the provider chain
   */
  slashAcks: string[];
}

export interface UnbondingOp {
  id: number;
  /** consumer chains that are still unbonding */
  unbondingConsumerChains: string[];
}

/**
 * This packet is sent from the consumer chain to the provider chain
 * to notify that a VSC packet reached maturity on the consumer chain.
 */
export interface VSCMaturedPacketData {
  /** the id of the VSC packet that reached maturity */
  valsetUpdateId: number;
}

/**
 * This packet is sent from the consumer chain to the provider chain
 * to request the slashing of a validator as a result of an infraction
 * committed on the consumer chain.
 */
export interface SlashPacketData {
  validator:
    | Validator
    | undefined;
  /** map to the infraction block height on the provider */
  valsetUpdateId: number;
  /** tell if the slashing is for a downtime or a double-signing infraction */
  infraction: InfractionType;
}

/** UnbondingOpsIndex defines a list of unbonding operation ids. */
export interface UnbondingOpsIndex {
  ids: number[];
}

/** MaturedUnbondingOps defines a list of ids corresponding to ids of matured unbonding operations. */
export interface MaturedUnbondingOps {
  ids: number[];
}

function createBaseValidatorSetChangePacketData(): ValidatorSetChangePacketData {
  return { validatorUpdates: [], valsetUpdateId: 0, slashAcks: [] };
}

export const ValidatorSetChangePacketData = {
  encode(message: ValidatorSetChangePacketData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.validatorUpdates) {
      ValidatorUpdate.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.valsetUpdateId !== 0) {
      writer.uint32(16).uint64(message.valsetUpdateId);
    }
    for (const v of message.slashAcks) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ValidatorSetChangePacketData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidatorSetChangePacketData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validatorUpdates.push(ValidatorUpdate.decode(reader, reader.uint32()));
          break;
        case 2:
          message.valsetUpdateId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.slashAcks.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ValidatorSetChangePacketData {
    return {
      validatorUpdates: Array.isArray(object?.validatorUpdates)
        ? object.validatorUpdates.map((e: any) => ValidatorUpdate.fromJSON(e))
        : [],
      valsetUpdateId: isSet(object.valsetUpdateId) ? Number(object.valsetUpdateId) : 0,
      slashAcks: Array.isArray(object?.slashAcks) ? object.slashAcks.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: ValidatorSetChangePacketData): unknown {
    const obj: any = {};
    if (message.validatorUpdates) {
      obj.validatorUpdates = message.validatorUpdates.map((e) => e ? ValidatorUpdate.toJSON(e) : undefined);
    } else {
      obj.validatorUpdates = [];
    }
    message.valsetUpdateId !== undefined && (obj.valsetUpdateId = Math.round(message.valsetUpdateId));
    if (message.slashAcks) {
      obj.slashAcks = message.slashAcks.map((e) => e);
    } else {
      obj.slashAcks = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ValidatorSetChangePacketData>, I>>(object: I): ValidatorSetChangePacketData {
    const message = createBaseValidatorSetChangePacketData();
    message.validatorUpdates = object.validatorUpdates?.map((e) => ValidatorUpdate.fromPartial(e)) || [];
    message.valsetUpdateId = object.valsetUpdateId ?? 0;
    message.slashAcks = object.slashAcks?.map((e) => e) || [];
    return message;
  },
};

function createBaseUnbondingOp(): UnbondingOp {
  return { id: 0, unbondingConsumerChains: [] };
}

export const UnbondingOp = {
  encode(message: UnbondingOp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    for (const v of message.unbondingConsumerChains) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnbondingOp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnbondingOp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.unbondingConsumerChains.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnbondingOp {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      unbondingConsumerChains: Array.isArray(object?.unbondingConsumerChains)
        ? object.unbondingConsumerChains.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: UnbondingOp): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    if (message.unbondingConsumerChains) {
      obj.unbondingConsumerChains = message.unbondingConsumerChains.map((e) => e);
    } else {
      obj.unbondingConsumerChains = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnbondingOp>, I>>(object: I): UnbondingOp {
    const message = createBaseUnbondingOp();
    message.id = object.id ?? 0;
    message.unbondingConsumerChains = object.unbondingConsumerChains?.map((e) => e) || [];
    return message;
  },
};

function createBaseVSCMaturedPacketData(): VSCMaturedPacketData {
  return { valsetUpdateId: 0 };
}

export const VSCMaturedPacketData = {
  encode(message: VSCMaturedPacketData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.valsetUpdateId !== 0) {
      writer.uint32(8).uint64(message.valsetUpdateId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VSCMaturedPacketData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVSCMaturedPacketData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.valsetUpdateId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VSCMaturedPacketData {
    return { valsetUpdateId: isSet(object.valsetUpdateId) ? Number(object.valsetUpdateId) : 0 };
  },

  toJSON(message: VSCMaturedPacketData): unknown {
    const obj: any = {};
    message.valsetUpdateId !== undefined && (obj.valsetUpdateId = Math.round(message.valsetUpdateId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VSCMaturedPacketData>, I>>(object: I): VSCMaturedPacketData {
    const message = createBaseVSCMaturedPacketData();
    message.valsetUpdateId = object.valsetUpdateId ?? 0;
    return message;
  },
};

function createBaseSlashPacketData(): SlashPacketData {
  return { validator: undefined, valsetUpdateId: 0, infraction: 0 };
}

export const SlashPacketData = {
  encode(message: SlashPacketData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.validator !== undefined) {
      Validator.encode(message.validator, writer.uint32(10).fork()).ldelim();
    }
    if (message.valsetUpdateId !== 0) {
      writer.uint32(16).uint64(message.valsetUpdateId);
    }
    if (message.infraction !== 0) {
      writer.uint32(24).int32(message.infraction);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SlashPacketData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSlashPacketData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator = Validator.decode(reader, reader.uint32());
          break;
        case 2:
          message.valsetUpdateId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.infraction = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SlashPacketData {
    return {
      validator: isSet(object.validator) ? Validator.fromJSON(object.validator) : undefined,
      valsetUpdateId: isSet(object.valsetUpdateId) ? Number(object.valsetUpdateId) : 0,
      infraction: isSet(object.infraction) ? infractionTypeFromJSON(object.infraction) : 0,
    };
  },

  toJSON(message: SlashPacketData): unknown {
    const obj: any = {};
    message.validator !== undefined
      && (obj.validator = message.validator ? Validator.toJSON(message.validator) : undefined);
    message.valsetUpdateId !== undefined && (obj.valsetUpdateId = Math.round(message.valsetUpdateId));
    message.infraction !== undefined && (obj.infraction = infractionTypeToJSON(message.infraction));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SlashPacketData>, I>>(object: I): SlashPacketData {
    const message = createBaseSlashPacketData();
    message.validator = (object.validator !== undefined && object.validator !== null)
      ? Validator.fromPartial(object.validator)
      : undefined;
    message.valsetUpdateId = object.valsetUpdateId ?? 0;
    message.infraction = object.infraction ?? 0;
    return message;
  },
};

function createBaseUnbondingOpsIndex(): UnbondingOpsIndex {
  return { ids: [] };
}

export const UnbondingOpsIndex = {
  encode(message: UnbondingOpsIndex, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.ids) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnbondingOpsIndex {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnbondingOpsIndex();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ids.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.ids.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnbondingOpsIndex {
    return { ids: Array.isArray(object?.ids) ? object.ids.map((e: any) => Number(e)) : [] };
  },

  toJSON(message: UnbondingOpsIndex): unknown {
    const obj: any = {};
    if (message.ids) {
      obj.ids = message.ids.map((e) => Math.round(e));
    } else {
      obj.ids = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnbondingOpsIndex>, I>>(object: I): UnbondingOpsIndex {
    const message = createBaseUnbondingOpsIndex();
    message.ids = object.ids?.map((e) => e) || [];
    return message;
  },
};

function createBaseMaturedUnbondingOps(): MaturedUnbondingOps {
  return { ids: [] };
}

export const MaturedUnbondingOps = {
  encode(message: MaturedUnbondingOps, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.ids) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MaturedUnbondingOps {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMaturedUnbondingOps();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ids.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.ids.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MaturedUnbondingOps {
    return { ids: Array.isArray(object?.ids) ? object.ids.map((e: any) => Number(e)) : [] };
  },

  toJSON(message: MaturedUnbondingOps): unknown {
    const obj: any = {};
    if (message.ids) {
      obj.ids = message.ids.map((e) => Math.round(e));
    } else {
      obj.ids = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MaturedUnbondingOps>, I>>(object: I): MaturedUnbondingOps {
    const message = createBaseMaturedUnbondingOps();
    message.ids = object.ids?.map((e) => e) || [];
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
