package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nutcas3/payment-rails/fnb"
	"github.com/nutcas3/payment-rails/fnb/pkg/api"
)

func main() {
	// Get credentials from environment variables
	clientID := os.Getenv("FNB_CLIENT_ID")
	clientSecret := os.Getenv("FNB_CLIENT_SECRET")
	apiKey := os.Getenv("FNB_API_KEY")

	if clientID == "" || clientSecret == "" || apiKey == "" {
		log.Fatal("Please set FNB_CLIENT_ID, FNB_CLIENT_SECRET, and FNB_API_KEY environment variables")
	}

	// Create FNB client
	client := fnb.NewClient(
		clientID,
		clientSecret,
		apiKey,
		fnb.WithEnvironment(fnb.EnvironmentSandbox),
	)

	ctx := context.Background()

	// Example 1: Send a single EFT payment
	fmt.Println("=== Example 1: Single EFT Payment ===")
	singlePaymentExample(ctx, client)

	// Example 2: Check payment status
	fmt.Println("\n=== Example 2: Check Payment Status ===")
	paymentStatusExample(ctx, client)

	// Example 3: Send batch payments
	fmt.Println("\n=== Example 3: Batch Payments ===")
	batchPaymentExample(ctx, client)

	// Example 4: Cancel a payment
	fmt.Println("\n=== Example 4: Cancel Payment ===")
	cancelPaymentExample(ctx, client)
}

func singlePaymentExample(ctx context.Context, client *fnb.Client) {
	req := api.EFTPaymentRequest{
		SourceAccountNumber:      "1234567890",
		SourceAccountType:        "CURRENT",
		BeneficiaryAccountNumber: "9876543210",
		BeneficiaryName:          "John Doe",
		BeneficiaryBankCode:      "250655", // FNB universal branch code
		BeneficiaryReference:     "Payment from ABC Company",
		Amount:                   1500.00,
		Currency:                 "ZAR",
		PaymentReference:         "INV-2025-001",
		PaymentDescription:       "Invoice payment for services rendered",
		NotificationEmail:        "john.doe@example.com",
		IdempotencyKey:           "payment-001-20251002",
	}

	resp, err := client.CreateEFTPayment(ctx, req)
	if err != nil {
		log.Printf("Payment failed: %v", err)
		return
	}

	fmt.Printf("✓ Payment submitted successfully!\n")
	fmt.Printf("  Transaction ID: %s\n", resp.TransactionID)
	fmt.Printf("  Status: %s\n", resp.Status)
	fmt.Printf("  Reference: %s\n", resp.PaymentReference)
	fmt.Printf("  Amount: %.2f %s\n", resp.Amount, resp.Currency)
	fmt.Printf("  Beneficiary: %s\n", resp.BeneficiaryName)
	fmt.Printf("  Processing Date: %s\n", resp.ProcessingDate.Format("2006-01-02 15:04:05"))
}

func paymentStatusExample(ctx context.Context, client *fnb.Client) {
	// Replace with actual transaction ID
	transactionID := "TXN123456789"

	status, err := client.GetPaymentStatus(ctx, transactionID)
	if err != nil {
		log.Printf("Failed to get payment status: %v", err)
		return
	}

	fmt.Printf("Transaction Status:\n")
	fmt.Printf("  Transaction ID: %s\n", status.TransactionID)
	fmt.Printf("  Reference: %s\n", status.PaymentReference)
	fmt.Printf("  Status: %s - %s\n", status.Status, status.StatusDescription)
	fmt.Printf("  Amount: %.2f %s\n", status.Amount, status.Currency)
	fmt.Printf("  Source Account: %s\n", status.SourceAccount)
	fmt.Printf("  Beneficiary: %s (%s)\n", status.BeneficiaryName, status.BeneficiaryAccount)
	fmt.Printf("  Processing Date: %s\n", status.ProcessingDate.Format("2006-01-02 15:04:05"))

	if status.Status == "COMPLETED" {
		fmt.Printf("  ✓ Settlement Date: %s\n", status.SettlementDate)
	} else if status.Status == "FAILED" {
		fmt.Printf("  ✗ Failure Reason: %s\n", status.FailureReason)
	}
}

func batchPaymentExample(ctx context.Context, client *fnb.Client) {
	req := api.BatchPaymentRequest{
		BatchReference:      "BATCH-2025-001",
		SourceAccountNumber: "1234567890",
		TotalAmount:         7500.00,
		TotalCount:          3,
		ProcessingDate:      "2025-10-03",
		Payments: []api.EFTPaymentRequest{
			{
				BeneficiaryAccountNumber: "1111111111",
				BeneficiaryName:          "Supplier A Ltd",
				BeneficiaryBankCode:      "250655",
				Amount:                   2500.00,
				Currency:                 "ZAR",
				PaymentReference:         "INV-001",
				PaymentDescription:       "Invoice 001 - Office supplies",
			},
			{
				BeneficiaryAccountNumber: "2222222222",
				BeneficiaryName:          "Supplier B (Pty) Ltd",
				BeneficiaryBankCode:      "198765", // Different bank
				Amount:                   3000.00,
				Currency:                 "ZAR",
				PaymentReference:         "INV-002",
				PaymentDescription:       "Invoice 002 - IT services",
			},
			{
				BeneficiaryAccountNumber: "3333333333",
				BeneficiaryName:          "Contractor C",
				BeneficiaryBankCode:      "250655",
				Amount:                   2000.00,
				Currency:                 "ZAR",
				PaymentReference:         "INV-003",
				PaymentDescription:       "Invoice 003 - Consulting fees",
			},
		},
	}

	resp, err := client.CreateBatchPayment(ctx, req)
	if err != nil {
		log.Printf("Batch payment failed: %v", err)
		return
	}

	fmt.Printf("✓ Batch payment submitted successfully!\n")
	fmt.Printf("  Batch ID: %s\n", resp.BatchID)
	fmt.Printf("  Batch Reference: %s\n", resp.BatchReference)
	fmt.Printf("  Status: %s\n", resp.Status)
	fmt.Printf("  Total Amount: %.2f ZAR\n", resp.TotalAmount)
	fmt.Printf("  Total Payments: %d\n", resp.TotalCount)
	fmt.Printf("  Processing Date: %s\n", resp.ProcessingDate.Format("2006-01-02"))

	if len(resp.PaymentResults) > 0 {
		fmt.Printf("\n  Individual Payment Results:\n")
		for i, result := range resp.PaymentResults {
			fmt.Printf("    %d. %s - %s (%.2f ZAR)\n",
				i+1,
				result.PaymentReference,
				result.Status,
				result.Amount,
			)
		}
	}
}

func cancelPaymentExample(ctx context.Context, client *fnb.Client) {
	// Replace with actual transaction ID of a pending payment
	transactionID := "TXN123456789"
	reason := "Customer requested cancellation"

	err := client.CancelPayment(ctx, transactionID, reason)
	if err != nil {
		log.Printf("Failed to cancel payment: %v", err)
		return
	}

	fmt.Printf("✓ Payment cancelled successfully\n")
	fmt.Printf("  Transaction ID: %s\n", transactionID)
	fmt.Printf("  Reason: %s\n", reason)
}
