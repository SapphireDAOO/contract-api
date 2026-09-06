import grpc from "@grpc/grpc-js";
import { FeeReceiverUnavailableError } from "../errors";
import { InvalidArgument } from "./validate";

const statusError = (code: grpc.status, message: string): grpc.ServiceError =>
  Object.assign(new Error(message), {
    code,
    details: message,
    metadata: new grpc.Metadata(),
  }) as grpc.ServiceError;

export const toStatusError = (error: unknown): grpc.ServiceError => {
  if (error instanceof InvalidArgument) {
    return statusError(grpc.status.INVALID_ARGUMENT, error.message);
  }

  console.error("fee-receiver rpc error", error);

  if (error instanceof FeeReceiverUnavailableError) {
    return statusError(
      grpc.status.UNAVAILABLE,
      "Fee relayer is not funded on the selected network",
    );
  }

  return statusError(grpc.status.INTERNAL, "Failed to prepare fee receiver");
};

export const unary =
  <Req, Res>(handler: (request: Req) => Promise<Res>) =>
  (
    call: grpc.ServerUnaryCall<Req, Res>,
    callback: grpc.sendUnaryData<Res>,
  ): void => {
    handler(call.request)
      .then((response) => callback(null, response))
      .catch((error: unknown) => callback(toStatusError(error), null));
  };
