/**
 * Structured logging for the service, mirroring the Go API's slog setup:
 * JSON under PRODUCTION so a log shipper can parse it, readable text
 * otherwise. LOG_LEVEL selects verbosity and defaults to info.
 *
 * Nothing here ever logs key material. Callers pass public values only —
 * stealth addresses and ephemeral public keys — never a private or spending
 * key.
 */
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
    // No pid or hostname, so a line carries the same keys as the Go API's.
    base: undefined,
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

  // In-process and synchronous, so a line logged just before process.exit
  // still reaches the terminal.
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

// Built on first use rather than at import, so a LOG_LEVEL from the root .env
// is already loaded by the time it is read.
const base = (): pino.Logger => (instance ??= create());

export const logger = {
  debug: (msg: string, fields: Fields = {}) => base().debug(fields, msg),
  info: (msg: string, fields: Fields = {}) => base().info(fields, msg),
  warn: (msg: string, fields: Fields = {}) => base().warn(fields, msg),
  error: (msg: string, fields: Fields = {}) => base().error(fields, msg),
};
