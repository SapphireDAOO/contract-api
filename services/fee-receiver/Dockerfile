FROM oven/bun:1.4.0 AS deps

WORKDIR /app/services/fee-receiver
COPY services/fee-receiver/package.json services/fee-receiver/bun.lock ./
RUN bun install --frozen-lockfile --production

FROM oven/bun:1.4.0

WORKDIR /app/services/fee-receiver

COPY --from=deps /app/services/fee-receiver/node_modules ./node_modules
COPY services/fee-receiver ./

COPY config.yaml /app/config.yaml

ENV GRPC_ADDRESS=0.0.0.0:50051

USER bun

EXPOSE 50051

CMD ["bun", "run", "index.ts"]
