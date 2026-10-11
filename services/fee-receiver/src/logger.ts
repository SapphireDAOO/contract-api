// Never pass key material here: addresses and ephemeral public keys only.
import pino from "pino";
import pretty from "pino-pretty";

const LEVELS = ["debug", "info", "warn", "error"] as const;

export type Level = (typeof LEVELS)[number];

type Fields = Record<string, unknown>;

const levelFromEnv = (): Level => {
  const raw = (process.env.LOG_LEVEL ?? "").trim().toLowerCase();
  if (raw === "warning") return "warn";
  return (LEVELS as readonly string[]).includes(raw) ? (raw as Level) : "info";
};

const create = (): pino.Logger => {
  const options: pino.LoggerOptions = {
    level: levelFromEnv(),
    base: undefined, // no pid or hostname, matching the Go API's lines
    timestamp: pino.stdTimeFunctions.isoTime,
    // Callers pass errors as `error`, not pino's default `err`.
    serializers: { error: pino.stdSerializers.err },
  };

  if (process.env.PRODUCTION !== undefined) {
    return pino({
      ...options,
      formatters: { level: (label) => ({ level: label.toUpperCase() }) },
    });
  }

  // Synchronous, so a line logged just before process.exit is not lost.
  return pino(
    options,
    pretty({
      sync: true,
      singleLine: true,
      translateTime: "UTC:yyyy-mm-dd'T'HH:MM:ss.l'Z'",
    }),
  );
};

let instance: pino.Logger | undefined;

// Lazy, so a LOG_LEVEL from the root .env is loaded before it is read.
const base = (): pino.Logger => (instance ??= create());

export const logger = {
  debug: (msg: string, fields: Fields = {}) => base().debug(fields, msg),
  info: (msg: string, fields: Fields = {}) => base().info(fields, msg),
  warn: (msg: string, fields: Fields = {}) => base().warn(fields, msg),
  error: (msg: string, fields: Fields = {}) => base().error(fields, msg),
};
