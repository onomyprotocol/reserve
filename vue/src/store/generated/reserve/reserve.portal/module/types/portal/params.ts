/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve.portal";

/** Params defines the parameters for the module. */
export interface Params {
  providerChannel: string;
  marketChannel: string;
}

const baseParams: object = { providerChannel: "", marketChannel: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.providerChannel !== "") {
      writer.uint32(10).string(message.providerChannel);
    }
    if (message.marketChannel !== "") {
      writer.uint32(18).string(message.marketChannel);
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
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.providerChannel !== undefined &&
      (obj.providerChannel = message.providerChannel);
    message.marketChannel !== undefined &&
      (obj.marketChannel = message.marketChannel);
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
