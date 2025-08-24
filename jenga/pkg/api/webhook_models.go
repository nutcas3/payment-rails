package api

import (
	"encoding/json"
	"time"
)

type WebhookEventType string

const (
	WebhookEventTypeTransactionSuccess WebhookEventType = "transaction.success"
	WebhookEventTypeTransactionFailed  WebhookEventType = "transaction.failed"
	WebhookEventTypeAccountUpdated     WebhookEventType = "account.updated"
	WebhookEventTypeKYCUpdated         WebhookEventType = "kyc.updated"
)

type WebhookEvent struct {
	ID            string          `json:"id"`
	EventType     WebhookEventType `json:"event_type"`
	CreatedAt     time.Time       `json:"created_at"`
	Data          json.RawMessage `json:"data"`
	MerchantCode  string          `json:"merchant_code"`
	Reference     string          `json:"reference,omitempty"`
	TransactionID string          `json:"transaction_id,omitempty"`
}

type TransactionWebhookData struct {
	TransactionID   string    `json:"transaction_id"`
	Reference       string    `json:"reference"`
	Amount          string    `json:"amount"`
	Currency        string    `json:"currency"`
	Status          string    `json:"status"`
	StatusReason    string    `json:"status_reason,omitempty"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CustomerInfo    struct {
		Name        string `json:"name,omitempty"`
		AccountID   string `json:"account_id,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
	} `json:"customer_info,omitempty"`
}

type AccountWebhookData struct {
	AccountID   string    `json:"account_id"`
	AccountName string    `json:"account_name"`
	Status      string    `json:"status"`
	UpdatedAt   time.Time `json:"updated_at"`
	Balance     string    `json:"balance,omitempty"`
	Currency    string    `json:"currency,omitempty"`
}

type KYCWebhookData struct {
	CustomerID     string    `json:"customer_id"`
	Status         string    `json:"status"`
	VerificationType string  `json:"verification_type"`
	UpdatedAt      time.Time `json:"updated_at"`
	Reason         string    `json:"reason,omitempty"`
}

type WebhookHandlerFunc func(*WebhookEvent)

type WebhookHandlers struct {
	TransactionSuccessHandler WebhookHandlerFunc
	TransactionFailedHandler  WebhookHandlerFunc
	AccountUpdatedHandler     WebhookHandlerFunc
	KYCUpdatedHandler         WebhookHandlerFunc
	DefaultHandler            WebhookHandlerFunc
}
