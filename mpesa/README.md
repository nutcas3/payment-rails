# Mpesa SDK for Go (Daraja API)

A Golang client for the Safaricom Mpesa Daraja API. This SDK provides a simple interface to interact with the Mpesa API for various payment operations.

## Features

- STK Push (Lipa Na M-Pesa Online)
- M-Pesa Express Query (STK Push Query)
- Customer to Business (C2B) URL Registration
- Customer to Business (C2B) Simulation
- Business to Customer (B2C) Payment
- Business to Business (B2B) Payment
- Business Pay Bill
- B2C Account Top Up
- B2B Express CheckOut (USSD Push to Till)
- Tax Remittance to KRA
- Transaction Status Query
- Account Balance Query
- Payment Reversal
- Dynamic QR Code Generation
- M-Pesa Ratiba (Standing Order) API

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

### M-Pesa Express Query (STK Push Query)

```go
// Check the status of a previous STK Push transaction
response, err := client.QueryStkPushStatus(mpesa.STKPushQueryRequest{
    BusinessShortCode: "174379",                      // Your business shortcode
    CheckoutRequestID: "ws_CO_260520211133524545",    // The CheckoutRequestID from the STK Push response
})
if err != nil {
    log.Fatalf("Failed to query STK Push status: %v", err)
}

// Interpret the response
if response.ResultCode == "0" {
    fmt.Println("Transaction was successful!")
} else if response.ResultCode == "1032" {
    fmt.Println("Transaction was cancelled by the user")
} else {
    fmt.Println("Transaction failed or is still being processed")
}
```

### Customer to Business (C2B) URL Registration

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

### Business Pay Bill

```go
// Make a Business Pay Bill payment
response, err := client.BusinessPayBill(mpesa.BusinessPayBillRequest{
    Initiator:          "testapi",                      // API Username
    SecurityCredential: securityCredential,             // Encrypted password
    Amount:             "100",                          // Amount to pay
    PartyA:             "123456",                       // Your business shortcode
    PartyB:             "000000",                       // Recipient paybill number
    AccountReference:   "INV001",                       // Account reference (up to 13 chars)
    Requester:          "254700000000",                 // Optional: Customer phone number
    Remarks:            "Payment for utility bill",     // Transaction remarks
    QueueTimeOutURL:    "https://example.com/timeout",  // Timeout URL
    ResultURL:          "https://example.com/result",   // Result URL
    Occasion:           "Monthly Payment",              // Optional: Additional information
})
if err != nil {
    log.Fatalf("Failed to make Business Pay Bill payment: %v", err)
}
fmt.Printf("Business Pay Bill Response: %+v\n", response)
```

### B2C Account Top Up

```go
// Load funds to a B2C shortcode for disbursement
response, err := client.B2CAccountTopUp(mpesa.B2CTopUpRequest{
    Initiator:          "testapi",                      // API Username
    SecurityCredential: securityCredential,             // Encrypted password
    Amount:             "10000",                        // Amount to load (in smallest currency unit)
    PartyA:             "600979",                       // Your business shortcode
    PartyB:             "600000",                       // B2C shortcode to be funded
    AccountReference:   "TopUp001",                     // Optional: Account reference
    Requester:          "254708374149",                 // Optional: Requester phone number
    Remarks:            "B2C account funding",          // Transaction remarks
    QueueTimeOutURL:    "https://example.com/timeout",  // Timeout URL
    ResultURL:          "https://example.com/result",   // Result URL
})
if err != nil {
    log.Fatalf("Failed to make B2C Account Top Up: %v", err)
}
fmt.Printf("B2C Account Top Up Response: %+v\n", response)
```

### B2B Express CheckOut (USSD Push to Till)

```go
// Generate a unique request reference ID
requestRefID := fmt.Sprintf("REF-%d", time.Now().Unix())

// Initiate USSD Push to Till (B2B Express CheckOut)
response, err := client.UssdPush(mpesa.UssdPushRequest{
    PrimaryShortCode:  "000001",                    // Merchant's till number (sending money)
    ReceiverShortCode: "000002",                    // Vendor's paybill (receiving money)
    Amount:            "100",                       // Amount to send
    PaymentRef:        "INV12345",                  // Payment reference
    CallbackURL:       "https://example.com/callback", // Callback URL for transaction result
    PartnerName:       "ACME Store",                // Vendor's friendly name
    RequestRefID:      requestRefID,                // Unique request reference ID
})
if err != nil {
    log.Fatalf("Failed to initiate USSD Push: %v", err)
}
fmt.Printf("USSD Push Response: %+v\n", response)

// The merchant will receive a USSD prompt to enter their operator ID, PIN, and confirm payment
// The final result will be sent to your callback URL
```

### Tax Remittance to KRA

```go
// Remit tax to KRA
response, err := client.RemitTax(mpesa.TaxRemittanceRequest{
    Initiator:          "testapi",                      // API Username
    SecurityCredential: securityCredential,             // Encrypted password
    Amount:             "239",                          // Amount to remit
    PartyA:             "888880",                       // Your business shortcode
    AccountReference:   "353353",                       // Payment Registration Number (PRN) from KRA
    Remarks:            "Tax payment for Q2 2025",      // Transaction remarks
    QueueTimeOutURL:    "https://example.com/timeout",  // Timeout URL
    ResultURL:          "https://example.com/result",   // Result URL
})
if err != nil {
    log.Fatalf("Failed to remit tax: %v", err)
}
fmt.Printf("Tax Remittance Response: %+v\n", response)

// Note: Prior integration with KRA is required for tax declaration and PRN generation
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

### Dynamic QR Code Generation

```go
// Generate a Dynamic QR Code for M-Pesa transactions
qrResponse, err := client.GenerateQRCode(mpesa.QRCodeRequest{
    MerchantName: "TEST Business",               // Merchant Name
    RefNo:        "Invoice123",                  // Reference Number
    Amount:       1000,                          // Amount (in smallest currency unit)
    TrxCode:      mpesa.TrxCodeBuyGoods,         // Transaction Code (BG, WA, PB, SM, SB)
    CPI:          "254708374149",                // Customer Phone/Till/Paybill
    Size:         "300",                         // QR Code size in pixels
})
if err != nil {
    log.Fatalf("Failed to generate QR code: %v", err)
}
fmt.Printf("QR Code Response: %+v\n", qrResponse)

// The QR Code is returned as a base64-encoded string in qrResponse.QRCode
// You can convert this to an image and display it in your application
```

### M-Pesa Ratiba (Standing Order) API

```go
// Create a standing order (M-Pesa Ratiba) for recurring payments
ratibaResponse, err := client.CreateStandingOrder(mpesa.RatibaRequest{
    StandingOrderName: "Monthly Rent Payment",
    StartDate:         "20250901",                     // Format: YYYYMMDD
    EndDate:           "20260901",                     // Format: YYYYMMDD
    BusinessShortCode: "174379",                       // Your business short code
    TransactionType:   mpesa.TransactionTypePayBill,   // For Paybill
    IdentifierType:    mpesa.ReceiverTypePaybill,      // For Paybill
    Amount:            "5000",                         // Amount in smallest currency unit
    PhoneNumber:       "254708374149",                 // Customer's phone number
    CallBackURL:       "https://example.com/callback", // Your callback URL
    AccountReference:  "Rent123",                      // Account reference
    TransactionDesc:   "Monthly Rent",                 // Transaction description
    Frequency:         mpesa.FrequencyMonthly,         // Monthly payments
})
if err != nil {
    log.Fatalf("Failed to create standing order: %v", err)
}
fmt.Printf("Standing Order Response: %+v\n", ratibaResponse)

// Note: The M-Pesa Ratiba API is a commercial API that requires a contract with Safaricom.
// For production use, you need to contact Safaricom at apisupport@safaricom.co.ke after testing.
```

## Testing

To run the tests:

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
