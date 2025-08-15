package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"payment-rails/airtel/pkg/api"
)

func main() {
	// Replace these with your actual credentials
	clientID := os.Getenv("AIRTEL_CLIENT_ID")
	clientSecret := os.Getenv("AIRTEL_CLIENT_SECRET")
	publicKey := os.Getenv("AIRTEL_PUBLIC_KEY")
	
	if clientID == "" || clientSecret == "" {
		log.Fatal("AIRTEL_CLIENT_ID and AIRTEL_CLIENT_SECRET environment variables must be set")
	}

	// Initialize the Airtel Money client
	// Use "KE" for Kenya and "KES" for Kenyan Shilling
	airtel, err := api.New(clientID, clientSecret, publicKey, api.SANDBOX, "KE", "KES")
	if err != nil {
		log.Fatalf("Failed to initialize Airtel Money client: %v", err)
	}

	// Example 1: Check account balance
	fmt.Println("Checking account balance...")
	balance, err := airtel.GetAccountBalance()
	if err != nil {
		log.Printf("Failed to get account balance: %v", err)
	} else {
		fmt.Printf("Account Balance: %.2f %s\n", balance.Data.Balance, balance.Data.Currency)
	}

	// Example 2: USSD Push (Collection)
	fmt.Println("\nInitiating USSD Push payment...")
	reference := "TEST-REF-" + time.Now().Format("20060102150405")
	phone := "700000000" // Replace with actual phone number without country code
	amount := 10.0       // Amount in the specified currency
	txID := "TX-" + time.Now().Format("20060102150405")

	collectionResp, err := airtel.UssdPush(reference, phone, amount, txID)
	if err != nil {
		log.Printf("Failed to initiate USSD Push: %v", err)
	} else {
		fmt.Printf("USSD Push initiated: Success=%v, Message=%s\n", 
			collectionResp.Status.Success, 
			collectionResp.Status.Message)
		
		if collectionResp.Status.Success {
			// Check transaction status after a few seconds
			fmt.Println("\nChecking transaction status...")
			time.Sleep(5 * time.Second)
			
			statusResp, err := airtel.GetTransactionStatus(txID)
			if err != nil {
				log.Printf("Failed to check transaction status: %v", err)
			} else {
				fmt.Printf("Transaction Status: %s, Message: %s\n", 
					statusResp.Data.Transaction.Status, 
					statusResp.Status.Message)
			}
		}
	}

	// Example 3: Disbursement (Send Money)
	fmt.Println("\nInitiating disbursement...")
	disbReference := "DISB-REF-" + time.Now().Format("20060102150405")
	disbTxID := "DISB-TX-" + time.Now().Format("20060102150405")
	pin := "1234" // Replace with your actual PIN

	disbResp, err := airtel.Disburse(disbReference, phone, amount, disbTxID, pin)
	if err != nil {
		log.Printf("Failed to initiate disbursement: %v", err)
	} else {
		fmt.Printf("Disbursement initiated: Success=%v, Message=%s\n", 
			disbResp.Status.Success, 
			disbResp.Status.Message)
		
		if disbResp.Status.Success {
			// Check disbursement status after a few seconds
			fmt.Println("\nChecking disbursement status...")
			time.Sleep(5 * time.Second)
			
			disbStatusResp, err := airtel.GetDisbursementStatus(disbTxID)
			if err != nil {
				log.Printf("Failed to check disbursement status: %v", err)
			} else {
				fmt.Printf("Disbursement Status: %s, Message: %s\n", 
					disbStatusResp.Data.Transaction.Status, 
					disbStatusResp.Status.Message)
			}
		}
	}
}
