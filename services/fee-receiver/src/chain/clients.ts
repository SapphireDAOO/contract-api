import { createPublicClient, createWalletClient, http } from "viem";
import { config } from "../config/config";
import { relayerAccount } from "./keys";

export const getClients = () => {
  const { chain, rpcUrl } = config();
  const transport = http(rpcUrl);
  const account = relayerAccount();

  return {
    relayerAccount: account,
    publicClient: createPublicClient({ chain, transport }),
    walletClient: createWalletClient({ account, chain, transport }),
  };
};
