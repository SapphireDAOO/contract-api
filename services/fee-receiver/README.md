# fee-receiver

gRPC service that issues one-time EIP-5564 stealth fee receivers for an invoice
and signs the `FeeAuthorizationLib` authorization the payment processors verify.

It reads the repo's `config.yaml` — the same file the Go API reads — so the
processor addresses cannot drift between the two.

Clients do not reach it directly. The Go API fronts the RPC as
[`POST /v1/fee-receivers`](../../README.md#endpoint-create-fee-receivers),
and dials the address in `services.feeReceiver`. That connection is plaintext,
so this process must stay on loopback: it holds the platform's spending and
viewing keys.

## RPC

Defined in [`proto/fee_receiver.proto`](../../proto/fee_receiver.proto).

**`CreateFeeReceivers(CreateFeeReceiversRequest)`** validates the whole request
first, so a bad one never spends gas. It then derives `quantity` stealth
addresses, signs the fee authorization over them and answers straight away with
the addresses, the signature to pass on-chain as `_feeReceiver(s)` / `_data`,
and the ephemeral public keys. None of this touches the chain.

**Approvals run in the background, in parallel.** After answering, the service
upgrades every receiver at once to a MetaMask Stateless7702 delegator, hands
the relayer a delegation over it and — in the same sponsored transaction —
approves the Sweeper for max uint256 of the fee token. The relayer's nonces are
tracked locally, so the parallel transactions do not collide. A run waits for
every receiver before deciding whether it failed; a failed run is retried
after 5 s, 30 s and 2 min, skipping receivers already done, then logged as
`fee receiver approval failed, giving up` with each receiver's error. A successful response therefore means the receivers can be paid, not that
they are ready to sweep: paying a receiver needs no approval, and a fee that
arrives first simply waits for it. Calling again for the same invoice finishes
whatever a failed or interrupted run left behind.

Callers should store the keys: with the `SPENDING` and `VIEWING` keys they
recover each receiver's private key. Receivers can also be re-derived from
`RECEIVER_SEED` (see below). At most 5 receivers per call, each costing one
sponsored transaction.

**Retries are idempotent.** Each receiver's ephemeral key is
HMAC-SHA256(`RECEIVER_SEED`, chain id, processor address, invoice kind,
invoice id, index), not random, so repeating a request returns the same
addresses, keys and signature. Before sending, each receiver's on-chain state
is checked: one already delegated and approved costs no transaction, and one
delegated but missing the token's approval gets only the approval. A retry
after a partial failure therefore sends only what is left. Requests that
overlap on a receiver share one run for it, and a receiver this process has
already set up is not read from the chain again. The payment token is not part of the
derivation, because the signature does not cover it: switching tokens keeps the
receivers and adds an approval. Rotating `RECEIVER_SEED` gives unpaid invoices
new receivers on their next request.

A meta invoice needs one receiver per sub-invoice, index-aligned with its
`subInvoiceIds`, all covered by the single signature over the whole array. Pass
`INVOICE_KIND_META` explicitly: a meta invoice holding one sub-invoice still
signs the array form, so the kind is never inferred from `quantity`. A single
invoice takes `quantity` 1.

## Layout

```
index.ts              bootstrap: resolve config, bind, serve
generated/            ts-proto stubs — do not edit
src/
  config/env.ts       repo-root .env + ${VAR} expansion
  config/config.ts    config.yaml -> the resolved network
  chain/keys.ts       private keys from the environment
  chain/clients.ts    viem public/wallet clients
  fees/stealth.ts     EIP-5564 derive and restore
  fees/approval.ts    7702 delegation + sweeper approval
  fees/authorization.ts  the fee authorization digest and signature
  rpc/validate.ts     request parsing, all rejections in one place
  rpc/handlers.ts     the RPC
  rpc/status.ts       error -> gRPC status
  rpc/server.ts       the assembled service
```

## Configuration

The network comes from `NETWORK` (or `config.yaml`'s own `network` key), and
these keys are read from the selected section:

| Key | Use |
| --- | --- |
| `rpc.http` | Chain the relayer sends through |
| `contracts.paymentProcessor` | Intermediated processor, for the digest |
| `contracts.simplePaymentProcessor` | Simple processor, for the digest |
| `contracts.sweeper` | Approved for max uint256 of the fee token |
| `contracts.wrappedNative` | Approved instead when the payment is native |

`sweeper` and `wrappedNative` are new and read only here; the Go API ignores
them. Both may be empty — the service starts and warns, then fails the calls
that need them.

Secrets stay in the environment. The service loads the **repo-root `.env`** — the
same file the Go API reads — because Bun only auto-loads a `.env` from the
working directory, and this service runs from its own. A variable already set in
the environment always wins over the file, and the file is skipped entirely when
`PRODUCTION` is set.

| Variable | Use |
| --- | --- |
| `SPONSOR` | Relayer key: sponsors the delegations and signs the fee authorization |
| `SPENDING` / `VIEWING` | EIP-5564 keys the stealth addresses derive from |
| `META_STEALTH_ADDRESS` | Stealth meta-address URI new receivers derive from |
| `RECEIVER_SEED` | 32-byte hex secret the receivers' ephemeral keys are derived from, which makes retries idempotent. Keep it apart from the other keys; a leak links receivers to invoices but cannot spend from them |
| `SWEEPER_CONTRACT` | Expanded into `contracts.sweeper` |
| `GRPC_ADDRESS` | Listen address, default `127.0.0.1:50051` |
| `CONFIG_PATH` | Override the path to `config.yaml` |

## Running

```bash
bun install
bun run start
```

## Regenerating the stubs

After editing the proto:

```bash
bun run gen
```
