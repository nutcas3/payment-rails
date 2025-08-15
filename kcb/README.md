# KCB Buni API Integration for Go

This package provides a Go implementation for integrating with the KCB Buni API, allowing you to easily implement financial operations like account information retrieval, forex operations, and Vooma payments in your Go applications.

## Features

- **Authentication**: Bearer token authentication
- **Account Information**: Retrieve account details and balance
- **Forex Operations**: Get exchange rates and perform currency conversions
- **Vooma Payments**: Process payments through KCB's Vooma platform

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

### Get Account Information

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

### Make Vooma Payment

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
```

## API Documentation

For more information about the KCB Buni API, please refer to the [official KCB Buni Developer Portal](https://sandbox.buni.kcbgroup.com/devportal/apis).

## Requirements

- Go 1.13 or higher
- KCB Buni Developer Account
- API Token from the KCB Buni Developer Portal

## License

This package is released under the MIT License.
