# Sapphire Contract API – REST Endpoints

This API provides HTTP endpoints for interacting with the Sapphire DAO's `IntermediatedPaymentProcessor`, `SimplePaymentProcessor` and `Notes` smart contracts on **Base Sepolia**. It supports invoice creation, cancellation, refunding, dispute creation, dispute resolution, fund release, and invoice notes.

Contract addresses and endpoints are not compiled in — they come from [`config.yaml`](config.yaml), which holds one section per network (`local`, `testnet`, `mainnet`). See [Configuration](#configuration).

Platform fees are collected through one-time stealth addresses issued by the [fee-receiver sidecar](services/fee-receiver/README.md), a separate process this API calls over gRPC. It reads the same `config.yaml`, so contract addresses cannot drift between the two.

**Base URL**: `https://sapphiredaotesting.com/`

## Endpoints

| Method | Path                                            | Description                          |
| ------ | ----------------------------------------------- | ------------------------------------ |
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
| POST   | [`/v1/fee-receivers`](#endpoint-create-fee-receivers) | Derive and approve stealth fee receivers |
| POST   | [`/v1/fee-receivers/authorization`](#endpoint-authorize-fee-receivers) | Sign the fee authorization |
| POST   | [`/notes`](#endpoint-notes)                      | Invoice notes (all actions)          |

The invoice id is a path segment on every invoice operation, so the request body carries only what is specific to that operation. `release`, `cancel` and `disputes` take no body at all.

Every endpoint requires an `X-API-KEY` header, enforced by `AccessControlMiddleWare`, except `GET /` and the two [fee receiver](#endpoint-create-fee-receivers) endpoints.

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
[
  {
    "orderId": "550e8400-e29b-41d4-a716-446655440000",
    "seller": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
    "price": 8680000000,
    "escrowHoldPeriod": 604800,
    "currency": "USD",
    "paymentTokens": ["ETH", "USDC"]
  }
]
```

#### Field Details

| Field              | Type     | Required | Description                                                                              |
| ------------------ | -------- | -------- | ---------------------------------------------------------------------------------------- |
| `orderId`          | string   | ✅       | Unique client-side identifier for the invoice (e.g., a UUID or any string).              |
| `seller`           | string   | ✅       | Ethereum address of the seller. Must not be the zero address.                            |
| `price`            | number   | ✅       | Invoice price in cents; scaled on the server using the `currency` precision.              |
| `escrowHoldPeriod` | number   | ✅       | Duration in seconds for holding funds in escrow (e.g., `604800` = 7 days).                |
| `currency`         | string   | ✅       | Pricing currency; sets the decimal precision applied to `price` (`USD` = 8 decimals).     |
| `paymentTokens`    | string[] | ✅       | Token **symbols** the buyer may pay with, e.g. `["ETH", "USDC"]`. At least one.            |

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
  "url": "https://sapphire-dao-website-six.vercel.app/checkout/?id=kWMRqBU-7H64tqp04CM5IfzKRMP1DbH2Ytg5",
  "orders": {
    "550e8400-e29b-41d4-a716-446655440000": {
      "seller": "0x329C3E1bEa46Abc22F307eE30Cbb522B82Fe7082",
      "orderId": "59808737901387817475691215581034097896123425895641016234844280889"
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
- The client-provided `orderId` is hashed to a `uint216` by `invoice.OrderIDToUint216` for on-chain storage, producing the numeric `orderId` in the response. That numeric id is the `{invoiceId}` used by every other endpoint.

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
-d '[
  {
    "orderId": "550e8400-e29b-41d4-a716-446655440000",
    "seller": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
    "price": 8680000000,
    "escrowHoldPeriod": 604800,
    "currency": "USD",
    "paymentTokens": ["ETH", "USDC"]
  }
]'
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
| ------------- | ------ | -------- | ---------------------------------------------------------------------- |
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
| ------------- | ------- | ------------------------------------------- | ------------------------------------------------------------------------ |
| `resolution`  | integer | ✅                                          | Enum value specifying the action type (see MarketplaceAction below).     |
| `sellerShare` | string  | ❌ Only if `resolution = 2` (SettleDispute) | Seller's share in basis points (e.g., `"10000"` = 100%, `"9000"` = 90%). |

#### MarketplaceAction Enum (`resolution`)

| Value | Name           | Description                                                                                            | Contract Function                        |
| ----- | -------------- | ------------------------------------------------------------------------------------------------------ | ---------------------------------------- |
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
| ------ | -------- | ----------------------------------------------------------------------------------------------- |
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
- **Description**: Derives one-time [EIP-5564](https://eips.ethereum.org/EIPS/eip-5564) stealth addresses to collect platform fees, upgrades each to an EIP-7702 delegator, and approves the sweeper to move the fee token out of them later.

The work is done by the [fee-receiver sidecar](services/fee-receiver/README.md), reached over gRPC at `services.feeReceiver`. This API forwards the request and translates the sidecar's gRPC status into an HTTP one.

**No `X-API-KEY`.** Both fee receiver endpoints are unauthenticated, unlike the rest of the API. Each call to this one sends a relayer-sponsored transaction per receiver, so whatever fronts the API is what limits who can spend that gas.

Fee receivers are issued in **two steps**, and the split is the point. This call spends gas but hands back only the *ephemeral public keys* — never the addresses. Those keys are the only way to recover the addresses, so **store them before doing anything else**: they cannot be re-derived, and the gas is spent whether or not they are kept.

#### **Request Body**

```json
{ "processor": "intermediated", "quantity": 3, "paymentToken": "USDC" }
```

| Field          | Type    | Required | Description                                                                                   |
| -------------- | ------- | -------- | --------------------------------------------------------------------------------------------- |
| `processor`    | string  | ✅       | `simple` or `intermediated` — which processor will verify the authorization.                   |
| `quantity`     | integer | ❌       | How many receivers to derive, `1`–`5`. Defaults to `1`. A meta invoice needs one per sub-invoice. |
| `paymentToken` | string  | ❌       | Token **symbol** the fee is collected in, e.g. `"USDC"`. Omit for a native-token payment.       |

**About `paymentToken`**: as with `paymentTokens` on invoice creation, callers name the token by symbol and it is resolved through the `tokens` table for the selected network. Omitting it — or naming `ETH`, which maps to the zero address — means the payment is native, and the sidecar approves the chain's wrapped native token (`contracts.wrappedNative`) instead.

#### **Response**

**Success (200)** — one key per receiver, in the order they were derived:

```json
{
  "status": "success",
  "ephemeralPublicKeys": [
    "0x02f7a1c3b45d6e89f0123456789abcdef0123456789abcdef0123456789abcdef01"
  ]
}
```

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/fee-receivers \
-H "Content-Type: application/json" \
-d '{ "processor": "intermediated", "quantity": 1, "paymentToken": "USDC" }'
```

---

### Endpoint: Authorize fee receivers

- **Method**: POST `/v1/fee-receivers/authorization`
- **Description**: Turns the stored ephemeral public keys back into fee receiver addresses and returns the fee signer's signature over them, to pass on-chain as the processor's `_feeReceiver(s)` / `_data` arguments.

Re-deriving is what makes it safe to accept these keys back from a client: a key can only ever resolve to an address the platform's spending and viewing keys control, never to one the caller chose.

#### **Request Body**

```json
{
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "processor": "intermediated",
  "kind": "meta",
  "paymentToken": "USDC",
  "ephemeralPublicKeys": ["0x02f7a1c3...", "0x03b8d2e4..."]
}
```

| Field                 | Type     | Required | Description                                                                           |
| --------------------- | -------- | -------- | --------------------------------------------------------------------------------------- |
| `invoiceId`           | string   | ✅       | The on-chain invoice id, base-10. For `kind: "meta"`, the meta-invoice id.               |
| `processor`           | string   | ✅       | `simple` or `intermediated`. Must match the processor the invoice lives on.              |
| `kind`                | string   | ❌       | `single` (default) or `meta`. `meta` requires `processor: "intermediated"`.              |
| `ephemeralPublicKeys` | string[] | ✅       | The keys from the previous call, **in sub-invoice order**. At most 5.                    |
| `paymentToken`        | string   | ❌       | Token symbol, as above.                                                                  |

**About `kind`**: it is never inferred from how many keys you send. A meta invoice signs one digest over the whole address array, and a meta invoice holding a single sub-invoice still needs that array form — so a `single` request must carry exactly one key, and a meta invoice must say so explicitly.

#### **Response**

**Success (200)** — `feeReceivers` is index-aligned with the meta-invoice's `subInvoiceIds`, and the one `signature` covers the whole array:

```json
{
  "status": "success",
  "feeReceivers": ["0x9ba1...", "0x4cd2..."],
  "signature": "0x7f3e..."
}
```

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/v1/fee-receivers/authorization \
-H "Content-Type: application/json" \
-d '{
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "processor": "intermediated",
  "kind": "single",
  "ephemeralPublicKeys": ["0x02f7a1c3..."]
}'
```

#### Errors (both endpoints)

| Status | Meaning                                                                                                         |
| ------ | ----------------------------------------------------------------------------------------------------------------- |
| `400`  | Rejected by this API (unknown `processor`, `kind` or token symbol; a non-integer `invoiceId`; no keys) or by the sidecar (`quantity` out of range, a meta invoice on a simple processor, a `single` invoice with several keys). |
| `502`  | The sidecar failed internally — a chain error, or one of its keys is unset. Its own message is generic by design.  |
| `503`  | `services.feeReceiver` is not configured for this network, or the relayer has no native balance to sponsor the delegations. |
| `504`  | The sidecar did not answer within its 25s call budget.                                                            |

A rejection from the sidecar keeps the sidecar's own wording in `reason`:

```json
{ "error": "error creating fee receivers", "reason": "quantity must be between 1 and 5" }
```

---

### Endpoint: `/notes`

- **Method**: POST `/notes`
- **Description**: Reads and writes invoice notes on the `Notes` contract. Notes are encrypted with a server-held key before they reach the chain, so the ciphertext is public but the content is not. A single endpoint serves four actions, selected by the `action` field.

> This endpoint keeps its action-dispatching shape, and its unversioned path, so the website can forward a request body unchanged.

#### Actions

| `action`    | Description                                                                              | On-chain |
| ----------- | ---------------------------------------------------------------------------------------- | -------- |
| `create`    | Encrypts a note and writes it for an invoice via `createNote`.                            | ✅        |
| `setOpened` | Marks a note as opened by the author via `setOpened`. `open: false` is a no-op.           | ✅        |
| `encrypt`   | Encrypts content and returns it as hex, for the `storageRef` note a client sends itself.  | ❌        |
| `decrypt`   | Reads notes by id from the chain and decrypts the ones the caller is allowed to see.      | ❌        |

#### Authorization

The `X-API-KEY` header is the only check. This API pays the gas and signs on the author's behalf, so the caller is trusted to have authenticated whoever the `author` and `viewer` fields name, and to have confirmed that they are a party on the invoice. Do not expose this endpoint to browsers directly.

#### **Request Body**

```json
{
  "action": "create",
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "author": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
  "content": "Left at the door",
  "share": true
}
```

#### Field Details

| Field       | Type     | Required                    | Description                                                             |
| ----------- | -------- | --------------------------- | ----------------------------------------------------------------------- |
| `action`    | string   | ✅                          | `create`, `setOpened`, `encrypt` or `decrypt`.                          |
| `invoiceId` | string   | ✅ except `encrypt`         | On-chain invoice ID.                                                    |
| `author`    | string   | ✅ for `create`/`setOpened` | Address the note is attributed to; must be a party on the invoice.      |
| `content`   | string   | ✅ for `create`/`encrypt`   | Note text, at most 20 characters.                                       |
| `share`     | boolean  | ❌                          | `true` publishes the note to both parties; private otherwise.           |
| `noteId`    | string   | ✅ for `setOpened`          | Id of the note to mark opened.                                          |
| `open`      | boolean  | ❌                          | Only `true` is written on chain.                                        |
| `noteIds`   | string[] | ✅ for `decrypt`            | Up to 50 note ids to read.                                              |
| `viewer`    | string   | ❌ for `decrypt`            | Address reading the notes; required to see that address's private notes. |

#### **Response**

**Success (200)** — `create` and `setOpened`:

```json
{ "success": true, "txHash": "0x123456..." }
```

**Success (200)** — `encrypt`:

```json
{ "success": true, "payload": "0x4a6f..." }
```

**Success (200)** — `decrypt`. `content` is `null` for a note the caller may not read, or one that could not be read back:

```json
{
  "success": true,
  "notes": [
    { "noteId": "1", "content": "Left at the door" },
    { "noteId": "2", "content": null }
  ]
}
```

**Error (400 / 413)**:

```json
{ "success": false, "error": "Invalid author address" }
```

- `400` for a malformed body, an unknown action, or a missing required field.
- `413` for content over 20 characters or more than 50 `noteIds`.

**Notes**:

- Notes are encrypted with AES-256-CBC under `sha256(NOTES_SECRET_KEY)` and stored as `<base64 iv>:<base64 ciphertext>`. This matches the scheme the website used, so notes written by either side stay readable by both. With `NOTES_SECRET_KEY` unset, notes are written and read in the clear.
- `decrypt` resolves notes by `(invoiceId, noteId)` from the chain rather than accepting ciphertext from the caller, so a party cannot decrypt a blob lifted from another invoice. Shared notes are readable by anyone; a private note decrypts only for the `viewer` that authored it.
- Write actions return as soon as the transaction is broadcast; the receipt is not awaited.

**Example**:

```bash
curl -X POST https://sapphiredaotesting.com/notes \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "action": "decrypt",
  "invoiceId": "59808737901387817475691215581034097896123425895641016234844280889",
  "noteIds": ["1", "2"],
  "viewer": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4"
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
```

- **`${VAR}` references** are resolved from the environment at load time. Endpoints whose URL embeds a credential (the RPC provider key, the Discord webhook) are written this way so the secret stays in `.env` and out of the repository. An unset variable fails startup naming the variable.
- **`urls.explorer`** is a block explorer root with no trailing slash. When empty, transaction links fall back to the bare hash.
- **`tokens`** is the one place a payment token is defined. `paymentTokens` in a create request names a symbol from this table, and callback payloads render an event's token address back to its symbol and decimals through the same table.
- **`signerKey`** selects which key signs transactions, per network: `${LOCAL_CALLER}` on `local` and `${PASS}` on the deployed networks, so a local run cannot touch a deployed network's key. It is parsed once at startup, and a malformed key fails startup rather than the first transaction.
- **`services.feeReceiver`** is the fee-receiver sidecar as `host:port` — a gRPC target, not a URL. Leaving it empty disables `/v1/fee-receivers` with a `503` and a line in the startup log, the same way an unset `oracleManager` disables exchange rates. The connection is plaintext and expects a loopback sidecar; it holds the platform's stealth keys and must not be exposed off-host.
- **`contracts.sweeper`** and **`contracts.wrappedNative`** are read only by the sidecar — the Go API ignores them. `sweeper` is the contract each fee receiver approves for the fee token; `wrappedNative` is what a native-token payment is approved in. Either may be empty, and the sidecar warns at startup then fails the calls that need them.
- **`CONFIG_PATH`** points at a different config file.

Run against a local chain with:

```bash
NETWORK=local go run ./cmd/server
```

### Environment variables

Values that are secrets rather than settings stay in `.env`:

| Variable                   | Purpose                                                    |
| -------------------------- | ---------------------------------------------------------- |
| `KEY`                      | Value clients must send in `X-API-KEY`.                    |
| `API_KEY`                  | Key this API sends when posting callbacks.                 |
| `NOTES_SECRET_KEY`         | Encrypts note content. Unset means notes are stored plain. |
| `PASS`                     | Private key that signs transactions on the deployed networks. |
| `LOCAL_CALLER`             | Private key that signs on `local` (anvil's default account). |
| `PORT`                     | HTTP port; defaults to `8080`.                             |
| `LOCAL_RPC_URL`, `LOCAL_WSS` | Referenced by the `local` section. `127.0.0.1` on the host; compose overrides them with `host.docker.internal` so a container can reach the host's node. |
| `TEST_NET_RPC_URL`, `TEST_NET_WSS` | Referenced by the `testnet` section.               |
| `MAIN_NET_RPC_URL`, `MAIN_NET_WSS` | Referenced by the `mainnet` section.               |
| `NETWORK`                  | Selects the network section. Required — `network:` in the file is a `${NETWORK}` reference. |
| `CONFIG_PATH`              | Path to the config file. Defaults to `config.yaml`.        |
| `END_POINT`, `URL`, `DISCORD_WEBHOOK_URL` | Subgraph, callback and Discord endpoints.   |
| `PRODUCTION`               | When set, `.env` is not read; the container supplies the environment. |
| `AUTOMATION_POLL_INTERVAL` | How often to poll for due automation tasks.                |
| `FEE_RECEIVER_ADDRESS`     | `host:port` of the fee-receiver sidecar. Unset disables `/v1/fee-receivers`. |
| `SWEEPER_CONTRACT`         | Sweeper address, referenced by `contracts.sweeper`.        |
| `SPONSOR`                  | Sidecar: relayer key that sponsors the delegations and signs the fee authorization. |
| `SPENDING`, `VIEWING`      | Sidecar: EIP-5564 keys the stealth addresses are derived from. |
| `META_STEALTH_ADDRESS`     | Sidecar: stealth meta-address URI new receivers derive from. |

The last four are read by the sidecar rather than by this API, but live in the same repo-root `.env` — it loads that file explicitly, since it runs from its own directory.

---

## Notes

- All endpoints require an `X-API-KEY` header, enforced by `AccessControlMiddleWare`, except `GET /` and the two fee receiver endpoints.
- Invoice states are: `INITIATED` (1), `PAID` (2), `REFUNDED` (3), `CANCELED` (4), `DISPUTED` (5), `DISPUTE_RESOLVED` (6), `DISPUTE_DISMISSED` (7), `DISPUTE_SETTLED` (8), `RELEASED` (9).
- The contracts use Chainlink price feeds (`AggregatorV3Interface`) for USD-to-token conversions, supporting the native token and ERC20 tokens.
- The intermediated platform operator, retrieved via `PaymentProcessorStorage.GetIntermediatedPlatformsOperator`, controls privileged operations (`createSingleInvoice`, `createMetaInvoice`, `createDispute`).
- Transaction links are built from `urls.explorer` for the selected network — `https://sepolia.basescan.org` on Base Sepolia.
- The client-provided `orderId` is hashed to a `uint216` by `invoice.OrderIDToUint216` for on-chain storage, producing a numeric string (e.g., `"59808737901387817475691215581034097896123425895641016234844280889"`). That value is the `{invoiceId}` path segment.
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
