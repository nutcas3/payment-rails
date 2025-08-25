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
