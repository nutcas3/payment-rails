package absa

import (
	"fmt"
	"net/http"
	"payment-rails/absa/pkg/api"
)

type Client struct {
	apiClient *api.Client
	webhookHandler *api.WebhookHandler
}

func NewClient(clientID, clientSecret, apiKey, environment string) (*Client, error) {
	apiClient, err := api.NewClient(clientID, clientSecret, apiKey, environment)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiClient: apiClient,
		webhookHandler: nil,
	}, nil
}

// SetWebhookSecret sets the webhook secret for validating webhook signatures
func (c *Client) SetWebhookSecret(webhookSecret string) {
	c.webhookHandler = api.NewWebhookHandler(webhookSecret)
}

// HandleWebhook processes incoming webhook requests
func (c *Client) HandleWebhook(w http.ResponseWriter, r *http.Request, handlers api.WebhookHandlers) error {
	if c.webhookHandler == nil {
		return fmt.Errorf("webhook handler not initialized, call SetWebhookSecret first")
	}
	c.webhookHandler.HandleWebhook(w, r, handlers)
	return nil
}

// GetAccountBalance retrieves the balance for a specified account
func (c *Client) GetAccountBalance(req api.AccountBalanceRequest) (*api.AccountBalanceResponse, error) {
	return c.apiClient.GetAccountBalance(req)
}

// GetMiniStatement retrieves a mini statement for a specified account
func (c *Client) GetMiniStatement(req api.MiniStatementRequest) (*api.MiniStatementResponse, error) {
	return c.apiClient.GetMiniStatement(req)
}

// GetFullStatement retrieves a full statement for a specified account
func (c *Client) GetFullStatement(req api.FullStatementRequest) (*api.FullStatementResponse, error) {
	return c.apiClient.GetFullStatement(req)
}

// ValidateAccount validates if an account exists and is active
func (c *Client) ValidateAccount(req api.AccountValidateRequest) (*api.AccountValidateResponse, error) {
	return c.apiClient.ValidateAccount(req)
}

// SendMoney transfers funds to another bank account
func (c *Client) SendMoney(req api.SendMoneyRequest) (*api.SendMoneyResponse, error) {
	return c.apiClient.SendMoney(req)
}

// SendToMobileWallet transfers funds to a mobile wallet
func (c *Client) SendToMobileWallet(req api.MobileWalletRequest) (*api.MobileWalletResponse, error) {
	return c.apiClient.SendToMobileWallet(req)
}

// SendInternalBankTransfer transfers funds within the same bank
func (c *Client) SendInternalBankTransfer(req api.SendMoneyRequest) (*api.SendMoneyResponse, error) {
	return c.apiClient.SendInternalBankTransfer(req)
}

// PayBill pays a bill to a specified biller
func (c *Client) PayBill(req api.BillPaymentRequest) (*api.BillPaymentResponse, error) {
	return c.apiClient.PayBill(req)
}

// ReceiveMoney initiates a request to receive money
func (c *Client) ReceiveMoney(req api.ReceiveMoneyRequest) (*api.ReceiveMoneyResponse, error) {
	return c.apiClient.ReceiveMoney(req)
}

// QueryTransaction checks the status of a transaction
func (c *Client) QueryTransaction(req api.TransactionQueryRequest) (*api.TransactionQueryResponse, error) {
	return c.apiClient.QueryTransaction(req)
}

// PurchaseAirtime buys airtime for a mobile number
func (c *Client) PurchaseAirtime(req api.AirtimePurchaseRequest) (*api.AirtimePurchaseResponse, error) {
	return c.apiClient.PurchaseAirtime(req)
}

// GenerateReference generates a unique reference ID for transactions
func GenerateReference() string {
	return api.GenerateReference()
}

// Bulk Payments

// ProcessBulkPayment processes multiple payments in a single batch
func (c *Client) ProcessBulkPayment(req api.BulkPaymentRequest) (*api.BulkPaymentResponse, error) {
	return c.apiClient.ProcessBulkPayment(req)
}

// GetBulkPaymentStatus retrieves the status of a bulk payment batch
func (c *Client) GetBulkPaymentStatus(req api.BulkPaymentStatusRequest) (*api.BulkPaymentStatusResponse, error) {
	return c.apiClient.GetBulkPaymentStatus(req)
}

// Standing Orders

// CreateStandingOrder creates a new recurring payment
func (c *Client) CreateStandingOrder(req api.StandingOrderRequest) (*api.StandingOrderResponse, error) {
	return c.apiClient.CreateStandingOrder(req)
}

// GetStandingOrderStatus retrieves the status of a standing order
func (c *Client) GetStandingOrderStatus(req api.StandingOrderStatusRequest) (*api.StandingOrderStatusResponse, error) {
	return c.apiClient.GetStandingOrderStatus(req)
}

// CancelStandingOrder cancels an existing standing order
func (c *Client) CancelStandingOrder(req api.StandingOrderCancelRequest) (*api.StandingOrderCancelResponse, error) {
	return c.apiClient.CancelStandingOrder(req)
}

// ListStandingOrders retrieves all standing orders for an account
func (c *Client) ListStandingOrders(req api.StandingOrderListRequest) (*api.StandingOrderListResponse, error) {
	return c.apiClient.ListStandingOrders(req)
}

// Beneficiary Management

// CreateBeneficiary creates a new beneficiary
func (c *Client) CreateBeneficiary(req api.BeneficiaryCreateRequest) (*api.BeneficiaryCreateResponse, error) {
	return c.apiClient.CreateBeneficiary(req)
}

// ListBeneficiaries retrieves all beneficiaries
func (c *Client) ListBeneficiaries(req api.BeneficiaryListRequest) (*api.BeneficiaryListResponse, error) {
	return c.apiClient.ListBeneficiaries(req)
}

// GetBeneficiary retrieves a specific beneficiary by ID
func (c *Client) GetBeneficiary(req api.BeneficiaryGetRequest) (*api.BeneficiaryGetResponse, error) {
	return c.apiClient.GetBeneficiary(req)
}

// UpdateBeneficiary updates an existing beneficiary
func (c *Client) UpdateBeneficiary(req api.BeneficiaryUpdateRequest) (*api.BeneficiaryUpdateResponse, error) {
	return c.apiClient.UpdateBeneficiary(req)
}

// DeleteBeneficiary deletes a beneficiary
func (c *Client) DeleteBeneficiary(req api.BeneficiaryDeleteRequest) (*api.BeneficiaryDeleteResponse, error) {
	return c.apiClient.DeleteBeneficiary(req)
}

// Foreign Exchange

// GetForexRate retrieves the exchange rate between two currencies
func (c *Client) GetForexRate(req api.ForexRateRequest) (*api.ForexRateResponse, error) {
	return c.apiClient.GetForexRate(req)
}

// ProcessForexTransfer initiates a foreign currency transfer
func (c *Client) ProcessForexTransfer(req api.ForexTransferRequest) (*api.ForexTransferResponse, error) {
	return c.apiClient.ProcessForexTransfer(req)
}

// GetForexTransferStatus retrieves the status of a forex transfer
func (c *Client) GetForexTransferStatus(req api.ForexTransferStatusRequest) (*api.ForexTransferStatusResponse, error) {
	return c.apiClient.GetForexTransferStatus(req)
}

// Authentication Methods

// RequestOTP requests a one-time password to be sent to a phone number or email
func (c *Client) RequestOTP(req api.OTPRequest) (*api.OTPResponse, error) {
	return c.apiClient.RequestOTP(req)
}

// VerifyOTP verifies a one-time password
func (c *Client) VerifyOTP(req api.OTPVerifyRequest) (*api.OTPVerifyResponse, error) {
	return c.apiClient.VerifyOTP(req)
}

// AuthenticateTransaction authenticates a transaction using the specified method
func (c *Client) AuthenticateTransaction(req api.TransactionAuthRequest) (*api.TransactionAuthResponse, error) {
	return c.apiClient.AuthenticateTransaction(req)
}

// RegisterDevice registers a new device for authentication
func (c *Client) RegisterDevice(req api.DeviceRegistrationRequest) (*api.DeviceRegistrationResponse, error) {
	return c.apiClient.RegisterDevice(req)
}
