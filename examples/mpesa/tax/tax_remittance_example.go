package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"payment-rails/mpesa"
	"time"
)

func main() {
	apiKey := os.Getenv("MPESA_API_KEY")
	consumerSecret := os.Getenv("MPESA_CONSUMER_SECRET")
	passKey := os.Getenv("MPESA_PASS_KEY")

	if apiKey == "" || consumerSecret == "" || passKey == "" {
		log.Fatal("Please set MPESA_API_KEY, MPESA_CONSUMER_SECRET, and MPESA_PASS_KEY environment variables")
	}

	client, err := mpesa.NewClient(
		apiKey,
		consumerSecret,
		passKey,
		mpesa.SANDBOX, // Use SANDBOX for testing, PRODUCTION for live environment
	)
	if err != nil {
		log.Fatalf("Failed to initialize Mpesa client: %v", err)
	}

	client.SetHttpClient(&http.Client{
		Timeout: 30 * time.Second,
	})

	// Get security credential
	// In production, you would encrypt your API password using the Safaricom public key
	// For sandbox testing, you can use a placeholder value
	securityCredential := "Safaricom123!"

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

	fmt.Printf("Tax Remittance Response:\n")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Description: %s\n", response.ResponseDescription)
	fmt.Printf("Originator Conversation ID: %s\n", response.OriginatorConversationID)
	fmt.Printf("Conversation ID: %s\n", response.ConversationID)
	
	fmt.Println("\nNote: The final result will be sent to your callback URL with details including:")
	fmt.Println("- Result code and description")
	fmt.Println("- Transaction ID")
	fmt.Println("- Account balances and other transaction details")
}
