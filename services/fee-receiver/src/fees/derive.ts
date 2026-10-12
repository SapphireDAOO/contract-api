import {
  type Address,
  encodeAbiParameters,
  hexToBytes,
  bytesToBigInt,
  numberToBytes,
} from "viem";
import { requirePrivateKey } from "../chain/keys";
import type { InvoiceKind } from "../types";
import { createHmac } from "node:crypto";

const SECP256K1_N =
  0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141n;

/**
 * Derives the ephemeral private key for one fee receiver from RECEIVER_SEED,
 * so creating receivers is idempotent: the same invoice and index always give
 * the same key, and so the same stealth address and authorization signature.
 *
 * The message is ABI-encoded (fixed width, unambiguous) and versioned. The
 * payment token is left out on purpose: the fee authorization does not cover
 * it, so a buyer switching tokens keeps the same receivers and signature.
 *
 * A leaked seed links receivers to invoices but cannot spend from them, which
 * also needs the spending key.
 */
export const deriveEphemeralKey = (
  chainId: number,
  processor: Address,
  kind: InvoiceKind,
  invoiceId: bigint,
  index: number,
): Uint8Array => {
  const message = encodeAbiParameters(
    [
      { type: "string" },
      { type: "uint256" },
      { type: "address" },
      { type: "string" },
      { type: "uint216" },
      { type: "uint8" },
    ],
    [
      "sapphire/fee-receiver/v1",
      BigInt(chainId),
      processor,
      kind,
      invoiceId,
      index,
    ],
  );

  //   k = function(seed, chainId, processor, kind, invoiceId, index)

  const mac = createHmac(
    "sha256",
    hexToBytes(requirePrivateKey("RECEIVER_SEED")),
  )
    .update(hexToBytes(message))
    .digest();
  // Map into [1, n-1]; the bias from the reduction is ~2^-128.
  const k = (bytesToBigInt(mac) % (SECP256K1_N - 1n)) + 1n;
  return numberToBytes(k, { size: 32 });
};
