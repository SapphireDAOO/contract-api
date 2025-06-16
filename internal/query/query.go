package query

const userQuery = `
query ($address: String!, $first: Int = 10, $skip: Int = 0) {
  user(id: $address) {
    issuedInvoices(first: $first, skip: $skip) {
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
    receivedInvoices(first: $first, skip: $skip) {
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
    metaInvoices(first: $first, skip: $skip) {
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

const invoiceQuery = `
  query($id: String!, $first: Int = 10, $skip: Int = 0){
    smartInvoice(id: $id){
      invoiceId
      state
      escrow
      paymentToken
      amountPaid
      price
      paidAt
      buyer {
        id
      }
      seller {
        id
      }
      metaInvoice {
        id
      }
    }
  }
`
