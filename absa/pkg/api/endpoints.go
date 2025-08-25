package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

// Account Balance
func (c *Client) GetAccountBalance(req AccountBalanceRequest) (*AccountBalanceResponse, error) {
	endpoint := "/accounts/balance"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response AccountBalanceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing account balance response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Mini Statement
func (c *Client) GetMiniStatement(req MiniStatementRequest) (*MiniStatementResponse, error) {
	endpoint := "/accounts/mini-statement"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response MiniStatementResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing mini statement response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Full Statement
func (c *Client) GetFullStatement(req FullStatementRequest) (*FullStatementResponse, error) {
	endpoint := "/accounts/full-statement"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response FullStatementResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing full statement response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Account Validation
func (c *Client) ValidateAccount(req AccountValidateRequest) (*AccountValidateResponse, error) {
	endpoint := "/accounts/validate"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response AccountValidateResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing account validation response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Send Money (Bank Transfer)
func (c *Client) SendMoney(req SendMoneyRequest) (*SendMoneyResponse, error) {
	endpoint := "/payments/bank-transfer"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response SendMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing send money response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Internal Bank Transfer
func (c *Client) SendInternalBankTransfer(req SendMoneyRequest) (*SendMoneyResponse, error) {
	endpoint := "/payments/internal-transfer"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response SendMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing internal transfer response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Mobile Wallet
func (c *Client) SendToMobileWallet(req MobileWalletRequest) (*MobileWalletResponse, error) {
	endpoint := "/payments/mobile-wallet"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response MobileWalletResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing mobile wallet response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Bill Payment
func (c *Client) PayBill(req BillPaymentRequest) (*BillPaymentResponse, error) {
	endpoint := "/payments/bill-payment"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BillPaymentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing bill payment response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Receive Money
func (c *Client) ReceiveMoney(req ReceiveMoneyRequest) (*ReceiveMoneyResponse, error) {
	endpoint := "/payments/receive"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response ReceiveMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing receive money response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}
	
	// Set expiry date if not provided but expiry minutes is set in request
	if response.ExpiryDate.IsZero() && req.ExpiryMinutes > 0 {
		response.ExpiryDate = time.Now().UTC().Add(time.Duration(req.ExpiryMinutes) * time.Minute)
	}

	return &response, nil
}

// Transaction Query
func (c *Client) QueryTransaction(req TransactionQueryRequest) (*TransactionQueryResponse, error) {
	endpoint := "/transactions/status"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response TransactionQueryResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing transaction query response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Helper function to validate amount is positive
func validateAmount(amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return fmt.Errorf("amount must be greater than zero")
	}
	return nil
}

// Airtime Purchase
func (c *Client) PurchaseAirtime(req AirtimePurchaseRequest) (*AirtimePurchaseResponse, error) {
	endpoint := "/payments/airtime"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response AirtimePurchaseResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing airtime purchase response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}
