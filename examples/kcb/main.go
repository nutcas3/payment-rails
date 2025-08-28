package main

import (
	"fmt"
	"log"
	"github.com/nutcas3/payment-rails/kcb"
	"time"
)

func main() {
	// Initialize the KCB Buni client
	client, err := kcb.New(
		"your-api-token", // API token from KCB Buni Developer Portal
		true,             // true for sandbox, false for production
	)
	if err != nil {
		log.Fatalf("Failed to initialize KCB Buni client: %v", err)
	}

	fmt.Println("=== KCB Buni API Examples ===")

	// Account Information Examples
	accountExamples(client)

	// Forex Examples
	forexExamples(client)

	// Payment Examples
	paymentExamples(client)

	// PesaLink Examples
	pesalinkExamples(client)

	// Mobile Money Examples
	mobileMoneyExamples(client)

	// Utility Payment Examples
	utilityExamples(client)
}

func accountExamples(client *kcb.Client) {
	fmt.Println("\n=== Account Information Examples ===")

	// Example 1: Get Account Information
	fmt.Println("\n1. Get Account Information:")
	accountInfo, err := client.GetAccountInfo()
	if err != nil {
		log.Printf("Failed to get account information: %v", err)
	} else {
		fmt.Printf("  Account Number: %s\n", accountInfo.Data.AccountNumber)
		fmt.Printf("  Account Name: %s\n", accountInfo.Data.AccountName)
		fmt.Printf("  Balance: %.2f %s\n", accountInfo.Data.Balance, accountInfo.Data.Currency)
		fmt.Printf("  Account Type: %s\n", accountInfo.Data.AccountType)
		fmt.Printf("  Branch: %s\n", accountInfo.Data.Branch)
		fmt.Printf("  Status: %s\n", accountInfo.Data.Status)
	}

	// Example 2: Get Account Balance
	fmt.Println("\n2. Get Account Balance:")
	balance, err := client.GetAccountBalance("1234567890")
	if err != nil {
		log.Printf("Failed to get account balance: %v", err)
	} else {
		fmt.Printf("  Account Number: %s\n", balance.Data.AccountNumber)
		fmt.Printf("  Account Name: %s\n", balance.Data.AccountName)
		fmt.Printf("  Balance: %.2f %s\n", balance.Data.Balance, balance.Data.Currency)
		fmt.Printf("  As Of: %s\n", balance.Data.AsOf)
	}

	// Example 3: Get Account Statement
	fmt.Println("\n3. Get Account Statement:")
	// Get statement for the last month
	now := time.Now()
	endDate := now.Format("2006-01-02")
	startDate := now.AddDate(0, -1, 0).Format("2006-01-02")

	statement, err := client.GetAccountStatement("1234567890", startDate, endDate)
	if err != nil {
		log.Printf("Failed to get account statement: %v", err)
	} else {
		fmt.Printf("  Account Number: %s\n", statement.Data.AccountNumber)
		fmt.Printf("  Account Name: %s\n", statement.Data.AccountName)
		fmt.Printf("  Period: %s to %s\n", statement.Data.StartDate, statement.Data.EndDate)
		fmt.Printf("  Currency: %s\n", statement.Data.Currency)
		fmt.Printf("  Transactions: %d\n", len(statement.Data.Transactions))

		// Print first 5 transactions or all if less than 5
		txCount := len(statement.Data.Transactions)
		if txCount > 5 {
			txCount = 5
		}

		for i := 0; i < txCount; i++ {
			tx := statement.Data.Transactions[i]
			fmt.Printf("    %s: %s %.2f - %s\n",
				tx.TransactionDate.Format("2006-01-02"),
				tx.Type,
				tx.Amount,
				tx.Description)
		}
	}

	// Example 4: Transfer Funds
	fmt.Println("\n4. Transfer Funds:")
	transfer, err := client.TransferFunds(
		"1234567890",     // Source account
		"0987654321",     // Destination account
		1000.0,           // Amount
		"KES",            // Currency
		"INV123456",      // Reference
		"Invoice payment", // Narration
	)
	if err != nil {
		log.Printf("Failed to transfer funds: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", transfer.Data.TransactionID)
		fmt.Printf("  From Account: %s\n", transfer.Data.SourceAccount)
		fmt.Printf("  To Account: %s\n", transfer.Data.DestAccount)
		fmt.Printf("  Amount: %.2f %s\n", transfer.Data.Amount, transfer.Data.Currency)
		fmt.Printf("  Status: %s\n", transfer.Data.Status)
		fmt.Printf("  Transaction Date: %s\n", transfer.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", transfer.Data.Reference)
	}
}

func forexExamples(client *kcb.Client) {
	fmt.Println("\n=== Forex Examples ===")

	// Example 1: Get Forex Rates
	fmt.Println("\n1. Get Forex Rates:")
	forexRates, err := client.GetForexRates("USD")
	if err != nil {
		log.Printf("Failed to get forex rates: %v", err)
	} else {
		fmt.Printf("  Base Currency: %s\n", forexRates.Data.BaseCurrency)
		fmt.Printf("  Timestamp: %s\n", forexRates.Data.Timestamp)
		fmt.Println("  Rates:")
		for currency, rate := range forexRates.Data.Rates {
			fmt.Printf("    %s: %.4f\n", currency, rate)
		}
	}

	// Example 2: Exchange Currency
	fmt.Println("\n2. Exchange Currency:")
	exchange, err := client.ExchangeCurrency("EUR", "USD", 100.0)
	if err != nil {
		log.Printf("Failed to exchange currency: %v", err)
	} else {
		fmt.Printf("  From: %s %.2f\n", exchange.Data.FromCurrency, exchange.Data.Amount)
		fmt.Printf("  To: %s %.2f\n", exchange.Data.ToCurrency, exchange.Data.ConvertedAmount)
		fmt.Printf("  Exchange Rate: %.4f\n", exchange.Data.ExchangeRate)
		fmt.Printf("  Timestamp: %s\n", exchange.Data.Timestamp)
	}
}

func paymentExamples(client *kcb.Client) {
	fmt.Println("\n=== Payment Examples ===")

	// Example 1: Make Vooma Payment
	fmt.Println("\n1. Make Vooma Payment:")
	payment, err := client.VoomaPay(100.0)
	if err != nil {
		log.Printf("Failed to make Vooma payment: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", payment.Data.TransactionID)
		fmt.Printf("  Amount: %.2f %s\n", payment.Data.Amount, payment.Data.Currency)
		fmt.Printf("  Status: %s\n", payment.Data.Status)
		fmt.Printf("  Transaction Date: %s\n", payment.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", payment.Data.Reference)
	}

	// Example 2: Check Vooma Payment Status
	fmt.Println("\n2. Check Vooma Payment Status:")
	// Using a sample transaction ID - in real usage, use the ID from the payment response
	voomaStatus, err := client.CheckVoomaStatus("TX123456789")
	if err != nil {
		log.Printf("Failed to check Vooma payment status: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", voomaStatus.Data.TransactionID)
		fmt.Printf("  Amount: %.2f %s\n", voomaStatus.Data.Amount, voomaStatus.Data.Currency)
		fmt.Printf("  Status: %s\n", voomaStatus.Data.Status)
		if voomaStatus.Data.StatusReason != "" {
			fmt.Printf("  Status Reason: %s\n", voomaStatus.Data.StatusReason)
		}
		fmt.Printf("  Transaction Date: %s\n", voomaStatus.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", voomaStatus.Data.Reference)
	}
}

func pesalinkExamples(client *kcb.Client) {
	fmt.Println("\n=== PesaLink Examples ===")

	// Example 1: PesaLink Transfer
	fmt.Println("\n1. PesaLink Transfer:")
	transfer, err := client.PesalinkTransfer(
		"1234567890",     // Source account
		"0987654321",     // Destination account
		"01",             // Destination bank code (e.g., 01 for KCB)
		1000.0,           // Amount
		"KES",            // Currency
		"INV123456",      // Reference
		"Invoice payment", // Narration
		"254712345678",    // Phone number
	)
	if err != nil {
		log.Printf("Failed to make PesaLink transfer: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", transfer.Data.TransactionID)
		fmt.Printf("  From Account: %s\n", transfer.Data.SourceAccount)
		fmt.Printf("  To Account: %s (Bank: %s)\n", transfer.Data.DestAccount, transfer.Data.DestBank)
		fmt.Printf("  Amount: %.2f %s\n", transfer.Data.Amount, transfer.Data.Currency)
		fmt.Printf("  Status: %s\n", transfer.Data.Status)
		fmt.Printf("  Transaction Date: %s\n", transfer.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", transfer.Data.Reference)
	}

	// Example 2: Check PesaLink Transfer Status
	fmt.Println("\n2. Check PesaLink Transfer Status:")
	// Using a sample transaction ID - in real usage, use the ID from the transfer response
	pesalinkStatus, err := client.CheckPesalinkStatus("TX123456789")
	if err != nil {
		log.Printf("Failed to check PesaLink transfer status: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", pesalinkStatus.Data.TransactionID)
		fmt.Printf("  From Account: %s\n", pesalinkStatus.Data.SourceAccount)
		fmt.Printf("  To Account: %s (Bank: %s)\n", pesalinkStatus.Data.DestAccount, pesalinkStatus.Data.DestBank)
		fmt.Printf("  Amount: %.2f %s\n", pesalinkStatus.Data.Amount, pesalinkStatus.Data.Currency)
		fmt.Printf("  Status: %s\n", pesalinkStatus.Data.Status)
		if pesalinkStatus.Data.StatusReason != "" {
			fmt.Printf("  Status Reason: %s\n", pesalinkStatus.Data.StatusReason)
		}
		fmt.Printf("  Transaction Date: %s\n", pesalinkStatus.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", pesalinkStatus.Data.Reference)
	}
}

func mobileMoneyExamples(client *kcb.Client) {
	fmt.Println("\n=== Mobile Money Examples ===")

	// Example 1: Mobile Money Transfer
	fmt.Println("\n1. Mobile Money Transfer:")
	transfer, err := client.MobileMoneyTransfer(
		"1234567890",     // Source account
		"254712345678",   // Phone number
		1000.0,           // Amount
		"KES",            // Currency
		"INV123456",      // Reference
		"Invoice payment", // Narration
		"MPESA",           // Provider
	)
	if err != nil {
		log.Printf("Failed to make mobile money transfer: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", transfer.Data.TransactionID)
		fmt.Printf("  From Account: %s\n", transfer.Data.SourceAccount)
		fmt.Printf("  To Phone: %s (Provider: %s)\n", transfer.Data.PhoneNumber, transfer.Data.Provider)
		fmt.Printf("  Amount: %.2f %s\n", transfer.Data.Amount, transfer.Data.Currency)
		fmt.Printf("  Status: %s\n", transfer.Data.Status)
		fmt.Printf("  Transaction Date: %s\n", transfer.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", transfer.Data.Reference)
	}

	// Example 2: Check Mobile Money Transfer Status
	fmt.Println("\n2. Check Mobile Money Transfer Status:")
	// Using a sample transaction ID - in real usage, use the ID from the transfer response
	mobileStatus, err := client.CheckMobileMoneyStatus("TX123456789")
	if err != nil {
		log.Printf("Failed to check mobile money transfer status: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", mobileStatus.Data.TransactionID)
		fmt.Printf("  From Account: %s\n", mobileStatus.Data.SourceAccount)
		fmt.Printf("  To Phone: %s (Provider: %s)\n", mobileStatus.Data.PhoneNumber, mobileStatus.Data.Provider)
		fmt.Printf("  Amount: %.2f %s\n", mobileStatus.Data.Amount, mobileStatus.Data.Currency)
		fmt.Printf("  Status: %s\n", mobileStatus.Data.Status)
		if mobileStatus.Data.StatusReason != "" {
			fmt.Printf("  Status Reason: %s\n", mobileStatus.Data.StatusReason)
		}
		fmt.Printf("  Transaction Date: %s\n", mobileStatus.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", mobileStatus.Data.Reference)
	}
}

func utilityExamples(client *kcb.Client) {
	fmt.Println("\n=== Utility Payment Examples ===")

	// Example 1: Get Utility Providers
	fmt.Println("\n1. Get Utility Providers:")
	providers, err := client.GetUtilityProviders()
	if err != nil {
		log.Printf("Failed to get utility providers: %v", err)
	} else {
		fmt.Printf("  Available Providers: %d\n", len(providers.Data.Providers))
		// Print first 5 providers or all if less than 5
		provCount := len(providers.Data.Providers)
		if provCount > 5 {
			provCount = 5
		}

		for i := 0; i < provCount; i++ {
			prov := providers.Data.Providers[i]
			fmt.Printf("    %s - %s (%s)\n",
				prov.ProviderID,
				prov.ProviderName,
				prov.Category)
		}
	}

	// Example 2: Pay Utility Bill
	fmt.Println("\n2. Pay Utility Bill:")
	payment, err := client.PayUtility(
		"1234567890",     // Source account
		"KPLC",           // Provider ID
		"12345678",       // Account number with provider
		1000.0,           // Amount
		"KES",            // Currency
		"BILL123456",     // Reference
		"254712345678",   // Phone number for notifications
	)
	if err != nil {
		log.Printf("Failed to pay utility bill: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", payment.Data.TransactionID)
		fmt.Printf("  From Account: %s\n", payment.Data.SourceAccount)
		fmt.Printf("  Provider: %s (%s)\n", payment.Data.ProviderName, payment.Data.ProviderID)
		fmt.Printf("  Customer Account: %s\n", payment.Data.AccountNumber)
		fmt.Printf("  Amount: %.2f %s\n", payment.Data.Amount, payment.Data.Currency)
		fmt.Printf("  Status: %s\n", payment.Data.Status)
		fmt.Printf("  Transaction Date: %s\n", payment.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", payment.Data.Reference)
		if payment.Data.ReceiptNumber != "" {
			fmt.Printf("  Receipt Number: %s\n", payment.Data.ReceiptNumber)
		}
	}

	// Example 3: Check Utility Payment Status
	fmt.Println("\n3. Check Utility Payment Status:")
	// Using a sample transaction ID - in real usage, use the ID from the payment response
	utilityStatus, err := client.CheckUtilityPaymentStatus("TX123456789")
	if err != nil {
		log.Printf("Failed to check utility payment status: %v", err)
	} else {
		fmt.Printf("  Transaction ID: %s\n", utilityStatus.Data.TransactionID)
		fmt.Printf("  From Account: %s\n", utilityStatus.Data.SourceAccount)
		fmt.Printf("  Provider: %s (%s)\n", utilityStatus.Data.ProviderName, utilityStatus.Data.ProviderID)
		fmt.Printf("  Customer Account: %s\n", utilityStatus.Data.AccountNumber)
		fmt.Printf("  Amount: %.2f %s\n", utilityStatus.Data.Amount, utilityStatus.Data.Currency)
		fmt.Printf("  Status: %s\n", utilityStatus.Data.Status)
		if utilityStatus.Data.StatusReason != "" {
			fmt.Printf("  Status Reason: %s\n", utilityStatus.Data.StatusReason)
		}
		fmt.Printf("  Transaction Date: %s\n", utilityStatus.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", utilityStatus.Data.Reference)
		if utilityStatus.Data.ReceiptNumber != "" {
			fmt.Printf("  Receipt Number: %s\n", utilityStatus.Data.ReceiptNumber)
		}
	}
}
