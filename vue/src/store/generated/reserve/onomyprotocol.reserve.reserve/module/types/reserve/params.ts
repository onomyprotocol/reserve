/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "onomyprotocol.reserve.reserve";

/** Params defines the parameters for the module. */
export interface Params {
  /** minimum collateralization ratio (parameter / 10000), 19999 representing as 199.99% */
  min_collateralization_ratio: string;
  /** liquidation ratio (parameter / 10000), 19999 representing as 199.99% */
  liquidation_ratio: string;
  /** interest rate (parameter / 10000), 9999 representing as 99.99% */
  interest_rate: string;
  /** savings rate (parameter / 10000), 9999 representing as 99.99% */
  savings_rate: string;
  /** provider chain channel */
  provider_channel: string;
  /** market chain channel */
  market_channel: string;
  /** market_coin is the ibc address for collateral on market chain */
  market_collateral: string;
  /** reserve_coin is the ibc address for collateral on reserve chain */
  reserve_collateral: string;
}

const baseParams: object = {
  min_collateralization_ratio: "",
  liquidation_ratio: "",
  interest_rate: "",
  savings_rate: "",
  provider_channel: "",
  market_channel: "",
  market_collateral: "",
  reserve_collateral: "",
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.min_collateralization_ratio !== "") {
      writer.uint32(10).string(message.min_collateralization_ratio);
    }
    if (message.liquidation_ratio !== "") {
      writer.uint32(18).string(message.liquidation_ratio);
    }
    if (message.interest_rate !== "") {
      writer.uint32(26).string(message.interest_rate);
    }
    if (message.savings_rate !== "") {
      writer.uint32(34).string(message.savings_rate);
    }
    if (message.provider_channel !== "") {
      writer.uint32(42).string(message.provider_channel);
    }
    if (message.market_channel !== "") {
      writer.uint32(50).string(message.market_channel);
    }
    if (message.market_collateral !== "") {
      writer.uint32(58).string(message.market_collateral);
    }
    if (message.reserve_collateral !== "") {
      writer.uint32(66).string(message.reserve_collateral);
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
          message.min_collateralization_ratio = reader.string();
          break;
        case 2:
          message.liquidation_ratio = reader.string();
          break;
        case 3:
          message.interest_rate = reader.string();
          break;
        case 4:
          message.savings_rate = reader.string();
          break;
        case 5:
          message.provider_channel = reader.string();
          break;
        case 6:
          message.market_channel = reader.string();
          break;
        case 7:
          message.market_collateral = reader.string();
          break;
        case 8:
          message.reserve_collateral = reader.string();
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
      object.min_collateralization_ratio !== undefined &&
      object.min_collateralization_ratio !== null
    ) {
      message.min_collateralization_ratio = String(
        object.min_collateralization_ratio
      );
    } else {
      message.min_collateralization_ratio = "";
    }
    if (
      object.liquidation_ratio !== undefined &&
      object.liquidation_ratio !== null
    ) {
      message.liquidation_ratio = String(object.liquidation_ratio);
    } else {
      message.liquidation_ratio = "";
    }
    if (object.interest_rate !== undefined && object.interest_rate !== null) {
      message.interest_rate = String(object.interest_rate);
    } else {
      message.interest_rate = "";
    }
    if (object.savings_rate !== undefined && object.savings_rate !== null) {
      message.savings_rate = String(object.savings_rate);
    } else {
      message.savings_rate = "";
    }
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
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.min_collateralization_ratio !== undefined &&
      (obj.min_collateralization_ratio = message.min_collateralization_ratio);
    message.liquidation_ratio !== undefined &&
      (obj.liquidation_ratio = message.liquidation_ratio);
    message.interest_rate !== undefined &&
      (obj.interest_rate = message.interest_rate);
    message.savings_rate !== undefined &&
      (obj.savings_rate = message.savings_rate);
    message.provider_channel !== undefined &&
      (obj.provider_channel = message.provider_channel);
    message.market_channel !== undefined &&
      (obj.market_channel = message.market_channel);
    message.market_collateral !== undefined &&
      (obj.market_collateral = message.market_collateral);
    message.reserve_collateral !== undefined &&
      (obj.reserve_collateral = message.reserve_collateral);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (
      object.min_collateralization_ratio !== undefined &&
      object.min_collateralization_ratio !== null
    ) {
      message.min_collateralization_ratio = object.min_collateralization_ratio;
    } else {
      message.min_collateralization_ratio = "";
    }
    if (
      object.liquidation_ratio !== undefined &&
      object.liquidation_ratio !== null
    ) {
      message.liquidation_ratio = object.liquidation_ratio;
    } else {
      message.liquidation_ratio = "";
    }
    if (object.interest_rate !== undefined && object.interest_rate !== null) {
      message.interest_rate = object.interest_rate;
    } else {
      message.interest_rate = "";
    }
    if (object.savings_rate !== undefined && object.savings_rate !== null) {
      message.savings_rate = object.savings_rate;
    } else {
      message.savings_rate = "";
    }
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
