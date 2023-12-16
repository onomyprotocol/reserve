// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDepositCollateral } from "./types/reserve/tx";
import { MsgMintDenom } from "./types/reserve/tx";
import { MsgCreateVault } from "./types/reserve/tx";


const types = [
  ["/onomyprotocol.reserve.reserve.MsgDepositCollateral", MsgDepositCollateral],
  ["/onomyprotocol.reserve.reserve.MsgMintDenom", MsgMintDenom],
  ["/onomyprotocol.reserve.reserve.MsgCreateVault", MsgCreateVault],
  
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
    msgDepositCollateral: (data: MsgDepositCollateral): EncodeObject => ({ typeUrl: "/onomyprotocol.reserve.reserve.MsgDepositCollateral", value: MsgDepositCollateral.fromPartial( data ) }),
    msgMintDenom: (data: MsgMintDenom): EncodeObject => ({ typeUrl: "/onomyprotocol.reserve.reserve.MsgMintDenom", value: MsgMintDenom.fromPartial( data ) }),
    msgCreateVault: (data: MsgCreateVault): EncodeObject => ({ typeUrl: "/onomyprotocol.reserve.reserve.MsgCreateVault", value: MsgCreateVault.fromPartial( data ) }),
    
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
