import { isAddress, maxUint216, type Address } from "viem";
import {
  InvoiceKind as InvoiceKindProto,
  ProcessorKind as ProcessorKindProto,
} from "../../generated/fee_receiver";
import { chainId } from "../config/config";
import type { InvoiceKind, ProcessorKind } from "../types";

export const MAX_FEE_RECEIVERS = 5;

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

// Empty means a native payment, not a missing field.
export const parsePaymentToken = (value: string): Address | undefined => {
  if (!value) return undefined;
  if (!isAddress(value)) invalid("Invalid payment_token");
  return value as Address;
};

export const parseInvoiceId = (value: string): bigint => {
  let invoiceId: bigint;
  try {
    invoiceId = BigInt(value);
  } catch {
    return invalid("Invalid invoice_id");
  }
  if (invoiceId < BigInt(0) || invoiceId > maxUint216)
    invalid("Invalid invoice_id");
  return invoiceId;
};
