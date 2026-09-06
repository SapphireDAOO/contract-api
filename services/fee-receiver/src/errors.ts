export class FeeReceiverUnavailableError extends Error {
  readonly code = "FEE_RELAYER_UNFUNDED";

  constructor(readonly chainId: number) {
    super(`Fee relayer has no native balance on chain ${chainId}`);
    this.name = "FeeReceiverUnavailableError";
  }
}
