# Sapphire Contract API – REST Endpoints

This API exposes smart contract functions through HTTP endpoints for invoice creation, escrow management, dispute resolution, cancellation, refund, and invoice data retrieval. It interacts with the Sapphire DAO's `AdvancedPaymentProcessor` contract on the Ethereum Sepolia testnet at address `0x953ff222255730544c8e118a2ccb5dfb856bfbad`.

**Base URL**: `https://contract-api-production.up.railway.app/`

## Endpoints

- [POST `/create`](#endpoint-create)
- [POST `/release`](#endpoint-release)
- [POST `/handleDispute`](#endpoint-handledispute)
- [POST `/cancel`](#endpoint-cancel)
- [POST `/refund`](#endpoint-refund)
- [GET `/invoices/{id}`](#endpoint-invoicesid)

---

### Endpoint: `/create`

- **Method**: POST
- **Description**: Creates one or more on-chain invoices using the `AdvancedPaymentProcessor` contract's `createSingleInvoice` or `createMetaInvoice` functions.
- **Request Body**:
  ```json
  [
    {
      "orderId": "inv001",
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
| `orderId`               | string  | ✅       | Client-side identifier for the invoice (e.g., "inv001")                        |
| `seller`                | string  | ✅       | Ethereum address of the seller (e.g., `0xabc123...`)                           |
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
  - For a single invoice:
    ```json
    {
      "status": "success",
      "url": "https://contract-api-production.up.railway.app/<token>",
      "response": {
        "key": "0x123...",
        "orderId": "0xabc123..."
      }
    }
    ```
  - For multiple invoices:
    ```json
    {
      "status": "success",
      "url": "https://contract-api-production.up.railway.app/<token>",
      "response": {
        "key": "0x123...",
        "orders": {
          "0xdef456...": {
            "seller": "0xabc123...",
            "sub_order_ids": {
              "inv001": "0x789abc..."
            }
          }
        }
      }
    }
    ```
  - `status`: Indicates the transaction was successful.
  - `url`: A checkout URL with a generated token for the created invoice(s).
  - `response`: Contains the invoice key and either `orderId` (single invoice) or `orders` (multiple invoices with seller and sub-order IDs).
- **Error (400)**:
  ```json
  {
    "error": "invalid request body",
    "reason": "<decoding error message>"
  }
  ```
  - Returned if the JSON request body cannot be decoded (e.g., malformed JSON or incorrect field types).
- **Error (400)**:
  ```json
  {
    "error": "no invoice parameters provided",
    "reason": "invoice array is empty"
  }
  ```
  - Returned if the invoice array is empty.
- **Error (500)**:
  ```json
  {
    "error": "error processing invoice",
    "reason": "<blockchain error message>"
  }
  ```
  - Returned if the blockchain transaction fails (e.g., invalid parameters, insufficient gas, or contract reversion).
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
- **Description**: Releases escrow funds for a specific invoice using the `AdvancedPaymentProcessor` contract's `releasePayment` function.
- **Request Body**:
  ```json
  {
    "id": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1"
  }
  ```

#### Field Details

| Field | Type   | Required | Description                                                   |
| ----- | ------ | -------- | ------------------------------------------------------------- |
| `id`  | string | ✅       | Invoice ID as a 66-character hex string (e.g., `0xabc123...`) |

**Notes**:

- The `id` must be a valid 66-character hex string starting with `0x`, representing the invoice ID.
- The transaction hash is returned with a Ethereum Sepolia explorer URL prefix (`https://sepolia.etherscan.io/tx/`).

**Response**:

- **Success (200)**:
  ```json
  {
    "status": "success",
    "hash": "https://sepolia.etherscan.io/tx/0x123456..."
  }
  ```
  - `status`: Indicates the transaction was successful.
  - `hash`: Transaction hash with the Ethereum Sepolia explorer URL.
- **Error (400)**:
  ```json
  {
    "error": "invalid request body",
    "reason": "<decoding error message>"
  }
  ```
  - Returned if the JSON request body cannot be decoded (e.g., malformed JSON or missing `id`).
- **Error (500)**:
  ```json
  {
    "error": "Error sending transaction",
    "reason": "<blockchain error message>"
  }
  ```
  - Returned if the blockchain transaction fails (e.g., invalid `id` or contract reversion).

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/release \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "id": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1"
}'
```

---

### Endpoint: `/handleDispute`

- **Method**: POST
- **Description**: Resolves a dispute for a specific invoice using the `AdvancedPaymentProcessor` contract's `handleDispute` function.
- **Request Body**:
  ```json
  {
    "orderId": "inv001",
    "resolution": 2,
    "resolver": "0x789abc...",
    "sellersShare": "9000"
  }
  ```

#### Field Details

| Field          | Type    | Required                                    | Description                                                             |
| -------------- | ------- | ------------------------------------------- | ----------------------------------------------------------------------- |
| `orderId`      | string  | ✅                                          | Invoice ID (hex string starting with `0x` or client-side identifier)    |
| `resolution`   | integer | ✅                                          | Enum value specifying the action type (see MarketplaceAction below)     |
| `resolver`     | string  | ❌ Optional                                 | Ethereum address of the resolver (optional)                             |
| `sellersShare` | string  | ❌ Only if `resolution = 3` (SettleDispute) | Seller's share in **basis points** (e.g., `10000` = 100%, `9000` = 90%) |

#### MarketplaceAction Enum (`resolution`)

| Value | Name           | Description                          | Contract Function                       |
| ----- | -------------- | ------------------------------------ | --------------------------------------- |
| `0`   | Pending        | Default state, no action taken       | N/A                                     |
| `1`   | Release        | Release funds to the seller          | `releasePayment`                        |
| `2`   | DismissDispute | Dismiss an active dispute            | `handleDispute` with `DISPUTEDISMISSED` |
| `3`   | SettleDispute  | Resolve a dispute by splitting funds | `handleDispute` with `DISPUTESETTLED`   |

**Response**:

- **Success (200)**:
  ```json
  {
    "status": "success",
    "hash": "https://sepolia.etherscan.io/tx/0x123456..."
  }
  ```
  - `status`: Indicates the transaction was successful.
  - `hash`: Transaction hash with the Ethereum Sepolia explorer URL.
- **Error (400)**:
  ```json
  {
    "error": "invalid request body",
    "reason": "<decoding error message>"
  }
  ```
  - Returned if the JSON request body cannot be decoded (e.g., malformed JSON, missing `orderId` or `resolution`).
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
  - Returned if the blockchain transaction fails (e.g., invalid `orderId`, unsupported `resolution`, or contract reversion).

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/handleDispute \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "orderId": "inv001",
  "resolution": 2,
  "resolver": "0x789abc...",
  "sellersShare": "9000"
}'
```

---

### Endpoint: `/cancel`

- **Method**: POST
- **Description**: Cancels a specific invoice using the `AdvancedPaymentProcessor` contract's `cancelInvoice` function.
- **Request Body**:
  ```json
  {
    "id": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1"
  }
  ```

#### Field Details

| Field | Type   | Required | Description                                                   |
| ----- | ------ | -------- | ------------------------------------------------------------- |
| `id`  | string | ✅       | Invoice ID as a 66-character hex string (e.g., `0xabc123...`) |

**Notes**:

- The `id` must be a valid 66-character hex string starting with `0x`, representing the invoice ID.
- The transaction hash is returned with a Ethereum Sepolia explorer URL prefix (`https://sepolia.etherscan.io/tx/`).

**Response**:

- **Success (200)**:
  ```json
  {
    "status": "success",
    "hash": "https://sepolia.etherscan.io/tx/0x123456..."
  }
  ```
  - `status`: Indicates the transaction was successful.
  - `hash`: Transaction hash with the Ethereum Sepolia explorer URL.
- **Error (400)**:
  ```json
  {
    "error": "invalid request body",
    "reason": "<decoding error message>"
  }
  ```
  - Returned if the JSON request body cannot be decoded (e.g., malformed JSON or missing `id`).
- **Error (500)**:
  ```json
  {
    "error": "Error sending transaction",
    "reason": "<blockchain error message>"
  }
  ```
  - Returned if the blockchain transaction fails (e.g., invalid `id` or contract reversion).

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/cancel \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "id": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1"
}'
```

---

### Endpoint: `/refund`

- **Method**: POST
- **Description**: Claims a refund for an expired invoice using the `AdvancedPaymentProcessor` contract's `claimExpiredInvoiceRefunds` function.
- **Request Body**:
  ```json
  {
    "id": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1"
  }
  ```

#### Field Details

| Field | Type   | Required | Description                                                   |
| ----- | ------ | -------- | ------------------------------------------------------------- |
| `id`  | string | ✅       | Invoice ID as a 66-character hex string (e.g., `0xabc123...`) |

**Notes**:

- The `id` must be a valid 66-character hex string starting with `0x`, representing the invoice ID.
- The transaction hash is returned with a Ethereum Sepolia explorer URL prefix (`https://sepolia.etherscan.io/tx/`).

**Response**:

- **Success (200)**:
  ```json
  {
    "status": "success",
    "hash": "https://sepolia.etherscan.io/tx/0x123456..."
  }
  ```
  - `status`: Indicates the transaction was successful.
  - `hash`: Transaction hash with the Ethereum Sepolia explorer URL.
- **Error (400)**:
  ```json
  {
    "error": "invalid request body",
    "reason": "<decoding error message>"
  }
  ```
  - Returned if the JSON request body cannot be decoded (e.g., malformed JSON or missing `id`).
- **Error (500)**:
  ```json
  {
    "error": "Error sending transaction",
    "reason": "<blockchain error message>"
  }
  ```
  - Returned if the blockchain transaction fails (e.g., invalid `id` or contract reversion).

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/refund \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "id": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1"
}'
```

**Example**:

```bash
curl "https://contract-api-production.up.railway.app/invoices/inv001?seller=0xabc123...&status=Pending"
```

---

## Notes

- All endpoints are routed through a multiplexer (`http.ServeMux`) defined in the `Route` function.
- The `/create`, `/release`, `/handleDispute`, `/cancel`, and `/refund` endpoints are protected by `AccessControlMiddleWare` for authentication/authorization using the `X-API-KEY` header.
- The `/create` endpoint generates a checkout URL with a token derived from the invoice key in the transaction logs.
- The `/handleDispute` endpoint supports `Release`, `DismissDispute`, and `SettleDispute` actions, with appropriate contract function calls (`releasePayment` or `handleDispute`).
- The `/invoices/{id}` endpoint queries The Graph using a Keccak256-hashed `orderId` extracted from the URL path, with optional filtering by `seller`, `buyer`, or `status`.
- Prices are in USD with 8 decimal places and converted to token amounts on-chain using an oracle.
- All blockchain interactions occur on the Ethereum Sepolia testnet, with transaction hashes linked to `https://sepolia.etherscan.io/tx/`.
- The `IAdvancedPaymentProcessorInvoiceCreationParam` struct defines the structure for invoice creation parameters, ensuring type safety for `orderId` (string), `seller` and `buyer` (Ethereum addresses), `invoiceExpiryDuration`, `timeBeforeCancelation`, and `releaseWindow` (uint32), and `price` (big.Int).
