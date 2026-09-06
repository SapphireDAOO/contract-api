import type { Address, Hex } from "viem";
import {
  InvoiceKind as InvoiceKindProto,
  ProcessorKind as ProcessorKindProto,
} from "../../generated/fee_receiver";
import { chainId } from "../config/config";
import type { InvoiceKind, ProcessorKind } from "../types";

const ADDRESS_PATTERN = /^0x[0-9a-fA-F]{40}$/;
// Compressed or uncompressed secp256k1 point, as the SDK may return either.
const EPHEMERAL_KEY_PATTERN = /^0x[0-9a-fA-F]{66,130}$/;
const MAX_UINT216 = (BigInt(1) << BigInt(216)) - BigInt(1);

export const MAX_FEE_RECEIVERS = 5;

/** A rejected request, surfaced to the caller as INVALID_ARGUMENT. */
export class InvalidArgument extends Error {}

export const invalid = (message: string): never => {
  throw new InvalidArgument(message);
};

export const parseProcessor = (value: ProcessorKindProto): ProcessorKind => {
  switch (value) {
    case ProcessorKindProto.PROCESSOR_KIND_SIMPLE:
      return "simple";
    case ProcessorKindProto.PROCESSOR_KIND_INTERMEDIATED:
      return "intermediated";
    default:
      return invalid("Invalid processor_kind");
  }
};

export const parseInvoiceKind = (value: InvoiceKindProto): InvoiceKind => {
  switch (value) {
    case InvoiceKindProto.INVOICE_KIND_SINGLE:
      return "single";
    case InvoiceKindProto.INVOICE_KIND_META:
      return "meta";
    default:
      return invalid("Invalid invoice_kind");
  }
};

/**
 * The service is configured for exactly one network, so a request naming a
 * different chain is refused rather than quietly signed for the configured one.
 */
export const requireConfiguredChain = (value: number): void => {
  if (value !== chainId()) {
    invalid(
      `Unsupported chain ${value}: this service serves chain ${chainId()}`,
    );
  }
};

export const parseQuantity = (value: number): number => {
  if (!Number.isInteger(value) || value < 1 || value > MAX_FEE_RECEIVERS) {
    invalid(`quantity must be between 1 and ${MAX_FEE_RECEIVERS}`);
  }
  return value;
};

/** An empty token means the payment is native, not that the field is missing. */
export const parsePaymentToken = (value: string): Address | undefined => {
  if (!value) return undefined;
  if (!ADDRESS_PATTERN.test(value)) invalid("Invalid payment_token");
  return value as Address;
};

export const parseInvoiceId = (value: string): bigint => {
  let invoiceId: bigint;
  try {
    invoiceId = BigInt(value);
  } catch {
    return invalid("Invalid invoice_id");
  }
  if (invoiceId < BigInt(0) || invoiceId > MAX_UINT216)
    invalid("Invalid invoice_id");
  return invoiceId;
};

export const parseEphemeralKeys = (
  keys: string[],
  kind: InvoiceKind,
): Hex[] => {
  if (
    !Array.isArray(keys) ||
    keys.length < 1 ||
    keys.length > MAX_FEE_RECEIVERS ||
    !keys.every(
      (key) => typeof key === "string" && EPHEMERAL_KEY_PATTERN.test(key),
    )
  ) {
    invalid("Invalid ephemeral_public_key");
  }
  if (kind === "single" && keys.length !== 1) {
    invalid("A single invoice takes one fee receiver");
  }
  return keys as Hex[];
};
