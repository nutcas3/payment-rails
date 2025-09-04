package sasapay

import (
	"fmt"
	"os"
	"time"

	"payment-rails/sasapay/pkg/api"
)

// Client is the main SasaPay client that provides access to the SasaPay API
type Client struct {
	apiClient *api.Client
}

// NewClient creates a new SasaPay client
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

// SetWebhookSecret sets the secret used to validate webhook signatures
func (c *Client) SetWebhookSecret(secret string) {
	c.apiClient.SetWebhookSecret(secret)
}

// GenerateReference generates a unique reference for transactions
func GenerateReference() string {
	return fmt.Sprintf("SASAPAY-%d", time.Now().UnixNano())
}

// C2B API Methods

// CustomerToBusiness initiates a payment from a customer to a business
func (c *Client) CustomerToBusiness(req api.C2BRequest) (*api.C2BResponse, error) {
	return c.apiClient.CustomerToBusiness(req)
}

// B2C API Methods

// BusinessToCustomer initiates a payment from a business to a customer
func (c *Client) BusinessToCustomer(req api.B2CRequest) (*api.B2CResponse, error) {
	return c.apiClient.BusinessToCustomer(req)
}

// B2B API Methods

// BusinessToBusiness initiates a payment from one business to another
func (c *Client) BusinessToBusiness(req api.B2BRequest) (*api.B2BResponse, error) {
	return c.apiClient.BusinessToBusiness(req)
}

// Wallet as a Service Methods

// CreateWallet creates a new wallet for a customer
func (c *Client) CreateWallet(req api.CreateWalletRequest) (*api.CreateWalletResponse, error) {
	return c.apiClient.CreateWallet(req)
}

// GetWalletBalance retrieves the balance of a wallet
func (c *Client) GetWalletBalance(req api.WalletBalanceRequest) (*api.WalletBalanceResponse, error) {
	return c.apiClient.GetWalletBalance(req)
}

// TransferToWallet transfers funds from one wallet to another
func (c *Client) TransferToWallet(req api.WalletTransferRequest) (*api.WalletTransferResponse, error) {
	return c.apiClient.TransferToWallet(req)
}

// GetWalletStatement retrieves a statement for a wallet
func (c *Client) GetWalletStatement(req api.WalletStatementRequest) (*api.WalletStatementResponse, error) {
	return c.apiClient.GetWalletStatement(req)
}

// Transaction Status Methods

// CheckTransactionStatus checks the status of a transaction
func (c *Client) CheckTransactionStatus(req api.TransactionStatusRequest) (*api.TransactionStatusResponse, error) {
	return c.apiClient.CheckTransactionStatus(req)
}

// VerifyTransaction verifies a transaction
func (c *Client) VerifyTransaction(req api.VerifyTransactionRequest) (*api.VerifyTransactionResponse, error) {
	return c.apiClient.VerifyTransaction(req)
}

// Cross-region transfers

// CrossRegionTransfer initiates a cross-region money transfer
func (c *Client) CrossRegionTransfer(req api.CrossRegionTransferRequest) (*api.CrossRegionTransferResponse, error) {
	return c.apiClient.CrossRegionTransfer(req)
}

// GetCrossRegionQuote gets a quote for a cross-region transfer
func (c *Client) GetCrossRegionQuote(req api.CrossRegionQuoteRequest) (*api.CrossRegionQuoteResponse, error) {
	return c.apiClient.GetCrossRegionQuote(req)
}

// Webhook handling

// HandleWebhook processes incoming webhook events
func (c *Client) HandleWebhook(payload []byte, signature string, handlers api.WebhookHandlers) error {
	// Webhook handling will be implemented in the webhook.go file
	return fmt.Errorf("webhook handling not yet implemented")
}
