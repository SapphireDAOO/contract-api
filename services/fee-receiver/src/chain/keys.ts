import { nonceManager, privateKeyToAccount } from "viem/accounts";
import type { Hex, PrivateKeyAccount } from "viem";

const PK_LENGTH = 66;

const normalizePrivateKey = (value: string): Hex =>
  (value.startsWith("0x") ? value : `0x${value}`) as Hex;

// Errors name only the variable, never the key.
export const requirePrivateKey = (name: string): Hex => {
  const raw = process.env[name];
  if (!raw) throw new Error(`Missing ${name}`);
  const key = normalizePrivateKey(raw);
  if (key.length !== PK_LENGTH) throw new Error(`Invalid ${name}`);
  return key;
};

export const relayerAccount = (): PrivateKeyAccount =>
  privateKeyToAccount(requirePrivateKey("SPONSOR"), { nonceManager });
