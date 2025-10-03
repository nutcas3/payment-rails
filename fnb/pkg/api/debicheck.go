package api

import (
	"context"
	"fmt"
	"time"
)


type MandateRequest struct {

	CreditorName          string `json:"creditorName"`
	CreditorAbbreviation  string `json:"creditorAbbreviation"` 
	CreditorAccountNumber string `json:"creditorAccountNumber"`
	DebtorName          string `json:"debtorName"`
	DebtorIDNumber      string `json:"debtorIdNumber,omitempty"` 
	DebtorAccountNumber string `json:"debtorAccountNumber"`
	DebtorAccountType   string `json:"debtorAccountType,omitempty"` 
	DebtorBankCode      string `json:"debtorBankCode"` 
	DebtorEmail         string `json:"debtorEmail,omitempty"`
	DebtorMobile        string `json:"debtorMobile,omitempty"`

	ContractReference   string  `json:"contractReference"`   
	MaximumAmount       float64 `json:"maximumAmount"` 
	Currency            string  `json:"currency"`            
	FrequencyType       string  `json:"frequencyType"`       
	FirstCollectionDate string  `json:"firstCollectionDate"` 
	LastCollectionDate  string  `json:"lastCollectionDate,omitempty"` 
	CollectionDay       int     `json:"collectionDay,omitempty"` 

	MandateDescription  string `json:"mandateDescription"`
	CategoryCode        string `json:"categoryCode,omitempty"` 
	IdempotencyKey      string `json:"idempotencyKey,omitempty"`
}
type MandateResponse struct {
	MandateID           string    `json:"mandateId"`
	Status              string    `json:"status"` 
	StatusDescription   string    `json:"statusDescription"`
	ContractReference   string    `json:"contractReference"`
	DebtorName          string    `json:"debtorName"`
	MaximumAmount       float64   `json:"maximumAmount"`
	Currency            string    `json:"currency"`
	FrequencyType       string    `json:"frequencyType"`
	FirstCollectionDate string    `json:"firstCollectionDate"`
	CreatedDate         time.Time `json:"createdDate"`
	ApprovalDate        string    `json:"approvalDate,omitempty"`
	ExpiryDate          string    `json:"expiryDate,omitempty"`
	Message             string    `json:"message,omitempty"`
}
type MandateStatusResponse struct {
	MandateID           string    `json:"mandateId"`
	ContractReference   string    `json:"contractReference"`
	Status              string    `json:"status"`
	StatusDescription   string    `json:"statusDescription"`
	CreditorName        string    `json:"creditorName"`
	DebtorName          string    `json:"debtorName"`
	DebtorAccountNumber string    `json:"debtorAccountNumber"`
	MaximumAmount       float64   `json:"maximumAmount"`
	Currency            string    `json:"currency"`
	FrequencyType       string    `json:"frequencyType"`
	FirstCollectionDate string    `json:"firstCollectionDate"`
	LastCollectionDate  string    `json:"lastCollectionDate,omitempty"`
	NextCollectionDate  string    `json:"nextCollectionDate,omitempty"`
	CreatedDate         time.Time `json:"createdDate"`
	ApprovalDate        string    `json:"approvalDate,omitempty"`
	LastModifiedDate    time.Time `json:"lastModifiedDate"`
	RejectionReason     string    `json:"rejectionReason,omitempty"`
}
type MandateUpdateRequest struct {
	MandateID           string  `json:"mandateId"`
	MaximumAmount       float64 `json:"maximumAmount,omitempty"`
	LastCollectionDate  string  `json:"lastCollectionDate,omitempty"`
	CollectionDay       int     `json:"collectionDay,omitempty"`
	UpdateReason        string  `json:"updateReason"`
}
type MandateCancellationRequest struct {
	MandateID          string `json:"mandateId"`
	CancellationReason string `json:"cancellationReason"`
	EffectiveDate      string `json:"effectiveDate,omitempty"` 
}
type MandateSuspensionRequest struct {
	MandateID        string `json:"mandateId"`
	SuspensionReason string `json:"suspensionReason"`
	SuspensionPeriod int    `json:"suspensionPeriod,omitempty"` 
}
type MandateCollectionRequest struct {
	MandateID           string  `json:"mandateId"`
	Amount              float64 `json:"amount"`
	CollectionReference string  `json:"collectionReference"`
	CollectionDate      string  `json:"collectionDate,omitempty"` 
	Description         string  `json:"description"`
	IdempotencyKey      string  `json:"idempotencyKey,omitempty"`
}
type MandateListRequest struct {
	Status            string `json:"status,omitempty"`            
	ContractReference string `json:"contractReference,omitempty"` 
	DebtorIDNumber    string `json:"debtorIdNumber,omitempty"`    
	ToDate            string `json:"toDate,omitempty"`            
	PageNumber        int    `json:"pageNumber,omitempty"`        
	PageSize          int    `json:"pageSize,omitempty"`          
}

type MandateListResponse struct {
	Mandates    []MandateStatusResponse `json:"mandates"`
	TotalCount  int                     `json:"totalCount"`
	PageNumber  int                     `json:"pageNumber"`
	PageSize    int                     `json:"pageSize"`
	TotalPages  int                     `json:"totalPages"`
}

func (c *Client) CreateMandate(ctx context.Context, req MandateRequest) (*MandateResponse, error) {
	if req.Currency == "" {
		req.Currency = "ZAR"
	}

	var result MandateResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/debicheck/mandates", req, &result); err != nil {
		return nil, fmt.Errorf("failed to create mandate: %w", err)
	}

	return &result, nil
}

func (c *Client) GetMandateStatus(ctx context.Context, mandateID string) (*MandateStatusResponse, error) {
	path := fmt.Sprintf("/api/v1/debicheck/mandates/%s", mandateID)
	
	var result MandateStatusResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get mandate status: %w", err)
	}

	return &result, nil
}

func (c *Client) GetMandateByContractReference(ctx context.Context, contractReference string) (*MandateStatusResponse, error) {
	path := fmt.Sprintf("/api/v1/debicheck/mandates/contract/%s", contractReference)
	
	var result MandateStatusResponse
	if err := c.DoRequest(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get mandate by contract reference: %w", err)
	}

	return &result, nil
}

func (c *Client) UpdateMandate(ctx context.Context, req MandateUpdateRequest) (*MandateResponse, error) {
	path := fmt.Sprintf("/api/v1/debicheck/mandates/%s", req.MandateID)
	
	var result MandateResponse
	if err := c.DoRequest(ctx, "PUT", path, req, &result); err != nil {
		return nil, fmt.Errorf("failed to update mandate: %w", err)
	}

	return &result, nil
}

func (c *Client) CancelMandate(ctx context.Context, req MandateCancellationRequest) error {
	if req.EffectiveDate == "" {
		req.EffectiveDate = time.Now().Format("2006-01-02")
	}

	path := fmt.Sprintf("/api/v1/debicheck/mandates/%s/cancel", req.MandateID)
	if err := c.DoRequest(ctx, "POST", path, req, nil); err != nil {
		return fmt.Errorf("failed to cancel mandate: %w", err)
	}

	return nil
}

func (c *Client) SuspendMandate(ctx context.Context, req MandateSuspensionRequest) error {
	path := fmt.Sprintf("/api/v1/debicheck/mandates/%s/suspend", req.MandateID)
	if err := c.DoRequest(ctx, "POST", path, req, nil); err != nil {
		return fmt.Errorf("failed to suspend mandate: %w", err)
	}

	return nil
}

func (c *Client) ReinstateMandate(ctx context.Context, mandateID string, reason string) error {
	payload := map[string]string{
		"reason": reason,
	}

	path := fmt.Sprintf("/api/v1/debicheck/mandates/%s/reinstate", mandateID)
	if err := c.DoRequest(ctx, "POST", path, payload, nil); err != nil {
		return fmt.Errorf("failed to reinstate mandate: %w", err)
	}

	return nil
}

func (c *Client) CollectAgainstMandate(ctx context.Context, req MandateCollectionRequest) (*EFTCollectionResponse, error) {
	if req.CollectionDate == "" {
		req.CollectionDate = time.Now().Format("2006-01-02")
	}

	var result EFTCollectionResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/debicheck/collect", req, &result); err != nil {
		return nil, fmt.Errorf("failed to collect against mandate: %w", err)
	}

	return &result, nil
}

func (c *Client) ListMandates(ctx context.Context, req MandateListRequest) (*MandateListResponse, error) {
	if req.PageNumber == 0 {
		req.PageNumber = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 50
	}

	var result MandateListResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/debicheck/mandates/list", req, &result); err != nil {
		return nil, fmt.Errorf("failed to list mandates: %w", err)
	}

	return &result, nil
}

func (c *Client) VerifyMandate(ctx context.Context, mandateID string, amount float64) (bool, error) {
	payload := map[string]interface{}{
		"mandateId": mandateID,
		"amount":    amount,
	}

	var result struct {
		Valid   bool   `json:"valid"`
		Message string `json:"message"`
	}

	if err := c.DoRequest(ctx, "POST", "/api/v1/debicheck/mandates/verify", payload, &result); err != nil {
		return false, fmt.Errorf("failed to verify mandate: %w", err)
	}

	return result.Valid, nil
}
