/**
 * Structured logging for the service, mirroring the Go API's slog setup:
 * JSON under PRODUCTION so a log shipper can parse it, readable text
 * otherwise. LOG_LEVEL selects verbosity and defaults to info.
 *
 * Nothing here ever logs key material. Callers pass public values only —
 * stealth addresses and ephemeral public keys — never a private or spending
 * key.
 */

const LEVELS = ["debug", "info", "warn", "error"] as const;

export type Level = (typeof LEVELS)[number];

type Fields = Record<string, unknown>;

const levelFromEnv = (): Level => {
  const raw = (process.env.LOG_LEVEL ?? "").trim().toLowerCase();
  if (raw === "warning") return "warn";
  return (LEVELS as readonly string[]).includes(raw) ? (raw as Level) : "info";
};

const enabled = (level: Level): boolean =>
  LEVELS.indexOf(level) >= LEVELS.indexOf(levelFromEnv());

const asJSON = (): boolean => process.env.PRODUCTION !== undefined;

/** Errors do not survive JSON.stringify, so unwrap them to useful fields. */
const normalise = (value: unknown): unknown => {
  if (value instanceof Error) {
    return { message: value.message, name: value.name };
  }
  if (typeof value === "bigint") return value.toString();
  return value;
};

const render = (level: Level, msg: string, fields: Fields): string => {
  const normalised: Fields = {};
  for (const [key, value] of Object.entries(fields)) {
    normalised[key] = normalise(value);
  }

  if (asJSON()) {
    return JSON.stringify({
      time: new Date().toISOString(),
      level: level.toUpperCase(),
      msg,
      ...normalised,
    });
  }

  const pairs = Object.entries(normalised)
    .map(([key, value]) => {
      const text = typeof value === "string" ? value : JSON.stringify(value);
      return `${key}=${text}`;
    })
    .join(" ");

  const head = `${new Date().toISOString()} ${level.toUpperCase()} ${msg}`;
  return pairs ? `${head} ${pairs}` : head;
};

const emit = (level: Level, msg: string, fields: Fields = {}): void => {
  if (!enabled(level)) return;

  const line = render(level, msg, fields);
  if (level === "error") {
    console.error(line);
    return;
  }
  if (level === "warn") {
    console.warn(line);
    return;
  }
  console.log(line);
};

export const logger = {
  debug: (msg: string, fields?: Fields) => emit("debug", msg, fields),
  info: (msg: string, fields?: Fields) => emit("info", msg, fields),
  warn: (msg: string, fields?: Fields) => emit("warn", msg, fields),
  error: (msg: string, fields?: Fields) => emit("error", msg, fields),
};
