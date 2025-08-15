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

	// Create bulk invoices
	invoices := []mpesa.BillManagerSingleInvoiceRequest{
		{
			ExternalReference: "#9932341",                // Unique invoice reference in your system
			BilledFullName:    "John Doe",                // Name of the recipient
			BilledPhoneNumber: "0722000000",              // Safaricom phone number to receive invoice
			BilledPeriod:      "August 2025",             // Month and Year
			InvoiceName:       "Monthly Subscription",    // Descriptive name for what customer is being billed
			DueDate:           "2025-09-15",              // Date customer is expected to pay
			AccountReference:  "1ASD678H",                // Account number that uniquely identifies a customer
			Amount:            "800",                     // Total invoice amount in KES
			InvoiceItems: []mpesa.BillManagerInvoiceItem{ // Optional: Additional billable items
				{
					ItemName: "Subscription",
					Amount:   "700",
				},
				{
					ItemName: "Processing Fee",
					Amount:   "100",
				},
			},
		},
		{
			ExternalReference: "#9932342",                // Unique invoice reference in your system
			BilledFullName:    "Jane Smith",              // Name of the recipient
			BilledPhoneNumber: "0722000001",              // Safaricom phone number to receive invoice
			BilledPeriod:      "August 2025",             // Month and Year
			InvoiceName:       "Monthly Subscription",    // Descriptive name for what customer is being billed
			DueDate:           "2025-09-15",              // Date customer is expected to pay
			AccountReference:  "2BSD679J",                // Account number that uniquely identifies a customer
			Amount:            "1200",                    // Total invoice amount in KES
			InvoiceItems: []mpesa.BillManagerInvoiceItem{ // Optional: Additional billable items
				{
					ItemName: "Premium Subscription",
					Amount:   "1000",
				},
				{
					ItemName: "Processing Fee",
					Amount:   "200",
				},
			},
		},
	}

	response, err := client.CreateBulkInvoices(invoices)
	if err != nil {
		log.Fatalf("Failed to create bulk invoices: %v", err)
	}

	fmt.Printf("Bulk Invoice Creation Response:\n")
	fmt.Printf("Status Message: %s\n", response.StatusMessage)
	fmt.Printf("Response Message: %s\n", response.ResMsg)
	fmt.Printf("Response Code: %s\n", response.ResCode)

	if response.ResCode == "200" {
		fmt.Println("\nBulk invoices created successfully!")
		fmt.Println("All customers will receive SMS notifications with their invoice details.")
		fmt.Println("They can pay via:")
		fmt.Println("- USSD")
		fmt.Println("- SIM Toolkit")
		fmt.Println("- M-PESA App")
		fmt.Println("- Safaricom App")
		fmt.Println("using your paybill number and the account reference provided.")
	} else {
		fmt.Println("\nFailed to create bulk invoices. Please check your request parameters.")
	}
}
