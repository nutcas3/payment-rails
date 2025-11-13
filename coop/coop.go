package coop

import (
	"fmt"
	"os"
	"time"

	"github.com/nutcas3/payment-rails/coop/pkg/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(clientID, clientSecret string, environment api.Environment) (*Client, error) {
	if clientID == "" {
		clientID = os.Getenv("COOP_CLIENT_ID")
	}

	if clientSecret == "" {
		clientSecret = os.Getenv("COOP_CLIENT_SECRET")
	}

	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("missing required credentials: client ID and client secret must be provided")
	}

	if environment == "" {
		environment = api.SANDBOX
	}

	apiClient, err := api.NewClient(clientID, clientSecret, environment)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiClient: apiClient,
	}, nil
}

func GenerateReference() string {
	return fmt.Sprintf("COOP-%d", time.Now().UnixNano())
}

func (c *Client) AccountBalance(accountNumber string) (*api.AccountBalanceResponse, error) {
	req := api.AccountBalanceRequest{
		BaseRequest: api.BaseRequest{
			MessageReference: GenerateReference(),
		},
		AccountNumber: accountNumber,
	}

	return c.apiClient.AccountBalance(req)
}

func (c *Client) AccountTransactions(accountNumber string, noOfTransactions string) (*api.AccountTransactionsResponse, error) {
	if noOfTransactions == "" {
		noOfTransactions = "10" // Default to 10 transactions
	}

	req := api.AccountTransactionsRequest{
		BaseRequest: api.BaseRequest{
			MessageReference: GenerateReference(),
		},
		AccountNumber:    accountNumber,
		NoOfTransactions: noOfTransactions,
	}

	return c.apiClient.AccountTransactions(req)
}

func (c *Client) ExchangeRate(fromCurrency, toCurrency string) (*api.ExchangeRateResponse, error) {
	if fromCurrency == "" {
		fromCurrency = "KES"
	}
	if toCurrency == "" {
		toCurrency = "USD"
	}

	req := api.ExchangeRateRequest{
		BaseRequest: api.BaseRequest{
			MessageReference: GenerateReference(),
		},
		FromCurrencyCode: fromCurrency,
		ToCurrencyCode:   toCurrency,
	}

	return c.apiClient.ExchangeRate(req)
}

func (c *Client) InternalFundsTransfer(sourceAccount string, amount float64, currency, narration string, destinations []api.Destination) (*api.IFTResponse, error) {
	if currency == "" {
		currency = "KES"
	}
	if narration == "" {
		narration = "Internal Transfer"
	}

	req := api.IFTRequest{
		BaseRequest: api.BaseRequest{
			MessageReference: GenerateReference(),
		},
		AccountNumber:       sourceAccount,
		Amount:              amount,
		TransactionCurrency: currency,
		Narration:           narration,
		Destinations:        destinations,
	}

	return c.apiClient.InternalFundsTransfer(req)
}

func (c *Client) PesaLinkTransfer(sourceAccount string, amount float64, currency, narration string, destinations []api.PesaLinkDestination) (*api.PesaLinkResponse, error) {
	if currency == "" {
		currency = "KES"
	}
	if narration == "" {
		narration = "PesaLink Transfer"
	}

	req := api.PesaLinkRequest{
		BaseRequest: api.BaseRequest{
			MessageReference: GenerateReference(),
		},
		AccountNumber:       sourceAccount,
		Amount:              amount,
		TransactionCurrency: currency,
		Narration:           narration,
		Destinations:        destinations,
	}

	return c.apiClient.PesaLinkSendToAccount(req)
}

func (c *Client) TransactionStatus(messageReference string) (*api.TransactionStatusResponse, error) {
	req := api.TransactionStatusRequest{
		BaseRequest: api.BaseRequest{
			MessageReference: messageReference,
		},
	}

	return c.apiClient.TransactionStatus(req)
}
