package query

const invoiceQuery = `
query ($address: String!) {
  user(id: $address) {
    ownedSmartInvoices {
      invoiceId
      amountPaid
      price
      state
      createdAt
      paidAt
      releasedAt
      buyer {
        id
      }
      metaInvoice {
        invoiceId
      }
    }

    paidSmartInvoices {
      invoiceId
      amountPaid
      price
      state
      createdAt
      paidAt
      releasedAt
      seller {
        id
      }
      metaInvoice {
        invoiceId
      }
    }

    metaInvoices {
      invoiceId
      price
      invoices {
        invoiceId
        amountPaid
        price
        state
        paidAt
        releasedAt
      }
    }
  }
}
`
