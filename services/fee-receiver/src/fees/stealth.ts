import {
  computeStealthKey,
  generateStealthAddress,
  VALID_SCHEME_ID,
} from "@scopelift/stealth-address-sdk";
import { privateKeyToAccount } from "viem/accounts";
import type { Hex, PrivateKeyAccount } from "viem";
import { requirePrivateKey } from "../chain/keys";
import { logger } from "../logger";

export type StealthFeeReceiver = {
  stealthAccount: PrivateKeyAccount;
  stealthPrivateKey: Hex;
  ephemeralPublicKey: Hex;
};

export const restoreStealthFeeReceiver = (
  ephemeralPublicKey: Hex,
): StealthFeeReceiver => {
  const stealthPrivateKey = computeStealthKey({
    ephemeralPublicKey,
    spendingPrivateKey: requirePrivateKey("SPENDING"),
    viewingPrivateKey: requirePrivateKey("VIEWING"),
    schemeId: VALID_SCHEME_ID.SCHEME_ID_1,
  });

  return {
    stealthAccount: privateKeyToAccount(stealthPrivateKey),
    stealthPrivateKey,
    ephemeralPublicKey,
  };
};

export const generateStealthFeeReceiver = (): StealthFeeReceiver => {
  const stealthMetaAddressURI = process.env.META_STEALTH_ADDRESS;
  if (!stealthMetaAddressURI) throw new Error("Missing META_STEALTH_ADDRESS");

  const { stealthAddress, ephemeralPublicKey } = generateStealthAddress({
    stealthMetaAddressURI,
  });

  const receiver = restoreStealthFeeReceiver(ephemeralPublicKey);

  if (
    receiver.stealthAccount.address.toLowerCase() !==
    stealthAddress.toLowerCase()
  ) {
    logger.error("derived stealth key does not control the stealth address", {
      ephemeralPublicKey,
      derivedAddress: stealthAddress,
      keyAddress: receiver.stealthAccount.address,
    });
    throw new Error(
      "Computed stealth key does not control the derived stealth address",
    );
  }

  // The ephemeral public key is the only way to recover this wallet later,
  // so it is recorded here alongside the address it controls. Both are
  // public; the stealth private key is never logged.
  logger.info("stealth fee receiver created", {
    ephemeralPublicKey: receiver.ephemeralPublicKey,
    address: receiver.stealthAccount.address,
  });

  return receiver;
};
