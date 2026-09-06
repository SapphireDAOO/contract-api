import type { FeeReceiverServer } from "../../generated/fee_receiver";
import { generateaddresses, send } from "./handlers";
import { unary } from "./status";

/** The service implementation, wired to the generated FeeReceiver definition. */
export const feeReceiverServer: FeeReceiverServer = {
  generateaddresses: unary(generateaddresses),
  send: unary(send),
};
