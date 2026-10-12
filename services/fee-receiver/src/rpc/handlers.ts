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
  const results = await Promise.allSettled(
    receivers.map((receiver) =>
      delegateAndApprove(
        receiver.stealthAccount,
        receiver.stealthPrivateKey,
        paymentToken,
      ),
    ),
  );

  const failures = results.filter(
    (result): result is PromiseRejectedResult => result.status === "rejected",
  );
  if (failures.length > 0) {
    throw new AggregateError(
      failures.map((failure) => failure.reason),
      `${failures.length} of ${receivers.length} fee receivers failed to approve`,
    );
  }
};

const RETRY_DELAYS_MS = [5_000, 30_000, 120_000];

const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

const approveInBackground = (
  invoiceId: bigint,
  receivers: StealthFeeReceiver[],
  paymentToken?: Address,
): void => {
  const addresses = receivers.map((r) => r.stealthAccount.address);

  void (async () => {
    for (let attempt = 0; ; attempt++) {
      try {
        await approveAll(receivers, paymentToken);
        logger.info("fee receivers approved", { invoiceId, addresses });
        return;
      } catch (error) {
        const retryIn = RETRY_DELAYS_MS[attempt];
        if (retryIn === undefined) {
          logger.error("fee receiver approval failed, giving up", {
            invoiceId,
            addresses,
            attempts: attempt + 1,
            error,
          });
          return;
        }
        logger.warn("fee receiver approval failed, retrying", {
          invoiceId,
          addresses,
          attempt: attempt + 1,
          retryInMs: retryIn,
          error,
        });
        await sleep(retryIn);
      }
    }
  })();
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
  const processorAddress = getProcessorAddress(processor);

  return prepareFeeReceivers(
    request.chainId,
    processorAddress,
    kind,
    invoiceId,
    quantity,
    paymentToken,
  );
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
  const addresses = receivers.map((r) => r.stealthAccount.address);
  const ephemeralPublicKey = receivers.map((r) => r.ephemeralPublicKey);

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
    ephemeralPublicKey,
  });

  approveInBackground(invoiceId, receivers, paymentToken);

  return { addresses, signature, ephemeralPublicKey };
};
