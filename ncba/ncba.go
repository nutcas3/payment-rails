package ncba

import (
	"payment-rails/ncba/pkg/api"
)

type Client struct {
	*api.Client
}

func NewClient(apiKey, username, password string) *Client {
	return &Client{
		Client: api.NewClient(apiKey, username, password),
	}
}
