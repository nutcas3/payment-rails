package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/nutcas3/payment-rails/mpesa"
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

	fmt.Printf("Business Pay Bill Response:\n")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Description: %s\n", response.ResponseDescription)
	fmt.Printf("Originator Conversation ID: %s\n", response.OriginatorConversationID)
	fmt.Printf("Conversation ID: %s\n", response.ConversationID)
}
