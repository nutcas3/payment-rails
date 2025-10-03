package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebhookHandler(t *testing.T) {
	secret := "test-webhook-secret"
	handler := NewWebhookHandler(secret)

	// Register a test handler
	var receivedEvent WebhookEvent
	handler.RegisterHandler("payment.completed", func(event WebhookEvent) error {
		receivedEvent = event
		return nil
	})

	// Create test event
	event := WebhookEvent{
		EventID:      "EVT123456",
		EventType:    "payment.completed",
		Timestamp:    time.Now(),
		ResourceType: "payment",
		ResourceID:   "TXN123456",
		Data: map[string]interface{}{
			"amount":   1000.00,
			"currency": "ZAR",
		},
	}

	eventJSON, _ := json.Marshal(event)

	// Generate signature
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(eventJSON)
	signature := hex.EncodeToString(mac.Sum(nil))

	// Create HTTP request
	req := httptest.NewRequest("POST", "/webhook", bytes.NewReader(eventJSON))
	req.Header.Set("X-FNB-Signature", signature)

	// Create response recorder
	w := httptest.NewRecorder()

	// Handle webhook
	handler.HandleWebhook(w, req)

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Check that handler was called
	if receivedEvent.EventID != event.EventID {
		t.Errorf("Expected event ID %s, got %s", event.EventID, receivedEvent.EventID)
	}

	if receivedEvent.EventType != event.EventType {
		t.Errorf("Expected event type %s, got %s", event.EventType, receivedEvent.EventType)
	}
}

func TestWebhookHandlerInvalidSignature(t *testing.T) {
	secret := "test-webhook-secret"
	handler := NewWebhookHandler(secret)

	event := WebhookEvent{
		EventID:   "EVT123456",
		EventType: "payment.completed",
		Timestamp: time.Now(),
	}

	eventJSON, _ := json.Marshal(event)

	// Create HTTP request with invalid signature
	req := httptest.NewRequest("POST", "/webhook", bytes.NewReader(eventJSON))
	req.Header.Set("X-FNB-Signature", "invalid-signature")

	w := httptest.NewRecorder()
	handler.HandleWebhook(w, req)

	// Should return 401 Unauthorized
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestWebhookHandlerNoHandler(t *testing.T) {
	secret := "test-webhook-secret"
	handler := NewWebhookHandler(secret)

	event := WebhookEvent{
		EventID:   "EVT123456",
		EventType: "unknown.event",
		Timestamp: time.Now(),
	}

	eventJSON, _ := json.Marshal(event)

	// Generate signature
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(eventJSON)
	signature := hex.EncodeToString(mac.Sum(nil))

	req := httptest.NewRequest("POST", "/webhook", bytes.NewReader(eventJSON))
	req.Header.Set("X-FNB-Signature", signature)

	w := httptest.NewRecorder()
	handler.HandleWebhook(w, req)

	// Should return 200 but acknowledge no handler
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestWebhookHandlerMethodNotAllowed(t *testing.T) {
	handler := NewWebhookHandler("secret")

	req := httptest.NewRequest("GET", "/webhook", nil)
	w := httptest.NewRecorder()

	handler.HandleWebhook(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

func TestParseWebhookEvent(t *testing.T) {
	eventJSON := `{
		"eventId": "EVT123456",
		"eventType": "payment.completed",
		"timestamp": "2025-10-02T20:00:00Z",
		"resourceType": "payment",
		"resourceId": "TXN123456",
		"data": {
			"amount": 1000.00,
			"currency": "ZAR"
		}
	}`

	event, err := ParseWebhookEvent([]byte(eventJSON))
	if err != nil {
		t.Fatalf("ParseWebhookEvent failed: %v", err)
	}

	if event.EventID != "EVT123456" {
		t.Errorf("Expected event ID 'EVT123456', got %s", event.EventID)
	}

	if event.EventType != "payment.completed" {
		t.Errorf("Expected event type 'payment.completed', got %s", event.EventType)
	}

	if event.ResourceType != "payment" {
		t.Errorf("Expected resource type 'payment', got %s", event.ResourceType)
	}
}

func TestNotificationHandler(t *testing.T) {
	handler := NewNotificationHandler()

	var receivedNotification NotificationEvent
	handler.RegisterHandler("CREDIT", func(event NotificationEvent) error {
		receivedNotification = event
		return nil
	})

	notification := NotificationEvent{
		NotificationID:   "NOT123456",
		AccountNumber:    "1234567890",
		NotificationType: "CREDIT",
		Timestamp:        time.Now(),
		Amount:           500.00,
		Balance:          10000.00,
		Description:      "Payment received",
	}

	err := handler.HandleNotification(notification)
	if err != nil {
		t.Fatalf("HandleNotification failed: %v", err)
	}

	if receivedNotification.NotificationID != notification.NotificationID {
		t.Errorf("Expected notification ID %s, got %s", notification.NotificationID, receivedNotification.NotificationID)
	}
}
