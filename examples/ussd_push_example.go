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

	// Generate a unique request reference ID (in production, use a UUID generator)
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

	fmt.Printf("USSD Push Response:\n")
	fmt.Printf("Code: %s\n", response.Code)
	fmt.Printf("Status: %s\n", response.Status)
	
	fmt.Println("\nNote: If successful, the merchant will receive a USSD prompt to:")
	fmt.Println("1. Enter their operator ID")
	fmt.Println("2. Enter their operator PIN")
	fmt.Println("3. Confirm the payment")
	
	fmt.Println("\nThe final result will be sent to your callback URL with details including:")
	fmt.Println("- Result code and description")
	fmt.Println("- Transaction ID (for successful transactions)")
	fmt.Println("- Status (SUCCESS/FAILED)")
}
