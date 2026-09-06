import { readFileSync } from "node:fs";
import { resolve } from "node:path";
import { isAddress, zeroAddress, type Address, type Chain } from "viem";
import { base, baseSepolia, hardhat } from "viem/chains";
import { expandEnv, loadRootEnv } from "./env";

type ConfigFile = {
  network?: string;
  networks?: Record<string, NetworkSection | undefined>;
};

type NetworkSection = {
  rpc?: { http?: string };
  contracts?: {
    paymentProcessor?: string;
    simplePaymentProcessor?: string;
    sweeper?: string;
    wrappedNative?: string;
  };
};

export type FeeReceiverConfig = {
  network: string;
  chain: Chain;
  rpcUrl: string;
  /** `contracts.paymentProcessor` — the intermediated processor. */
  intermediatedProcessor: Address;
  simpleProcessor: Address;
  /** Unset until the sweeper is deployed on the selected network. */
  sweeper?: Address;
  /** Unset where no wrapped native token exists, as on a local chain. */
  wrappedNative?: Address;
};

/** The viem chain each config.yaml network section describes. */
const CHAINS: Record<string, Chain> = {
  local: hardhat,
  testnet: baseSepolia,
  mainnet: base,
};

const DEFAULT_CONFIG_PATH = resolve(import.meta.dir, "../../../../config.yaml");

/**
 * Parses an address that must be present. The zero address is rejected for the
 * same reason the Go loader rejects it: it is never a real deployment.
 */
const requireAddress = (key: string, value: string | undefined): Address => {
  if (!value) throw new Error(`config: contracts.${key} is required`);
  if (!isAddress(value)) {
    throw new Error(
      `config: contracts.${key} "${value}" is not a valid address`,
    );
  }
  if (value.toLowerCase() === zeroAddress) {
    throw new Error(`config: contracts.${key} cannot be the zero address`);
  }
  return value;
};

const optionalAddress = (
  key: string,
  value: string | undefined,
): Address | undefined => {
  if (!value) return undefined;
  if (!isAddress(value)) {
    throw new Error(
      `config: contracts.${key} "${value}" is not a valid address`,
    );
  }
  return value.toLowerCase() === zeroAddress ? undefined : value;
};

const parseConfig = (parsed: ConfigFile, path: string): FeeReceiverConfig => {
  const name = process.env.NETWORK || parsed.network;
  if (!name) {
    throw new Error(
      `config ${path}: no network selected: set the network key or NETWORK`,
    );
  }

  const section = parsed.networks?.[name];
  if (!section) {
    const available = Object.keys(parsed.networks ?? {})
      .sort()
      .join(", ");
    throw new Error(
      `config ${path}: network "${name}" is not defined (available: ${available})`,
    );
  }

  const chain = CHAINS[name];
  if (!chain) {
    throw new Error(
      `config ${path}: network "${name}" has no chain mapping in this service`,
    );
  }

  const rpcUrl = section.rpc?.http;
  if (!rpcUrl)
    throw new Error(`config ${path}: networks.${name}.rpc.http is required`);

  const contracts = section.contracts ?? {};

  return {
    network: name,
    chain,
    rpcUrl,
    intermediatedProcessor: requireAddress(
      "paymentProcessor",
      contracts.paymentProcessor,
    ),
    simpleProcessor: requireAddress(
      "simplePaymentProcessor",
      contracts.simplePaymentProcessor,
    ),
    sweeper: optionalAddress("sweeper", contracts.sweeper),
    wrappedNative: optionalAddress("wrappedNative", contracts.wrappedNative),
  };
};

export const loadConfig = (
  path: string = process.env.CONFIG_PATH ?? DEFAULT_CONFIG_PATH,
): FeeReceiverConfig => {
  loadRootEnv();
  return parseConfig(
    Bun.YAML.parse(expandEnv(readFileSync(path, "utf8"))) as ConfigFile,
    path,
  );
};

let cached: FeeReceiverConfig | undefined;

/** The resolved config, loaded once on first use. */
export const config = (): FeeReceiverConfig => (cached ??= loadConfig());

/** The chain this service is configured for; requests must match it. */
export const chainId = (): number => config().chain.id;
