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
    "github.com/nutcas3/payment-rails/momo/pkg/api"
)
```

## Usage

### Initialize the Client

```go
// Initialize the MTN MoMo client
client, err := momo.NewClient(
    "your-api-user-uuid",           // API User from MTN Developer Portal
    "your-api-key",                 // API Key from MTN Developer Portal
    "your-subscription-key",        // Subscription Key from MTN Developer Portal
    momo.SANDBOX,                   // Environment (momo.SANDBOX or momo.PRODUCTION)
)
if err != nil {
    log.Fatalf("Failed to initialize MTN MoMo client: %v", err)
}

// Use custom HTTP client if needed
client.SetHttpClient(&http.Client{
    Timeout: 30 * time.Second,
})

// Get authentication token
token, err := client.GetCollectionToken()
if err != nil {
    log.Fatalf("Failed to get auth token: %v", err)
}
fmt.Printf("Auth Token: %s\n", token)
```

## Examples

### Collections API

#### Request-to-Pay

```go
// Initiate a request-to-pay transaction
response, err := client.RequestToPay(api.RequestToPayRequest{
    Amount:       "100",
    Currency:     "EUR",
    ExternalID:   "123456789",
    Payer: api.Payer{
        PartyIDType: "MSISDN",
        PartyID:     "256774290781",
    },
    PayerMessage: "Payment for services",
    PayeeNote:    "Thank you for your payment",
    CallbackURL:  "https://example.com/callback",
})
if err != nil {
    log.Fatalf("Request-to-pay failed: %v", err)
}
fmt.Printf("Reference ID: %s\n", response.ReferenceID)
fmt.Printf("Status: %s\n", response.Status)
```

#### Check Transaction Status

```go
// Check the status of a request-to-pay transaction
status, err := client.GetRequestToPayStatus(response.ReferenceID)
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
balance, err := client.GetAccountBalance()
if err != nil {
    log.Fatalf("Failed to get account balance: %v", err)
}
fmt.Printf("Available Balance: %s %s\n", balance.AvailableBalance, balance.Currency)
```

#### Validate Account Holder

```go
// Validate if an account holder is active
holder, err := client.ValidateAccountHolderStatus("256774290781")
if err != nil {
    log.Fatalf("Failed to validate account holder: %v", err)
}
if holder.Result {
    fmt.Println("Account holder is active")
} else {
    fmt.Println("Account holder is not active")
}
```

#### Get Basic User Info

```go
// Get basic user information
userInfo, err := client.GetBasicUserInfo("256774290781")
if err != nil {
    log.Fatalf("Failed to get user info: %v", err)
}
fmt.Printf("Name: %s %s\n", userInfo.GivenName, userInfo.FamilyName)
fmt.Printf("Gender: %s\n", userInfo.Gender)
```

### Disbursements API

#### Transfer Funds

```go
// Initiate a disbursement transfer
response, err := client.Transfer(api.TransferRequest{
    Amount:       "100",
    Currency:     "EUR",
    ExternalID:   "987654321",
    Payee: api.Payee{
        PartyIDType: "MSISDN",
        PartyID:     "256774290781",
    },
    PayerMessage: "Salary payment",
    PayeeNote:    "Your salary for January",
    CallbackURL:  "https://example.com/callback",
})
if err != nil {
    log.Fatalf("Transfer failed: %v", err)
}
fmt.Printf("Reference ID: %s\n", response.ReferenceID)
fmt.Printf("Status: %s\n", response.Status)
```

#### Check Transfer Status

```go
// Check the status of a transfer
status, err := client.GetTransferStatus(response.ReferenceID)
if err != nil {
    log.Fatalf("Failed to get transfer status: %v", err)
}
fmt.Printf("Transfer Status: %s\n", status.Status)
fmt.Printf("Amount: %s %s\n", status.Amount, status.Currency)
```

#### Get Disbursement Balance

```go
// Get disbursement account balance
balance, err := client.GetDisbursementBalance()
if err != nil {
    log.Fatalf("Failed to get disbursement balance: %v", err)
}
fmt.Printf("Available Balance: %s %s\n", balance.AvailableBalance, balance.Currency)
```

### Remittances API

#### Send Remittance

```go
// Initiate a remittance transfer
response, err := client.Remit(api.RemittanceRequest{
    Amount:       "500",
    Currency:     "EUR",
    ExternalID:   "REM123456",
    Payee: api.Payee{
        PartyIDType: "MSISDN",
        PartyID:     "256774290781",
    },
    PayerMessage: "International transfer",
    PayeeNote:    "Money from abroad",
    CallbackURL:  "https://example.com/callback",
})
if err != nil {
    log.Fatalf("Remittance failed: %v", err)
}
fmt.Printf("Reference ID: %s\n", response.ReferenceID)
fmt.Printf("Status: %s\n", response.Status)
```

#### Check Remittance Status

```go
// Check the status of a remittance
status, err := client.GetRemittanceStatus(response.ReferenceID)
if err != nil {
    log.Fatalf("Failed to get remittance status: %v", err)
}
fmt.Printf("Remittance Status: %s\n", status.Status)
fmt.Printf("Amount: %s %s\n", status.Amount, status.Currency)
```

#### Get Remittance Balance

```go
// Get remittance account balance
balance, err := client.GetRemittanceBalance()
if err != nil {
    log.Fatalf("Failed to get remittance balance: %v", err)
}
fmt.Printf("Available Balance: %s %s\n", balance.AvailableBalance, balance.Currency)
```

### Refunds API

#### Refund Transaction (V1)

```go
// Initiate a refund
response, err := client.Refund(api.RefundRequest{
    Amount:              "100",
    Currency:            "EUR",
    ExternalID:          "REF123456",
    PayerMessage:        "Refund for order #12345",
    PayeeNote:           "Refund processed",
    ReferenceIDToRefund: "original-transaction-reference-id",
    CallbackURL:         "https://example.com/callback",
})
if err != nil {
    log.Fatalf("Refund failed: %v", err)
}
fmt.Printf("Refund Reference ID: %s\n", response.ReferenceID)
fmt.Printf("Status: %s\n", response.Status)
```

#### Refund Transaction (V2)

```go
// Initiate a refund using V2 API
response, err := client.RefundV2(api.RefundRequest{
    Amount:              "100",
    Currency:            "EUR",
    ExternalID:          "REF123456",
    PayerMessage:        "Refund for order #12345",
    PayeeNote:           "Refund processed",
    ReferenceIDToRefund: "original-transaction-reference-id",
    CallbackURL:         "https://example.com/callback",
})
if err != nil {
    log.Fatalf("Refund V2 failed: %v", err)
}
fmt.Printf("Refund Reference ID: %s\n", response.ReferenceID)
```

#### Check Refund Status

```go
// Check the status of a refund
status, err := client.GetRefundStatus(response.ReferenceID)
if err != nil {
    log.Fatalf("Failed to get refund status: %v", err)
}
fmt.Printf("Refund Status: %s\n", status.Status)
fmt.Printf("Amount: %s %s\n", status.Amount, status.Currency)
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
client, err := momo.NewClient(apiUser, apiKey, subscriptionKey, momo.SANDBOX)

// Production environment (for live transactions)
client, err := momo.NewClient(apiUser, apiKey, subscriptionKey, momo.PRODUCTION)
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
