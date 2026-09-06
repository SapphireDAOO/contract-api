import { encodeAbiParameters, keccak256, type Address, type Hex } from "viem";
import { relayerAccount } from "../chain/keys";
import { chainId, config } from "../config/config";
import type { InvoiceKind, ProcessorKind } from "../types";

/** The processor whose storage the digest is verified against. */
export const getProcessorAddress = (processor: ProcessorKind): Address =>
  processor === "simple"
    ? config().simpleProcessor
    : config().intermediatedProcessor;

export const signFeeAuthorization = async (
  processorAddress: Address,
  invoiceId: bigint,
  feeReceivers: Address[],
  kind: InvoiceKind,
): Promise<Hex> => {
  const account = relayerAccount();
  const chain = BigInt(chainId());

  const digest =
    kind === "meta"
      ? keccak256(
          encodeAbiParameters(
            [
              { type: "address" },
              { type: "uint256" },
              { type: "uint216" },
              { type: "address[]" },
            ],
            [processorAddress, chain, invoiceId, feeReceivers],
          ),
        )
      : keccak256(
          encodeAbiParameters(
            [
              { type: "address" },
              { type: "uint256" },
              { type: "uint216" },
              { type: "address" },
            ],
            [processorAddress, chain, invoiceId, feeReceivers[0] as Address],
          ),
        );

  return account.signMessage({ message: { raw: digest } });
};
