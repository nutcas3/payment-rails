package main

import (
	"fmt"
	"log"
	"os"

	"payment-rails/jenga"
	"payment-rails/jenga/pkg/api"
)

func main() {
	// Get credentials from environment variables
	apiKey := os.Getenv("JENGA_API_KEY")
	username := os.Getenv("JENGA_USERNAME")
	password := os.Getenv("JENGA_PASSWORD")
	privateKey := os.Getenv("JENGA_PRIVATE_KEY")

	// Create a new Jenga client
	client, err := jenga.NewClient(apiKey, username, password, privateKey, "sandbox")
	if err != nil {
		log.Fatalf("Error creating Jenga client: %v", err)
	}

	// Example 1: Get Account Balance
	fmt.Println("Example 1: Get Account Balance")
	accountBalance, err := client.GetAccountBalance(api.AccountBalanceRequest{
		CountryCode:   "KE",
		AccountID:     "0011547896523",
	})
	if err != nil {
		log.Printf("Error getting account balance: %v", err)
	} else {
		fmt.Printf("Account Balance: %s %s\n", accountBalance.Data.Currency, accountBalance.Data.Balance)
	}

	// Example 2: Get Mini Statement
	fmt.Println("\nExample 2: Get Mini Statement")
	miniStatement, err := client.GetMiniStatement(api.MiniStatementRequest{
		CountryCode:   "KE",
		AccountID:     "0011547896523",
	})
	if err != nil {
		log.Printf("Error getting mini statement: %v", err)
	} else {
		fmt.Printf("Mini Statement: Account %s, Balance %s %s, %d transactions\n", 
			miniStatement.Data.AccountNumber, 
			miniStatement.Data.Balance, 
			miniStatement.Data.Currency, 
			len(miniStatement.Data.Transactions))
		for i, tx := range miniStatement.Data.Transactions {
			fmt.Printf("%d. %s: %s %s (%s)\n", i+1, tx.Date, tx.Amount, tx.Description, tx.Type)
		}
	}

	// Example 3: Send Money to Bank Account
	fmt.Println("\nExample 3: Send Money to Bank Account")
	reference := jenga.GenerateReference()
	sendMoneyReq := api.SendMoneyRequest{
		Source: api.Source{
			CountryCode:   "KE",
			Name:          "John Doe",
			AccountNumber: "0011547896523",
		},
		Destination: api.Destination{
			Type:          "bank",
			CountryCode:   "KE",
			Name:          "Jane Doe",
			AccountNumber: "0022547896523",
		},
		Transfer: api.Transfer{
			Type:         "EFT",
			Amount:       "1000.00",
			CurrencyCode: "KES",
			Reference:    reference,
			Date:         "2023-10-13",
			Description:  "Payment for services",
		},
	}
	sendMoneyResp, err := client.SendMoney(sendMoneyReq)
	if err != nil {
		log.Printf("Error sending money: %v", err)
	} else {
		fmt.Printf("Send Money Response: %s\n", sendMoneyResp.Data.TransactionID)
	}

	// Example 4: Send Money to Mobile Wallet
	fmt.Println("\nExample 4: Send Money to Mobile Wallet")
	mobileWalletReq := api.MobileWalletRequest{
		Source: api.Source{
			CountryCode:   "KE",
			Name:          "John Doe",
			AccountNumber: "0011547896523",
		},
		Destination: struct {
			Type          string `json:"type"`
			CountryCode   string `json:"countryCode"`
			Name          string `json:"name"`
			MobileNumber  string `json:"mobileNumber"`
			WalletName    string `json:"walletName"`
		}{
			Type:         "mobile",
			CountryCode:  "KE",
			Name:         "Jane Doe",
			MobileNumber: "254712345678",
			WalletName:   "Mpesa",
		},
		Transfer: struct {
			Type         string `json:"type"`
			Amount       string `json:"amount"`
			CurrencyCode string `json:"currencyCode"`
			Reference    string `json:"reference"`
			Date         string `json:"date"`
			Description  string `json:"description"`
			CallbackUrl  string `json:"callbackUrl"`
		}{
			Type:         "MobileWallet",
			Amount:       "500.00",
			CurrencyCode: "KES",
			Reference:    jenga.GenerateReference(),
			Date:         "2023-10-13",
			Description:  "Payment for services",
			CallbackUrl:  "https://webhook.site/your-webhook-id",
		},
	}
	mobileWalletResp, err := client.SendToMobileWallet(mobileWalletReq)
	if err != nil {
		log.Printf("Error sending to mobile wallet: %v", err)
	} else {
		fmt.Printf("Mobile Wallet Response: %s\n", mobileWalletResp.Data.TransactionID)
	}

	// Example 5: Pay Bill
	fmt.Println("\nExample 5: Pay Bill")
	billPaymentReq := api.BillPaymentRequest{
		BillerCode:   "KPLC",
		AccountNumber: "12345678",
		Amount:       "1000.00",
		Reference:    jenga.GenerateReference(),
		CurrencyCode: "KES",
		Narration:    "Electricity bill payment",
	}
	billPaymentResp, err := client.PayBill(billPaymentReq)
	if err != nil {
		log.Printf("Error paying bill: %v", err)
	} else {
		fmt.Printf("Bill Payment Response: %s\n", billPaymentResp.Data.TransactionID)
	}

	// Example 6: Purchase Airtime
	fmt.Println("\nExample 6: Purchase Airtime")
	airtimeReq := api.AirtimePurchaseRequest{
		CustomerMobile: "254712345678",
		TelcoCode:      "Safaricom",
		Amount:         "100.00",
		Reference:      jenga.GenerateReference(),
		CurrencyCode:   "KES",
	}
	airtimeResp, err := client.PurchaseAirtime(airtimeReq)
	if err != nil {
		log.Printf("Error purchasing airtime: %v", err)
	} else {
		fmt.Printf("Airtime Purchase Response: %s\n", airtimeResp.Data.TransactionID)
	}

	// Example 7: KYC Verification
	fmt.Println("\nExample 7: KYC Verification")
	kycReq := api.KYCRequest{
		DocumentType:   "ID",
		DocumentNumber: "12345678",
		CountryCode:    "KE",
		DateOfBirth:    "1990-01-01",
	}
	kycResp, err := client.VerifyIdentity(kycReq)
	if err != nil {
		log.Printf("Error verifying identity: %v", err)
	} else {
		fmt.Printf("KYC Verification Response: %s %s\n", kycResp.Data.FirstName, kycResp.Data.LastName)
	}

	// Example 8: Send Money via RTGS
	fmt.Println("\nExample 8: Send Money via RTGS")
	rtgsReq := api.SendMoneyRequest{
		Source: api.Source{
			CountryCode:   "KE", // Kenya
			Name:          "John Doe",
			AccountNumber: "0011547896523",
		},
		Destination: api.Destination{
			Type:          "bank",
			CountryCode:   "KE", // Kenya
			Name:          "Jane Smith",
			AccountNumber: "0987654321",
			BankCode:      "01", // Bank code for the receiving bank
			BranchCode:    "112", // Branch code for the receiving bank
		},
		Transfer: api.Transfer{
			Type:         api.TransferTypeRTGS, // Specify RTGS transfer type
			Amount:       "50000.00", // RTGS is typically used for larger amounts
			CurrencyCode: "KES",
			Reference:    jenga.GenerateReference(),
			Date:         "2023-10-13",
			Description:  "RTGS Transfer to Jane Smith",
		},
	}
	rtgsResp, err := client.SendMoney(rtgsReq)
	if err != nil {
		log.Printf("Error sending money via RTGS: %v", err)
	} else {
		fmt.Printf("RTGS Transfer Response: %s\n", rtgsResp.Data.TransactionID)
	}

	// Example 9: Send Money via SWIFT
	fmt.Println("\nExample 9: Send Money via SWIFT")
	swiftReq := api.SendMoneyRequest{
		Source: api.Source{
			CountryCode:   "KE", // Kenya
			Name:          "John Doe",
			AccountNumber: "0011547896523",
		},
		Destination: api.Destination{
			Type:          "bank",
			CountryCode:   "US", // United States
			Name:          "Jane Smith",
			AccountNumber: "0987654321",
			// SWIFT specific fields
			BankName:      "Bank of America",
			BankAddress:   "100 North Tryon Street, Charlotte, NC 28255, USA",
			SwiftCode:     "BOFAUS3N", // Example SWIFT code for Bank of America
			RoutingNumber: "026009593", // ABA routing number for US banks
		},
		Transfer: api.Transfer{
			Type:         api.TransferTypeSWIFT, // Specify SWIFT transfer type
			Amount:       "1000.00",
			CurrencyCode: "USD", // International transfer in USD
			Reference:    jenga.GenerateReference(),
			Date:         "2023-10-13",
			Description:  "International SWIFT Transfer to Jane Smith",
		},
	}
	swiftResp, err := client.SendMoney(swiftReq)
	if err != nil {
		log.Printf("Error sending money via SWIFT: %v", err)
	} else {
		fmt.Printf("SWIFT Transfer Response: %s\n", swiftResp.Data.TransactionID)
	}

	// Example 10: Get Forex Rates
	fmt.Println("\nExample 10: Get Forex Rates")
	forexReq := api.ForexRatesRequest{
		CountryCode:  "KE",
		CurrencyCode: "USD",
		BaseCurrency: "KES",
	}
	forexResp, err := client.GetForexRates(forexReq)
	if err != nil {
		log.Printf("Error getting forex rates: %v", err)
	} else {
		fmt.Printf("Forex Rate: 1 %s = %s %s\n", forexResp.Data.BaseCurrency, forexResp.Data.Rates[0].MeanRate, forexResp.Data.Rates[0].CurrencyCode)
	}
}
