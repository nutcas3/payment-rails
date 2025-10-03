package fnb

import (
	"payment-rails/fnb/pkg/api"
)


type Client struct {
	*api.Client
}

func NewClient(clientID, clientSecret, apiKey string, opts ...ClientOption) *Client {
	config := &ClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		APIKey:       apiKey,
		Environment:  EnvironmentSandbox,
	}

	for _, opt := range opts {
		opt(config)
	}

	apiConfig := &api.ClientConfig{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		APIKey:       config.APIKey,
		Environment:  string(config.Environment),
		BaseURL:      config.BaseURL,
		H2HConfig:    (*api.H2HConfig)(config.H2HConfig),
	}

	return &Client{
		Client: api.NewClient(apiConfig),
	}
}


type ClientConfig struct {
	ClientID     string
	ClientSecret string
	APIKey       string
	Environment  Environment
	BaseURL      string
	H2HConfig    *H2HConfig
}

type H2HConfig struct {
	CertPath     string
	KeyPath      string
	FTPHost      string
	FTPUser      string
	FTPPassword  string
	FTPDirectory string
}

type Environment string

const (
	EnvironmentSandbox    Environment = "sandbox"
	EnvironmentProduction Environment = "production"
)

type ClientOption func(*ClientConfig)


func WithEnvironment(env Environment) ClientOption {
	return func(c *ClientConfig) {
		c.Environment = env
	}
}

func WithBaseURL(url string) ClientOption {
	return func(c *ClientConfig) {
		c.BaseURL = url
	}
}

func WithH2HConfig(config *H2HConfig) ClientOption {
	return func(c *ClientConfig) {
		c.H2HConfig = config
	}
}
