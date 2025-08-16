# KCB Buni API Integration for Go

This package provides a Go implementation for integrating with the KCB Buni API, allowing you to easily implement financial operations like account information retrieval, forex operations, payments, and utility bill payments in your Go applications.

## Features

- **Authentication**: Bearer token authentication
- **Account Operations**: 
  - Retrieve account details and balance
  - Get account statements
  - Transfer funds between accounts
- **Forex Operations**: Get exchange rates and perform currency conversions
- **Payment Services**:
  - Vooma payments
  - PesaLink transfers
  - Mobile money transfers
  - Utility bill payments

## Installation

```bash
go get github.com/nutcas3/payment-rails/kcb
```

## Usage

### Initialize the Client

```go
import (
    "github.com/nutcas3/payment-rails/kcb"
)

// Create a new client
client, err := kcb.New(
    "YOUR_API_TOKEN", // API token from KCB Buni Developer Portal
    true,             // true for sandbox, false for production
)
if err != nil {
    log.Fatalf("Failed to initialize KCB Buni client: %v", err)
}
```

### Account Operations

#### Get Account Information

```go
// Retrieve account information
accountInfo, err := client.GetAccountInfo()
if err != nil {
    log.Printf("Failed to get account information: %v", err)
} else {
    fmt.Printf("Account Number: %s\n", accountInfo.Data.AccountNumber)
    fmt.Printf("Account Name: %s\n", accountInfo.Data.AccountName)
    fmt.Printf("Balance: %.2f %s\n", accountInfo.Data.Balance, accountInfo.Data.Currency)
}
```

#### Get Account Balance

```go
// Get balance for a specific account
balance, err := client.GetAccountBalance("1234567890")
if err != nil {
    log.Printf("Failed to get account balance: %v", err)
} else {
    fmt.Printf("Account: %s\n", balance.Data.AccountNumber)
    fmt.Printf("Balance: %.2f %s\n", balance.Data.Balance, balance.Data.Currency)
    fmt.Printf("As of: %s\n", balance.Data.AsOf)
}
```

#### Get Account Statement

```go
// Get account statement for a date range
statement, err := client.GetAccountStatement("1234567890", "2025-01-01", "2025-01-31")
if err != nil {
    log.Printf("Failed to get account statement: %v", err)
} else {
    fmt.Printf("Account: %s\n", statement.Data.AccountNumber)
    fmt.Printf("Period: %s to %s\n", statement.Data.StartDate, statement.Data.EndDate)
    fmt.Printf("Transactions: %d\n", len(statement.Data.Transactions))
    
    for _, tx := range statement.Data.Transactions {
        fmt.Printf("  %s: %s %.2f - %s\n", 
            tx.TransactionDate.Format("2006-01-02"),
            tx.Type,
            tx.Amount,
            tx.Description)
    }
}
```

#### Transfer Funds

```go
// Transfer funds between accounts
transfer, err := client.TransferFunds(
    "1234567890",     // Source account
    "0987654321",     // Destination account
    1000.0,           // Amount
    "KES",            // Currency
    "INV123456",      // Reference
    "Invoice payment" // Narration
)
if err != nil {
    log.Printf("Failed to transfer funds: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", transfer.Data.TransactionID)
    fmt.Printf("Amount: %.2f %s\n", transfer.Data.Amount, transfer.Data.Currency)
    fmt.Printf("Status: %s\n", transfer.Data.Status)
}
```

### Get Forex Rates

```go
// Get exchange rates for a specific currency
forexRates, err := client.GetForexRates("USD")
if err != nil {
    log.Printf("Failed to get forex rates: %v", err)
} else {
    fmt.Printf("Base Currency: %s\n", forexRates.Data.BaseCurrency)
    for currency, rate := range forexRates.Data.Rates {
        fmt.Printf("%s: %.4f\n", currency, rate)
    }
}
```

### Exchange Currency

```go
// Convert an amount from one currency to another
exchange, err := client.ExchangeCurrency("EUR", "USD", 100.0)
if err != nil {
    log.Printf("Failed to exchange currency: %v", err)
} else {
    fmt.Printf("From: %s %.2f\n", exchange.Data.FromCurrency, exchange.Data.Amount)
    fmt.Printf("To: %s %.2f\n", exchange.Data.ToCurrency, exchange.Data.ConvertedAmount)
    fmt.Printf("Exchange Rate: %.4f\n", exchange.Data.ExchangeRate)
}
```

### Payment Services

#### Make Vooma Payment

```go
// Process a payment using Vooma
payment, err := client.VoomaPay(100.0)
if err != nil {
    log.Printf("Failed to make Vooma payment: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", payment.Data.TransactionID)
    fmt.Printf("Amount: %.2f %s\n", payment.Data.Amount, payment.Data.Currency)
    fmt.Printf("Status: %s\n", payment.Data.Status)
}

// Check status of a Vooma payment
status, err := client.CheckVoomaStatus("TX123456789")
if err != nil {
    log.Printf("Failed to check payment status: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", status.Data.TransactionID)
    fmt.Printf("Status: %s\n", status.Data.Status)
}
```

#### PesaLink Transfer

```go
// Transfer funds to another bank via PesaLink
transfer, err := client.PesalinkTransfer(
    "1234567890",     // Source account
    "0987654321",     // Destination account
    "01",             // Destination bank code
    1000.0,           // Amount
    "KES",            // Currency
    "INV123456",      // Reference
    "Invoice payment", // Narration
    "254712345678"    // Phone number
)
if err != nil {
    log.Printf("Failed to make PesaLink transfer: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", transfer.Data.TransactionID)
    fmt.Printf("Amount: %.2f %s\n", transfer.Data.Amount, transfer.Data.Currency)
    fmt.Printf("Status: %s\n", transfer.Data.Status)
}

// Check status of a PesaLink transfer
status, err := client.CheckPesalinkStatus("TX123456789")
if err != nil {
    log.Printf("Failed to check transfer status: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", status.Data.TransactionID)
    fmt.Printf("Status: %s\n", status.Data.Status)
}
```

#### Mobile Money Transfer

```go
// Transfer funds to a mobile money account
transfer, err := client.MobileMoneyTransfer(
    "1234567890",     // Source account
    "254712345678",   // Phone number
    1000.0,           // Amount
    "KES",            // Currency
    "INV123456",      // Reference
    "Invoice payment", // Narration
    "MPESA"           // Provider
)
if err != nil {
    log.Printf("Failed to make mobile money transfer: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", transfer.Data.TransactionID)
    fmt.Printf("Amount: %.2f %s\n", transfer.Data.Amount, transfer.Data.Currency)
    fmt.Printf("Status: %s\n", transfer.Data.Status)
}

// Check status of a mobile money transfer
status, err := client.CheckMobileMoneyStatus("TX123456789")
if err != nil {
    log.Printf("Failed to check transfer status: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", status.Data.TransactionID)
    fmt.Printf("Status: %s\n", status.Data.Status)
}
```

#### Utility Payments

```go
// Get list of utility providers
providers, err := client.GetUtilityProviders()
if err != nil {
    log.Printf("Failed to get utility providers: %v", err)
} else {
    fmt.Printf("Available providers: %d\n", len(providers.Data.Providers))
    for _, provider := range providers.Data.Providers {
        fmt.Printf("  %s - %s (%s)\n", 
            provider.ProviderID,
            provider.ProviderName,
            provider.Category)
    }
}

// Pay a utility bill
payment, err := client.PayUtility(
    "1234567890",     // Source account
    "KPLC",           // Provider ID
    "12345678",       // Account number with provider
    1000.0,           // Amount
    "KES",            // Currency
    "BILL123456",     // Reference
    "254712345678"    // Phone number for notifications
)
if err != nil {
    log.Printf("Failed to pay utility bill: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", payment.Data.TransactionID)
    fmt.Printf("Amount: %.2f %s\n", payment.Data.Amount, payment.Data.Currency)
    fmt.Printf("Status: %s\n", payment.Data.Status)
    fmt.Printf("Receipt Number: %s\n", payment.Data.ReceiptNumber)
}

// Check status of a utility payment
status, err := client.CheckUtilityPaymentStatus("TX123456789")
if err != nil {
    log.Printf("Failed to check payment status: %v", err)
} else {
    fmt.Printf("Transaction ID: %s\n", status.Data.TransactionID)
    fmt.Printf("Status: %s\n", status.Data.Status)
}
```

## API Documentation

For more information about the KCB Buni API, please refer to the [official KCB Buni Developer Portal](https://sandbox.buni.kcbgroup.com/devportal/apis).

## Requirements

- Go 1.13 or higher
- KCB Buni Developer Account
- API Token from the KCB Buni Developer Portal

## License

This package is released under the MIT License.
