import { readFileSync } from "node:fs";
import { resolve } from "node:path";

/**
 * The repo-root .env, where the Go API keeps its secrets too. Bun only loads a
 * .env from the working directory, and this service runs from its own, so the
 * root file has to be loaded explicitly.
 */
const DEFAULT_ENV_PATH = resolve(import.meta.dir, "../../../../.env");

const ENV_LINE = /^\s*(?:export\s+)?([A-Za-z_][A-Za-z0-9_]*)\s*=\s*(.*)$/;

/** Strips surrounding quotes, or a trailing comment from a bare value. */
const unquote = (value: string): string => {
  const trimmed = value.trim();
  const quote = trimmed[0];
  if (
    trimmed.length > 1 &&
    (quote === '"' || quote === "'") &&
    trimmed.endsWith(quote)
  ) {
    return trimmed.slice(1, -1);
  }
  const comment = trimmed.indexOf(" #");
  return (comment === -1 ? trimmed : trimmed.slice(0, comment)).trim();
};

let loaded = false;

/**
 * Loads the repo-root .env the way internal/server does: skipped entirely when
 * PRODUCTION is set, and never overriding a variable already in the
 * environment, so a real env var always wins over the file.
 *
 * A missing file is not fatal here — the environment may be supplied another
 * way, and config validation reports what is actually absent.
 */
export const loadRootEnv = (path: string = DEFAULT_ENV_PATH): void => {
  if (loaded) return;
  loaded = true;

  if (process.env.PRODUCTION !== undefined) return;

  let text: string;
  try {
    text = readFileSync(path, "utf8");
  } catch {
    console.warn(`No .env at ${path}; relying on the ambient environment`);
    return;
  }

  for (const line of text.split("\n")) {
    const match = ENV_LINE.exec(line);
    if (!match) continue;

    const [, key, value] = match;
    if (key && !(key in process.env)) process.env[key] = unquote(value ?? "");
  }
};

/** Go's os.ExpandEnv: `$VAR` and `${VAR}`, with an unset name expanding to "". */
export const expandEnv = (text: string): string =>
  text.replace(
    /\$(?:\{([A-Za-z_][A-Za-z0-9_]*)\}|([A-Za-z_][A-Za-z0-9_]*))/g,
    (_match, braced?: string, bare?: string) =>
      process.env[(braced ?? bare) as string] ?? "",
  );
