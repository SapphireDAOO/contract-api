/** Which payment processor the invoice lives on. */
export type ProcessorKind = "simple" | "intermediated";

/** Which FeeAuthorizationLib overload the processor will verify against. */
export type InvoiceKind = "single" | "meta";
