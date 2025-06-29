# Sapphire Contract API – REST Endpoints

This API exposes smart contract functions through HTTP endpoints for invoice creation, escrow management, and invoice data retrieval. It interacts with the Sapphire DAO's blockchain contracts on the Polygon Amoy testnet using the `AdvancedPaymentProcessor` contract at address `0x57BFD7c3D1d14b82AB7Ad135B2E56e330F65D27f`.

**Base URL**: `https://contract-api-production.up.railway.app/`

## Endpoints

- [POST `/create`](#endpoint-create)
- [POST `/release`](#endpoint-release)
- [GET `/invoices/{id}`](#endpoint-invoicesid)

---

### Endpoint: `/create`

- **Method**: POST
- **Description**: Creates one or more on-chain invoices using the `AdvancedPaymentProcessor` contract's `createSingleInvoice` or `createMetaInvoice` functions.
- **Request Body**:
  ```json
  [
    {
      "orderId": "string",
      "seller": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
      "buyer": "0x60D7dD3b4248D53Abba8DA999B22023656A2E4B3",
      "invoiceExpiryDuration": 864000,
      "timeBeforeCancelation": 864000,
      "releaseWindow": 864000,
      "price": "8680000000"
    }
  ]
  ```

#### Field Details

| Field                   | Type    | Required | Description                                                                    |
| ----------------------- | ------- | -------- | ------------------------------------------------------------------------------ |
| `orderId`               | string  | ✅       | Client-side identifier for the invoice (e.g., "string")                        |
| `seller`                | string  | ✅       | Ethereum address of the seller (e.g., `0xabc123...`)                           |
| `buyer`                 | string  | ✅       | Ethereum address of the buyer (e.g., `0xdef456...`)                            |
| `invoiceExpiryDuration` | integer | ✅       | Invoice expiration time in seconds (e.g., `864000` for 10 days)                |
| `timeBeforeCancelation` | integer | ✅       | Time window (in seconds) before the buyer can cancel the invoice               |
| `releaseWindow`         | integer | ✅       | Time window (in seconds) the seller has to release funds after payment         |
| `price`                 | string  | ✅       | Invoice price in **USD** with **8 decimal places** (e.g., `100000000` = $1.00) |

**Notes**:

- `price` is specified in USD with 8 decimal places (e.g., `100000000` = $1.00, `2500000000` = $25.00).
- The price is converted to the payment token amount using on-chain oracle pricing.
- A single invoice uses `createSingleInvoice`; multiple invoices in the array use `createMetaInvoice` with the first `buyer` address.
- The endpoint returns a checkout URL with a token generated from the invoice key (derived from transaction logs).

**Response**:

- **Success (200)**:
  ```json
  {
    "url": "https://contract-api-production.up.railway.app/<token>"
  }
  ```
  - `url`: A checkout URL with a generated token for the created invoice(s).
- **Error (400)**:
  ```json
  {
    "error": "invalid request body",
    "reason": "<decoding error message>"
  }
  ```
  - Returned if the JSON request body cannot be decoded (e.g., malformed JSON or incorrect field types).
- **Error (500)**:
  ```json
  {
    "error": "Error sending transaction",
    "reason": "<blockchain error message>"
  }
  ```
  - Returned if the blockchain transaction for creating an invoice fails (e.g., invalid parameters, insufficient gas, or contract reversion).
- **Error (500)**:
  ```json
  {
    "error": "failed to generate token",
    "reason": "<token generation error message>"
  }
  ```
  - Returned if token generation for the checkout URL fails.

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/create \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '[
  {
    "orderId": "inv001",
    "seller": "0xabc123...",
    "buyer": "0xdef456...",
    "invoiceExpiryDuration": 864000,
    "timeBeforeCancelation": 864000,
    "releaseWindow": 864000,
    "price": "100000000"
  }
]'
```

---

### Endpoint: `/release`

- **Method**: POST
- **Description**: Releases escrow funds or resolves a dispute for a specific invoice using the `AdvancedPaymentProcessor` contract's `releasePayment` or `handleDispute` functions.
- **Request Body**:
  ```json
  {
    "orderId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
    "resolution": 1,
    "sellersShare": "9000"
  }
  ```

#### Field Details

| Field          | Type    | Required                                    | Description                                                             |
| -------------- | ------- | ------------------------------------------- | ----------------------------------------------------------------------- |
| `orderId`      | string  | ✅                                          | Client-side identifier for the invoice (e.g., "string")                 |
| `resolution`   | integer | ✅                                          | Enum value specifying the action type (see MarketplaceAction below)     |
| `sellersShare` | string  | ❌ Only if `resolution = 3` (SettleDispute) | Seller's share in **basis points** (e.g., `10000` = 100%, `9000` = 90%) |

#### MarketplaceAction Enum (`resolution`)

| Value | Name           | Description                          | Contract Function                       |
| ----- | -------------- | ------------------------------------ | --------------------------------------- |
| `0`   | Pending        | Default state, no action taken       | N/A                                     |
| `1`   | Release        | Release funds to the seller          | `releasePayment`                        |
| `2`   | DismissDispute | Dismiss an active dispute            | `handleDispute` with `DISPUTEDISMISSED` |
| `3`   | SettleDispute  | Resolve a dispute by splitting funds | `handleDispute` with `DISPUTESETTLED`   |

**Notes**:

- If `orderId` is a 66-character hex string starting with `0x`, it is used directly; otherwise, it is hashed using Keccak256.
- `sellersShare` is required only for `SettleDispute` and defaults to `0` if not provided.
- The transaction hash is returned with a Polygon Amoy explorer URL prefix (`https://amoy.polygonscan.com/tx/`).

**Response**:

- **Success (200)**:
  ```json
  {
    "status": "success",
    "hash": "https://amoy.polygonscan.com/tx/0x123456..."
  }
  ```
  - `status`: Indicates the transaction was successful.
  - `hash`: Transaction hash with the Polygon Amoy explorer URL.
- **Error (400)**:
  ```json
  {
    "error": "invalid request body",
    "reason": "<decoding error message>"
  }
  ```
  - Returned if the JSON request body cannot be decoded (e.g., malformed JSON, missing `orderId` or `resolution`, or invalid `sellersShare`).
- **Error (500)**:
  ```json
  {
    "error": "failed to generate order ID hash",
    "reason": "<hashing error message>"
  }
  ```
  - Returned if the Keccak256 hashing of the `orderId` fails.
- **Error (500)**:
  ```json
  {
    "error": "Error sending transaction",
    "reason": "<blockchain error message>"
  }
  ```
  - Returned if the blockchain transaction for releasing escrow fails (e.g., invalid `orderId`, unsupported `resolution`, or contract reversion).

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/release \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "orderId": "inv001",
  "resolution": 1
}'
```

---

### Endpoint: `/invoices/{id}`

- **Method**: GET
- **Description**: Retrieves invoice data for a given `orderId` from The Graph, with the `orderId` hashed via Keccak256.

#### Path Parameters

| Name | Type   | Required | Description                           |
| ---- | ------ | -------- | ------------------------------------- |
| `id` | string | ✅       | Order ID to fetch associated invoices |

**Notes**:

- The `id` is extracted from the URL path (`/invoices/{id}`) and hashed using Keccak256 before querying The Graph.
- Pagination parameters (`first`, `skip`) are not supported in this implementation.

**Response**:

- **Success (200)**:
  ```json
  {
    "invoices": [
      {
        "orderId": "0xabc123...",
        "seller": "0xabc123...",
        "buyer": "0xdef456...",
        "createdAt": "2025-06-17T22:32:00Z",
        "price": "100000000",
        "status": "Pending"
      }
    ]
  }
  ```
  - Returns an array of invoices associated with the hashed `orderId`.
- **Error (400)**:
  ```json
  {
    "error": "Invalid or missing invoice ID",
    "reason": "<specific error message>"
  }
  ```
  - Returned if the `id` path parameter is missing or the URL is malformed.
- **Error (400)**:
  ```json
  {
    "error": "Missing orderId parameter",
    "reason": "<specific error message>"
  }
  ```
  - Returned if the `id` is empty.
- **Error (500)**:
  ```json
  {
    "error": "failed to generate order ID hash",
    "reason": "<hashing error message>"
  }
  ```
  - Returned if the Keccak256 hashing of the `id` fails.
- **Error (500)**:
  ```json
  {
    "error": "failed to fetch invoice data",
    "reason": "<GraphQL query error message>"
  }
  ```
  - Returned if the query to The Graph fails (e.g., network issues or invalid response).

**Example**:

```bash
curl "https://contract-api-production.up.railway.app/invoices/inv001"
```

---

## Notes

- All endpoints are routed through a multiplexer (`http.ServeMux`) defined in the `Route` function.
- The `/create` and `/release` endpoints are protected by `AccessControlMiddleWare` for authentication/authorization.
- The `/create` endpoint generates a checkout URL with a token derived from the invoice key in the transaction logs.
- The `/release` endpoint supports `Release`, `DismissDispute`, and `SettleDispute` actions, with appropriate contract function calls (`releasePayment` or `handleDispute`).
- The `/invoices/{id}` endpoint queries The Graph using a Keccak256-hashed `orderId` extracted from the URL path.
- Prices are in USD with 8 decimal places and converted to token amounts on-chain using an oracle.
- All blockchain interactions occur on the Polygon Amoy testnet, with transactions linked