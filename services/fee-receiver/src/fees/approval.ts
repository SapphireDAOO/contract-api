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
  concat,
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
import { logger } from "../logger";

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

type Clients = ReturnType<typeof getClients>;
type SmartAccountsEnvironment = ReturnType<typeof getSmartAccountsEnvironment>;
type Authorization = Awaited<
  ReturnType<PrivateKeyAccount["signAuthorization"]>
>;

const loadEnvironment = (chain: number) => {
  let environment: SmartAccountsEnvironment;
  try {
    environment = getSmartAccountsEnvironment(chain);
  } catch {
    logger.warn("no smart accounts environment; skipping 7702 delegation", {
      chainId: chain,
    });
    return undefined;
  }

  const delegatorImpl =
    environment.implementations.EIP7702StatelessDeleGatorImpl;
  if (!delegatorImpl) {
    throw new Error(
      `No EIP7702StatelessDeleGatorImpl in the smart accounts environment for chain ${chain}`,
    );
  }
  return { environment, delegatorImpl };
};

const requireSweeper = (): Address => {
  const { sweeper, network } = config();
  if (!sweeper) {
    throw new Error(`No sweeper contract configured for network ${network}`);
  }
  return sweeper;
};

const readReceiverState = async (
  publicClient: Clients["publicClient"],
  receiver: Address,
  approvalToken: Address,
  sweeper: Address,
  delegatorImpl: Address,
) => {
  const [code, allowance] = await Promise.all([
    publicClient.getCode({ address: receiver }),
    publicClient.readContract({
      address: approvalToken,
      abi: erc20Abi,
      functionName: "allowance",
      args: [receiver, sweeper],
    }),
  ]);

  return {
    delegated:
      code?.toLowerCase() ===
      concat(["0xef0100", delegatorImpl]).toLowerCase(),
    approved: allowance === maxUint256,
  };
};

const authorizeDelegator = async (
  publicClient: Clients["publicClient"],
  stealthAccount: PrivateKeyAccount,
  delegatorImpl: Address,
  chain: number,
) =>
  stealthAccount.signAuthorization({
    address: delegatorImpl,
    chainId: chain,
    nonce: await publicClient.getTransactionCount({
      address: stealthAccount.address,
    }),
  });

const encodeApproval = async (
  environment: SmartAccountsEnvironment,
  relayer: Address,
  stealthAccount: PrivateKeyAccount,
  stealthPrivateKey: Hex,
  chain: number,
  approvalToken: Address,
  sweeper: Address,
): Promise<Hex> => {
  const delegation = {
    delegate: relayer as Hex,
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

  return contracts.DelegationManager.encode.redeemDelegations({
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
};

const sendAndConfirm = async (
  { publicClient, walletClient }: Clients,
  to: Address,
  data: Hex,
  authorizationList: Authorization[] | undefined,
  receiver: Address,
) => {
  const hash = await walletClient.sendTransaction({
    to,
    data,
    authorizationList,
  });
  const receipt = await publicClient.waitForTransactionReceipt({ hash });
  if (receipt.status !== "success") {
    logger.error("delegation transaction reverted", {
      address: receiver,
      txHash: hash,
      block: receipt.blockNumber,
    });
    throw new Error("Stealth fee receiver delegation transaction failed");
  }
  return { hash, receipt };
};

export const delegateAndApprove = async (
  stealthAccount: PrivateKeyAccount,
  stealthPrivateKey: Hex,
  paymentToken?: Address,
): Promise<void> => {
  const chain = chainId();
  const setup = loadEnvironment(chain);
  if (!setup) return;
  const { environment, delegatorImpl } = setup;

  const approvalToken = resolveApprovalToken(paymentToken);
  const sweeper = requireSweeper();
  const clients = getClients();
  const receiver = stealthAccount.address;

  const { delegated, approved } = await readReceiverState(
    clients.publicClient,
    receiver,
    approvalToken,
    sweeper,
    delegatorImpl,
  );
  if (delegated && approved) {
    logger.info("stealth fee receiver already delegated and approved", {
      address: receiver,
      approvalToken,
    });
    return;
  }

  const authorizationList = delegated
    ? undefined
    : [
        await authorizeDelegator(
          clients.publicClient,
          stealthAccount,
          delegatorImpl,
          chain,
        ),
      ];
  const data = await encodeApproval(
    environment,
    clients.relayerAccount.address,
    stealthAccount,
    stealthPrivateKey,
    chain,
    approvalToken,
    sweeper,
  );

  logger.debug("sending delegation and approval", {
    address: receiver,
    approvalToken,
    sweeper,
    alreadyDelegated: delegated,
  });

  const { hash, receipt } = await sendAndConfirm(
    clients,
    environment.DelegationManager,
    data,
    authorizationList,
    receiver,
  );

  logger.info("stealth fee receiver delegated and approved", {
    address: receiver,
    approvalToken,
    txHash: hash,
    gasUsed: receipt.gasUsed,
  });
};
