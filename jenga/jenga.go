package jenga

import (
	"fmt"
	"net/http"
	"payment-rails/jenga/pkg/api"
)

type Client struct {
	apiClient *api.Client
	webhookHandler *api.WebhookHandler
}

func NewClient(apiKey, username, password, privateKey, environment string) (*Client, error) {
	apiClient, err := api.NewClient(apiKey, username, password, privateKey, environment)
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

func (c *Client) GetAccountBalance(req api.AccountBalanceRequest) (*api.AccountBalanceResponse, error) {
	return c.apiClient.GetAccountBalance(req)
}

func (c *Client) GetMiniStatement(req api.MiniStatementRequest) (*api.MiniStatementResponse, error) {
	return c.apiClient.GetMiniStatement(req)
}

func (c *Client) GetFullStatement(req api.FullStatementRequest) (*api.FullStatementResponse, error) {
	return c.apiClient.GetFullStatement(req)
}

func (c *Client) ValidateAccount(req api.AccountValidateRequest) (*api.AccountValidateResponse, error) {
	return c.apiClient.ValidateAccount(req)
}


func (c *Client) SendMoney(req api.SendMoneyRequest) (*api.SendMoneyResponse, error) {
	return c.apiClient.SendMoney(req)
}

func (c *Client) SendToMobileWallet(req api.MobileWalletRequest) (*api.MobileWalletResponse, error) {
	return c.apiClient.SendToMobileWallet(req)
}

func (c *Client) SendInternalBankTransfer(req api.SendMoneyRequest) (*api.SendMoneyResponse, error) {
	return c.apiClient.SendInternalBankTransfer(req)
}


func (c *Client) PayBill(req api.BillPaymentRequest) (*api.BillPaymentResponse, error) {
	return c.apiClient.PayBill(req)
}

func (c *Client) ReceiveMoney(req api.ReceiveMoneyRequest) (*api.ReceiveMoneyResponse, error) {
	return c.apiClient.ReceiveMoney(req)
}

func (c *Client) QueryReceiveMoneyTransaction(req api.ReceiveMoneyQueryRequest) (*api.ReceiveMoneyQueryResponse, error) {
	return c.apiClient.QueryReceiveMoneyTransaction(req)
}

func (c *Client) PurchaseAirtime(req api.AirtimePurchaseRequest) (*api.AirtimePurchaseResponse, error) {
	return c.apiClient.PurchaseAirtime(req)
}


func (c *Client) VerifyIdentity(req api.KYCRequest) (*api.KYCResponse, error) {
	return c.apiClient.VerifyIdentity(req)
}

func (c *Client) PerformAMLScreening(req api.AMLScreeningRequest) (*api.AMLScreeningResponse, error) {
	return c.apiClient.PerformAMLScreening(req)
}

func (c *Client) PerformCustomerDueDiligence(req api.CDDRequest) (*api.CDDResponse, error) {
	return c.apiClient.PerformCustomerDueDiligence(req)
}

func (c *Client) GetForexRates(req api.ForexRatesRequest) (*api.ForexRatesResponse, error) {
	return c.apiClient.GetForexRates(req)
}

// GenerateReference generates a unique reference ID for transactions
func GenerateReference() string {
	return api.GenerateReference()
}
