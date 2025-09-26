package stripe

import (
	"github.com/stripe/stripe-go/v82"
)

type Client struct {
	*stripe.Client
}
func NewClient(config Config) *Client {

	stripe.Key = config.APIKey

	if config.TelemetryEnabled {
		stripe.EnableTelemetry = true
	}

	return &Client{
		Client: stripe.NewClient(config.APIKey),
	}
}

func NewClientWithKey(apiKey string) *Client {
	return NewClient(Config{
		APIKey:      apiKey,
		Environment: Sandbox,
	})
}

type Config struct {
	APIKey      string
	Environment Environment
	TelemetryEnabled bool
}
type Environment string

const (
	Production Environment = "production"
	Sandbox Environment = "sandbox"
)
