package main

import (
	"fmt"
	"log"

	"payment-rails/ncba"
	"payment-rails/ncba/pkg/api"
)

func main() {
	// Initialize NCBA client
	client := ncba.NewClient(
		"your-api-key",
		"your-username",
		"your-password",
	)

	// Get account details
	accountDetails, err := client.GetAccountDetails("KE", "1234567890")
	if err != nil {
		log.Fatalf("Error getting account details: %v", err)
	}
	fmt.Printf("Account Details: %+v\n", accountDetails)

	// Get mini statement
	miniStatement, err := client.GetMiniStatement("KE", "1234567890")
	if err != nil {
		log.Fatalf("Error getting mini statement: %v", err)
	}
	fmt.Printf("Mini Statement: %+v\n", miniStatement)

	// Send internal transfer
	internalTransfer := api.InternalTransferRequest{
		TransferRequest: api.TransferRequest{
			SourceAccount:      "1234567890",
			DestinationAccount: "0987654321",
			Amount:            1000.00,
			Currency:          "KES",
			Reference:         "INV001",
			Narration:         "Payment for services",
		},
		DestinationName: "John Doe",
	}

	transferResp, err := client.SendInternalTransfer(internalTransfer)
	if err != nil {
		log.Fatalf("Error sending internal transfer: %v", err)
	}
	fmt.Printf("Transfer Response: %+v\n", transferResp)

	// Check transaction status
	status, err := client.CheckTransactionStatus(transferResp.TransactionID)
	if err != nil {
		log.Fatalf("Error checking transaction status: %v", err)
	}
	fmt.Printf("Transaction Status: %+v\n", status)

	// Send PesaLink transfer
	pesalinkTransfer := api.PesaLinkTransferRequest{
		TransferRequest: api.TransferRequest{
			SourceAccount:      "1234567890",
			DestinationAccount: "9876543210",
			Amount:            5000.00,
			Currency:          "KES",
			Reference:         "INV002",
			Narration:         "PesaLink transfer",
		},
		DestinationBank: "KCB",
		PhoneNumber:     "+254712345678",
	}

	pesalinkResp, err := client.SendPesaLinkTransfer(pesalinkTransfer)
	if err != nil {
		log.Fatalf("Error sending PesaLink transfer: %v", err)
	}
	fmt.Printf("PesaLink Transfer Response: %+v\n", pesalinkResp)
}
