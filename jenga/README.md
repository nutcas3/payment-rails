# Jenga V3 API SDK for Go

A Go client for the Equity (Finserve) Jenga V3 API. This SDK provides a simple interface to interact with the Jenga API for various financial operations.

## Features

- Authentication and Token Management
- Send Money (Within Equity, To Mobile Wallets, To Other Banks, RTGS, SWIFT)
- Receive Money and Query Transactions
- Account Services (Balance Inquiry, Mini Statement, Full Statement, Account Validation)
- Bill Payments
- Airtime Purchase
- RegTech Services (KYC, AML, Customer Due Diligence)
- Forex Rates
- Webhook/Callback Handling for Notifications

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

## Account Services

### Account Validation

Validate Equity bank accounts across different countries.

```go
// Create an account validation request
validateReq := api.AccountValidateRequest{
    CountryCode:     "UG",            // Country code (e.g., UG for Uganda)
    AccountNumber:   "1036200681230", // Account number to validate
    AccountFullName: "DICKSON MAITEI", // Full name on the account
    ChargeAccount:   "1036200681230", // Optional: account to charge for this operation
}

// Validate the account
response, err := client.ValidateAccount(validateReq)
if err != nil {
    log.Fatalf("Error validating account: %v", err)
}

// Process the response
if response.Status {
    fmt.Println("Account validation successful!")
    fmt.Printf("Full Name: %s\n", response.Data.Account.FullNames)
    fmt.Printf("Account Number: %s\n", response.Data.Account.AccountNumber)
    fmt.Printf("Currency: %s\n", response.Data.Account.Currency)
    fmt.Printf("Account Status: %s\n", response.Data.Account.Status)
} else {
    fmt.Printf("Account validation failed. Code: %d, Message: %s\n", response.Code, response.Message)
}
```

### Account Balance

```go
// Get account balance
balance, err := client.GetAccountBalance(api.AccountBalanceRequest{
    CountryCode: "KE",
    AccountID:   "0011547896523",
})
```

### Mini Statement

```go
// Get mini statement
miniStatement, err := client.GetMiniStatement(api.MiniStatementRequest{
    CountryCode: "KE",
    AccountID:   "0011547896523",
})
```

### Full Statement

```go
// Get full statement
fullStatement, err := client.GetFullStatement(api.FullStatementRequest{
    CountryCode: "KE",
    AccountID:   "0011547896523",
    FromDate:    "2023-01-01",
    ToDate:      "2023-01-31",
})
```

## Receive Money

Receive money from various sources into an Equity Bank account.

```go
// Create a receive money request
receiveReq := api.ReceiveMoneyRequest{
    MerchantCode:    "1234567",                // Your merchant code
    MerchantAccount: "0011547896523",          // Account to receive funds
    CustomerName:    "John Doe",               // Name of the customer sending money
    Amount:          "500.00",                 // Amount to receive
    CurrencyCode:    "KES",                    // Currency code
    Reference:       jenga.GenerateReference(), // Unique reference
    Description:     "Payment for services",    // Description
}

// Process the receive money request
response, err := client.ReceiveMoney(receiveReq)
if err != nil {
    log.Fatalf("Error receiving money: %v", err)
}

// Check the response
if response.Status {
    fmt.Printf("Transaction ID: %s\n", response.Data.TransactionID)
    fmt.Printf("Status: %s\n", response.Data.Status)
} else {
    fmt.Printf("Failed to receive money: %s\n", response.Message)
}
```

### Query Receive Money Transaction

```go
// Create a query request
queryReq := api.ReceiveMoneyQueryRequest{
    MerchantCode:  "1234567",     // Your merchant code
    TransactionID: "TRX12345678", // Transaction ID from the receive money response
}

// Query the transaction status
response, err := client.QueryReceiveMoneyTransaction(queryReq)
if err != nil {
    log.Fatalf("Error querying transaction: %v", err)
}

// Process the response
fmt.Printf("Transaction Status: %s\n", response.Data.Status)
fmt.Printf("Transaction Date: %s\n", response.Data.TransactionDate)
```

## Airtime Purchase

Purchase airtime for mobile numbers across different telcos.

```go
// Create an airtime purchase request
airtimeReq := api.AirtimePurchaseRequest{
    CustomerMobile: "254722000000", // Mobile number with country code
    TelcoCode:      "SAF",         // Telco code (SAF for Safaricom, AIR for Airtel, etc.)
    Amount:         "100",         // Amount of airtime to purchase
    Reference:      jenga.GenerateReference(), // Unique reference
    CurrencyCode:   "KES",        // Currency code
}

// Purchase airtime
response, err := client.PurchaseAirtime(airtimeReq)
if err != nil {
    log.Fatalf("Error purchasing airtime: %v", err)
}

// Process the response
if response.Status {
    fmt.Println("Airtime purchase successful!")
} else {
    fmt.Printf("Airtime purchase failed: %s\n", response.Message)
}
```

## RegTech Services

### KYC Verification

```go
// Create a KYC verification request
kycReq := api.KYCRequest{
    DocumentType:   "ALIENID", // ID, PASSPORT, ALIENID, etc.
    DocumentNumber: "AB123456",
    CountryCode:    "KE",
    FirstName:      "John",     // Optional
    LastName:       "Doe",      // Optional
    DateOfBirth:    "1990-01-01", // Optional
}

// Verify identity
response, err := client.VerifyIdentity(kycReq)
if err != nil {
    log.Fatalf("Error verifying identity: %v", err)
}

// Process the response
if response.Status {
    fmt.Printf("Verified: %t\n", response.Data.Verified)
    fmt.Printf("Full Name: %s\n", response.Data.FullName)
} else {
    fmt.Printf("Verification failed: %s\n", response.Message)
}
```

### AML Screening

```go
// Create an AML screening request
amlReq := api.AMLScreeningRequest{
    FirstName:      "John",
    LastName:       "Doe",
    CountryCode:    "KE",
    DateOfBirth:    "1990-01-01", // Optional
    DocumentType:   "ALIENID",    // Optional
    DocumentNumber: "AB123456",   // Optional
    Reference:      jenga.GenerateReference(), // Optional
}

// Perform AML screening
response, err := client.PerformAMLScreening(amlReq)
if err != nil {
    log.Fatalf("Error performing AML screening: %v", err)
}

// Process the response
if response.Status {
    fmt.Printf("Risk Score: %d\n", response.Data.RiskScore)
    fmt.Printf("Risk Level: %s\n", response.Data.RiskLevel)
    fmt.Printf("Match Status: %t\n", response.Data.MatchStatus)
} else {
    fmt.Printf("AML screening failed: %s\n", response.Message)
}
```

### Customer Due Diligence (CDD)

```go
// Create a CDD request
cddReq := api.CDDRequest{
    CustomerID:     "CUS12345",
    CountryCode:    "KE",
    RiskLevel:      "medium",    // Optional: low, medium, high
    BusinessType:   "retail",    // Optional
    DocumentType:   "ALIENID",   // Optional
    DocumentNumber: "AB123456",  // Optional
    Reference:      jenga.GenerateReference(), // Optional
}

// Perform Customer Due Diligence
response, err := client.PerformCustomerDueDiligence(cddReq)
if err != nil {
    log.Fatalf("Error performing CDD: %v", err)
}

// Process the response
if response.Status {
    fmt.Printf("Risk Score: %d\n", response.Data.RiskScore)
    fmt.Printf("Risk Level: %s\n", response.Data.RiskLevel)
    fmt.Printf("Verification Status: %s\n", response.Data.VerificationStatus)
} else {
    fmt.Printf("CDD failed: %s\n", response.Message)
}
```

## Forex Rates

Get foreign exchange rates for different currencies.

```go
// Create a forex rates request
forexReq := api.ForexRatesRequest{
    CountryCode:  "KE",  // Country code
    CurrencyCode: "USD", // Currency code
    BaseCurrency: "KES", // Optional: Base currency for conversion
}

// Get forex rates
response, err := client.GetForexRates(forexReq)
if err != nil {
    log.Fatalf("Error getting forex rates: %v", err)
}

// Process the response
if response.Status {
    fmt.Printf("Buy Rate: %s\n", response.Data.BuyRate)
    fmt.Printf("Sell Rate: %s\n", response.Data.SellRate)
} else {
    fmt.Printf("Failed to get forex rates: %s\n", response.Message)
}
```

## Webhook Handling

Handle webhook notifications from Jenga API for real-time updates.

```go
// Initialize the client and set webhook secret
client, err := jenga.NewClient(apiKey, username, password, privateKeyPath, "sandbox")
if err != nil {
    log.Fatalf("Error initializing Jenga client: %v", err)
}

// Set webhook secret for signature validation
client.SetWebhookSecret(webhookSecret)

// Define webhook handlers
handlers := api.WebhookHandlers{
    TransactionSuccessHandler: func(event *api.WebhookEvent) {
        fmt.Println("Transaction successful:", event.ID)
        // Parse transaction data
        var data api.TransactionWebhookData
        json.Unmarshal(event.Data, &data)
        // Process the transaction data
        fmt.Printf("Transaction ID: %s\n", data.TransactionID)
        fmt.Printf("Amount: %s %s\n", data.Amount, data.Currency)
    },
    TransactionFailedHandler: func(event *api.WebhookEvent) {
        fmt.Println("Transaction failed:", event.ID)
        // Handle failed transaction
    },
    AccountUpdatedHandler: func(event *api.WebhookEvent) {
        fmt.Println("Account updated:", event.ID)
        // Handle account update
    },
    KYCUpdatedHandler: func(event *api.WebhookEvent) {
        fmt.Println("KYC updated:", event.ID)
        // Handle KYC update
    },
    DefaultHandler: func(event *api.WebhookEvent) {
        fmt.Printf("Received unhandled event type: %s\n", event.EventType)
    },
}

// Set up HTTP handler for webhooks
http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
    if err := client.HandleWebhook(w, r, handlers); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
})

// Start the server
http.ListenAndServe(":8080", nil)
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
