import type { FeeReceiverServer } from "../../generated/fee_receiver";
import { createFeeReceivers } from "./handlers";
import { unary } from "./status";

/** The service implementation, wired to the generated FeeReceiver definition. */
export const feeReceiverServer: FeeReceiverServer = {
  createFeeReceivers: unary(createFeeReceivers),
};
