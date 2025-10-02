package api

import (
	"context"
	"fmt"
	"time"
)

type EFTPaymentRequest struct {
	SourceAccountNumber string `json:"sourceAccountNumber"`
	SourceAccountType   string `json:"sourceAccountType,omitempty"`

	BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber"`
	BeneficiaryAccountType   string `json:"beneficiaryAccountType,omitempty"`
	BeneficiaryName          string `json:"beneficiaryName"`
	BeneficiaryBankCode      string `json:"beneficiaryBankCode"`
	BeneficiaryReference     string `json:"beneficiaryReference,omitempty"`

	Amount              float64 `json:"amount"`
	Currency            string  `json:"currency"`
	PaymentReference    string  `json:"paymentReference"`
	PaymentDescription  string  `json:"paymentDescription"`
	PaymentDate         string  `json:"paymentDate,omitempty"`

	NotificationEmail   string `json:"notificationEmail,omitempty"`
	NotificationMobile  string `json:"notificationMobile,omitempty"`
	IdempotencyKey      string `json:"idempotencyKey,omitempty"`
}

type EFTPaymentResponse struct {
	TransactionID       string    `json:"transactionId"`
	Status              string    `json:"status"`
	StatusDescription   string    `json:"statusDescription"`
	PaymentReference    string    `json:"paymentReference"`
	Amount              float64   `json:"amount"`
	Currency            string    `json:"currency"`
	ProcessingDate      time.Time `json:"processingDate"`
	SettlementDate      string    `json:"settlementDate,omitempty"`
	BeneficiaryName     string    `json:"beneficiaryName"`
	Message             string    `json:"message,omitempty"`
}

type UrgentPaymentRequest struct {
	EFTPaymentRequest
	UrgencyReason string `json:"urgencyReason,omitempty"`
}

type PaymentStatusRequest struct {
	TransactionID    string `json:"transactionId,omitempty"`
	PaymentReference string `json:"paymentReference,omitempty"`
}
type PaymentStatusResponse struct {
	TransactionID       string    `json:"transactionId"`
	PaymentReference    string    `json:"paymentReference"`
	Status              string    `json:"status"`
	StatusDescription   string    `json:"statusDescription"`
	Amount              float64   `json:"amount"`
	Currency            string    `json:"currency"`
	SourceAccount       string    `json:"sourceAccount"`
	BeneficiaryAccount  string    `json:"beneficiaryAccount"`
	BeneficiaryName     string    `json:"beneficiaryName"`
	ProcessingDate      time.Time `json:"processingDate"`
	SettlementDate      string    `json:"settlementDate,omitempty"`
	FailureReason       string    `json:"failureReason,omitempty"`
	LastUpdated         time.Time `json:"lastUpdated"`
}

type BatchPaymentRequest struct {
	BatchReference      string              `json:"batchReference"`
	SourceAccountNumber string              `json:"sourceAccountNumber"`
	TotalAmount         float64             `json:"totalAmount"`
	TotalCount          int                 `json:"totalCount"`
	ProcessingDate      string              `json:"processingDate,omitempty"`
	Payments            []EFTPaymentRequest `json:"payments"`
}

type BatchPaymentResponse struct {
	BatchID            string                   `json:"batchId"`
	BatchReference     string                   `json:"batchReference"`
	Status             string                   `json:"status"`
	TotalAmount        float64                  `json:"totalAmount"`
	TotalCount         int                      `json:"totalCount"`
	SuccessCount       int                      `json:"successCount"`
	FailureCount       int                      `json:"failureCount"`
	ProcessingDate     time.Time                `json:"processingDate"`
	PaymentResults     []PaymentResult          `json:"paymentResults,omitempty"`
}

type PaymentResult struct {
	PaymentReference  string  `json:"paymentReference"`
	TransactionID     string  `json:"transactionId,omitempty"`
	Status            string  `json:"status"`
	StatusDescription string  `json:"statusDescription"`
	Amount            float64 `json:"amount"`
	BeneficiaryName   string  `json:"beneficiaryName"`
	FailureReason     string  `json:"failureReason,omitempty"`
}

func (c *Client) CreateEFTPayment(ctx context.Context, req EFTPaymentRequest) (*EFTPaymentResponse, error) {
	if req.Currency == "" {
		req.Currency = "ZAR"
	}
	if req.PaymentDate == "" {
		req.PaymentDate = time.Now().Format("2006-01-02")
	}

	var result EFTPaymentResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/payments/eft", req, &result); err != nil {
		return nil, fmt.Errorf("failed to create EFT payment: %w", err)
	}

	return &result, nil
}

func (c *Client) CreateUrgentPayment(ctx context.Context, req UrgentPaymentRequest) (*EFTPaymentResponse, error) {
	if req.Currency == "" {
		req.Currency = "ZAR"
	}

	var result EFTPaymentResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/payments/urgent", req, &result); err != nil {
		return nil, fmt.Errorf("failed to create urgent payment: %w", err)
	}

	return &result, nil
}

func (c *Client) GetPaymentStatus(ctx context.Context, transactionID string) (*PaymentStatusResponse, error) {
	path := fmt.Sprintf("/api/v1/payments/%s/status", transactionID)
	
	var result PaymentStatusResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get payment status: %w", err)
	}

	return &result, nil
}

func (c *Client) GetPaymentStatusByReference(ctx context.Context, reference string) (*PaymentStatusResponse, error) {
	path := fmt.Sprintf("/api/v1/payments/status?reference=%s", reference)
	
	var result PaymentStatusResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get payment status by reference: %w", err)
	}

	return &result, nil
}

func (c *Client) CreateBatchPayment(ctx context.Context, req BatchPaymentRequest) (*BatchPaymentResponse, error) {
	if req.ProcessingDate == "" {
		req.ProcessingDate = time.Now().Format("2006-01-02")
	}

	var result BatchPaymentResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/payments/batch", req, &result); err != nil {
		return nil, fmt.Errorf("failed to create batch payment: %w", err)
	}

	return &result, nil
}

func (c *Client) GetBatchPaymentStatus(ctx context.Context, batchID string) (*BatchPaymentResponse, error) {
	path := fmt.Sprintf("/api/v1/payments/batch/%s/status", batchID)
	
	var result BatchPaymentResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get batch payment status: %w", err)
	}

	return &result, nil
}

func (c *Client) CancelPayment(ctx context.Context, transactionID string, reason string) error {
	payload := map[string]string{
		"reason": reason,
	}

	path := fmt.Sprintf("/api/v1/payments/%s/cancel", transactionID)
	if err := c.DoRequest(ctx, "POST", path, payload, nil); err != nil {
		return fmt.Errorf("failed to cancel payment: %w", err)
	}

	return nil
}
