package mpesa

import (
	"net/http"
	"github.com/nutcas3/payment-rails/mpesa/pkg/daraja"
)

type Environment string

const (
	SANDBOX Environment = "sandbox"
	PRODUCTION Environment = "production"
)

type Client struct {
	Service *daraja.Service
}

func NewClient(apiKey, consumerSecret, passKey string, environment Environment) (*Client, error) {
	service, err := daraja.New(apiKey, consumerSecret, passKey, daraja.Environment(environment))
	if err != nil {
		return nil, err
	}

	return &Client{
		Service: service,
	}, nil
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.Service.SetHttpClient(httpClient)
}

func (c *Client) GetAuthToken() (string, error) {
	return c.Service.GetAuthToken()
}
