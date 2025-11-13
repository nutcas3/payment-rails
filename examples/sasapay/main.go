package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shopspring/decimal"

	"github.com/nutcas3/payment-rails/sasapay"
	"github.com/nutcas3/payment-rails/sasapay/pkg/api"
)

func main() {
	// Initialize client with credentials from environment variables
	// You can set these environment variables or pass them directly
	clientID := os.Getenv("SASAPAY_CLIENT_ID")
	clientSecret := os.Getenv("SASAPAY_CLIENT_SECRET")
	environment := "sandbox" // Use "production" for live environment

	client, err := sasapay.NewClient(clientID, clientSecret, environment)
	if err != nil {
		fmt.Printf("Error initializing client: %v\n", err)
		return
	}

	// Set webhook secret for validating webhook signatures
	client.SetWebhookSecret(os.Getenv("SASAPAY_WEBHOOK_SECRET"))

	// Example 1: Customer to Business (C2B) Payment
	fmt.Println("\n=== Example 1: Customer to Business (C2B) Payment ===")
	c2bExample(client)

	// Example 2: Business to Customer (B2C) Payment
	fmt.Println("\n=== Example 2: Business to Customer (B2C) Payment ===")
	b2cExample(client)

	// Example 3: Business to Business (B2B) Payment
	fmt.Println("\n=== Example 3: Business to Business (B2B) Payment ===")
	b2bExample(client)

	// Example 4: Wallet as a Service
	fmt.Println("\n=== Example 4: Wallet as a Service ===")
	waasExample(client)

	// Example 5: Transaction Status and Verification
	fmt.Println("\n=== Example 5: Transaction Status and Verification ===")
	transactionStatusExample(client)

	// Example 6: Cross-Region Transfer
	fmt.Println("\n=== Example 6: Cross-Region Transfer ===")
	crossRegionExample(client)
}

// Example 1: Customer to Business (C2B) Payment
func c2bExample(client *sasapay.Client) {
	// Generate a unique reference for the transaction
	reference := sasapay.GenerateReference()

	// Create C2B request
	amount, _ := decimal.NewFromString("100.00")
	c2bReq := api.C2BRequest{
		MerchantCode: "MERCHANT123",
		PhoneNumber:  "254712345678",
		Amount:       amount,
		Reference:    reference,
		Description:  "Payment for goods",
		CallbackURL:  "https://example.com/callbacks/c2b",
	}

	// Send C2B request
	c2bResp, err := client.CustomerToBusiness(c2bReq)
	if err != nil {
		fmt.Printf("C2B Error: %v\n", err)
		return
	}

	// Print response
	fmt.Printf("C2B Transaction ID: %s\n", c2bResp.TransactionID)
	fmt.Printf("C2B Status: %s\n", c2bResp.Status)
	fmt.Printf("C2B Message: %s\n", c2bResp.Message)
	fmt.Printf("C2B Timestamp: %s\n", c2bResp.Timestamp.Format(time.RFC3339))
}

// Example 2: Business to Customer (B2C) Payment
func b2cExample(client *sasapay.Client) {
	// Generate a unique reference for the transaction
	reference := sasapay.GenerateReference()

	// Create B2C request
	amount, _ := decimal.NewFromString("50.00")
	b2cReq := api.B2CRequest{
		MerchantCode: "MERCHANT123",
		PhoneNumber:  "254712345678",
		Amount:       amount,
		Reference:    reference,
		Description:  "Refund for returned goods",
		CallbackURL:  "https://example.com/callbacks/b2c",
	}

	// Send B2C request
	b2cResp, err := client.BusinessToCustomer(b2cReq)
	if err != nil {
		fmt.Printf("B2C Error: %v\n", err)
		return
	}

	// Print response
	fmt.Printf("B2C Transaction ID: %s\n", b2cResp.TransactionID)
	fmt.Printf("B2C Status: %s\n", b2cResp.Status)
	fmt.Printf("B2C Message: %s\n", b2cResp.Message)
	fmt.Printf("B2C Timestamp: %s\n", b2cResp.Timestamp.Format(time.RFC3339))
}

// Example 3: Business to Business (B2B) Payment
func b2bExample(client *sasapay.Client) {
	// Generate a unique reference for the transaction
	reference := sasapay.GenerateReference()

	// Create B2B request
	amount, _ := decimal.NewFromString("1000.00")
	b2bReq := api.B2BRequest{
		SourceMerchantCode:      "MERCHANT123",
		DestinationMerchantCode: "MERCHANT456",
		Amount:                  amount,
		Reference:               reference,
		Description:             "Payment for services rendered",
		CallbackURL:             "https://example.com/callbacks/b2b",
	}

	// Send B2B request
	b2bResp, err := client.BusinessToBusiness(b2bReq)
	if err != nil {
		fmt.Printf("B2B Error: %v\n", err)
		return
	}

	// Print response
	fmt.Printf("B2B Transaction ID: %s\n", b2bResp.TransactionID)
	fmt.Printf("B2B Status: %s\n", b2bResp.Status)
	fmt.Printf("B2B Message: %s\n", b2bResp.Message)
	fmt.Printf("B2B Timestamp: %s\n", b2bResp.Timestamp.Format(time.RFC3339))
}

// Example 4: Wallet as a Service
func waasExample(client *sasapay.Client) {
	// Create a wallet
	createWalletReq := api.CreateWalletRequest{
		PhoneNumber: "254712345678",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		IDNumber:    "12345678",
		CallbackURL: "https://example.com/callbacks/wallet",
	}

	createWalletResp, err := client.CreateWallet(createWalletReq)
	if err != nil {
		fmt.Printf("Create Wallet Error: %v\n", err)
		return
	}

	fmt.Printf("Wallet ID: %s\n", createWalletResp.WalletID)
	fmt.Printf("Wallet Status: %s\n", createWalletResp.Status)
	fmt.Printf("Wallet Message: %s\n", createWalletResp.Message)

	// Get wallet balance
	balanceReq := api.WalletBalanceRequest{
		WalletID: createWalletResp.WalletID,
	}

	balanceResp, err := client.GetWalletBalance(balanceReq)
	if err != nil {
		fmt.Printf("Get Balance Error: %v\n", err)
		return
	}

	fmt.Printf("Wallet Balance: %s %s\n", balanceResp.Balance.String(), balanceResp.Currency)

	// Transfer between wallets
	transferReq := api.WalletTransferRequest{
		SourceWalletID:      createWalletResp.WalletID,
		DestinationWalletID: "WALLET456",
		Amount:              decimal.NewFromFloat(25.00),
		Reference:           sasapay.GenerateReference(),
		Description:         "Transfer funds between wallets",
		CallbackURL:         "https://example.com/callbacks/transfer",
	}

	transferResp, err := client.TransferToWallet(transferReq)
	if err != nil {
		fmt.Printf("Transfer Error: %v\n", err)
		return
	}

	fmt.Printf("Transfer Transaction ID: %s\n", transferResp.TransactionID)
	fmt.Printf("Transfer Status: %s\n", transferResp.Status)
	fmt.Printf("Transfer Message: %s\n", transferResp.Message)

	// Get wallet statement
	statementReq := api.WalletStatementRequest{
		WalletID:  createWalletResp.WalletID,
		StartDate: time.Now().AddDate(0, -1, 0), // 1 month ago
		EndDate:   time.Now(),
	}

	statementResp, err := client.GetWalletStatement(statementReq)
	if err != nil {
		fmt.Printf("Statement Error: %v\n", err)
		return
	}

	fmt.Printf("Statement Wallet ID: %s\n", statementResp.WalletID)
	fmt.Printf("Statement Transactions: %d\n", len(statementResp.Transactions))

	// Print first few transactions if available
	if len(statementResp.Transactions) > 0 {
		fmt.Println("Recent transactions:")
		limit := 3
		if len(statementResp.Transactions) < limit {
			limit = len(statementResp.Transactions)
		}
		
		for i := 0; i < limit; i++ {
			tx := statementResp.Transactions[i]
			fmt.Printf("  - %s: %s %s (%s)\n", 
				tx.Type, 
				tx.Amount.String(), 
				tx.Description,
				tx.Timestamp.Format(time.RFC3339))
		}
	}
}

// Example 5: Transaction Status and Verification
func transactionStatusExample(client *sasapay.Client) {
	// Use a transaction ID from a previous example
	// In a real application, you would use an actual transaction ID
	transactionID := "TXN123456789"

	// Check transaction status
	statusReq := api.TransactionStatusRequest{
		TransactionID: transactionID,
	}

	statusResp, err := client.CheckTransactionStatus(statusReq)
	if err != nil {
		fmt.Printf("Status Check Error: %v\n", err)
		return
	}

	fmt.Printf("Transaction Status: %s\n", statusResp.Status)
	fmt.Printf("Transaction Message: %s\n", statusResp.Message)

	// Verify transaction
	verifyReq := api.VerifyTransactionRequest{
		TransactionID: transactionID,
	}

	verifyResp, err := client.VerifyTransaction(verifyReq)
	if err != nil {
		fmt.Printf("Verification Error: %v\n", err)
		return
	}

	fmt.Printf("Verified Transaction ID: %s\n", verifyResp.TransactionID)
	fmt.Printf("Verified Amount: %s %s\n", verifyResp.Amount.String(), verifyResp.Currency)
	fmt.Printf("Verified Status: %s\n", verifyResp.Status)
	fmt.Printf("Verified Message: %s\n", verifyResp.Message)
}

// Example 6: Cross-Region Transfer
func crossRegionExample(client *sasapay.Client) {
	// Get a quote for cross-region transfer
	quoteReq := api.CrossRegionQuoteRequest{
		SourceRegion:      api.RegionKenya,
		DestinationRegion: api.RegionUganda,
		SourceCurrency:    "KES",
		DestCurrency:      "UGX",
		Amount:            decimal.NewFromFloat(1000.00),
	}

	quoteResp, err := client.GetCrossRegionQuote(quoteReq)
	if err != nil {
		fmt.Printf("Quote Error: %v\n", err)
		return
	}

	fmt.Printf("Quote ID: %s\n", quoteResp.QuoteID)
	fmt.Printf("Source Amount: %s %s\n", quoteResp.SourceAmount.String(), quoteResp.SourceCurrency)
	fmt.Printf("Destination Amount: %s %s\n", quoteResp.DestinationAmount.String(), quoteResp.DestinationCurrency)
	fmt.Printf("Exchange Rate: %s\n", quoteResp.ExchangeRate.String())
	fmt.Printf("Fee: %s %s\n", quoteResp.Fee.String(), quoteResp.SourceCurrency)
	fmt.Printf("Total Cost: %s %s\n", quoteResp.TotalCost.String(), quoteResp.SourceCurrency)
	fmt.Printf("Expires At: %s\n", quoteResp.ExpiresAt.Format(time.RFC3339))

	// Initiate cross-region transfer
	transferReq := api.CrossRegionTransferRequest{
		SourceRegion:      api.RegionKenya,
		DestinationRegion: api.RegionUganda,
		SourceCurrency:    "KES",
		DestCurrency:      "UGX",
		Amount:            decimal.NewFromFloat(1000.00),
		PhoneNumber:       "256712345678", // Uganda phone number
		Reference:         sasapay.GenerateReference(),
		Description:       "Cross-region transfer to Uganda",
		CallbackURL:       "https://example.com/callbacks/cross-region",
	}

	transferResp, err := client.CrossRegionTransfer(transferReq)
	if err != nil {
		fmt.Printf("Transfer Error: %v\n", err)
		return
	}

	fmt.Printf("Transaction ID: %s\n", transferResp.TransactionID)
	fmt.Printf("Status: %s\n", transferResp.Status)
	fmt.Printf("Message: %s\n", transferResp.Message)
	fmt.Printf("Source Amount: %s %s\n", transferResp.SourceAmount.String(), transferResp.SourceCurrency)
	fmt.Printf("Destination Amount: %s %s\n", transferResp.DestinationAmount.String(), transferResp.DestinationCurrency)
	fmt.Printf("Exchange Rate: %s\n", transferResp.ExchangeRate.String())
}

// Example of webhook handling (would be used in an HTTP handler)
// This function is commented out to avoid the unused variable lint error
/*
func webhookHandlingExample(client *sasapay.Client) {
	// Define webhook handlers
	handlers := api.WebhookHandlers{
		PaymentReceived: func(event api.WebhookEvent) {
			fmt.Printf("Payment received: %s for %s %s\n", 
				event.TransactionID, 
				event.Amount.String(), 
				event.Currency)
		},
		PaymentCompleted: func(event api.WebhookEvent) {
			fmt.Printf("Payment completed: %s for %s %s\n", 
				event.TransactionID, 
				event.Amount.String(), 
				event.Currency)
		},
		PaymentFailed: func(event api.WebhookEvent) {
			fmt.Printf("Payment failed: %s - %s\n", 
				event.TransactionID, 
				event.Message)
		},
		WalletCreated: func(event api.WebhookEvent) {
			fmt.Printf("Wallet created: %s\n", event.Reference)
		},
		WalletTransferred: func(event api.WebhookEvent) {
			fmt.Printf("Wallet transfer: %s for %s %s\n", 
				event.TransactionID, 
				event.Amount.String(), 
				event.Currency)
		},
	}

	// In an HTTP handler, you would use:
	// client.ProcessWebhookRequest(r.Body, r.Header.Get("X-SasaPay-Signature"), handlers)
}
*/
