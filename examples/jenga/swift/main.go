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

	// Create a request for SWIFT international transfer
	req := api.SendMoneyRequest{
		Source: api.Source{
			CountryCode:   "KE", // Kenya
			Name:          "John Doe",
			AccountNumber: "1234567890",
		},
		Destination: api.Destination{
			Type:          "bank",
			CountryCode:   "US", // United States
			Name:          "Jane Smith",
			AccountNumber: "0987654321",
			// SWIFT specific fields
			BankName:      "Bank of America",
			BankAddress:   "100 North Tryon Street, Charlotte, NC 28255, USA",
			SwiftCode:     "BOFAUS3N", // Example SWIFT code for Bank of America
			RoutingNumber: "026009593", // ABA routing number for US banks
		},
		Transfer: api.Transfer{
			Type:         api.TransferTypeSWIFT, // Specify SWIFT transfer type
			Amount:       "1000.00",
			CurrencyCode: "USD", // International transfer in USD
			Reference:    fmt.Sprintf("SWIFT-REF-%d", time.Now().Unix()),
			Date:         time.Now().Format("2006-01-02"),
			Description:  "International SWIFT Transfer to Jane Smith",
		},
	}

	// Send the money using SWIFT
	response, err := client.SendMoney(req)
	if err != nil {
		log.Fatalf("Error sending money via SWIFT: %v", err)
	}

	// Print the response
	fmt.Printf("SWIFT Transfer Status: %v\n", response.Status)
	fmt.Printf("Message: %s\n", response.Message)
	fmt.Printf("Reference: %s\n", response.Reference)
	fmt.Printf("Transaction ID: %s\n", response.Data.TransactionID)
	fmt.Printf("Transaction Status: %s\n", response.Data.Status)
}
