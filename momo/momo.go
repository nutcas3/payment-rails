package momo

import (
	"net/http"

	"github.com/nutcas3/payment-rails/momo/collection"
	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/disbursement"
	"github.com/nutcas3/payment-rails/momo/remittance"
)

type Environment string

const (
	SANDBOX    Environment = "sandbox"
	PRODUCTION Environment = "production"
)

// ClientConfig holds the configuration for creating a new Momo client
type ClientConfig struct {
	Environment                 string
	APIKey                      string
	APISecret                   string
	CollectionSubscriptionKey   string
	DisbursementSubscriptionKey string
	RemittanceSubscriptionKey   string
	HTTPClient                  *http.Client
}

type Client struct {
	Collection   collection.Service
	Disbursement disbursement.Service
	Remittance   remittance.Service
	backend      common.Backend
	cache        common.CacheStore
}

// New creates a new Momo client with the given configuration
func New(cfg ClientConfig) (*Client, error) {
	backendCfg := &common.BackendConfig{
		Environment: cfg.Environment,
		HTTPClient:  cfg.HTTPClient,
	}

	backend, err := common.NewBackend(backendCfg)
	if err != nil {
		return nil, err
	}

	cache := common.NewCache()

	client := &Client{
		backend: backend,
		cache:   cache,
	}

	// Initialize Collection service if subscription key is provided
	if cfg.CollectionSubscriptionKey != "" {
		client.Collection = collection.NewCollection(
			cfg.CollectionSubscriptionKey,
			cfg.APIKey,
			cfg.APISecret,
			cfg.Environment,
			backend,
			cache,
		)
	}

	// Initialize Disbursement service if subscription key is provided
	if cfg.DisbursementSubscriptionKey != "" {
		client.Disbursement = disbursement.NewDisbursement(
			cfg.DisbursementSubscriptionKey,
			cfg.APIKey,
			cfg.APISecret,
			cfg.Environment,
			backend,
			cache,
		)
	}

	// Initialize Remittance service if subscription key is provided
	if cfg.RemittanceSubscriptionKey != "" {
		client.Remittance = remittance.NewRemittance(
			cfg.RemittanceSubscriptionKey,
			cfg.APIKey,
			cfg.APISecret,
			cfg.Environment,
			backend,
			cache,
		)
	}

	return client, nil
}
