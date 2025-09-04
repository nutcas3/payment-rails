package api

import (
	"time"

	"github.com/shopspring/decimal"
)

// Common status constants
const (
	StatusSuccess     = "success"
	StatusPending     = "pending"
	StatusFailed      = "failed"
	StatusProcessing  = "processing"
	StatusCompleted   = "completed"
)

// Common error codes
const (
	ErrInvalidRequest      = "invalid_request"
	ErrInsufficientFunds   = "insufficient_funds"
	ErrAuthenticationError = "authentication_error"
	ErrInvalidAccount      = "invalid_account"
	ErrSystemError         = "system_error"
)

// APIError represents an error returned by the SasaPay API
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// AuthTokenRequest represents a request for an authentication token
type AuthTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// AuthTokenResponse represents the response from an authentication token request
type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// C2B API Models

// C2BRequest represents a request to initiate a Customer to Business payment
type C2BRequest struct {
	MerchantCode string          `json:"merchant_code"`
	PhoneNumber  string          `json:"phone_number"`
	Amount       decimal.Decimal `json:"amount"`
	Reference    string          `json:"reference"`
	Description  string          `json:"description"`
	CallbackURL  string          `json:"callback_url,omitempty"`
}

// C2BResponse represents the response from a Customer to Business payment request
type C2BResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

// B2C API Models

// B2CRequest represents a request to initiate a Business to Customer payment
type B2CRequest struct {
	MerchantCode string          `json:"merchant_code"`
	PhoneNumber  string          `json:"phone_number"`
	Amount       decimal.Decimal `json:"amount"`
	Reference    string          `json:"reference"`
	Description  string          `json:"description"`
	CallbackURL  string          `json:"callback_url,omitempty"`
}

// B2CResponse represents the response from a Business to Customer payment request
type B2CResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

// B2B API Models

// B2BRequest represents a request to initiate a Business to Business payment
type B2BRequest struct {
	SourceMerchantCode      string          `json:"source_merchant_code"`
	DestinationMerchantCode string          `json:"destination_merchant_code"`
	Amount                  decimal.Decimal `json:"amount"`
	Reference               string          `json:"reference"`
	Description             string          `json:"description"`
	CallbackURL             string          `json:"callback_url,omitempty"`
}

// B2BResponse represents the response from a Business to Business payment request
type B2BResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

// Wallet as a Service Models

// CreateWalletRequest represents a request to create a new wallet
type CreateWalletRequest struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email,omitempty"`
	IDNumber    string `json:"id_number"`
	CallbackURL string `json:"callback_url,omitempty"`
}

// CreateWalletResponse represents the response from a create wallet request
type CreateWalletResponse struct {
	WalletID  string    `json:"wallet_id"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// WalletBalanceRequest represents a request to get a wallet balance
type WalletBalanceRequest struct {
	WalletID string `json:"wallet_id"`
}

// WalletBalanceResponse represents the response from a wallet balance request
type WalletBalanceResponse struct {
	WalletID  string          `json:"wallet_id"`
	Balance   decimal.Decimal `json:"balance"`
	Currency  string          `json:"currency"`
	Status    string          `json:"status"`
	Timestamp time.Time       `json:"timestamp"`
}

// WalletTransferRequest represents a request to transfer funds between wallets
type WalletTransferRequest struct {
	SourceWalletID      string          `json:"source_wallet_id"`
	DestinationWalletID string          `json:"destination_wallet_id"`
	Amount              decimal.Decimal `json:"amount"`
	Reference           string          `json:"reference"`
	Description         string          `json:"description"`
	CallbackURL         string          `json:"callback_url,omitempty"`
}

// WalletTransferResponse represents the response from a wallet transfer request
type WalletTransferResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

// WalletStatementRequest represents a request to get a wallet statement
type WalletStatementRequest struct {
	WalletID  string    `json:"wallet_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// StatementTransaction represents a single transaction in a wallet statement
type StatementTransaction struct {
	TransactionID string          `json:"transaction_id"`
	Type          string          `json:"type"`
	Amount        decimal.Decimal `json:"amount"`
	Balance       decimal.Decimal `json:"balance"`
	Description   string          `json:"description"`
	Reference     string          `json:"reference"`
	Timestamp     time.Time       `json:"timestamp"`
}

// WalletStatementResponse represents the response from a wallet statement request
type WalletStatementResponse struct {
	WalletID     string                `json:"wallet_id"`
	Transactions []StatementTransaction `json:"transactions"`
	Status       string                `json:"status"`
	Timestamp    time.Time             `json:"timestamp"`
}

// Transaction Status Models

// TransactionStatusRequest represents a request to check a transaction status
type TransactionStatusRequest struct {
	TransactionID string `json:"transaction_id"`
}

// TransactionStatusResponse represents the response from a transaction status request
type TransactionStatusResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

// VerifyTransactionRequest represents a request to verify a transaction
type VerifyTransactionRequest struct {
	TransactionID string `json:"transaction_id"`
}

// VerifyTransactionResponse represents the response from a verify transaction request
type VerifyTransactionResponse struct {
	TransactionID string          `json:"transaction_id"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Status        string          `json:"status"`
	Message       string          `json:"message"`
	Timestamp     time.Time       `json:"timestamp"`
}

// Webhook Models

// WebhookEvent represents an event received via webhook
type WebhookEvent struct {
	EventType     string          `json:"event_type"`
	TransactionID string          `json:"transaction_id"`
	MerchantCode  string          `json:"merchant_code"`
	PhoneNumber   string          `json:"phone_number,omitempty"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Reference     string          `json:"reference"`
	Status        string          `json:"status"`
	Message       string          `json:"message"`
	Timestamp     time.Time       `json:"timestamp"`
}

// WebhookHandlers contains handler functions for different webhook events
type WebhookHandlers struct {
	PaymentReceived   func(event WebhookEvent)
	PaymentCompleted  func(event WebhookEvent)
	PaymentFailed     func(event WebhookEvent)
	WalletCreated     func(event WebhookEvent)
	WalletTransferred func(event WebhookEvent)
}
