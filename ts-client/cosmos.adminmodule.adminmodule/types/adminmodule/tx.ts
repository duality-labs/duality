/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Any } from "../google/protobuf/any";

export const protobufPackage = "cosmos.adminmodule.adminmodule";

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgDeleteAdmin {
  creator: string;
  admin: string;
}

export interface MsgDeleteAdminResponse {
}

export interface MsgAddAdmin {
  creator: string;
  admin: string;
}

export interface MsgAddAdminResponse {
}

/**
 * MsgSubmitProposal defines an sdk.Msg type that supports submitting arbitrary
 * proposal Content.
 */
export interface MsgSubmitProposal {
  content: Any | undefined;
  proposer: string;
}

/** MsgSubmitProposalResponse defines the Msg/SubmitProposal response type. */
export interface MsgSubmitProposalResponse {
  proposalId: number;
}

function createBaseMsgDeleteAdmin(): MsgDeleteAdmin {
  return { creator: "", admin: "" };
}

export const MsgDeleteAdmin = {
  encode(message: MsgDeleteAdmin, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.admin !== "") {
      writer.uint32(18).string(message.admin);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAdmin {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAdmin();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.admin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteAdmin {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      admin: isSet(object.admin) ? String(object.admin) : "",
    };
  },

  toJSON(message: MsgDeleteAdmin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.admin !== undefined && (obj.admin = message.admin);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAdmin>, I>>(object: I): MsgDeleteAdmin {
    const message = createBaseMsgDeleteAdmin();
    message.creator = object.creator ?? "";
    message.admin = object.admin ?? "";
    return message;
  },
};

function createBaseMsgDeleteAdminResponse(): MsgDeleteAdminResponse {
  return {};
}

export const MsgDeleteAdminResponse = {
  encode(_: MsgDeleteAdminResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAdminResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAdminResponse();
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

  fromJSON(_: any): MsgDeleteAdminResponse {
    return {};
  },

  toJSON(_: MsgDeleteAdminResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAdminResponse>, I>>(_: I): MsgDeleteAdminResponse {
    const message = createBaseMsgDeleteAdminResponse();
    return message;
  },
};

function createBaseMsgAddAdmin(): MsgAddAdmin {
  return { creator: "", admin: "" };
}

export const MsgAddAdmin = {
  encode(message: MsgAddAdmin, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.admin !== "") {
      writer.uint32(18).string(message.admin);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddAdmin {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddAdmin();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.admin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddAdmin {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      admin: isSet(object.admin) ? String(object.admin) : "",
    };
  },

  toJSON(message: MsgAddAdmin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.admin !== undefined && (obj.admin = message.admin);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddAdmin>, I>>(object: I): MsgAddAdmin {
    const message = createBaseMsgAddAdmin();
    message.creator = object.creator ?? "";
    message.admin = object.admin ?? "";
    return message;
  },
};

function createBaseMsgAddAdminResponse(): MsgAddAdminResponse {
  return {};
}

export const MsgAddAdminResponse = {
  encode(_: MsgAddAdminResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddAdminResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddAdminResponse();
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

  fromJSON(_: any): MsgAddAdminResponse {
    return {};
  },

  toJSON(_: MsgAddAdminResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddAdminResponse>, I>>(_: I): MsgAddAdminResponse {
    const message = createBaseMsgAddAdminResponse();
    return message;
  },
};

function createBaseMsgSubmitProposal(): MsgSubmitProposal {
  return { content: undefined, proposer: "" };
}

export const MsgSubmitProposal = {
  encode(message: MsgSubmitProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.content !== undefined) {
      Any.encode(message.content, writer.uint32(10).fork()).ldelim();
    }
    if (message.proposer !== "") {
      writer.uint32(18).string(message.proposer);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.content = Any.decode(reader, reader.uint32());
          break;
        case 2:
          message.proposer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitProposal {
    return {
      content: isSet(object.content) ? Any.fromJSON(object.content) : undefined,
      proposer: isSet(object.proposer) ? String(object.proposer) : "",
    };
  },

  toJSON(message: MsgSubmitProposal): unknown {
    const obj: any = {};
    message.content !== undefined && (obj.content = message.content ? Any.toJSON(message.content) : undefined);
    message.proposer !== undefined && (obj.proposer = message.proposer);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitProposal>, I>>(object: I): MsgSubmitProposal {
    const message = createBaseMsgSubmitProposal();
    message.content = (object.content !== undefined && object.content !== null)
      ? Any.fromPartial(object.content)
      : undefined;
    message.proposer = object.proposer ?? "";
    return message;
  },
};

function createBaseMsgSubmitProposalResponse(): MsgSubmitProposalResponse {
  return { proposalId: 0 };
}

export const MsgSubmitProposalResponse = {
  encode(message: MsgSubmitProposalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.proposalId !== 0) {
      writer.uint32(8).uint64(message.proposalId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitProposalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitProposalResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposalId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitProposalResponse {
    return { proposalId: isSet(object.proposalId) ? Number(object.proposalId) : 0 };
  },

  toJSON(message: MsgSubmitProposalResponse): unknown {
    const obj: any = {};
    message.proposalId !== undefined && (obj.proposalId = Math.round(message.proposalId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitProposalResponse>, I>>(object: I): MsgSubmitProposalResponse {
    const message = createBaseMsgSubmitProposalResponse();
    message.proposalId = object.proposalId ?? 0;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  DeleteAdmin(request: MsgDeleteAdmin): Promise<MsgDeleteAdminResponse>;
  AddAdmin(request: MsgAddAdmin): Promise<MsgAddAdminResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SubmitProposal(request: MsgSubmitProposal): Promise<MsgSubmitProposalResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.DeleteAdmin = this.DeleteAdmin.bind(this);
    this.AddAdmin = this.AddAdmin.bind(this);
    this.SubmitProposal = this.SubmitProposal.bind(this);
  }
  DeleteAdmin(request: MsgDeleteAdmin): Promise<MsgDeleteAdminResponse> {
    const data = MsgDeleteAdmin.encode(request).finish();
    const promise = this.rpc.request("cosmos.adminmodule.adminmodule.Msg", "DeleteAdmin", data);
    return promise.then((data) => MsgDeleteAdminResponse.decode(new _m0.Reader(data)));
  }

  AddAdmin(request: MsgAddAdmin): Promise<MsgAddAdminResponse> {
    const data = MsgAddAdmin.encode(request).finish();
    const promise = this.rpc.request("cosmos.adminmodule.adminmodule.Msg", "AddAdmin", data);
    return promise.then((data) => MsgAddAdminResponse.decode(new _m0.Reader(data)));
  }

  SubmitProposal(request: MsgSubmitProposal): Promise<MsgSubmitProposalResponse> {
    const data = MsgSubmitProposal.encode(request).finish();
    const promise = this.rpc.request("cosmos.adminmodule.adminmodule.Msg", "SubmitProposal", data);
    return promise.then((data) => MsgSubmitProposalResponse.decode(new _m0.Reader(data)));
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
