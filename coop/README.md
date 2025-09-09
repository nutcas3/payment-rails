# Co-op Bank Kenya SDK

A comprehensive Go SDK for integrating with Co-operative Bank of Kenya APIs.

## Features

- **Account Balance**: Check account balance for any Co-op Bank account
- **Account Transactions**: Retrieve transaction history with configurable limits
- **Exchange Rates**: Get current SPOT exchange rates between currencies
- **Internal Funds Transfer**: Transfer funds between Co-op Bank accounts
- **PesaLink**: Transfer funds to other banks via IPSL network
- **Transaction Status**: Check the status of previously requested transactions

## Installation

```bash
go get payment-rails/coop
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "payment-rails/coop"
    "payment-rails/coop/pkg/api"
)

func main() {
    // Initialize client
    client, err := coop.NewClient("your-client-id", "your-client-secret", api.SANDBOX)
    if err != nil {
        log.Fatal(err)
    }

    // Check account balance
    balance, err := client.AccountBalance("1234567890")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Account Balance: %f %s\n", balance.ClearedBalance, balance.Currency)
}
```

## Configuration

### Environment Variables

You can set your credentials using environment variables:

```bash
export COOP_CLIENT_ID="your-client-id"
export COOP_CLIENT_SECRET="your-client-secret"
```

### Environments

- `api.SANDBOX` - For testing and development
- `api.PRODUCTION` - For live transactions

## API Methods

### Account Balance

```go
balance, err := client.AccountBalance("1234567890")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Cleared Balance: %f\n", balance.ClearedBalance)
```

### Account Transactions

```go
// Get last 5 transactions
transactions, err := client.AccountTransactions("1234567890", "5")
if err != nil {
    log.Fatal(err)
}

for _, txn := range transactions.Transactions {
    fmt.Printf("Date: %s, Amount: %f, Narration: %s\n", 
        txn.TransactionDate, txn.Amount, txn.Narration)
}
```

### Exchange Rate

```go
rate, err := client.ExchangeRate("KES", "USD")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Exchange Rate: 1 %s = %f %s\n", 
    rate.FromCurrencyCode, rate.Rate, rate.ToCurrencyCode)
```

### Internal Funds Transfer

```go
destinations := []api.Destination{
    {
        ReferenceNumber:     "REF001",
        AccountNumber:       "0987654321",
        Amount:              1000.00,
        TransactionCurrency: "KES",
        Narration:           "Payment for services",
    },
}

response, err := client.InternalFundsTransfer(
    "1234567890", // source account
    1000.00,      // total amount
    "KES",        // currency
    "Bulk transfer", // narration
    destinations,
)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Transaction ID: %s\n", response.TransactionID)
```

### PesaLink Transfer

```go
destinations := []api.PesaLinkDestination{
    {
        ReferenceNumber:     "REF001",
        DestinationBank:     "01", // Bank code
        AccountNumber:       "1122334455",
        Amount:              500.00,
        TransactionCurrency: "KES",
        Narration:           "Payment",
    },
}

response, err := client.PesaLinkTransfer(
    "1234567890", // source account
    500.00,       // total amount
    "KES",        // currency
    "PesaLink payment", // narration
    destinations,
)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Transaction ID: %s\n", response.TransactionID)
```

### Transaction Status

```go
status, err := client.TransactionStatus("COOP-1234567890123")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Status: %s, Amount: %f\n", 
    status.TransactionStatus, status.TransactionAmount)
```

## Error Handling

The SDK provides detailed error messages for different scenarios:

```go
balance, err := client.AccountBalance("invalid-account")
if err != nil {
    // Handle specific error cases
    fmt.Printf("Error: %s\n", err.Error())
}
```

## Authentication

The SDK handles OAuth2 authentication automatically. Access tokens are cached and refreshed as needed.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions, please open an issue in the GitHub repository.
