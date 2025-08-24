package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"payment-rails/jenga/pkg/api"

	"payment-rails/jenga"
)

func main() {
	// Get environment variables
	apiKey := os.Getenv("JENGA_API_KEY")
	username := os.Getenv("JENGA_MERCHANT_CODE")
	password := os.Getenv("JENGA_CONSUMER_SECRET")
	privateKeyPath := os.Getenv("JENGA_PRIVATE_KEY_PATH")
	webhookSecret := os.Getenv("JENGA_WEBHOOK_SECRET")
	
	// Initialize Jenga client
	client, err := jenga.NewClient(apiKey, username, password, privateKeyPath, "sandbox")
	if err != nil {
		log.Fatalf("Error initializing Jenga client: %v", err)
	}
	
	// Set webhook secret for signature validation
	client.SetWebhookSecret(webhookSecret)
	
	// Define webhook handlers
	handlers := api.WebhookHandlers{
		TransactionSuccessHandler: handleTransactionSuccess,
		TransactionFailedHandler:  handleTransactionFailed,
		AccountUpdatedHandler:     handleAccountUpdated,
		KYCUpdatedHandler:         handleKYCUpdated,
		DefaultHandler:            handleDefaultEvent,
	}
	
	// Set up HTTP server
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if err := client.HandleWebhook(w, r, handlers); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Printf("Starting webhook server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// Handler for successful transactions
func handleTransactionSuccess(event *api.WebhookEvent) {
	fmt.Println("Received transaction success event:", event.ID)
	
	// Parse transaction data
	var data api.TransactionWebhookData
	if err := json.Unmarshal(event.Data, &data); err != nil {
		fmt.Printf("Error parsing transaction data: %v\n", err)
		return
	}
	
	fmt.Printf("Transaction ID: %s\n", data.TransactionID)
	fmt.Printf("Reference: %s\n", data.Reference)
	fmt.Printf("Amount: %s %s\n", data.Amount, data.Currency)
	fmt.Printf("Status: %s\n", data.Status)
	
	// Process the transaction (e.g., update database, send notification)
	// ...
}

// Handler for failed transactions
func handleTransactionFailed(event *api.WebhookEvent) {
	fmt.Println("Received transaction failed event:", event.ID)
	
	// Parse transaction data
	var data api.TransactionWebhookData
	if err := json.Unmarshal(event.Data, &data); err != nil {
		fmt.Printf("Error parsing transaction data: %v\n", err)
		return
	}
	
	fmt.Printf("Transaction ID: %s\n", data.TransactionID)
	fmt.Printf("Reference: %s\n", data.Reference)
	fmt.Printf("Status Reason: %s\n", data.StatusReason)
	
	// Handle the failed transaction (e.g., retry, notify user)
	// ...
}

// Handler for account updates
func handleAccountUpdated(event *api.WebhookEvent) {
	fmt.Println("Received account updated event:", event.ID)
	
	// Parse account data
	var data api.AccountWebhookData
	if err := json.Unmarshal(event.Data, &data); err != nil {
		fmt.Printf("Error parsing account data: %v\n", err)
		return
	}
	
	fmt.Printf("Account ID: %s\n", data.AccountID)
	fmt.Printf("Account Name: %s\n", data.AccountName)
	fmt.Printf("Status: %s\n", data.Status)
	
	// Process the account update (e.g., update database)
	// ...
}

// Handler for KYC updates
func handleKYCUpdated(event *api.WebhookEvent) {
	fmt.Println("Received KYC updated event:", event.ID)
	
	// Parse KYC data
	var data api.KYCWebhookData
	if err := json.Unmarshal(event.Data, &data); err != nil {
		fmt.Printf("Error parsing KYC data: %v\n", err)
		return
	}
	
	fmt.Printf("Customer ID: %s\n", data.CustomerID)
	fmt.Printf("Status: %s\n", data.Status)
	fmt.Printf("Verification Type: %s\n", data.VerificationType)
	
	// Process the KYC update (e.g., update customer status)
	// ...
}

// Default handler for other event types
func handleDefaultEvent(event *api.WebhookEvent) {
	fmt.Printf("Received unhandled event type: %s, ID: %s\n", event.EventType, event.ID)
}
