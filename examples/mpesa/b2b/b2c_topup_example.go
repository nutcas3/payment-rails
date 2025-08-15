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

	fmt.Printf("B2C Account Top Up Response:\n")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Description: %s\n", response.ResponseDescription)
	fmt.Printf("Originator Conversation ID: %s\n", response.OriginatorConversationID)
	fmt.Printf("Conversation ID: %s\n", response.ConversationID)
}
