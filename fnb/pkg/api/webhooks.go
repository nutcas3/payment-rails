package api

import (
	"context"
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
	EventID       string                 `json:"eventId"`
	EventType     string                 `json:"eventType"` 
	Timestamp     time.Time              `json:"timestamp"`
	ResourceType  string                 `json:"resourceType"` 
	ResourceID    string                 `json:"resourceId"`
	Data          map[string]interface{} `json:"data"`
	Signature     string                 `json:"signature,omitempty"`
}

type NotificationEvent struct {
	NotificationID   string    `json:"notificationId"`
	AccountNumber    string    `json:"accountNumber"`
	NotificationType string    `json:"notificationType"` 
	Timestamp        time.Time `json:"timestamp"`
	TransactionID    string    `json:"transactionId,omitempty"`
	Amount           float64   `json:"amount"`
	Balance          float64   `json:"balance"`
	Description      string    `json:"description"`
	Reference        string    `json:"reference,omitempty"`
	CounterParty     string    `json:"counterParty,omitempty"`
}

type WebhookHandler struct {
	webhookSecret string
	handlers      map[string]WebhookEventHandler
}
type WebhookEventHandler func(event WebhookEvent) error

func NewWebhookHandler(webhookSecret string) *WebhookHandler {
	return &WebhookHandler{
		webhookSecret: webhookSecret,
		handlers:      make(map[string]WebhookEventHandler),
	}
}

func (wh *WebhookHandler) RegisterHandler(eventType string, handler WebhookEventHandler) {
	wh.handlers[eventType] = handler
}
func (wh *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	
	signature := r.Header.Get("X-FNB-Signature")
	if !wh.verifySignature(body, signature) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	
	var event WebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	
	handler, exists := wh.handlers[event.EventType]
	if !exists {
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "acknowledged",
			"message": fmt.Sprintf("No handler for event type: %s", event.EventType),
		})
		return
	}

	
	if err := handler(event); err != nil {
		http.Error(w, fmt.Sprintf("Handler error: %v", err), http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "processed",
		"eventId": event.EventID,
	})
}

func (wh *WebhookHandler) verifySignature(payload []byte, signature string) bool {
	if wh.webhookSecret == "" {
		return true
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

type NotificationHandler struct {
	handlers map[string]NotificationEventHandler
}
type NotificationEventHandler func(event NotificationEvent) error

func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{
		handlers: make(map[string]NotificationEventHandler),
	}
}

func (nh *NotificationHandler) RegisterHandler(notificationType string, handler NotificationEventHandler) {
	nh.handlers[notificationType] = handler
}
func (nh *NotificationHandler) HandleNotification(event NotificationEvent) error {
	handler, exists := nh.handlers[event.NotificationType]
	if !exists {
		// No handler registered, silently ignore
		return nil
	}

	return handler(event)
}

func (c *Client) SubscribeToNotifications(accountNumber string, notificationTypes []string, callbackURL string) error {
	payload := map[string]interface{}{
		"accountNumber":     accountNumber,
		"notificationTypes": notificationTypes,
		"callbackUrl":       callbackURL,
	}

	if err := c.DoRequest(context.Background(), "POST", "/api/v1/notifications/subscribe", payload, nil); err != nil {
		return fmt.Errorf("failed to subscribe to notifications: %w", err)
	}

	return nil
}

func (c *Client) UnsubscribeFromNotifications(accountNumber string) error {
	payload := map[string]string{
		"accountNumber": accountNumber,
	}

	if err := c.DoRequest(context.Background(), "POST", "/api/v1/notifications/unsubscribe", payload, nil); err != nil {
		return fmt.Errorf("failed to unsubscribe from notifications: %w", err)
	}

	return nil
}

const (
	EventPaymentCompleted     = "payment.completed"
	EventPaymentFailed        = "payment.failed"
	EventPaymentCancelled     = "payment.cancelled"
	EventCollectionCompleted  = "collection.completed"
	EventCollectionFailed     = "collection.failed"
	EventCollectionRejected   = "collection.rejected"
	EventMandateApproved      = "mandate.approved"
	EventMandateRejected      = "mandate.rejected"
	EventMandateCancelled     = "mandate.cancelled"
	EventMandateSuspended     = "mandate.suspended"
	EventDisputeSubmitted     = "dispute.submitted"
	EventDisputeResolved      = "dispute.resolved"
)

const (
	NotificationCredit        = "CREDIT"
	NotificationDebit         = "DEBIT"
	NotificationBalanceAlert  = "BALANCE_ALERT"
	NotificationPayment       = "PAYMENT"
	NotificationCollection    = "COLLECTION"
)
