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

	// Create a standing order (M-Pesa Ratiba) for recurring payments
	ratibaResponse, err := client.CreateStandingOrder(mpesa.RatibaRequest{
		StandingOrderName: "Monthly Rent Payment",
		StartDate:         "20250901",                     // Format: YYYYMMDD
		EndDate:           "20260901",                     // Format: YYYYMMDD
		BusinessShortCode: "174379",                       // Your business short code
		TransactionType:   mpesa.TransactionTypePayBill,   // For Paybill
		IdentifierType:    mpesa.ReceiverTypePaybill,      // For Paybill
		Amount:            "5000",                         // Amount in smallest currency unit
		PhoneNumber:       "254708374149",                 // Customer's phone number
		CallBackURL:       "https://example.com/callback", // Your callback URL
		AccountReference:  "Rent123",                      // Account reference
		TransactionDesc:   "Monthly Rent",                 // Transaction description
		Frequency:         mpesa.FrequencyMonthly,         // Monthly payments
	})
	if err != nil {
		log.Fatalf("Failed to create standing order: %v", err)
	}

	fmt.Printf("Standing Order Response:\n")
	fmt.Printf("Response Code: %s\n", ratibaResponse.ResponseCode)
	fmt.Printf("Response Description: %s\n", ratibaResponse.ResponseDescription)
	fmt.Printf("Response Reference ID: %s\n", ratibaResponse.ResponseRefID)
	
	fmt.Println("\nNote: The M-Pesa Ratiba API is a commercial API that requires a contract with Safaricom.")
	fmt.Println("For production use, you need to contact Safaricom at apisupport@safaricom.co.ke after testing.")
}
