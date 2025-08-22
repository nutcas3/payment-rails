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

	// Example 1: Cancel a single invoice
	singleResponse, err := client.CancelSingleInvoice(mpesa.BillManagerCancelInvoiceRequest{
		ExternalReference: "#9932340", // External reference of the invoice to cancel
	})
	if err != nil {
		log.Fatalf("Failed to cancel single invoice: %v", err)
	}

	fmt.Printf("Single Invoice Cancellation Response:\n")
	fmt.Printf("Status Message: %s\n", singleResponse.StatusMessage)
	fmt.Printf("Response Message: %s\n", singleResponse.ResMsg)
	fmt.Printf("Response Code: %s\n", singleResponse.ResCode)
	if len(singleResponse.Errors) > 0 {
		fmt.Println("Errors:")
		for _, err := range singleResponse.Errors {
			fmt.Printf("- %s\n", err)
		}
	}

	if singleResponse.ResCode == "200" {
		fmt.Println("\nInvoice cancelled successfully!")
		fmt.Println("The customer will no longer be able to pay this invoice.")
	} else {
		fmt.Println("\nFailed to cancel invoice. Please check your request parameters.")
	}

	fmt.Println("\n------------------------------------")

	// Example 2: Cancel multiple invoices in bulk
	bulkRequests := []mpesa.BillManagerCancelInvoiceRequest{
		{
			ExternalReference: "#9932341", // First invoice to cancel
		},
		{
			ExternalReference: "#9932342", // Second invoice to cancel
		},
	}

	bulkResponse, err := client.CancelBulkInvoices(bulkRequests)
	if err != nil {
		log.Fatalf("Failed to cancel bulk invoices: %v", err)
	}

	fmt.Printf("Bulk Invoice Cancellation Response:\n")
	fmt.Printf("Status Message: %s\n", bulkResponse.StatusMessage)
	fmt.Printf("Response Message: %s\n", bulkResponse.ResMsg)
	fmt.Printf("Response Code: %s\n", bulkResponse.ResCode)
	if len(bulkResponse.Errors) > 0 {
		fmt.Println("Errors:")
		for _, err := range bulkResponse.Errors {
			fmt.Printf("- %s\n", err)
		}
	}

	if bulkResponse.ResCode == "200" {
		fmt.Println("\nBulk invoices cancelled successfully!")
		fmt.Println("All specified invoices have been cancelled.")
	} else {
		fmt.Println("\nFailed to cancel bulk invoices. Please check your request parameters.")
		fmt.Println("Note: Even if some invoices failed to cancel, others may have been successful.")
		fmt.Println("Check the errors array for details on specific failures.")
	}
}
