package api

import (
	"context"
	"fmt"
	"time"
)

type EFTCollectionRequest struct {
	CreditorAccountNumber string `json:"creditorAccountNumber"`
	CreditorAccountType   string `json:"creditorAccountType,omitempty"`
	CreditorName          string `json:"creditorName"`
	DebtorAccountNumber string `json:"debtorAccountNumber"`
	DebtorAccountType   string `json:"debtorAccountType,omitempty"`
	DebtorName          string `json:"debtorName"`
	DebtorBankCode      string `json:"debtorBankCode"` // Universal branch code
	Amount              float64 `json:"amount"`
	Currency            string  `json:"currency"`           // Default: ZAR
	CollectionReference string  `json:"collectionReference"` // Your reference
	CollectionDate      string  `json:"collectionDate,omitempty"` // Format: YYYY-MM-DD
	Description         string  `json:"description"`
	MandateID           string `json:"mandateId,omitempty"`
	ContractReference   string `json:"contractReference,omitempty"`
	NotificationEmail   string `json:"notificationEmail,omitempty"`
	IdempotencyKey      string `json:"idempotencyKey,omitempty"`
}

type EFTCollectionResponse struct {
	TransactionID        string    `json:"transactionId"`
	Status               string    `json:"status"` // PENDING, PROCESSING, COMPLETED, FAILED, REJECTED
	StatusDescription    string    `json:"statusDescription"`
	CollectionReference  string    `json:"collectionReference"`
	Amount               float64   `json:"amount"`
	Currency             string    `json:"currency"`
	ProcessingDate       time.Time `json:"processingDate"`
	SettlementDate       string    `json:"settlementDate,omitempty"`
	DebtorName           string    `json:"debtorName"`
	Message              string    `json:"message,omitempty"`
}

type CollectionStatusResponse struct {
	TransactionID        string    `json:"transactionId"`
	CollectionReference  string    `json:"collectionReference"`
	Status               string    `json:"status"`
	StatusDescription    string    `json:"statusDescription"`
	Amount               float64   `json:"amount"`
	Currency             string    `json:"currency"`
	CreditorAccount      string    `json:"creditorAccount"`
	DebtorAccount        string    `json:"debtorAccount"`
	DebtorName           string    `json:"debtorName"`
	ProcessingDate       time.Time `json:"processingDate"`
	SettlementDate       string    `json:"settlementDate,omitempty"`
	FailureReason        string    `json:"failureReason,omitempty"`
	RejectionReason      string    `json:"rejectionReason,omitempty"`
	LastUpdated          time.Time `json:"lastUpdated"`
}

type BatchCollectionRequest struct {
	BatchReference        string                 `json:"batchReference"`
	CreditorAccountNumber string                 `json:"creditorAccountNumber"`
	TotalAmount           float64                `json:"totalAmount"`
	TotalCount            int                    `json:"totalCount"`
	ProcessingDate        string                 `json:"processingDate,omitempty"`
	Collections           []EFTCollectionRequest `json:"collections"`
}

type BatchCollectionResponse struct {
	BatchID            string              `json:"batchId"`
	BatchReference     string              `json:"batchReference"`
	Status             string              `json:"status"`
	TotalAmount        float64             `json:"totalAmount"`
	TotalCount         int                 `json:"totalCount"`
	SuccessCount       int                 `json:"successCount"`
	FailureCount       int                 `json:"failureCount"`
	ProcessingDate     time.Time           `json:"processingDate"`
	CollectionResults  []CollectionResult  `json:"collectionResults,omitempty"`
}

type CollectionResult struct {
	CollectionReference string  `json:"collectionReference"`
	TransactionID       string  `json:"transactionId,omitempty"`
	Status              string  `json:"status"`
	StatusDescription   string  `json:"statusDescription"`
	Amount              float64 `json:"amount"`
	DebtorName          string  `json:"debtorName"`
	FailureReason       string  `json:"failureReason,omitempty"`
}

type DisputeRequest struct {
	TransactionID string `json:"transactionId"`
	Reason        string `json:"reason"`
	Description   string `json:"description,omitempty"`
}

type DisputeResponse struct {
	DisputeID         string    `json:"disputeId"`
	TransactionID     string    `json:"transactionId"`
	Status            string    `json:"status"` // SUBMITTED, UNDER_REVIEW, RESOLVED, REJECTED
	Reason            string    `json:"reason"`
	SubmittedDate     time.Time `json:"submittedDate"`
	ResolutionDate    string    `json:"resolutionDate,omitempty"`
	ResolutionOutcome string    `json:"resolutionOutcome,omitempty"`
}

func (c *Client) CreateEFTCollection(ctx context.Context, req EFTCollectionRequest) (*EFTCollectionResponse, error) {
	if req.Currency == "" {
		req.Currency = "ZAR"
	}
	if req.CollectionDate == "" {
		req.CollectionDate = time.Now().Format("2006-01-02")
	}

	var result EFTCollectionResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/collections/eft", req, &result); err != nil {
		return nil, fmt.Errorf("failed to create EFT collection: %w", err)
	}

	return &result, nil
}

func (c *Client) GetCollectionStatus(ctx context.Context, transactionID string) (*CollectionStatusResponse, error) {
	path := fmt.Sprintf("/api/v1/collections/%s/status", transactionID)
	
	var result CollectionStatusResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get collection status: %w", err)
	}

	return &result, nil
}

func (c *Client) GetCollectionStatusByReference(ctx context.Context, reference string) (*CollectionStatusResponse, error) {
	path := fmt.Sprintf("/api/v1/collections/status?reference=%s", reference)
	
	var result CollectionStatusResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get collection status by reference: %w", err)
	}

	return &result, nil
}

func (c *Client) CreateBatchCollection(ctx context.Context, req BatchCollectionRequest) (*BatchCollectionResponse, error) {
	if req.ProcessingDate == "" {
		req.ProcessingDate = time.Now().Format("2006-01-02")
	}

	var result BatchCollectionResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/collections/batch", req, &result); err != nil {
		return nil, fmt.Errorf("failed to create batch collection: %w", err)
	}

	return &result, nil
}

func (c *Client) GetBatchCollectionStatus(ctx context.Context, batchID string) (*BatchCollectionResponse, error) {
	path := fmt.Sprintf("/api/v1/collections/batch/%s/status", batchID)
	
	var result BatchCollectionResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get batch collection status: %w", err)
	}

	return &result, nil
}

func (c *Client) CancelCollection(ctx context.Context, transactionID string, reason string) error {
	payload := map[string]string{
		"reason": reason,
	}

	path := fmt.Sprintf("/api/v1/collections/%s/cancel", transactionID)
	if err := c.DoRequest(ctx, "POST", path, payload, nil); err != nil {
		return fmt.Errorf("failed to cancel collection: %w", err)
	}

	return nil
}

func (c *Client) DisputeCollection(ctx context.Context, req DisputeRequest) (*DisputeResponse, error) {
	var result DisputeResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/collections/dispute", req, &result); err != nil {
		return nil, fmt.Errorf("failed to dispute collection: %w", err)
	}

	return &result, nil
}

func (c *Client) GetDisputeStatus(ctx context.Context, disputeID string) (*DisputeResponse, error) {
	path := fmt.Sprintf("/api/v1/collections/dispute/%s", disputeID)
	
	var result DisputeResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get dispute status: %w", err)
	}

	return &result, nil
}
