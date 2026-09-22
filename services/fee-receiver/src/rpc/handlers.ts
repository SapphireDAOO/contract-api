import type { Address } from "viem";
import type {
  PrepareAddressRequest,
  PrepareAddressResponse,
  VerifyAddressesRequest,
  VerifyAddressesResponse,
} from "../../generated/fee_receiver";
import {
  getProcessorAddress,
  signFeeAuthorization,
} from "../fees/authorization";
import { delegateAndApprove } from "../fees/approval";
import {
  generateStealthFeeReceiver,
  restoreStealthFeeReceiver,
  type StealthFeeReceiver,
} from "../fees/stealth";
import { logger } from "../logger";
import {
  invalid,
  parseEphemeralKeys,
  parseInvoiceId,
  parseInvoiceKind,
  parsePaymentToken,
  parseProcessor,
  parseQuantity,
  requireConfiguredChain,
} from "./validate";

const approveAll = async (
  receivers: StealthFeeReceiver[],
  paymentToken?: Address,
): Promise<void> => {
  for (const receiver of receivers) {
    await delegateAndApprove(
      receiver.stealthAccount,
      receiver.stealthPrivateKey,
      paymentToken,
    );
  }
};

export const generateaddresses = async (
  request: PrepareAddressRequest,
): Promise<PrepareAddressResponse> => {
  parseProcessor(request.processorKind);
  requireConfiguredChain(request.chainId);
  const quantity = parseQuantity(request.quantity);
  const paymentToken = parsePaymentToken(request.paymentToken);

  logger.info("generating fee receivers", {
    quantity,
    chainId: request.chainId,
    paymentToken: paymentToken ?? "default",
  });

  const receivers = Array.from({ length: quantity }, () =>
    generateStealthFeeReceiver(),
  );
  await approveAll(receivers, paymentToken);

  const ephemeralPublicKey = receivers.map((r) => r.ephemeralPublicKey);
  logger.info("fee receivers ready", {
    quantity: receivers.length,
    addresses: receivers.map((r) => r.stealthAccount.address),
    ephemeralPublicKey,
  });

  return { ephemeralPublicKey };
};

export const send = async (
  request: VerifyAddressesRequest,
): Promise<VerifyAddressesResponse> => {
  const processor = parseProcessor(request.processorKind);
  const kind = parseInvoiceKind(request.invoiceKind);
  if (kind === "meta" && processor !== "intermediated") {
    invalid("Meta invoices are intermediated only");
  }

  requireConfiguredChain(request.chainId);
  const invoiceId = parseInvoiceId(request.invoiceId);
  const keys = parseEphemeralKeys(request.ephemeralPublicKey, kind);

  let receivers: StealthFeeReceiver[];
  try {
    receivers = keys.map((key) => restoreStealthFeeReceiver(key));
  } catch (error) {
    logger.warn("unusable ephemeral public key", {
      invoiceId,
      keys: keys.length,
      error,
    });
    return invalid("Invalid ephemeral_public_key");
  }

  const feeReceivers = receivers.map((r) => r.stealthAccount.address);
  const signature = await signFeeAuthorization(
    getProcessorAddress(processor),
    invoiceId,
    feeReceivers,
    kind,
  );

  logger.info("fee receivers authorized", {
    invoiceId,
    processor,
    kind,
    addresses: feeReceivers,
  });

  return { addresses: feeReceivers, signature };
};
