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
	stkResponse, err := client.InitiateStkPush(
		config.ShortCode,
		"CustomerPayBillOnline",
		"1",
		config.PhoneNumber,
		config.ShortCode,
		config.PhoneNumber,
		"https://example.com/callback",
		"Test Payment",
		"Test Payment",
	)
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
	b2cResponse, err := client.B2CPayment(
		"TestInitiator",                              // Initiator Name
		"SecurityCredential",                         // Security Credential
		"BusinessPayment",                            // Command ID
		1,                                            // Amount
		shortCode,                                    // Party A (Short code)
		phoneNumber,                                  // Party B (Phone number)
		"Test B2C Payment",                           // Remarks
		"https://example.com/b2c/timeout",            // Queue Timeout URL
		"https://example.com/b2c/result",             // Result URL
		"Test",                                       // Occasion
	)
	if err != nil {
		log.Printf("Failed to make B2C payment: %v", err)
	} else {
		fmt.Printf("✓ B2C Payment Response: %+v\n", b2cResponse)
	}

	// Example: B2B Payment
	fmt.Println("\n7. Making B2B Payment...")
	b2bResponse, err := client.B2BPayment(
		"TestInitiator",                              // Initiator
		"SecurityCredential",                         // Security Credential
		"BusinessPayBill",                            // Command ID
		"4",                                          // Sender Identifier Type
		"4",                                          // Receiver Identifier Type
		"100",                                        // Amount
		config.ShortCode,                             // Party A (Short code)
		"600001",                                     // Party B (Short code)
		"Test",                                       // Account Reference
		config.PhoneNumber,                           // Requester
		"Test B2B Payment",                           // Remarks
		"https://example.com/b2b/timeout",            // Queue Timeout URL
		"https://example.com/b2b/result",             // Result URL
	)
	if err != nil {
		log.Printf("Failed to make B2B payment: %v", err)
	} else {
		fmt.Printf("✓ B2B Payment Response: %+v\n", b2bResponse)
	}

	// Example: Transaction Status
	fmt.Println("\n8. Querying Transaction Status...")
	statusResponse, err := client.TransactionStatus(
		"TestInitiator",                              // Initiator
		"SecurityCredential",                         // Security Credential
		"TransactionStatusQuery",                     // Command ID
		"LKXXXX1234",                                 // Transaction ID
		shortCode,                                    // Party A (Short code)
		4,                                            // Identifier Type
		"https://example.com/status/result",          // Result URL
		"https://example.com/status/timeout",         // Queue Timeout URL
		"Test Transaction Status",                    // Remarks
		"Test",                                       // Occasion
	)
	if err != nil {
		log.Printf("Failed to query transaction status: %v", err)
	} else {
		fmt.Printf("✓ Transaction Status Response: %+v\n", statusResponse)
	}

	// Example: Account Balance
	fmt.Println("\n9. Querying Account Balance...")
	balanceResponse, err := client.AccountBalance(
		"TestInitiator",                              // Initiator
		"SecurityCredential",                         // Security Credential
		"AccountBalance",                             // Command ID
		shortCode,                                    // Party A (Short code)
		4,                                            // Identifier Type
		"Test Account Balance",                       // Remarks
		"https://example.com/balance/timeout",        // Queue Timeout URL
		"https://example.com/balance/result",         // Result URL
	)
	if err != nil {
		log.Printf("Failed to query account balance: %v", err)
	} else {
		fmt.Printf("✓ Account Balance Response: %+v\n", balanceResponse)
	}

	// Example: Payment Reversal
	fmt.Println("\n10. Reversing Transaction...")
	reversalResponse, err := client.Reversal(
		"TestInitiator",                              // Initiator
		"SecurityCredential",                         // Security Credential
		"TransactionReversal",                        // Command ID
		"LKXXXX1234",                                 // Transaction ID
		1,                                            // Amount
		shortCode,                                    // Receiver Party
		4,                                            // Receiver Identifier Type
		"https://example.com/reversal/result",        // Result URL
		"https://example.com/reversal/timeout",       // Queue Timeout URL
		"Test Reversal",                              // Remarks
		"Test",                                       // Occasion
	)
	if err != nil {
		log.Printf("Failed to reverse transaction: %v", err)
	} else {
		fmt.Printf("✓ Reversal Response: %+v\n", reversalResponse)
	}
	*/

	fmt.Println("\nExample completed successfully!")
}
