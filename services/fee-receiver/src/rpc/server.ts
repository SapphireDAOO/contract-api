import type { FeeReceiverServer } from "../../generated/fee_receiver";
import { createFeeReceivers } from "./handlers";
import { unary } from "./status";

export const feeReceiverServer: FeeReceiverServer = {
  createFeeReceivers: unary(createFeeReceivers),
};
