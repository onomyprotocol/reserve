// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateVault } from "./types/reserve/tx";
import { MsgBond } from "./types/reserve/tx";
import { MsgWithdraw } from "./types/reserve/tx";
import { MsgDeposit } from "./types/reserve/tx";
import { MsgUnbond } from "./types/reserve/tx";
import { MsgLiquidate } from "./types/reserve/tx";


const types = [
  ["/reserve.MsgCreateVault", MsgCreateVault],
  ["/reserve.MsgBond", MsgBond],
  ["/reserve.MsgWithdraw", MsgWithdraw],
  ["/reserve.MsgDeposit", MsgDeposit],
  ["/reserve.MsgUnbond", MsgUnbond],
  ["/reserve.MsgLiquidate", MsgLiquidate],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateVault: (data: MsgCreateVault): EncodeObject => ({ typeUrl: "/reserve.MsgCreateVault", value: MsgCreateVault.fromPartial( data ) }),
    msgBond: (data: MsgBond): EncodeObject => ({ typeUrl: "/reserve.MsgBond", value: MsgBond.fromPartial( data ) }),
    msgWithdraw: (data: MsgWithdraw): EncodeObject => ({ typeUrl: "/reserve.MsgWithdraw", value: MsgWithdraw.fromPartial( data ) }),
    msgDeposit: (data: MsgDeposit): EncodeObject => ({ typeUrl: "/reserve.MsgDeposit", value: MsgDeposit.fromPartial( data ) }),
    msgUnbond: (data: MsgUnbond): EncodeObject => ({ typeUrl: "/reserve.MsgUnbond", value: MsgUnbond.fromPartial( data ) }),
    msgLiquidate: (data: MsgLiquidate): EncodeObject => ({ typeUrl: "/reserve.MsgLiquidate", value: MsgLiquidate.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
