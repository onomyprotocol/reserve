/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve.reserve";

/** Params defines the parameters for the module. */
export interface Params {
  providerChannel: string;
  marketChannel: string;
  marketCollateral: string;
  reserveCollateral: string;
  collateralDeposit: string;
}

const baseParams: object = {
  providerChannel: "",
  marketChannel: "",
  marketCollateral: "",
  reserveCollateral: "",
  collateralDeposit: "",
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.providerChannel !== "") {
      writer.uint32(10).string(message.providerChannel);
    }
    if (message.marketChannel !== "") {
      writer.uint32(18).string(message.marketChannel);
    }
    if (message.marketCollateral !== "") {
      writer.uint32(26).string(message.marketCollateral);
    }
    if (message.reserveCollateral !== "") {
      writer.uint32(34).string(message.reserveCollateral);
    }
    if (message.collateralDeposit !== "") {
      writer.uint32(42).string(message.collateralDeposit);
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
          message.providerChannel = reader.string();
          break;
        case 2:
          message.marketChannel = reader.string();
          break;
        case 3:
          message.marketCollateral = reader.string();
          break;
        case 4:
          message.reserveCollateral = reader.string();
          break;
        case 5:
          message.collateralDeposit = reader.string();
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
      object.providerChannel !== undefined &&
      object.providerChannel !== null
    ) {
      message.providerChannel = String(object.providerChannel);
    } else {
      message.providerChannel = "";
    }
    if (object.marketChannel !== undefined && object.marketChannel !== null) {
      message.marketChannel = String(object.marketChannel);
    } else {
      message.marketChannel = "";
    }
    if (
      object.marketCollateral !== undefined &&
      object.marketCollateral !== null
    ) {
      message.marketCollateral = String(object.marketCollateral);
    } else {
      message.marketCollateral = "";
    }
    if (
      object.reserveCollateral !== undefined &&
      object.reserveCollateral !== null
    ) {
      message.reserveCollateral = String(object.reserveCollateral);
    } else {
      message.reserveCollateral = "";
    }
    if (
      object.collateralDeposit !== undefined &&
      object.collateralDeposit !== null
    ) {
      message.collateralDeposit = String(object.collateralDeposit);
    } else {
      message.collateralDeposit = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.providerChannel !== undefined &&
      (obj.providerChannel = message.providerChannel);
    message.marketChannel !== undefined &&
      (obj.marketChannel = message.marketChannel);
    message.marketCollateral !== undefined &&
      (obj.marketCollateral = message.marketCollateral);
    message.reserveCollateral !== undefined &&
      (obj.reserveCollateral = message.reserveCollateral);
    message.collateralDeposit !== undefined &&
      (obj.collateralDeposit = message.collateralDeposit);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (
      object.providerChannel !== undefined &&
      object.providerChannel !== null
    ) {
      message.providerChannel = object.providerChannel;
    } else {
      message.providerChannel = "";
    }
    if (object.marketChannel !== undefined && object.marketChannel !== null) {
      message.marketChannel = object.marketChannel;
    } else {
      message.marketChannel = "";
    }
    if (
      object.marketCollateral !== undefined &&
      object.marketCollateral !== null
    ) {
      message.marketCollateral = object.marketCollateral;
    } else {
      message.marketCollateral = "";
    }
    if (
      object.reserveCollateral !== undefined &&
      object.reserveCollateral !== null
    ) {
      message.reserveCollateral = object.reserveCollateral;
    } else {
      message.reserveCollateral = "";
    }
    if (
      object.collateralDeposit !== undefined &&
      object.collateralDeposit !== null
    ) {
      message.collateralDeposit = object.collateralDeposit;
    } else {
      message.collateralDeposit = "";
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
