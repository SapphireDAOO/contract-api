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
import { deriveEphemeralKey } from "../fees/derive";
import {
  generateStealthFeeReceiver,
  type StealthFeeReceiver,
} from "../fees/stealth";
import { logger } from "../logger";
import type { InvoiceKind } from "../types";
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

const inFlight = new Map<string, Promise<CreateFeeReceiversResponse>>();

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
  const processorAddress = getProcessorAddress(processor);

  const key = [
    request.chainId,
    processorAddress.toLowerCase(),
    kind,
    invoiceId,
    quantity,
    paymentToken?.toLowerCase() ?? "native",
  ].join(":");

  const pending = inFlight.get(key);
  if (pending) {
    logger.info("joining an identical fee receiver request in flight", {
      invoiceId,
      quantity,
    });
    return pending;
  }

  const run = prepareFeeReceivers(
    request.chainId,
    processorAddress,
    kind,
    invoiceId,
    quantity,
    paymentToken,
  ).finally(() => inFlight.delete(key));
  inFlight.set(key, run);
  return run;
};

const prepareFeeReceivers = async (
  chain: number,
  processorAddress: Address,
  kind: InvoiceKind,
  invoiceId: bigint,
  quantity: number,
  paymentToken?: Address,
): Promise<CreateFeeReceiversResponse> => {
  logger.info("generating fee receivers", {
    invoiceId,
    quantity,
    chainId: chain,
    paymentToken: paymentToken ?? "default",
  });

  const receivers = Array.from({ length: quantity }, (_, index) =>
    generateStealthFeeReceiver(
      deriveEphemeralKey(chain, processorAddress, kind, invoiceId, index),
    ),
  );
  await approveAll(receivers, paymentToken);

  const addresses = receivers.map((r) => r.stealthAccount.address);
  const ephemeralPublicKey = receivers.map((r) => r.ephemeralPublicKey);
  // Logged before signing, so a signer failure cannot lose keys whose gas is spent.
  logger.info("fee receivers ready", {
    invoiceId,
    addresses,
    ephemeralPublicKey,
  });

  const signature = await signFeeAuthorization(
    processorAddress,
    invoiceId,
    addresses,
    kind,
  );

  logger.info("fee receivers authorized", {
    invoiceId,
    processorAddress,
    kind,
    addresses,
  });

  return { addresses, signature, ephemeralPublicKey };
};
