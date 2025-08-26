package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shopspring/decimal"
	"payment-rails/absa"
	"payment-rails/absa/pkg/api"
)

func main() {
	// Replace these with your actual credentials
	clientID := os.Getenv("ABSA_CLIENT_ID")
	clientSecret := os.Getenv("ABSA_CLIENT_SECRET")
	apiKey := os.Getenv("ABSA_API_KEY")
	environment := "sandbox" // Use "production" for live environment

	if clientID == "" || clientSecret == "" || apiKey == "" {
		log.Fatal("Missing required environment variables: ABSA_CLIENT_ID, ABSA_CLIENT_SECRET, ABSA_API_KEY")
	}

	// Initialize the Absa client
	client, err := absa.NewClient(clientID, clientSecret, apiKey, environment)
	if err != nil {
		log.Fatalf("Failed to initialize Absa client: %v", err)
	}

	// Example 1: Account Balance
	fmt.Println("Example 1: Get Account Balance")
	balanceReq := api.AccountBalanceRequest{
		AccountNumber: "1234567890",
	}
	
	balance, err := client.GetAccountBalance(balanceReq)
	if err != nil {
		fmt.Printf("Error getting account balance: %v\n", err)
	} else {
		fmt.Printf("Account Balance: %s %s\n", api.FormatAmount(balance.AvailableBalance), balance.Currency)
	}
	fmt.Println()

	// Example 2: Account Validation
	fmt.Println("Example 2: Validate Account")
	validateReq := api.AccountValidateRequest{
		AccountNumber: "1234567890",
		BankCode:      "123",
	}
	
	validation, err := client.ValidateAccount(validateReq)
	if err != nil {
		fmt.Printf("Error validating account: %v\n", err)
	} else {
		fmt.Printf("Account Name: %s, Active: %t\n", validation.AccountName, validation.IsActive)
	}
	fmt.Println()

	// Example 3: Send Money (Bank Transfer)
	fmt.Println("Example 3: Send Money")
	reference := absa.GenerateReference()
	amount, _ := decimal.NewFromString("1000.00")
	sendMoneyReq := api.SendMoneyRequest{
		SourceAccount:       "1234567890",
		DestinationAccount:  "0987654321",
		DestinationBankCode: "123",
		Amount:              amount,
		Currency:            "KES",
		Reference:           reference,
		Description:         "Payment for services",
		BeneficiaryName:     "John Doe",
	}
	
	sendMoney, err := client.SendMoney(sendMoneyReq)
	if err != nil {
		fmt.Printf("Error sending money: %v\n", err)
	} else {
		// Display status using the defined constants
		statusText := "Unknown"
		switch sendMoney.Status {
		case api.StatusSuccess:
			statusText = "Successful"
		case api.StatusPending:
			statusText = "Pending"
		case api.StatusFailed:
			statusText = "Failed"
		case api.StatusProcessing:
			statusText = "Processing"
		default:
			statusText = sendMoney.Status
		}
		fmt.Printf("Transaction ID: %s, Status: %s\n", sendMoney.TransactionID, statusText)
	}
	fmt.Println()

	// Example 4: Send to Mobile Wallet
	fmt.Println("Example 4: Send to Mobile Wallet")
	mobileRef := absa.GenerateReference()
	mobileAmount, _ := decimal.NewFromString("500.00")
	mobileReq := api.MobileWalletRequest{
		SourceAccount: "1234567890",
		MobileNumber:  "254712345678",
		Amount:        mobileAmount,
		Currency:      "KES",
		Reference:     mobileRef,
		Description:   "Mobile money transfer",
		Provider:      "MPESA",
		CountryCode:   "KE",
	}
	
	mobileTransfer, err := client.SendToMobileWallet(mobileReq)
	if err != nil {
		fmt.Printf("Error sending to mobile wallet: %v\n", err)
	} else {
		fmt.Printf("Transaction ID: %s, Status: %s\n", mobileTransfer.TransactionID, mobileTransfer.Status)
	}
	fmt.Println()

	// Example 5: Query Transaction Status
	fmt.Println("Example 5: Query Transaction Status")
	queryReq := api.TransactionQueryRequest{
		Reference: reference,
		FromDate:  time.Now().AddDate(0, 0, -7), // Query transactions from the last 7 days
		ToDate:    time.Now(),
	}
	
	query, err := client.QueryTransaction(queryReq)
	if err != nil {
		fmt.Printf("Error querying transaction: %v\n", err)
	} else {
		fmt.Printf("Transaction Status: %s, Amount: %s %s, Date: %s\n", 
		query.Status, 
		api.FormatAmount(query.Amount), 
		query.Currency, 
		query.Timestamp.Format(time.RFC3339))
	}
	fmt.Println()

	// Example 6: Bulk Payments
	fmt.Println("Example 6: Bulk Payments")
	bulkRef := absa.GenerateReference()
	bulkItems := []api.BulkPaymentItem{
		{
			DestinationAccount:  "1111222233",
			DestinationBankCode: "123",
			Amount:              decimal.NewFromFloat(500.00),
			Reference:           absa.GenerateReference(),
			Description:         "Salary payment",
			BeneficiaryName:     "Employee One",
		},
		{
			DestinationAccount:  "4444555566",
			DestinationBankCode: "123",
			Amount:              decimal.NewFromFloat(750.00),
			Reference:           absa.GenerateReference(),
			Description:         "Salary payment",
			BeneficiaryName:     "Employee Two",
		},
	}
	
	bulkReq := api.BulkPaymentRequest{
		SourceAccount:  "1234567890",
		Currency:       "KES",
		BatchReference: bulkRef,
		Items:          bulkItems,
	}
	
	bulkPayment, err := client.ProcessBulkPayment(bulkReq)
	if err != nil {
		fmt.Printf("Error processing bulk payment: %v\n", err)
	} else {
		fmt.Printf("Batch ID: %s, Status: %s, Success Count: %d\n", 
			bulkPayment.BatchID, 
			bulkPayment.Status, 
			bulkPayment.SuccessCount)
	}
}
