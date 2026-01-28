# MTN Mobile Money (MoMo) SDK for Go

A comprehensive Golang client for the MTN Mobile Money API. This SDK provides a simple interface to interact with MTN MoMo for Collections, Disbursements, Remittances, and Refunds.

## Features

- **Collections API**
  - Request-to-Pay (Customer payment requests)
  - Transaction status checking
  - Account balance inquiry
  - Account holder validation
  - Basic user information retrieval

- **Disbursements API**
  - Transfer funds to customers
  - Transfer status checking
  - Account balance inquiry
  - Account holder validation
  - Basic user information retrieval

- **Remittances API**
  - International money transfers
  - Remittance status checking
  - Account balance inquiry
  - Account holder validation
  - Basic user information retrieval

- **Refunds API**
  - Refund completed transactions (V1 & V2)
  - Refund status checking

- **Authentication & Token Management**
  - Automatic OAuth2 token requests
  - Token caching and refresh
  - Separate tokens for Collections, Disbursements, and Remittances

## Prerequisites

To use this SDK, you'll need:

1. Sign up with [MTN MoMo Developer Portal](https://momodeveloper.mtn.com/)
2. Subscribe to the required products (Collections, Disbursements, Remittances)
3. Create API User and API Key
4. You'll need the following credentials:
   - API User (UUID)
   - API Key
   - Subscription Key (Primary or Secondary)

## Installation

Install with the go get command:

```bash
go get github.com/nutcas3/payment-rails/momo
```

Then import it to your package:

```go
package main

import (
    "github.com/nutcas3/payment-rails/momo"
)
```

## Usage

### Initialize the Client

```go
// Initialize the MTN MoMo client with flexible product configuration
client, err := momo.New(momo.ClientConfig{
    Environment:                 string(momo.SANDBOX), // or momo.PRODUCTION
    APIKey:                      "your-api-key",
    APISecret:                   "your-api-secret", 
    CollectionSubscriptionKey:   "your-collection-subscription-key",   // Optional: only if using Collection
    DisbursementSubscriptionKey: "your-disbursement-subscription-key", // Optional: only if using Disbursement
    RemittanceSubscriptionKey:   "your-remittance-subscription-key",   // Optional: only if using Remittance
})
if err != nil {
    log.Fatalf("Failed to initialize MTN MoMo client: %v", err)
}

// Use Collection API if subscription key was provided
if client.Collection != nil {
    balance, err := client.Collection.GetAccountBalance(context.Background())
    if err != nil {
        log.Printf("Failed to get collection balance: %v", err)
    } else {
        fmt.Printf("Collection Balance: %s %s\n", balance.AvailableBalance, balance.Currency)
    }
}

// Use Disbursement API if subscription key was provided
if client.Disbursement != nil {
    balance, err := client.Disbursement.GetAccountBalance(context.Background())
    if err != nil {
        log.Printf("Failed to get disbursement balance: %v", err)
    } else {
        fmt.Printf("Disbursement Balance: %s %s\n", balance.AvailableBalance, balance.Currency)
    }
}

// Use Remittance API if subscription key was provided
if client.Remittance != nil {
    balance, err := client.Remittance.GetAccountBalance(context.Background())
    if err != nil {
        log.Printf("Failed to get remittance balance: %v", err)
    } else {
        fmt.Printf("Remittance Balance: %s %s\n", balance.AvailableBalance, balance.Currency)
    }
}
```

## Examples

### Collections API

#### Request-to-Pay

```go
// Initiate a request-to-pay transaction using Collection service
ctx := context.Background()
refID := uuid.New()

err := client.Collection.RequestToPay(ctx, refID, "https://example.com/callback", true, momo.RequestToPayInput{
    Amount:       "100",
    Currency:     "EUR",
    ExternalID:   "123456789",
    Payer: momo.Payer{
        PartyIDType: "MSISDN",
        PartyID:     "256774290781",
    },
    PayerMessage: "Payment for services",
    PayeeNote:    "Thank you for your payment",
})
if err != nil {
    log.Fatalf("Request-to-pay failed: %v", err)
}

// Check transaction status
status, err := client.Collection.RequestToPayTransactionStatus(ctx, refID)
if err != nil {
    log.Fatalf("Failed to get transaction status: %v", err)
}
fmt.Printf("Transaction Status: %s\n", status.Status)
```

#### Check Transaction Status

```go
// Check the status of a request-to-pay transaction
ctx := context.Background()
status, err := client.Collection.RequestToPayTransactionStatus(ctx, refID)
if err != nil {
    log.Fatalf("Failed to get transaction status: %v", err)
}
fmt.Printf("Transaction Status: %s\n", status.Status)
fmt.Printf("Amount: %s %s\n", status.Amount, status.Currency)
fmt.Printf("Financial Transaction ID: %s\n", status.FinancialTransactionID)
```

#### Get Account Balance

```go
// Get collection account balance
ctx := context.Background()
balance, err := client.Collection.GetAccountBalance(ctx)
if err != nil {
    log.Fatalf("Failed to get account balance: %v", err)
}
fmt.Printf("Available Balance: %s %s\n", balance.AvailableBalance, balance.Currency)
```

#### Validate Account Holder

```go
// Validate if an account holder is active
ctx := context.Background()
active, err := client.Collection.ValidateAccountHolderStatus(ctx, "256774290781", "MSISDN")
if err != nil {
    log.Fatalf("Failed to validate account holder: %v", err)
}
if active {
    fmt.Println("Account holder is active")
} else {
    fmt.Println("Account holder is not active")
}
```

#### Get Basic User Info

```go
// Get basic user information
ctx := context.Background()
userInfo, err := client.Collection.GetBasicUserInfo(ctx, "MSISDN", "256774290781")
if err != nil {
    log.Fatalf("Failed to get user info: %v", err)
}
fmt.Printf("Name: %s %s\n", userInfo.GivenName, userInfo.FamilyName)
fmt.Printf("Gender: %s\n", userInfo.Gender)
```

### Disbursements API

Use the `client.Disbursement` service to access disbursement methods:

```go
// Transfer funds
ctx := context.Background()
refID := uuid.New()

err := client.Disbursement.Transfer(ctx, refID, "https://example.com/callback", types.TransferInput{
    Amount:       "100",
    Currency:     types.EUR,
    ExternalID:   "987654321",
    Payee: types.Party{
        PartyIDType: types.MSISDN,
        PartyID:     "256774290781",
    },
    PayerMessage: "Salary payment",
    PayeeNote:    "Your salary for January",
})

// Check transfer status
status, err := client.Disbursement.GetTransferStatus(ctx, refID)

// Get disbursement balance
balance, err := client.Disbursement.GetAccountBalance(ctx)
```

### Remittances API

Use the `client.Remittance` service to access remittance methods:

```go
// Send remittance
ctx := context.Background()
refID := uuid.New()

err := client.Remittance.Transfer(ctx, refID, "https://example.com/callback", types.CashTransferInput{
    Amount:       "500",
    Currency:     types.EUR,
    ExternalID:   "REM123456",
    Payee: types.Party{
        PartyIDType: types.MSISDN,
        PartyID:     "256774290781",
    },
    PayerMessage: "International transfer",
    PayeeNote:    "Money from abroad",
})

// Check remittance status
status, err := client.Remittance.GetTransferStatus(ctx, refID)

// Get remittance balance
balance, err := client.Remittance.GetAccountBalance(ctx)
```

### Refunds API

Use the `client.Disbursement` service to access refund methods:

```go
// Process refund (V1)
ctx := context.Background()
refID := uuid.New()

err := client.Disbursement.RefundV1(ctx, refID, "https://example.com/callback", types.RefundInput{
    Amount:              "100",
    Currency:            types.EUR,
    ExternalID:          "REF123456",
    PayerMessage:        "Refund for order #12345",
    PayeeNote:           "Refund processed",
    ReferenceIDToRefund: "original-transaction-reference-id",
})

// Check refund status
status, err := client.Disbursement.GetRefundStatus(ctx, refID)
```

## Transaction Status Codes

- **PENDING**: Transaction is being processed
- **SUCCESSFUL**: Transaction completed successfully
- **FAILED**: Transaction failed

## Error Handling

The SDK returns detailed error messages for failed operations:

```go
response, err := client.RequestToPay(request)
if err != nil {
    // Handle error
    log.Printf("Error: %v", err)
    return
}

// Check transaction status
status, err := client.GetRequestToPayStatus(response.ReferenceID)
if err != nil {
    log.Printf("Failed to get status: %v", err)
    return
}

if status.Status == "FAILED" {
    log.Printf("Transaction failed: %s", status.Reason)
}
```

## Environment Configuration

The SDK supports both sandbox and production environments:

```go
// Sandbox environment (for testing)
client, err := momo.New(momo.ClientConfig{
    Environment:                 string(momo.SANDBOX),
    APIKey:                      "your-api-key",
    APISecret:                   "your-api-secret",
    CollectionSubscriptionKey:   "your-collection-subscription-key",
})

// Production environment (for live transactions)
client, err := momo.New(momo.ClientConfig{
    Environment:                 string(momo.PRODUCTION),
    APIKey:                      "your-api-key",
    APISecret:                   "your-api-secret",
    CollectionSubscriptionKey:   "your-collection-subscription-key",
})
```

## API Endpoints

### Sandbox
- Base URL: `https://sandbox.momodeveloper.mtn.com`

### Production
- Base URL: `https://proxy.momoapi.mtn.com`

## Testing

The SDK has been tested end-to-end in the sandbox environment with the following workflows:
- ✅ Request-to-Pay
- ✅ Disbursements & Balance
- ✅ Remittances
- ✅ Refunds
- ✅ Transaction Status

To run tests:

```bash
go test ./...
```

## Security Best Practices

1. **Never hardcode credentials** - Use environment variables or secure configuration management
2. **Use HTTPS** - All API calls are made over HTTPS
3. **Validate callbacks** - Verify callback signatures when receiving webhook notifications
4. **Rotate API keys** - Regularly rotate your API keys and subscription keys
5. **Monitor transactions** - Implement logging and monitoring for all transactions

## Support

For issues and questions:
- MTN MoMo Developer Portal: [https://momodeveloper.mtn.com](https://momodeveloper.mtn.com)
- API Documentation: [https://momodeveloper.mtn.com/api-documentation](https://momodeveloper.mtn.com/api-documentation)

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
