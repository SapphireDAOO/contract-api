# 📘 Sapphire Contract API – REST Endpoints

This API exposes selected smart contract functions through HTTP endpoints for invoice creation, escrow resolution, and token removal.

---

### Endpoint: `/createInvoice`

- **Method:** `POST`
- **Description:** Creates one or more on-chain invoices.
- **Request Body:**

  ```json
  [
    {
      "Seller": "0xabc123...",
      "Buyer": "0xdef456...",
      "InvoiceExpiryDuration": 3600,
      "TimeBeforeCancelation": 1800,
      "ReleaseWindow": 900,
      "Price": "1000000000000000000"
    }
  ]
  ```

  **Note:**

- `Price` must be in **USD**, with **8 decimal places**.

For example:

- `100000000` = **$1.00**
- `2500000000` = **$25.00**

This value will be converted to the payment token amount using on-chain oracle pricing.

### Endpoint: `/release`

- **Method:** POST
- **Description:** Releases escrow or resolves a dispute for a specific invoice.
- **Request Body:**

```json
{
  "orderId": "0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1",
  "resolution": 1,
  "sellersShare": "900000000000000000"
}
```

#### Field Details

| Field          | Type    | Required                                    | Description                                                            |
| -------------- | ------- | ------------------------------------------- | ---------------------------------------------------------------------- |
| `orderId`      | string  | ✅                                          | 32-byte hex string representing the invoice/order ID                   |
| `resolution`   | integer | ✅                                          | Enum value specifying the type of action (see MarketplaceAction below) |
| `sellersShare` | string  | ❌ Only if `resolution = 4` (SettleDispute) | Seller's share in **basis points** (e.g., `10000` = 100%, `6500` = 65%)             |

#### MarketplaceAction Enum (`resolution`)

| Value | Name           | Description                            |
| ----- | -------------- | -------------------------------------- |
| `0`   | Pending        | Default state, no action taken         |
| `1`   | Release        | Release funds to the seller            |
| `2`   | Refund         | Refund funds to the buyer              |
| `3`   | DismissDispute | Dismiss an active dispute              |
| `4`   | SettleDispute  | Resolve a dispute by splitting funds   |
| `5`   | ResolveDispute | Resolve a dispute without custom share |
