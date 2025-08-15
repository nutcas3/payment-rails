# Airtel Money Integration for Go

This package provides a Go implementation for integrating with the Airtel Money API, allowing you to easily implement payment collection and disbursement functionalities in your Go applications.

## Features

- **Authentication**: OAuth2 token management with automatic refresh
- **Collection API**: USSD Push payments and status checking
- **Disbursement API**: Send money to mobile wallets
- **Refunds**: Process refunds for transactions
- **Account Balance**: Check your Airtel Money account balance
- **Multi-country support**: Works with all countries where Airtel Money is available

## Installation

```bash
go get github.com/nutcase/payment-rails/airtel
```

## Usage

### Initialize the Client

```go
import (
    "github.com/nutcase/payment-rails/airtel"
)

// Create a new client
client, err := airtel.New(
    "YOUR_CLIENT_ID",
    "YOUR_CLIENT_SECRET",
    "YOUR_PUBLIC_KEY",
    true, // true for sandbox, false for production
    "KE", // Country code (e.g., KE for Kenya)
    "KES" // Currency code (e.g., KES for Kenyan Shilling)
)
if err != nil {
    log.Fatalf("Failed to initialize Airtel Money client: %v", err)
}
```

### USSD Push (Collection)

```go
// Initiate a USSD Push payment
response, err := client.UssdPush(
    "YOUR_REFERENCE",
    "700000000", // Phone number without country code
    10.0,        // Amount
    "TX123"      // Transaction ID
)
if err != nil {
    log.Printf("Failed to initiate USSD Push: %v", err)
} else {
    fmt.Printf("USSD Push initiated: Success=%v, Message=%s\n", 
        response.Status.Success, 
        response.Status.Message)
}
```

### Check Transaction Status

```go
// Check the status of a transaction
status, err := client.GetTransactionStatus("TX123")
if err != nil {
    log.Printf("Failed to check transaction status: %v", err)
} else {
    fmt.Printf("Transaction Status: %s\n", status.Data.Transaction.Status)
}
```

### Disbursement (Send Money)

```go
// Send money to a customer
response, err := client.Disburse(
    "YOUR_REFERENCE",
    "700000000", // Phone number without country code
    10.0,        // Amount
    "TX123",     // Transaction ID
    "1234"       // PIN
)
if err != nil {
    log.Printf("Failed to initiate disbursement: %v", err)
} else {
    fmt.Printf("Disbursement initiated: Success=%v, Message=%s\n", 
        response.Status.Success, 
        response.Status.Message)
}
```

### Check Disbursement Status

```go
// Check the status of a disbursement
status, err := client.GetDisbursementStatus("TX123")
if err != nil {
    log.Printf("Failed to check disbursement status: %v", err)
} else {
    fmt.Printf("Disbursement Status: %s\n", status.Data.Transaction.Status)
}
```

### Process Refund

```go
// Refund a transaction
response, err := client.RefundTransaction("AM123456789", 10.0)
if err != nil {
    log.Printf("Failed to initiate refund: %v", err)
} else {
    fmt.Printf("Refund initiated: Success=%v, Message=%s\n", 
        response.Status.Success, 
        response.Status.Message)
}
```

### Check Account Balance

```go
// Get account balance
balance, err := client.GetAccountBalance()
if err != nil {
    log.Printf("Failed to get account balance: %v", err)
} else {
    fmt.Printf("Account Balance: %.2f %s\n", balance.Data.Balance, balance.Data.Currency)
}
```

## API Documentation

For more information about the Airtel Money API, please refer to the [official Airtel Developer Portal](https://developers.airtel.africa/).

## Requirements

- Go 1.13 or higher
- Airtel Money Developer Account
- Client ID and Client Secret from the Airtel Developer Portal

## License

This package is released under the MIT License.
