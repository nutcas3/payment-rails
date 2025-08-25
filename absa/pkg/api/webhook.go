package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
)

const (
	WebhookSignatureHeader = "X-Absa-Signature"
)

type WebhookHandler struct {
	Secret string
}

type WebhookHandlers struct {
	PaymentSuccessHandler      func(PaymentSuccessWebhook) error
	PaymentFailureHandler      func(PaymentFailureWebhook) error
	TransactionStatusHandler   func(TransactionStatusWebhook) error
	AccountUpdateHandler       func(AccountUpdateWebhook) error
	DefaultHandler             func(map[string]interface{}) error
}

type WebhookEvent struct {
	EventType string `json:"eventType"`
}

type PaymentSuccessWebhook struct {
	EventType       string `json:"eventType"`
	TransactionID   string `json:"transactionId"`
	Reference       string `json:"reference"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	Status          string `json:"status"`
	Timestamp       string `json:"timestamp"`
	AccountNumber   string `json:"accountNumber"`
	RecipientName   string `json:"recipientName"`
	RecipientBank   string `json:"recipientBank"`
	TransactionFee  string `json:"transactionFee"`
	PaymentMethod   string `json:"paymentMethod"`
	AdditionalInfo  map[string]interface{} `json:"additionalInfo,omitempty"`
}

type PaymentFailureWebhook struct {
	EventType       string `json:"eventType"`
	TransactionID   string `json:"transactionId"`
	Reference       string `json:"reference"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	Status          string `json:"status"`
	ErrorCode       string `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
	Timestamp       string `json:"timestamp"`
	AccountNumber   string `json:"accountNumber"`
	PaymentMethod   string `json:"paymentMethod"`
	AdditionalInfo  map[string]interface{} `json:"additionalInfo,omitempty"`
}

type TransactionStatusWebhook struct {
	EventType       string `json:"eventType"`
	TransactionID   string `json:"transactionId"`
	Reference       string `json:"reference"`
	Status          string `json:"status"`
	Timestamp       string `json:"timestamp"`
	AdditionalInfo  map[string]interface{} `json:"additionalInfo,omitempty"`
}

type AccountUpdateWebhook struct {
	EventType       string `json:"eventType"`
	AccountNumber   string `json:"accountNumber"`
	UpdateType      string `json:"updateType"`
	Timestamp       string `json:"timestamp"`
	AdditionalInfo  map[string]interface{} `json:"additionalInfo,omitempty"`
}

func NewWebhookHandler(secret string) *WebhookHandler {
	return &WebhookHandler{
		Secret: secret,
	}
}

func (h *WebhookHandler) ValidateSignature(body []byte, signature string) bool {
	if h.Secret == "" || signature == "" {
		return false
	}

	mac := hmac.New(sha256.New, []byte(h.Secret))
	mac.Write(body)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(expectedMAC))
}

func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request, handlers WebhookHandlers) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	signature := r.Header.Get(WebhookSignatureHeader)
	if !h.ValidateSignature(body, signature) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	var event WebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Error parsing webhook event", http.StatusBadRequest)
		return
	}

	switch event.EventType {
	case "payment.success":
		var webhook PaymentSuccessWebhook
		if err := json.Unmarshal(body, &webhook); err != nil {
			http.Error(w, "Error parsing webhook data", http.StatusBadRequest)
			return
		}
		if handlers.PaymentSuccessHandler != nil {
			if err := handlers.PaymentSuccessHandler(webhook); err != nil {
				http.Error(w, "Error processing webhook", http.StatusInternalServerError)
				return
			}
		}
	case "payment.failure":
		var webhook PaymentFailureWebhook
		if err := json.Unmarshal(body, &webhook); err != nil {
			http.Error(w, "Error parsing webhook data", http.StatusBadRequest)
			return
		}
		if handlers.PaymentFailureHandler != nil {
			if err := handlers.PaymentFailureHandler(webhook); err != nil {
				http.Error(w, "Error processing webhook", http.StatusInternalServerError)
				return
			}
		}
	case "transaction.status":
		var webhook TransactionStatusWebhook
		if err := json.Unmarshal(body, &webhook); err != nil {
			http.Error(w, "Error parsing webhook data", http.StatusBadRequest)
			return
		}
		if handlers.TransactionStatusHandler != nil {
			if err := handlers.TransactionStatusHandler(webhook); err != nil {
				http.Error(w, "Error processing webhook", http.StatusInternalServerError)
				return
			}
		}
	case "account.update":
		var webhook AccountUpdateWebhook
		if err := json.Unmarshal(body, &webhook); err != nil {
			http.Error(w, "Error parsing webhook data", http.StatusBadRequest)
			return
		}
		if handlers.AccountUpdateHandler != nil {
			if err := handlers.AccountUpdateHandler(webhook); err != nil {
				http.Error(w, "Error processing webhook", http.StatusInternalServerError)
				return
			}
		}
	default:
		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			http.Error(w, "Error parsing webhook data", http.StatusBadRequest)
			return
		}
		if handlers.DefaultHandler != nil {
			if err := handlers.DefaultHandler(data); err != nil {
				http.Error(w, "Error processing webhook", http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}
