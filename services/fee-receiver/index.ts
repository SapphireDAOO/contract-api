import grpc from "@grpc/grpc-js";
import { FeeReceiverService } from "./generated/fee_receiver";
import { config } from "./src/config/config";
import { logger } from "./src/logger";
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
        logger.error("fee-receiver failed to bind", { address: ADDRESS, error });
        process.exit(1);
      }

      logger.info("fee-receiver listening", {
        port,
        network,
        chainId: chain.id,
      });

      if (!sweeper) {
        logger.warn("no sweeper contract configured", { network });
      }
      if (!wrappedNative) {
        logger.warn("no wrapped native token configured", { network });
      }
    },
  );

  const shutdown = () => {
    logger.info("shutdown signal received");
    server.tryShutdown(() => {
      logger.info("shutdown complete");
      process.exit(0);
    });
  };
  process.on("SIGINT", shutdown);
  process.on("SIGTERM", shutdown);
};

main();
