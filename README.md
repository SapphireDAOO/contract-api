# Sapphire Contract API – REST Endpoints

This API provides HTTP endpoints for interacting with the Sapphire DAO's `AdvancedPaymentProcessor` smart contract on the Ethereum Sepolia testnet at address `0xb9DD118C880759E62516fa2c88Eb76ba3fd42eae`. It supports invoice creation, escrow management, dispute resolution, cancellation, and refund operations.

**Base URL**: `https://contract-api-production.up.railway.app/`

## Endpoints

- [POST `/create`](#endpoint-create)
- [POST `/release`](#endpoint-release)
- [POST `/handleDispute`](#endpoint-handledispute)
- [POST `/cancel`](#endpoint-cancel)
- [POST `/refund`](#endpoint-refund)

---

### Endpoint: `/create`

- **Method**: POST
- **Description**: Creates one or more on-chain invoices using the `AdvancedPaymentProcessor` contract's `createSingleInvoice` or `createMetaInvoice` functions.

## 📦 Invoice Examples

### 🧾 Single Invoice

In a single-item order, the API accepts an array with one invoice object.

#### **Request**

```json
[
  {
    "orderId": "88121%",
    "Seller": "0x60D7dD3b4248D53Abba8DA999B22023656A2E4B3",
    "InvoiceExpiryDuration": 864000,
    "TimeBeforeCancelation": 864000,
    "ReleaseWindow": 864000,
    "Price": 190000000000
  }
]
```

#### **Response**

```json
{
  "url": "https://sapphire-dao-website-six.vercel.app/checkout/?data=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbnZvaWNlS2V5IjoiMHg2YWI1OGRiYzA0OGNhNmI2ZmI0NmU0OTcyYjhkNmY3NTE1YjU3MjhjNmM3ZTRiNjFjNmFiYTRkZDNlMDEzODQ0In0.O7BojYAD3oo9HFuQvXmTBTbShz2NCVn5-fGD9iIA2M4",
  "order_id": "88121%",
  "hashed_order_id": "0x6ab58dbc048ca6b6fb46e4972b8d6f7515b5728c6c7e4b61c6aba4dd3e013844"
}
```

---

### 📦 Meta Invoice (Multiple Invoices)

When multiple items from one or more sellers are submitted in a single transaction, a _meta-invoice_ is created. Each seller receives a unique `hashed_order_id`, and each individual invoice under that seller is identified by a `hashed_sub_order_id`.

#### **Request**

```json
[
  {
    "orderId": "-1@#k1",
    "Seller": "0x60D7dD3b4248D53Abba8DA999B22023656A2E4B3",
    "InvoiceExpiryDuration": 864000,
    "TimeBeforeCancelation": 864000,
    "ReleaseWindow": 864000,
    "Price": 610000000000
  },
  {
    "orderId": "t-='1'",
    "Seller": "0x60D7dD3b4248D53Abba8DA999B22023656A2E4B3",
    "InvoiceExpiryDuration": 864000,
    "TimeBeforeCancelation": 864000,
    "ReleaseWindow": 864000,
    "Price": 15000000000
  },
  {
    "orderId": "t1'",
    "Seller": "0x329C3E1bEa46Abc22F307eE30Cbb522B82Fe7082",
    "InvoiceExpiryDuration": 864000,
    "TimeBeforeCancelation": 864000,
    "ReleaseWindow": 864000,
    "Price": 35000000000
  }
]
```

#### **Response**

```json
{
  "url": "https://sapphire-dao-website-six.vercel.app/checkout/?data=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbnZvaWNlS2V5IjoiMHg4N2YzNzM2OGIxNTE0MjUzZWE3MDJhZGQ3OThiMjU5OGQ1ODQxNjJhYzAwNjJlZWYwYjQ5YTlkZjk0MjhkZDA2In0.TvuMSZ4_sh2-v36FurszprTAU4RvVO4ffCsgpGnaO3o",
  "orders": {
    "0x329C3E1bEa46Abc22F307eE30Cbb522B82Fe7082": {
      "hashed_order_id": "0x09b04d922875a8a3188b87e5853c30cd2b38f003d4eaa5b3fb3e4dd679e3cbb1",
      "hashed_sub_order_ids": {
        "t1'": "0xce83a6685d3e54d85861fc842ea16f531a9f65542115ed1a55164f6c468394ab"
      }
    },
    "0x60D7dD3b4248D53Abba8DA999B22023656A2E4B3": {
      "hashed_order_id": "0x9ce028293da230f7fdd0497bd734cc70b8c77d4867ee173075213925d3d355ad",
      "hashed_sub_order_ids": {
        "-1@#k1": "0xb0ebc31f44ecffabd39b8be38e77992b4e681a3848641681c15a3730a30e383f",
        "t-='1'": "0xfb007570f437909c3580d112b836e68f59ca0d748a93d7524552b34834c56aa6"
      }
    }
  }
}
```

- **Request Body**:
  ```json
  [
    {
      "orderId": "inv001",
      "seller": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
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
- Prices are converted to payment token amounts using on-chain oracle pricing.
- A single invoice triggers `createSingleInvoice`; multiple invoices trigger `createMetaInvoice`.
- The endpoint returns a checkout URL with a token generated from the invoice key (derived from transaction logs).

**Response**:

- **Success (200)**:
  - For a single invoice:
    ```json
    {
      "url": "https://contract-api-production.up.railway.app/<token>",
      "orderId": "inv001",
      "hashed_order_id": "0x123..."
    }
    ```
  - For multiple invoices:
    ```json
    {
      "url": "https://contract-api-production.up.railway.app/<token>",
      "seller": {
        "0xabc123...": {
          "hashed_order_id": "0xdef456...",
          "hashed_sub_order_ids": {
            "inv001": "0x789abc..."
          }
        }
      }
    }
    ```
  - `url`: Checkout URL with a generated token for the invoice(s).
  - `key`: idenfier of the order in the contract.
  - `hashed_order_id`: identifier of the order on-chain
  - `orderId` (single invoice) or `seller` (multiple invoices): Contains order details, including hashed sub-order IDs (identifiers of each sub-invoice in a single order).
- **Error (400)**:
  ```json
  {
    "error": "invalid request body",
    "reason": "<decoding error message>"
  }
  ```
  - Returned for malformed JSON or incorrect field types.
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

| Field | Type   | Required | Description                                                               |
| ----- | ------ | -------- | ------------------------------------------------------------------------- |
| `id`  | string | ✅       | Represents the hashed_order_id, which is derived during invoice creation. |

**Notes**:

- The `id` must be a valid 66-character hex string starting with `0x`.
- The transaction hash is returned with an Ethereum Sepolia explorer URL prefix (`https://sepolia.etherscan.io/tx/`).

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
  - Returned for malformed JSON or missing `id`.
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
- **Description**: Resolves a dispute for a specific invoice using the `AdvancedPaymentProcessor` contract's `resolveDispute` or `handleDispute` functions.
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

| Field          | Type    | Required                                    | Description                                                               |
| -------------- | ------- | ------------------------------------------- | ------------------------------------------------------------------------- |
| `orderId`      | string  | ✅                                          | Represents the hashed_order_id, which is derived during invoice creation. |
| `resolution`   | integer | ✅                                          | Enum value specifying the action type (see MarketplaceAction below)       |
| `resolver`     | string  | ❌ Optional                                 | Ethereum address of the resolver (required for `ResolveDispute`)          |
| `sellersShare` | string  | ❌ Only if `resolution = 3` (SettleDispute) | Seller's share in **basis points** (e.g., `10000` = 100%, `9000` = 90%)   |

#### MarketplaceAction Enum (`resolution`)

| Value | Name           | Description                                                                                                                   | Contract Function                       |
| ----- | -------------- | ----------------------------------------------------------------------------------------------------------------------------- | --------------------------------------- |
| `0`   | Pending        | Default state, no action taken                                                                                                | N/A                                     |
| `1`   | ResolveDispute | Both the buyer and seller agreed to dismiss the dispute, with the escrow allocation remaining unchanged (fully to the seller) | `resolveDispute`                        |
| `2`   | SettleDispute  | The dispute was resolved by an arbitrator, with X% awarded to the seller and the remaining (100 − X)% to the buyer            | `handleDispute` with `DISPUTESETTLED`   |
| `3`   | DismissDispute | The arbitrator ruled to dismiss the dispute, leaving the escrow allocation unchanged (all to the seller)                      | `handleDispute` with `DISPUTEDISMISSED` |

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
  - Returned for malformed JSON, missing `orderId`, or invalid `resolution`.
- **Error (500)**:
  ```json
  {
    "error": "failed to generate order ID hash",
    "reason": "<hashing error message>"
  }
  ```
  - Returned if Keccak256 hashing of `orderId` fails.
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

| Field | Type   | Required | Description                                                                                                                                                                 |
| ----- | ------ | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id`  | string | ✅       | Represents the hashed_order_id (for a single-item order) or one of the hashed_sub_order_ids (for a multiple-item order), both of which are derived during invoice creation. |

**Notes**:

- The transaction hash is returned with an Ethereum Sepolia explorer URL prefix (`https://sepolia.etherscan.io/tx/`).

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
  - Returned for malformed JSON or missing `id`.
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

| Field | Type   | Required | Description                                                                                                                                                                 |
| ----- | ------ | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id`  | string | ✅       | Represents the hashed_order_id (for a single-item order) or one of the hashed_sub_order_ids (for a multiple-item order), both of which are derived during invoice creation. |

**Notes**:

- The `id` must be a valid 66-character hex string starting with `0x`.
- The transaction hash is returned with an Ethereum Sepolia explorer URL prefix (`https://sepolia.etherscan.io/tx/`).

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
  - Returned for malformed JSON or missing `id`.
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

---

## Notes

- All endpoints are routed through a multiplexer (`http.ServeMux`) defined in the `Route` function.
- The `/create`, `/release`, `/handleDispute`, `/cancel`, and `/refund` endpoints are protected by `AccessControlMiddleWare` for authentication/authorization using the `X-API-KEY` header.
- The `/create` endpoint generates a checkout URL with a token derived from the invoice key in the transaction logs.
- The `/handleDispute` endpoint supports `Release`, `ResolveDispute`, `SettleDispute`, and `DismissDispute` actions, with appropriate contract function calls (`releasePayment`, `resolveDispute`, or `handleDispute`).
- Prices are in USD with 8 decimal places and converted to token amounts on-chain using an oracle.
- All blockchain interactions occur on the Ethereum Sepolia testnet, with transaction hashes linked to `https://sepolia.etherscan.io/tx/`.
- The `IAdvancedPaymentProcessorInvoiceCreationParam` struct defines the structure for invoice creation parameters, ensuring type safety for `orderId` (string), `seller` (Ethereum addresses), `invoiceExpiryDuration`, `timeBeforeCancelation`, and `releaseWindow` (uint32), and `price` (big.Int).
