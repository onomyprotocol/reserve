/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "onomyprotocol.reserve.portal";

export interface PortalPacketData {
  noData: NoData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  subscribeRatePacket: SubscribeRatePacketData | undefined;
}

export interface NoData {}

/** SubscribeRatePacketData defines a struct for the packet payload */
export interface SubscribeRatePacketData {
  chain: string;
  denom: string;
}

/** SubscribeRatePacketAck defines a struct for the packet acknowledgment */
export interface SubscribeRatePacketAck {
  success: boolean;
}

const basePortalPacketData: object = {};

export const PortalPacketData = {
  encode(message: PortalPacketData, writer: Writer = Writer.create()): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    if (message.subscribeRatePacket !== undefined) {
      SubscribeRatePacketData.encode(
        message.subscribeRatePacket,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PortalPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePortalPacketData } as PortalPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        case 2:
          message.subscribeRatePacket = SubscribeRatePacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PortalPacketData {
    const message = { ...basePortalPacketData } as PortalPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.subscribeRatePacket !== undefined &&
      object.subscribeRatePacket !== null
    ) {
      message.subscribeRatePacket = SubscribeRatePacketData.fromJSON(
        object.subscribeRatePacket
      );
    } else {
      message.subscribeRatePacket = undefined;
    }
    return message;
  },

  toJSON(message: PortalPacketData): unknown {
    const obj: any = {};
    message.noData !== undefined &&
      (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    message.subscribeRatePacket !== undefined &&
      (obj.subscribeRatePacket = message.subscribeRatePacket
        ? SubscribeRatePacketData.toJSON(message.subscribeRatePacket)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<PortalPacketData>): PortalPacketData {
    const message = { ...basePortalPacketData } as PortalPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.subscribeRatePacket !== undefined &&
      object.subscribeRatePacket !== null
    ) {
      message.subscribeRatePacket = SubscribeRatePacketData.fromPartial(
        object.subscribeRatePacket
      );
    } else {
      message.subscribeRatePacket = undefined;
    }
    return message;
  },
};

const baseNoData: object = {};

export const NoData = {
  encode(_: NoData, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNoData } as NoData;
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

  fromJSON(_: any): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },

  toJSON(_: NoData): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<NoData>): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },
};

const baseSubscribeRatePacketData: object = { chain: "", denom: "" };

export const SubscribeRatePacketData = {
  encode(
    message: SubscribeRatePacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.chain !== "") {
      writer.uint32(10).string(message.chain);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SubscribeRatePacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseSubscribeRatePacketData,
    } as SubscribeRatePacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chain = reader.string();
          break;
        case 2:
          message.denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubscribeRatePacketData {
    const message = {
      ...baseSubscribeRatePacketData,
    } as SubscribeRatePacketData;
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    return message;
  },

  toJSON(message: SubscribeRatePacketData): unknown {
    const obj: any = {};
    message.chain !== undefined && (obj.chain = message.chain);
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial(
    object: DeepPartial<SubscribeRatePacketData>
  ): SubscribeRatePacketData {
    const message = {
      ...baseSubscribeRatePacketData,
    } as SubscribeRatePacketData;
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    return message;
  },
};

const baseSubscribeRatePacketAck: object = { success: false };

export const SubscribeRatePacketAck = {
  encode(
    message: SubscribeRatePacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SubscribeRatePacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSubscribeRatePacketAck } as SubscribeRatePacketAck;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubscribeRatePacketAck {
    const message = { ...baseSubscribeRatePacketAck } as SubscribeRatePacketAck;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: SubscribeRatePacketAck): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<SubscribeRatePacketAck>
  ): SubscribeRatePacketAck {
    const message = { ...baseSubscribeRatePacketAck } as SubscribeRatePacketAck;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
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
