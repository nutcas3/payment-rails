# Jenga V3 API SDK for Go

A Go client for the Equity (Finserve) Jenga V3 API. This SDK provides a simple interface to interact with the Jenga API for various financial operations.

## Features

- Authentication and Token Management
- Send Money (Within Equity, To Mobile Wallets, To Other Banks, RTGS, SWIFT)
- Receive Money
- Account Services (Balance Inquiry, Mini Statement, Full Statement)
- Bill Payments
- Airtime Purchase
- KYC and Identity Verification
- Forex Rates

## Pre-requisites

To use this SDK, you'll need:

1. Sign up with [Jenga API](https://developer.jengahq.io/) to get your API credentials
2. You'll need the following credentials:
   - API Key
   - Username
   - Password
   - Private Key (for generating signatures)

## Installation

Simply install with the go get command:

```bash
go get github.com/nutcas3/payment-rails/jenga
```

Then import it to your main package as:

```go
package main

import (
    "github.com/nutcas3/payment-rails/jenga"
)
```

## Usage

```go
// Initialize the Jenga client
client, err := jenga.NewClient(
    "your-api-key",          // API Key from Jenga Developer Portal
    "your-username",         // Username from Jenga Developer Portal
    "your-password",         // Password from Jenga Developer Portal
    "your-private-key",      // Private Key from Jenga Developer Portal
    jenga.SANDBOX,           // Environment (jenga.SANDBOX or jenga.PRODUCTION)
)
if err != nil {
    log.Fatalf("Failed to initialize Jenga client: %v", err)
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

See the [examples](../examples/jenga) directory for more usage examples.

### Send Money via RTGS

```go
// Create a request for RTGS transfer
req := api.SendMoneyRequest{
    Source: api.Source{
        CountryCode:   "KE", // Kenya
        Name:          "John Doe",
        AccountNumber: "1234567890",
    },
    Destination: api.Destination{
        Type:          "bank",
        CountryCode:   "KE", // Kenya
        Name:          "Jane Smith",
        AccountNumber: "0987654321",
        BankCode:      "01", // Bank code for the receiving bank
        BranchCode:    "112", // Branch code for the receiving bank
    },
    Transfer: api.Transfer{
        Type:         api.TransferTypeRTGS, // Specify RTGS transfer type
        Amount:       "50000.00", // RTGS is typically used for larger amounts
        CurrencyCode: "KES",
        Reference:    fmt.Sprintf("RTGS-REF-%d", time.Now().Unix()),
        Date:         time.Now().Format("2006-01-02"),
        Description:  "RTGS Transfer to Jane Smith",
    },
}

// Send the money using RTGS
response, err := client.SendMoney(req)
if err != nil {
    log.Fatalf("Error sending money via RTGS: %v", err)
}
```

### Send Money via SWIFT

```go
// Create a request for SWIFT international transfer
req := api.SendMoneyRequest{
    Source: api.Source{
        CountryCode:   "KE", // Kenya
        Name:          "John Doe",
        AccountNumber: "1234567890",
    },
    Destination: api.Destination{
        Type:          "bank",
        CountryCode:   "US", // United States
        Name:          "Jane Smith",
        AccountNumber: "0987654321",
        // SWIFT specific fields
        BankName:      "Bank of America",
        BankAddress:   "100 North Tryon Street, Charlotte, NC 28255, USA",
        SwiftCode:     "BOFAUS3N", // Example SWIFT code for Bank of America
        RoutingNumber: "026009593", // ABA routing number for US banks
    },
    Transfer: api.Transfer{
        Type:         api.TransferTypeSWIFT, // Specify SWIFT transfer type
        Amount:       "1000.00",
        CurrencyCode: "USD", // International transfer in USD
        Reference:    fmt.Sprintf("SWIFT-REF-%d", time.Now().Unix()),
        Date:         time.Now().Format("2006-01-02"),
        Description:  "International SWIFT Transfer to Jane Smith",
    },
}

// Send the money using SWIFT
response, err := client.SendMoney(req)
if err != nil {
    log.Fatalf("Error sending money via SWIFT: %v", err)
}
```

### Send Money to Mobile Wallet

```go
// Create a request for mobile wallet transfer (e.g., M-Pesa)
mobileWalletReq := api.MobileWalletRequest{
    Source: api.Source{
        CountryCode:   "KE", // Kenya
        Name:          "John Doe",
        AccountNumber: "0011547896523", // Source account number
    },
    Destination: struct {
        Type          string `json:"type"`
        CountryCode   string `json:"countryCode"`
        Name          string `json:"name"`
        MobileNumber  string `json:"mobileNumber"`
        WalletName    string `json:"walletName"`
    }{
        Type:         "mobile",
        CountryCode:  "KE", // Kenya
        Name:         "Jane Smith",
        MobileNumber: "254722000000", // Recipient's mobile number with country code
        WalletName:   "Mpesa", // Mobile wallet provider
    },
    Transfer: struct {
        Type         string `json:"type"`
        Amount       string `json:"amount"`
        CurrencyCode string `json:"currencyCode"`
        Reference    string `json:"reference"`
        Date         string `json:"date"`
        Description  string `json:"description"`
        CallbackUrl  string `json:"callbackUrl"`
    }{
        Type:         "MobileWallet",
        Amount:       "200.00", // Amount to send
        CurrencyCode: "KES", // Kenyan Shillings
        Reference:    jenga.GenerateReference(), // Unique transaction reference
        Date:         time.Now().Format("2006-01-02"),
        Description:  "Payment for services rendered",
        CallbackUrl:  "https://webhook.site/your-webhook-id", // Replace with your actual webhook URL
    },
}

// Send the money to mobile wallet
response, err := client.SendToMobileWallet(mobileWalletReq)
if err != nil {
    log.Fatalf("Error sending money to mobile wallet: %v", err)
}

// The response will include a transaction ID and initial status
// The final status will be sent to the callback URL provided
```

### Send Money Within Equity Bank (Internal Bank Transfer)

```go
// Create a request for internal bank transfer within Equity Bank
internalTransferReq := api.SendMoneyRequest{
    Source: api.Source{
        CountryCode:   "KE", // Kenya
        Name:          "John Doe",
        AccountNumber: "0011547896523", // Source account number
    },
    Destination: api.Destination{
        Type:          "bank",
        CountryCode:   "KE", // Kenya (same as source for internal transfer)
        Name:          "Jane Smith",
        AccountNumber: "0020154789652", // Destination account number
    },
    Transfer: api.Transfer{
        Type:         api.TransferTypeEFT, // EFT for internal transfers
        Amount:       "500.00", // Amount to transfer
        CurrencyCode: "KES", // Kenyan Shillings
        Reference:    jenga.GenerateReference(), // Unique transaction reference
        Date:         time.Now().Format("2006-01-02"),
        Description:  "Monthly rent payment",
    },
}

// Send the internal bank transfer request
response, err := client.SendInternalBankTransfer(internalTransferReq)
if err != nil {
    log.Fatalf("Error sending internal bank transfer: %v", err)
}

// The response will include a transaction ID and status
fmt.Printf("Transaction ID: %s, Status: %s\n", response.Data.TransactionID, response.Data.Status)
```

#### Cross-Border Internal Transfer

```go
// Example: Cross-border internal transfer (e.g., Kenya to Uganda)
crossBorderReq := api.SendMoneyRequest{
    Source: api.Source{
        CountryCode:   "KE", // Kenya
        Name:          "John Doe",
        AccountNumber: "0011547896523", // Source account number
    },
    Destination: api.Destination{
        Type:          "bank",
        CountryCode:   "UG", // Uganda (different from source for cross-border transfer)
        Name:          "Robert Smith",
        AccountNumber: "0020154789652", // Destination account number
    },
    Transfer: api.Transfer{
        Type:         api.TransferTypeEFT, // EFT for internal transfers
        Amount:       "1000.00", // Amount to transfer
        CurrencyCode: "USD", // USD for cross-border transfers
        Reference:    jenga.GenerateReference(), // Unique reference
        Date:         time.Now().Format("2006-01-02"),
        Description:  "Business payment",
    },
}

// Send the cross-border internal bank transfer request
response, err := client.SendInternalBankTransfer(crossBorderReq)
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
