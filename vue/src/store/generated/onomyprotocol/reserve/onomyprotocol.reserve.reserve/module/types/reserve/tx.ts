/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "onomyprotocol.reserve.reserve";

export interface MsgCreateVault {
  creator: string;
  collateral: string;
}

export interface MsgCreateVaultResponse {
  id: number;
}

export interface MsgDepositCollateral {
  creator: string;
  uid: number;
  collateral: string;
}

export interface MsgDepositCollateralResponse {}

export interface MsgMintDenom {
  creator: string;
  denom: string;
  amount: string;
}

export interface MsgMintDenomResponse {}

const baseMsgCreateVault: object = { creator: "", collateral: "" };

export const MsgCreateVault = {
  encode(message: MsgCreateVault, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collateral !== "") {
      writer.uint32(18).string(message.collateral);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVault {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateVault } as MsgCreateVault;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collateral = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVault {
    const message = { ...baseMsgCreateVault } as MsgCreateVault;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = String(object.collateral);
    } else {
      message.collateral = "";
    }
    return message;
  },

  toJSON(message: MsgCreateVault): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateVault>): MsgCreateVault {
    const message = { ...baseMsgCreateVault } as MsgCreateVault;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = object.collateral;
    } else {
      message.collateral = "";
    }
    return message;
  },
};

const baseMsgCreateVaultResponse: object = { id: 0 };

export const MsgCreateVaultResponse = {
  encode(
    message: MsgCreateVaultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVaultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateVaultResponse } as MsgCreateVaultResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVaultResponse {
    const message = { ...baseMsgCreateVaultResponse } as MsgCreateVaultResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateVaultResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateVaultResponse>
  ): MsgCreateVaultResponse {
    const message = { ...baseMsgCreateVaultResponse } as MsgCreateVaultResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDepositCollateral: object = {
  creator: "",
  uid: 0,
  collateral: "",
};

export const MsgDepositCollateral = {
  encode(
    message: MsgDepositCollateral,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uid !== 0) {
      writer.uint32(16).uint64(message.uid);
    }
    if (message.collateral !== "") {
      writer.uint32(26).string(message.collateral);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDepositCollateral {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDepositCollateral } as MsgDepositCollateral;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.collateral = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDepositCollateral {
    const message = { ...baseMsgDepositCollateral } as MsgDepositCollateral;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = String(object.collateral);
    } else {
      message.collateral = "";
    }
    return message;
  },

  toJSON(message: MsgDepositCollateral): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uid !== undefined && (obj.uid = message.uid);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDepositCollateral>): MsgDepositCollateral {
    const message = { ...baseMsgDepositCollateral } as MsgDepositCollateral;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = object.collateral;
    } else {
      message.collateral = "";
    }
    return message;
  },
};

const baseMsgDepositCollateralResponse: object = {};

export const MsgDepositCollateralResponse = {
  encode(
    _: MsgDepositCollateralResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDepositCollateralResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDepositCollateralResponse,
    } as MsgDepositCollateralResponse;
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

  fromJSON(_: any): MsgDepositCollateralResponse {
    const message = {
      ...baseMsgDepositCollateralResponse,
    } as MsgDepositCollateralResponse;
    return message;
  },

  toJSON(_: MsgDepositCollateralResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDepositCollateralResponse>
  ): MsgDepositCollateralResponse {
    const message = {
      ...baseMsgDepositCollateralResponse,
    } as MsgDepositCollateralResponse;
    return message;
  },
};

const baseMsgMintDenom: object = { creator: "", denom: "", amount: "" };

export const MsgMintDenom = {
  encode(message: MsgMintDenom, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintDenom {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintDenom } as MsgMintDenom;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denom = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMintDenom {
    const message = { ...baseMsgMintDenom } as MsgMintDenom;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: MsgMintDenom): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denom !== undefined && (obj.denom = message.denom);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMintDenom>): MsgMintDenom {
    const message = { ...baseMsgMintDenom } as MsgMintDenom;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseMsgMintDenomResponse: object = {};

export const MsgMintDenomResponse = {
  encode(_: MsgMintDenomResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintDenomResponse } as MsgMintDenomResponse;
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

  fromJSON(_: any): MsgMintDenomResponse {
    const message = { ...baseMsgMintDenomResponse } as MsgMintDenomResponse;
    return message;
  },

  toJSON(_: MsgMintDenomResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgMintDenomResponse>): MsgMintDenomResponse {
    const message = { ...baseMsgMintDenomResponse } as MsgMintDenomResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateVault(request: MsgCreateVault): Promise<MsgCreateVaultResponse>;
  DepositCollateral(
    request: MsgDepositCollateral
  ): Promise<MsgDepositCollateralResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  MintDenom(request: MsgMintDenom): Promise<MsgMintDenomResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateVault(request: MsgCreateVault): Promise<MsgCreateVaultResponse> {
    const data = MsgCreateVault.encode(request).finish();
    const promise = this.rpc.request(
      "onomyprotocol.reserve.reserve.Msg",
      "CreateVault",
      data
    );
    return promise.then((data) =>
      MsgCreateVaultResponse.decode(new Reader(data))
    );
  }

  DepositCollateral(
    request: MsgDepositCollateral
  ): Promise<MsgDepositCollateralResponse> {
    const data = MsgDepositCollateral.encode(request).finish();
    const promise = this.rpc.request(
      "onomyprotocol.reserve.reserve.Msg",
      "DepositCollateral",
      data
    );
    return promise.then((data) =>
      MsgDepositCollateralResponse.decode(new Reader(data))
    );
  }

  MintDenom(request: MsgMintDenom): Promise<MsgMintDenomResponse> {
    const data = MsgMintDenom.encode(request).finish();
    const promise = this.rpc.request(
      "onomyprotocol.reserve.reserve.Msg",
      "MintDenom",
      data
    );
    return promise.then((data) =>
      MsgMintDenomResponse.decode(new Reader(data))
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
