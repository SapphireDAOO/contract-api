import { resolve } from "node:path";
import dotenv from "dotenv";
import { logger } from "../logger";

// Bun only auto-loads a .env from the working directory, not the repo root.
const DEFAULT_ENV_PATH = resolve(import.meta.dir, "../../../../.env");

let loaded = false;

export const loadRootEnv = (path: string = DEFAULT_ENV_PATH): void => {
  if (loaded) return;
  loaded = true;

  if (process.env.PRODUCTION !== undefined) return;

  // dotenv leaves variables that are already set untouched.
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
