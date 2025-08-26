# Absa Retail Payments API Integration

A Go client library for integrating with Absa Bank's Retail Payments API, providing a comprehensive set of banking and payment services.

## Features

- **Account Services**
  - Account Balance Inquiry
  - Mini and Full Statement Retrieval
  - Account Validation

- **Payment Services**
  - Bank Transfers (Inter-bank and Intra-bank)
  - Mobile Wallet Transfers
  - Bill Payments
  - Airtime Purchases
  - Transaction Status Queries

- **Advanced Payment Features**
  - Bulk Payments
  - Standing Orders/Recurring Payments
  - Foreign Exchange Transfers

- **Beneficiary Management**
  - Create, List, Update, and Delete Beneficiaries
  - Support for Bank, Mobile, and Biller Beneficiary Types

- **Authentication Methods**
  - OTP Request and Verification
  - Transaction Authentication
  - Device Registration

## Installation

```bash
go get -u payment-rails/absa
```

## Usage

### Initialize the Client

```go
import (
    "os"
    "payment-rails/absa"
)

func main() {
    // Initialize with environment variables
    clientID := os.Getenv("ABSA_CLIENT_ID")
    clientSecret := os.Getenv("ABSA_CLIENT_SECRET")
    apiKey := os.Getenv("ABSA_API_KEY")
    environment := "sandbox" // Use "production" for live environment

    client, err := absa.NewClient(clientID, clientSecret, apiKey, environment)
    if err != nil {
        // Handle error
    }
    
    // Now you can use the client to make API calls
}
```

### Basic Operations

#### Check Account Balance

```go
balanceReq := api.AccountBalanceRequest{
    AccountNumber: "1234567890",
}

balance, err := client.GetAccountBalance(balanceReq)
if err != nil {
    // Handle error
}

fmt.Printf("Available Balance: %s %s\n", api.FormatAmount(balance.AvailableBalance), balance.Currency)
```

#### Send Money (Bank Transfer)

```go
reference := absa.GenerateReference()
amount, _ := decimal.NewFromString("1000.00")

sendMoneyReq := api.SendMoneyRequest{
    SourceAccount:       "1234567890",
    DestinationAccount:  "0987654321",
    DestinationBankCode: "123",
    Amount:              amount,
    Currency:            "KES",
    Reference:           reference,
    Description:         "Payment for services",
    BeneficiaryName:     "John Doe",
}

sendMoney, err := client.SendMoney(sendMoneyReq)
if err != nil {
    // Handle error
}

fmt.Printf("Transaction ID: %s, Status: %s\n", sendMoney.TransactionID, sendMoney.Status)
```

### Advanced Features

#### Process Bulk Payments

```go
bulkRef := absa.GenerateReference()
bulkItems := []api.BulkPaymentItem{
    {
        DestinationAccount:  "1111222233",
        DestinationBankCode: "123",
        Amount:              decimal.NewFromFloat(500.00),
        Reference:           absa.GenerateReference(),
        Description:         "Salary payment",
        BeneficiaryName:     "Employee One",
    },
    {
        DestinationAccount:  "4444555566",
        DestinationBankCode: "123",
        Amount:              decimal.NewFromFloat(750.00),
        Reference:           absa.GenerateReference(),
        Description:         "Salary payment",
        BeneficiaryName:     "Employee Two",
    },
}

bulkReq := api.BulkPaymentRequest{
    SourceAccount:  "1234567890",
    Currency:       "KES",
    BatchReference: bulkRef,
    Items:          bulkItems,
}

bulkPayment, err := client.ProcessBulkPayment(bulkReq)
if err != nil {
    // Handle error
}

fmt.Printf("Batch ID: %s, Status: %s, Success Count: %d\n", 
    bulkPayment.BatchID, 
    bulkPayment.Status, 
    bulkPayment.SuccessCount)
```

#### Create Standing Order

```go
standingOrderRef := absa.GenerateReference()
standingOrderAmount, _ := decimal.NewFromString("1500.00")
startDate := time.Now().AddDate(0, 0, 1) // Start tomorrow
endDate := time.Now().AddDate(0, 6, 0)   // End after 6 months

standingOrderReq := api.StandingOrderRequest{
    SourceAccount:       "1234567890",
    DestinationAccount:  "9876543210",
    DestinationBankCode: "123",
    Amount:              standingOrderAmount,
    Currency:            "KES",
    Reference:           standingOrderRef,
    Description:         "Monthly rent payment",
    Frequency:           api.FrequencyMonthly,
    StartDate:           startDate,
    EndDate:             endDate,
    BeneficiaryName:     "Landlord Company Ltd",
}

standingOrder, err := client.CreateStandingOrder(standingOrderReq)
if err != nil {
    // Handle error
}

fmt.Printf("Standing Order ID: %s, Status: %s\n", 
    standingOrder.OrderID, 
    standingOrder.Status)
```

#### Foreign Exchange Transfer

```go
// First get the exchange rate
forexRateReq := api.ForexRateRequest{
    SourceCurrency:      "KES",
    DestinationCurrency: "USD",
}

forexRate, err := client.GetForexRate(forexRateReq)
if err != nil {
    // Handle error
}

fmt.Printf("Exchange Rate: 1 %s = %s %s\n", 
    forexRate.SourceCurrency, 
    api.FormatAmount(forexRate.Rate), 
    forexRate.DestinationCurrency)

// Then process the forex transfer
forexAmount, _ := decimal.NewFromString("5000.00")
forexTransferReq := api.ForexTransferRequest{
    SourceAccount:       "1234567890",
    DestinationAccount:  "9876543210",
    DestinationBankCode: "123",
    SourceAmount:        forexAmount,
    SourceCurrency:      "KES",
    DestinationCurrency: "USD",
    Reference:           absa.GenerateReference(),
    Description:         "International payment",
    BeneficiaryName:     "Global Supplier Inc.",
}

forexTransfer, err := client.ProcessForexTransfer(forexTransferReq)
if err != nil {
    // Handle error
}

fmt.Printf("Transaction ID: %s, Status: %s\n", 
    forexTransfer.TransactionID, 
    forexTransfer.Status)
```

## Webhook Handling

To handle webhooks from Absa:

```go
// Set the webhook secret
client.SetWebhookSecret("your-webhook-secret")

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
}

// In your HTTP handler
http.HandleFunc("/webhooks/absa", func(w http.ResponseWriter, r *http.Request) {
    err := client.HandleWebhook(w, r, handlers)
    if err != nil {
        // Handle error
    }
})
```

## Error Handling

The library provides detailed error messages to help diagnose issues:

```go
resp, err := client.SendMoney(req)
if err != nil {
    if apiErr, ok := err.(*api.APIError); ok {
        fmt.Printf("API Error: %s, Code: %s\n", apiErr.Message, apiErr.Code)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    return
}
```

## Environment Variables

The following environment variables are used:

- `ABSA_CLIENT_ID`: Your Absa API client ID
- `ABSA_CLIENT_SECRET`: Your Absa API client secret
- `ABSA_API_KEY`: Your Absa API key

## License

This project is licensed under the MIT License - see the LICENSE file for details.
