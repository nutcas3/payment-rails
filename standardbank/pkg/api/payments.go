package api

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type PaymentRequest struct {
	Amount              float64                `json:"amount"`
	Currency            string                 `json:"currency"`
	Reference           string                 `json:"reference"`
	Description         string                 `json:"description,omitempty"`
	SourceAccount       string                 `json:"sourceAccount,omitempty"`
	DestinationAccount  string                 `json:"destinationAccount,omitempty"`
	BeneficiaryName     string                 `json:"beneficiaryName,omitempty"`
	BeneficiaryBankCode string                 `json:"beneficiaryBankCode,omitempty"`
	PaymentDate         string                 `json:"paymentDate,omitempty"`
	IdempotencyKey      string                 `json:"idempotencyKey,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
}

type PaymentResponse struct {
	PaymentID          string                 `json:"paymentId"`
	TransactionID      string                 `json:"transactionId"`
	Status             string                 `json:"status"`
	StatusDescription  string                 `json:"statusDescription,omitempty"`
	Amount             float64                `json:"amount"`
	Currency           string                 `json:"currency"`
	Reference          string                 `json:"reference"`
	Description        string                 `json:"description,omitempty"`
	SourceAccount      string                 `json:"sourceAccount,omitempty"`
	DestinationAccount string                 `json:"destinationAccount,omitempty"`
	ProcessingDate     time.Time              `json:"processingDate"`
	SettlementDate     *time.Time             `json:"settlementDate,omitempty"`
	CreatedAt          time.Time              `json:"createdAt"`
	UpdatedAt          time.Time              `json:"updatedAt"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}

type PaymentStatusResponse struct {
	PaymentID          string                 `json:"paymentId"`
	TransactionID      string                 `json:"transactionId"`
	Status             string                 `json:"status"`
	StatusDescription  string                 `json:"statusDescription,omitempty"`
	Amount             float64                `json:"amount"`
	Currency           string                 `json:"currency"`
	Reference          string                 `json:"reference"`
	SourceAccount      string                 `json:"sourceAccount,omitempty"`
	DestinationAccount string                 `json:"destinationAccount,omitempty"`
	ProcessingDate     time.Time              `json:"processingDate"`
	SettlementDate     *time.Time             `json:"settlementDate,omitempty"`
	FailureReason      string                 `json:"failureReason,omitempty"`
	LastUpdated        time.Time              `json:"lastUpdated"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}

func (c *Client) CreatePayment(ctx context.Context, req PaymentRequest) (*PaymentResponse, error) {
	if req.Currency == "" {
		req.Currency = "ZAR"
	}

	if req.PaymentDate == "" {
		req.PaymentDate = time.Now().Format("2006-01-02")
	}

	var result PaymentResponse
	if err := c.DoRequest(ctx, "POST", "/api/payments", req, &result); err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return &result, nil
}

func (c *Client) GetPayment(ctx context.Context, paymentID string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/api/payments/%s", paymentID)

	var result PaymentResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}

	return &result, nil
}

func (c *Client) GetPaymentStatus(ctx context.Context, paymentID string) (*PaymentStatusResponse, error) {
	path := fmt.Sprintf("/api/payments/%s/status", paymentID)

	var result PaymentStatusResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get payment status: %w", err)
	}

	return &result, nil
}

func (c *Client) GetPaymentByReference(ctx context.Context, reference string) (*PaymentResponse, error) {
	path := fmt.Sprintf("/api/payments?reference=%s", url.QueryEscape(reference))

	var result PaymentResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get payment by reference: %w", err)
	}

	return &result, nil
}
