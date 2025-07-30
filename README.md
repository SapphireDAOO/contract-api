# Sapphire Contract API – REST Endpoints

This API provides HTTP endpoints for interacting with the Sapphire DAO's `AdvancedPaymentProcessor` smart contract on the Ethereum Sepolia testnet at address `0x1A1b771B7e6cE617d22A148d08d0395Ca29f208a`. It facilitates invoice creation, payment processing, escrow management, dispute resolution, cancellation, and refund operations for secure and decentralized transactions.

**Base URL**: `https://contract-api-production.up.railway.app/`

## Endpoints

- [POST `/create`](#endpoint-create)
- [POST `/release`](#endpoint-release)
- [POST `/handleDispute`](#endpoint-handledispute)
- [POST `/cancel`](#endpoint-cancel)
- [POST `/refund`](#endpoint-refund)
- [GET `/invoices/{id}`](#endpoint-getinvoicedata)

---

### Endpoint: `/create`

- **Method**: POST
- **Description**: Creates one or more on-chain invoices using the `AdvancedPaymentProcessor` contract's `createSingleInvoice` or `createMetaInvoice` functions. Invoices are stored in the contract's `invoice` mapping for single invoices or `metaInvoice` mapping for multiple invoices, with unique IDs.

#### **Request Body**

```json
[
  {
    "orderId": "inv001",
    "seller": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
    "price": 8680000000
  }
]
```

#### Field Details

| Field     | Type    | Required | Description                                                                     |
| --------- | ------- | -------- | ------------------------------------------------------------------------------- |
| `orderId` | string  | ✅       | Client-side identifier for the invoice (e.g., "inv001").                        |
| `seller`  | string  | ✅       | Ethereum address of the seller (e.g., `0xabc123...`).                           |
| `price`   | integer | ✅       | Invoice price in **USD** with **8 decimal places** (e.g., `100000000` = $1.00). |

#### **Response**

**Success (200)**:

- For a single invoice:
  ```json
  {
    "url": "https://contract-api-production.up.railway.app/<token>",
    "orderId": "order-1",
    "invoiceId": "0x6ab58dbc048ca6b6fb46e4972b8d6f7515b5728c6c7e4b61c6aba4dd3e013844"
  }
  ```
- For multiple invoices:
  ```json
  {
    "url": "https://contract-api-production.up.railway.app/<token>",
    "orders": {
      "order-id-1": {
        "seller": "0x329C3E1bEa46Abc22F307eE30Cbb522B82Fe7082",
        "invoiceId": "0xb0ebc31f44ecffabd39b8be38e77992b4e681a3848641681c15a3730a30e383f"
      },
      "order-id-2": {
        "seller": "0x60D7dD3b4248D53Abba8DA999B22023656A2E4B3",
        "invoiceId": "0xce83a6685d3e54d85861fc842ea16f531a9f65542115ed1a55164f6c468394ab"
      }
    }
  }
  ```

**Notes**:

- The `price` is specified in USD with 8 decimal places and converted to token amounts using Chainlink price feeds via the contract's `getTokenValueFromUsd` function.
- A single invoice triggers `createSingleInvoice`, emitting an `InvoiceCreated` event. Multiple invoices trigger `createMetaInvoice`, emitting a `MetaInvoiceCreated` event and storing sub-invoices in the `metaInvoice` mapping.
- The contract enforces that `price` cannot be zero and checks for existing invoices to prevent duplicates.
- Only the `marketplace` address, set during contract initialization, can call this function.

**Error Responses**:

**Error (400)**:

```json
{
  "error": "invalid request body",
  "reason": "<decoding error message>"
}
```

- Returned for malformed JSON or incorrect field types.

**Error (400)**:

```json
{
  "error": "no invoice parameters provided",
  "reason": "invoice array is empty"
}
```

- Returned if the invoice array is empty.

**Error (500)**:

```json
{
  "error": "error processing invoice",
  "reason": "<blockchain error message>"
}
```

- Returned if the blockchain transaction fails (e.g., `InvoiceAlreadyExists`, `PriceCannotBeZero`, or `NotAuthorized` errors).

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/create \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '[
  {
    "orderId": "inv001",
    "seller": "0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4",
    "price": 8680000000
  }
]'
```

---

### Endpoint: `/release`

- **Method**: POST
- **Description**: Releases escrow funds for a specific invoice using the `AdvancedPaymentProcessor` contract's `release` function, distributing funds between the seller and buyer based on the specified `sellerShare`.

#### **Request Body**

```json
{
  "invoiceId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
  "sellerShare": 90000
}
```

#### Field Details

| Field         | Type    | Required | Description                                                              |
| ------------- | ------- | -------- | ------------------------------------------------------------------------ |
| `invoiceId`   | string  | ✅       | The invoice ID retrieved when the invoice is created                     |
| `sellerShare` | integer | ✅       | Seller's share in **basis points** (e.g., `10000` = 100%, `9000` = 90%). |

**Response**:

**Success (200)**:

```json
{
  "status": "success",
  "transaction url": "https://sepolia.etherscan.io/tx/0x123456..."
}
```

**Error (400)**:

```json
{
  "error": "invalid request body",
  "reason": "<decoding error message>"
}
```

**Error (500)**:

```json
{
  "error": "Error sending transaction",
  "reason": "<blockchain error message, e.g., InvalidInvoiceState or NotAuthorized>"
}
```

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/release \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "invoiceId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
  "sellerShare": "9000000000"
}'
```

---

### Endpoint: `/handleDispute`

- **Method**: POST
- **Description**: Resolves a dispute for a specific invoice using the `AdvancedPaymentProcessor` contract's `resolveDispute` or `handleDispute` functions, updating the invoice state and distributing funds if applicable.

#### **Request Body**

```json
{
  "invoiceId": "0x008d38d35c74aafbf1e0437e73a53a62e439a6479c566360215c587132ad5ee0",
  "resolution": 2,
  "resolver": "0x789abc...",
  "sellersShare": "9000"
}
```

#### Field Details

| Field          | Type    | Required                                    | Description                                                              |
| -------------- | ------- | ------------------------------------------- | ------------------------------------------------------------------------ |
| `invoiceId`    | string  | ✅                                          | The invoice ID retrieved when the invoice is created                     |
| `resolution`   | integer | ✅                                          | Enum value specifying the action type (see MarketplaceAction below).     |
| `sellersShare` | integer | ❌ Only if `resolution = 2` (SettleDispute) | Seller's share in **basis points** (e.g., `10000` = 100%, `9000` = 90%). |

#### MarketplaceAction Enum (`resolution`)

| Value | Name           | Description                                                                                            | Contract Function                           |
| ----- | -------------- | ------------------------------------------------------------------------------------------------------ | ------------------------------------------- |
| `0`   | Pending        | Default state, no action taken                                                                         | N/A                                         |
| `1`   | ResolveDispute | Both buyer and seller agree to dismiss the dispute, with escrow allocation unchanged (fully to seller) | `resolveDispute`                            |
| `2`   | SettleDispute  | Dispute resolved by an arbitrator, with `sellersShare` to the seller and the rest to the buyer         | `handleDispute` with `DISPUTE_SETTLED`      |
| `3`   | DismissDispute | Arbitrator dismisses the dispute, leaving escrow allocation unchanged (fully to seller)                | `handleDispute` with `DISPUTE_DISDISMISSED` |

**Notes**:

- The contract validates that the invoice is in the `DISPUTED` state and that `sellersShare` does not exceed `BASIS_POINTS` (10,000).
- For `SettleDispute`, funds are distributed via `_distributeFunds`, with platform fees applied. Emits `DisputeSettled` with payout details.
- For `DismissDispute`, emits `DisputeDismissed` without fund distribution.
- For `ResolveDispute`, emits `DisputeResolved` and sets the state to `DISPUTE_RESOLVED`.
- Only the `marketplace` address can call this function.

**Response**:

**Success (200)**:

```json
{
  "status": "success",
  "transaction url": "https://sepolia.etherscan.io/tx/0x123456..."
}
```

**Error (400)**:

```json
{
  "error": "invalid request body",
  "reason": "<decoding error message>"
}
```

**Error (500)**:

```json
{
  "error": "Error sending transaction",
  "reason": "<blockchain error message, e.g., InvalidInvoiceState or InvalidDisputeResolution>"
}
```

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/handleDispute \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "invoiceId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
  "resolution": 2,
  "resolver": "0x789abc...",
  "sellersShare": "9000"
}'
```

---

### Endpoint: `/cancel`

- **Method**: POST
- **Description**: Cancels a specific invoice using the `AdvancedPaymentProcessor` contract's `cancelInvoice` function, setting the invoice state to `CANCELED`. This can only be done before payment.

#### **Request Body**

```json
{
  "invoiceId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1"
}
```

#### Field Details

| Field       | Type   | Required | Description                                          |
| ----------- | ------ | -------- | ---------------------------------------------------- |
| `invoiceId` | string | ✅       | The invoice ID retrieved when the invoice is created |

**Response**:

- **Success (200)**:
  ```json
  {
    "status": "success",
    "transaction url": "https://sepolia.etherscan.io/tx/0x123456..."
  }
  ```

**Error (400)**:

```json
{
  "error": "invalid request body",
  "reason": "<decoding error message>"
}
```

**Error (500)**:

```json
{
  "error": "Error sending transaction",
  "reason": "<blockchain error message, e.g., InvalidInvoiceState or Unauthorized>"
}
```

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/cancel \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "invoiceId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1"
}'
```

---

### Endpoint: `/refund`

- **Method**: POST
- **Description**: Issues a refund for a specific invoice using the `AdvancedPaymentProcessor` contract's `refund` function, withdrawing funds from the escrow to the buyer.

#### **Request Body**

```json
{
  "invoiceId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
  "amount": 5000000000
}
```

#### Field Details

| Field       | Type    | Required | Description                                                               |
| ----------- | ------- | -------- | ------------------------------------------------------------------------- |
| `invoiceId` | string  | ✅       | The invoice ID retrieved when the invoice is created                      |
| `amount`    | integer | ✅       | Refund amount in USD with 8 decimal places (e.g., `5000000000` = $50.00). |

**Response**:

**Success (200)**:

```json
{
  "status": "success",
  "transaction url": "https://sepolia.etherscan.io/tx/0x123456..."
}
```

**Error (400)**:

```json
{
  "error": "invalid request body",
  "reason": "<decoding error message>"
}
```

**Error (500)**:

```json
{
  "error": "Error sending transaction",
  "reason": "<blockchain error message, e.g., InvalidInvoiceState or InsufficientBalance>"
}
```

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/refund \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE" \
-d '{
  "invoiceId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
  "amount": 5000000000
}'
```

---

### Endpoint: `/invoices/{invoiceId}`

- **Method**: GET
- **Description**: Retrieves invoice data from the `AdvancedPaymentProcessor` contract's `getInvoice` function using the invoice ID.

#### **Request**

- **URL Path**: `/invoices/{invoiceId}`
  - `invoiceId`: Invoice ID retrieved when the invoice created

**Response**:

- **Success (200)**:
  ```json
  {
    "invoiceId": "0x008d38d35c74aafbf1e0437e73a53a62e439a6479c566360215c587132ad5ee0",
    "createdAt": "1753051512",
    "price": "35000000000",
    "state": "PAID",
    "amountPaid": "350000000",
    "paidAt": "1753052436",
    "balance": "350000000",
    "releasedAt": "",
    "buyer": "0x0f447989b14a3f0bbf08808020ec1a6de0b8cbc4",
    "seller": "0x329c3e1bea46abc22f307ee30cbb522b82fe7082",
    "metaInvoice": "",
    "paymentToken": "Mock Usdc"
  }
  ```

**Error (400)**:

```json
{
  "error": "invalid URL path structure",
  "reason": "Invalid or missing invoice ID"
}
```

**Error (400)**:

```json
{
  "error": "empty invoice id in path",
  "reason": "Missing invoiceId parameter"
}
```

**Error (500)**:

```json
{
  "error": "failed to fetch invoice data",
  "reason": "<query error message, e.g., InvoiceDoesNotExist>"
}
```

**Example**:

```bash
curl -X GET https://contract-api-production.up.railway.app/invoices/inv001 \
-H "Content-Type: application/json" \
-H "X-API-KEY: YOUR_API_KEY_HERE"
```

---

## Notes

- All endpoints are routed through an `http.ServeMux` multiplexer defined in the `Route` function of the Go backend.
- Endpoints are protected by `AccessControlMiddleWare`, requiring an `X-API-KEY` header for authentication/authorization.
- The `AdvancedPaymentProcessor` contract uses Solady libraries (`SafeTransferLib`, `SafeCastLib`, `FixedPointMathLib`) for secure token transfers, type casting, and fixed-point arithmetic.
- Invoice states are tracked with the following constants: `INITIATED` (1), `PAID` (2), `REFUNDED` (3), `CANCELED` (4), `DISPUTED` (5), `DISPUTE_RESOLVED` (6), `DISPUTE_DISMISSED` (7), `DISPUTE_SETTLED` (8), `RELEASED` (9).
- The contract uses Chainlink price feeds (`AggregatorV3Interface`) for USD-to-token conversions, supporting both native (ETH) and ERC20 tokens.
- Escrow contracts are created via the `EscrowFactory` parent contract, with funds managed securely using `IEscrow.withdraw`.
- Platform fees are applied during fund distribution, transferred to the fee receiver address stored in the `IPaymentProcessorStorage` contract.
- The `marketplace` address, set via `setMarketplace`, controls privileged operations (`createSingleInvoice`, `createMetaInvoice`, `release`, `handleDispute`, `refund`, `resolveDispute`).
- All blockchain interactions occur on the Ethereum Sepolia testnet, with transaction hashes linked to `https://sepolia.etherscan.io/tx/`.
- The `IAdvancedPaymentProcessorInvoiceCreationParam` struct in the Go backend ensures type safety for `orderId` (string), `seller` (Ethereum address), `invoiceExpiryDuration`, `timeBeforeCancelation`, and `releaseWindow` (uint32), and `price` (big.Int).
