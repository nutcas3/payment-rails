package momo

import (
	"net/http"
	"payment-rails/momo/pkg/api"
)

type Environment string

const (
	SANDBOX    Environment = "sandbox"
	PRODUCTION Environment = "production"
)

type Client struct {
	API *api.Client
}

func NewClient(apiUser, apiKey, subscriptionKey string, environment Environment) (*Client, error) {
	apiClient, err := api.New(apiUser, apiKey, subscriptionKey, api.Environment(environment))
	if err != nil {
		return nil, err
	}

	return &Client{
		API: apiClient,
	}, nil
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.API.SetHttpClient(httpClient)
}
func (c *Client) GetCollectionToken() (string, error) {
	return c.API.GetCollectionToken()
}

func (c *Client) GetDisbursementToken() (string, error) {
	return c.API.GetDisbursementToken()
}
func (c *Client) GetRemittanceToken() (string, error) {
	return c.API.GetRemittanceToken()
}
