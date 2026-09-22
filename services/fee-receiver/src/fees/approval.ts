import { randomBytes } from "node:crypto";
import {
  contracts,
  createExecution,
  ExecutionMode,
  getSmartAccountsEnvironment,
  ROOT_AUTHORITY,
  signDelegation,
} from "@metamask/smart-accounts-kit";
import {
  encodeFunctionData,
  erc20Abi,
  maxUint256,
  zeroAddress,
  type Address,
  type Hex,
  type PrivateKeyAccount,
} from "viem";
import { getClients } from "../chain/clients";
import { chainId, config } from "../config/config";
import { FeeReceiverUnavailableError } from "../errors";
import { logger } from "../logger";

/**
 * Resolves the token the stealth account must approve: the payment token
 * itself, or the chain's wrapped native token when the payment is in native
 * currency.
 */
export const resolveApprovalToken = (paymentToken?: Address): Address => {
  if (paymentToken && paymentToken.toLowerCase() !== zeroAddress) {
    return paymentToken;
  }

  const { wrappedNative, network } = config();
  if (!wrappedNative) {
    throw new Error(
      `No wrapped native token configured for network ${network}`,
    );
  }
  return wrappedNative;
};

export const delegateAndApprove = async (
  stealthAccount: PrivateKeyAccount,
  stealthPrivateKey: Hex,
  paymentToken?: Address,
): Promise<void> => {
  const chain = chainId();

  let environment;
  try {
    environment = getSmartAccountsEnvironment(chain);
  } catch {
    logger.warn("no smart accounts environment; skipping 7702 delegation", {
      chainId: chain,
    });
    return;
  }

  const approvalToken = resolveApprovalToken(paymentToken);

  const { sweeper, network } = config();
  if (!sweeper) {
    throw new Error(`No sweeper contract configured for network ${network}`);
  }

  const { relayerAccount, publicClient, walletClient } = getClients();

  const delegatorImpl =
    environment.implementations.EIP7702StatelessDeleGatorImpl;
  if (!delegatorImpl) {
    throw new Error(
      `No EIP7702StatelessDeleGatorImpl in the smart accounts environment for chain ${chain}`,
    );
  }

  const nonce = await publicClient.getTransactionCount({
    address: stealthAccount.address,
  });
  const authorization = await stealthAccount.signAuthorization({
    address: delegatorImpl,
    chainId: chain,
    nonce,
  });

  const delegation = {
    delegate: relayerAccount.address as Hex,
    delegator: stealthAccount.address as Hex,
    authority: ROOT_AUTHORITY as Hex,
    caveats: [],
    salt: `0x${randomBytes(32).toString("hex")}` as Hex,
  };
  const delegationSignature = await signDelegation({
    privateKey: stealthPrivateKey,
    delegation,
    delegationManager: environment.DelegationManager,
    chainId: chain,
    allowInsecureUnrestrictedDelegation: true,
  });

  const data = contracts.DelegationManager.encode.redeemDelegations({
    delegations: [[{ ...delegation, signature: delegationSignature }]],
    modes: [ExecutionMode.SingleDefault],
    executions: [
      [
        createExecution({
          target: approvalToken,
          value: BigInt(0),
          callData: encodeFunctionData({
            abi: erc20Abi,
            functionName: "approve",
            args: [sweeper, maxUint256],
          }),
        }),
      ],
    ],
  });

  logger.debug("sending delegation and approval", {
    address: stealthAccount.address,
    approvalToken,
    sweeper,
  });

  const hash = await walletClient.sendTransaction({
    to: environment.DelegationManager,
    data,
    authorizationList: [authorization],
  });
  const receipt = await publicClient.waitForTransactionReceipt({ hash });
  if (receipt.status !== "success") {
    logger.error("delegation transaction reverted", {
      address: stealthAccount.address,
      txHash: hash,
      block: receipt.blockNumber,
    });
    throw new Error("Stealth fee receiver delegation transaction failed");
  }

  logger.info("stealth fee receiver delegated and approved", {
    address: stealthAccount.address,
    approvalToken,
    txHash: hash,
    gasUsed: receipt.gasUsed,
  });
};
