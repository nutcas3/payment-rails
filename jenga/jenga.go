package jenga

import (
	"payment-rails/jenga/pkg/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(apiKey, username, password, privateKey, environment string) (*Client, error) {
	apiClient, err := api.NewClient(apiKey, username, password, privateKey, environment)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiClient: apiClient,
	}, nil
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

func (c *Client) PurchaseAirtime(req api.AirtimePurchaseRequest) (*api.AirtimePurchaseResponse, error) {
	return c.apiClient.PurchaseAirtime(req)
}


func (c *Client) VerifyIdentity(req api.KYCRequest) (*api.KYCResponse, error) {
	return c.apiClient.VerifyIdentity(req)
}

func (c *Client) GetForexRates(req api.ForexRatesRequest) (*api.ForexRatesResponse, error) {
	return c.apiClient.GetForexRates(req)
}

// GenerateReference generates a unique reference ID for transactions
func GenerateReference() string {
	return api.GenerateReference()
}
