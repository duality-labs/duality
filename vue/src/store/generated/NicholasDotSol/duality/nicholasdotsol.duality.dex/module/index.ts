// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreatePair } from "./types/dex/tx";
import { MsgAddLiquidity } from "./types/dex/tx";
import { MsgSwap } from "./types/dex/tx";
import { MsgRemoveLiquidity } from "./types/dex/tx";


const types = [
  ["/nicholasdotsol.duality.dex.MsgCreatePair", MsgCreatePair],
  ["/nicholasdotsol.duality.dex.MsgAddLiquidity", MsgAddLiquidity],
  ["/nicholasdotsol.duality.dex.MsgSwap", MsgSwap],
  ["/nicholasdotsol.duality.dex.MsgRemoveLiquidity", MsgRemoveLiquidity],
  
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
    msgCreatePair: (data: MsgCreatePair): EncodeObject => ({ typeUrl: "/nicholasdotsol.duality.dex.MsgCreatePair", value: MsgCreatePair.fromPartial( data ) }),
    msgAddLiquidity: (data: MsgAddLiquidity): EncodeObject => ({ typeUrl: "/nicholasdotsol.duality.dex.MsgAddLiquidity", value: MsgAddLiquidity.fromPartial( data ) }),
    msgSwap: (data: MsgSwap): EncodeObject => ({ typeUrl: "/nicholasdotsol.duality.dex.MsgSwap", value: MsgSwap.fromPartial( data ) }),
    msgRemoveLiquidity: (data: MsgRemoveLiquidity): EncodeObject => ({ typeUrl: "/nicholasdotsol.duality.dex.MsgRemoveLiquidity", value: MsgRemoveLiquidity.fromPartial( data ) }),
    
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
