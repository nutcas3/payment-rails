package api

import (
	"context"
	"fmt"
)

type Provider struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Description string            `json:"description,omitempty"`
	Status      string            `json:"status"`
	Currency    string            `json:"currency,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type ProviderListResponse struct {
	Providers []Provider `json:"providers"`
	Total     int        `json:"total"`
	Page      int        `json:"page,omitempty"`
	PageSize  int        `json:"pageSize,omitempty"`
}

type ProviderPaymentRequest struct {
	ProviderID     string                 `json:"providerId"`
	Amount         float64                `json:"amount"`
	Currency       string                 `json:"currency"`
	Reference      string                 `json:"reference"`
	Description    string                 `json:"description,omitempty"`
	SourceAccount  string                 `json:"sourceAccount,omitempty"`
	Destination    map[string]interface{} `json:"destination"`
	IdempotencyKey string                 `json:"idempotencyKey,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

type ProviderPaymentResponse struct {
	PaymentID         string                 `json:"paymentId"`
	TransactionID     string                 `json:"transactionId"`
	ProviderID        string                 `json:"providerId"`
	Status            string                 `json:"status"`
	StatusDescription string                 `json:"statusDescription,omitempty"`
	Amount            float64                `json:"amount"`
	Currency          string                 `json:"currency"`
	Reference         string                 `json:"reference"`
	ProviderReference string                 `json:"providerReference,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
}

func (c *Client) GetProviders(ctx context.Context) (*ProviderListResponse, error) {
	var result ProviderListResponse
	if err := c.DoRequest(ctx, "GET", "/api/providers", nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get providers: %w", err)
	}

	return &result, nil
}

func (c *Client) GetProvider(ctx context.Context, providerID string) (*Provider, error) {
	path := fmt.Sprintf("/api/providers/%s", providerID)

	var result Provider
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	return &result, nil
}

func (c *Client) ExecuteProviderPayment(ctx context.Context, req ProviderPaymentRequest) (*ProviderPaymentResponse, error) {
	if req.Currency == "" {
		req.Currency = "ZAR"
	}

	path := fmt.Sprintf("/api/providers/%s/pay", req.ProviderID)

	var result ProviderPaymentResponse
	if err := c.DoRequest(ctx, "POST", path, req, &result); err != nil {
		return nil, fmt.Errorf("failed to execute provider payment: %w", err)
	}

	return &result, nil
}
