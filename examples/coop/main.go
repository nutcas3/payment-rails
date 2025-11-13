package main

import (
	"fmt"
	"log"
	"github.com/nutcas3/payment-rails/coop"
	"github.com/nutcas3/payment-rails/coop/pkg/api"
)

func main() {
	// Initialize Co-op Bank client
	client, err := coop.NewClient("your-client-id", "your-client-secret", api.SANDBOX)
	if err != nil {
		log.Fatal("Failed to create Co-op Bank client:", err)
	}

	// Example 1: Check Account Balance
	fmt.Println("=== Account Balance Example ===")
	balance, err := client.AccountBalance("1234567890")
	if err != nil {
		log.Printf("Account balance error: %v", err)
	} else {
		fmt.Printf("Account: %s\n", balance.AccountNumber)
		fmt.Printf("Account Name: %s\n", balance.AccountName)
		fmt.Printf("Currency: %s\n", balance.Currency)
		fmt.Printf("Cleared Balance: %.2f\n", balance.ClearedBalance)
		fmt.Printf("Booked Balance: %.2f\n", balance.BookedBalance)
		fmt.Printf("Product: %s\n", balance.ProductName)
	}

	// Example 2: Get Account Transactions
	fmt.Println("\n=== Account Transactions Example ===")
	transactions, err := client.AccountTransactions("1234567890", "5")
	if err != nil {
		log.Printf("Account transactions error: %v", err)
	} else {
		fmt.Printf("Account: %s (%s)\n", transactions.AccountNumber, transactions.AccountName)
		fmt.Printf("Found %d transactions:\n", len(transactions.Transactions))
		for i, txn := range transactions.Transactions {
			fmt.Printf("  %d. Date: %s, Amount: %.2f %s, Type: %s\n",
				i+1, txn.TransactionDate.Format("2006-01-02"), txn.Amount, txn.Currency, txn.TransactionType)
			fmt.Printf("     Narration: %s\n", txn.Narration)
		}
	}

	// Example 3: Get Exchange Rate
	fmt.Println("\n=== Exchange Rate Example ===")
	rate, err := client.ExchangeRate("KES", "USD")
	if err != nil {
		log.Printf("Exchange rate error: %v", err)
	} else {
		fmt.Printf("Exchange Rate: 1 %s = %.4f %s\n", 
			rate.FromCurrencyCode, rate.Rate, rate.ToCurrencyCode)
		fmt.Printf("Rate Type: %s\n", rate.RateType)
		fmt.Printf("Tolerance: %.4f\n", rate.Tolerance)
	}

	// Example 4: Internal Funds Transfer
	fmt.Println("\n=== Internal Funds Transfer Example ===")
	destinations := []api.Destination{
		{
			ReferenceNumber:     "REF001",
			AccountNumber:       "0987654321",
			Amount:              1000.00,
			TransactionCurrency: "KES",
			Narration:           "Payment for services",
		},
		{
			ReferenceNumber:     "REF002",
			AccountNumber:       "1122334455",
			Amount:              500.00,
			TransactionCurrency: "KES",
			Narration:           "Salary payment",
		},
	}

	iftResponse, err := client.InternalFundsTransfer(
		"1234567890",     // source account
		1500.00,          // total amount
		"KES",            // currency
		"Bulk transfer",  // narration
		destinations,
	)
	if err != nil {
		log.Printf("Internal funds transfer error: %v", err)
	} else {
		fmt.Printf("Transfer initiated successfully\n")
		fmt.Printf("Transaction ID: %s\n", iftResponse.TransactionID)
		fmt.Printf("Response Code: %s\n", iftResponse.ResponseCode)
		fmt.Printf("Response Message: %s\n", iftResponse.ResponseMessage)
	}

	// Example 5: PesaLink Transfer
	fmt.Println("\n=== PesaLink Transfer Example ===")
	pesaLinkDestinations := []api.PesaLinkDestination{
		{
			ReferenceNumber:     "PL001",
			DestinationBank:     "01", // KCB Bank code
			AccountNumber:       "1122334455",
			Amount:              2000.00,
			TransactionCurrency: "KES",
			Narration:           "Payment to supplier",
		},
	}

	pesaLinkResponse, err := client.PesaLinkTransfer(
		"1234567890",         // source account
		2000.00,              // total amount
		"KES",                // currency
		"PesaLink payment",   // narration
		pesaLinkDestinations,
	)
	if err != nil {
		log.Printf("PesaLink transfer error: %v", err)
	} else {
		fmt.Printf("PesaLink transfer initiated successfully\n")
		fmt.Printf("Transaction ID: %s\n", pesaLinkResponse.TransactionID)
		fmt.Printf("Response Code: %s\n", pesaLinkResponse.ResponseCode)
		fmt.Printf("Response Message: %s\n", pesaLinkResponse.ResponseMessage)
	}

	// Example 6: Check Transaction Status
	fmt.Println("\n=== Transaction Status Example ===")
	if iftResponse != nil && iftResponse.ResponseCode == "00" {
		// Use the message reference from the IFT transaction
		status, err := client.TransactionStatus(iftResponse.MessageReference)
		if err != nil {
			log.Printf("Transaction status error: %v", err)
		} else {
			fmt.Printf("Transaction Status: %s\n", status.TransactionStatus)
			fmt.Printf("Transaction ID: %s\n", status.TransactionID)
			fmt.Printf("Amount: %.2f\n", status.TransactionAmount)
			fmt.Printf("Date: %s\n", status.TransactionDate.Format("2006-01-02 15:04:05"))
		}
	}

	fmt.Println("\n=== Examples completed ===")
}
