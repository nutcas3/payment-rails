package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"payment-rails/jenga"
	"payment-rails/jenga/pkg/api"
)

func main() {
	// Initialize Jenga client with credentials from environment variables
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

	// Create a request for RTGS transfer
	req := api.SendMoneyRequest{
		Source: api.Source{
			CountryCode:   "KE", // Kenya
			Name:          "John Doe",
			AccountNumber: "1234567890",
		},
		Destination: api.Destination{
			Type:          "bank",
			CountryCode:   "KE", // Kenya
			Name:          "Jane Smith",
			AccountNumber: "0987654321",
			BankCode:      "01", // Bank code for the receiving bank
			BranchCode:    "112", // Branch code for the receiving bank
		},
		Transfer: api.Transfer{
			Type:         api.TransferTypeRTGS, // Specify RTGS transfer type
			Amount:       "50000.00", // RTGS is typically used for larger amounts
			CurrencyCode: "KES",
			Reference:    fmt.Sprintf("RTGS-REF-%d", time.Now().Unix()),
			Date:         time.Now().Format("2006-01-02"),
			Description:  "RTGS Transfer to Jane Smith",
		},
	}

	// Send the money using RTGS
	response, err := client.SendMoney(req)
	if err != nil {
		log.Fatalf("Error sending money via RTGS: %v", err)
	}

	// Print the response
	fmt.Printf("RTGS Transfer Status: %v\n", response.Status)
	fmt.Printf("Message: %s\n", response.Message)
	fmt.Printf("Reference: %s\n", response.Reference)
	fmt.Printf("Transaction ID: %s\n", response.Data.TransactionID)
	fmt.Printf("Transaction Status: %s\n", response.Data.Status)
}
