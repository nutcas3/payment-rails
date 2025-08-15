package airtel

import (
	"payment-rails/airtel/pkg/api"
)

type Client struct {
	service *api.Service
}

func New(clientID, clientSecret, publicKey string, sandbox bool, country, currency string) (*Client, error) {
	environment := api.PRODUCTION
	if sandbox {
		environment = api.SANDBOX
	}

	service, err := api.New(clientID, clientSecret, publicKey, environment, country, currency)
	if err != nil {
		return nil, err
	}

	return &Client{
		service: service,
	}, nil
}

func (c *Client) UssdPush(reference, phone string, amount float64, transactionID string) (*api.CollectionResponse, error) {
	return c.service.UssdPush(reference, phone, amount, transactionID)
}

func (c *Client) GetTransactionStatus(transactionID string) (*api.TransactionStatusResponse, error) {
	return c.service.GetTransactionStatus(transactionID)
}

func (c *Client) RefundTransaction(airtelMoneyID string, amount float64) (*api.RefundResponse, error) {
	return c.service.RefundTransaction(airtelMoneyID, amount)
}

func (c *Client) Disburse(reference, phone string, amount float64, transactionID string, pin string) (*api.DisbursementResponse, error) {
	return c.service.Disburse(reference, phone, amount, transactionID, pin)
}

func (c *Client) GetDisbursementStatus(transactionID string) (*api.DisbursementStatusResponse, error) {
	return c.service.GetDisbursementStatus(transactionID)
}

func (c *Client) GetAccountBalance() (*api.AccountBalanceResponse, error) {
	return c.service.GetAccountBalance()
}
