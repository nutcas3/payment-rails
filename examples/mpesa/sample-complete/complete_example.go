package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"payment-rails/mpesa"
	"time"
)

// Configuration holds the Mpesa API credentials
type Configuration struct {
	APIKey         string
	ConsumerSecret string
	PassKey        string
	ShortCode      string
	PhoneNumber    string
}

// loadConfig loads the configuration from environment variables
func loadConfig() Configuration {
	return Configuration{
		APIKey:         getEnv("MPESA_API_KEY", ""),
		ConsumerSecret: getEnv("MPESA_CONSUMER_SECRET", ""),
		PassKey:        getEnv("MPESA_PASS_KEY", ""),
		ShortCode:      getEnv("MPESA_SHORT_CODE", "174379"), // Default sandbox shortcode
		PhoneNumber:    getEnv("MPESA_PHONE_NUMBER", "254708374149"), // Default test phone number
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	config := loadConfig()

	if config.APIKey == "" || config.ConsumerSecret == "" || config.PassKey == "" {
		log.Fatal("MPESA_API_KEY, MPESA_CONSUMER_SECRET, and MPESA_PASS_KEY environment variables are required")
	}

	fmt.Println("Initializing Mpesa SDK...")

	client, err := mpesa.NewClient(
		config.APIKey,
		config.ConsumerSecret,
		config.PassKey,
		mpesa.SANDBOX,
	)
	if err != nil {
		log.Fatalf("Failed to initialize Mpesa client: %v", err)
	}

	client.SetHttpClient(&http.Client{
		Timeout: 60 * time.Second,
	})
	fmt.Println("\n1. Getting authentication token...")
	token, err := client.GetAuthToken()
	if err != nil {
		log.Fatalf("Failed to get auth token: %v", err)
	}
	fmt.Printf("✓ Auth Token: %s\n", token)

	fmt.Println("\n2. Registering C2B URLs...")
	c2bRegisterResponse, err := client.C2BRegisterURL(
		config.ShortCode,
		"Completed",
		"https://example.com/c2b/confirmation",
		"https://example.com/c2b/validation",
	)
	if err != nil {
		log.Printf("Failed to register C2B URL: %v", err)
	} else {
		fmt.Printf("✓ C2B Register URL Response: %+v\n", c2bRegisterResponse)
	}

	fmt.Println("\n3. Simulating C2B transaction...")
	shortCode := 0
	phoneNumber := 0
	fmt.Sscanf(config.ShortCode, "%d", &shortCode)
	fmt.Sscanf(config.PhoneNumber, "%d", &phoneNumber)
	
	c2bSimulateResponse, err := client.C2BSimulate(
		shortCode,
		"CustomerPayBillOnline",
		1,
		phoneNumber,
		"Test",
	)
	if err != nil {
		log.Printf("Failed to simulate C2B: %v", err)
	} else {
		fmt.Printf("✓ C2B Simulate Response: %+v\n", c2bSimulateResponse)
	}

	fmt.Println("\n4. Initiating STK Push...")
	stkResponse, err := client.InitiateStkPush(mpesa.StkPushParams{
		BusinessShortCode: config.ShortCode,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            "1",
		PartyA:            config.PhoneNumber,
		PartyB:            config.ShortCode,
		PhoneNumber:       config.PhoneNumber,
		CallBackURL:       "https://example.com/callback",
		AccountReference:  "Test Payment",
		TransactionDesc:   "Test Payment",
	})
	if err != nil {
		log.Printf("Failed to initiate STK push: %v", err)
	} else {
		fmt.Printf("✓ STK Push Response: %+v\n", stkResponse)
		
		if stkResponse.CheckoutRequestID != "" {
			fmt.Println("\n5. Querying STK Push status...")
			time.Sleep(5 * time.Second)
			
			queryResponse, err := client.QueryStkPush(
				config.ShortCode,
				stkResponse.CheckoutRequestID,
			)
			if err != nil {
				log.Printf("Failed to query STK push: %v", err)
			} else {
				fmt.Printf("✓ STK Push Query Response: %+v\n", queryResponse)
			}
		}
	}

	// Note: The following examples are commented out as they require additional credentials
	// and are typically used in production environments

	/*
	// Example: B2C Payment
	fmt.Println("\n6. Making B2C Payment...")
	b2cResponse, err := client.B2CPayment(mpesa.B2CPaymentParams{
		InitiatorName:      "TestInitiator",
		SecurityCredential: "SecurityCredential",
		CommandID:          "BusinessPayment",
		Amount:             1,
		PartyA:             shortCode,
		PartyB:             phoneNumber,
		Remarks:            "Test B2C Payment",
		QueueTimeOutURL:    "https://example.com/b2c/timeout",
		ResultURL:          "https://example.com/b2c/result",
		Occasion:           "Test",
	})
	if err != nil {
		log.Printf("Failed to make B2C payment: %v", err)
	} else {
		fmt.Printf("✓ B2C Payment Response: %+v\n", b2cResponse)
	}

	// Example: B2B Payment
	fmt.Println("\n7. Making B2B Payment...")
	b2bResponse, err := client.B2BPayment(mpesa.B2BPaymentParams{
		Initiator:              "TestInitiator",
		SecurityCredential:     "SecurityCredential",
		CommandID:              "BusinessPayBill",
		SenderIdentifierType:   "4",
		ReceiverIdentifierType: "4",
		Amount:                 "100",
		PartyA:                 config.ShortCode,
		PartyB:                 "600001",
		AccountReference:       "Test",
		Requester:              config.PhoneNumber,
		Remarks:                "Test B2B Payment",
		QueueTimeOutURL:        "https://example.com/b2b/timeout",
		ResultURL:              "https://example.com/b2b/result",
	})
	if err != nil {
		log.Printf("Failed to make B2B payment: %v", err)
	} else {
		fmt.Printf("✓ B2B Payment Response: %+v\n", b2bResponse)
	}

	// Example: Transaction Status
	fmt.Println("\n8. Querying Transaction Status...")
	statusResponse, err := client.TransactionStatus(mpesa.TransactionStatusParams{
		Initiator:          "TestInitiator",
		SecurityCredential: "SecurityCredential",
		CommandID:          "TransactionStatusQuery",
		TransactionID:      "LKXXXX1234",
		PartyA:             shortCode,
		IdentifierType:     4,
		ResultURL:          "https://example.com/status/result",
		QueueTimeOutURL:    "https://example.com/status/timeout",
		Remarks:            "Test Transaction Status",
		Occasion:           "Test",
	})
	if err != nil {
		log.Printf("Failed to query transaction status: %v", err)
	} else {
		fmt.Printf("✓ Transaction Status Response: %+v\n", statusResponse)
	}

	// Example: Account Balance
	fmt.Println("\n9. Querying Account Balance...")
	balanceResponse, err := client.AccountBalance(mpesa.AccountBalanceParams{
		Initiator:          "TestInitiator",
		SecurityCredential: "SecurityCredential",
		CommandID:          "AccountBalance",
		PartyA:             shortCode,
		IdentifierType:     4,
		Remarks:            "Test Account Balance",
		QueueTimeOutURL:    "https://example.com/balance/timeout",
		ResultURL:          "https://example.com/balance/result",
	})
	if err != nil {
		log.Printf("Failed to query account balance: %v", err)
	} else {
		fmt.Printf("✓ Account Balance Response: %+v\n", balanceResponse)
	}

	// Example: Payment Reversal
	fmt.Println("\n10. Reversing Transaction...")
	reversalResponse, err := client.Reversal(mpesa.ReversalParams{
		Initiator:              "TestInitiator",
		SecurityCredential:     "SecurityCredential",
		CommandID:              "TransactionReversal",
		TransactionID:          "LKXXXX1234",
		Amount:                 1,
		ReceiverParty:          shortCode,
		ReceiverIdentifierType: 4,
		ResultURL:              "https://example.com/reversal/result",
		QueueTimeOutURL:        "https://example.com/reversal/timeout",
		Remarks:                "Test Reversal",
		Occasion:               "Test",
	})
	if err != nil {
		log.Printf("Failed to reverse transaction: %v", err)
	} else {
		fmt.Printf("✓ Reversal Response: %+v\n", reversalResponse)
	}
	*/

	fmt.Println("\nExample completed successfully!")
}
