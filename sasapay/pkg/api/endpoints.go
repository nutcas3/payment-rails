package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)


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

	respBody, err := c.SendRequest("POST", "/c2b/payment", req)
	if err != nil {
		return nil, err
	}

	var resp C2BResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}


func (c *Client) BusinessToCustomer(req B2CRequest) (*B2CResponse, error) {
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

	respBody, err := c.SendRequest("POST", "/b2c/payment", req)
	if err != nil {
		return nil, err
	}

	var resp B2CResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}


func (c *Client) BusinessToBusiness(req B2BRequest) (*B2BResponse, error) {
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

	respBody, err := c.SendRequest("POST", "/b2b/payment", req)
	if err != nil {
		return nil, err
	}

	var resp B2BResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}


func (c *Client) CreateWallet(req CreateWalletRequest) (*CreateWalletResponse, error) {
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

	respBody, err := c.SendRequest("POST", "/waas/wallet", req)
	if err != nil {
		return nil, err
	}

	var resp CreateWalletResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

func (c *Client) GetWalletBalance(req WalletBalanceRequest) (*WalletBalanceResponse, error) {
	if req.WalletID == "" {
		return nil, fmt.Errorf("wallet ID is required")
	}

	respBody, err := c.SendRequest("POST", "/waas/balance", req)
	if err != nil {
		return nil, err
	}

	var resp WalletBalanceResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

func (c *Client) TransferToWallet(req WalletTransferRequest) (*WalletTransferResponse, error) {
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

	respBody, err := c.SendRequest("POST", "/waas/transfer", req)
	if err != nil {
		return nil, err
	}

	var resp WalletTransferResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

func (c *Client) GetWalletStatement(req WalletStatementRequest) (*WalletStatementResponse, error) {
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

	respBody, err := c.SendRequest("POST", "/waas/statement", req)
	if err != nil {
		return nil, err
	}

	var resp WalletStatementResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}


func (c *Client) CheckTransactionStatus(req TransactionStatusRequest) (*TransactionStatusResponse, error) {
	if req.TransactionID == "" {
		return nil, fmt.Errorf("transaction ID is required")
	}

	respBody, err := c.SendRequest("POST", "/transaction/status", req)
	if err != nil {
		return nil, err
	}

	var resp TransactionStatusResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}

func (c *Client) VerifyTransaction(req VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	if req.TransactionID == "" {
		return nil, fmt.Errorf("transaction ID is required")
	}

	respBody, err := c.SendRequest("POST", "/transaction/verify", req)
	if err != nil {
		return nil, err
	}

	var resp VerifyTransactionResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if resp.Timestamp.IsZero() {
		resp.Timestamp = time.Now().UTC()
	}

	return &resp, nil
}
