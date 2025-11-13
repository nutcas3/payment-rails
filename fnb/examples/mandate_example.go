package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	// Example 1: Create a DebiCheck mandate
	fmt.Println("=== Example 1: Create DebiCheck Mandate ===")
	mandateID := createMandateExample(ctx, client)

	// Example 2: Check mandate status
	fmt.Println("\n=== Example 2: Check Mandate Status ===")
	checkMandateStatusExample(ctx, client, mandateID)

	// Example 3: Verify mandate before collection
	fmt.Println("\n=== Example 3: Verify Mandate ===")
	verifyMandateExample(ctx, client, mandateID)

	// Example 4: Collect against mandate
	fmt.Println("\n=== Example 4: Collect Against Mandate ===")
	collectAgainstMandateExample(ctx, client, mandateID)

	// Example 5: List mandates
	fmt.Println("\n=== Example 5: List Mandates ===")
	listMandatesExample(ctx, client)

	// Example 6: Suspend and reinstate mandate
	fmt.Println("\n=== Example 6: Suspend and Reinstate Mandate ===")
	suspendReinstateExample(ctx, client, mandateID)

	// Example 7: Cancel mandate
	fmt.Println("\n=== Example 7: Cancel Mandate ===")
	cancelMandateExample(ctx, client, mandateID)
}

func createMandateExample(ctx context.Context, client *fnb.Client) string {
	// Calculate first collection date (next month, 1st day)
	now := time.Now()
	firstCollection := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())

	req := api.MandateRequest{
		// Creditor (your company) details
		CreditorName:          "ABC Subscription Services (Pty) Ltd",
		CreditorAbbreviation:  "ABCSUB",
		CreditorAccountNumber: "1234567890",

		// Debtor (customer) details
		DebtorName:          "Jane Smith",
		DebtorIDNumber:      "8505155009087",
		DebtorAccountNumber: "9876543210",
		DebtorAccountType:   "CURRENT",
		DebtorBankCode:      "250655",
		DebtorEmail:         "jane.smith@example.com",
		DebtorMobile:        "+27821234567",

		// Mandate details
		ContractReference:   "SUB-2025-12345",
		MaximumAmount:       500.00,
		Currency:            "ZAR",
		FrequencyType:       "MONTHLY",
		FirstCollectionDate: firstCollection.Format("2006-01-02"),
		CollectionDay:       1, // 1st of each month

		MandateDescription: "Monthly subscription for premium services",
		CategoryCode:       "SUBSCRIPTION",
		IdempotencyKey:     "mandate-12345-20251002",
	}

	resp, err := client.CreateMandate(ctx, req)
	if err != nil {
		log.Printf("Mandate creation failed: %v", err)
		return ""
	}

	fmt.Printf("✓ Mandate created successfully!\n")
	fmt.Printf("  Mandate ID: %s\n", resp.MandateID)
	fmt.Printf("  Status: %s\n", resp.Status)
	fmt.Printf("  Contract Reference: %s\n", resp.ContractReference)
	fmt.Printf("  Debtor: %s\n", resp.DebtorName)
	fmt.Printf("  Maximum Amount: %.2f %s\n", resp.MaximumAmount, resp.Currency)
	fmt.Printf("  Frequency: %s\n", resp.FrequencyType)
	fmt.Printf("  First Collection: %s\n", resp.FirstCollectionDate)

	if resp.Status == "PENDING_APPROVAL" {
		fmt.Printf("\n  ⏳ Waiting for customer approval via their banking app...\n")
	}

	return resp.MandateID
}

func checkMandateStatusExample(ctx context.Context, client *fnb.Client, mandateID string) {
	if mandateID == "" {
		mandateID = "MAN123456789" // Replace with actual mandate ID
	}

	status, err := client.GetMandateStatus(ctx, mandateID)
	if err != nil {
		log.Printf("Failed to get mandate status: %v", err)
		return
	}

	fmt.Printf("Mandate Status:\n")
	fmt.Printf("  Mandate ID: %s\n", status.MandateID)
	fmt.Printf("  Contract Reference: %s\n", status.ContractReference)
	fmt.Printf("  Status: %s - %s\n", status.Status, status.StatusDescription)
	fmt.Printf("  Creditor: %s\n", status.CreditorName)
	fmt.Printf("  Debtor: %s (%s)\n", status.DebtorName, status.DebtorAccountNumber)
	fmt.Printf("  Maximum Amount: %.2f %s\n", status.MaximumAmount, status.Currency)
	fmt.Printf("  Frequency: %s\n", status.FrequencyType)
	fmt.Printf("  First Collection: %s\n", status.FirstCollectionDate)

	if status.NextCollectionDate != "" {
		fmt.Printf("  Next Collection: %s\n", status.NextCollectionDate)
	}

	if status.ApprovalDate != "" {
		fmt.Printf("  ✓ Approved on: %s\n", status.ApprovalDate)
	}

	if status.RejectionReason != "" {
		fmt.Printf("  ✗ Rejection Reason: %s\n", status.RejectionReason)
	}
}

func verifyMandateExample(ctx context.Context, client *fnb.Client, mandateID string) {
	if mandateID == "" {
		mandateID = "MAN123456789"
	}

	amount := 250.00

	valid, err := client.VerifyMandate(ctx, mandateID, amount)
	if err != nil {
		log.Printf("Mandate verification failed: %v", err)
		return
	}

	if valid {
		fmt.Printf("✓ Mandate is valid for collection\n")
		fmt.Printf("  Mandate ID: %s\n", mandateID)
		fmt.Printf("  Collection Amount: %.2f ZAR\n", amount)
		fmt.Printf("  Status: Ready to collect\n")
	} else {
		fmt.Printf("✗ Mandate is NOT valid for collection\n")
		fmt.Printf("  Mandate ID: %s\n", mandateID)
		fmt.Printf("  Please check mandate status\n")
	}
}

func collectAgainstMandateExample(ctx context.Context, client *fnb.Client, mandateID string) {
	if mandateID == "" {
		mandateID = "MAN123456789"
	}

	// First verify the mandate
	valid, err := client.VerifyMandate(ctx, mandateID, 250.00)
	if err != nil {
		log.Printf("Verification failed: %v", err)
		return
	}

	if !valid {
		fmt.Printf("✗ Cannot collect: Mandate is not valid\n")
		return
	}

	// Proceed with collection
	req := api.MandateCollectionRequest{
		MandateID:           mandateID,
		Amount:              250.00,
		CollectionReference: fmt.Sprintf("COLL-%s", time.Now().Format("20060102")),
		Description:         fmt.Sprintf("Monthly subscription - %s", time.Now().Format("January 2006")),
		IdempotencyKey:      fmt.Sprintf("coll-%s-%s", mandateID, time.Now().Format("20060102")),
	}

	resp, err := client.CollectAgainstMandate(ctx, req)
	if err != nil {
		log.Printf("Collection failed: %v", err)
		return
	}

	fmt.Printf("✓ Collection initiated successfully!\n")
	fmt.Printf("  Transaction ID: %s\n", resp.TransactionID)
	fmt.Printf("  Status: %s\n", resp.Status)
	fmt.Printf("  Collection Reference: %s\n", resp.CollectionReference)
	fmt.Printf("  Amount: %.2f %s\n", resp.Amount, resp.Currency)
	fmt.Printf("  Processing Date: %s\n", resp.ProcessingDate.Format("2006-01-02"))
}

func listMandatesExample(ctx context.Context, client *fnb.Client) {
	req := api.MandateListRequest{
		Status:     "ACTIVE",
		PageNumber: 1,
		PageSize:   10,
	}

	resp, err := client.ListMandates(ctx, req)
	if err != nil {
		log.Printf("Failed to list mandates: %v", err)
		return
	}

	fmt.Printf("Active Mandates (Page %d of %d):\n", resp.PageNumber, resp.TotalPages)
	fmt.Printf("Total: %d mandates\n\n", resp.TotalCount)

	for i, mandate := range resp.Mandates {
		fmt.Printf("%d. %s\n", i+1, mandate.ContractReference)
		fmt.Printf("   Mandate ID: %s\n", mandate.MandateID)
		fmt.Printf("   Debtor: %s\n", mandate.DebtorName)
		fmt.Printf("   Amount: %.2f %s (%s)\n", mandate.MaximumAmount, mandate.Currency, mandate.FrequencyType)
		fmt.Printf("   Next Collection: %s\n", mandate.NextCollectionDate)
		fmt.Printf("\n")
	}
}

func suspendReinstateExample(ctx context.Context, client *fnb.Client, mandateID string) {
	if mandateID == "" {
		mandateID = "MAN123456789"
	}

	// Suspend mandate
	suspendReq := api.MandateSuspensionRequest{
		MandateID:        mandateID,
		SuspensionReason: "Customer requested temporary suspension",
		SuspensionPeriod: 30, // 30 days
	}

	err := client.SuspendMandate(ctx, suspendReq)
	if err != nil {
		log.Printf("Failed to suspend mandate: %v", err)
		return
	}

	fmt.Printf("✓ Mandate suspended\n")
	fmt.Printf("  Mandate ID: %s\n", mandateID)
	fmt.Printf("  Suspension Period: 30 days\n")

	// Simulate waiting...
	fmt.Printf("\n  ... (time passes) ...\n\n")

	// Reinstate mandate
	err = client.ReinstateMandate(ctx, mandateID, "Customer requested reinstatement")
	if err != nil {
		log.Printf("Failed to reinstate mandate: %v", err)
		return
	}

	fmt.Printf("✓ Mandate reinstated\n")
	fmt.Printf("  Mandate ID: %s\n", mandateID)
	fmt.Printf("  Status: Active again\n")
}

func cancelMandateExample(ctx context.Context, client *fnb.Client, mandateID string) {
	if mandateID == "" {
		mandateID = "MAN123456789"
	}

	req := api.MandateCancellationRequest{
		MandateID:          mandateID,
		CancellationReason: "Customer cancelled subscription",
		EffectiveDate:      time.Now().Format("2006-01-02"),
	}

	err := client.CancelMandate(ctx, req)
	if err != nil {
		log.Printf("Failed to cancel mandate: %v", err)
		return
	}

	fmt.Printf("✓ Mandate cancelled successfully\n")
	fmt.Printf("  Mandate ID: %s\n", mandateID)
	fmt.Printf("  Reason: %s\n", req.CancellationReason)
	fmt.Printf("  Effective Date: %s\n", req.EffectiveDate)
}
