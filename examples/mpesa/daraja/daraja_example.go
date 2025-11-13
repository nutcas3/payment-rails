package main

import (
	"fmt"
	"log"
	"os"
	"github.com/nutcas3/payment-rails/mpesa"
)

func main() {
	// Get API credentials from environment variables
	apiKey := os.Getenv("MPESA_API_KEY")
	consumerSecret := os.Getenv("MPESA_CONSUMER_SECRET")
	passKey := os.Getenv("MPESA_PASS_KEY")

	if apiKey == "" || consumerSecret == "" || passKey == "" {
		log.Fatal("MPESA_API_KEY, MPESA_CONSUMER_SECRET, and MPESA_PASS_KEY environment variables are required")
	}

	// Initialize the Mpesa client
	client, err := mpesa.NewClient(apiKey, consumerSecret, passKey, mpesa.SANDBOX)
	if err != nil {
		log.Fatalf("Failed to initialize Mpesa client: %v", err)
	}

	// Example: Get authentication token
	token, err := client.GetAuthToken()
	if err != nil {
		log.Fatalf("Failed to get auth token: %v", err)
	}
	fmt.Printf("Auth Token: %s\n", token)

	// Example: Initiate STK Push
	stkResponse, err := client.InitiateStkPush(mpesa.StkPushParams{
		BusinessShortCode: "174379",
		TransactionType:   "CustomerPayBillOnline",
		Amount:            "1",
		PartyA:            "254708374149",
		PartyB:            "174379",
		PhoneNumber:       "254708374149",
		CallBackURL:       "https://example.com/callback",
		AccountReference:  "Test Payment",
		TransactionDesc:   "Test Payment",
	})
	if err != nil {
		log.Fatalf("Failed to initiate STK push: %v", err)
	}
	fmt.Printf("STK Push Response: %+v\n", stkResponse)

	// Example: Query STK Push status
	// Note: You would need a valid CheckoutRequestID from a previous STK Push
	/*
	queryResponse, err := client.QueryStkPush(
		"174379",                         // Business Short Code
		"ws_CO_DMZ_12345678901234567",   // Checkout Request ID from previous STK Push
	)
	if err != nil {
		log.Fatalf("Failed to query STK push: %v", err)
	}
	fmt.Printf("STK Push Query Response: %+v\n", queryResponse)
	*/

	// Example: Register C2B URL
	c2bRegisterResponse, err := client.C2BRegisterURL(
		"600000",                                     // Short Code
		"Completed",                                  // Response Type
		"https://example.com/c2b/confirmation",       // Confirmation URL
		"https://example.com/c2b/validation",         // Validation URL
	)
	if err != nil {
		log.Fatalf("Failed to register C2B URL: %v", err)
	}
	fmt.Printf("C2B Register URL Response: %+v\n", c2bRegisterResponse)

	// Example: C2B Simulate (only works in sandbox)
	c2bSimulateResponse, err := client.C2BSimulate(
		600000,                // Short Code
		"CustomerPayBillOnline", // Command ID
		1,                     // Amount
		254708374149,          // MSISDN (Phone number)
		"Test",                // Bill Reference Number
	)
	if err != nil {
		log.Fatalf("Failed to simulate C2B: %v", err)
	}
	fmt.Printf("C2B Simulate Response: %+v\n", c2bSimulateResponse)

	// Example: B2C Payment
	// Note: You need a valid security credential
	/*
	b2cResponse, err := client.B2CPayment(mpesa.B2CPaymentParams{
		InitiatorName:      "TestInitiator",
		SecurityCredential: "SecurityCredential",
		CommandID:          "BusinessPayment",
		Amount:             1,
		PartyA:             600000,
		PartyB:             254708374149,
		Remarks:            "Test B2C Payment",
		QueueTimeOutURL:    "https://example.com/b2c/timeout",
		ResultURL:          "https://example.com/b2c/result",
		Occasion:           "Test",
	})
	if err != nil {
		log.Fatalf("Failed to make B2C payment: %v", err)
	}
	fmt.Printf("B2C Payment Response: %+v\n", b2cResponse)
	*/

	// Example: Transaction Status
	// Note: You need a valid transaction ID and security credential
	/*
	statusResponse, err := client.TransactionStatus(mpesa.TransactionStatusParams{
		Initiator:          "TestInitiator",
		SecurityCredential: "SecurityCredential",
		CommandID:          "TransactionStatusQuery",
		TransactionID:      "LKXXXX1234",
		PartyA:             600000,
		IdentifierType:     4,
		ResultURL:          "https://example.com/status/result",
		QueueTimeOutURL:    "https://example.com/status/timeout",
		Remarks:            "Test Transaction Status",
		Occasion:           "Test",
	})
	if err != nil {
		log.Fatalf("Failed to query transaction status: %v", err)
	}
	fmt.Printf("Transaction Status Response: %+v\n", statusResponse)
	*/

	fmt.Println("Examples completed successfully")
}
