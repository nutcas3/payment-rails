package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

const (
	EventPaymentReceived   = "payment.received"
	EventPaymentCompleted  = "payment.completed"
	EventPaymentFailed     = "payment.failed"
	EventWalletCreated     = "wallet.created"
	EventWalletTransferred = "wallet.transferred"
)

func (c *Client) HandleWebhook(payload []byte, signature string, handlers WebhookHandlers) error {
	if c.WebhookSecret != "" {
		if err := verifyWebhookSignature(payload, signature, c.WebhookSecret); err != nil {
			return fmt.Errorf("webhook signature verification failed: %w", err)
		}
	}

	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return fmt.Errorf("error unmarshalling webhook payload: %w", err)
	}

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

func verifyWebhookSignature(payload []byte, signature, secret string) error {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return fmt.Errorf("invalid webhook signature")
	}

	return nil
}

func (c *Client) RegisterWebhookURL(url string) error {
	reqBody := map[string]string{
		"webhook_url": url,
	}

	_, err := c.SendRequest("POST", "/webhook/register", reqBody)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ProcessWebhookRequest(body io.ReadCloser, signature string, handlers WebhookHandlers) error {
	payload, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading webhook request body: %w", err)
	}

	return c.HandleWebhook(payload, signature, handlers)
}
