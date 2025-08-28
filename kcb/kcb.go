package kcb

import (
	"github.com/nutcas3/payment-rails/kcb/pkg/api"
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

func (c *Client) GetAccountBalance(accountNumber string) (*api.AccountBalanceResponse, error) {
	return c.service.GetAccountBalance(accountNumber)
}

func (c *Client) GetAccountStatement(accountNumber, startDate, endDate string) (*api.StatementResponse, error) {
	return c.service.GetAccountStatement(accountNumber, startDate, endDate)
}

func (c *Client) TransferFunds(sourceAccount, destinationAccount string, amount float64, currency, reference, narration string) (*api.TransferResponse, error) {
	return c.service.TransferFunds(sourceAccount, destinationAccount, amount, currency, reference, narration)
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

func (c *Client) CheckVoomaStatus(transactionID string) (*api.VoomaStatusResponse, error) {
	return c.service.CheckVoomaStatus(transactionID)
}

func (c *Client) PesalinkTransfer(sourceAccount, destinationAccount, destinationBank string, amount float64, currency, reference, narration, phoneNumber string) (*api.PesalinkResponse, error) {
	return c.service.PesalinkTransfer(sourceAccount, destinationAccount, destinationBank, amount, currency, reference, narration, phoneNumber)
}

func (c *Client) CheckPesalinkStatus(transactionID string) (*api.PesalinkStatusResponse, error) {
	return c.service.CheckPesalinkStatus(transactionID)
}

func (c *Client) MobileMoneyTransfer(sourceAccount, phoneNumber string, amount float64, currency, reference, narration, provider string) (*api.MobileMoneyResponse, error) {
	return c.service.MobileMoneyTransfer(sourceAccount, phoneNumber, amount, currency, reference, narration, provider)
}

func (c *Client) CheckMobileMoneyStatus(transactionID string) (*api.MobileMoneyStatusResponse, error) {
	return c.service.CheckMobileMoneyStatus(transactionID)
}

func (c *Client) GetUtilityProviders() (*api.UtilityProvidersResponse, error) {
	return c.service.GetUtilityProviders()
}

func (c *Client) PayUtility(sourceAccount, providerID, accountNumber string, amount float64, currency, reference, phoneNumber string) (*api.UtilityPaymentResponse, error) {
	return c.service.PayUtility(sourceAccount, providerID, accountNumber, amount, currency, reference, phoneNumber)
}

func (c *Client) CheckUtilityPaymentStatus(transactionID string) (*api.UtilityStatusResponse, error) {
	return c.service.CheckUtilityPaymentStatus(transactionID)
}
