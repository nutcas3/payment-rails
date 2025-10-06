# FNB Integration Channel SDK

A comprehensive Go SDK for integrating with FNB's Integration Channel services in South Africa. This SDK provides seamless access to EFT payments, collections, DebiCheck mandates, transaction history, and more.

## Features

### Payment Services
- **EFT Domestic Payments** - Electronic funds transfers to any South African bank account
- **Urgent Payments** - Immediate payment processing
- **Batch Payments** - Process multiple payments in a single request
- **Payment Status Tracking** - Real-time payment status updates
- **Payment Cancellation** - Cancel pending payments

### Collection Services
- **EFT Domestic Collections** - Collect funds from customer accounts
- **Batch Collections** - Process multiple collections efficiently
- **Collection Status Tracking** - Monitor collection progress
- **Dispute Management** - Handle collection disputes

### DebiCheck Mandates
- **Mandate Creation** - Create authenticated debit order mandates
- **Mandate Management** - Update, suspend, cancel, and reinstate mandates
- **Mandate Verification** - Verify mandate validity before collection
- **Collection Against Mandates** - Collect funds using active mandates

### Account Services
- **Account Verification (AVS)** - Verify account details before transactions
- **Transaction History** - Retrieve detailed transaction records
- **Account Balance** - Check current and available balances
- **Statements** - Generate and retrieve account statements
- **Proof of Payment** - Download payment confirmation documents

### Notifications & Webhooks
- **Real-Time Notifications** - Receive instant transaction notifications
- **Webhook Support** - Handle asynchronous event callbacks
- **Event Types** - Payment completed, collection status, mandate approvals, etc.

### Host-to-Host (H2H) Support
- **File-Based Integration** - Support for legacy H2H file formats
- **FTP Upload/Download** - Automated file transfer
- **Response File Parsing** - Parse FNB response files

## Installation

```bash
go get github.com/nutcas3/payment-rails/fnb
```

## Quick Start

### Initialize the Client

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "payment-rails/fnb"
)

func main() {
    // Create client with credentials
    client := fnb.NewClient(
        "your-client-id",
        "your-client-secret",
        "your-api-key",
        fnb.WithEnvironment(fnb.EnvironmentSandbox),
    )
    
    ctx := context.Background()
    
    // Client is ready to use
    fmt.Println("FNB client initialized successfully")
}
```

## Usage Examples

### 1. Send an EFT Payment

```go
import (
    "context"
    "fmt"
    "payment-rails/fnb/pkg/api"
)

func sendPayment(client *fnb.Client) {
    ctx := context.Background()
    
    req := api.EFTPaymentRequest{
        SourceAccountNumber:      "1234567890",
        BeneficiaryAccountNumber: "9876543210",
        BeneficiaryName:          "John Doe",
        BeneficiaryBankCode:      "250655", // Universal branch code
        Amount:                   1500.00,
        Currency:                 "ZAR",
        PaymentReference:         "INV-2025-001",
        PaymentDescription:       "Invoice payment",
        NotificationEmail:        "customer@example.com",
    }
    
    resp, err := client.CreateEFTPayment(ctx, req)
    if err != nil {
        log.Fatalf("Payment failed: %v", err)
    }
    
    fmt.Printf("Payment submitted successfully!\n")
    fmt.Printf("Transaction ID: %s\n", resp.TransactionID)
    fmt.Printf("Status: %s\n", resp.Status)
    fmt.Printf("Reference: %s\n", resp.PaymentReference)
}
```

### 2. Check Payment Status

```go
func checkPaymentStatus(client *fnb.Client, transactionID string) {
    ctx := context.Background()
    
    status, err := client.GetPaymentStatus(ctx, transactionID)
    if err != nil {
        log.Fatalf("Failed to get status: %v", err)
    }
    
    fmt.Printf("Transaction ID: %s\n", status.TransactionID)
    fmt.Printf("Status: %s\n", status.Status)
    fmt.Printf("Amount: %.2f %s\n", status.Amount, status.Currency)
    fmt.Printf("Beneficiary: %s\n", status.BeneficiaryName)
    
    if status.Status == "COMPLETED" {
        fmt.Printf("Settlement Date: %s\n", status.SettlementDate)
    } else if status.Status == "FAILED" {
        fmt.Printf("Failure Reason: %s\n", status.FailureReason)
    }
}
```

### 3. Create a DebiCheck Mandate

```go
func createMandate(client *fnb.Client) {
    ctx := context.Background()
    
    req := api.MandateRequest{
        CreditorName:          "My Company (Pty) Ltd",
        CreditorAbbreviation:  "MYCO",
        CreditorAccountNumber: "1234567890",
        
        DebtorName:          "Jane Smith",
        DebtorIDNumber:      "8505155009087",
        DebtorAccountNumber: "9876543210",
        DebtorBankCode:      "250655",
        DebtorEmail:         "jane@example.com",
        DebtorMobile:        "+27821234567",
        
        ContractReference:   "SUB-2025-001",
        MaximumAmount:       500.00,
        Currency:            "ZAR",
        FrequencyType:       "MONTHLY",
        FirstCollectionDate: "2025-11-01",
        CollectionDay:       1, // 1st of each month
        
        MandateDescription: "Monthly subscription payment",
    }
    
    resp, err := client.CreateMandate(ctx, req)
    if err != nil {
        log.Fatalf("Mandate creation failed: %v", err)
    }
    
    fmt.Printf("Mandate created successfully!\n")
    fmt.Printf("Mandate ID: %s\n", resp.MandateID)
    fmt.Printf("Status: %s\n", resp.Status)
    fmt.Printf("Contract Reference: %s\n", resp.ContractReference)
    
    if resp.Status == "PENDING_APPROVAL" {
        fmt.Println("Waiting for customer approval...")
    }
}
```

### 4. Collect Against a Mandate

```go
func collectFromMandate(client *fnb.Client, mandateID string) {
    ctx := context.Background()
    
    // First, verify the mandate is valid
    valid, err := client.VerifyMandate(ctx, mandateID, 250.00)
    if err != nil {
        log.Fatalf("Mandate verification failed: %v", err)
    }
    
    if !valid {
        log.Fatal("Mandate is not valid for collection")
    }
    
    // Proceed with collection
    req := api.MandateCollectionRequest{
        MandateID:           mandateID,
        Amount:              250.00,
        CollectionReference: "COLL-2025-001",
        Description:         "Monthly subscription - October 2025",
    }
    
    resp, err := client.CollectAgainstMandate(ctx, req)
    if err != nil {
        log.Fatalf("Collection failed: %v", err)
    }
    
    fmt.Printf("Collection initiated successfully!\n")
    fmt.Printf("Transaction ID: %s\n", resp.TransactionID)
    fmt.Printf("Status: %s\n", resp.Status)
}
```

### 5. Retrieve Transaction History

```go
func getTransactionHistory(client *fnb.Client, accountNumber string) {
    ctx := context.Background()
    
    req := api.TransactionHistoryRequest{
        AccountNumber:   accountNumber,
        FromDate:        "2025-09-01",
        ToDate:          "2025-10-02",
        TransactionType: "ALL", // DEBIT, CREDIT, or ALL
        PageNumber:      1,
        PageSize:        50,
    }
    
    resp, err := client.GetTransactionHistory(ctx, req)
    if err != nil {
        log.Fatalf("Failed to get transaction history: %v", err)
    }
    
    fmt.Printf("Account: %s (%s)\n", resp.AccountNumber, resp.AccountName)
    fmt.Printf("Period: %s to %s\n", resp.FromDate, resp.ToDate)
    fmt.Printf("Total Transactions: %d\n\n", resp.TotalCount)
    
    for _, txn := range resp.Transactions {
        fmt.Printf("%s | %s | %s | %.2f | %.2f\n",
            txn.Date.Format("2006-01-02"),
            txn.Type,
            txn.Description,
            txn.Amount,
            txn.Balance,
        )
    }
}
```

### 6. Verify Account Details

```go
func verifyAccount(client *fnb.Client, accountNumber, bankCode string) {
    ctx := context.Background()
    
    req := api.AccountVerificationRequest{
        AccountNumber: accountNumber,
        BankCode:      bankCode,
    }
    
    resp, err := client.VerifyAccount(ctx, req)
    if err != nil {
        log.Fatalf("Account verification failed: %v", err)
    }
    
    if resp.IsValid && resp.Status == "ACTIVE" {
        fmt.Printf("✓ Account verified successfully\n")
        fmt.Printf("Account Name: %s\n", resp.AccountName)
        fmt.Printf("Account Type: %s\n", resp.AccountType)
        fmt.Printf("Bank: %s\n", resp.BankName)
    } else {
        fmt.Printf("✗ Account verification failed\n")
        fmt.Printf("Status: %s\n", resp.Status)
        fmt.Printf("Message: %s\n", resp.Message)
    }
}
```

### 7. Handle Webhooks

```go
import (
    "net/http"
    "payment-rails/fnb/pkg/api"
)

func setupWebhooks() {
    // Create webhook handler with your secret
    webhookHandler := api.NewWebhookHandler("your-webhook-secret")
    
    // Register handlers for different event types
    webhookHandler.RegisterHandler(api.EventPaymentCompleted, func(event api.WebhookEvent) error {
        fmt.Printf("Payment completed: %s\n", event.ResourceID)
        // Update your database, send notifications, etc.
        return nil
    })
    
    webhookHandler.RegisterHandler(api.EventMandateApproved, func(event api.WebhookEvent) error {
        fmt.Printf("Mandate approved: %s\n", event.ResourceID)
        // Start processing collections
        return nil
    })
    
    webhookHandler.RegisterHandler(api.EventCollectionFailed, func(event api.WebhookEvent) error {
        fmt.Printf("Collection failed: %s\n", event.ResourceID)
        // Handle failed collection, retry logic, etc.
        return nil
    })
    
    // Set up HTTP endpoint
    http.HandleFunc("/webhooks/fnb", webhookHandler.HandleWebhook)
    http.ListenAndServe(":8080", nil)
}
```

### 8. Batch Payments

```go
func sendBatchPayments(client *fnb.Client) {
    ctx := context.Background()
    
    req := api.BatchPaymentRequest{
        BatchReference:      "BATCH-2025-001",
        SourceAccountNumber: "1234567890",
        TotalAmount:         5000.00,
        TotalCount:          3,
        Payments: []api.EFTPaymentRequest{
            {
                BeneficiaryAccountNumber: "1111111111",
                BeneficiaryName:          "Supplier A",
                BeneficiaryBankCode:      "250655",
                Amount:                   2000.00,
                PaymentReference:         "INV-001",
                PaymentDescription:       "Invoice 001",
            },
            {
                BeneficiaryAccountNumber: "2222222222",
                BeneficiaryName:          "Supplier B",
                BeneficiaryBankCode:      "250655",
                Amount:                   1500.00,
                PaymentReference:         "INV-002",
                PaymentDescription:       "Invoice 002",
            },
            {
                BeneficiaryAccountNumber: "3333333333",
                BeneficiaryName:          "Supplier C",
                BeneficiaryBankCode:      "250655",
                Amount:                   1500.00,
                PaymentReference:         "INV-003",
                PaymentDescription:       "Invoice 003",
            },
        },
    }
    
    resp, err := client.CreateBatchPayment(ctx, req)
    if err != nil {
        log.Fatalf("Batch payment failed: %v", err)
    }
    
    fmt.Printf("Batch submitted successfully!\n")
    fmt.Printf("Batch ID: %s\n", resp.BatchID)
    fmt.Printf("Status: %s\n", resp.Status)
    fmt.Printf("Total: %.2f (%d payments)\n", resp.TotalAmount, resp.TotalCount)
}
```

### 9. Host-to-Host Integration

```go
func h2hPaymentExample() {
    // Configure H2H settings
    h2hConfig := &fnb.H2HConfig{
        CertPath:     "/path/to/client-cert.pem",
        KeyPath:      "/path/to/client-key.pem",
        FTPHost:      "ftp.fnb.co.za:21",
        FTPUser:      "your-ftp-username",
        FTPPassword:  "your-ftp-password",
        FTPDirectory: "/upload",
    }
    
    client := fnb.NewClient(
        "client-id",
        "client-secret",
        "api-key",
        fnb.WithH2HConfig(h2hConfig),
    )
    
    // Create H2H client
    h2hClient := api.NewH2HClient(h2hConfig)
    
    // Build payment file
    paymentFile := api.H2HPaymentFile{
        Header: api.H2HFileHeader{
            FileType:       "PAYMENT",
            FileReference:  "PAY20251002001",
            CreationDate:   time.Now(),
            OriginatorCode: "YOUR_CODE",
            OriginatorName: "Your Company",
            TestIndicator:  "T", // "P" for production
        },
        Payments: []api.H2HPaymentRecord{
            {
                RecordType:          "P",
                SequenceNumber:      1,
                BeneficiaryAccount:  "9876543210",
                BeneficiaryName:     "John Doe",
                BankCode:            "250655",
                BranchCode:          "000000",
                Amount:              1500.00,
                PaymentReference:    "INV001",
                BeneficiaryReference: "Payment for services",
                ActionDate:          time.Now().Format("20060102"),
            },
        },
        Trailer: api.H2HFileTrailer{
            RecordType:  "T",
            RecordCount: 1,
            TotalAmount: 1500.00,
            HashTotal:   "9876543210",
        },
    }
    
    // Generate file content
    content, err := h2hClient.GeneratePaymentFile(paymentFile)
    if err != nil {
        log.Fatalf("Failed to generate file: %v", err)
    }
    
    // Upload to FNB
    ctx := context.Background()
    filename := fmt.Sprintf("PAY_%s.txt", time.Now().Format("20060102150405"))
    err = h2hClient.UploadFile(ctx, filename, content)
    if err != nil {
        log.Fatalf("Failed to upload file: %v", err)
    }
    
    fmt.Printf("H2H payment file uploaded: %s\n", filename)
}
```

## Configuration

### Environment Options

```go
// Sandbox environment (for testing)
client := fnb.NewClient(
    clientID, 
    clientSecret, 
    apiKey,
    fnb.WithEnvironment(fnb.EnvironmentSandbox),
)

// Production environment
client := fnb.NewClient(
    clientID, 
    clientSecret, 
    apiKey,
    fnb.WithEnvironment(fnb.EnvironmentProduction),
)

// Custom base URL
client := fnb.NewClient(
    clientID, 
    clientSecret, 
    apiKey,
    fnb.WithBaseURL("https://custom-api.fnb.co.za"),
)
```

### Authentication

The SDK uses OAuth2 client credentials flow for authentication. Access tokens are automatically managed and refreshed as needed.

### Security Best Practices

1. **Store credentials securely** - Use environment variables or secret management services
2. **Enable webhook signature verification** - Always verify webhook signatures in production
3. **Use HTTPS** - Ensure all API calls use HTTPS
4. **Implement idempotency** - Use idempotency keys for payment requests
5. **Certificate management** - Keep H2H certificates secure and rotate regularly

## Error Handling

```go
resp, err := client.CreateEFTPayment(ctx, req)
if err != nil {
    // Check if it's an API error
    if apiErr, ok := err.(*api.ErrorResponse); ok {
        fmt.Printf("API Error [%d]: %s\n", apiErr.Status, apiErr.Error())
        // Handle specific error codes
        switch apiErr.Code {
        case "INSUFFICIENT_FUNDS":
            // Handle insufficient funds
        case "INVALID_ACCOUNT":
            // Handle invalid account
        default:
            // Handle other errors
        }
    } else {
        // Handle network or other errors
        fmt.Printf("Error: %v\n", err)
    }
    return
}
```

## Testing

Run the test suite:

```bash
cd fnb
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run specific tests:

```bash
go test -run TestCreateEFTPayment ./pkg/api
```

## API Reference

### Payment Methods
- `CreateEFTPayment(ctx, req)` - Create an EFT payment
- `CreateUrgentPayment(ctx, req)` - Create an urgent payment
- `GetPaymentStatus(ctx, transactionID)` - Get payment status
- `CreateBatchPayment(ctx, req)` - Create batch payment
- `CancelPayment(ctx, transactionID, reason)` - Cancel a payment

### Collection Methods
- `CreateEFTCollection(ctx, req)` - Create an EFT collection
- `GetCollectionStatus(ctx, transactionID)` - Get collection status
- `CreateBatchCollection(ctx, req)` - Create batch collection
- `DisputeCollection(ctx, req)` - Dispute a collection

### Mandate Methods
- `CreateMandate(ctx, req)` - Create a DebiCheck mandate
- `GetMandateStatus(ctx, mandateID)` - Get mandate status
- `UpdateMandate(ctx, req)` - Update a mandate
- `CancelMandate(ctx, req)` - Cancel a mandate
- `CollectAgainstMandate(ctx, req)` - Collect using a mandate
- `VerifyMandate(ctx, mandateID, amount)` - Verify mandate validity

### Account Methods
- `VerifyAccount(ctx, req)` - Verify account details
- `GetTransactionHistory(ctx, req)` - Get transaction history
- `GetAccountBalance(ctx, accountNumber)` - Get account balance
- `GetStatement(ctx, req)` - Get account statement
- `GetProofOfPayment(ctx, req)` - Get proof of payment

## Support

For issues, questions, or contributions, please visit:
- GitHub: https://github.com/nutcas3/payment-rails
- Documentation: https://www.fnb.co.za/integration-channel/

## License

This SDK is part of the payment-rails project.

## Changelog

### v1.0.0 (2025-10-02)
- Initial release
- EFT Payments and Collections
- DebiCheck Mandates
- Account Verification and Transaction History
- Webhook support
- Host-to-Host integration
- Comprehensive test suite
