package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"payment-rails/standardbank"
	"payment-rails/standardbank/pkg/api"
)

func main() {
	// Get credentials from environment variables
	clientID := os.Getenv("STANDARD_BANK_CLIENT_ID")
	clientSecret := os.Getenv("STANDARD_BANK_CLIENT_SECRET")
	apiKey := os.Getenv("STANDARD_BANK_API_KEY")

	if clientID == "" || clientSecret == "" || apiKey == "" {
		log.Fatal("Please set STANDARD_BANK_CLIENT_ID, STANDARD_BANK_CLIENT_SECRET, and STANDARD_BANK_API_KEY environment variables")
	}

	// Initialize the client
	client := standardbank.NewClient(
		clientID,
		clientSecret,
		apiKey,
		standardbank.WithEnvironment(standardbank.EnvironmentSandbox),
	)

	ctx := context.Background()

	// Example 1: Create a payment
	fmt.Println("=== Creating Payment ===")
	payment, err := client.Payments().Create(ctx, api.PaymentRequest{
		Amount:            1000.00,
		Currency:          "ZAR",
		Reference:         "INV-2024-001",
		Description:       "Payment for services",
		SourceAccount:     "ACC123456",
		DestinationAccount: "ACC789012",
		BeneficiaryName:   "John Doe",
		IdempotencyKey:    "unique-key-123",
	})

	if err != nil {
		log.Printf("Error creating payment: %v", err)
	} else {
		fmt.Printf("Payment created successfully!\n")
		fmt.Printf("  Payment ID: %s\n", payment.PaymentID)
		fmt.Printf("  Transaction ID: %s\n", payment.TransactionID)
		fmt.Printf("  Status: %s\n", payment.Status)
		fmt.Printf("  Amount: %.2f %s\n", payment.Amount, payment.Currency)
	}

	// Example 2: Get payment status
	if payment != nil {
		fmt.Println("\n=== Getting Payment Status ===")
		status, err := client.Payments().GetStatus(ctx, payment.PaymentID)
		if err != nil {
			log.Printf("Error getting payment status: %v", err)
		} else {
			fmt.Printf("Payment Status: %s\n", status.Status)
			fmt.Printf("Last Updated: %s\n", status.LastUpdated)
		}
	}

	// Example 3: Create an internal transfer
	fmt.Println("\n=== Creating Internal Transfer ===")
	transfer, err := client.Transfers().Create(ctx, api.InternalTransferRequest{
		SourceAccount:      "ACC001",
		DestinationAccount: "ACC002",
		Amount:            500.00,
		Currency:          "ZAR",
		Reference:         "TRANSFER-001",
		Description:       "Internal account transfer",
		DestinationName:   "Jane Smith",
		IdempotencyKey:    "transfer-key-123",
	})

	if err != nil {
		log.Printf("Error creating transfer: %v", err)
	} else {
		fmt.Printf("Transfer created successfully!\n")
		fmt.Printf("  Transfer ID: %s\n", transfer.TransferID)
		fmt.Printf("  Status: %s\n", transfer.Status)
		fmt.Printf("  Amount: %.2f %s\n", transfer.Amount, transfer.Currency)
	}

	// Example 4: List available providers
	fmt.Println("\n=== Listing Providers ===")
	providers, err := client.Providers().List(ctx)
	if err != nil {
		log.Printf("Error listing providers: %v", err)
	} else {
		fmt.Printf("Total Providers: %d\n", providers.Total)
		for _, provider := range providers.Providers {
			fmt.Printf("  - %s (%s): %s - %s\n", provider.Name, provider.ID, provider.Type, provider.Status)
		}
	}

	// Example 5: Execute provider payment (if providers are available)
	if providers != nil && len(providers.Providers) > 0 {
		fmt.Println("\n=== Executing Provider Payment ===")
		providerID := providers.Providers[0].ID

		providerPayment, err := client.Providers().Pay(ctx, api.ProviderPaymentRequest{
			ProviderID:  providerID,
			Amount:      200.00,
			Currency:    "ZAR",
			Reference:   "PROV-PAY-001",
			Description: "Mobile wallet payment",
			Destination: map[string]interface{}{
				"accountNumber": "1234567890",
				"name":          "John Doe",
				"phoneNumber":   "+27123456789",
			},
			IdempotencyKey: "provider-key-123",
		})

		if err != nil {
			log.Printf("Error executing provider payment: %v", err)
		} else {
			fmt.Printf("Provider payment executed successfully!\n")
			fmt.Printf("  Payment ID: %s\n", providerPayment.PaymentID)
			fmt.Printf("  Provider ID: %s\n", providerPayment.ProviderID)
			fmt.Printf("  Status: %s\n", providerPayment.Status)
			if providerPayment.ProviderReference != "" {
				fmt.Printf("  Provider Reference: %s\n", providerPayment.ProviderReference)
			}
		}
	}

	fmt.Println("\n=== Examples completed ===")
}

