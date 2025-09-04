package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

// Webhook event types
const (
	EventPaymentReceived   = "payment.received"
	EventPaymentCompleted  = "payment.completed"
	EventPaymentFailed     = "payment.failed"
	EventWalletCreated     = "wallet.created"
	EventWalletTransferred = "wallet.transferred"
)

// HandleWebhook processes incoming webhook events
func (c *Client) HandleWebhook(payload []byte, signature string, handlers WebhookHandlers) error {
	// Verify webhook signature if a secret is set
	if c.WebhookSecret != "" {
		if err := verifyWebhookSignature(payload, signature, c.WebhookSecret); err != nil {
			return fmt.Errorf("webhook signature verification failed: %w", err)
		}
	}

	// Parse webhook event
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return fmt.Errorf("error unmarshalling webhook payload: %w", err)
	}

	// Handle event based on type
	switch event.EventType {
	case EventPaymentReceived:
		if handlers.PaymentReceived != nil {
			handlers.PaymentReceived(event)
		}
	case EventPaymentCompleted:
		if handlers.PaymentCompleted != nil {
			handlers.PaymentCompleted(event)
		}
	case EventPaymentFailed:
		if handlers.PaymentFailed != nil {
			handlers.PaymentFailed(event)
		}
	case EventWalletCreated:
		if handlers.WalletCreated != nil {
			handlers.WalletCreated(event)
		}
	case EventWalletTransferred:
		if handlers.WalletTransferred != nil {
			handlers.WalletTransferred(event)
		}
	default:
		return fmt.Errorf("unknown webhook event type: %s", event.EventType)
	}

	return nil
}

// verifyWebhookSignature verifies the signature of a webhook payload
func verifyWebhookSignature(payload []byte, signature, secret string) error {
	// Create HMAC-SHA256 hash
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	// Compare signatures
	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return fmt.Errorf("invalid webhook signature")
	}

	return nil
}

// RegisterWebhookURL registers a webhook URL with SasaPay
func (c *Client) RegisterWebhookURL(url string) error {
	// Create request body
	reqBody := map[string]string{
		"webhook_url": url,
	}

	// Send request
	_, err := c.SendRequest("POST", "/webhook/register", reqBody)
	if err != nil {
		return err
	}

	return nil
}

// ProcessWebhookRequest processes a webhook request from an HTTP handler
func (c *Client) ProcessWebhookRequest(body io.ReadCloser, signature string, handlers WebhookHandlers) error {
	// Read request body
	payload, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading webhook request body: %w", err)
	}

	// Handle webhook
	return c.HandleWebhook(payload, signature, handlers)
}
