package sasapay

import (
	"fmt"
	"os"
	"time"

	"github.com/nutcas3/payment-rails/sasapay/pkg/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(clientID, clientSecret, environment string) (*Client, error) {
	if clientID == "" {
		clientID = os.Getenv("SASAPAY_CLIENT_ID")
	}

	if clientSecret == "" {
		clientSecret = os.Getenv("SASAPAY_CLIENT_SECRET")
	}

	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("missing required credentials: client ID and client secret must be provided")
	}

	if environment == "" {
		environment = "sandbox"
	}

	if environment != "sandbox" && environment != "production" {
		return nil, fmt.Errorf("invalid environment: must be 'sandbox' or 'production'")
	}

	apiClient, err := api.NewClient(clientID, clientSecret, environment)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiClient: apiClient,
	}, nil
}

func (c *Client) SetWebhookSecret(secret string) {
	c.apiClient.SetWebhookSecret(secret)
}

func GenerateReference() string {
	return fmt.Sprintf("SASAPAY-%d", time.Now().UnixNano())
}


func (c *Client) CustomerToBusiness(req api.C2BRequest) (*api.C2BResponse, error) {
	return c.apiClient.CustomerToBusiness(req)
}


func (c *Client) BusinessToCustomer(req api.B2CRequest) (*api.B2CResponse, error) {
	return c.apiClient.BusinessToCustomer(req)
}


func (c *Client) BusinessToBusiness(req api.B2BRequest) (*api.B2BResponse, error) {
	return c.apiClient.BusinessToBusiness(req)
}


func (c *Client) CreateWallet(req api.CreateWalletRequest) (*api.CreateWalletResponse, error) {
	return c.apiClient.CreateWallet(req)
}

func (c *Client) GetWalletBalance(req api.WalletBalanceRequest) (*api.WalletBalanceResponse, error) {
	return c.apiClient.GetWalletBalance(req)
}

func (c *Client) TransferToWallet(req api.WalletTransferRequest) (*api.WalletTransferResponse, error) {
	return c.apiClient.TransferToWallet(req)
}

func (c *Client) GetWalletStatement(req api.WalletStatementRequest) (*api.WalletStatementResponse, error) {
	return c.apiClient.GetWalletStatement(req)
}


func (c *Client) CheckTransactionStatus(req api.TransactionStatusRequest) (*api.TransactionStatusResponse, error) {
	return c.apiClient.CheckTransactionStatus(req)
}

func (c *Client) VerifyTransaction(req api.VerifyTransactionRequest) (*api.VerifyTransactionResponse, error) {
	return c.apiClient.VerifyTransaction(req)
}


func (c *Client) CrossRegionTransfer(req api.CrossRegionTransferRequest) (*api.CrossRegionTransferResponse, error) {
	return c.apiClient.CrossRegionTransfer(req)
}

func (c *Client) GetCrossRegionQuote(req api.CrossRegionQuoteRequest) (*api.CrossRegionQuoteResponse, error) {
	return c.apiClient.GetCrossRegionQuote(req)
}


func (c *Client) HandleWebhook(payload []byte, signature string, handlers api.WebhookHandlers) error {
	return fmt.Errorf("webhook handling not yet implemented")
}
