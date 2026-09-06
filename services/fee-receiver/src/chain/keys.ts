import { privateKeyToAccount } from "viem/accounts";
import type { Hex, PrivateKeyAccount } from "viem";

const PK_LENGTH = 66;

const normalizePrivateKey = (value: string): Hex =>
  (value.startsWith("0x") ? value : `0x${value}`) as Hex;

/**
 * Reads a private key from the environment. Keys are never logged or returned,
 * so the error names only the variable that is wrong.
 */
export const requirePrivateKey = (name: string): Hex => {
  const raw = process.env[name];
  if (!raw) throw new Error(`Missing ${name}`);
  const key = normalizePrivateKey(raw);
  if (key.length !== PK_LENGTH) throw new Error(`Invalid ${name}`);
  return key;
};

/**
 * The relayer: it sponsors every delegation transaction and is the fee signer
 * registered on PaymentProcessorStorage, so the same key that pays is the one
 * the processors verify against.
 */
export const relayerAccount = (): PrivateKeyAccount =>
  privateKeyToAccount(requirePrivateKey("SPONSOR"));
