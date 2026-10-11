import { resolve } from "node:path";
import dotenv from "dotenv";
import { logger } from "../logger";

/**
 * The repo-root .env, where the Go API keeps its secrets too. Bun only loads a
 * .env from the working directory, and this service runs from its own, so the
 * root file has to be loaded explicitly.
 */
const DEFAULT_ENV_PATH = resolve(import.meta.dir, "../../../../.env");

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

  // dotenv never overrides a variable that is already set unless asked to.
  const { error } = dotenv.config({ path, quiet: true });
  if (error) {
    logger.warn("no .env file; relying on the ambient environment", { path });
  }
};

/** Go's os.ExpandEnv: `$VAR` and `${VAR}`, with an unset name expanding to "". */
export const expandEnv = (text: string): string =>
  text.replace(
    /\$(?:\{([A-Za-z_][A-Za-z0-9_]*)\}|([A-Za-z_][A-Za-z0-9_]*))/g,
    (_match, braced?: string, bare?: string) =>
      process.env[(braced ?? bare) as string] ?? "",
  );
