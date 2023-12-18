/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "onomyprotocol.reserve.reserve";

/** Params defines the parameters for the module. */
export interface Params {
  /** provider chain channel */
  provider_channel: string;
  /** market chain channel */
  market_channel: string;
  /** market_collateral is the ibc address for collateral on market chain */
  market_collateral: string;
  /** reserve_collateral is the ibc address for collateral on reserve chain */
  reserve_collateral: string;
  /** collateral_deposit is the amount of collateral needed to create a new denom */
  collateral_deposit: string;
}

const baseParams: object = {
  provider_channel: "",
  market_channel: "",
  market_collateral: "",
  reserve_collateral: "",
  collateral_deposit: "",
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.provider_channel !== "") {
      writer.uint32(10).string(message.provider_channel);
    }
    if (message.market_channel !== "") {
      writer.uint32(18).string(message.market_channel);
    }
    if (message.market_collateral !== "") {
      writer.uint32(26).string(message.market_collateral);
    }
    if (message.reserve_collateral !== "") {
      writer.uint32(34).string(message.reserve_collateral);
    }
    if (message.collateral_deposit !== "") {
      writer.uint32(42).string(message.collateral_deposit);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.provider_channel = reader.string();
          break;
        case 2:
          message.market_channel = reader.string();
          break;
        case 3:
          message.market_collateral = reader.string();
          break;
        case 4:
          message.reserve_collateral = reader.string();
          break;
        case 5:
          message.collateral_deposit = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (
      object.provider_channel !== undefined &&
      object.provider_channel !== null
    ) {
      message.provider_channel = String(object.provider_channel);
    } else {
      message.provider_channel = "";
    }
    if (object.market_channel !== undefined && object.market_channel !== null) {
      message.market_channel = String(object.market_channel);
    } else {
      message.market_channel = "";
    }
    if (
      object.market_collateral !== undefined &&
      object.market_collateral !== null
    ) {
      message.market_collateral = String(object.market_collateral);
    } else {
      message.market_collateral = "";
    }
    if (
      object.reserve_collateral !== undefined &&
      object.reserve_collateral !== null
    ) {
      message.reserve_collateral = String(object.reserve_collateral);
    } else {
      message.reserve_collateral = "";
    }
    if (
      object.collateral_deposit !== undefined &&
      object.collateral_deposit !== null
    ) {
      message.collateral_deposit = String(object.collateral_deposit);
    } else {
      message.collateral_deposit = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.provider_channel !== undefined &&
      (obj.provider_channel = message.provider_channel);
    message.market_channel !== undefined &&
      (obj.market_channel = message.market_channel);
    message.market_collateral !== undefined &&
      (obj.market_collateral = message.market_collateral);
    message.reserve_collateral !== undefined &&
      (obj.reserve_collateral = message.reserve_collateral);
    message.collateral_deposit !== undefined &&
      (obj.collateral_deposit = message.collateral_deposit);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (
      object.provider_channel !== undefined &&
      object.provider_channel !== null
    ) {
      message.provider_channel = object.provider_channel;
    } else {
      message.provider_channel = "";
    }
    if (object.market_channel !== undefined && object.market_channel !== null) {
      message.market_channel = object.market_channel;
    } else {
      message.market_channel = "";
    }
    if (
      object.market_collateral !== undefined &&
      object.market_collateral !== null
    ) {
      message.market_collateral = object.market_collateral;
    } else {
      message.market_collateral = "";
    }
    if (
      object.reserve_collateral !== undefined &&
      object.reserve_collateral !== null
    ) {
      message.reserve_collateral = object.reserve_collateral;
    } else {
      message.reserve_collateral = "";
    }
    if (
      object.collateral_deposit !== undefined &&
      object.collateral_deposit !== null
    ) {
      message.collateral_deposit = object.collateral_deposit;
    } else {
      message.collateral_deposit = "";
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
