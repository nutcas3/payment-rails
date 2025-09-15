# NCBA Bank Integration

This package provides a Go client for integrating with NCBA Bank's API services. It supports account operations, funds transfers, and transaction status checking.

## Features

- Account Operations
  - Get account details and balance
  - Retrieve mini statements
  - Get detailed account statements
- Funds Transfer
  - Internal transfers (NCBA to NCBA)
  - External transfers (to other banks)
  - RTGS transfers
  - PesaLink transfers
- Transaction Status Checking

## Installation

```bash
go get github.com/nutcas3/payment-rails/ncba
```

## Usage

### Initialize Client

```go
import "github.com/nutcas3/payment-rails/ncba"

client := ncba.NewClient(
    "your-api-key",
    "your-username",
    "your-password",
)
```

### Account Operations

```go
// Get account details
details, err := client.GetAccountDetails("KE", "1234567890")

// Get mini statement
statement, err := client.GetMiniStatement("KE", "1234567890")

// Get account statement
statement, err := client.GetAccountStatement("KE", "1234567890", "01012024", "15012024")
```

### Funds Transfer

```go
// Internal Transfer
transfer := ncba.InternalTransferRequest{
    TransferRequest: ncba.TransferRequest{
        SourceAccount:      "1234567890",
        DestinationAccount: "0987654321",
        Amount:            1000.00,
        Currency:          "KES",
        Reference:         "INV001",
        Narration:         "Payment for services",
    },
    DestinationName: "John Doe",
}

resp, err := client.SendInternalTransfer(transfer)

// PesaLink Transfer
pesalink := ncba.PesaLinkTransferRequest{
    TransferRequest: ncba.TransferRequest{
        SourceAccount:      "1234567890",
        DestinationAccount: "9876543210",
        Amount:            5000.00,
        Currency:          "KES",
        Reference:         "INV002",
        Narration:         "PesaLink transfer",
    },
    DestinationBank: "KCB",
    PhoneNumber:     "+254712345678",
}

resp, err := client.SendPesaLinkTransfer(pesalink)
```

### Check Transaction Status

```go
status, err := client.CheckTransactionStatus("transaction-id")
```

## Error Handling

All methods return errors that should be checked before using the response. The client handles authentication automatically, refreshing the token when needed.

## Configuration

The client accepts the following configuration parameters:
- API Key
- Username
- Password

These can be provided during client initialization.
