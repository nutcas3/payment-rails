# Standard Bank API SDK

This package provides a Go client for integrating with Standard Bank's API Marketplace, enabling seamless payment processing, internal transfers, and provider-based transactions.

## Features

- Payment Processing
  - Create and retrieve payments with full status tracking
- Internal Transfers
  - Execute transfers between accounts
- Provider Payments
  - Query and execute payments through various payment providers

## Installation

```bash
go get payment-rails/standardbank
```

## Quick Start

### Basic Setup

```go
package main

import (
    "context"
    "fmt"
    "log"

    "payment-rails/standardbank"
    "payment-rails/standardbank/pkg/api"
)

func main() {
    // Initialize the client
    client := standardbank.NewClient(
        "your-client-id",
        "your-client-secret",
        "your-api-key",
        standardbank.WithEnvironment(standardbank.EnvironmentSandbox),
    )

    ctx := context.Background()

    // Create a payment
    payment, err := client.Payments().Create(ctx, api.PaymentRequest{
        Amount:    100.50,
        Currency:   "ZAR",
        Reference: "PAY-001",
        Description: "Payment for services",
    })

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Payment created: %s\n", payment.PaymentID)
}
```

## Usage Examples

### Payments

#### Create a Payment

```go
ctx := context.Background()

payment, err := client.Payments().Create(ctx, api.PaymentRequest{
    Amount:            1000.00,
    Currency:          "ZAR",
    Reference:         "INV-2024-001",
    Description:       "Invoice payment",
    SourceAccount:     "ACC123456",
    DestinationAccount: "ACC789012",
    BeneficiaryName:   "John Doe",
    IdempotencyKey:    "unique-key-123",
    Metadata: map[string]interface{}{
        "invoiceId": "INV-001",
        "customerId": "CUST-123",
    },
})

if err != nil {
    log.Fatal(err)
}

fmt.Printf("Payment ID: %s\n", payment.PaymentID)
fmt.Printf("Status: %s\n", payment.Status)
```

#### Get Payment Status

```go
// By payment ID
status, err := client.Payments().GetStatus(ctx, "pay_123456")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Payment Status: %s\n", status.Status)
fmt.Printf("Last Updated: %s\n", status.LastUpdated)

// By reference
payment, err := client.Payments().GetByReference(ctx, "INV-2024-001")
if err != nil {
    log.Fatal(err)
}
```

#### Retrieve a Payment

```go
payment, err := client.Payments().Get(ctx, "pay_123456")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Payment Amount: %.2f %s\n", payment.Amount, payment.Currency)
fmt.Printf("Processing Date: %s\n", payment.ProcessingDate)
```

### Internal Transfers

#### Create an Internal Transfer

```go
transfer, err := client.Transfers().Create(ctx, api.InternalTransferRequest{
    SourceAccount:      "ACC001",
    DestinationAccount: "ACC002",
    Amount:            500.00,
    Currency:          "ZAR",
    Reference:         "TRANSFER-001",
    Description:       "Internal account transfer",
    DestinationName:   "Jane Smith",
    IdempotencyKey:    "transfer-key-123",
})

if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transfer ID: %s\n", transfer.TransferID)
fmt.Printf("Status: %s\n", transfer.Status)
```

#### Get Transfer Status

```go
status, err := client.Transfers().GetStatus(ctx, "transfer_123456")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transfer Status: %s\n", status.Status)
if status.FailureReason != "" {
    fmt.Printf("Failure Reason: %s\n", status.FailureReason)
}
```

### Provider Payments

#### List Available Providers

```go
providers, err := client.Providers().List(ctx)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Total Providers: %d\n", providers.Total)
for _, provider := range providers.Providers {
    fmt.Printf("- %s (%s): %s\n", provider.Name, provider.ID, provider.Type)
}
```

#### Get Provider Details

```go
provider, err := client.Providers().Get(ctx, "provider_1")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Provider: %s\n", provider.Name)
fmt.Printf("Type: %s\n", provider.Type)
fmt.Printf("Status: %s\n", provider.Status)
```

#### Execute Provider Payment

```go
payment, err := client.Providers().Pay(ctx, api.ProviderPaymentRequest{
    ProviderID:  "provider_1",
    Amount:      200.00,
    Currency:    "ZAR",
    Reference:   "PROV-PAY-001",
    Description: "Mobile wallet payment",
    Destination: map[string]interface{}{
        "accountNumber": "1234567890",
        "name":          "John Doe",
        "phoneNumber":   "+27123456789",
    },
    IdempotencyKey: "provider-key-123",
})

if err != nil {
    log.Fatal(err)
}

fmt.Printf("Provider Payment ID: %s\n", payment.PaymentID)
fmt.Printf("Provider Reference: %s\n", payment.ProviderReference)
```

## Webhooks

### Setting Up Webhook Handler

```go
// Set webhook secret for signature verification
client.SetWebhookSecret("your-webhook-secret")

// Register handlers for specific event types
client.RegisterWebhookHandler(api.EventPaymentCompleted, func(event api.WebhookEvent) error {
    paymentID := event.Data["paymentId"].(string)
    status := event.Data["status"].(string)

    fmt.Printf("Payment %s completed with status: %s\n", paymentID, status)

    // Update your database, send notifications, etc.
    return nil
})

client.RegisterWebhookHandler(api.EventPaymentFailed, func(event api.WebhookEvent) error {
    paymentID := event.Data["paymentId"].(string)
    reason := event.Data["failureReason"].(string)

    fmt.Printf("Payment %s failed: %s\n", paymentID, reason)

    // Handle payment failure
    return nil
})
```

### Webhook Event Types

- `api.EventPaymentCompleted` - Payment successfully completed
- `api.EventPaymentFailed` - Payment failed
- `api.EventPaymentPending` - Payment is pending
- `api.EventPaymentCancelled` - Payment was cancelled
- `api.EventTransferCompleted` - Transfer successfully completed
- `api.EventTransferFailed` - Transfer failed
- `api.EventProviderPaymentCompleted` - Provider payment completed
- `api.EventProviderPaymentFailed` - Provider payment failed

## Environment Variables

```bash
export STANDARD_BANK_CLIENT_ID="your-client-id"
export STANDARD_BANK_CLIENT_SECRET="your-client-secret"
export STANDARD_BANK_API_KEY="your-api-key"
export STANDARD_BANK_ENVIRONMENT="sandbox" # or "production"
export STANDARD_BANK_WEBHOOK_SECRET="your-webhook-secret"
```

## API Reference

### Client Methods

- `Payments()` - Returns payments service
- `Transfers()` - Returns transfers service
- `Providers()` - Returns providers service
- `SetWebhookSecret(secret string)` - Set webhook secret
- `HandleWebhook(w, r)` - Handle incoming webhook
- `RegisterWebhookHandler(eventType, handler)` - Register webhook handler

### Payments Service

- `Create(ctx, req)` - Create a payment
- `Get(ctx, paymentID)` - Get payment by ID
- `GetStatus(ctx, paymentID)` - Get payment status
- `GetByReference(ctx, reference)` - Get payment by reference

### Transfers Service

- `Create(ctx, req)` - Create internal transfer
- `Get(ctx, transferID)` - Get transfer by ID
- `GetStatus(ctx, transferID)` - Get transfer status

### Providers Service

- `List(ctx)` - List all providers
- `Get(ctx, providerID)` - Get provider by ID
- `Pay(ctx, req)` - Execute provider payment
