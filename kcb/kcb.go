package kcb

import (
	"payment-rails/kcb/pkg/api"
)

type Client struct {
	service *api.Service
}

func New(token string, sandbox bool) (*Client, error) {
	environment := api.PRODUCTION
	if sandbox {
		environment = api.SANDBOX
	}

	service, err := api.New(token, environment)
	if err != nil {
		return nil, err
	}

	return &Client{
		service: service,
	}, nil
}

func (c *Client) GetAccountInfo() (*api.AccountInfoResponse, error) {
	return c.service.GetAccountInfo()
}

func (c *Client) GetForexRates(currency string) (*api.ForexRatesResponse, error) {
	return c.service.GetForexRates(currency)
}

func (c *Client) ExchangeCurrency(from, to string, amount float64) (*api.ForexExchangeResponse, error) {
	return c.service.ExchangeCurrency(from, to, amount)
}

func (c *Client) VoomaPay(amount float64) (*api.VoomaPayResponse, error) {
	return c.service.VoomaPay(amount)
}
