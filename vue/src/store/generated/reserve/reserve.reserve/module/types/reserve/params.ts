/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve.reserve";

/** Params defines the parameters for the module. */
export interface Params {
  marketCollateral: string;
  reserveCollateral: string;
  collateralDeposit: string;
}

const baseParams: object = {
  marketCollateral: "",
  reserveCollateral: "",
  collateralDeposit: "",
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.marketCollateral !== "") {
      writer.uint32(10).string(message.marketCollateral);
    }
    if (message.reserveCollateral !== "") {
      writer.uint32(18).string(message.reserveCollateral);
    }
    if (message.collateralDeposit !== "") {
      writer.uint32(26).string(message.collateralDeposit);
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
          message.marketCollateral = reader.string();
          break;
        case 2:
          message.reserveCollateral = reader.string();
          break;
        case 3:
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
