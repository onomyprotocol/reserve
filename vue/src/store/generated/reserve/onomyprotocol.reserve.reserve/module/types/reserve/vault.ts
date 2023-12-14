/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "onomyprotocol.reserve.reserve";

export interface Vault {
  uid: number;
  collateral: string;
  principal: string;
  interest: string;
}

const baseVault: object = {
  uid: 0,
  collateral: "",
  principal: "",
  interest: "",
};

export const Vault = {
  encode(message: Vault, writer: Writer = Writer.create()): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    if (message.collateral !== "") {
      writer.uint32(18).string(message.collateral);
    }
    if (message.principal !== "") {
      writer.uint32(26).string(message.principal);
    }
    if (message.interest !== "") {
      writer.uint32(34).string(message.interest);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Vault {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVault } as Vault;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.collateral = reader.string();
          break;
        case 3:
          message.principal = reader.string();
          break;
        case 4:
          message.interest = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Vault {
    const message = { ...baseVault } as Vault;
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
    if (object.principal !== undefined && object.principal !== null) {
      message.principal = String(object.principal);
    } else {
      message.principal = "";
    }
    if (object.interest !== undefined && object.interest !== null) {
      message.interest = String(object.interest);
    } else {
      message.interest = "";
    }
    return message;
  },

  toJSON(message: Vault): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    message.principal !== undefined && (obj.principal = message.principal);
    message.interest !== undefined && (obj.interest = message.interest);
    return obj;
  },

  fromPartial(object: DeepPartial<Vault>): Vault {
    const message = { ...baseVault } as Vault;
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
    if (object.principal !== undefined && object.principal !== null) {
      message.principal = object.principal;
    } else {
      message.principal = "";
    }
    if (object.interest !== undefined && object.interest !== null) {
      message.interest = object.interest;
    } else {
      message.interest = "";
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
