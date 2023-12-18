/* eslint-disable */
import { Metadata } from "../cosmos/bank/v1beta1/bank";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "onomyprotocol.dao.v1";

/** CreateDenomProposal details proposal that creates a new denom */
export interface CreateDenomProposal {
  sender: string;
  title: string;
  description: string;
  metadata: Metadata | undefined;
  reservedata: Reservedata | undefined;
  rate: string[];
}

export interface Escrow {
  proposer: string;
  amount: string;
}

export interface Reservedata {
  /** minimum collateralization ratio (parameter / 10000), 19999 representing as 199.99% */
  min_collateralization_ratio: string;
  /** liquidation ratio (parameter / 10000), 19999 representing as 199.99% */
  liquidation_ratio: string;
  /** interest rate (parameter / 10000), 9999 representing as 99.99% */
  interest_rate: string;
  /** savings rate (parameter / 10000), 9999 representing as 99.99% */
  savings_rate: string;
}

const baseCreateDenomProposal: object = {
  sender: "",
  title: "",
  description: "",
  rate: "",
};

export const CreateDenomProposal = {
  encode(
    message: CreateDenomProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(34).fork()).ldelim();
    }
    if (message.reservedata !== undefined) {
      Reservedata.encode(
        message.reservedata,
        writer.uint32(42).fork()
      ).ldelim();
    }
    for (const v of message.rate) {
      writer.uint32(50).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CreateDenomProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCreateDenomProposal } as CreateDenomProposal;
    message.rate = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        case 5:
          message.reservedata = Reservedata.decode(reader, reader.uint32());
          break;
        case 6:
          message.rate.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateDenomProposal {
    const message = { ...baseCreateDenomProposal } as CreateDenomProposal;
    message.rate = [];
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromJSON(object.metadata);
    } else {
      message.metadata = undefined;
    }
    if (object.reservedata !== undefined && object.reservedata !== null) {
      message.reservedata = Reservedata.fromJSON(object.reservedata);
    } else {
      message.reservedata = undefined;
    }
    if (object.rate !== undefined && object.rate !== null) {
      for (const e of object.rate) {
        message.rate.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: CreateDenomProposal): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    message.reservedata !== undefined &&
      (obj.reservedata = message.reservedata
        ? Reservedata.toJSON(message.reservedata)
        : undefined);
    if (message.rate) {
      obj.rate = message.rate.map((e) => e);
    } else {
      obj.rate = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<CreateDenomProposal>): CreateDenomProposal {
    const message = { ...baseCreateDenomProposal } as CreateDenomProposal;
    message.rate = [];
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromPartial(object.metadata);
    } else {
      message.metadata = undefined;
    }
    if (object.reservedata !== undefined && object.reservedata !== null) {
      message.reservedata = Reservedata.fromPartial(object.reservedata);
    } else {
      message.reservedata = undefined;
    }
    if (object.rate !== undefined && object.rate !== null) {
      for (const e of object.rate) {
        message.rate.push(e);
      }
    }
    return message;
  },
};

const baseEscrow: object = { proposer: "", amount: "" };

export const Escrow = {
  encode(message: Escrow, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== "") {
      writer.uint32(10).string(message.proposer);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Escrow {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEscrow } as Escrow;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Escrow {
    const message = { ...baseEscrow } as Escrow;
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer);
    } else {
      message.proposer = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: Escrow): unknown {
    const obj: any = {};
    message.proposer !== undefined && (obj.proposer = message.proposer);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<Escrow>): Escrow {
    const message = { ...baseEscrow } as Escrow;
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer;
    } else {
      message.proposer = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseReservedata: object = {
  min_collateralization_ratio: "",
  liquidation_ratio: "",
  interest_rate: "",
  savings_rate: "",
};

export const Reservedata = {
  encode(message: Reservedata, writer: Writer = Writer.create()): Writer {
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
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Reservedata {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseReservedata } as Reservedata;
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Reservedata {
    const message = { ...baseReservedata } as Reservedata;
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
    return message;
  },

  toJSON(message: Reservedata): unknown {
    const obj: any = {};
    message.min_collateralization_ratio !== undefined &&
      (obj.min_collateralization_ratio = message.min_collateralization_ratio);
    message.liquidation_ratio !== undefined &&
      (obj.liquidation_ratio = message.liquidation_ratio);
    message.interest_rate !== undefined &&
      (obj.interest_rate = message.interest_rate);
    message.savings_rate !== undefined &&
      (obj.savings_rate = message.savings_rate);
    return obj;
  },

  fromPartial(object: DeepPartial<Reservedata>): Reservedata {
    const message = { ...baseReservedata } as Reservedata;
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
