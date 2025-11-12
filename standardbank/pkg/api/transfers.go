package api

import (
	"context"
	"fmt"
	"time"
)

type InternalTransferRequest struct {
	SourceAccount      string  `json:"sourceAccount"`
	DestinationAccount string  `json:"destinationAccount"`
	Amount             float64 `json:"amount"`
	Currency           string  `json:"currency"`
	Reference          string  `json:"reference"`
	Description        string  `json:"description,omitempty"`
	DestinationName    string  `json:"destinationName,omitempty"`
	TransferDate       string  `json:"transferDate,omitempty"`
	IdempotencyKey     string  `json:"idempotencyKey,omitempty"`
}

type InternalTransferResponse struct {
	TransferID         string     `json:"transferId"`
	TransactionID      string     `json:"transactionId"`
	Status             string     `json:"status"`
	StatusDescription  string     `json:"statusDescription,omitempty"`
	SourceAccount      string     `json:"sourceAccount"`
	DestinationAccount string     `json:"destinationAccount"`
	Amount             float64    `json:"amount"`
	Currency           string     `json:"currency"`
	Reference          string     `json:"reference"`
	Description        string     `json:"description,omitempty"`
	ProcessingDate     time.Time  `json:"processingDate"`
	SettlementDate     *time.Time `json:"settlementDate,omitempty"`
	CreatedAt          time.Time  `json:"createdAt"`
}

type TransferStatusResponse struct {
	TransferID         string     `json:"transferId"`
	TransactionID      string     `json:"transactionId"`
	Status             string     `json:"status"`
	StatusDescription  string     `json:"statusDescription,omitempty"`
	SourceAccount      string     `json:"sourceAccount"`
	DestinationAccount string     `json:"destinationAccount"`
	Amount             float64    `json:"amount"`
	Currency           string     `json:"currency"`
	Reference          string     `json:"reference"`
	ProcessingDate     time.Time  `json:"processingDate"`
	SettlementDate     *time.Time `json:"settlementDate,omitempty"`
	FailureReason      string     `json:"failureReason,omitempty"`
	LastUpdated        time.Time  `json:"lastUpdated"`
}

func (c *Client) CreateInternalTransfer(ctx context.Context, req InternalTransferRequest) (*InternalTransferResponse, error) {
	if req.Currency == "" {
		req.Currency = "ZAR"
	}

	if req.TransferDate == "" {
		req.TransferDate = time.Now().Format("2006-01-02")
	}

	var result InternalTransferResponse
	if err := c.DoRequest(ctx, "POST", "/api/transfers", req, &result); err != nil {
		return nil, fmt.Errorf("failed to create internal transfer: %w", err)
	}

	return &result, nil
}

func (c *Client) GetTransfer(ctx context.Context, transferID string) (*InternalTransferResponse, error) {
	path := fmt.Sprintf("/api/transfers/%s", transferID)

	var result InternalTransferResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get transfer: %w", err)
	}

	return &result, nil
}

func (c *Client) GetTransferStatus(ctx context.Context, transferID string) (*TransferStatusResponse, error) {
	path := fmt.Sprintf("/api/transfers/%s/status", transferID)

	var result TransferStatusResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get transfer status: %w", err)
	}

	return &result, nil
}
