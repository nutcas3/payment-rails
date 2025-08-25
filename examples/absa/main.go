package main

import (
	"fmt"
	"log"
	"os"
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
		fmt.Printf("Account Balance: %s %s\n", balance.AvailableBalance, balance.Currency)
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
	sendMoneyReq := api.SendMoneyRequest{
		SourceAccount:       "1234567890",
		DestinationAccount:  "0987654321",
		DestinationBankCode: "123",
		Amount:              "1000.00",
		Currency:            "KES",
		Reference:           reference,
		Description:         "Payment for services",
		BeneficiaryName:     "John Doe",
	}
	
	sendMoney, err := client.SendMoney(sendMoneyReq)
	if err != nil {
		fmt.Printf("Error sending money: %v\n", err)
	} else {
		fmt.Printf("Transaction ID: %s, Status: %s\n", sendMoney.TransactionID, sendMoney.Status)
	}
	fmt.Println()

	// Example 4: Send to Mobile Wallet
	fmt.Println("Example 4: Send to Mobile Wallet")
	mobileRef := absa.GenerateReference()
	mobileReq := api.MobileWalletRequest{
		SourceAccount: "1234567890",
		MobileNumber:  "254712345678",
		Amount:        "500.00",
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
	}
	
	query, err := client.QueryTransaction(queryReq)
	if err != nil {
		fmt.Printf("Error querying transaction: %v\n", err)
	} else {
		fmt.Printf("Transaction Status: %s, Amount: %s %s\n", query.Status, query.Amount, query.Currency)
	}
}
