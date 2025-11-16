package standardbank

import (
	"context"
	"fmt"
	"net/http"
	"payment-rails/standardbank/pkg/api"
	"time"
)

type Client struct {
	apiClient      *api.Client
	webhookHandler *api.WebhookHandler
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	APIKey       string
	Environment  Environment
	BaseURL      string
	Timeout      int // Timeout in seconds
	Logger       api.Logger
}

type Environment string

const (
	EnvironmentSandbox    Environment = "sandbox"
	EnvironmentProduction Environment = "production"
)

type ClientOption func(*ClientConfig)

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
		Logger:       config.Logger,
	}

	if config.Timeout > 0 {
		apiConfig.Timeout = time.Duration(config.Timeout) * time.Second
	}

	return &Client{
		apiClient: api.NewClient(apiConfig),
	}
}

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

func WithTimeout(seconds int) ClientOption {
	return func(c *ClientConfig) {
		c.Timeout = seconds
	}
}

func WithLogger(logger api.Logger) ClientOption {
	return func(c *ClientConfig) {
		c.Logger = logger
	}
}

func (c *Client) SetWebhookSecret(webhookSecret string) {
	c.webhookHandler = api.NewWebhookHandler(webhookSecret)
}

func (c *Client) HandleWebhook(w http.ResponseWriter, r *http.Request) error {
	if c.webhookHandler == nil {
		return fmt.Errorf("webhook handler not initialized, call SetWebhookSecret first")
	}
	return c.webhookHandler.HandleWebhook(w, r)
}

type Payments struct {
	client *Client
}

func (c *Client) Payments() *Payments {
	return &Payments{client: c}
}

func (p *Payments) Create(ctx context.Context, req api.PaymentRequest) (*api.PaymentResponse, error) {
	return p.client.apiClient.CreatePayment(ctx, req)
}

func (p *Payments) Get(ctx context.Context, paymentID string) (*api.PaymentResponse, error) {
	return p.client.apiClient.GetPayment(ctx, paymentID)
}

func (p *Payments) GetStatus(ctx context.Context, paymentID string) (*api.PaymentStatusResponse, error) {
	return p.client.apiClient.GetPaymentStatus(ctx, paymentID)
}

func (p *Payments) GetByReference(ctx context.Context, reference string) (*api.PaymentResponse, error) {
	return p.client.apiClient.GetPaymentByReference(ctx, reference)
}

type Transfers struct {
	client *Client
}

func (c *Client) Transfers() *Transfers {
	return &Transfers{client: c}
}

func (t *Transfers) Create(ctx context.Context, req api.InternalTransferRequest) (*api.InternalTransferResponse, error) {
	return t.client.apiClient.CreateInternalTransfer(ctx, req)
}

func (t *Transfers) Get(ctx context.Context, transferID string) (*api.InternalTransferResponse, error) {
	return t.client.apiClient.GetTransfer(ctx, transferID)
}

func (t *Transfers) GetStatus(ctx context.Context, transferID string) (*api.TransferStatusResponse, error) {
	return t.client.apiClient.GetTransferStatus(ctx, transferID)
}

type Providers struct {
	client *Client
}

func (c *Client) Providers() *Providers {
	return &Providers{client: c}
}

func (p *Providers) List(ctx context.Context) (*api.ProviderListResponse, error) {
	return p.client.apiClient.GetProviders(ctx)
}

func (p *Providers) Get(ctx context.Context, providerID string) (*api.Provider, error) {
	return p.client.apiClient.GetProvider(ctx, providerID)
}

func (p *Providers) Pay(ctx context.Context, req api.ProviderPaymentRequest) (*api.ProviderPaymentResponse, error) {
	return p.client.apiClient.ExecuteProviderPayment(ctx, req)
}

func (c *Client) RegisterWebhookHandler(eventType string, handler api.WebhookEventHandler) error {
	if c.webhookHandler == nil {
		return fmt.Errorf("webhook handler not initialized, call SetWebhookSecret first")
	}
	c.webhookHandler.RegisterHandler(eventType, handler)
	return nil
}
