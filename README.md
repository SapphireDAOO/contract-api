# Sapphire Contract API – REST Endpoints

This API provides HTTP endpoints for interacting with the Sapphire DAO's `IntermediatedPaymentProcessor`, `SimplePaymentProcessor` and `Notes` smart contracts on **Base Sepolia**. It supports invoice creation, cancellation, refunding, dispute creation, dispute resolution, fund release, and invoice notes.

Contract addresses and endpoints are not compiled in — they come from [`config.yaml`](config.yaml), which holds one section per network (`local`, `testnet`, `mainnet`). See [Configuration](#configuration).

Platform fees are collected through one-time stealth addresses issued by the [fee-receiver sidecar](services/fee-receiver/README.md), a separate process this API calls over gRPC. It reads the same `config.yaml`, so contract addresses cannot drift between the two.

**Base URL**: `https://sapphiredaotesting.com/`

## Endpoints

| Method | Path                                            | Description                          |
| :----: | :---------------------------------------------: | :----------------------------------: |
| GET    | [`/`](#endpoint-health)                          | Health check                         |
| POST   | [`/v1/invoices`](#endpoint-create-invoices)      | Create one or more invoices          |
| GET    | [`/v1/invoices/{invoiceId}`](#endpoint-get-invoice) | Read invoice data from the subgraph |
| POST   | [`/v1/invoices/{invoiceId}/release`](#endpoint-release) | Release escrowed funds       |
| POST   | [`/v1/invoices/{invoiceId}/cancel`](#endpoint-cancel) | Cancel an unpaid invoice       |
| POST   | [`/v1/invoices/{invoiceId}/refund`](#endpoint-refund) | Refund a paid invoice          |
| POST   | [`/v1/invoices/{invoiceId}/disputes`](#endpoint-create-dispute) | Open a dispute       |
| POST   | [`/v1/invoices/{invoiceId}/disputes/resolution`](#endpoint-resolve-dispute) | Resolve a dispute |
| GET    | [`/v1/settlements/status`](#endpoint-settlement-status) | Simple processor settlement window |
| GET    | [`/v1/exchangeRate`](#endpoint-exchange-rates)    | How much of a token one USD buys     |
| POST   | [`/v1/fee-receivers`](#endpoint-create-fee-receivers) | Derive, approve and authorize stealth fee receivers |
| GET    | [`/v1/auth/nonce`](#endpoint-sign-in)            | Issue a single-use sign-in nonce     |
| POST   | [`/v1/auth/verify`](#endpoint-sign-in)           | Sign in with Ethereum; returns a session token |
| POST   | [`/v1/notes`](#endpoint-write-note)              | Write an encrypted note to an invoice |
| POST   | [`/v1/notes/open`](#endpoint-open-note)          | Mark a note as opened                |

The invoice id is a path segment on every invoice operation, so the request body carries only what is specific to that operation. `release`, `cancel` and `disputes` take no body at all.

The invoice endpoints that send transactions (create, release, cancel, refund and both dispute endpoints) require an `X-API-KEY` header, enforced by `AccessControlMiddleWare`. The read endpoints, the [fee receiver](#endpoint-create-fee-receivers) endpoint and [sign-in](#endpoint-sign-in) need no key. The notes endpoints act for a user, so instead of a key they require a session token from sign-in.

---

### Endpoint: Health

- **Method**: GET `/`

Returns `200` with the current time. Any other unrouted path returns `404`.

```json
{ "status": "ok", "time": "2026-09-05T10:48:12Z" }
```

---

### Endpoint: Create invoices

- **Method**: POST `/v1/invoices`
- **Description**: Creates one or more on-chain invoices using the `IntermediatedPaymentProcessor` contract's `createSingleInvoice` (for one invoice) or `createMetaInvoice` (for multiple) functions.

#### **Request Body**

```json
{
  "invoices": [
    {
      "orderId": "550e8400-e29b-41d4-a716-446655440000",
      "seller": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
      "price": 8680000000,
      "escrowHoldPeriod": 604800,
      "currency": "USD"
    }
  ],
  "paymentTokens": ["ETH", "USDC"]
}
```

`invoices` holds one entry per invoice; one creates a single invoice, several create a meta invoice. `paymentTokens` sits beside it and applies to every invoice in the batch.

#### Field Details

| Field              | Type     | Required | Description                                                                              |
| :----------------: | :------: | :------: | :--------------------------------------------------------------------------------------: |
| `invoices`         | object[] | ✅       | The invoices to create. At least one.                                                    |
| `orderId`          | string   | ✅       | Unique client-side identifier for the invoice (e.g., a UUID or any string).              |
| `seller`           | string   | ✅       | Ethereum address of the seller. Must not be the zero address.                            |
| `price`            | number   | ✅       | Invoice price in cents; scaled on the server using the `currency` precision.              |
| `escrowHoldPeriod` | number   | ✅       | Duration in seconds for holding funds in escrow (e.g., `604800` = 7 days).                |
| `currency`         | string   | ✅       | Pricing currency; sets the decimal precision applied to `price` (`USD` = 8 decimals).     |
| `paymentTokens`    | string[] | ✅       | Token **symbols** the buyer may pay with, e.g. `["ETH", "USDC"]`. At least one. Applies to every invoice in the batch, and sits beside `invoices` rather than inside it. |

**About `paymentTokens`**: callers name tokens by **symbol**, not address. Each symbol is resolved to the address deployed on the selected network using the `tokens` table in [`config.yaml`](config.yaml), so the same request body works against local, testnet and mainnet. Matching is case-insensitive (`usdc` resolves `USDC`).

The contract's `InvoiceCreationParam` takes an `address[]` and reverts with `NoPaymentTokens` on an empty list, so an empty or missing list is rejected with `400`. An unknown symbol is rejected with the list of configured ones:

```json
{ "error": "invoice 0: unknown payment token \"DOGE\" (known: ETH, USDC, wBTC)", "reason": "" }
```

`ETH` maps to the zero address, which is how the contracts denote the native token.

#### **Response**

**Success (200)** — the same shape for one invoice or many:

```json
{
  "url": "https://sapphire-dao-website.pages.dev/checkout/?id=kWMRqBU-7H64tqp04CM5IfzKRMP1DbH2Ytg5",
  "invoices": {
    "550e8400-e29b-41d4-a716-446655440000": {
      "seller": "0x329C3E1bEa46Abc22F307eE30Cbb522B82Fe7082",
      "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889"
    }
  }
}
```

**Notes**:

- The `url` is `urls.checkout` for the selected network followed by the invoice id encoded as unpadded URL-safe base64 of its big-endian bytes — the same encoding the website uses, so it can decode the id straight from the link. For a single invoice that is the invoice's own id. For several it is the **meta-invoice** id, prefixed with `mt-` so the two kinds of identifier cannot be confused:

  ```
  single: .../checkout/?id=dJNOQid37cGYw04ceDk1kxjAME8ElIMK9yLm
  meta:   .../checkout/?id=mt-dJNOQid37cGYw04ceDk1kxjAME8ElIMK9yLm
  ```
- `price` is converted to token amounts using Chainlink price feeds via the contract's `getTokenValueFromUsd` function.
- A single invoice triggers `createSingleInvoice`, emitting `InvoiceCreated`. Multiple invoices trigger `createMetaInvoice`, emitting `MetaInvoiceCreated`.
- Only the intermediated platform operator (retrieved via `PaymentProcessorStorage.GetIntermediatedPlatformsOperator`) can call these functions.
- The client-provided `orderId` is hashed to a `uint216` by `invoice.InvoiceIDToUint216` for on-chain storage, producing the numeric `invoiceId` in the response. `invoices` is keyed by the `orderId`. That id is the `{invoiceId}` every other endpoint takes.

**Error Responses**:

**Error (400)** — malformed JSON, or failed validation such as an invalid seller address, a non-positive price, or a missing `paymentTokens` entry:

```json
{
  "error": "invoice 0: at least one payment token is required",
  "reason": "<validation error>"
}
```

**Error (500)** — fetching the operator address failed, or the transaction reverted:

```json
{
  "error": "error creating invoice",
  "reason": "An invoice with this identifier already exists."
}
```

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/invoices \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "invoices": [
    {
      "orderId": "550e8400-e29b-41d4-a716-446655440000",
      "seller": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
      "price": 8680000000,
      "escrowHoldPeriod": 604800,
      "currency": "USD"
    }
  ],
  "paymentTokens": ["ETH", "USDC"]
}'
```

---

### Endpoint: Get invoice

- **Method**: GET `/v1/invoices/{invoiceId}`
- **Description**: Reads invoice data from the subgraph (`urls.subgraph`) rather than from the chain directly.

**Success (200)**: the invoice record as stored in the subgraph.

**Error (500)**:

```json
{
  "error": "failed to fetch invoice data",
  "reason": "<subgraph error message>"
}
```

**Example**:

```bash
curl https://sapphiredaotesting.com/v1/invoices/59808737901387817475691215581034097896123425895641016234844280889 \
-H "X-API-KEY: YOUR_API_KEY_HERE"
```

---

### Endpoint: Release

- **Method**: POST `/v1/invoices/{invoiceId}/release`
- **Description**: Releases escrow funds for an invoice using the contract's `release` function, distributing funds to the seller and platform. **No request body.**

**Success (200)**:

```json
{
  "status": "success",
  "transactionUrl": "https://sepolia.basescan.org/tx/0x123456..."
}
```

**Error (400)**: `invoiceId` in the path is not a base-10 integer.

**Error (500)**:

```json
{
  "error": "Error sending transaction",
  "reason": "The invoice is not in a valid state for this action."
}
```

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/invoices/59808737901387817475691215581034097896123425895641016234844280889/release \
-H "X-API-KEY: YOUR_API_KEY_HERE"
```

---

### Endpoint: Cancel

- **Method**: POST `/v1/invoices/{invoiceId}/cancel`
- **Description**: Cancels an invoice using the contract's `cancelInvoice` function, setting its state to `CANCELED`. Can only be called before payment. **No request body.**

**Success (200)**:

```json
{
  "status": "success",
  "transactionUrl": "https://sepolia.basescan.org/tx/0x123456..."
}
```

**Error (500)**: transaction reverted (e.g., `InvalidInvoiceState`, `NotAuthorized`).

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/invoices/59808737901387817475691215581034097896123425895641016234844280889/cancel \
-H "X-API-KEY: YOUR_API_KEY_HERE"
```

---

### Endpoint: Refund

- **Method**: POST `/v1/invoices/{invoiceId}/refund`
- **Description**: Issues a refund for an invoice using the contract's `refund` function, withdrawing funds from escrow to the buyer.

#### **Request Body**

```json
{ "refundShare": "5000" }
```

| Field         | Type   | Required | Description                                                            |
| :-----------: | :----: | :------: | :--------------------------------------------------------------------: |
| `refundShare` | string | ✅       | Refund share in basis points (e.g., `"10000"` = 100%, `"5000"` = 50%). |

**Success (200)**:

```json
{
  "status": "success",
  "transactionUrl": "https://sepolia.basescan.org/tx/0x123456..."
}
```

**Error (400)**:

```json
{ "error": "refundShare is required", "reason": "" }
```

- Also returned when `refundShare` is zero (`"share can not be zero"`).

**Error (500)**: transaction reverted (e.g., `InsufficientBalance`, `InvalidInvoiceState`).

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/invoices/59808737901387817475691215581034097896123425895641016234844280889/refund \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{ "refundShare": "5000" }'
```

---

### Endpoint: Create dispute

- **Method**: POST `/v1/invoices/{invoiceId}/disputes`
- **Description**: Opens a dispute for an invoice using the contract's `createDispute` function, setting the invoice state to `DISPUTED`. **No request body.**

**Success (200)**:

```json
{
  "status": "success",
  "transactionUrl": "https://sepolia.basescan.org/tx/0x123456..."
}
```

**Error (500)**: fetching the operator address failed, or the transaction reverted.

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/invoices/59808737901387817475691215581034097896123425895641016234844280889/disputes \
-H "X-API-KEY: YOUR_API_KEY_HERE"
```

---

### Endpoint: Resolve dispute

- **Method**: POST `/v1/invoices/{invoiceId}/disputes/resolution`
- **Description**: Resolves an open dispute using the contract's `resolveDispute` or `handleDispute` functions, updating the invoice state and distributing funds if applicable.

#### **Request Body**

```json
{
  "resolution": 2,
  "sellerShare": "9000"
}
```

| Field         | Type    | Required                                    | Description                                                              |
| :-----------: | :-----: | :-----------------------------------------: | :----------------------------------------------------------------------: |
| `resolution`  | integer | ✅                                          | Enum value specifying the action type (see MarketplaceAction below).     |
| `sellerShare` | string  | ❌ Only if `resolution = 2` (SettleDispute) | Seller's share in basis points (e.g., `"10000"` = 100%, `"9000"` = 90%). |

#### MarketplaceAction Enum (`resolution`)

| Value | Name           | Description                                                                                            | Contract Function                        |
| :---: | :------------: | :----------------------------------------------------------------------------------------------------: | :--------------------------------------: |
| `1`   | ResolveDispute | Both buyer and seller agree to dismiss the dispute, with escrow allocation unchanged (fully to seller) | `resolveDispute`                         |
| `2`   | SettleDispute  | Dispute resolved by an arbitrator, with `sellerShare` to the seller and the rest to the buyer          | `handleDispute` with `DISPUTE_SETTLED`   |
| `3`   | DismissDispute | Arbitrator dismisses the dispute, leaving escrow allocation unchanged (fully to seller)                | `handleDispute` with `DISPUTE_DISMISSED` |

**Notes**:

- The contract validates that the invoice is in the `DISPUTED` state and that `sellerShare` does not exceed 10,000 basis points.
- For `SettleDispute`, funds are distributed with platform fees applied, emitting `DisputeSettled`.
- For `DismissDispute`, emits `DisputeDismissed` without fund distribution.
- For `ResolveDispute`, emits `DisputeResolved` and sets the state to `DISPUTE_RESOLVED`.

**Success (200)**:

```json
{
  "status": "success",
  "transactionUrl": "https://sepolia.basescan.org/tx/0x123456..."
}
```

**Error (500)**: transaction reverted (e.g., `InvalidDisputeResolution`, `InvalidInvoiceState`).

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/invoices/59808737901387817475691215581034097896123425895641016234844280889/disputes/resolution \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{ "resolution": 2, "sellerShare": "9000" }'
```

---

### Endpoint: Settlement status

- **Method**: GET `/v1/settlements/status`
- **Description**: Reports whether the `SimplePaymentProcessor` settlement window is still open.

- **200** with an empty body: the window is open.
- **400**: the window has passed.

```json
{ "error": "settlement time passed", "reason": "settlement window has expired" }
```

---

### Endpoint: Exchange rates

- **Method**: GET `/v1/exchangeRate?From=USD&to=wBTC&to=ETH`
- **Description**: Reports how much of each requested token one USD buys, inverting the `OracleManager` contract's `getUsdPerToken`. The oracle address comes from `contracts.oracleManager` for the selected network.

| Query  | Required | Description                                                                                     |
| :----: | :------: | :---------------------------------------------------------------------------------------------: |
| `from` | ❌       | Must be `USD`, the only currency the oracle prices against. Defaults to `USD`. `From` also works. |
| `to`   | ✅       | A token symbol from the network's `tokens` table. Repeat the parameter for several.              |

Each token is its own `to` parameter: `to=wBTC&to=ETH`. A comma-separated list is not split.

**Success (200)** — the rate is *from* USD, so each value is how much of that token one USD buys:

```json
{ "from": "USD", "to": { "ETH": 0.000510204081632653, "wBTC": 0.00001111 } }
```

- Values are JSON numbers carrying the full precision of the token's own decimals, so an 18-decimal token keeps all 18.
- A token the oracle cannot price falls back to a rate of `1`, treating it as a dollar-pegged token. The response does not distinguish that from a quoted rate.
- The conversion truncates, as any fixed-point division must: at $90,000 wBTC resolves to `0.00001111`, not a repeating `0.0000111111…`, because the token has 8 decimals.

**Error (400)** — an unknown symbol, a `from` other than `USD`, or a missing `to`:

```json
{ "error": "unknown token BTC (known: ETH, USDC, wBTC)", "reason": "unknown token BTC" }
```

**Error (502)** — the oracle call failed for a reason other than a missing feed (a stale price, or the sequencer being down). **503** — `contracts.oracleManager` is not configured for this network, so rates are disabled.

**Example**:

```bash
curl "https://sapphiredaotesting.com/v1/exchangeRate?From=USD&to=ETH&to=wBTC" \
-H "X-API-KEY: YOUR_API_KEY_HERE"
```

---

### Endpoint: Create fee receivers

- **Method**: POST `/v1/fee-receivers`
- **Description**: Derives one-time [EIP-5564](https://eips.ethereum.org/EIPS/eip-5564) stealth addresses to collect platform fees for an invoice, upgrades each to an EIP-7702 delegator, approves the sweeper to move the fee token out of them later, and returns the addresses with the fee signer's signature to pass on-chain as the processor's `_feeReceiver(s)` / `_data` arguments.

The work is done by the [fee-receiver sidecar](services/fee-receiver/README.md), reached over gRPC at `services.feeReceiver`. This API forwards the request and translates the sidecar's gRPC status into an HTTP one.

**No `X-API-KEY`.** This endpoint is unauthenticated, unlike the rest of the API. Each call sends a relayer-sponsored transaction per receiver, so whatever fronts the API is what limits who can spend that gas.

**The response does not wait for the approvals.** The sidecar derives the receivers, signs over them and answers at once; each receiver's delegation and approval is then sent in the background, retried after 5 s, 30 s and 2 min, and logged by the sidecar if it still fails. A `200` therefore means the receivers can be paid — paying needs no approval — not that they are ready to sweep. Receivers are derived from `RECEIVER_SEED`, so calling again for the same invoice returns the same addresses and finishes any approval a failed or interrupted run left behind.

#### **Request Body**

```json
{
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "processor": "intermediated",
  "kind": "meta",
  "quantity": 2,
  "paymentToken": "USDC"
}
```

| Field          | Type    | Required | Description                                                                                     |
| :------------: | :-----: | :------: | :---------------------------------------------------------------------------------------------: |
| `invoiceId`    | string  | ✅       | The on-chain invoice id, base-10. For `kind: "meta"`, the meta-invoice id.                       |
| `processor`    | string  | ✅       | `simple` or `intermediated`. Must match the processor the invoice lives on.                      |
| `kind`         | string  | ❌       | `single` (default) or `meta`. `meta` requires `processor: "intermediated"`.                      |
| `quantity`     | integer | ❌       | How many receivers to derive, `1`–`5`. Defaults to `1`. A meta invoice needs one per sub-invoice; a `single` invoice takes exactly one. |
| `paymentToken` | string  | ❌       | Token **symbol** the fee is collected in, e.g. `"USDC"`. Omit for a native-token payment.         |

**About `kind`**: it is never inferred from `quantity`. A meta invoice signs one digest over the whole address array, and a meta invoice holding a single sub-invoice still needs that array form — so a meta invoice must say so explicitly.

**About `paymentToken`**: as with `paymentTokens` on invoice creation, callers name the token by symbol and it is resolved through the `tokens` table for the selected network. Omitting it — or naming `ETH`, which maps to the zero address — means the payment is native, and the sidecar approves the chain's wrapped native token (`contracts.wrappedNative`) instead.

#### **Response**

**Success (200)** — `feeReceivers` is index-aligned with the meta-invoice's `subInvoiceIds`, and the one `signature` covers the whole array. `ephemeralPublicKeys` holds one key per receiver in the same order; they are the only way to recover a receiver's key, so store them.

```json
{
  "status": "success",
  "feeReceivers": ["0x9ba1...", "0x4cd2..."],
  "signature": "0x7f3e...",
  "ephemeralPublicKeys": ["0x02f7a1c3...", "0x03b8d2e4..."]
}
```

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/fee-receivers \
-H "Content-Type: application/json" \
-d '{
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "processor": "intermediated",
  "kind": "single",
  "paymentToken": "USDC"
}'
```

#### Errors

| Status | Meaning                                                                                                         |
| :----: | :---------------------------------------------------------------------------------------------------------------: |
| `400`  | Rejected by this API (unknown `processor`, `kind` or token symbol; a non-integer `invoiceId`; a meta invoice on the simple processor) or by the sidecar (`quantity` out of range, a meta invoice on a simple processor, a `single` invoice with `quantity` above 1). Nothing is sent on-chain. |
| `409`  | The invoice is not in a state that takes fee receivers: a simple invoice must be `PAID`, an intermediated invoice `CREATED`, and a meta invoice needs at least one `PAID` sub-invoice. An id that was never created is refused the same way. Checked before the sidecar is called, so nothing is sent on-chain. |
| `502`  | The sidecar failed internally — a chain error, a reverted approval, or one of its keys is unset. Its own message is generic by design. Also returned when this API cannot read the invoice's state from the chain. |
| `503`  | `services.feeReceiver` is not configured for this network, or the relayer has no native balance to sponsor the delegations. |
| `504`  | The sidecar did not answer within its 25s call budget.                                                            |

A rejection from the sidecar keeps the sidecar's own wording in `reason`:

```json
{ "error": "error creating fee receivers", "reason": "quantity must be between 1 and 5" }
```

---

### Endpoint: Sign in <a id="endpoint-sign-in"></a>

- **Method**: GET `/v1/auth/nonce`, then POST `/v1/auth/verify`
- **Description**: Signs a wallet in with [Sign-In with Ethereum](https://eips.ethereum.org/EIPS/eip-4361) and returns a session token. The notes endpoints take the author from that session, so a caller can only write or open notes as the account it signed in with. The user signs once per session, with no gas.

**Flow**:

1. `GET /v1/auth/nonce` returns a single-use nonce, valid for 5 minutes.
2. The website builds an EIP-4361 message with that nonce (viem's `createSiweMessage` does this) and the user signs it with `personal_sign`.
3. `POST /v1/auth/verify` with the message and signature returns a token valid for 12 hours.
4. Send the token as `Authorization: Bearer <token>` on `/v1/notes` and `/v1/notes/open`.

The message must name one of the network's `auth.domains` on its first line and the network's chain id, and must still be within any `Expiration Time` it sets. Each nonce works once.

**One sign-in covers every note.** The user signs only in step 2. After that, the same token is sent on every `/v1/notes` and `/v1/notes/open` call until `expiresAt`, any number of times; the API checks the token itself and never asks the wallet again. Keep the token in `sessionStorage` or `localStorage` so a page reload does not lose it.

Sign in again when:

- a notes call returns `401` (`Session is invalid or expired`) — the token has passed `expiresAt`, or `SESSION_SECRET` was changed;
- the token was lost, e.g. it was held only in memory and the page reloaded;
- the user switches wallet account — the token belongs to the address that signed in, so notes sent with it are written as that address. Sign in again for the new account and replace the stored token.

**Nonce response (200)**:

```json
{ "success": true, "nonce": "3f9a...c21e", "expiresAt": "2026-10-12T12:05:00Z" }
```

**Verify request**:

```json
{
  "message": "app.example.com wants you to sign in with your Ethereum account:\n0x329C...7082\n\nSign in to write and read invoice notes.\n\nURI: https://app.example.com\nVersion: 1\nChain ID: 84532\nNonce: 3f9a...c21e\nIssued At: 2026-10-12T12:00:00Z",
  "signature": "0x5c1b...1b"
}
```

**Verify response (200)**:

```json
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "address": "0x329C3E1bEa46Abc22F307eE30Cbb522B82Fe7082",
  "expiresAt": "2026-10-13T00:00:00Z"
}
```

**Errors**: `400` for a malformed body or a message that is not EIP-4361; `401` with the reason for another domain or chain, an expired message, a nonce that is unknown, expired or already used, or a signature that does not match the message's address; `503` when sign-in is not configured (no `SESSION_SECRET` or `auth.domains`).

**Notes**:

- Only signatures from ordinary wallets (EOAs) are accepted. Smart-contract wallets that sign through ERC-1271, such as Coinbase Smart Wallet, cannot sign in yet.
- Sessions are signed tokens, not stored server-side, so one cannot be revoked before it expires; rotating `SESSION_SECRET` ends every session at once.
- Used nonces are remembered in memory, so they reset on restart and are not shared between API instances. A nonce still expires 5 minutes after it was issued either way.

---

### Endpoint: `/v1/notes` <a id="endpoint-write-note"></a>

- **Method**: POST `/v1/notes`
- **Description**: Writes a note against an invoice via `createNote`. The content is encrypted by the caller and reaches this API as an opaque hex payload, so the plaintext is never seen or stored server-side.

#### Authorization

Requires `Authorization: Bearer <token>` from [sign-in](#endpoint-sign-in). The note is written as the signed-in account. This API pays the gas; it does not yet check that the account is a party on the invoice.

#### **Request Body**

```json
{
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "author": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
  "content": "0x4a6f8b2c1d...",
  "share": true
}
```

#### Field Details

| Field       | Type    | Required | Description                                                        |
| :---------: | :-----: | :------: | :----------------------------------------------------------------: |
| `invoiceId` | string  | ✅       | On-chain invoice ID, base-10.                                      |
| `author`    | string  | ❌       | Defaults to the signed-in account; if sent, it must be that account. |
| `content`   | string  | ✅       | Already-encrypted note, `0x`-prefixed hex, at most 4096 bytes.     |
| `share`     | boolean | ❌       | `true` publishes the note to both parties; private otherwise.      |

#### **Response**

**Success (200)**:

```json
{ "success": true, "txHash": "0x123456..." }
```

**Error (400 / 413)**:

```json
{ "success": false, "error": "Invalid author address" }
```

- `400` for a malformed body, a bad `invoiceId` or `author`, or a `content` that is missing or not `0x`-prefixed hex.
- `401` without a valid session token.
- `403` when `author` is not the signed-in account.
- `413` for a `content` payload over 4096 bytes.
- `503` when sign-in is not configured.

**Notes**:

- Encryption is entirely the caller's concern. This API validates that `content` is well-formed hex within the size cap and relays the bytes to the contract unchanged; it holds no note key and cannot read a note back.
- The write returns as soon as the transaction is broadcast; the receipt is not awaited.

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/notes \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_SESSION_TOKEN" \
-d '{
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "content": "0x4a6f8b2c1d",
  "share": true
}'
```

---

### Endpoint: `/v1/notes/open` <a id="endpoint-open-note"></a>

- **Method**: POST `/v1/notes/open`
- **Description**: Records that the signed-in account has opened a note, via `setOpened`. Only opening is written on chain; closing is a client-side state. Requires `Authorization: Bearer <token>` from [sign-in](#endpoint-sign-in).

#### **Request Body**

```json
{
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "author": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
  "noteId": "1"
}
```

#### Field Details

| Field       | Type   | Required | Description                          |
| :---------: | :----: | :------: | :----------------------------------: |
| `invoiceId` | string | ✅       | On-chain invoice ID, base-10.        |
| `author`    | string | ❌       | Defaults to the signed-in account; if sent, it must be that account. |
| `noteId`    | string | ✅       | Id of the note to mark opened.       |

#### **Response**

**Success (200)**:

```json
{ "success": true, "txHash": "0x123456..." }
```

**Error (400)**:

```json
{ "success": false, "error": "noteId must be a base-10 integer" }
```

As for `/v1/notes`: `401` without a valid session, `403` when `author` is not the signed-in account, `503` when sign-in is not configured.

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/notes/open \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_SESSION_TOKEN" \
-d '{
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "noteId": "1"
}'
```

---

## Configuration

[`config.yaml`](config.yaml) holds one section per network. `network:` selects which one; it is itself an `${NETWORK}` reference, so the network is chosen by the environment — `.env` locally, and `NETWORK=testnet` in [`docker-compose.yml`](docker-compose.yml). Setting `NETWORK` always wins, and leaving it unset is an error rather than a silent default. Only the selected section is validated, so a network that is not deployed yet cannot break startup.

```yaml
network: ${NETWORK}

networks:
  local:
    rpc:
      # A reference, not a literal: a container cannot reach the host's node on
      # 127.0.0.1, so compose overrides these with host.docker.internal.
      http: "${LOCAL_RPC_URL}"
      ws: "${LOCAL_WSS}"
      dialTimeout: 10s
    urls:
      explorer: ""                       # a local chain has no explorer
      checkout: "http://localhost:3000/checkout/?id="
      subgraph: "${END_POINT}"
      callback: "${URL}"
      discordWebhook: "${DISCORD_WEBHOOK_URL}"
    signerKey: "${LOCAL_CALLER}"
    tokens:
      ETH:
        address: "0x0000000000000000000000000000000000000000"
        decimals: 18
      USDC:
        address: "0x2d19afC50EaaCe1CE730ab2A9a5D87712b0d4bCc"
        decimals: 6
    contracts:
      paymentProcessor: "0x..."
      oracleManager: "0x..."   # optional; without it /v1/exchangeRate is 503
      sweeper: "${SWEEPER_CONTRACT}"   # read by the sidecar, not by this API
      wrappedNative: "0x4200000000000000000000000000000000000006"
      # ...
    services:
      # host:port of the fee-receiver sidecar; empty disables /v1/fee-receivers
      feeReceiver: "${FEE_RECEIVER_ADDRESS}"
    auth:
      # hosts the website signs in from; empty disables sign-in and notes
      domains: ["localhost:3000"]
```

- **`${VAR}` references** are resolved from the environment at load time. Endpoints whose URL embeds a credential (the RPC provider key, the Discord webhook) are written this way so the secret stays in `.env` and out of the repository. An unset variable fails startup naming the variable.
- **`urls.explorer`** is a block explorer root with no trailing slash. When empty, transaction links fall back to the bare hash.
- **`tokens`** is the one place a payment token is defined. `paymentTokens` in a create request names a symbol from this table, and callback payloads render an event's token address back to its symbol and decimals through the same table.
- **`signerKey`** selects which key signs transactions, per network: `${LOCAL_CALLER}` on `local` and `${PASS}` on the deployed networks, so a local run cannot touch a deployed network's key. It is parsed once at startup, and a malformed key fails startup rather than the first transaction.
- **`services.feeReceiver`** is the fee-receiver sidecar as `host:port` — a gRPC target, not a URL. Leaving it empty disables `/v1/fee-receivers` with a `503` and a line in the startup log, the same way an unset `oracleManager` disables exchange rates. The connection is plaintext and expects a loopback sidecar; it holds the platform's stealth keys and must not be exposed off-host.
- **`contracts.sweeper`** and **`contracts.wrappedNative`** are read only by the sidecar — the Go API ignores them. `sweeper` is the contract each fee receiver approves for the fee token; `wrappedNative` is what a native-token payment is approved in. Either may be empty, and the sidecar warns at startup then fails the calls that need them.
- **`auth.domains`** lists the hosts the website [signs in](#endpoint-sign-in) from, exactly as they appear on the first line of the SIWE message: a host, with a port if it has one (`localhost:3000`), never a URL. A message from any other host is refused, which stops a lookalike site from reusing a signature it collected. Sign-in also needs `SESSION_SECRET`; without either, sign-in and the notes endpoints return `503` and the startup log says why.
- **`CONFIG_PATH`** points at a different config file.

Run against a local chain with:

```bash
NETWORK=local go run ./cmd/server
```

### Environment variables

Values that are secrets rather than settings stay in `.env`:

| Variable                   | Purpose                                                    |
| :------------------------: | :--------------------------------------------------------: |
| `KEY`                      | Value clients must send in `X-API-KEY`.                    |
| `SESSION_SECRET`           | Signs sign-in nonces and session tokens. At least 32 characters, e.g. `openssl rand -hex 32`; shorter fails startup. Unset disables sign-in and notes. Changing it ends every session. |
| `API_KEY`                  | Key this API sends when posting callbacks.                 |
| `PASS`                     | Private key that signs transactions on the deployed networks. |
| `LOCAL_CALLER`             | Private key that signs on `local` (anvil's default account). |
| `PORT`                     | HTTP port; defaults to `8080`.                             |
| `LOCAL_RPC_URL`, `LOCAL_WSS` | Referenced by the `local` section. `127.0.0.1` on the host; compose overrides them with `host.docker.internal` so a container can reach the host's node. |
| `TEST_NET_RPC_URL`, `TEST_NET_WSS` | Referenced by the `testnet` section.               |
| `MAIN_NET_RPC_URL`, `MAIN_NET_WSS` | Referenced by the `mainnet` section.               |
| `NETWORK`                  | Selects the network section. Required — `network:` in the file is a `${NETWORK}` reference. |
| `CONFIG_PATH`              | Path to the config file. Defaults to `config.yaml`.        |
| `END_POINT`, `URL`, `DISCORD_WEBHOOK_URL` | Subgraph, callback and Discord endpoints.   |
| `PRODUCTION`               | When set, `.env` is not read; the container supplies the environment. Callbacks to the intermediated platform are only sent over HTTP when it is `1` or `true`; when it is empty, `0`, `false` or unset, each one is built and logged, and treated as accepted. |
| `AUTOMATION_POLL_INTERVAL` | How often to poll for due automation tasks.                |
| `FEE_RECEIVER_ADDRESS`     | `host:port` of the fee-receiver sidecar. Unset disables `/v1/fee-receivers`. |
| `SWEEPER_CONTRACT`         | Sweeper address, referenced by `contracts.sweeper`.        |
| `SPONSOR`                  | Sidecar: relayer key that sponsors the delegations and signs the fee authorization. |
| `SPENDING`, `VIEWING`      | Sidecar: EIP-5564 keys the stealth addresses are derived from. |
| `META_STEALTH_ADDRESS`     | Sidecar: stealth meta-address URI new receivers derive from. |
| `RECEIVER_SEED`            | Sidecar: 32-byte hex secret receivers are derived from, so a repeated `/v1/fee-receivers` call returns the same receivers. |

The last four are read by the sidecar rather than by this API, but live in the same repo-root `.env` — it loads that file explicitly, since it runs from its own directory.

---

## Notes

- All endpoints require an `X-API-KEY` header, enforced by `AccessControlMiddleWare`, except `GET /` and the fee receiver endpoint.
- Invoice states are: `INITIATED` (1), `PAID` (2), `REFUNDED` (3), `CANCELED` (4), `DISPUTED` (5), `DISPUTE_RESOLVED` (6), `DISPUTE_DISMISSED` (7), `DISPUTE_SETTLED` (8), `RELEASED` (9).
- The contracts use Chainlink price feeds (`AggregatorV3Interface`) for USD-to-token conversions, supporting the native token and ERC20 tokens.
- The intermediated platform operator, retrieved via `PaymentProcessorStorage.GetIntermediatedPlatformsOperator`, controls privileged operations (`createSingleInvoice`, `createMetaInvoice`, `createDispute`).
- Transaction links are built from `urls.explorer` for the selected network — `https://sepolia.basescan.org` on Base Sepolia.
- The client-provided `orderId` is hashed to a `uint216` by `invoice.InvoiceIDToUint216` for on-chain storage, producing a numeric string (e.g., `"59808737901387817475691215581034097896123425895641016234844280889"`). That value is the `{invoiceId}` path segment.
- Blockchain reverts are mapped to human-readable messages by `revert.Descriptions`, and to status codes by `revert.StatusCodes`:
  - `The buyer and seller addresses cannot be the same.`
  - `The account balance is insufficient to perform this action.`
  - `The provided dispute resolution is invalid.`
  - `The invoice is not in a valid state for this action.`
  - `The native token payment is invalid for this invoice.`
  - `The specified payment token is not supported or invalid.`
  - `The seller's payout share is invalid.`
  - `An invoice with this identifier already exists.`
  - `The specified invoice does not exist.`
  - `A meta-invoice with this identifier already exists.`
  - `The caller is not authorized to perform this action.`
  - `The price cannot be zero.`
  - `The price specified is too low.`
