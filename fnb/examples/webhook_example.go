package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nutcas3/payment-rails/fnb/pkg/api"
)

// This example demonstrates how to set up a webhook endpoint to receive
// real-time notifications from FNB about payment and mandate events

func main() {
	// Get webhook secret from environment
	webhookSecret := os.Getenv("FNB_WEBHOOK_SECRET")
	if webhookSecret == "" {
		log.Fatal("Please set FNB_WEBHOOK_SECRET environment variable")
	}

	// Create webhook handler
	webhookHandler := api.NewWebhookHandler(webhookSecret)

	// Register handlers for different event types
	registerEventHandlers(webhookHandler)

	// Set up HTTP server
	http.HandleFunc("/webhooks/fnb", webhookHandler.HandleWebhook)
	http.HandleFunc("/health", healthCheckHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Webhook server starting on port %s\n", port)
	fmt.Printf("📡 Listening for FNB webhooks at /webhooks/fnb\n")
	fmt.Printf("💚 Health check available at /health\n\n")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func registerEventHandlers(handler *api.WebhookHandler) {
	// Payment completed event
	handler.RegisterHandler(api.EventPaymentCompleted, func(event api.WebhookEvent) error {
		fmt.Printf("✓ Payment Completed: %s\n", event.ResourceID)

		// Extract payment data
		amount, _ := event.Data["amount"].(float64)
		currency, _ := event.Data["currency"].(string)
		reference, _ := event.Data["reference"].(string)
		beneficiary, _ := event.Data["beneficiaryName"].(string)

		fmt.Printf("  Amount: %.2f %s\n", amount, currency)
		fmt.Printf("  Reference: %s\n", reference)
		fmt.Printf("  Beneficiary: %s\n", beneficiary)

		// Update your database
		if err := updatePaymentStatus(event.ResourceID, "COMPLETED", event.Data); err != nil {
			log.Printf("Failed to update payment status: %v", err)
			return err
		}

		// Send notification to customer
		if err := sendCustomerNotification(reference, "completed"); err != nil {
			log.Printf("Failed to send notification: %v", err)
			// Don't return error - notification failure shouldn't fail webhook
		}

		return nil
	})

	// Payment failed event
	handler.RegisterHandler(api.EventPaymentFailed, func(event api.WebhookEvent) error {
		fmt.Printf("✗ Payment Failed: %s\n", event.ResourceID)

		failureReason, _ := event.Data["failureReason"].(string)
		reference, _ := event.Data["reference"].(string)

		fmt.Printf("  Reference: %s\n", reference)
		fmt.Printf("  Reason: %s\n", failureReason)

		// Update database
		if err := updatePaymentStatus(event.ResourceID, "FAILED", event.Data); err != nil {
			return err
		}

		// Notify customer and support team
		sendCustomerNotification(reference, "failed")
		alertSupportTeam(event.ResourceID, failureReason)

		return nil
	})

	// Collection completed event
	handler.RegisterHandler(api.EventCollectionCompleted, func(event api.WebhookEvent) error {
		fmt.Printf("✓ Collection Completed: %s\n", event.ResourceID)

		amount, _ := event.Data["amount"].(float64)
		reference, _ := event.Data["reference"].(string)

		fmt.Printf("  Amount: %.2f ZAR\n", amount)
		fmt.Printf("  Reference: %s\n", reference)

		// Update database
		if err := updateCollectionStatus(event.ResourceID, "COMPLETED", event.Data); err != nil {
			return err
		}

		// Process successful collection (e.g., activate subscription)
		processSuccessfulCollection(reference, amount)

		return nil
	})

	// Collection failed event
	handler.RegisterHandler(api.EventCollectionFailed, func(event api.WebhookEvent) error {
		fmt.Printf("✗ Collection Failed: %s\n", event.ResourceID)

		failureReason, _ := event.Data["failureReason"].(string)
		reference, _ := event.Data["reference"].(string)

		fmt.Printf("  Reference: %s\n", reference)
		fmt.Printf("  Reason: %s\n", failureReason)

		// Update database
		if err := updateCollectionStatus(event.ResourceID, "FAILED", event.Data); err != nil {
			return err
		}

		// Handle failed collection (e.g., retry logic, suspend service)
		handleFailedCollection(reference, failureReason)

		return nil
	})

	// Mandate approved event
	handler.RegisterHandler(api.EventMandateApproved, func(event api.WebhookEvent) error {
		fmt.Printf("✓ Mandate Approved: %s\n", event.ResourceID)

		contractRef, _ := event.Data["contractReference"].(string)
		debtorName, _ := event.Data["debtorName"].(string)

		fmt.Printf("  Contract: %s\n", contractRef)
		fmt.Printf("  Debtor: %s\n", debtorName)

		// Update database
		if err := updateMandateStatus(event.ResourceID, "ACTIVE", event.Data); err != nil {
			return err
		}

		// Activate subscription or service
		activateSubscription(contractRef)

		// Send welcome notification
		sendCustomerNotification(contractRef, "mandate_approved")

		return nil
	})

	// Mandate rejected event
	handler.RegisterHandler(api.EventMandateRejected, func(event api.WebhookEvent) error {
		fmt.Printf("✗ Mandate Rejected: %s\n", event.ResourceID)

		contractRef, _ := event.Data["contractReference"].(string)
		rejectionReason, _ := event.Data["rejectionReason"].(string)

		fmt.Printf("  Contract: %s\n", contractRef)
		fmt.Printf("  Reason: %s\n", rejectionReason)

		// Update database
		if err := updateMandateStatus(event.ResourceID, "REJECTED", event.Data); err != nil {
			return err
		}

		// Handle rejection (e.g., offer alternative payment method)
		handleMandateRejection(contractRef, rejectionReason)

		return nil
	})

	// Dispute submitted event
	handler.RegisterHandler(api.EventDisputeSubmitted, func(event api.WebhookEvent) error {
		fmt.Printf("⚠️  Dispute Submitted: %s\n", event.ResourceID)

		transactionID, _ := event.Data["transactionId"].(string)
		reason, _ := event.Data["reason"].(string)

		fmt.Printf("  Transaction: %s\n", transactionID)
		fmt.Printf("  Reason: %s\n", reason)

		// Update database
		if err := logDispute(event.ResourceID, transactionID, reason); err != nil {
			return err
		}

		// Alert support team
		alertSupportTeam(transactionID, fmt.Sprintf("Dispute: %s", reason))

		return nil
	})
}

// Mock database functions - replace with your actual database logic

func updatePaymentStatus(transactionID, status string, data map[string]interface{}) error {
	fmt.Printf("  → Updating payment %s to status: %s\n", transactionID, status)
	// TODO: Update your database
	// db.Exec("UPDATE payments SET status = ?, updated_at = NOW() WHERE transaction_id = ?", status, transactionID)
	return nil
}

func updateCollectionStatus(transactionID, status string, data map[string]interface{}) error {
	fmt.Printf("  → Updating collection %s to status: %s\n", transactionID, status)
	// TODO: Update your database
	return nil
}

func updateMandateStatus(mandateID, status string, data map[string]interface{}) error {
	fmt.Printf("  → Updating mandate %s to status: %s\n", mandateID, status)
	// TODO: Update your database
	return nil
}

func logDispute(disputeID, transactionID, reason string) error {
	fmt.Printf("  → Logging dispute %s for transaction %s\n", disputeID, transactionID)
	// TODO: Insert into disputes table
	return nil
}

func sendCustomerNotification(reference, eventType string) error {
	fmt.Printf("  → Sending notification to customer (ref: %s, type: %s)\n", reference, eventType)
	// TODO: Send email/SMS notification
	return nil
}

func alertSupportTeam(transactionID, message string) {
	fmt.Printf("  → Alerting support team: %s - %s\n", transactionID, message)
	// TODO: Send alert to support team (Slack, email, etc.)
}

func processSuccessfulCollection(reference string, amount float64) {
	fmt.Printf("  → Processing successful collection: %s (%.2f ZAR)\n", reference, amount)
	// TODO: Activate service, extend subscription, etc.
}

func handleFailedCollection(reference, reason string) {
	fmt.Printf("  → Handling failed collection: %s - %s\n", reference, reason)
	// TODO: Implement retry logic, suspend service, etc.
}

func activateSubscription(contractRef string) {
	fmt.Printf("  → Activating subscription: %s\n", contractRef)
	// TODO: Activate subscription in your system
}

func handleMandateRejection(contractRef, reason string) {
	fmt.Printf("  → Handling mandate rejection: %s - %s\n", contractRef, reason)
	// TODO: Offer alternative payment methods
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "healthy",
		"service": "fnb-webhook-handler",
	})
}

// Example: Database connection (replace with your actual DB setup)
var db *sql.DB

func initDB() {
	// TODO: Initialize your database connection
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
}
