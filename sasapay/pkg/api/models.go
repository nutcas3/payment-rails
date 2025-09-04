package api

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	StatusSuccess     = "success"
	StatusPending     = "pending"
	StatusFailed      = "failed"
	StatusProcessing  = "processing"
	StatusCompleted   = "completed"
)

const (
	ErrInvalidRequest      = "invalid_request"
	ErrInsufficientFunds   = "insufficient_funds"
	ErrAuthenticationError = "authentication_error"
	ErrInvalidAccount      = "invalid_account"
	ErrSystemError         = "system_error"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

type AuthTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}


type C2BRequest struct {
	MerchantCode string          `json:"merchant_code"`
	PhoneNumber  string          `json:"phone_number"`
	Amount       decimal.Decimal `json:"amount"`
	Reference    string          `json:"reference"`
	Description  string          `json:"description"`
	CallbackURL  string          `json:"callback_url,omitempty"`
}

type C2BResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}


type B2CRequest struct {
	MerchantCode string          `json:"merchant_code"`
	PhoneNumber  string          `json:"phone_number"`
	Amount       decimal.Decimal `json:"amount"`
	Reference    string          `json:"reference"`
	Description  string          `json:"description"`
	CallbackURL  string          `json:"callback_url,omitempty"`
}

type B2CResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}


type B2BRequest struct {
	SourceMerchantCode      string          `json:"source_merchant_code"`
	DestinationMerchantCode string          `json:"destination_merchant_code"`
	Amount                  decimal.Decimal `json:"amount"`
	Reference               string          `json:"reference"`
	Description             string          `json:"description"`
	CallbackURL             string          `json:"callback_url,omitempty"`
}

type B2BResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}


type CreateWalletRequest struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email,omitempty"`
	IDNumber    string `json:"id_number"`
	CallbackURL string `json:"callback_url,omitempty"`
}

type CreateWalletResponse struct {
	WalletID  string    `json:"wallet_id"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type WalletBalanceRequest struct {
	WalletID string `json:"wallet_id"`
}

type WalletBalanceResponse struct {
	WalletID  string          `json:"wallet_id"`
	Balance   decimal.Decimal `json:"balance"`
	Currency  string          `json:"currency"`
	Status    string          `json:"status"`
	Timestamp time.Time       `json:"timestamp"`
}

type WalletTransferRequest struct {
	SourceWalletID      string          `json:"source_wallet_id"`
	DestinationWalletID string          `json:"destination_wallet_id"`
	Amount              decimal.Decimal `json:"amount"`
	Reference           string          `json:"reference"`
	Description         string          `json:"description"`
	CallbackURL         string          `json:"callback_url,omitempty"`
}

type WalletTransferResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

type WalletStatementRequest struct {
	WalletID  string    `json:"wallet_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type StatementTransaction struct {
	TransactionID string          `json:"transaction_id"`
	Type          string          `json:"type"`
	Amount        decimal.Decimal `json:"amount"`
	Balance       decimal.Decimal `json:"balance"`
	Description   string          `json:"description"`
	Reference     string          `json:"reference"`
	Timestamp     time.Time       `json:"timestamp"`
}

type WalletStatementResponse struct {
	WalletID     string                `json:"wallet_id"`
	Transactions []StatementTransaction `json:"transactions"`
	Status       string                `json:"status"`
	Timestamp    time.Time             `json:"timestamp"`
}


type TransactionStatusRequest struct {
	TransactionID string `json:"transaction_id"`
}

type TransactionStatusResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

type VerifyTransactionRequest struct {
	TransactionID string `json:"transaction_id"`
}

type VerifyTransactionResponse struct {
	TransactionID string          `json:"transaction_id"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Status        string          `json:"status"`
	Message       string          `json:"message"`
	Timestamp     time.Time       `json:"timestamp"`
}


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

type WebhookHandlers struct {
	PaymentReceived   func(event WebhookEvent)
	PaymentCompleted  func(event WebhookEvent)
	PaymentFailed     func(event WebhookEvent)
	WalletCreated     func(event WebhookEvent)
	WalletTransferred func(event WebhookEvent)
}
