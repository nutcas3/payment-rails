package main

import (
	"fmt"
	"log"
	"os"
	"payment-rails/mpesa"
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
	stkResponse, err := client.InitiateStkPush(
		"174379",                                     // Business Short Code
		"CustomerPayBillOnline",                      // Transaction Type
		"1",                                          // Amount
		"254708374149",                               // Party A (Phone number)
		"174379",                                     // Party B (Short code)
		"254708374149",                               // Phone Number
		"https://example.com/callback",               // Callback URL
		"Test Payment",                               // Account Reference
		"Test Payment",                               // Transaction Description
	)
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
	b2cResponse, err := client.B2CPayment(
		"TestInitiator",                              // Initiator Name
		"SecurityCredential",                         // Security Credential
		"BusinessPayment",                            // Command ID
		1,                                            // Amount
		600000,                                       // Party A (Short code)
		254708374149,                                 // Party B (Phone number)
		"Test B2C Payment",                           // Remarks
		"https://example.com/b2c/timeout",            // Queue Timeout URL
		"https://example.com/b2c/result",             // Result URL
		"Test",                                       // Occasion
	)
	if err != nil {
		log.Fatalf("Failed to make B2C payment: %v", err)
	}
	fmt.Printf("B2C Payment Response: %+v\n", b2cResponse)
	*/

	// Example: Transaction Status
	// Note: You need a valid transaction ID and security credential
	/*
	statusResponse, err := client.TransactionStatus(
		"TestInitiator",                              // Initiator
		"SecurityCredential",                         // Security Credential
		"TransactionStatusQuery",                     // Command ID
		"LKXXXX1234",                                 // Transaction ID
		600000,                                       // Party A (Short code)
		4,                                            // Identifier Type
		"https://example.com/status/result",          // Result URL
		"https://example.com/status/timeout",         // Queue Timeout URL
		"Test Transaction Status",                    // Remarks
		"Test",                                       // Occasion
	)
	if err != nil {
		log.Fatalf("Failed to query transaction status: %v", err)
	}
	fmt.Printf("Transaction Status Response: %+v\n", statusResponse)
	*/

	fmt.Println("Examples completed successfully")
}
