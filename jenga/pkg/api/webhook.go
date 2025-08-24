package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type WebhookHandler struct {
	WebhookSecret string
}

func NewWebhookHandler(webhookSecret string) *WebhookHandler {
	return &WebhookHandler{
		WebhookSecret: webhookSecret,
	}
}

func (h *WebhookHandler) ValidateSignature(payload []byte, signature string) bool {
	if h.WebhookSecret == "" || signature == "" {
		return false
	}

	mac := hmac.New(sha256.New, []byte(h.WebhookSecret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	return strings.EqualFold(signature, expectedSignature)
}

func (h *WebhookHandler) ParseWebhookRequest(r *http.Request) (*WebhookEvent, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading request body: %w", err)
	}

	signature := r.Header.Get("X-Jenga-Signature")
	if signature == "" {
		return nil, fmt.Errorf("missing X-Jenga-Signature header")
	}

	if !h.ValidateSignature(body, signature) {
		return nil, fmt.Errorf("invalid webhook signature")
	}

	var event WebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return nil, fmt.Errorf("error parsing webhook event: %w", err)
	}

	return &event, nil
}

func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request, handlers WebhookHandlers) {
	event, err := h.ParseWebhookRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch event.EventType {
	case WebhookEventTypeTransactionSuccess:
		if handlers.TransactionSuccessHandler != nil {
			handlers.TransactionSuccessHandler(event)
		}
	case WebhookEventTypeTransactionFailed:
		if handlers.TransactionFailedHandler != nil {
			handlers.TransactionFailedHandler(event)
		}
	case WebhookEventTypeAccountUpdated:
		if handlers.AccountUpdatedHandler != nil {
			handlers.AccountUpdatedHandler(event)
		}
	case WebhookEventTypeKYCUpdated:
		if handlers.KYCUpdatedHandler != nil {
			handlers.KYCUpdatedHandler(event)
		}
	default:
		if handlers.DefaultHandler != nil {
			handlers.DefaultHandler(event)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}
