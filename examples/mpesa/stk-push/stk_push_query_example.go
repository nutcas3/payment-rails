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

	// Check the status of a previous STK Push transaction
	response, err := client.QueryStkPushStatus(mpesa.STKPushQueryRequest{
		BusinessShortCode: "174379",                      // Your business shortcode
		CheckoutRequestID: "ws_CO_260520211133524545",    // The CheckoutRequestID from the STK Push response
	})
	if err != nil {
		log.Fatalf("Failed to query STK Push status: %v", err)
	}

	fmt.Printf("STK Push Query Response:\n")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Description: %s\n", response.ResponseDescription)
	fmt.Printf("Merchant Request ID: %s\n", response.MerchantRequestID)
	fmt.Printf("Checkout Request ID: %s\n", response.CheckoutRequestID)
	fmt.Printf("Result Code: %s\n", response.ResultCode)
	fmt.Printf("Result Description: %s\n", response.ResultDesc)
	
	// Interpret the response
	if response.ResultCode == "0" {
		fmt.Println("\nTransaction was successful!")
	} else if response.ResultCode == "1032" {
		fmt.Println("\nTransaction was cancelled by the user")
	} else {
		fmt.Println("\nTransaction failed or is still being processed")
	}
}
