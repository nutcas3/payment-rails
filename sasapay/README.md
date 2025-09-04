# SasaPay API Integration

This package provides a Go client for the SasaPay payment gateway API, enabling seamless integration with SasaPay's payment services.

## Features

- **Standard API Integration**
  - Customer to Business (C2B) payments
  - Business to Customer (B2C) payments
  - Business to Business (B2B) payments

- **Wallet as a Service (WAAS)**
  - Wallet creation
  - Balance inquiry
  - Wallet-to-wallet transfers
  - Transaction statements

- **Transaction Management**
  - Transaction status checking
  - Transaction verification

- **Webhook Integration**
  - Webhook event handling
  - Signature verification
  - Event-specific callbacks

## Installation

```bash
go get -u payment-rails/sasapay
```

## Quick Start

```go
package main

import (
    "os"
    "payment-rails/sasapay"
    "payment-rails/sasapay/pkg/api"
)

func main() {
    // Initialize client
    client, err := sasapay.NewClient(
        os.Getenv("SASAPAY_CLIENT_ID"),
        os.Getenv("SASAPAY_CLIENT_SECRET"),
        "sandbox", // Use "production" for live environment
    )
    if err != nil {
        panic(err)
    }

    // Optional: Set webhook secret for signature verification
    client.SetWebhookSecret(os.Getenv("SASAPAY_WEBHOOK_SECRET"))

    // Make API calls...
}
```

## Authentication

The client automatically handles OAuth2 authentication with SasaPay. It retrieves and caches access tokens, refreshing them as needed.

## Usage Examples

### Customer to Business (C2B) Payment

```go
import (
    "github.com/shopspring/decimal"
    "payment-rails/sasapay/pkg/api"
)

// Create C2B request
amount, _ := decimal.NewFromString("100.00")
c2bReq := api.C2BRequest{
    MerchantCode: "MERCHANT123",
    PhoneNumber:  "254712345678",
    Amount:       amount,
    Reference:    sasapay.GenerateReference(),
    Description:  "Payment for goods",
    CallbackURL:  "https://example.com/callbacks/c2b",
}

// Send C2B request
c2bResp, err := client.CustomerToBusiness(c2bReq)
if err != nil {
    // Handle error
}

// Use response
fmt.Printf("Transaction ID: %s\n", c2bResp.TransactionID)
```

### Business to Customer (B2C) Payment

```go
// Create B2C request
amount, _ := decimal.NewFromString("50.00")
b2cReq := api.B2CRequest{
    MerchantCode: "MERCHANT123",
    PhoneNumber:  "254712345678",
    Amount:       amount,
    Reference:    sasapay.GenerateReference(),
    Description:  "Refund for returned goods",
    CallbackURL:  "https://example.com/callbacks/b2c",
}

// Send B2C request
b2cResp, err := client.BusinessToCustomer(b2cReq)
```

### Wallet as a Service

```go
// Create a wallet
createWalletReq := api.CreateWalletRequest{
    PhoneNumber: "254712345678",
    FirstName:   "John",
    LastName:    "Doe",
    Email:       "john.doe@example.com",
    IDNumber:    "12345678",
    CallbackURL: "https://example.com/callbacks/wallet",
}

createWalletResp, err := client.CreateWallet(createWalletReq)

// Get wallet balance
balanceReq := api.WalletBalanceRequest{
    WalletID: createWalletResp.WalletID,
}

balanceResp, err := client.GetWalletBalance(balanceReq)

// Transfer between wallets
transferReq := api.WalletTransferRequest{
    SourceWalletID:      "WALLET123",
    DestinationWalletID: "WALLET456",
    Amount:              decimal.NewFromFloat(25.00),
    Reference:           sasapay.GenerateReference(),
    Description:         "Transfer funds between wallets",
    CallbackURL:         "https://example.com/callbacks/transfer",
}

transferResp, err := client.TransferToWallet(transferReq)
```

### Transaction Status and Verification

```go
// Check transaction status
statusReq := api.TransactionStatusRequest{
    TransactionID: "TXN123456789",
}

statusResp, err := client.CheckTransactionStatus(statusReq)

// Verify transaction
verifyReq := api.VerifyTransactionRequest{
    TransactionID: "TXN123456789",
}

verifyResp, err := client.VerifyTransaction(verifyReq)
```

### Webhook Handling

```go
// In your HTTP handler
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    // Define webhook handlers
    handlers := api.WebhookHandlers{
        PaymentReceived: func(event api.WebhookEvent) {
            // Handle payment received event
        },
        PaymentCompleted: func(event api.WebhookEvent) {
            // Handle payment completed event
        },
        PaymentFailed: func(event api.WebhookEvent) {
            // Handle payment failed event
        },
        WalletCreated: func(event api.WebhookEvent) {
            // Handle wallet created event
        },
        WalletTransferred: func(event api.WebhookEvent) {
            // Handle wallet transferred event
        },
    }

    // Process webhook
    err := client.ProcessWebhookRequest(r.Body, r.Header.Get("X-SasaPay-Signature"), handlers)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK)
}
```

## Configuration

### Environment Variables

- `SASAPAY_CLIENT_ID` - SasaPay API client ID
- `SASAPAY_CLIENT_SECRET` - SasaPay API client secret
- `SASAPAY_WEBHOOK_SECRET` - Secret for webhook signature verification

### Environment Selection

The client supports both sandbox and production environments:

```go
// Sandbox environment (default)
client, err := sasapay.NewClient(clientID, clientSecret, "sandbox")

// Production environment
client, err := sasapay.NewClient(clientID, clientSecret, "production")
```

## Error Handling

All client methods return errors that can be checked and handled:

```go
resp, err := client.CustomerToBusiness(req)
if err != nil {
    if apiErr, ok := err.(*api.Error); ok {
        // Handle API-specific error
        fmt.Printf("API Error: %s (Code: %s)\n", apiErr.Message, apiErr.Code)
    } else {
        // Handle other errors (network, parsing, etc.)
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Complete Example

See the [examples/sasapay/main.go](../examples/sasapay/main.go) file for a complete example of using the SasaPay client.

## License

This package is part of the payment-rails library.
