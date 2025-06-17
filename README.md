# Sapphire Contract API – REST Endpoints

This API exposes selected smart contract functions through HTTP endpoints for invoice creation, escrow resolution, and invoice data retrieval. It interacts with the Sapphire DAO's blockchain contracts on the Polygon Amoy testnet.

**Base URL**: `https://contract-api-production.up.railway.app/`

## Endpoints

- [POST `/create-invoice`](#endpoint-create-invoice)
- [POST `/release`](#endpoint-release)
- [GET `/invoice-data`](#endpoint-invoice-data)

---

### Endpoint: `/create-invoice`

- **Method**: POST
- **Description**: Creates one or more on-chain invoices using the `AdvancedPaymentProcessor` contract.
- **Request Body**:
  ```json
  [
    {
      "Seller": "0xabc123...",
      "Buyer": "0xdef456...",
      "InvoiceExpiryDuration": 3600,
      "TimeBeforeCancelation": 1800,
      "ReleaseWindow": 900,
      "Price": "100000000"
    }
  ]
  ```

#### Field Details

| Field                   | Type    | Required | Description                                                                    |
| ----------------------- | ------- | -------- | ------------------------------------------------------------------------------ |
| `Seller`                | string  | ✅       | Ethereum address of the seller (e.g., `0xabc123...`)                           |
| `Buyer`                 | string  | ✅       | Ethereum address of the buyer (e.g., `0xdef456...`)                            |
| `InvoiceExpiryDuration` | integer | ✅       | Invoice expiration time in seconds (e.g., `3600` for 1 hour)                   |
| `TimeBeforeCancelation` | integer | ✅       | Time window (in seconds) before the buyer can cancel the invoice               |
| `ReleaseWindow`         | integer | ✅       | Time window (in seconds) the seller has to release funds after payment         |
| `Price`                 | string  | ✅       | Invoice price in **USD** with **8 decimal places** (e.g., `100000000` = $1.00) |

**Notes**:

- `Price` is specified in USD with 8 decimal places (e.g., `100000000` = $1.00, `2500000000` = $25.00).
- The price is converted to the payment token amount using on-chain oracle pricing.
- Multiple invoices can be created in a single request by including multiple objects in the array.

**Response**:

- **Success (200)**:
  ```json
  {
    "url": "https://sapphire-dao-website-six.vercel.app/checkout/?data=<token>"
  }
  ```
  - `url`: A checkout URL with a generated token for the created invoice(s).
- **Error (400)**:
  ```json
  {
    "error": "invalid request body"
  }
  ```
  - Returned if the request body is malformed.
- **Error (400)**:
  ```json
  {
    "error": "Error sending transaction: <reason>"
  }
  ```
  - Returned if the blockchain transaction fails.
- **Error (500)**:
  ```json
  {
    "error": "failed to generate token"
  }
  ```
  - Returned if token generation fails.

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/create-invoice \
-H "Content-Type: application/json" \
-d '[
  {
    "Seller": "0xabc123...",
    "Buyer": "0xdef456...",
    "InvoiceExpiryDuration": 3600,
    "TimeBeforeCancelation": 1800,
    "ReleaseWindow": 900,
    "Price": "100000000"
  }
]'
```

---

### Endpoint: `/release`

- **Method**: POST
- **Description**: Releases escrow funds or resolves a dispute for a specific invoice using the `AdvancedPaymentProcessor` contract.
- **Request Body**:
  ```json
  {
    "orderId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
    "resolution": 1,
    "sellersShare": "9000"
  }
  ```

#### Field Details

| Field          | Type    | Required                                    | Description                                                                |
| -------------- | ------- | ------------------------------------------- | -------------------------------------------------------------------------- |
| `orderId`      | string  | ✅                                          | 32-byte hex string representing the invoice/order ID (e.g., `0xabc123...`) |
| `resolution`   | integer | ✅                                          | Enum value specifying the action type (see MarketplaceAction below)        |
| `sellersShare` | string  | ❌ Only if `resolution = 3` (SettleDispute) | Seller's share in **basis points** (e.g., `10000` = 100%, `6500` = 65%)    |

#### MarketplaceAction Enum (`resolution`)

| Value | Name           | Description                          |
| ----- | -------------- | ------------------------------------ |
| `0`   | Pending        | Default state, no action taken       |
| `1`   | Release        | Release funds to the seller          |
| `2`   | DismissDispute | Dismiss an active dispute            |
| `3`   | SettleDispute  | Resolve a dispute by splitting funds |

**Response**:

- **Success (200)**:
  ```json
  {
    "status": "success",
    "hash": "0x123456..."
  }
  ```
  - `status`: Indicates the transaction was successful.
  - `hash`: Transaction hash of the escrow release or dispute resolution.
- **Error (400)**:
  ```json
  {
    "error": "Invalid request body"
  }
  ```
  - Returned if the request body is malformed.
- **Error (400)**:
  ```json
  {
    "error": "Error sending transaction: <reason>"
  }
  ```
  - Returned if the blockchain transaction fails.

**Example**:

```bash
curl -X POST https://contract-api-production.up.railway.app/release \
-H "Content-Type: application/json" \
-d '{
  "orderId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
  "resolution": 1,
  "sellersShare": "9000"
}'
```

---

### Endpoint: `/invoice-data`

- **Method**: GET
- **Description**: Retrieves invoice data for a given wallet address or invoice ID from The Graph.

#### Query Parameters

| Name    | Type    | Required | Description                                               |
| ------- | ------- | -------- | --------------------------------------------------------- |
|`orderId`| string  | ✅       | orderId ID to fetch associated invoices |
| `first` | integer | ❌       | Number of invoices to fetch (default: `10`)               |
| `skip`  | integer | ❌       | Number of invoices to skip for pagination (default: `0`)  |

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
  - Returns an array of invoices associated with the provided `orderId`.
- **Error (400)**:
  ```json
  {
    "error": "Missing orderId parameter"
  }
  ```
  - Returned if the `orderId` query parameter is missing.
- **Error (500)**:
  ```json
  {
    "error": "failed to fetch invoice data: <reason>"
  }
  ```
  - Returned if the query to The Graph fails.

**Example**:

```bash
curl "https://contract-api-production.up.railway.app/invoice-data?orderId=0xabc123...&first=5&skip=0"
```

---

## Notes

- All endpoints are protected by an `AccessControlMiddleWare` for authentication/authorization (implementation details not provided).
- The `/create-invoice` endpoint generates a checkout URL with a token for the frontend to process payments.
- The `/release` endpoint supports multiple actions (`Release`, `DismissDispute`, `SettleDispute`) based on the `resolution` value.
- The `/invoice-data` endpoint supports pagination via `first` and `skip` parameters, defaulting to 10 invoices with no offset.
- Prices are in USD with 8 decimal places and converted to token amounts on-chain using an oracle.
- All blockchain interactions occur on the Polygon Amoy testnet.

---

## Error Handling

- **400 Bad Request**: Invalid request body or missing required parameters.
- **500 Internal Server Error**: Failures in blockchain transactions, token generation, or GraphQL queries.

---
