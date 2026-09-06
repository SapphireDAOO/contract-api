import grpc from "@grpc/grpc-js";
import { FeeReceiverService } from "./generated/fee_receiver";
import { config } from "./src/config/config";
import { feeReceiverServer } from "./src/rpc/server";

const ADDRESS = process.env.GRPC_ADDRESS ?? "127.0.0.1:50051";

const main = () => {
  const { network, chain, sweeper, wrappedNative } = config();

  const server = new grpc.Server();
  server.addService(FeeReceiverService, feeReceiverServer);

  server.bindAsync(
    ADDRESS,
    grpc.ServerCredentials.createInsecure(),
    (error, port) => {
      if (error) {
        console.error("fee-receiver failed to bind", error);
        process.exit(1);
      }

      console.log(
        `fee-receiver listening on port ${port} (network ${network}, chain ${chain.id})`,
      );

      if (!sweeper) {
        console.warn(`No sweeper contract configured for network ${network}`);
      }
      if (!wrappedNative) {
        console.warn(
          `No wrapped native token configured for network ${network}`,
        );
      }
    },
  );

  const shutdown = () => server.tryShutdown(() => process.exit(0));
  process.on("SIGINT", shutdown);
  process.on("SIGTERM", shutdown);
};

main();
