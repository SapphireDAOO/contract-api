import { readFileSync } from "node:fs";
import { resolve } from "node:path";
import { isAddress, zeroAddress, type Address, type Chain } from "viem";
import { base, baseSepolia, hardhat } from "viem/chains";
import { z } from "zod";
import { expandEnv, loadRootEnv } from "./env";

// Missing keys read "is required", as they did before the schema.
const required = { error: "is required" };

const address = z
  .string(required)
  .refine((value) => isAddress(value), "is not a valid address")
  .transform((value) => value as Address);

/**
 * An address that must be present. The zero address is rejected for the same
 * reason the Go loader rejects it: it is never a real deployment.
 */
const requiredAddress = address.refine(
  (value) => value.toLowerCase() !== zeroAddress,
  "cannot be the zero address",
);

/**
 * Empty, null (a bare YAML key) and the zero address all mean not deployed on
 * this network.
 */
const optionalAddress = z
  .union([z.literal(""), address])
  .nullish()
  .transform((value) =>
    !value || value.toLowerCase() === zeroAddress ? undefined : value,
  );

const NetworkSection = z.object({
  rpc: z.object({ http: z.string(required).min(1, required) }, required),
  contracts: z.object(
    {
      paymentProcessor: requiredAddress,
      simplePaymentProcessor: requiredAddress,
      sweeper: optionalAddress,
      wrappedNative: optionalAddress,
    },
    required,
  ),
});

// Only the selected network's section is validated: the others may reference
// variables this environment does not set.
const ConfigFile = z.object({
  network: z.string().nullish(),
  networks: z.record(z.string(), z.unknown()).nullish(),
});

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

/** Every issue on one line, each named by its full key in config.yaml. */
const describe = (prefix: string, error: z.ZodError): string =>
  error.issues
    .map((issue) => `${[prefix, ...issue.path].join(".")} ${issue.message}`)
    .join("; ");

const parseConfig = (raw: unknown, path: string): FeeReceiverConfig => {
  const file = ConfigFile.safeParse(raw);
  if (!file.success) {
    throw new Error(`config ${path}: ${z.prettifyError(file.error)}`);
  }

  const name = process.env.NETWORK || file.data.network;
  if (!name) {
    throw new Error(
      `config ${path}: no network selected: set the network key or NETWORK`,
    );
  }

  const networks = file.data.networks ?? {};
  if (networks[name] == null) {
    const available = Object.keys(networks).sort().join(", ");
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

  const section = NetworkSection.safeParse(networks[name]);
  if (!section.success) {
    throw new Error(
      `config ${path}: ${describe(`networks.${name}`, section.error)}`,
    );
  }

  const { rpc, contracts } = section.data;
  return {
    network: name,
    chain,
    rpcUrl: rpc.http,
    intermediatedProcessor: contracts.paymentProcessor,
    simpleProcessor: contracts.simplePaymentProcessor,
    sweeper: contracts.sweeper,
    wrappedNative: contracts.wrappedNative,
  };
};

export const loadConfig = (
  path: string = process.env.CONFIG_PATH ?? DEFAULT_CONFIG_PATH,
): FeeReceiverConfig => {
  loadRootEnv();
  return parseConfig(
    Bun.YAML.parse(expandEnv(readFileSync(path, "utf8"))),
    path,
  );
};

let cached: FeeReceiverConfig | undefined;

/** The resolved config, loaded once on first use. */
export const config = (): FeeReceiverConfig => (cached ??= loadConfig());

/** The chain this service is configured for; requests must match it. */
export const chainId = (): number => config().chain.id;
