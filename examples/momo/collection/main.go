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
	CollectionSubKey := os.Getenv("COLLECTION_SUBSCRIPTION_KEY")

	cfg := momo.ClientConfig{
		Environment:               targetEnv,
		APIKey:                    apiKey,
		APISecret:                 apiSecret,
		CollectionSubscriptionKey: CollectionSubKey,
	}

	c, err := momo.New(cfg)
	if err != nil {
		log.Fatalf("error starting client: %v", err)
	}

	ctx := context.Background()
	// if listening for callback pass the URL configured in your Momo Account
	callback := os.Getenv("CALLBACK_URL")

	// 1. GetAccountBalance example
	bal, err := c.Collection.GetAccountBalance(ctx)
	if err != nil {
		fmt.Printf("error getting account balance: %v", err)
	}

	fmt.Printf("Account balance: %#v\n", bal)

	// 2. GetAccountBalanceInSpecificCurrency example
	bal, err = c.Collection.GetAccountBalanceInSpecificCurrency(ctx, types.EUR)
	if err != nil {
		fmt.Printf("error getting account balance in specific currency: %v", err)
	}

	fmt.Printf("Account balance in specific currency: %#v\n", bal)

	// 3. CreateInvoice example
	invoiceBody := types.CreateInvoiceInput{
		ExternalID:       uuid.NewString(),
		Amount:           "100",
		Currency:         types.EUR,
		ValidityDuration: "60",
		IntendedPayer: types.Party{
			PartyIDType: types.PARTYCODE,
			PartyID:     uuid.NewString(),
		},
		Payee: types.Party{
			PartyIDType: types.PARTYCODE,
			PartyID:     uuid.NewString(),
		},
		Description: "test invoice",
	}

	invoiceID := uuid.New()

	err = c.Collection.CreateInvoice(ctx, invoiceID, callback, invoiceBody)
	if err != nil {
		fmt.Printf("error getting creating invoice: %v", err)
	}

	// 4. GetInvoiceStatus example
	invoiceStatus, err := c.Collection.GetInvoiceStatus(ctx, invoiceID)
	if err != nil {
		fmt.Printf("error getting invoice status: %v", err)
	}

	fmt.Printf("Invoice status: %#v\n", invoiceStatus)

	// 5. CancelInvoice example. Refer to docs for which transactionID to use.
	err = c.Collection.CancelInvoice(ctx, invoiceID, uuid.New(), callback)
	if err != nil {
		fmt.Printf("error canceling invoice: %v", err)
	}

	// 6. CreatePayments example
	paymentBody := types.PaymentInput{
		ExternalTransactionID: uuid.NewString(),
		Money: types.Money{
			Amount:   "100",
			Currency: types.EUR,
		},
		CustomerReference:       uuid.NewString(),
		ServiceProviderUsername: "sukuna",
		CouponID:                uuid.NewString(),
		ProductID:               uuid.NewString(),
		ProductOfferingID:       uuid.NewString(),
		ReceiverMessage:         "pay now",
		SenderNote:              "rent due",
		MaxNumberOfRetries:      3,
		IncludeSenderCharges:    false,
	}
	paymentID := uuid.New()

	err = c.Collection.CreatePayments(ctx, paymentID, callback, paymentBody)
	if err != nil {
		fmt.Printf("error creating payments: %v", err)
	}

	// 7. GetPaymentStatus example
	paymentStatus, err := c.Collection.GetPaymentStatus(ctx, paymentID)
	if err != nil {
		fmt.Printf("error getting payment status: %v", err)
	}

	fmt.Printf("Payment status: %#v\n", paymentStatus)

	// 8. PreApproval example
	preapprovalInput := types.PreApprovalInput{
		Payer: types.Party{
			PartyIDType: types.PARTYCODE,
			PartyID:     uuid.NewString(),
		},
		PayerCurrency: types.EUR,
		PayerMessage:  "approve payment",
		ValidityTime:  30,
	}
	preapprovalID := uuid.New()

	err = c.Collection.PreApproval(ctx, preapprovalID, callback, preapprovalInput)
	if err != nil {
		fmt.Printf("error getting making preapproval: %v", err)
	}

	// 9. GetPreApprovalStatus example
	preApprovalStatus, err := c.Collection.GetPreApprovalStatus(ctx, preapprovalID)
	if err != nil {
		fmt.Printf("error getting preapproval status: %v", err)
	}

	fmt.Printf("GetPreApprovalStatus: %#v\n", preApprovalStatus)

	// 10. CancelPreApproval example
	err = c.Collection.CancelPreApproval(ctx, preapprovalID)
	if err != nil {
		fmt.Printf("error canceling preapproval: %v", err)
	}

	// 11. GetApprovedApprovals example
	accHolderIDType := ""
	accHolderID := ""

	preApprovalDetails, err := c.Collection.GetApprovedPreApprovals(ctx, accHolderIDType, accHolderID)
	if err != nil {
		fmt.Printf("error getting approved preapproval: %v", err)
	}

	fmt.Printf("GetApprovedPreApprovals %#v\n", preApprovalDetails)

	// 12. RequestToPay example
	reqToPayID := uuid.New()
	reqToPayInput := types.RequestToPayInput{
		Amount:       "100",
		ExternalID:   uuid.NewString(),
		PayerMessage: "test",
		PayeeNote:    "test",
		Currency:     types.EUR,
		Payer: types.Party{
			PartyIDType: types.PARTYCODE,
			PartyID:     uuid.NewString(),
		},
	}

	reqToPayStatus, err := c.Collection.RequestToPay(ctx, reqToPayID, callback, true, reqToPayInput)
	if err != nil {
		fmt.Printf("error requesting payment: %v", err)
	}

	fmt.Printf("RequestToPay status %#v\n", reqToPayStatus)

	// 13. RequestToPayTransactionStatus example
	reqToPayTxStatus, err := c.Collection.RequestToPayTransactionStatus(ctx, reqToPayID)
	if err != nil {
		fmt.Printf("error getting request to pay status: %v", err)
	}

	fmt.Printf("RequestToPayTransactionStatus: %#v\n", reqToPayTxStatus)

	// 14. RequestToPayDeliveryNotification
	deliveryNotification, err := c.Collection.RequestToPayDeliveryNotification(ctx, reqToPayID, "test", "EN")
	if err != nil {
		fmt.Printf("error sending delivery notification: %v", err)
	}

	fmt.Printf("RequestToPayDeliveryNotification: %#v\n", deliveryNotification)

	// 15. RequestToWithdrawv1 example
	reqToWithdrawID := uuid.New()

	err = c.Collection.RequestToWithdrawV1(ctx, reqToWithdrawID, callback, reqToPayInput)
	if err != nil {
		fmt.Printf("error requesting withdrawal(V1): %v", err)
	}

	err = c.Collection.RequestToWithdrawV2(ctx, reqToWithdrawID, callback, reqToPayInput)
	if err != nil {
		fmt.Printf("error requesting withdrawal(v2): %v", err)
	}

	// 16. RequestToWithdrawStatus example
	reqToWithdrawStatus, err := c.Collection.RequestToWithdrawTransactionStatus(ctx, reqToWithdrawID)
	if err != nil {
		fmt.Printf("error requesting withdrawal status: %v", err)
	}

	fmt.Printf("RequestToWithdrawTransactionStatus: %#v\n", reqToWithdrawStatus)

	// 17. CreateAccessToken example
	_, err = c.Collection.CreateAccessToken(ctx)
	if err != nil {
		fmt.Printf("error creating access token: %v", err)
	}

	// 18. CreateOauth2Token example
	_, err = c.Collection.CreateOauth2Token(ctx)
	if err != nil {
		fmt.Printf("error creating oauth2 token: %v", err)
	}

	// 19. BcAuthorize example
	_, err = c.Collection.BcAuthorize(ctx, callback)
	if err != nil {
		fmt.Printf("error doing BcAuthorize: %v", err)
	}

	// 20. GetBasicUserInfo example
	userInfo, err := c.Collection.GetBasicUserInfo(ctx, accHolderIDType, accHolderID)
	if err != nil {
		fmt.Printf("error getting user info: %v", err)
	}

	fmt.Printf("GetBasicUserInfo: %#v\n", userInfo)

	// 21. GetUserInfoWithConsent example
	userInfoConsent, err := c.Collection.GetUserInfoWithConsent(ctx)
	if err != nil {
		fmt.Printf("error getting user infor with consent: %v", err)
	}

	fmt.Printf("GetUserInfoWithConsent: %#v\n", userInfoConsent)
}
