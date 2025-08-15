package main

import (
	"fmt"
	"log"
	"payment-rails/kcb"
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

	// Example 1: Get Account Information
	accountInfo, err := client.GetAccountInfo()
	if err != nil {
		log.Printf("Failed to get account information: %v", err)
	} else {
		fmt.Printf("Account Information:\n")
		fmt.Printf("  Account Number: %s\n", accountInfo.Data.AccountNumber)
		fmt.Printf("  Account Name: %s\n", accountInfo.Data.AccountName)
		fmt.Printf("  Balance: %.2f %s\n", accountInfo.Data.Balance, accountInfo.Data.Currency)
		fmt.Printf("  Account Type: %s\n", accountInfo.Data.AccountType)
		fmt.Printf("  Branch: %s\n", accountInfo.Data.Branch)
		fmt.Printf("  Status: %s\n", accountInfo.Data.Status)
	}

	// Example 2: Get Forex Rates
	forexRates, err := client.GetForexRates("USD")
	if err != nil {
		log.Printf("Failed to get forex rates: %v", err)
	} else {
		fmt.Printf("\nForex Rates for %s:\n", forexRates.Data.BaseCurrency)
		for currency, rate := range forexRates.Data.Rates {
			fmt.Printf("  %s: %.4f\n", currency, rate)
		}
		fmt.Printf("  Timestamp: %s\n", forexRates.Data.Timestamp)
	}

	// Example 3: Exchange Currency
	exchange, err := client.ExchangeCurrency("EUR", "USD", 100.0)
	if err != nil {
		log.Printf("Failed to exchange currency: %v", err)
	} else {
		fmt.Printf("\nCurrency Exchange:\n")
		fmt.Printf("  From: %s %.2f\n", exchange.Data.FromCurrency, exchange.Data.Amount)
		fmt.Printf("  To: %s %.2f\n", exchange.Data.ToCurrency, exchange.Data.ConvertedAmount)
		fmt.Printf("  Exchange Rate: %.4f\n", exchange.Data.ExchangeRate)
		fmt.Printf("  Timestamp: %s\n", exchange.Data.Timestamp)
	}

	// Example 4: Make Vooma Payment
	payment, err := client.VoomaPay(100.0)
	if err != nil {
		log.Printf("Failed to make Vooma payment: %v", err)
	} else {
		fmt.Printf("\nVooma Payment:\n")
		fmt.Printf("  Transaction ID: %s\n", payment.Data.TransactionID)
		fmt.Printf("  Amount: %.2f %s\n", payment.Data.Amount, payment.Data.Currency)
		fmt.Printf("  Status: %s\n", payment.Data.Status)
		fmt.Printf("  Transaction Date: %s\n", payment.Data.TransactionDate)
		fmt.Printf("  Reference: %s\n", payment.Data.Reference)
	}
}
