/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "onomyprotocol.reserve.reserve";

/** Params defines the parameters for the module. */
export interface Params {
  /** minimum collateralization ratio (parameter / 10000), 19999 representing as 199.99% */
  m_c_r: string;
  /** liquidation ratio (parameter / 10000), 19999 representing as 199.99% */
  l_r: string;
  /** interest rate (parameter / 10000), 9999 representing as 99.99% */
  i_r: string;
  /** savings rate (parameter / 10000), 9999 representing as 99.99% */
  s_r: string;
  /** market chain channel */
  market_channel: string;
  /** market_coin is the ibc address for collateral on market chain */
  market_coin: string;
  /** reserve_coin is the ibc address for collateral on reserve chain */
  reserve_coin: string;
}

const baseParams: object = {
  m_c_r: "",
  l_r: "",
  i_r: "",
  s_r: "",
  market_channel: "",
  market_coin: "",
  reserve_coin: "",
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.m_c_r !== "") {
      writer.uint32(10).string(message.m_c_r);
    }
    if (message.l_r !== "") {
      writer.uint32(18).string(message.l_r);
    }
    if (message.i_r !== "") {
      writer.uint32(26).string(message.i_r);
    }
    if (message.s_r !== "") {
      writer.uint32(34).string(message.s_r);
    }
    if (message.market_channel !== "") {
      writer.uint32(42).string(message.market_channel);
    }
    if (message.market_coin !== "") {
      writer.uint32(50).string(message.market_coin);
    }
    if (message.reserve_coin !== "") {
      writer.uint32(58).string(message.reserve_coin);
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
          message.m_c_r = reader.string();
          break;
        case 2:
          message.l_r = reader.string();
          break;
        case 3:
          message.i_r = reader.string();
          break;
        case 4:
          message.s_r = reader.string();
          break;
        case 5:
          message.market_channel = reader.string();
          break;
        case 6:
          message.market_coin = reader.string();
          break;
        case 7:
          message.reserve_coin = reader.string();
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
    if (object.m_c_r !== undefined && object.m_c_r !== null) {
      message.m_c_r = String(object.m_c_r);
    } else {
      message.m_c_r = "";
    }
    if (object.l_r !== undefined && object.l_r !== null) {
      message.l_r = String(object.l_r);
    } else {
      message.l_r = "";
    }
    if (object.i_r !== undefined && object.i_r !== null) {
      message.i_r = String(object.i_r);
    } else {
      message.i_r = "";
    }
    if (object.s_r !== undefined && object.s_r !== null) {
      message.s_r = String(object.s_r);
    } else {
      message.s_r = "";
    }
    if (object.market_channel !== undefined && object.market_channel !== null) {
      message.market_channel = String(object.market_channel);
    } else {
      message.market_channel = "";
    }
    if (object.market_coin !== undefined && object.market_coin !== null) {
      message.market_coin = String(object.market_coin);
    } else {
      message.market_coin = "";
    }
    if (object.reserve_coin !== undefined && object.reserve_coin !== null) {
      message.reserve_coin = String(object.reserve_coin);
    } else {
      message.reserve_coin = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.m_c_r !== undefined && (obj.m_c_r = message.m_c_r);
    message.l_r !== undefined && (obj.l_r = message.l_r);
    message.i_r !== undefined && (obj.i_r = message.i_r);
    message.s_r !== undefined && (obj.s_r = message.s_r);
    message.market_channel !== undefined &&
      (obj.market_channel = message.market_channel);
    message.market_coin !== undefined &&
      (obj.market_coin = message.market_coin);
    message.reserve_coin !== undefined &&
      (obj.reserve_coin = message.reserve_coin);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.m_c_r !== undefined && object.m_c_r !== null) {
      message.m_c_r = object.m_c_r;
    } else {
      message.m_c_r = "";
    }
    if (object.l_r !== undefined && object.l_r !== null) {
      message.l_r = object.l_r;
    } else {
      message.l_r = "";
    }
    if (object.i_r !== undefined && object.i_r !== null) {
      message.i_r = object.i_r;
    } else {
      message.i_r = "";
    }
    if (object.s_r !== undefined && object.s_r !== null) {
      message.s_r = object.s_r;
    } else {
      message.s_r = "";
    }
    if (object.market_channel !== undefined && object.market_channel !== null) {
      message.market_channel = object.market_channel;
    } else {
      message.market_channel = "";
    }
    if (object.market_coin !== undefined && object.market_coin !== null) {
      message.market_coin = object.market_coin;
    } else {
      message.market_coin = "";
    }
    if (object.reserve_coin !== undefined && object.reserve_coin !== null) {
      message.reserve_coin = object.reserve_coin;
    } else {
      message.reserve_coin = "";
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
