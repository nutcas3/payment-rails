package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/nutcas3/payment-rails/momo"
	"github.com/nutcas3/payment-rails/momo/common/types"
)

func main() {
	targetEnv := os.Getenv("TARGET_ENVIRONMENT")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	RemittanceSubKey := os.Getenv("REMITTANCE_SUBSCRIPTION_KEY")

	cfg := momo.ClientConfig{
		Environment:               targetEnv,
		APIKey:                    apiKey,
		APISecret:                 apiSecret,
		RemittanceSubscriptionKey: RemittanceSubKey,
	}

	c, err := momo.New(cfg)
	if err != nil {
		log.Fatalf("error starting client: %v", err)
	}

	ctx := context.Background()
	callback := os.Getenv("CALLBACK_URL")

	// 1. GetAccountBalance example
	bal, err := c.Remittance.GetAccountBalance(ctx)
	if err != nil {
		fmt.Printf("error getting account balance: %v", err)
	}

	fmt.Printf("GetAccountBalance: %#v\n", bal)

	// 2. GetAccountBalanceInSpecificCurrency example
	bal, err = c.Remittance.GetAccountBalanceInSpecificCurrency(ctx, types.EUR)
	if err != nil {
		fmt.Printf("error getting account balance in specific currency: %v", err)
	}

	fmt.Printf("GetAccountBalanceInSpecificCurrency: %#v\n", bal)

	// 3. CreateAccessToken example
	_, err = c.Remittance.CreateAccessToken(ctx)
	if err != nil {
		fmt.Printf("error creating access token: %v", err)
	}

	// 4. CreateOauth2Token example
	_, err = c.Remittance.CreateOauth2Token(ctx)
	if err != nil {
		fmt.Printf("error creating oauth2 token: %v", err)
	}

	// 5. BcAuthorize example
	_, err = c.Remittance.BcAuthorize(ctx, callback)
	if err != nil {
		fmt.Printf("error doing BcAuthorize: %v", err)
	}

	// 6. GetBasicUserInfo example
	accHolderMsisdn := ""

	userInfo, err := c.Remittance.GetBasicUserInfo(ctx, accHolderMsisdn)
	if err != nil {
		fmt.Printf("error getting user info: %v", err)
	}

	fmt.Printf("GetBasicUserInfo: %#v\n", userInfo)

	// 7. GetUserInfoWithConsent example
	userInfoConsent, err := c.Remittance.GetUserInfoWithConsent(ctx)
	if err != nil {
		fmt.Printf("error getting user infor with consent: %v", err)
	}

	fmt.Printf("GetUserInfoWithConsent: %#v\n", userInfoConsent)

	// 8. CashTransfer example
	cashTransferInput := types.CashTransferInput{}
	cashTransferID := uuid.New()

	err = c.Remittance.CashTransfer(ctx, cashTransferID, callback, cashTransferInput)
	if err != nil {
		fmt.Printf("error initiating cash transfer: %v", err)
	}

	// 9. GetCashTransferStatus example
	cashTransferStatus, err := c.Remittance.GetCashTransferStatus(ctx, cashTransferID)
	if err != nil {
		fmt.Printf("error getting cash transfer status: %v", err)
	}

	fmt.Printf("GetCashTransferStatus: %#v\n", cashTransferStatus)

	// 10. Transfer example
	transferInput := types.TransferInput{
		Amount:     "100",
		Currency:   types.EUR,
		ExternalID: uuid.NewString(),
		Payee: types.Party{
			PartyIDType: types.PARTYCODE,
			PartyID:     uuid.NewString(),
		},
		PayerMessage: "test",
		PayeeNote:    "test",
	}
	transferID := uuid.New()

	err = c.Remittance.Transfer(ctx, transferID, callback, transferInput)
	if err != nil {
		fmt.Printf("error initiating transfer: %v", err)
	}

	// 11. GetTransferStatus example
	transferStatus, err := c.Remittance.GetTransferStatus(ctx, transferID)
	if err != nil {
		fmt.Printf("error getting transfer status: %v", err)
	}

	fmt.Printf("GetTransferStatus: %#v\n", transferStatus)
}
