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

	// Example: Send Money to Mobile Wallet (M-Pesa)
	fmt.Println("Sending money to M-Pesa mobile wallet...")

	// Create the mobile wallet request
	mobileWalletReq := api.MobileWalletRequest{
		Source: api.Source{
			CountryCode:   "KE", // Kenya
			Name:          "John Doe",
			AccountNumber: "0011547896523", // Source account number
		},
		Destination: struct {
			Type          string `json:"type"`
			CountryCode   string `json:"countryCode"`
			Name          string `json:"name"`
			MobileNumber  string `json:"mobileNumber"`
			WalletName    string `json:"walletName"`
		}{
			Type:         "mobile",
			CountryCode:  "KE", // Kenya
			Name:         "Jane Smith",
			MobileNumber: "254722000000", // Recipient's mobile number with country code
			WalletName:   "Mpesa", // Mobile wallet provider
		},
		Transfer: struct {
			Type         string `json:"type"`
			Amount       string `json:"amount"`
			CurrencyCode string `json:"currencyCode"`
			Reference    string `json:"reference"`
			Date         string `json:"date"`
			Description  string `json:"description"`
			CallbackUrl  string `json:"callbackUrl"`
		}{
			Type:         "MobileWallet",
			Amount:       "200.00", // Amount to send
			CurrencyCode: "KES", // Kenyan Shillings
			Reference:    reference, // Unique transaction reference
			Date:         currentDate,
			Description:  "Payment for services rendered",
			CallbackUrl:  "https://webhook.site/your-webhook-id", // Replace with your actual webhook URL
		},
	}

	// Send the request to the Jenga API
	response, err := client.SendToMobileWallet(mobileWalletReq)
	if err != nil {
		log.Fatalf("Error sending money to mobile wallet: %v", err)
	}

	// Print the response
	fmt.Printf("\nTransaction Response:\n")
	fmt.Printf("Status: %t\n", response.Status)
	fmt.Printf("Code: %d\n", response.Code)
	fmt.Printf("Message: %s\n", response.Message)
	fmt.Printf("Reference: %s\n", response.Reference)
	fmt.Printf("Transaction ID: %s\n", response.Data.TransactionID)
	fmt.Printf("Transaction Status: %s\n", response.Data.Status)

	fmt.Println("\nNote: The actual transaction status will be sent to the callback URL provided in the request.")
	fmt.Println("The callback will contain details such as ResponseCode, ReceiverMsisdn, Reference, ResponseDescription, etc.")
}
