package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WebhookEvent struct {
	EventID      string                 `json:"eventId"`
	EventType    string                 `json:"eventType"`
	Timestamp    time.Time              `json:"timestamp"`
	ResourceType string                 `json:"resourceType"`
	ResourceID   string                 `json:"resourceId"`
	Data         map[string]interface{} `json:"data"`
	Signature    string                 `json:"signature,omitempty"`
}

type WebhookHandler struct {
	webhookSecret string
	handlers      map[string]WebhookEventHandler
}

type WebhookEventHandler func(event WebhookEvent) error

const (
	EventPaymentCompleted         = "payment.completed"
	EventPaymentFailed            = "payment.failed"
	EventPaymentPending           = "payment.pending"
	EventPaymentCancelled         = "payment.cancelled"
	EventTransferCompleted        = "transfer.completed"
	EventTransferFailed           = "transfer.failed"
	EventProviderPaymentCompleted = "provider.payment.completed"
	EventProviderPaymentFailed    = "provider.payment.failed"
)

func NewWebhookHandler(webhookSecret string) *WebhookHandler {
	if webhookSecret == "" {
		return nil
	}
	return &WebhookHandler{
		webhookSecret: webhookSecret,
		handlers:      make(map[string]WebhookEventHandler),
	}
}

func (wh *WebhookHandler) RegisterHandler(eventType string, handler WebhookEventHandler) {
	wh.handlers[eventType] = handler
}

func (wh *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return fmt.Errorf("invalid method: %s", r.Method)
	}

	if r.Body == nil {
		http.Error(w, "Request body is required", http.StatusBadRequest)
		return fmt.Errorf("request body is nil")
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return fmt.Errorf("failed to read request body: %w", err)
	}

	signature := r.Header.Get("X-StandardBank-Signature")
	if !wh.verifySignature(body, signature) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return fmt.Errorf("invalid webhook signature")
	}

	// Parse the webhook event
	var event WebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return fmt.Errorf("failed to parse webhook event: %w", err)
	}

	handler, exists := wh.handlers[event.EventType]
	if !exists {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "acknowledged",
			"message": fmt.Sprintf("No handler for event type: %s", event.EventType),
		})
		return nil
	}

	if err := handler(event); err != nil {
		http.Error(w, fmt.Sprintf("Handler error: %v", err), http.StatusInternalServerError)
		return fmt.Errorf("handler error: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "processed",
		"eventId": event.EventID,
	})

	return nil
}

func (wh *WebhookHandler) verifySignature(payload []byte, signature string) bool {
	if wh.webhookSecret == "" {
		return false
	}

	if signature == "" {
		return false
	}

	mac := hmac.New(sha256.New, []byte(wh.webhookSecret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

func ParseWebhookEvent(data []byte) (*WebhookEvent, error) {
	var event WebhookEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return nil, fmt.Errorf("failed to parse webhook event: %w", err)
	}
	return &event, nil
}

func ValidateWebhookSignature(payload []byte, signature, secret string) bool {
	if secret == "" || signature == "" {
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
