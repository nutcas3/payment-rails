package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// Supported regions for cross-region transfers
const (
	RegionKenya     = "KE"
	RegionUganda    = "UG"
	RegionTanzania  = "TZ"
	RegionRwanda    = "RW"
	RegionEthiopia  = "ET"
	RegionSouthAfrica = "ZA"
	RegionNigeria   = "NG"
	RegionGhana     = "GH"
)

// CrossRegionTransferRequest represents a request to transfer money across regions
type CrossRegionTransferRequest struct {
	SourceRegion      string          `json:"source_region"`
	DestinationRegion string          `json:"destination_region"`
	SourceCurrency    string          `json:"source_currency"`
	DestCurrency      string          `json:"dest_currency"`
	Amount            decimal.Decimal `json:"amount"`
	PhoneNumber       string          `json:"phone_number"`
	Reference         string          `json:"reference"`
	Description       string          `json:"description,omitempty"`
	CallbackURL       string          `json:"callback_url,omitempty"`
}

// CrossRegionTransferResponse represents the response from a cross-region transfer
type CrossRegionTransferResponse struct {
	TransactionID      string          `json:"transaction_id"`
	Status             string          `json:"status"`
	Message            string          `json:"message"`
	SourceAmount       decimal.Decimal `json:"source_amount"`
	SourceCurrency     string          `json:"source_currency"`
	DestinationAmount  decimal.Decimal `json:"destination_amount"`
	DestinationCurrency string         `json:"destination_currency"`
	ExchangeRate       decimal.Decimal `json:"exchange_rate"`
	Timestamp          time.Time       `json:"timestamp"`
}

// CrossRegionQuoteRequest represents a request to get a quote for a cross-region transfer
type CrossRegionQuoteRequest struct {
	SourceRegion      string          `json:"source_region"`
	DestinationRegion string          `json:"destination_region"`
	SourceCurrency    string          `json:"source_currency"`
	DestCurrency      string          `json:"dest_currency"`
	Amount            decimal.Decimal `json:"amount"`
}

// CrossRegionQuoteResponse represents the response from a cross-region quote request
type CrossRegionQuoteResponse struct {
	QuoteID            string          `json:"quote_id"`
	SourceAmount       decimal.Decimal `json:"source_amount"`
	SourceCurrency     string          `json:"source_currency"`
	DestinationAmount  decimal.Decimal `json:"destination_amount"`
	DestinationCurrency string         `json:"destination_currency"`
	ExchangeRate       decimal.Decimal `json:"exchange_rate"`
	Fee                decimal.Decimal `json:"fee"`
	TotalCost          decimal.Decimal `json:"total_cost"`
	ExpiresAt          time.Time       `json:"expires_at"`
	Timestamp          time.Time       `json:"timestamp"`
}

// CrossRegionTransfer initiates a cross-region money transfer
func (c *Client) CrossRegionTransfer(req CrossRegionTransferRequest) (*CrossRegionTransferResponse, error) {
	// Validate request
	if req.SourceRegion == "" {
		return nil, fmt.Errorf("source region is required")
	}
	if req.DestinationRegion == "" {
		return nil, fmt.Errorf("destination region is required")
	}
	if req.SourceRegion == req.DestinationRegion {
		return nil, fmt.Errorf("source and destination regions must be different")
	}
	if req.SourceCurrency == "" {
		return nil, fmt.Errorf("source currency is required")
	}
	if req.DestCurrency == "" {
		return nil, fmt.Errorf("destination currency is required")
	}
	if req.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("amount must be greater than zero")
	}
	if req.PhoneNumber == "" {
		return nil, fmt.Errorf("phone number is required")
	}
	if req.Reference == "" {
		return nil, fmt.Errorf("reference is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/cross-region/transfer", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp CrossRegionTransferResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// GetCrossRegionQuote gets a quote for a cross-region transfer
func (c *Client) GetCrossRegionQuote(req CrossRegionQuoteRequest) (*CrossRegionQuoteResponse, error) {
	// Validate request
	if req.SourceRegion == "" {
		return nil, fmt.Errorf("source region is required")
	}
	if req.DestinationRegion == "" {
		return nil, fmt.Errorf("destination region is required")
	}
	if req.SourceRegion == req.DestinationRegion {
		return nil, fmt.Errorf("source and destination regions must be different")
	}
	if req.SourceCurrency == "" {
		return nil, fmt.Errorf("source currency is required")
	}
	if req.DestCurrency == "" {
		return nil, fmt.Errorf("destination currency is required")
	}
	if req.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/cross-region/quote", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp CrossRegionQuoteResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}
