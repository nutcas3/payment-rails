package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nutcas3/payment-rails/jenga"
	"github.com/nutcas3/payment-rails/jenga/pkg/api"
)

func main() {
	// Initialize Jenga client with credentials
	client, err := jenga.NewClient(
		os.Getenv("JENGA_API_KEY"),
		os.Getenv("JENGA_MERCHANT_CODE"),
		os.Getenv("JENGA_CONSUMER_SECRET"),
		os.Getenv("JENGA_PRIVATE_KEY_PATH"),
		"sandbox", // Use "production" for production environment
	)
	if err != nil {
		log.Fatalf("Error initializing Jenga client: %v", err)
	}

	// Generate a unique reference number
	reference := jenga.GenerateReference()

	// Current date in YYYY-MM-DD format
	currentDate := time.Now().Format("2006-01-02")

	// Example: Send Money within Equity Bank (Internal Bank Transfer)
	fmt.Println("Sending money within Equity Bank...")

	// Create the internal bank transfer request
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
			Reference:    reference, // Unique transaction reference
			Date:         currentDate,
			Description:  "Monthly rent payment",
		},
	}

	// Send the internal bank transfer request
	response, err := client.SendInternalBankTransfer(internalTransferReq)
	if err != nil {
		log.Fatalf("Error sending internal bank transfer: %v", err)
	}

	// Print the response
	fmt.Printf("\nTransaction Response:\n")
	fmt.Printf("Status: %t\n", response.Status)
	fmt.Printf("Code: %d\n", response.Code)
	fmt.Printf("Message: %s\n", response.Message)
	fmt.Printf("Reference: %s\n", response.Reference)
	fmt.Printf("Transaction ID: %s\n", response.Data.TransactionID)
	fmt.Printf("Transaction Status: %s\n", response.Data.Status)

	// Example: Cross-border internal transfer (e.g., Kenya to Uganda)
	fmt.Println("\nSending money across borders within Equity Bank...")

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
			Reference:    jenga.GenerateReference(), // New unique reference
			Date:         currentDate,
			Description:  "Business payment",
		},
	}

	// Send the cross-border internal bank transfer request
	crossBorderResponse, err := client.SendInternalBankTransfer(crossBorderReq)
	if err != nil {
		log.Fatalf("Error sending cross-border internal bank transfer: %v", err)
	}

	// Print the response
	fmt.Printf("\nCross-Border Transaction Response:\n")
	fmt.Printf("Status: %t\n", crossBorderResponse.Status)
	fmt.Printf("Code: %d\n", crossBorderResponse.Code)
	fmt.Printf("Message: %s\n", crossBorderResponse.Message)
	fmt.Printf("Reference: %s\n", crossBorderResponse.Reference)
	fmt.Printf("Transaction ID: %s\n", crossBorderResponse.Data.TransactionID)
	fmt.Printf("Transaction Status: %s\n", crossBorderResponse.Data.Status)
}
