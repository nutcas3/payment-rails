# Mpesa SDK for Go (Daraja API)

A Golang client for the Safaricom Mpesa Daraja API. This SDK provides a simple interface to interact with the Mpesa API for various payment operations.

## Features

- STK Push (Lipa Na M-Pesa Online)
- STK Push Query
- Customer to Business (C2B) URL Registration
- Customer to Business (C2B) Simulation
- Business to Customer (B2C) Payment
- Business to Business (B2B) Payment
- Transaction Status Query
- Account Balance Query
- Payment Reversal

## Pre-requisites

To use this SDK, you'll need:

1. Sign up with [Safaricom Developer Portal](https://developer.safaricom.co.ke/) to get your API credentials
2. You'll need the following credentials:
   - API Key (Consumer Key)
   - Consumer Secret
   - Pass Key (for STK Push)

## Installation

Simply install with the go get command:

```bash
go get github.com/nutcas3/payment-rails/mpesa
```

Then import it to your main package as:

```go
package main

import (
    "github.com/nutcas3/payment-rails/mpesa"
)
```

## Usage

```go
// Initialize the Mpesa client
client, err := mpesa.NewClient(
    "your-api-key",          // Consumer Key from Safaricom Developer Portal
    "your-consumer-secret", // Consumer Secret from Safaricom Developer Portal
    "your-pass-key",        // Pass Key from Safaricom Developer Portal
    mpesa.SANDBOX,          // Environment (mpesa.SANDBOX or mpesa.PRODUCTION)
)
if err != nil {
    log.Fatalf("Failed to initialize Mpesa client: %v", err)
}

// Use custom HTTP client if needed
client.SetHttpClient(&http.Client{
    Timeout: 30 * time.Second,
})

// Get authentication token
token, err := client.GetAuthToken()
if err != nil {
    log.Fatalf("Failed to get auth token: %v", err)
}
fmt.Printf("Auth Token: %s\n", token)
```

## Examples

### STK Push (Lipa Na M-Pesa Online)

```go
// Initiate STK Push
stkResponse, err := client.InitiateStkPush(
    "174379",                                     // Business Short Code
    "CustomerPayBillOnline",                      // Transaction Type
    "1",                                          // Amount
    "254708374149",                               // Party A (Phone number)
    "174379",                                     // Party B (Short code)
    "254708374149",                               // Phone Number
    "https://example.com/callback",               // Callback URL
    "Test Payment",                               // Account Reference
    "Test Payment",                               // Transaction Description
)
if err != nil {
    log.Fatalf("Failed to initiate STK push: %v", err)
}
fmt.Printf("STK Push Response: %+v\n", stkResponse)

// Query STK Push status
queryResponse, err := client.QueryStkPush(
    "174379",                         // Business Short Code
    "ws_CO_DMZ_12345678901234567",   // Checkout Request ID from previous STK Push
)
if err != nil {
    log.Fatalf("Failed to query STK push: %v", err)
}
fmt.Printf("STK Push Query Response: %+v\n", queryResponse)
```

### Customer to Business (C2B)

```go
// Register C2B URL
c2bRegisterResponse, err := client.C2BRegisterURL(
    "600000",                                     // Short Code
    "Completed",                                  // Response Type
    "https://example.com/c2b/confirmation",       // Confirmation URL
    "https://example.com/c2b/validation",         // Validation URL
)
if err != nil {
    log.Fatalf("Failed to register C2B URL: %v", err)
}
fmt.Printf("C2B Register URL Response: %+v\n", c2bRegisterResponse)

// C2B Simulate (only works in sandbox)
c2bSimulateResponse, err := client.C2BSimulate(
    600000,                // Short Code
    "CustomerPayBillOnline", // Command ID
    1,                     // Amount
    254708374149,          // MSISDN (Phone number)
    "Test",                // Bill Reference Number
)
if err != nil {
    log.Fatalf("Failed to simulate C2B: %v", err)
}
fmt.Printf("C2B Simulate Response: %+v\n", c2bSimulateResponse)
```

### Business to Customer (B2C) Payment

```go
b2cResponse, err := client.B2CPayment(
    "TestInitiator",                              // Initiator Name
    "SecurityCredential",                         // Security Credential
    "BusinessPayment",                            // Command ID
    1,                                            // Amount
    600000,                                       // Party A (Short code)
    254708374149,                                 // Party B (Phone number)
    "Test B2C Payment",                           // Remarks
    "https://example.com/b2c/timeout",            // Queue Timeout URL
    "https://example.com/b2c/result",             // Result URL
    "Test",                                       // Occasion
)
if err != nil {
    log.Fatalf("Failed to make B2C payment: %v", err)
}
fmt.Printf("B2C Payment Response: %+v\n", b2cResponse)
```

### Business to Business (B2B) Payment

```go
b2bResponse, err := client.B2BPayment(
    "TestInitiator",                              // Initiator
    "SecurityCredential",                         // Security Credential
    "BusinessPayBill",                            // Command ID
    "4",                                          // Sender Identifier Type
    "4",                                          // Receiver Identifier Type
    "100",                                        // Amount
    "600000",                                     // Party A (Short code)
    "600001",                                     // Party B (Short code)
    "Test",                                       // Account Reference
    "254708374149",                               // Requester
    "Test B2B Payment",                           // Remarks
    "https://example.com/b2b/timeout",            // Queue Timeout URL
    "https://example.com/b2b/result",             // Result URL
)
if err != nil {
    log.Fatalf("Failed to make B2B payment: %v", err)
}
fmt.Printf("B2B Payment Response: %+v\n", b2bResponse)
```

### Transaction Status

```go
statusResponse, err := client.TransactionStatus(
    "TestInitiator",                              // Initiator
    "SecurityCredential",                         // Security Credential
    "TransactionStatusQuery",                     // Command ID
    "LKXXXX1234",                                 // Transaction ID
    600000,                                       // Party A (Short code)
    4,                                            // Identifier Type
    "https://example.com/status/result",          // Result URL
    "https://example.com/status/timeout",         // Queue Timeout URL
    "Test Transaction Status",                    // Remarks
    "Test",                                       // Occasion
)
if err != nil {
    log.Fatalf("Failed to query transaction status: %v", err)
}
fmt.Printf("Transaction Status Response: %+v\n", statusResponse)
```

### Account Balance

```go
balanceResponse, err := client.AccountBalance(
    "TestInitiator",                              // Initiator
    "SecurityCredential",                         // Security Credential
    "AccountBalance",                             // Command ID
    600000,                                       // Party A (Short code)
    4,                                            // Identifier Type
    "Test Account Balance",                       // Remarks
    "https://example.com/balance/timeout",        // Queue Timeout URL
    "https://example.com/balance/result",         // Result URL
)
if err != nil {
    log.Fatalf("Failed to query account balance: %v", err)
}
fmt.Printf("Account Balance Response: %+v\n", balanceResponse)
```

### Payment Reversal

```go
reversalResponse, err := client.Reversal(
    "TestInitiator",                              // Initiator
    "SecurityCredential",                         // Security Credential
    "TransactionReversal",                        // Command ID
    "LKXXXX1234",                                 // Transaction ID
    1,                                            // Amount
    600000,                                       // Receiver Party
    4,                                            // Receiver Identifier Type
    "https://example.com/reversal/result",        // Result URL
    "https://example.com/reversal/timeout",       // Queue Timeout URL
    "Test Reversal",                              // Remarks
    "Test",                                       // Occasion
)
if err != nil {
    log.Fatalf("Failed to reverse transaction: %v", err)
}
fmt.Printf("Reversal Response: %+v\n", reversalResponse)
```

## Testing

To run the tests:

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
