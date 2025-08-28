package momo

import (
	"fmt"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/collection"
	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/disbursement"
	"github.com/nutcas3/payment-rails/momo/remittance"
)

// Client with services for collection, disbursement and remittance.
type Client struct {
	Collection   collection.Service   // APIs for remote collection of bills, fees and taxes.
	Disbursement disbursement.Service // APIs to automatically deposit funds to multiple users.
	Remittance   remittance.Service   // APIs to remit funds to local recipients from the diaspora with ease.
}

// ClientConfig are configs needed when creating a Momo client options for the client
type ClientConfig struct {
	// API environment being used i.e. sandbox or production. Default is sandbox.
	Environment string

	// API key assigned to the user. This option is required.
	APIKey string

	// API secret assigned to the user. This option is required.
	APISecret string

	// Subscription key obtained after subscribing to Collection product
	//
	// Used in requests to the /collection/ endpoint. If empty Collection will be nil.
	CollectionSubscriptionKey string

	// Subscription key is the key obtained after subscribing to Disbursement product
	//
	// Used in requests to the /disbursement/ endpoint. If empty Disbursement will be nil.
	DisbursementSubscriptionKey string

	// Subscription key is the key obtained after subscribing to Remittance product
	//
	// Used in requests to the /remittance/ endpoint. If empty Remittance will be nil.
	RemittanceSubscriptionKey string

	// A custom HTTP client
	HTTPClient *http.Client
}

// New creates a Momo client.
func New(cfg ClientConfig) (*Client, error) {
	if cfg.APIKey == "" || cfg.APISecret == "" {
		return nil, fmt.Errorf("APIKey and APISecret are required")
	}

	if cfg.Environment == "" {
		cfg.Environment = "sandbox"
	}

	backend, err := common.NewBackend(&common.BackendConfig{
		Environment: cfg.Environment,
		HTTPClient:  cfg.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	cache := common.NewCache()

	c := &Client{}

	if cfg.CollectionSubscriptionKey != "" {
		c.Collection = collection.NewCollection(
			cfg.CollectionSubscriptionKey,
			cfg.APIKey,
			cfg.APISecret,
			cfg.Environment,
			backend,
			cache,
		)
	}

	if cfg.DisbursementSubscriptionKey != "" {
		c.Disbursement = disbursement.NewDisbursement(
			cfg.DisbursementSubscriptionKey,
			cfg.APIKey,
			cfg.APISecret,
			cfg.Environment,
			backend,
			cache,
		)
	}

	if cfg.RemittanceSubscriptionKey != "" {
		c.Remittance = remittance.NewRemittance(
			cfg.RemittanceSubscriptionKey,
			cfg.APIKey,
			cfg.APISecret,
			cfg.Environment,
			backend,
			cache,
		)
	}

	return c, nil
}
