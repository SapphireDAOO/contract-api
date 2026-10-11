import type { Address } from "viem";
import type {
  CreateFeeReceiversRequest,
  CreateFeeReceiversResponse,
} from "../../generated/fee_receiver";
import {
  getProcessorAddress,
  signFeeAuthorization,
} from "../fees/authorization";
import { delegateAndApprove } from "../fees/approval";
import {
  generateStealthFeeReceiver,
  type StealthFeeReceiver,
} from "../fees/stealth";
import { logger } from "../logger";
import {
  invalid,
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

export const createFeeReceivers = async (
  request: CreateFeeReceiversRequest,
): Promise<CreateFeeReceiversResponse> => {
  const processor = parseProcessor(request.processorKind);
  const kind = parseInvoiceKind(request.invoiceKind);
  if (kind === "meta" && processor !== "intermediated") {
    invalid("Meta invoices are intermediated only");
  }

  requireConfiguredChain(request.chainId);
  const invoiceId = parseInvoiceId(request.invoiceId);
  const quantity = parseQuantity(request.quantity);
  if (kind === "single" && quantity !== 1) {
    invalid("A single invoice takes one fee receiver");
  }
  const paymentToken = parsePaymentToken(request.paymentToken);

  logger.info("generating fee receivers", {
    invoiceId,
    quantity,
    chainId: request.chainId,
    paymentToken: paymentToken ?? "default",
  });

  const receivers = Array.from({ length: quantity }, () =>
    generateStealthFeeReceiver(),
  );

  // Resolves only once every delegation and approval has been mined.
  await approveAll(receivers, paymentToken);

  const addresses = receivers.map((r) => r.stealthAccount.address);
  const ephemeralPublicKey = receivers.map((r) => r.ephemeralPublicKey);
  // Logged before signing so a signer failure still leaves the keys of
  // receivers whose gas is already spent.
  logger.info("fee receivers ready", {
    invoiceId,
    addresses,
    ephemeralPublicKey,
  });

  const signature = await signFeeAuthorization(
    getProcessorAddress(processor),
    invoiceId,
    addresses,
    kind,
  );

  logger.info("fee receivers authorized", {
    invoiceId,
    processor,
    kind,
    addresses,
  });

  return { addresses, signature, ephemeralPublicKey };
};
