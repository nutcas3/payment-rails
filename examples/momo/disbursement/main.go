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
	DisbursementSubKey := os.Getenv("DISBURSEMENT_SUBSCRIPTION_KEY")

	cfg := momo.ClientConfig{
		Environment:                 targetEnv,
		APIKey:                      apiKey,
		APISecret:                   apiSecret,
		DisbursementSubscriptionKey: DisbursementSubKey,
	}

	c, err := momo.New(cfg)
	if err != nil {
		log.Fatalf("error starting client: %v", err)
	}

	ctx := context.Background()
	callback := os.Getenv("CALLBACK_URL")

	// 1. GetAccountBalance example
	bal, err := c.Disbursement.GetAccountBalance(ctx)
	if err != nil {
		fmt.Printf("error getting account balance: %v", err)
	}

	fmt.Printf("Account balance: %#v\n", bal)

	// 2. GetAccountBalanceInSpecificCurrency example
	bal, err = c.Disbursement.GetAccountBalanceInSpecificCurrency(ctx, types.EUR)
	if err != nil {
		fmt.Printf("error getting account balance in specific currency: %v", err)
	}

	fmt.Printf("Account balance in specific currency: %#v\n", bal)

	// 3. CreateAccessToken example
	_, err = c.Disbursement.CreateAccessToken(ctx)
	if err != nil {
		fmt.Printf("error creating access token: %v", err)
	}

	// 4. CreateOauth2Token example
	_, err = c.Disbursement.CreateOauth2Token(ctx)
	if err != nil {
		fmt.Printf("error creating oauth2 token: %v", err)
	}

	// 5. BcAuthorize example
	_, err = c.Disbursement.BcAuthorize(ctx, callback)
	if err != nil {
		fmt.Printf("error doing BcAuthorize: %v", err)
	}

	// 6. GetBasicUserInfo example
	accHolderIDType := ""
	accHolderID := ""

	userInfo, err := c.Disbursement.GetBasicUserInfo(ctx, accHolderIDType, accHolderID)
	if err != nil {
		fmt.Printf("error getting user info: %v", err)
	}

	fmt.Printf("GetBasicUserInfo: %#v\n", userInfo)

	// 7. GetUserInfoWithConsent example
	userInfoConsent, err := c.Disbursement.GetUserInfoWithConsent(ctx)
	if err != nil {
		fmt.Printf("error getting user infor with consent: %v", err)
	}

	fmt.Printf("GetUserInfoWithConsent: %#v\n", userInfoConsent)

	// 8. DepositV1 example
	depositInput := types.TransferInput{
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
	depositID := uuid.New()

	err = c.Disbursement.DepositV1(ctx, depositID, callback, depositInput)
	if err != nil {
		fmt.Printf("error initiating deposit(V1): %v", err)
	}

	// 9. DepositV2 example
	err = c.Disbursement.DepositV2(ctx, depositID, callback, depositInput)
	if err != nil {
		fmt.Printf("error initiating deposit(V2): %v", err)
	}

	// 10. GetDepositStatus example
	depositStatus, err := c.Disbursement.GetDepositStatus(ctx, depositID)
	if err != nil {
		fmt.Printf("error deposit status: %v", err)
	}

	fmt.Printf("GetDepositStatus: %#v\n", depositStatus)

	// 11. Transfer example
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

	err = c.Disbursement.Transfer(ctx, transferID, callback, transferInput)
	if err != nil {
		fmt.Printf("error initiating transfer: %v", err)
	}

	// 12. GetTransferStatus example
	transferStatus, err := c.Disbursement.GetTransferStatus(ctx, transferID)
	if err != nil {
		fmt.Printf("error getting transfer status: %v", err)
	}

	fmt.Printf("GetTransferStatus: %#v\n", transferStatus)

	// 13. RefundV1 example
	refundInput := types.RefundInput{
		Amount:              "100",
		Currency:            types.EUR,
		ExternalID:          uuid.NewString(),
		PayerMessage:        "test",
		PayeeNote:           "test",
		ReferenceIDToRefund: uuid.NewString(),
	}
	refundID := uuid.New()

	err = c.Disbursement.RefundV1(ctx, refundID, callback, refundInput)
	if err != nil {
		fmt.Printf("error initiating refund(V1): %v", err)
	}

	// 14. RefundV2 example
	err = c.Disbursement.RefundV2(ctx, refundID, callback, refundInput)
	if err != nil {
		fmt.Printf("error initiating refund(V2): %v", err)
	}

	// 15. GetRefundStatus example
	refundStatus, err := c.Disbursement.GetRefundStatus(ctx, refundID)
	if err != nil {
		fmt.Printf("error getting refund status: %v", err)
	}

	fmt.Printf("GetRefundStatus: %#v\n", refundStatus)
}
