package stripe

import (
	"github.com/stripe/stripe-go/v82"
	stripeClient "github.com/nutcas3/payment-rails/stripe/pkg/stripe"
)

// Re-export types from the internal package
type (
	CustomerParams       = stripeClient.CustomerParams
	PaymentIntentParams  = stripeClient.PaymentIntentParams
	SubscriptionParams   = stripeClient.SubscriptionParams
	SubscriptionItem     = stripeClient.SubscriptionItem
	InvoiceParams        = stripeClient.InvoiceParams
	InvoiceLineItem      = stripeClient.InvoiceLineItem
	RefundParams         = stripeClient.RefundParams
)

type Environment string

const (
	Production Environment = "production"
	Sandbox Environment = "sandbox"
)

type Config struct {
	APIKey      string
	Environment Environment
	TelemetryEnabled bool
}

type Client struct {
	*stripeClient.Client
}

func NewClient(config Config) *Client {
	stripe.Key = config.APIKey

	if config.TelemetryEnabled {
		stripe.EnableTelemetry = true
	}

	return &Client{
		Client: stripeClient.NewClient(stripeClient.Config{
			APIKey:           config.APIKey,
			Environment:      stripeClient.Environment(config.Environment),
			TelemetryEnabled: config.TelemetryEnabled,
		}),
	}
}

func NewClientWithKey(apiKey string) *Client {
	return NewClient(Config{
		APIKey:      apiKey,
		Environment: Sandbox, // Default to sandbox for safety
	})
}
