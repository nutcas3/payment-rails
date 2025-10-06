package api

import (
	"context"
	"fmt"
	"time"
)

type AccountVerificationRequest struct {
	AccountNumber string `json:"accountNumber"`
	BankCode      string `json:"bankCode,omitempty"` 
	AccountType   string `json:"accountType,omitempty"` 
}

type AccountVerificationResponse struct {
	AccountNumber   string `json:"accountNumber"`
	AccountName     string `json:"accountName"`
	AccountType     string `json:"accountType"`
	BankName        string `json:"bankName"`
	BranchCode      string `json:"branchCode"`
	IsValid         bool   `json:"isValid"`
	Status          string `json:"status"` 
	VerificationID  string `json:"verificationId"`
	Message         string `json:"message,omitempty"`
}

type TransactionHistoryRequest struct {
	AccountNumber   string  `json:"accountNumber"`
	FromDate        string  `json:"fromDate"`
	ToDate          string  `json:"toDate"`
	TransactionType string  `json:"transactionType,omitempty"`
	MinAmount       float64 `json:"minAmount,omitempty"`
	MaxAmount       float64 `json:"maxAmount,omitempty"`
	PageNumber      int     `json:"pageNumber,omitempty"`
	PageSize        int     `json:"pageSize,omitempty"`  
}

type TransactionHistoryResponse struct {
	AccountNumber string        `json:"accountNumber"`
	AccountName   string        `json:"accountName"`
	FromDate      string        `json:"fromDate"`
	ToDate        string        `json:"toDate"`
	Transactions  []Transaction `json:"transactions"`
	TotalCount    int           `json:"totalCount"`
	PageNumber    int           `json:"pageNumber"`
	PageSize      int           `json:"pageSize"`
	TotalPages    int           `json:"totalPages"`
}

type Transaction struct {
	TransactionID   string    `json:"transactionId"`
	Date            time.Time `json:"date"`
	ValueDate       string    `json:"valueDate,omitempty"`
	Description     string    `json:"description"`
	Reference       string    `json:"reference"`
	Amount          float64   `json:"amount"`
	Balance         float64   `json:"balance"`
	Type            string    `json:"type"` 
	Category        string    `json:"category,omitempty"` 
	CounterParty    string    `json:"counterParty,omitempty"`
	CounterPartyAccount string `json:"counterPartyAccount,omitempty"`
}

type AccountBalanceRequest struct {
	AccountNumber string `json:"accountNumber"`
}
type AccountBalanceResponse struct {
	AccountNumber   string    `json:"accountNumber"`
	AccountName     string    `json:"accountName"`
	AccountType     string    `json:"accountType"`
	Currency        string    `json:"currency"`
	CurrentBalance  float64   `json:"currentBalance"`
	AvailableBalance float64  `json:"availableBalance"`
	OverdraftLimit  float64   `json:"overdraftLimit,omitempty"`
	LastUpdated     time.Time `json:"lastUpdated"`
}

type ProofOfPaymentRequest struct {
	TransactionID    string `json:"transactionId,omitempty"`
	PaymentReference string `json:"paymentReference,omitempty"`
	Format           string `json:"format,omitempty"` 
}
type ProofOfPaymentResponse struct {
	TransactionID       string    `json:"transactionId"`
	PaymentReference    string    `json:"paymentReference"`
	PaymentDate         time.Time `json:"paymentDate"`
	Amount              float64   `json:"amount"`
	Currency            string    `json:"currency"`
	SourceAccount       string    `json:"sourceAccount"`
	SourceAccountName   string    `json:"sourceAccountName"`
	BeneficiaryAccount  string    `json:"beneficiaryAccount"`
	BeneficiaryName     string    `json:"beneficiaryName"`
	BeneficiaryBank     string    `json:"beneficiaryBank"`
	Status              string    `json:"status"`
	Description         string    `json:"description"`
	DocumentURL         string    `json:"documentUrl,omitempty"` 
	DocumentData        string    `json:"documentData,omitempty"` 
}

type StatementRequest struct {
	AccountNumber string `json:"accountNumber"`
	FromDate      string `json:"fromDate"`                
	ToDate        string `json:"toDate"`               
	Format        string `json:"format,omitempty"`       
	Email         string `json:"email,omitempty"`       
}
type StatementResponse struct {
	AccountNumber    string        `json:"accountNumber"`
	AccountName      string        `json:"accountName"`
	FromDate         string        `json:"fromDate"`
	ToDate           string        `json:"toDate"`
	OpeningBalance   float64       `json:"openingBalance"`
	ClosingBalance   float64       `json:"closingBalance"`
	TotalCredits     float64       `json:"totalCredits"`
	TotalDebits      float64       `json:"totalDebits"`
	TransactionCount int           `json:"transactionCount"`
	Transactions     []Transaction `json:"transactions,omitempty"`
	DocumentURL      string        `json:"documentUrl,omitempty"` 
	DocumentData     string        `json:"documentData,omitempty"` 
	EmailSent        bool          `json:"emailSent,omitempty"`
}
type NotificationPreferencesRequest struct {
	AccountNumber      string   `json:"accountNumber"`
	EmailNotifications bool     `json:"emailNotifications"`
	SMSNotifications   bool     `json:"smsNotifications"`
	EmailAddresses     []string `json:"emailAddresses,omitempty"`
	MobileNumbers      []string `json:"mobileNumbers,omitempty"`
	NotificationTypes  []string `json:"notificationTypes,omitempty"` 
}

func (c *Client) VerifyAccount(ctx context.Context, req AccountVerificationRequest) (*AccountVerificationResponse, error) {
	var result AccountVerificationResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/accounts/verify", req, &result); err != nil {
		return nil, fmt.Errorf("failed to verify account: %w", err)
	}

	return &result, nil
}

func (c *Client) GetTransactionHistory(ctx context.Context, req TransactionHistoryRequest) (*TransactionHistoryResponse, error) {
	// Set defaults
	if req.TransactionType == "" {
		req.TransactionType = "ALL"
	}
	if req.PageNumber == 0 {
		req.PageNumber = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 50
	}

	var result TransactionHistoryResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/accounts/transactions", req, &result); err != nil {
		return nil, fmt.Errorf("failed to get transaction history: %w", err)
	}

	return &result, nil
}

func (c *Client) GetAccountBalance(ctx context.Context, accountNumber string) (*AccountBalanceResponse, error) {
	req := AccountBalanceRequest{
		AccountNumber: accountNumber,
	}

	var result AccountBalanceResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/accounts/balance", req, &result); err != nil {
		return nil, fmt.Errorf("failed to get account balance: %w", err)
	}

	return &result, nil
}

func (c *Client) GetProofOfPayment(ctx context.Context, req ProofOfPaymentRequest) (*ProofOfPaymentResponse, error) {
	if req.Format == "" {
		req.Format = "PDF"
	}

	var result ProofOfPaymentResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/payments/proof", req, &result); err != nil {
		return nil, fmt.Errorf("failed to get proof of payment: %w", err)
	}

	return &result, nil
}

func (c *Client) GetStatement(ctx context.Context, req StatementRequest) (*StatementResponse, error) {
	if req.Format == "" {
		req.Format = "PDF"
	}

	var result StatementResponse
	if err := c.DoRequest(ctx, "POST", "/api/v1/accounts/statement", req, &result); err != nil {
		return nil, fmt.Errorf("failed to get statement: %w", err)
	}

	return &result, nil
}

func (c *Client) UpdateNotificationPreferences(ctx context.Context, req NotificationPreferencesRequest) error {
	if err := c.DoRequest(ctx, "POST", "/api/v1/accounts/notifications", req, nil); err != nil {
		return fmt.Errorf("failed to update notification preferences: %w", err)
	}

	return nil
}
