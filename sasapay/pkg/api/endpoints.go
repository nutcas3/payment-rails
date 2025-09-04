package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// C2B API Endpoints

// CustomerToBusiness initiates a payment from a customer to a business
func (c *Client) CustomerToBusiness(req C2BRequest) (*C2BResponse, error) {
	if req.MerchantCode == "" {
		return nil, fmt.Errorf("merchant code is required")
	}
	if req.PhoneNumber == "" {
		return nil, fmt.Errorf("phone number is required")
	}
	if req.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("amount must be greater than zero")
	}
	if req.Reference == "" {
		return nil, fmt.Errorf("reference is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/c2b/payment", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp C2BResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// B2C API Endpoints

// BusinessToCustomer initiates a payment from a business to a customer
func (c *Client) BusinessToCustomer(req B2CRequest) (*B2CResponse, error) {
	// Validate request
	if req.MerchantCode == "" {
		return nil, fmt.Errorf("merchant code is required")
	}
	if req.PhoneNumber == "" {
		return nil, fmt.Errorf("phone number is required")
	}
	if req.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("amount must be greater than zero")
	}
	if req.Reference == "" {
		return nil, fmt.Errorf("reference is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/b2c/payment", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp B2CResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// B2B API Endpoints

// BusinessToBusiness initiates a payment from one business to another
func (c *Client) BusinessToBusiness(req B2BRequest) (*B2BResponse, error) {
	// Validate request
	if req.SourceMerchantCode == "" {
		return nil, fmt.Errorf("source merchant code is required")
	}
	if req.DestinationMerchantCode == "" {
		return nil, fmt.Errorf("destination merchant code is required")
	}
	if req.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("amount must be greater than zero")
	}
	if req.Reference == "" {
		return nil, fmt.Errorf("reference is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/b2b/payment", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp B2BResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// Wallet as a Service Endpoints

// CreateWallet creates a new wallet for a customer
func (c *Client) CreateWallet(req CreateWalletRequest) (*CreateWalletResponse, error) {
	// Validate request
	if req.PhoneNumber == "" {
		return nil, fmt.Errorf("phone number is required")
	}
	if req.FirstName == "" {
		return nil, fmt.Errorf("first name is required")
	}
	if req.LastName == "" {
		return nil, fmt.Errorf("last name is required")
	}
	if req.IDNumber == "" {
		return nil, fmt.Errorf("ID number is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/waas/wallet", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp CreateWalletResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// GetWalletBalance retrieves the balance of a wallet
func (c *Client) GetWalletBalance(req WalletBalanceRequest) (*WalletBalanceResponse, error) {
	// Validate request
	if req.WalletID == "" {
		return nil, fmt.Errorf("wallet ID is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/waas/balance", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp WalletBalanceResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// TransferToWallet transfers funds from one wallet to another
func (c *Client) TransferToWallet(req WalletTransferRequest) (*WalletTransferResponse, error) {
	// Validate request
	if req.SourceWalletID == "" {
		return nil, fmt.Errorf("source wallet ID is required")
	}
	if req.DestinationWalletID == "" {
		return nil, fmt.Errorf("destination wallet ID is required")
	}
	if req.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("amount must be greater than zero")
	}
	if req.Reference == "" {
		return nil, fmt.Errorf("reference is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/waas/transfer", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp WalletTransferResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// GetWalletStatement retrieves a statement for a wallet
func (c *Client) GetWalletStatement(req WalletStatementRequest) (*WalletStatementResponse, error) {
	// Validate request
	if req.WalletID == "" {
		return nil, fmt.Errorf("wallet ID is required")
	}
	if req.StartDate.IsZero() {
		return nil, fmt.Errorf("start date is required")
	}
	if req.EndDate.IsZero() {
		return nil, fmt.Errorf("end date is required")
	}
	if req.EndDate.Before(req.StartDate) {
		return nil, fmt.Errorf("end date must be after start date")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/waas/statement", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp WalletStatementResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// Transaction Status Endpoints

// CheckTransactionStatus checks the status of a transaction
func (c *Client) CheckTransactionStatus(req TransactionStatusRequest) (*TransactionStatusResponse, error) {
	// Validate request
	if req.TransactionID == "" {
		return nil, fmt.Errorf("transaction ID is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/transaction/status", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp TransactionStatusResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

// VerifyTransaction verifies a transaction
func (c *Client) VerifyTransaction(req VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	// Validate request
	if req.TransactionID == "" {
		return nil, fmt.Errorf("transaction ID is required")
	}

	// Send request
	respBody, err := c.SendRequest("POST", "/transaction/verify", req)
	if err != nil {
		return nil, err
	}

	// Parse response
	var resp VerifyTransactionResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Set timestamp if not provided
	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}
