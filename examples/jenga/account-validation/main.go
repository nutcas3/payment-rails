package main

import (
	"fmt"
	"log"
	"os"
	"payment-rails/jenga"
	"payment-rails/jenga/pkg/api"
)

func main() {
	// Get environment variables
	apiKey := os.Getenv("JENGA_API_KEY")
	merchantCode := os.Getenv("JENGA_MERCHANT_CODE")
	consumerSecret := os.Getenv("JENGA_CONSUMER_SECRET")
	privateKeyPath := os.Getenv("JENGA_PRIVATE_KEY_PATH")

	if apiKey == "" || merchantCode == "" || consumerSecret == "" || privateKeyPath == "" {
		log.Fatal("Missing required environment variables. Please set JENGA_API_KEY, JENGA_MERCHANT_CODE, JENGA_CONSUMER_SECRET, and JENGA_PRIVATE_KEY_PATH")
	}

	// Initialize Jenga client
	client, err := jenga.NewClient(apiKey, merchantCode, consumerSecret, privateKeyPath, "sandbox")
	if err != nil {
		log.Fatalf("Error initializing Jenga client: %v", err)
	}

	fmt.Println("Example: Validate Equity Bank Account")
	fmt.Println("-------------------------------------")

	// Create account validation request
	validateReq := api.AccountValidateRequest{
		CountryCode:     "UG",
		AccountNumber:   "1036200681230",
		AccountFullName: "DICKSON MAITEI",
		ChargeAccount:   "1036200681230", // Optional
	}

	// Validate account
	response, err := client.ValidateAccount(validateReq)
	if err != nil {
		log.Fatalf("Error validating account: %v", err)
	}

	// Process response
	if response.Status {
		fmt.Println("Account validation successful!")
		fmt.Printf("Full Name: %s\n", response.Data.Account.FullNames)
		fmt.Printf("Account Number: %s\n", response.Data.Account.AccountNumber)
		fmt.Printf("Currency: %s\n", response.Data.Account.Currency)
		fmt.Printf("Account Status: %s\n", response.Data.Account.Status)
	} else {
		fmt.Printf("Account validation failed. Code: %d, Message: %s\n", response.Code, response.Message)
	}
}
